package params

type CreateUser struct {
	Name       string
	Email      string
	Passw      string
	Permission UserLevel
}

type UserLevel struct {
	ID int64
}
