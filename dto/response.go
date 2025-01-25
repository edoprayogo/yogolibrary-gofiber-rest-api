package dto

const _StatusSucess string = "200"
const _StatusInternalServerError string = "500"

type Response[T any] struct {
	Code    string `Json:"code"`
	Message string `Json:"message"`
	Data    T      `Json:"data"`
}

func CrateResponseError(msg string) Response[string] {
	return Response[string]{
		Code:    _StatusInternalServerError,
		Message: msg,
		Data:    "",
	}
}

func CrateResponseOk[T any](data T) Response[T] {
	return Response[T]{
		Code:    _StatusSucess,
		Message: "Success",
		Data:    data,
	}
}
