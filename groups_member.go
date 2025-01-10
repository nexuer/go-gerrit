package gerrit

import (
	"context"
	"fmt"
	"net/http"
)

// ListGroupMembersOptions specifies the different options for the ListGroupMembers call.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#group-members
type ListGroupMembersOptions struct {
	// To resolve the included groups of a group recursively and to list all members the parameter recursive can be set.
	// Members from included external groups and from included groups which are not visible to the calling user are ignored.
	Recursive bool `query:"recursive,omitempty"`
}

// ListGroupMembers lists the direct members of a Gerrit internal group.
// The entries in the list are sorted by full name, preferred email and id.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#group-members
func (s *GroupsService) ListGroupMembers(ctx context.Context, groupID string, opts *ListGroupMembersOptions) ([]*AccountInfo, error) {
	u := fmt.Sprintf("groups/%s/members/", groupID)

	var reply []*AccountInfo
	if _, err := s.client.InvokeWithCredential(ctx, http.MethodGet, u, opts, &reply); err != nil {
		return nil, err
	}

	return reply, nil
}
