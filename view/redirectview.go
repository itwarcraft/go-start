package view

///////////////////////////////////////////////////////////////////////////////
// RedirectView

// If rendered, this view will cause a HTTP redirect.
type RedirectView struct {
	ViewBase
	URL       string
	Permanent bool
}

func (self *RedirectView) Render(request *Request, session *Session, response *Response) (err error) {
	if self.Permanent {
		return PermanentRedirect(self.URL)
	}
	return Redirect(self.URL)
}
