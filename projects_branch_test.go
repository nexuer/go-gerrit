package gerrit

import (
	"context"
	"fmt"
	"net/url"
	"testing"
)

func TestProjectsService_ListBranches(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	branches, err := client.Projects.ListBranches(context.Background(), "All-Projects", &ListBranchesOptions{})

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("branches: %v", branches)
	for _, branch := range branches {
		fmt.Println(branch.Ref)
	}

}

func TestProjectsService_GetBranch(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
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
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	content, err := client.Projects.GetBranchContent(context.Background(), "",
		"", "")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("content: %v", content)
}
