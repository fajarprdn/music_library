package utils

import (
	"errors"
	"fmt"
)

type AppError struct {
	ErrorCode string
	Err       error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Error: %s, Code: %s", e.ErrorCode, e.Err)
}

func RequiredError() error {
	return &AppError{
		ErrorCode: "X01",
		Err:       errors.New("Input can not be empty"),
	}
}

func UnauthorizedError() error {
	return &AppError{
		ErrorCode: "X04",
		Err:       errors.New("Unauthorized user"),
	}
}

func DataNotFoundError() error {
	return &AppError{
		ErrorCode: "X02",
		Err:       errors.New("Data not found"),
	}
}
