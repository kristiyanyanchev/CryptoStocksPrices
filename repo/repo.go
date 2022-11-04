package repo

type Repo interface {
	Method() string
}

func NewRepo() Repo {
	return repo{}
}

type repo struct {
}

func (receiver repo) Method() string {
	return "Hello"
}
