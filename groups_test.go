package gerrit_test

import (
	"context"
	"testing"

	"github.com/nexuer/go-gerrit"
)

func TestGroupsService_ListGroups(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})
	reply, err := client.Groups.ListGroups(context.Background(), &gerrit.ListGroupsOptions{
		AdditionalFields: []gerrit.GroupAdditionalField{gerrit.MEMBERS},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %v", reply)
}
