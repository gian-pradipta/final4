package errorhandler

import "final2/internal/dto"

func NewHttpError(err string, code int) dto.HttpError {
	var httpErr dto.HttpError

	httpErr.Code = code
	httpErr.Err = err

	return httpErr
}
