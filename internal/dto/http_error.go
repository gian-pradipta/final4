package dto

type HttpError struct {
	Code int    `json:"code"`
	Err  string `json:"error"`
}
