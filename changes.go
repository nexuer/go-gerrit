package gerrit

// WebLinkInfo entity describes a link to an external site.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#web-link-info
type WebLinkInfo struct {
	Name     string `json:"name"`
	Tooltip  string `json:"tooltip,omitempty"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url,omitempty"`
}
