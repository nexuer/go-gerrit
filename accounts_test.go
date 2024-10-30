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

	reply, err := client.Accounts.QueryAccounts(context.Background(), "admin", &QueryAccountsOptions{
		Suggest: ptr.Ptr(true),
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("accounts: %v", len(reply))
}
