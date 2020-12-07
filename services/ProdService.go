package services

import (
	"context"
)

type ProdService struct {

}

func (this *ProdService) GetProdStock(ctx context.Context, in *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 20},nil
}