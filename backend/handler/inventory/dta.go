package inventory

type CreateItemStructureRequest struct {
	Name string `json:"name" binding:"required"`
	//Typ
	Attributes map[string]interface{} `json:"attributes" binding:"required"`
}

type CreateItemsRequest struct {
	Id string `json:"id" binding:"required"`
	Items []map[string]interface{} `json:"items" binding:"required"`
}