package entities

type SearchEntity interface {
	Search(data string) ([]string, error)
	Title() string
}
