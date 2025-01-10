package gerrit

import (
	"context"
	"net/http"
)

// GroupsService contains Group related REST endpoints
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html
type GroupsService service

type GroupAdditionalField string

const (
	INCLUDES GroupAdditionalField = "INCLUDES"
	MEMBERS  GroupAdditionalField = "MEMBERS"
)

// GroupOptionsInfo entity contains options of the group.
type GroupOptionsInfo struct {
	VisibleToAll bool `json:"visible_to_all,omitempty"`
}

// GroupInfo entity contains information about a group.
// This can be a Gerrit internal group, or an external group that is known to Gerrit.
type GroupInfo struct {
	ID          string           `json:"id"`
	Name        string           `json:"name,omitempty"`
	URL         string           `json:"url,omitempty"`
	Options     GroupOptionsInfo `json:"options"`
	Description string           `json:"description,omitempty"`
	GroupID     int              `json:"group_id,omitempty"`
	Owner       string           `json:"owner,omitempty"`
	OwnerID     string           `json:"owner_id,omitempty"`
	CreatedOn   *Timestamp       `json:"created_on,omitempty"`
	MoreGroups  bool             `json:"_more_groups,omitempty"`
	Members     []AccountInfo    `json:"members,omitempty"`
	Includes    []GroupInfo      `json:"includes,omitempty"`
}

// ListGroupsOptions specifies the different options for the ListGroups call.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#list-groups
type ListGroupsOptions struct {
	ListOptions

	// Group Options
	// Options fields can be obtained by adding o parameters, each option requires more lookups and slows down the query response time to the client so they are generally disabled by default.
	// Optional fields are:
	//	INCLUDES: include list of directly included groups.
	//	MEMBERS: include list of direct group members.
	AdditionalFields []GroupAdditionalField `query:"o,omitempty"`
}

// ListGroups lists the groups accessible by the caller.
// This is the same as using the ls-groups command over SSH, and accepts the same options as query parameters.
// The entries in the map are sorted by group name.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#list-groups
func (s *GroupsService) ListGroups(ctx context.Context, opts *ListGroupsOptions) (map[string]*GroupInfo, error) {
	var reply map[string]*GroupInfo
	if _, err := s.client.InvokeWithCredential(ctx, http.MethodGet, "groups/", opts, &reply); err != nil {
		return nil, err
	}
	return reply, nil
}
