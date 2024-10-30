package gerrit

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
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

	AdditionalFields []AdditionalField `query:"o,omitempty"`
	Suggest          *bool             `query:"suggest,omitempty"`
}

// QueryAccounts lists accounts visible to the caller.
// The query string must be provided by the q parameter.
// The n parameter can be used to limit the returned results.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#query-accounts
func (s *AccountsService) QueryAccounts(ctx context.Context, query string, opts *QueryAccountsOptions) ([]*AccountInfo, error) {
	u := fmt.Sprintf("accounts/?q=%s", url.QueryEscape(query))
	var reply []*AccountInfo
	if _, err := s.client.InvokeByCredential(ctx, http.MethodGet, u, opts, &reply); err != nil {
		return nil, err
	}
	return reply, nil
}

// GetAccount returns an account as an AccountInfo entity.
// If account is "self" the current authenticated account will be returned.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#get-account
func (s *AccountsService) GetAccount(ctx context.Context, account string) (*AccountInfo, error) {
	u := fmt.Sprintf("accounts/%s", account)

	var reply AccountInfo
	if _, err := s.client.InvokeByCredential(ctx, http.MethodGet, u, nil, &reply); err != nil {
		return nil, err
	}
	return &reply, nil
}

type ListAccountsOptions struct {
	ListOptions

	Active           *bool
	Inactive         *bool
	AdditionalFields []AdditionalField
}

// ListAccounts lists accounts visible to the caller.
// The query string must be provided by the q parameter.
// The n parameter can be used to limit the returned results.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#query-accounts
func (s *AccountsService) ListAccounts(ctx context.Context, opts *ListAccountsOptions) ([]*AccountInfo, error) {
	qs := []string{"is:active"}
	var queryOpts *QueryAccountsOptions
	if opts != nil {
		if opts.Active != nil && *opts.Active == false {
			qs = []string{}
		}
		if opts.Inactive != nil && *opts.Inactive {
			qs = append(qs, "is:inactive")
		}
		queryOpts = &QueryAccountsOptions{
			ListOptions:      opts.ListOptions,
			AdditionalFields: opts.AdditionalFields,
		}
	}
	return s.QueryAccounts(ctx, strings.Join(qs, " OR "), queryOpts)
}

// SetActive Sets the account state to active.
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#set-active
func (s *AccountsService) SetActive(ctx context.Context, accountID string) error {
	u := fmt.Sprintf("accounts/%s/active", accountID)
	if _, err := s.client.InvokeByCredential(ctx, http.MethodPut, u, nil, nil, true); err != nil {
		return err
	}
	return nil
}

// DeleteActive Sets the account state to inactive.
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#delete-active
func (s *AccountsService) DeleteActive(ctx context.Context, accountID string) error {
	u := fmt.Sprintf("accounts/%s/active", accountID)
	if _, err := s.client.InvokeByCredential(ctx, http.MethodDelete, u, nil, nil, true); err != nil {
		return err
	}
	return nil
}
