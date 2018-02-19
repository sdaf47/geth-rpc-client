package gethrpcclient

import (
	"testing"
)

func TestClient_NewAccount(t *testing.T) {
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

	_ = account

}

