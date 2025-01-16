package gerrit_test

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/nexuer/go-gerrit"
)

func TestProjectsService_ListBranches(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	branches, err := client.Projects.ListBranches(context.Background(), "All-Projects", &gerrit.ListBranchesOptions{})

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("branches: %v", branches)
	for _, branch := range branches {
		fmt.Println(branch.Ref)
	}

}

func TestProjectsService_GetBranch(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	branch, err := client.Projects.GetBranch(context.Background(), "All-Projects",
		url.QueryEscape("refs/meta/config"))

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("branches: %v", branch)
}

func TestProjectsService_GetBranchContent(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	content, err := client.Projects.GetBranchContent(context.Background(), "",
		"", "")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("content: %v", content)
}
