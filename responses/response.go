package response

import "FP1hacktiv8/models"

type SuccessCreate struct {
	Status   string      `json:"status" example:"Success"`
	Messages string      `json:"messages" example:"Success Add New Todos"`
	Data     models.Todo `json:"data"`
}

type SuccessDelete struct {
	Status   string `json:"status" example:"Success"`
	Messages string `json:"messages" example:"Success Delete Todos"`
}

type SuccessUpdate struct {
	Status   string `json:"status" example:"Success"`
	Messages string `json:"messages" example:"Success update todos"`
}
