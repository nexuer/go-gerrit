package gerrit_test

import (
	"context"
	"testing"

	"github.com/nexuer/go-gerrit"
)

func TestProjectsService_GetCommit(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	commitID := "292acc0fc02e62807b2977120e814ab49cbcd7f0"

	reply, err := client.Projects.GetCommit(context.Background(), "All-Projects", commitID)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %+v", reply)
}
