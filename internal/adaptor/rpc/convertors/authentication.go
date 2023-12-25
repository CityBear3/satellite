package convertors

import (
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/authentication"
	"github.com/CityBear3/satellite/internal/usecase"
	"github.com/CityBear3/satellite/pb/authentication/v1"
)

func ConvertToAuthenticationRequest(request *authPb.AuthenticateRequest) (usecase.AuthenticateRequest, error) {
	parseID, err := primitive.ParseID(request.GetId())
	if err != nil {
		return usecase.AuthenticateRequest{}, err
	}

	rowSecret := authentication.NewRawSecrets(request.GetSecret())

	return usecase.AuthenticateRequest{
		ID:      parseID,
		Secrets: rowSecret,
	}, nil
}
