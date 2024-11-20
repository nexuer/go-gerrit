package gerrit

import (
	"context"
	"testing"

	"github.com/nexuer/go-gerrit/query"

	"github.com/nexuer/utils/ptr"
)

func TestChangesService_QueryChanges(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	q := query.Or(
		query.Raw("status:open"),
		query.Raw("status:merged"),
		query.Raw("status:abandoned"),
	)

	reply, err := client.Changes.QueryChanges(context.Background(), &QueryChangesOptions{
		ListOptions:      NewListOptions(0, 100),
		Query:            ptr.Ptr(q.String()),
		AdditionalFields: []AdditionalField{CurrentRevision, CurrentCommit, WebLinks},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %v", len(reply))
}
