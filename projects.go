package gerrit

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ProjectsService
// Gerrit API Docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html
type ProjectsService service

type ProjectState string

const (
	Active   ProjectState = "ACTIVE"
	ReadOnly ProjectState = "READ_ONLY"
	Hidden   ProjectState = "HIDDEN"
)

type ProjectType string

const (
	All         ProjectType = "ALL"
	Code        ProjectType = "CODE"
	Permissions ProjectType = "PERMISSIONS"
)

// ProjectInfo entity contains information about a project.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#project-info
type ProjectInfo struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Parent      string            `json:"parent,omitempty"`
	Description string            `json:"description,omitempty"`
	State       ProjectState      `json:"state,omitempty"`
	Branches    map[string]string `json:"branches,omitempty"`
	WebLinks    []WebLinkInfo     `json:"web_links,omitempty"`
}

type ListProjectsOptions struct {
	// Limit the number of projects to be included in the results.
	Limit *int `query:"n,omitempty"`

	// Skip the given number of branches from the beginning of the list.
	Skip *int `query:"S,omitempty"`

	// Limit the results to the projects having the specified branch and include the sha1 of the branch in the results.
	Branch *string `query:"b,omitempty"`

	// Include project description in the results.
	Description *bool `query:"d,omitempty"`

	// Limit the results to those projects that start with the specified prefix.
	Prefix *string `query:"p,omitempty"`

	// Limit the results to those projects that match the specified regex.
	// Boundary matchers '^' and '$' are implicit.
	// For example: the regex 'test.*' will match any projects that start with 'test' and regex '.*test' will match any project that end with 'test'.
	Regex *string `query:"r,omitempty"`

	// Limit the results to those projects that match the specified substring.
	Substring *string `query:"m,omitempty"`

	// Get projects inheritance in a tree-like format.
	// This option does not work together with the branch option.
	Tree *bool `query:"t,omitempty"`

	// Get projects with specified type: ALL, CODE, PERMISSIONS.
	Type *ProjectType `query:"type,omitempty"`
}

// ListProjects gets a list of projects accessible by the authenticated user.
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#list-projects
func (ps *ProjectsService) ListProjects(ctx context.Context, opts *ListProjectsOptions) (map[string]*ProjectInfo, error) {
	var projects map[string]*ProjectInfo
	if err := ps.client.InvokeByCredential(ctx, http.MethodGet, "projects/", opts, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

// GetProject retrieves a project.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#get-project
func (ps *ProjectsService) GetProject(ctx context.Context, projectName string) (*ProjectInfo, error) {
	u := fmt.Sprintf("projects/%s", url.QueryEscape(projectName))

	var project ProjectInfo
	if err := ps.client.InvokeByCredential(ctx, http.MethodGet, u, nil, &project); err != nil {
		return nil, err
	}

	return &project, nil
}
