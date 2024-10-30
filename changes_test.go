package gerrit

import (
	"context"
	"testing"
)

func TestChangesService_QueryChanges(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	reply, err := client.Changes.QueryChanges(context.Background(), &QueryChangesOptions{
		ListOptions: NewListOptions(0, 100),
		//Query:       ptr.Ptr("project:kernel"),
		//Suggest:     ptr.Ptr(true),
		//AdditionalFields: []AdditionalField{Details},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %v", len(reply))
}
