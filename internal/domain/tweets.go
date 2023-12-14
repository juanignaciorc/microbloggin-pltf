package domain

type Tweet struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Message string `json:"message" validate:"max=280"`
}
