import {
  CallData,
  constants,
  RpcProvider,
  Contract,
  Account,
  json,
  ec,
} from "starknet";
import fs from "fs";
import * as dotenv from "dotenv";
import { getCompiledCode } from "./utils";
dotenv.config();

async function main() {
  //initialize Provider
  const provider = new RpcProvider({
    nodeUrl: process.env.RPC_ENDPOINT,
  });

  // initialize existing Argent X account
  const privateKey0 = process.env.DEPLOYER_PRIVATE_KEY ?? "";
  const accountAddress0: string = process.env.DEPLOYER_ADDRESS ?? "";
  const account0 = new Account(provider, accountAddress0, privateKey0);
  console.log("existing AX account1 connected.\n");

  // Connect the deployed Test instance in devnet
  const testAddress =
    "0x16c7d7713ded107f8e6287ba9beea4adbc96891cccbb3617dcaf3c80f0a916d"; // modify in accordance with result of script 4
  const compiledTest = json.parse(
    fs
      .readFileSync("target/dev/counter_WalletMapping.contract_class.json")
      .toString("ascii")
  );
  const myTestContract = new Contract(compiledTest.abi, testAddress, provider);
  console.log("Test Contract connected at =", myTestContract.address);

  // Interactions with the contract with call & invoke
  myTestContract.connect(account0);
  const par1 = CallData.compile({
    wallet:
      "0x07658dc41EaE0f9eDf42Eba61781E01279cA66c3B9f36B2c9fE5b218Ef306a00",
      data : "Hello World"
  });

  const res1 = await myTestContract.set_wallet_data_for(par1);

  // const tx = await myTestContract.increase_balance(
  //     CallData.compile({
  //         amount: 100,
  //     })
  // );
  // 🚨 do not work in V5.1.0
  //const bal1b = await myTestContract.call("get_balance");
  console.log("res1 =", res1);

  // await provider.waitForTransaction(tx.transaction_hash);

//   const balance = await myTestContract.get_counter({
//     parseRequest: false,
//     parseResponse: false,
//   });
//   console.log("res4 =", balance);
  // console.log("Initial balance =", bal1b.res.toString());
  // estimate fee
  // const { suggestedMaxFee: estimatedFee1 } = await account0.estimateInvokeFee({ contractAddress: testAddress, entrypoint: "increase_balance", calldata: ["10", "30"] });

  // const resu = await myTestContract.invoke("increase_balance", [10, 30]);
  // await provider.waitForTransaction(resu.transaction_hash);
  // const bal2 = await myTestContract.get_balance();
  // console.log("Initial balance =", bal2.res.toString());
  console.log("✅ Test completed.");
}
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
