package repository

type Migrate string

const (
	Up   Migrate = "up"
	Down Migrate = "down"
)
