package entities

type Profile struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenInformation struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type RefreshSession struct {
	ID           string `json:"id"`
	UserId       string `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
	UserAgent    string `json:"user_agent"`
	Fingerprint  string `json:"fingerprint"`
	ExpiresIn    int    `json:"expires_in"`
	CreatedAt    int    `json:"created_at"`
}
