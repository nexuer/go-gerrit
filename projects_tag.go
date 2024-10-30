package gerrit

import (
	"context"
	"fmt"
	"net/http"
)

// TagInfo entity contains information about a tag.
type TagInfo struct {
	Ref      string        `json:"ref"`
	Revision string        `json:"revision"`
	Object   string        `json:"object"`
	Message  string        `json:"message"`
	Tagger   GitPersonInfo `json:"tagger"`
	Created  *Timestamp    `json:"created,omitempty"`
}

type ListTagsOptions struct {
	ListOptions `query:",inline,omitempty"`
}

// ListTags list the tags of a project.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#list-tags
func (s *ProjectsService) ListTags(ctx context.Context, projectName string, opts *ListTagsOptions) ([]*TagInfo, error) {
	u := fmt.Sprintf("projects/%s/tags/", projectName)
	var reply []*TagInfo
	if _, err := s.client.InvokeByCredential(ctx, http.MethodGet, u, opts, &reply); err != nil {
		return nil, err
	}
	return reply, nil
}