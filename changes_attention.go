package gerrit

// AttentionSetInfo entity contains details of users that are in the attention set.
//
// https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#attention-set-info
type AttentionSetInfo struct {
	// AccountInfo entity.
	Account AccountInfo `json:"account"`
	// The timestamp of the last update.
	LastUpdate Timestamp `json:"last_update"`
	// The reason of for adding or removing the user.
	Reason        string `json:"reason"`
	ReasonAccount string `json:"reason_account"`
}
