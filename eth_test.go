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

	gas := big.NewInt(100000)

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

	fmt.Println("contract address", txReceiptResp.Result.ContractAddress)
}
