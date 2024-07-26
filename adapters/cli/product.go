package cli

import (
	"fmt"
	"github.com/aka-luana/arch-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {

	result := ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with name %s has been creates with price %f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s has been enabled", res.GetID())

	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s has been disabled", res.GetID())

	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s\n",
			res.GetID(), res.GetName(), res.GetPrice(), res.GetStatus())
	}

	return result, nil
}
