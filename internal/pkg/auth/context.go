package auth

import (
	"context"

	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
)

func GetSub(ctx context.Context) (primitive.ID, error) {
	sub, ok := ctx.Value("sub").(string)
	if !ok {
		return primitive.ID{}, apperrs.UnauthenticatedError
	}

	id, err := primitive.ParseID(sub)
	if err != nil {
		return primitive.ID{}, err
	}

	return id, nil
}
