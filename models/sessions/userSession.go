package sessions
type UserSession struct {
	Username string
	Email string
	Password string
	Errors []error
}