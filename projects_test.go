package gerrit

import (
	"context"
	"testing"

	"github.com/nexuer/utils/ptr"
)

func TestProjectsService_ListProjects(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	projects, err := client.Projects.ListProjects(context.Background(), &ListProjectsOptions{
		Skip:        ptr.Ptr(0),
		Limit:       ptr.Ptr(0),
		Description: ptr.Ptr(true),
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("projects: %v", projects)
}

func TestProjectsService_GetProject(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	project, err := client.Projects.GetProject(context.Background(), "All-Users")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("projects: %v", project)

}
