package middlewares

import (
	"context"
	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
)

func AuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpcAuth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	//TODO implement authentication
	ctx = context.WithValue(ctx, "id", token)
	return ctx, nil
}
