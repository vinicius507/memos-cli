package cli

type cmdErrorMsg struct{ err error }

func (e cmdErrorMsg) Error() string {
	return e.err.Error()
}
