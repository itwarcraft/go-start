package view

///////////////////////////////////////////////////////////////////////////////
// Authenticator

// Authenticator authenticates the user of a request context.
type Authenticator interface {
	// Authenticate returns the auth result in ok,
	// err is used for real errors not negative authentication
	Authenticate(request *Request, session *Session, response *Response) (ok bool, err error)
}

///////////////////////////////////////////////////////////////////////////////
// BoolAuth

// BoolAuth always returns its value at Authenticate().
// Can be used for debugging.
type BoolAuth bool

func (self BoolAuth) Authenticate(request *Request, session *Session, response *Response) (ok bool, err error) {
	return bool(self), nil
}

///////////////////////////////////////////////////////////////////////////////
// AnyAuthenticator

// AnyAuthenticator returns true if any of its authenticators returns true.
type AnyAuthenticator []Authenticator

func (self AnyAuthenticator) Authenticate(request *Request, session *Session, response *Response) (ok bool, err error) {
	for _, auth := range self {
		if ok, err = auth.Authenticate(request, session, response); ok || err != nil {
			return ok, err
		}
	}
	return false, nil
}

///////////////////////////////////////////////////////////////////////////////
// AllAuthenticators

// AllAuthenticators returns true if all of its authenticators return true.
type AllAuthenticators []Authenticator

func (self AllAuthenticators) Authenticate(request *Request, session *Session, response *Response) (ok bool, err error) {
	for _, auth := range self {
		if ok, err = auth.Authenticate(request, session, response); !ok {
			return false, err
		}
	}
	return true, nil
}
