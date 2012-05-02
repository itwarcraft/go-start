package view

import (
	"encoding/base64"
	"strings"
)

// NewBasicAuth creates a BasicAuth instance with a single username and password.
func NewBasicAuth(realm string, username string, password string) *BasicAuth {
	return &BasicAuth{
		Realm:        realm,
		UserPassword: map[string]string{username: password},
	}
}

///////////////////////////////////////////////////////////////////////////////
// BasicAuth

// BasicAuth implements HTTP basic auth as Authenticator.
type BasicAuth struct {
	Realm        string
	UserPassword map[string]string
}

func (self *BasicAuth) Authenticate(request *Request, session *Session, response *Response) (ok bool, err error) {
	header := request.Header().Get("Authorization")
	f := strings.Fields(header)
	if len(f) == 2 && f[0] == "Basic" {
		if b, err := base64.StdEncoding.DecodeString(f[1]); err == nil {
			a := strings.Split(string(b), ":")
			if len(a) == 2 {
				username := a[0]
				password := a[1]
				p, ok := self.UserPassword[username]
				if ok && p == password {
					return true, nil
				}
			}
		}
	}

	response.SetHeader("WWW-Authenticate", "Basic realm=\""+self.Realm+"\"", false)
	response.AuthorizationRequired401()
	return false, nil
}
