package gerrit_test

import (
	"context"
	"testing"

	"github.com/nexuer/go-gerrit"
)

func TestGroupsService_ListGroupMembers(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})
	reply, err := client.Groups.ListGroupMembers(context.Background(), "Administrators", nil)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %v", reply)
}
