package application

type APIBody struct {
	Action     string                 `json:"action"`
	EntityType string                 `json:"entityType"`
	LayerType  string                 `json:"layerType"`
	Data       map[string]interface{} `json:"data"`
}
