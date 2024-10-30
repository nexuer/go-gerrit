package gerrit

import (
	"context"
	"fmt"
	"net/http"
)

// AccountsService
// Gerrit API Docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#account-endpoints
type AccountsService service

// AccountInfo entity contains information about an account.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#account-info
type AccountInfo struct {
	AccountID   int    `json:"_account_id,omitempty"`
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Email       string `json:"email,omitempty"`
	Username    string `json:"username,omitempty"`

	// Avatars lists avatars of various sizes for the account.
	// This field is only populated if the avatars plugin is enabled.
	Avatars []struct {
		URL    string `json:"url,omitempty"`
		Height int    `json:"height,omitempty"`
	} `json:"avatars,omitempty"`
	MoreAccounts    bool     `json:"_more_accounts,omitempty"`
	SecondaryEmails []string `json:"secondary_emails,omitempty"`
	Status          string   `json:"status,omitempty"`
	Inactive        bool     `json:"inactive,omitempty"`
	Tags            []string `json:"tags,omitempty"`
}

// QueryAccountsOptions queries accounts visible to the caller.
type QueryAccountsOptions struct {
	ListOptions `query:",inline,omitempty"`

	AdditionalFields []string `query:"o,omitempty"`
	Suggest          *bool    `query:"suggest,omitempty"`
}

// QueryAccounts lists accounts visible to the caller.
// The query string must be provided by the q parameter.
// The n parameter can be used to limit the returned results.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#query-accounts
func (as *AccountsService) QueryAccounts(ctx context.Context, query string, opts *QueryAccountsOptions) ([]*AccountInfo, error) {
	u := fmt.Sprintf("accounts/?q=%s", query)
	var reply []*AccountInfo
	if _, err := as.client.InvokeByCredential(ctx, http.MethodGet, u, opts, &reply); err != nil {
		return nil, err
	}
	return reply, nil
}
