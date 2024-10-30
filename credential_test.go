package gerrit

import (
	"context"
	"os"
	"testing"
)

var testPasswordCredential = &PasswordCredential{
	Endpoint: os.Getenv("GERRIT_HOST"),
	Username: os.Getenv("GERRIT_USERNAME"),
	Password: os.Getenv("GERRIT_PASSWORD"),
}

func TestPasswordCredential(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	projects, err := client.Projects.ListProjects(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("projects: %v", projects)
}
