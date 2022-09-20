package api

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"

	"code-kata/utils/log"
)

func handleValidationError(err error) string {
	transErr := err.(validator.ValidationErrors).Translate(trans)
	log.Error(fmt.Sprintf("Failed to bind request: %s", transErr))
	errs := make([]string, 0, len(transErr))
	for _, v := range transErr {
		errs = append(errs, v)
	}
	return strings.Join(errs, "\n")
}
