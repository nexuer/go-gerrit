package gerrit

// AccessSectionInfo describes the access rights that are assigned on a ref.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-access.html#access-section-info
type AccessSectionInfo struct {
	Permissions map[string]PermissionInfo `json:"permissions"`
}

// PermissionInfo entity contains information about an assigned permission.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-access.html#permission-info
type PermissionInfo struct {
	Label     string                        `json:"label,omitempty"`
	Exclusive bool                          `json:"exclusive"`
	Rules     map[string]PermissionRuleInfo `json:"rules"`
}

// PermissionRuleInfo entity contains information about a permission rule that is assigned to group.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-access.html#permission-rule-info
type PermissionRuleInfo struct {
	Action PermissionAction `json:"action"`
	Force  bool             `json:"force"`
	Min    int              `json:"min"`
	Max    int              `json:"max"`
}

// ProjectAccessInfo entity contains information about the access rights for a project.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-access.html#project-access-info
type ProjectAccessInfo struct {
	Revision                     string                       `json:"revision"`
	InheritsFrom                 ProjectInfo                  `json:"inherits_from"`
	Local                        map[string]AccessSectionInfo `json:"local"`
	IsOwner                      bool                         `json:"is_owner"`
	OwnerOf                      []string                     `json:"owner_of"`
	CanUpload                    bool                         `json:"can_upload"`
	CanAdd                       bool                         `json:"can_add"`
	CanAddTags                   bool                         `json:"can_add_tags"`
	ConfigVisible                bool                         `json:"config_visible"`
	Groups                       map[string]GroupInfo         `json:"groups"`
	ConfigWebLinks               []WebLinkInfo                `json:"configWebLinks"`
	RequireChangeForConfigUpdate bool                         `json:"require_change_for_config_update"`
}
