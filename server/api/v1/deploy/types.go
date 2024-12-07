package deploy

type DeployRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required,hexadecimal"`
	GitURL        string `json:"git_url" binding:"required"`
	Type          string `json:"type"`
}

type DeployResponse struct {
	TransactionHash string `json:"transactionHash"`
}
