package gerrit

import (
	"context"
	"fmt"
	"github.com/nexuer/utils/gitutil"
	"testing"
)

func TestProjectsService_ListTags(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	tags, err := client.Projects.ListTags(context.Background(), "All-Projects", &ListTagsOptions{})

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("tags: %v", tags)
	for _, tag := range tags {
		fmt.Println(tag.Ref, gitutil.ShortName(tag.Ref))
	}
}
