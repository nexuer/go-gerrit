package gerrit

import (
	"context"
	"net/http"
)

// ConfigService contains Config related REST endpoints
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html
type ConfigService service

// GetVersion returns the version of the Gerrit server.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html#get-version
func (s *ConfigService) GetVersion(ctx context.Context) (string, error) {
	u := "config/server/version"

	var reply string
	if _, err := s.client.InvokeByCredential(ctx, http.MethodGet, u, nil, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

// AuthInfo entity contains information about the authentication configuration of the Gerrit server.
type AuthInfo struct {
	Type                     string   `json:"type"`
	UseContributorAgreements bool     `json:"use_contributor_agreements,omitempty"`
	EditableAccountFields    []string `json:"editable_account_fields"`
	LoginURL                 string   `json:"login_url,omitempty"`
	LoginText                string   `json:"login_text,omitempty"`
	SwitchAccountURL         string   `json:"switch_account_url,omitempty"`
	RegisterURL              string   `json:"register_url,omitempty"`
	RegisterText             string   `json:"register_text,omitempty"`
	EditFullNameURL          string   `json:"edit_full_name_url,omitempty"`
	HTTPPasswordURL          string   `json:"http_password_url,omitempty"`
	IsGitBasicAuth           bool     `json:"is_git_basic_auth,omitempty"`
}

// ChangeConfigInfo entity contains information about Gerrit configuration from the change section.
type ChangeConfigInfo struct {
	AllowDrafts      bool   `json:"allow_drafts,omitempty"`
	LargeChange      int    `json:"large_change"`
	ReplyLabel       string `json:"reply_label"`
	ReplyTooltip     string `json:"reply_tooltip"`
	UpdateDelay      int    `json:"update_delay"`
	SubmitWholeTopic bool   `json:"submit_whole_topic"`
}

// DownloadSchemeInfo entity contains information about a supported download scheme and its commands.
type DownloadSchemeInfo struct {
	URL             string            `json:"url"`
	IsAuthRequired  bool              `json:"is_auth_required,omitempty"`
	IsAuthSupported bool              `json:"is_auth_supported,omitempty"`
	Commands        map[string]string `json:"commands"`
	CloneCommands   map[string]string `json:"clone_commands"`
}

// DownloadInfo entity contains information about supported download options.
type DownloadInfo struct {
	Schemes  map[string]DownloadSchemeInfo `json:"schemes"`
	Archives []string                      `json:"archives"`
}

// Info entity contains information about Gerrit configuration from the gerrit section.
type Info struct {
	AllProjectsName string `json:"all_projects_name"`
	AllUsersName    string `json:"all_users_name"`
	DocURL          string `json:"doc_url,omitempty"`
	ReportBugURL    string `json:"report_bug_url,omitempty"`
	ReportBugText   string `json:"report_bug_text,omitempty"`
}

// ReceiveInfo entity contains information about the configuration of git-receive-pack behavior on the server.
type ReceiveInfo struct {
	EnableSignedPush bool `json:"enableSignedPush,omitempty"`
}

// PluginConfigInfo entity contains information about Gerrit extensions by plugins.
type PluginConfigInfo struct {
	// HasAvatars reports whether an avatar provider is registered.
	HasAvatars bool `json:"has_avatars,omitempty"`
}

// SshdInfo entity contains information about Gerrit configuration from the sshd section.
type SshdInfo struct{}

// SuggestInfo entity contains information about Gerrit configuration from the suggest section.
type SuggestInfo struct {
	From int `json:"from"`
}

// UserConfigInfo entity contains information about Gerrit configuration from the user section.
type UserConfigInfo struct {
	AnonymousCowardName string `json:"anonymous_coward_name"`
}

// ServerInfo entity contains information about the configuration of the Gerrit server.
type ServerInfo struct {
	Auth       AuthInfo          `json:"auth"`
	Change     ChangeConfigInfo  `json:"change"`
	Download   DownloadInfo      `json:"download"`
	Gerrit     Info              `json:"gerrit"`
	Gitweb     map[string]string `json:"gitweb,omitempty"`
	Plugin     PluginConfigInfo  `json:"plugin"`
	Receive    ReceiveInfo       `json:"receive,omitempty"`
	SSHd       SshdInfo          `json:"sshd,omitempty"`
	Suggest    SuggestInfo       `json:"suggest"`
	URLAliases map[string]string `json:"url_aliases,omitempty"`
	User       UserConfigInfo    `json:"user"`
}

// GetServerInfo returns the information about the Gerrit server configuration.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html#get-info
func (s *ConfigService) GetServerInfo(ctx context.Context) (*ServerInfo, error) {
	u := "config/server/info"

	var reply ServerInfo
	if _, err := s.client.InvokeByCredential(ctx, http.MethodGet, u, nil, &reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
