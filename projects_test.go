package gerrit

import (
	"context"
	"fmt"
	"testing"

	"github.com/nexuer/utils/ptr"
)

func TestProjectsService_ListProjects(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	projects, err := client.Projects.ListProjects(context.Background(), &ListProjectsOptions{
		Description: ptr.Ptr(true),
		ListOptions: NewListOptions(0, 50),
		All:         ptr.Ptr(true),
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("projects: %v", len(projects))

	fmt.Println(projects["All-Projects"].State)
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

func TestProjectsService_GetHEAD(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	head, err := client.Projects.GetHEAD(context.Background(), "All-Users")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("head: %v", head)
}

func TestProjectsService_GetRepositoryStatistics(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	reply, err := client.Projects.GetRepositoryStatistics(context.Background(), "All-Users")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %v", reply)
}

func TestProjectsService_CreateProject(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{Debug: true})

	project, err := client.Projects.CreateProject(context.Background(), "test-1", &CreateProjectOptions{})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("projects: %v", project)
}

func TestProjectsService_ListAccessRights(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})
	reply, err := client.Projects.ListAccessRights(context.Background(), "All-Users")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %v", reply)
}
