package domain

type Profile struct {
	ID        int
	UserID    int
	Firstname string
	Lastname  string
}

func NewProfile() *Profile {
	return &Profile{}
}
