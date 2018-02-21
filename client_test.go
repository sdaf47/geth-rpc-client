package gethrpcclient

import (
	"testing"
	"math/big"
	"time"
)

var DefaultClient = NewClient("http://127.0.0.1:7545")

func TestClient_LockAccount(t *testing.T) {
	client := DefaultClient

	pass := "super_test_pass"
	accountResp, err := client.NewAccount(pass)
	if accountResp.Error != nil {
		t.Fatal(accountResp.Error)
	}
	if err != nil {
		t.Fatal(err)
	}

	lockResp, err := client.LockAccount(accountResp.Result)
	if err != nil {
		t.Fatal(err)
	}
	if lockResp.Error != nil {
		t.Fatal(lockResp.Error.Message)
	}

}

func TestClient_UnlockAccount(t *testing.T) {
	client := DefaultClient

	pass := "super_test_pass"
	accountResp, err := client.NewAccount(pass)
	if err != nil {
		t.Fatal(err)
	}
	if accountResp.Error != nil {
		t.Fatal(accountResp.Error.Message)
	}

	res, err := client.UnlockAccount(accountResp.Result, pass, 25000)

	if err != nil {
		t.Fatal(err)
	}

	if !res.Result {
		t.Fatalf("Account %s not unlocked", accountResp.Result)
	}
}

func TestClient_SendTransaction(t *testing.T) {
	client := DefaultClient

	accounts, err := client.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	bResp0, err := client.GetBalance(accounts.Result[0], LatestBlock)
	if err != nil {
		t.Fatal(err)
	}
	if bResp0.Error != nil {
		t.Fatal(bResp0.Error.Message)
	}

	bResp1, err := client.GetBalance(accounts.Result[1], LatestBlock)
	if err != nil {
		t.Fatal(err)
	}
	if bResp1.Error != nil {
		t.Fatal(bResp1.Error.Message)
	}

	value := big.NewInt(1).Mul(big.NewInt(1000000000000000), big.NewInt(5000)) // *1000
	gas := big.NewInt(21000)
	gasPrice := big.NewInt(1000000000)

	tx := &Transaction{
		From:     &accounts.Result[0],
		To:       &accounts.Result[1],
		Gas:      NewQuantity(gas),
		GasPrice: NewQuantity(gasPrice),
		Value:    NewQuantity(value),
	}

	txHashResp, err := client.SendTransaction(tx)
	if err != nil {
		t.Fatal(err)
	}
	if txHashResp.Error != nil {
		t.Fatal("txHash is nil.", txHashResp.Error.Message)
	}
	txHash := *txHashResp.Result

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

	nbResp0, err := client.GetBalance(accounts.Result[0], LatestBlock)
	if err != nil {
		t.Fatal(err)
	}
	if nbResp0.Error != nil {
		t.Fatal(nbResp0.Error.Message)
	}

	nbResp1, err := client.GetBalance(accounts.Result[1], LatestBlock)
	if err != nil {
		t.Fatal(err)
	}
	if nbResp1.Error != nil {
		t.Fatal(nbResp1.Error.Message)
	}

}

func TestClient_SendTransaction2(t *testing.T) {
	client := DefaultClient

	accounts, err := client.Accounts()

	if err != nil {
		t.Fatal(err)
	}

	value := big.NewInt(1000000)
	gas := big.NewInt(21000)
	gasPrice := big.NewInt(1000000000)

	tx := &Transaction{
		From:     &accounts.Result[0],
		To:       &accounts.Result[1],
		Gas:      NewQuantity(gas),
		GasPrice: NewQuantity(gasPrice),
		Value:    NewQuantity(value),
	}

	_, err = client.SendTransaction(tx)
	if err != nil {
		t.Fatal(err)
	}
}
