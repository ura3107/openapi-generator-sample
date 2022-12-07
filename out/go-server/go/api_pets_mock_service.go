package openapi

import (
	"context"
	"net/http"
)

type PetsApiMockService struct {
}

func NewPetsApiMockService() PetsApiServicer {
	return &PetsApiMockService{}
}

func (p *PetsApiMockService) CreatePets(context.Context) (ImplResponse, error) {
	return Response(http.StatusOK, nil), nil
}
func (p *PetsApiMockService) ListPets(context.Context, int32) (ImplResponse, error) {
	return Response(http.StatusOK, nil), nil
}

func (p *PetsApiMockService) ShowPetById(context.Context, string) (ImplResponse, error) {
	return Response(http.StatusOK, nil), nil
}
