package gerrit

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// BranchInfo entity contains information about a branch.
type BranchInfo struct {
	Ref       string        `json:"ref"`
	Revision  string        `json:"revision"`
	CanDelete bool          `json:"can_delete"`
	WebLinks  []WebLinkInfo `json:"web_links,omitempty"`
}

// ListBranchesOptions specifies the parameters to the branch API endpoints.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#branch-options
type ListBranchesOptions struct {
	// Limit the number of branches to be included in the results.
	Limit *int `query:"n,omitempty"`

	// Skip the given number of branches from the beginning of the list.
	Skip *int `query:"S,omitempty"`

	// Substring limits the results to those projects that match the specified substring.
	Substring *string `query:"m,omitempty"`

	// Limit the results to those branches that match the specified regex.
	// Boundary matchers '^' and '$' are implicit.
	// For example: the regex 't*' will match any branches that start with 'test' and regex '*t' will match any branches that end with 'test'.
	Regex *string `query:"r,omitempty"`
}

// ListBranches list the branches of a project.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#list-branches
func (ps *ProjectsService) ListBranches(ctx context.Context, projectName string, opts *ListBranchesOptions) ([]*BranchInfo, error) {
	u := fmt.Sprintf("projects/%s/branches/", url.QueryEscape(projectName))
	var branches []*BranchInfo
	if err := ps.client.InvokeByCredential(ctx, http.MethodGet, u, opts, &branches); err != nil {
		return nil, err
	}
	return branches, nil
}
