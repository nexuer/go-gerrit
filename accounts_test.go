package gerrit

import (
	"context"
	"testing"

	"github.com/nexuer/utils/ptr"
)

func TestAccountsService_QueryAccounts(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	reply, err := client.Accounts.QueryAccounts(context.Background(), "is:active OR is:inactive", &QueryAccountsOptions{
		ListOptions: NewListOptions(0, 100),
		//Suggest:     ptr.Ptr(true),
		AdditionalFields: []AdditionalField{Details},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("accounts: %v", len(reply))
}

func TestAccountsService_ListAccounts(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	reply, err := client.Accounts.ListAccounts(context.Background(), &ListAccountsOptions{
		ListOptions:      NewListOptions(0, 100),
		Inactive:         ptr.Ptr(true),
		Active:           ptr.Ptr(false),
		AdditionalFields: []AdditionalField{Details},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("accounts: %v", len(reply))
}
