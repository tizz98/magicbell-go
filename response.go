package magicbell

type baseResponse struct {
	Errors APIErrors `json:"errors"`
}

func (r baseResponse) Err() error {
	if r.Errors == nil {
		return nil
	}
	return r.Errors
}
