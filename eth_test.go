package gethrpcclient

import (
	"testing"
	"fmt"
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

		fmt.Printf("%s: %d\n", account, balance.Result.BigInt())
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

func TestClient_SendTransaction3(t *testing.T) {
	client := DefaultClient

	accounts, err := client.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	if len(accounts.Result) <= 0 {
		t.Fatal("Empty acount list")
	}

	//client.SendTransaction()

}
