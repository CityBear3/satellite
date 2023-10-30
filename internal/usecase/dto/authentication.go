package dto

import "github.com/CityBear3/satellite/internal/domain/primitive"

type AuthenticateRequest struct {
	ID     primitive.ID
	Secret primitive.Secret
}

func NewAuthenticateRequest(id string, secret string) (AuthenticateRequest, error) {
	parseID, err := primitive.ParseID(id)
	if err != nil {
		return AuthenticateRequest{}, err
	}

	rowSecret := primitive.NewRowSecret(secret)

	return AuthenticateRequest{
		ID:     parseID,
		Secret: rowSecret,
	}, nil
}
