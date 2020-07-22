package handler

import (
	"encoding/json"
	"io/ioutil"

	corev1 "k8s.io/api/core/v1"
)

//WriteToFile writes the json result to file
func WriteToFile(pods *corev1.PodList, outputPath string) {

	podInfos := PodInfos{make([]PodDetail, 0)}

	for _, pod := range pods.Items {
		podInfo := PodDetail{
			IP:     pod.Status.PodIP,
			Name:   pod.Name,
			Status: string(pod.Status.Phase),
		}
		podInfos.AddItem(podInfo)

	}

	file, _ := json.MarshalIndent(podInfos, "", "  ")
	_ = ioutil.WriteFile(outputPath, file, 0644)

}
