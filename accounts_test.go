package gerrit

import (
	"context"
	"fmt"
	"testing"
)

func TestAccountsService_QueryAccounts(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	q := Or(
		F("is", "active"),
		F("is", "inactive"),
	)

	reply, err := client.Accounts.QueryAccounts(context.Background(), q.String(), &QueryAccountsOptions{
		ListOptions:      NewListOptions(0, 100),
		AdditionalFields: []AccountAdditionalField{Details},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("accounts: %v", len(reply))
}

func TestAccountsService_GetAccount(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	reply, err := client.Accounts.GetAccount(context.Background(), "test")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("account: %+v", reply)
}

func TestAccountsService_SetActive(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	err := client.Accounts.SetActive(context.Background(), "1000002")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("set active ok!")

}

func TestAccountsService_DeleteActive(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	err := client.Accounts.DeleteActive(context.Background(), "1000002")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("delete active ok!")

}

func TestAccountsService_ListAccounts(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	reply, err := client.Accounts.ListAccounts(context.Background(), &ListAccountsOptions{
		ListOptions:     NewListOptions(0, 100),
		IncludeInactive: true,
		//ExcludeActive:    true,
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("accounts: %v", len(reply))
}
