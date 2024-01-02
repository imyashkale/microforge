package dummy


type Dummy struct {
	Name   string
	URL    string
	Target string
}

type CreateDummy struct {
	Name string `json:"name"`
}
