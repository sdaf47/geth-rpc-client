package gethrpcclient

import (
	"testing"
	"fmt"
	"math/big"
	"time"
	"os"
	"io/ioutil"
)

func TestClient_Accounts(t *testing.T) {
	client := DefaultClient

	accounts, err := client.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	if len(accounts.Result) <= 0 {
		t.Fatal("Empty acount list")
	}
}

func TestClient_GetBalance(t *testing.T) {
	client := DefaultClient

	accounts, err := client.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	if len(accounts.Result) <= 0 {
		t.Fatal("Empty acount list")
	}

	for _, account := range accounts.Result {
		balance, err := client.GetBalance(account, LatestBlock)

		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("%s: %d\n", account, balance.Result)
	}
}

func TestClient_GetCode(t *testing.T) {
	client := DefaultClient

	accounts, err := client.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	if len(accounts.Result) <= 0 {
		t.Fatal("Empty acount list")
	}

	_, err = client.GetCode(accounts.Result[0], LatestBlock)
	if err != nil {
		t.Fatal(err)
	}
}

// wait time check
func TestClient_SendTransaction3(t *testing.T) {
	client := DefaultClient

	accounts, err := client.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	value := big.NewInt(10000000000000000)
	gas := big.NewInt(21000)
	gasPrice := big.NewInt(1000000000)

	tx := &Transaction{
		From:     &accounts.Result[0],
		To:       &accounts.Result[1],
		Gas:      NewQuantity(gas),
		GasPrice: NewQuantity(gasPrice),
		Value:    NewQuantity(value),
	}

	resp, err := client.SendTransaction(tx)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Result == nil {
		if resp.Error != nil {
			t.Fatal("txHash is nil.", resp.Error.Message)
		}
		t.Fatal("txHash is nil.")
	}
	txHash := *resp.Result

	pedning := 0
	for {
		time.Sleep(5 * time.Second)
		pedning += 5

		txResp, err := client.GetTransactionByHash(txHash)
		if err != nil {
			t.Fatal(err)
		}

		if txResp.Result.BlockNumber != nil {
			break
		}
		if pedning > 100 {
			t.Fatal("Too long to wait for transaction", pedning)
		}
	}

}

func TestClient_GetTransactionReceipt(t *testing.T) {
	client := DefaultClient

	accounts, err := client.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	contractSourceFile, err := os.OpenFile("example.contract.bin", os.O_RDONLY, 777)
	if err != nil {
		t.Fatal(err)
	}

	contractSource, err := ioutil.ReadAll(contractSourceFile)
	if err != nil {
		t.Fatal(err)
	}

	gas := big.NewInt(224000)

	tx := &Transaction{
		From: &accounts.Result[0],
		Gas:  NewQuantity(gas),
		Data: Data(contractSource),
	}

	resp, err := client.SendTransaction(tx)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Error != nil {
		t.Fatal("txHash is nil.", resp.Error.Message)
	}
	txHash := *resp.Result
	fmt.Println(txHash)

	pedning := 0
	for {
		time.Sleep(5 * time.Second)
		pedning += 5

		txResp, err := client.GetTransactionByHash(txHash)
		if err != nil {
			t.Fatal(err)
		}

		if txResp.Result.BlockNumber != nil {
			break
		}
		if pedning > 100 {
			t.Fatal("Too long to wait for transaction", pedning)
		}
	}

	txReceiptResp, err := client.GetTransactionReceipt(txHash)
	if err != nil {
		t.Fatal(err)
	}
	if txReceiptResp.Error != nil {
		t.Fatal(txReceiptResp.Error.Message)
	}

}

func TestClient_GetTransactionReceipt2(t *testing.T) {
	client := DefaultClient

	accountsResp, err := client.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	contractSourceFile, err := os.OpenFile("example.contract.bin", os.O_RDONLY, 777)
	if err != nil {
		t.Fatal(err)
	}

	contractSource, err := ioutil.ReadAll(contractSourceFile)
	if err != nil {
		t.Fatal(err)
	}

	gas := big.NewInt(225000)

	tx := &Transaction{
		From: &accountsResp.Result[0],
		Gas:  NewQuantity(gas),
		Data: Data(contractSource),
	}

	resp, err := client.SendTransaction(tx)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Error != nil {
		t.Fatal("txHash is nil.", resp.Error.Message)
	}
	txHash := *resp.Result

	pedning := 0
	for {
		time.Sleep(2 * time.Second)
		pedning += 2

		txResp, err := client.GetTransactionByHash(txHash)
		if err != nil {
			t.Fatal(err)
		}

		if txResp.Result.BlockNumber != nil {
			break
		}
		if pedning > 100 {
			t.Fatal("Too long to wait for transaction", pedning)
		}
	}

	txReceiptResp, err := client.GetTransactionReceipt(txHash)
	if err != nil {
		t.Fatal(err)
	}
	if txReceiptResp.Error != nil {
		t.Fatal(txReceiptResp.Error.Message)
	}

	contractHash := txReceiptResp.Result.ContractAddress

	codeResp, err := client.GetCode(contractHash, LatestBlock)
	if err != nil {
		t.Fatal(err)
	}
	if codeResp.Error != nil {
		t.Fatal(codeResp.Error.Message)
	}

	// calculate name
	// contract test
	shaResp, err := client.Sha3("transfer(address,uint256)")
	if err != nil {
		t.Fatal(err)
	}
	if shaResp.Error != nil {
		t.Fatal(shaResp.Error.Message)
	}

	callData := Data(shaResp.Result[:10])
	callData += Data(fmt.Sprintf("%064s", accountsResp.Result[0][2:]))
	//callData += Data(fmt.Sprintf("%064s", accountsResp.Result[0][2:]))
	callData += Data(fmt.Sprintf("%064s", "6"))
	fmt.Println("callData", callData)

	tx = &Transaction{
		From: &accountsResp.Result[0],
		To:   &contractHash,
		Data: callData,
	}

	txHashResp, err := client.SendTransaction(tx)
	if err != nil {
		t.Fatal(err)
	}
	if txHashResp.Error != nil {
		t.Fatal(txHashResp.Error.Message)
	}
	txHash = *txHashResp.Result

	pedning = 0
	for {
		time.Sleep(2 * time.Second)
		pedning += 2

		txResp, err := client.GetTransactionByHash(txHash)
		if err != nil {
			t.Fatal(err)
		}

		if txResp.Result.BlockNumber != nil {
			break
		}
		if pedning > 100 {
			t.Fatal("Too long to wait for transaction", pedning)
		}
	}

	txResp, err := client.GetTransactionByHash(txHash)
	if err != nil {
		t.Fatal(err)
	}
	if txResp.Error != nil {
		t.Fatal(txResp.Error.Message)
	}

	fmt.Printf(
		"Tx: %s\nInput: %s\nData: %s\nHash: %s\n",
		txHash,
		*txResp.Result.Input,
		txResp.Result.Data,
		*txResp.Result.Hash,
	)

	logsResp, err := client.GetLogs(&Filter{
		//Address: contractHash,
	})
	if err != nil {
		t.Fatal(err)
	}
	if logsResp.Error != nil {
		t.Fatal(logsResp.Error.Message)
	}

	fmt.Println("Len of logs:", len(logsResp.Result))
	for _, log := range logsResp.Result {
		fmt.Println(log.Address, log.BlockNumber, log.TransactionHash)
	}
}
