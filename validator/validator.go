package validator

import (
	"path/filepath"
	"runtime"
	"strings"
)

type Validation interface {
	Validate() (bool, map[string]string)
}

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

func FuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	nameFull := runtime.FuncForPC(pc).Name()
	nameEnd := filepath.Ext(nameFull)
	return strings.TrimPrefix(nameEnd, ".")
}
