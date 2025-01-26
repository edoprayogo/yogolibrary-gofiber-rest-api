package dto

const _StatusSucess string = "200"
const _StatusBadRequest string = "400"
const _StatusInternalServerError string = "500"

type Response[T any] struct {
	Code    string `Json:"code"`
	Message string `Json:"message"`
	Data    T      `Json:"data"`
}

func CrateResponseOk[T any](data T) Response[T] {
	return Response[T]{
		Code:    _StatusSucess,
		Message: "Success",
		Data:    data,
	}
}

func CrateResponseError(msg string) Response[string] {
	return Response[string]{
		Code:    _StatusInternalServerError,
		Message: msg,
		Data:    "",
	}
}

func CrateResponseErrorData(msg string, data map[string]string) Response[map[string]string] {
	return Response[map[string]string]{
		Code:    _StatusBadRequest,
		Message: msg,
		Data:    data,
	}
}
