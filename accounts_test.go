package gerrit_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/nexuer/go-gerrit"
)

func TestAccountsService_QueryAccounts(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	q := gerrit.Or(
		gerrit.F("is", "active"),
		gerrit.F("is", "inactive"),
	)

	reply, err := client.Accounts.QueryAccounts(context.Background(), q.String(), &gerrit.QueryAccountsOptions{
		ListOptions: gerrit.NewListOptions(0, 100),
		AdditionalFields: []gerrit.AccountAdditionalField{
			gerrit.DETAILS,
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("accounts: %v", len(reply))
}

func TestAccountsService_GetAccount(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	reply, err := client.Accounts.GetAccount(context.Background(), "test")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("account: %+v", reply)
}

func TestAccountsService_SetActive(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	err := client.Accounts.SetActive(context.Background(), "1000001")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("set active ok!")

}

func TestAccountsService_DeleteActive(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	err := client.Accounts.DeleteActive(context.Background(), "1000001")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("delete active ok!")

}

func TestAccountsService_ListAccounts(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	reply, err := client.Accounts.ListAccounts(context.Background(), &gerrit.ListAccountsOptions{
		ListOptions:     gerrit.NewListOptions(0, 100),
		IncludeInactive: true,
		//ExcludeActive:    true,
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("accounts: %v", len(reply))
}

func TestAccountsService_ListSSHKeys(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	reply, err := client.Accounts.ListSSHKeys(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("ssh keys: %v", len(reply))
}

func TestAccountsService_AddSSHKey(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	reply, err := client.Accounts.AddSSHKey(context.Background(), "ssh-rsa AAAAB3Nz...")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("ssh key: %v", reply)
}

func TestAccountsService_DeleteSSHKey(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	err := client.Accounts.DeleteSSHKey(context.Background(), 2)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("delete ssh key ok!")
}
