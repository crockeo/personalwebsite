package schema

// Authorization used for the admin system
type Auth struct {
	Username string `db:"username"` // The user's username
	Password string `db:"password"` // The user's password
}
