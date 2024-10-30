package gerrit

// ActionInfo entity describes a REST API call the client can make to manipulate a resource.
// These are frequently implemented by plugins and may be discovered at runtime.
type ActionInfo struct {
	Method  string `json:"method,omitempty"`
	Label   string `json:"label,omitempty"`
	Title   string `json:"title,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}

// CommitInfo entity contains information about a commit.
type CommitInfo struct {
	Commit    string        `json:"commit,omitempty"`
	Parents   []CommitInfo  `json:"parents"`
	Author    GitPersonInfo `json:"author"`
	Committer GitPersonInfo `json:"committer"`
	Subject   string        `json:"subject"`
	Message   string        `json:"message"`
	WebLinks  []WebLinkInfo `json:"web_links,omitempty"`
}

// GitPersonInfo entity contains information about the author/committer of a commit.
type GitPersonInfo struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Date  Timestamp `json:"date"`
	TZ    int       `json:"tz"`
}
