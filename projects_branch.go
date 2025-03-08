package gerrit

import (
	"context"
	"fmt"
	"net/http"
	"time"
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
	ListOptions `query:",inline,omitempty"`

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
func (s *ProjectsService) ListBranches(ctx context.Context, projectName string, opts *ListBranchesOptions) ([]*BranchInfo, error) {
	u := fmt.Sprintf("projects/%s/branches/", projectName)
	var branches []*BranchInfo
	if _, err := s.client.InvokeWithCredential(ctx, http.MethodGet, u, opts, &branches); err != nil {
		return nil, err
	}
	return branches, nil
}

// GetBranch retrieves a branch of a project.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#get-branch
func (s *ProjectsService) GetBranch(ctx context.Context, projectName, branchID string) (*BranchInfo, error) {
	u := fmt.Sprintf("projects/%s/branches/%s", projectName, branchID)

	var reply BranchInfo
	if _, err := s.client.InvokeWithCredential(ctx, http.MethodGet, u, nil, &reply); err != nil {
		return nil, err
	}
	return &reply, nil
}

// GetBranchContent gets the content of a file from the HEAD revision of a certain branch.
// The content is returned as base64 encoded string.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#get-content
func (s *ProjectsService) GetBranchContent(ctx context.Context, projectName, branchID, fileID string) (string, error) {
	u := fmt.Sprintf("projects/%s/branches/%s/files/%s/content",
		projectName,
		branchID,
		fileID)
	var reply string
	if _, err := s.client.InvokeWithCredential(ctx, http.MethodGet, u, nil, &reply); err != nil {
		return "", err
	}
	return reply, nil
}

// ReflogEntryInfo entity describes an entry in a reflog.
type ReflogEntryInfo struct {
	OldID   string        `json:"old_id"`
	NewID   string        `json:"new_id"`
	Who     GitPersonInfo `json:"who"`
	Comment string        `json:"comment"`
}

type GetReflogOptions struct {
	// Limit the number of projects to be included in the results.
	Limit int `query:"n,omitempty"`

	FromTime time.Time `query:"-"`
	ToTime   time.Time `query:"-"`

	// The timestamp for from and to must be given as UTC in the following format: yyyyMMdd_HHmm.
	From string `query:"from,omitempty"`
	To   string `query:"to,omitempty"`
}

// GetReflog gets the reflog of a certain branch.
// The caller must be project owner.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#get-reflog
func (s *ProjectsService) GetReflog(ctx context.Context, projectID, branchID string, opts ...*GetReflogOptions) ([]*ReflogEntryInfo, error) {
	u := fmt.Sprintf("projects/%s/branches/%s/reflog", projectID, branchID)

	var args *GetReflogOptions
	if len(opts) > 0 && opts[0] != nil {
		args = opts[0]
		if !args.FromTime.IsZero() {
			args.From = args.FromTime.Format("20060102_1504")
		}

		if !args.ToTime.IsZero() {
			args.To = args.FromTime.Format("20060102_1504")
		}
	}

	var reply []*ReflogEntryInfo
	if _, err := s.client.InvokeWithCredential(ctx, http.MethodGet, u, args, &reply); err != nil {
		return nil, err
	}
	return reply, nil
}
