package inventory

type CreateItemStructureRequest struct {
	Name               string                   `json:"name" binding:"required"`
	Attributes         []map[string]interface{} `json:"attributes" binding:"required"`
	GeneralInformation []map[string]interface{} `json:"general_information" binding:"required"`
}

type CreateItemsRequest struct {
	StructureId string                   `json:"structure_id" binding:"required"`
	Items       []map[string]interface{} `json:"items" binding:"required"`
}