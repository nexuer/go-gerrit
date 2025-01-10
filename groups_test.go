package gerrit

import (
	"context"
	"testing"
)

func TestGroupsService_ListGroups(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})
	reply, err := client.Groups.ListGroups(context.Background(), &ListGroupsOptions{
		AdditionalFields: []GroupAdditionalField{MEMBERS},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %v", reply)
}
