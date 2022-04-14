package postgres

type Profile struct {
	ID       int32
	Email    string
	Password string
	Balance  float32
}
