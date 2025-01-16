package gerrit_test

import (
	"context"
	"os"
	"testing"

	"github.com/nexuer/go-gerrit"
)

var testPasswordCredential = &gerrit.PasswordCredential{
	Endpoint: os.Getenv("GERRIT_HOST"),
	Username: os.Getenv("GERRIT_USERNAME"),
	Password: os.Getenv("GERRIT_PASSWORD"),
}

func TestPasswordCredential(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	projects, err := client.Projects.ListProjects(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("projects: %v", projects)
}
