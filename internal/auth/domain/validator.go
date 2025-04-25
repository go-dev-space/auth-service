package domain

type Validator interface {
	Struct(any) error
	Test(error) (map[string]string, error)
}
