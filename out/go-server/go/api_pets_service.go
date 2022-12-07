/*
 * Swagger Petstore
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// PetsApiService is a service that implements the logic for the PetsApiServicer
// This service should implement the business logic for every endpoint for the PetsApi API.
// Include any external packages or services that will be required by this service.
type PetsApiService struct {
}

// NewPetsApiService creates a default api service
func NewPetsApiService() PetsApiServicer {
	return &PetsApiService{}
}

// CreatePets - Create a pet
func (s *PetsApiService) CreatePets(ctx context.Context) (ImplResponse, error) {
	// TODO - update CreatePets with the required logic for this service method.
	// Add api_pets_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, {}) or use other options such as http.Ok ...
	//return Response(201, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreatePets method not implemented")
}

// ListPets - List all pets
func (s *PetsApiService) ListPets(ctx context.Context, limit int32) (ImplResponse, error) {
	// TODO - update ListPets with the required logic for this service method.
	// Add api_pets_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []Pet{}) or use other options such as http.Ok ...
	//return Response(200, []Pet{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListPets method not implemented")
}

// ShowPetById - Info for a specific pet
func (s *PetsApiService) ShowPetById(ctx context.Context, petId string) (ImplResponse, error) {
	// TODO - update ShowPetById with the required logic for this service method.
	// Add api_pets_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Pet{}) or use other options such as http.Ok ...
	//return Response(200, Pet{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ShowPetById method not implemented")
}
