package application

type APIBody struct {
	Action    string                   `json:"action"`
	LayerName string                   `json:"layerName"`
	LayerType string                   `json:"layerType"`
	Data      []map[string]interface{} `json:"data"`
}
