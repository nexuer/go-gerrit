package gerrit_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/nexuer/go-gerrit"
)

func TestProjectsService_ListTags(t *testing.T) {
	client := gerrit.NewClient(testPasswordCredential, &gerrit.Options{
		Debug: true,
	})

	tags, err := client.Projects.ListTags(context.Background(), "All-Projects", &gerrit.ListTagsOptions{
		//SortBy: gerrit.TagSortByCreationTime,
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("tags: %v", tags)
	for _, tag := range tags {
		fmt.Println(tag.Ref)
	}
}
