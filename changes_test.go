package gerrit_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/nexuer/go-gerrit"
	"github.com/nexuer/utils/ptr"
)

func TestChangesService_QueryChanges(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	q := gerrit.And(
		gerrit.Or(
			gerrit.F("status", "open"),
			gerrit.F("status", "merged"),
		),
		gerrit.F("since", gerrit.T(time.Date(2024, 11, 19, 8, 51, 36, 0, time.UTC))),
	)

	reply, err := client.Changes.QueryChanges(context.Background(), &gerrit.QueryChangesOptions{
		ListOptions: gerrit.NewListOptions(0, 100),
		Query:       ptr.Ptr(q.String()),
		AdditionalFields: []gerrit.AdditionalField{
			gerrit.CURRENT_REVISION,
			gerrit.CURRENT_COMMIT,
			gerrit.WEB_LINKS,
		},
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
