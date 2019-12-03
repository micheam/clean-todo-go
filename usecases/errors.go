package usecases

var _ error = (*ErrIllegalInputData)(nil)

type ErrIllegalInputData struct {
	detail string
}

func (e ErrIllegalInputData) Error() string {
	return e.detail
}

func NewErrIllegalInputData(detail string) error {
	return ErrIllegalInputData{detail: detail}
}
