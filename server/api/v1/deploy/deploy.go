package deploy

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"

	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Shachindra/altlife/server/api/types"
	"github.com/Shachindra/altlife/server/util/pkg/httphelper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/deploy")
	{
		v1.POST("", deploy)
	}
}

func deploy(c *gin.Context) {
	var request DeployRequest
	err := c.BindJSON(&request)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}

	userWalletAddr := common.HexToAddress(request.WalletAddress)
	gitUrl := request.GitURL
	codeType := request.Type

	baseFolder := "./deployments"

	_, err = git.PlainClone(baseFolder, false, &git.CloneOptions{
		URL:      gitUrl,
		Progress: os.Stdout,
	})
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "error in cloning the repo")
		return
	}

	// zip the file, make sure it is compressed for faster speed
	var buf bytes.Buffer
	err = Compress(baseFolder, &buf)
	if err != nil {
		panic(err)
	}

	// write the compressed file to disk
	err = os.WriteFile(fmt.Sprintf("%s.%s", filepath.Base(baseFolder), "tar.gz"), buf.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("/usr/local/bin/walrus", "store", "./deployments.tar.gz")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	// fmt.Println(string(stdout))
	// Split into lines
	lines := strings.Split(string(stdout), "\n")

	// Get lines from the end
	suiObjId := getLineFromEnd(lines, 4)
	blobId := getLineFromEnd(lines, 5)

	fmt.Println(userWalletAddr)
	fmt.Println(codeType)
	fmt.Println("Sui Object Id:", suiObjId)
	fmt.Println("Blob Id:", blobId)

	os.RemoveAll("./deployments/")
	os.MkdirAll("./deployments/", os.ModePerm)
	os.Remove("deployments.tar.gz")

	status := types.ApiResponse{
		Status: 200,
		Result: suiObjId + ", " + blobId,
		// Result: userWalletAddr.String() + ", " + gitUrl + ", " + codeType,
	}
	c.JSON(200, status)
}

func getLineFromEnd(lines []string, offsetFromEnd int) string {
	if len(lines) > offsetFromEnd {
		return lines[len(lines)-1-offsetFromEnd]
	}
	return ""
}

func Compress(src string, buf io.Writer) error {
	zr := gzip.NewWriter(buf)
	defer zr.Close()
	tw := tar.NewWriter(zr)
	defer tw.Close()

	return filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, file)
		if err != nil {
			return err
		}
		header.Name = filepath.ToSlash(relPath)

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			defer data.Close()

			_, err = io.Copy(tw, data)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
