package gerrit

import (
	"net/http"
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
