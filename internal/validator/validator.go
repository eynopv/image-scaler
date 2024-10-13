package validator

type Validator struct {
	Message string
}

func NewValidator() *Validator {
	return &Validator{
		Message: "",
	}
}

func (v *Validator) Check(isValid bool, message string) {
	if !isValid {
		v.Message = message
	}
}

func (v *Validator) IsValid() bool {
	if v.Message == "" {
		return true
	}
	return false
}
