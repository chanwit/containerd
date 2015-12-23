package peer

type Container struct {
	Id     string
	Name   string
	Status string
}

type Node map[string][]*Container
