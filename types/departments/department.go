package departments

type Department struct {
	Id     int    `json:"ID,string"`
	Name   string `json:"NAME"`
	Sort   int    `json:"SORT"`
	UfHead string `json:"UF_HEAD,omitempty"`
	Parent string `json:"PARENT,omitempty"`
}
