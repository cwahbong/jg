package backend

type Problem struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Tagged []Tag  `json:"tagged"`
}

type Tag struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Tagged []Tag  `json:"tagged"`
}
