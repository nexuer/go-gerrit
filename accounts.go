package gerrit

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
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

	AdditionalFields []AccountAdditionalField `query:"o,omitempty"`
	Suggest          *bool                    `query:"suggest,omitempty"`
}

// QueryAccounts lists accounts visible to the caller.
// The query string must be provided by the q parameter.
// The n parameter can be used to limit the returned results.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#query-accounts
func (s *AccountsService) QueryAccounts(ctx context.Context, query string, opts *QueryAccountsOptions) ([]*AccountInfo, error) {
	u := fmt.Sprintf("accounts/?q=%s", url.QueryEscape(query))
	var reply []*AccountInfo
	if _, err := s.client.InvokeWithCredential(ctx, http.MethodGet, u, opts, &reply); err != nil {
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
	if _, err := s.client.InvokeWithCredential(ctx, http.MethodGet, u, nil, &reply); err != nil {
		return nil, err
	}
	return &reply, nil
}

type AccountAdditionalField string

const (
	DETAILS    AccountAdditionalField = "DETAILS"
	ALL_EMAILS AccountAdditionalField = "ALL_EMAILS"
)

type ListAccountsOptions struct {
	ListOptions

	ExcludeActive   bool
	IncludeInactive bool
}

// ListAccounts lists accounts visible to the caller.
// The query string must be provided by the q parameter.
// The n parameter can be used to limit the returned results.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#query-accounts
func (s *AccountsService) ListAccounts(ctx context.Context, opts *ListAccountsOptions) ([]*AccountInfo, error) {
	f := []Query{
		Raw("is:active"),
	}
	var queryOpts *QueryAccountsOptions
	if opts != nil {
		if opts.ExcludeActive {
			f = []Query{}
		}
		if opts.IncludeInactive {
			f = append(f, Raw("is:inactive"))
		}
		queryOpts = &QueryAccountsOptions{
			ListOptions:      opts.ListOptions,
			AdditionalFields: []AccountAdditionalField{ALL_EMAILS, DETAILS},
		}
	}
	return s.QueryAccounts(ctx, Or(f...).String(), queryOpts)
}

// SetActive Sets the account state to active.
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#set-active
func (s *AccountsService) SetActive(ctx context.Context, accountID string) error {
	u := fmt.Sprintf("accounts/%s/active", accountID)
	if _, err := s.client.InvokeWithCredential(ctx, http.MethodPut, u, nil, nil, DelContentType()); err != nil {
		return err
	}
	return nil
}

// DeleteActive Sets the account state to inactive.
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#delete-active
func (s *AccountsService) DeleteActive(ctx context.Context, accountID string) error {
	u := fmt.Sprintf("accounts/%s/active", accountID)
	if _, err := s.client.InvokeWithCredential(ctx, http.MethodDelete, u, nil, nil, DelContentType()); err != nil {
		return err
	}
	return nil
}

// SSHKeyInfo entity contains information about an SSH key of a user.
type SSHKeyInfo struct {
	Seq          int    `json:"seq"`
	SSHPublicKey string `json:"ssh_public_key"`
	EncodedKey   string `json:"encoded_key"`
	Algorithm    string `json:"algorithm"`
	Comment      string `json:"comment,omitempty"`
	Valid        bool   `json:"valid"`
}

// ListSSHKeys Returns the SSH keys of an account.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#list-ssh-keys
func (s *AccountsService) ListSSHKeys(ctx context.Context) ([]*SSHKeyInfo, error) {
	var reply []*SSHKeyInfo
	if _, err := s.client.InvokeWithCredential(ctx, http.MethodGet, "accounts/self/sshkeys", nil, &reply); err != nil {
		return nil, err
	}
	return reply, nil
}

// AddSSHKey Adds an SSH key for self user.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#list-ssh-keys
func (s *AccountsService) AddSSHKey(ctx context.Context, key string) (*SSHKeyInfo, error) {
	var reply SSHKeyInfo
	if _, err := s.client.InvokeWithCredential(ctx, http.MethodPost, "accounts/self/sshkeys", nil, &reply,
		PlainText(key)); err != nil {
		return nil, err
	}
	return &reply, nil
}

// DeleteSSHKey Deletes an SSH key of self user.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#list-ssh-keys
func (s *AccountsService) DeleteSSHKey(ctx context.Context, id int) error {
	u := fmt.Sprintf("accounts/self/sshkeys/%d", id)
	if _, err := s.client.InvokeWithCredential(ctx, http.MethodDelete, u, nil, nil, DelContentType()); err != nil {
		return err
	}
	return nil
}
