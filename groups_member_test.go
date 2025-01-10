package gerrit

import (
	"context"
	"testing"
)

func TestGroupsService_ListGroupMembers(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})
	reply, err := client.Groups.ListGroupMembers(context.Background(), "Administrators", nil)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %v", reply)
}
