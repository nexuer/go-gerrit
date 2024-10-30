package gerrit

import (
	"net/http"
	"strings"
)

type Credential interface {
	GetEndpoint() string
	AuthURL(path string) string
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

func (p *PasswordCredential) AuthURL(path string) string {
	if !HasAuthURL(path) {
		return "a/" + path
	}
	return path
}

func HasAuthURL(path string) bool {
	return strings.HasPrefix(path, "a/") || strings.HasPrefix(path, "/a/")
}
