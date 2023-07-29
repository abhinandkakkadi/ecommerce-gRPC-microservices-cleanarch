package models

type AdminLogin struct {
	Email    string
	Password string
}

type AdminDetails struct {
	ID    uint
	Name  string
	Email string
}

type AdminSignUp struct {
	Name            string
	Email           string
	Password        string
	ConfirmPassword string
}
