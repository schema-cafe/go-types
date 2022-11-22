package ui

type User struct {
	ID          string   `json:"id"`
	Name        []string `json:"name"`
	GithubToken string   `json:"githubToken"`
}
