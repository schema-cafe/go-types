package ui

type User struct {
	ID           string   `json:"id"`
	Name         []string `json:"name"`
	PasswordHash string   `json:"passwordHash"`
	GithubToken  string   `json:"githubToken"`
}
