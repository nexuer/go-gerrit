package gerrit

import (
	"net/http"
	"strings"
)

type Credential interface {
	GetEndpoint() string
	Auth(req *http.Request) error
}

type PasswordCredential struct {
	Endpoint string `json:"endpoint" xml:"endpoint"`
	Username string `json:"username" xml:"username"`
	Password string `json:"password" xml:"password"`
}

func (p *PasswordCredential) GetEndpoint() string {
	return p.Endpoint
}

func (p *PasswordCredential) Auth(req *http.Request) error {
	req.SetBasicAuth(p.Username, p.Password)
	return nil
}

func hasAuthURL(path string) bool {
	return strings.HasPrefix(path, "a/") || strings.HasPrefix(path, "/a/")
}

func authUrl(path string) string {
	if strings.HasPrefix(path, "a/") || strings.HasPrefix(path, "/a/") {
		return path
	}
	return "/a/" + strings.TrimPrefix(path, "/")
}
