package gerrit

import (
	"context"
	"fmt"
	"net/http"
)

// GetCommit retrieves a commit of a project.
// The commit must be visible to the caller.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#get-commit
func (s *ProjectsService) GetCommit(ctx context.Context, projectName, commitID string) (*CommitInfo, error) {
	u := fmt.Sprintf("projects/%s/commits/%s", projectName, commitID)
	var reply CommitInfo
	if _, err := s.client.InvokeWithCredential(ctx, http.MethodGet, u, nil, &reply); err != nil {
		return nil, err
	}
	return &reply, nil
}
