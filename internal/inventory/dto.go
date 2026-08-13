package inventory

type CreateInventoryRequest struct {
	ProductID         uint `json:"productId" binding:"required"`
	Stock             int  `json:"stock" binding:"min=0"`
	LowStockThreshold int  `json:"lowStockThreshold" binding:"min=0"`
}

type UpdateInventoryRequest struct {
	Stock             int `json:"stock" binding:"min=0"`
	LowStockThreshold int `json:"lowStockThreshold" binding:"min=0"`
}

type InventoryResponse struct {
	ID                uint `json:"id"`
	ProductID         uint `json:"productId"`
	Stock             int  `json:"stock"`
	LowStockThreshold int  `json:"lowStockThreshold"`
}

func toInventoryResponse(
	inventory Inventory,
) InventoryResponse {
	return InventoryResponse{
		ID:                inventory.ID,
		ProductID:         inventory.ProductID,
		Stock:             inventory.Stock,
		LowStockThreshold: inventory.LowStockThreshold,
	}
}