package gethrpcclient

import (
	"testing"
	"math/big"
	"time"
)

var DefaultClient = NewClient("http://127.0.0.1:7545")

func TestClient_LockAccount(t *testing.T) {
	client := DefaultClient

	accounts, err := client.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	count := len(accounts.Result)

	pass := "super_test_pass"
	account, err := client.NewAccount(pass)
	if err != nil {
		t.Fatal(err)
	}

	accounts, err = client.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	if count+1 != len(accounts.Result) {
		t.Fatal("Account not created")
	}

	res, err := client.LockAccount(account.Result)
	if err != nil {
		t.Fatal(err)
	}

	if !res.Result {
		t.Fatalf("Account %s not locked", account.Result)
	}

}

func TestClient_UnlockAccount(t *testing.T) {
	client := DefaultClient

	accounts, err := client.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	count := len(accounts.Result)

	pass := "super_test_pass"
	account, err := client.NewAccount(pass)
	if err != nil {
		t.Fatal(err)
	}

	accounts, err = client.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	if count+1 != len(accounts.Result) {
		t.Fatal("Account not created")
	}

	res, err := client.UnlockAccount(account.Result, pass, 25000)

	if err != nil {
		t.Fatal(err)
	}

	if !res.Result {
		t.Fatalf("Account %s not unlocked", account.Result)
	}
}

func TestClient_SendTransaction(t *testing.T) {
	client := DefaultClient

	accounts, err := client.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	txBuilder := NewTransactionBuilder()

	bResp0, err := client.GetBalance(accounts.Result[0], LatestBlock)
	if err != nil {
		t.Fatal(err)
	}

	b0 := bResp0.Result.BigInt()

	bResp1, err := client.GetBalance(accounts.Result[1], LatestBlock)
	if err != nil {
		t.Fatal(err)
	}

	b1 := bResp1.Result.BigInt()

	value := big.NewInt(1).Mul(big.NewInt(1000000000000000), big.NewInt(5000)) // *1000
	gas := big.NewInt(21000)
	gasPrice := big.NewInt(1000000000)

	tx := txBuilder.
		From(accounts.Result[0]).
		To(accounts.Result[1]).
		Gas(gas).
		GasPrice(gasPrice).
		To(accounts.Result[1]).
		Value(value).
		Build()

	_, err = client.SendTransaction(tx)
	if err != nil {
		t.Fatal(err)
	}

	// todo check transaction
	time.Sleep(time.Second * 5)

	nbResp0, err := client.GetBalance(accounts.Result[0], LatestBlock)
	if err != nil {
		t.Fatal(err)
	}

	nb0 := nbResp0.Result.BigInt()

	nbResp1, err := client.GetBalance(accounts.Result[1], LatestBlock)
	if err != nil {
		t.Fatal(err)
	}

	nb1 := nbResp1.Result.BigInt()

	if b0.Cmp(nb0) != 1 || b1.Cmp(nb1) != -1 {
		t.Fatal("Transaction fail")
	}

}

func TestClient_SendTransaction2(t *testing.T) {
	client := DefaultClient

	accounts, err := client.Accounts()

	if err != nil {
		t.Fatal(err)
	}

	txBuilder := NewTransactionBuilder()

	value := big.NewInt(1000000)
	gas := big.NewInt(21000)
	gasPrice := big.NewInt(1000000000)

	tx := txBuilder.
		From(accounts.Result[0]).
		To(accounts.Result[1]).
		Gas(gas).
		GasPrice(gasPrice).
		To(accounts.Result[1]).
		Value(value).
		Build()

	_, err = client.SendTransaction(tx)
	if err != nil {
		t.Fatal(err)
	}
}
