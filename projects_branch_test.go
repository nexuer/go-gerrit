package gerrit

import (
	"context"
	"testing"
)

func TestProjectsService_ListBranches(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	branches, err := client.Projects.ListBranches(context.Background(), "test", &ListBranchesOptions{})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("branches: %v", branches)
}
