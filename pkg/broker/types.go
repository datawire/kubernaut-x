package broker

type Claim struct {
	ID   string `json:",omitempty"`
	Name string `json:",omitempty"`
}

type Cluster struct {
	ID     string `json:",omitempty"`
	Config string `json:",omitempty"`
}

type ClaimList struct {
	Claims []Claim `json:",omitempty"`
}
