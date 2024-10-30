package gerrit

import (
	"context"
	"fmt"
	"net/http"
)

// ChangesService contains Change related REST endpoints
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html
type ChangesService service

// WebLinkInfo entity describes a link to an external site.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#web-link-info
type WebLinkInfo struct {
	Name     string `json:"name"`
	Tooltip  string `json:"tooltip,omitempty"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url,omitempty"`
}

// LabelInfo entity contains information about a label on a change, always corresponding to the current patch set.
type LabelInfo struct {
	Optional bool `json:"optional,omitempty"`

	// Fields set by LABELS
	Approved     AccountInfo `json:"approved,omitempty"`
	Rejected     AccountInfo `json:"rejected,omitempty"`
	Recommended  AccountInfo `json:"recommended,omitempty"`
	Disliked     AccountInfo `json:"disliked,omitempty"`
	Blocking     bool        `json:"blocking,omitempty"`
	Value        int         `json:"value,omitempty"`
	DefaultValue int         `json:"default_value,omitempty"`

	// Fields set by DETAILED_LABELS
	All    []ApprovalInfo    `json:"all,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}

// ApprovalInfo entity contains information about an approval from a user for a label on a change.
type ApprovalInfo struct {
	AccountInfo
	Value int    `json:"value,omitempty"`
	Date  string `json:"date,omitempty"`
}

// ReviewerUpdateInfo entity contains information about updates
// to change's reviewers set.
type ReviewerUpdateInfo struct {
	Updated   Timestamp   `json:"updated"`    // Timestamp of the update.
	UpdatedBy AccountInfo `json:"updated_by"` // The account which modified state of the reviewer in question.
	Reviewer  AccountInfo `json:"reviewer"`   // The reviewer account added or removed from the change.
	State     string      `json:"state"`      // The reviewer state, one of "REVIEWER", "CC" or "REMOVED".
}

// ChangeMessageInfo entity contains information about a message attached to a change.
type ChangeMessageInfo struct {
	ID             string      `json:"id"`
	Author         AccountInfo `json:"author,omitempty"`
	Date           Timestamp   `json:"date"`
	Message        string      `json:"message"`
	Tag            string      `json:"tag,omitempty"`
	RevisionNumber int         `json:"_revision_number,omitempty"`
}

// FetchInfo entity contains information about how to fetch a patch set via a certain protocol.
type FetchInfo struct {
	URL      string            `json:"url"`
	Ref      string            `json:"ref"`
	Commands map[string]string `json:"commands,omitempty"`
}

// FileInfo entity contains information about a file in a patch set.
type FileInfo struct {
	Status        string `json:"status,omitempty"`
	Binary        bool   `json:"binary,omitempty"`
	OldPath       string `json:"old_path,omitempty"`
	LinesInserted int    `json:"lines_inserted,omitempty"`
	LinesDeleted  int    `json:"lines_deleted,omitempty"`
	SizeDelta     int    `json:"size_delta"`
	Size          int    `json:"size"`
}

// The ParentInfo entity contains information about the parent commit of a patch-set.
type ParentInfo struct {
	BranchName             string `json:"branch_name,omitempty"`
	CommitID               string `json:"commit_id,omitempty"`
	IsMergedInTargetBranch bool   `json:"is_merged_in_target_branch"`
	ChangeID               string `json:"change_id,omitempty"`
	ChangeNumber           int    `json:"change_number,omitempty"`
	PatchSetNumber         int    `json:"patch_set_number,omitempty"`
	ChangeStatus           string `json:"change_status,omitempty"`
}

// RevisionInfo entity contains information about a patch set.
type RevisionInfo struct {
	Kind              RevisionKind          `json:"kind,omitempty"`
	Draft             bool                  `json:"draft,omitempty"`
	Number            int                   `json:"_number"`
	Created           Timestamp             `json:"created"`
	Uploader          AccountInfo           `json:"uploader"`
	Ref               string                `json:"ref"`
	Fetch             map[string]FetchInfo  `json:"fetch"`
	Commit            CommitInfo            `json:"commit,omitempty"`
	Files             map[string]FileInfo   `json:"files,omitempty"`
	Actions           map[string]ActionInfo `json:"actions,omitempty"`
	Reviewed          bool                  `json:"reviewed,omitempty"`
	MessageWithFooter string                `json:"messageWithFooter,omitempty"`
	ParentsData       []ParentInfo          `json:"parents_data,omitempty"`
}

// ProblemInfo entity contains a description of a potential consistency problem with a change.
// These are not related to the code review process, but rather indicate some inconsistency in Gerrit’s database or repository metadata related to the enclosing change.
type ProblemInfo struct {
	Message string `json:"message"`
	Status  string `json:"status,omitempty"`
	Outcome string `json:"outcome,omitempty"`
}

// ChangeInfo entity contains information about a change.
type ChangeInfo struct {
	ID                     string                        `json:"id"`
	URL                    string                        `json:"url,omitempty"`
	Project                string                        `json:"project"`
	Branch                 string                        `json:"branch"`
	Topic                  string                        `json:"topic,omitempty"`
	AttentionSet           map[string]AttentionSetInfo   `json:"attention_set,omitempty"`
	Assignee               AccountInfo                   `json:"assignee,omitempty"`
	Hashtags               []string                      `json:"hashtags,omitempty"`
	ChangeID               string                        `json:"change_id"`
	Subject                string                        `json:"subject"`
	Status                 string                        `json:"status"`
	Created                Timestamp                     `json:"created"`
	Updated                Timestamp                     `json:"updated"`
	Submitted              *Timestamp                    `json:"submitted,omitempty"`
	Submitter              AccountInfo                   `json:"submitter,omitempty"`
	Starred                bool                          `json:"starred,omitempty"`
	Reviewed               bool                          `json:"reviewed,omitempty"`
	SubmitType             string                        `json:"submit_type,omitempty"`
	Mergeable              bool                          `json:"mergeable,omitempty"`
	Submittable            bool                          `json:"submittable,omitempty"`
	Insertions             int                           `json:"insertions"`
	Deletions              int                           `json:"deletions"`
	TotalCommentCount      int                           `json:"total_comment_count,omitempty"`
	UnresolvedCommentCount int                           `json:"unresolved_comment_count,omitempty"`
	Number                 int                           `json:"_number"`
	Owner                  AccountInfo                   `json:"owner"`
	Actions                map[string]ActionInfo         `json:"actions,omitempty"`
	Labels                 map[string]LabelInfo          `json:"labels,omitempty"`
	PermittedLabels        map[string][]string           `json:"permitted_labels,omitempty"`
	RemovableReviewers     []AccountInfo                 `json:"removable_reviewers,omitempty"`
	Reviewers              map[string][]AccountInfo      `json:"reviewers,omitempty"`
	PendingReviewers       map[string][]AccountInfo      `json:"pending_reviewers,omitempty"`
	ReviewerUpdates        []ReviewerUpdateInfo          `json:"reviewer_updates,omitempty"`
	Messages               []ChangeMessageInfo           `json:"messages,omitempty"`
	CurrentRevision        string                        `json:"current_revision,omitempty"`
	Revisions              map[string]RevisionInfo       `json:"revisions,omitempty"`
	MoreChanges            bool                          `json:"_more_changes,omitempty"`
	Problems               []ProblemInfo                 `json:"problems,omitempty"`
	IsPrivate              bool                          `json:"is_private,omitempty"`
	WorkInProgress         bool                          `json:"work_in_progress,omitempty"`
	HasReviewStarted       bool                          `json:"has_review_started,omitempty"`
	RevertOf               int                           `json:"revert_of,omitempty"`
	SubmissionID           string                        `json:"submission_id,omitempty"`
	CherryPickOfChange     int                           `json:"cherry_pick_of_change,omitempty"`
	CherryPickOfPatchSet   int                           `json:"cherry_pick_of_patch_set,omitempty"`
	ContainsGitConflicts   bool                          `json:"contains_git_conflicts,omitempty"`
	BaseChange             string                        `json:"base_change,omitempty"`
	SubmitRequirements     []SubmitRequirementResultInfo `json:"submit_requirements,omitempty"`
}

// SubmitRequirementExpressionInfo entity contains information about a submit requirement exppression.
//
// Docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#submit-requirement-expression-info
type SubmitRequirementExpressionInfo struct {
	Expression   string   `json:"expression,omitempty"`
	Fulfilled    bool     `json:"fulfilled"`
	Status       string   `json:"status"`
	PassingAtoms []string `json:"passing_atoms,omitempty"`
	FailingAtoms []string `json:"failing_atoms,omitempty"`
	ErrorMessage string   `json:"error_message,omitempty"`
}

// SubmitRequirementResultInfo entity describes the result of evaluating a submit requirement on a change.
//
// Docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#submit-requirement-result-info
type SubmitRequirementResultInfo struct {
	Name                           string                          `json:"name"`
	Description                    string                          `json:"description,omitempty"`
	Status                         string                          `json:"status"`
	IsLegacy                       bool                            `json:"is_legacy"`
	ApplicabilityExpressionResult  SubmitRequirementExpressionInfo `json:"applicability_expression_result,omitempty"`
	SubmittabilityExpressionResult SubmitRequirementExpressionInfo `json:"submittability_expression_result"`
	OverrideExpressionResult       SubmitRequirementExpressionInfo `json:"override_expression_result,omitempty"`
}

// QueryChangesOptions specifies the parameters to the ChangesService.QueryChanges.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-changes
type QueryChangesOptions struct {
	ListOptions

	Query *string `query:"q,omitempty"`
}

// QueryChanges lists changes visible to the caller.
// The query string must be provided by the q parameter.
// The n parameter can be used to limit the returned results.
//
// The change output is sorted by the last update time, most recently updated to oldest updated.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-changes
func (s *ChangesService) QueryChanges(ctx context.Context, opts *QueryChangesOptions) ([]*ChangeInfo, error) {
	u := fmt.Sprintf("changes/%s", "")
	var reply []*ChangeInfo
	if _, err := s.client.InvokeByCredential(ctx, http.MethodGet, u, opts, &reply); err != nil {
		return nil, err
	}
	return reply, nil
}
