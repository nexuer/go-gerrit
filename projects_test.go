package gerrit

import (
	"context"
	"testing"
)

func TestProjectsService_ListProjects(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	projects, err := client.Projects.ListProjects(context.Background(), &ListProjectsOptions{
		//Description: ptr.Ptr(true),
		ListOptions: NewListOptions(1, 50),
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("projects: %v", len(projects))
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

func TestProjectsService_CreateProject(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{Debug: true})

	project, err := client.Projects.CreateProject(context.Background(), "test-1", &CreateProjectOptions{})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("projects: %v", project)
}
