package util

type Namespace struct {
	APIVersion string    `json:"apiVersion"`
	Kind       string    `json:"kind"`
	Metadata   *Metadata `json:"metadata"`
}

type Metadata struct {
	Name string `json:"name"`
	UID  string `json:"uid"`
}
