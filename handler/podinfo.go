package handler

// PodInfos a structure to hold all pod information
type PodInfos struct {
	Details []PodDetail `json:"pods"`
}

// PodDetail a structure to each pod information
type PodDetail struct {
	Name   string `json:"name"`
	IP     string `json:"ip"`
	Status string `json:"status"`
}

//AddItem add the pod detail into the slice.
func (podInfos *PodInfos) AddItem(podDetail PodDetail) []PodDetail {
	podInfos.Details = append(podInfos.Details, podDetail)
	return podInfos.Details
}
