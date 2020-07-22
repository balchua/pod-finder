package handler

// PodInfos a structure to hold all pod information
type PodInfos struct {
	Info []PodInfo
}

// PodInfo a structure to each pod information
type PodInfo struct {
	Name   string `json:"name"`
	IP     string `json:"ip"`
	Status string `json:"status"`
}
