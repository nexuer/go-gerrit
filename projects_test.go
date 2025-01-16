package gerrit_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/nexuer/go-gerrit"

	"github.com/nexuer/utils/ptr"
)

func TestProjectsService_ListProjects(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	projects, err := client.Projects.ListProjects(context.Background(), &gerrit.ListProjectsOptions{
		Description: ptr.Ptr(true),
		ListOptions: gerrit.NewListOptions(0, 50),
		All:         ptr.Ptr(true),
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("projects: %v", len(projects))

	fmt.Println(projects["All-Projects"].State)
}

func TestProjectsService_GetProject(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	project, err := client.Projects.GetProject(context.Background(), "All-Users")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("projects: %v", project)

}

func TestProjectsService_GetHEAD(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	head, err := client.Projects.GetHEAD(context.Background(), "All-Users")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("head: %v", head)
}

func TestProjectsService_GetRepositoryStatistics(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	reply, err := client.Projects.GetRepositoryStatistics(context.Background(), "All-Users")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %v", reply)
}

func TestProjectsService_CreateProject(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{Debug: true})

	project, err := client.Projects.CreateProject(context.Background(), "test-1", &gerrit.CreateProjectOptions{})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("projects: %v", project)
}

func TestProjectsService_ListAccessRights(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})
	reply, err := client.Projects.ListAccessRights(context.Background(), "All-Projects")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("reply: %v", reply)
}
