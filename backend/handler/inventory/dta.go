package inventory

type CreateItemStructureRequest struct {
	Name string `json:"name" binding:"required"`
	//Typ
	Attributes map[string]interface{} `json:"attributes"`
}