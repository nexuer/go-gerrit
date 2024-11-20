package gerrit

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/nexuer/utils/ptr"
)

func TestChangesService_QueryChanges(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	q := And(
		Or(
			F("status", "open"),
			F("status", "merged"),
		),
		F("since", T(time.Date(2024, 11, 19, 8, 51, 36, 0, time.UTC))),
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
	for _, change := range reply {
		if len(change.Revisions) > 0 {
			for _, r := range change.Revisions {
				fmt.Println("commit_id:",
					r.Commit.Committer.Date.Local().Format(time.RFC3339),
				)
			}
		}
	}
}
