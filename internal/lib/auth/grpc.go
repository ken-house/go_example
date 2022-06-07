package auth

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

type Authentication interface {
	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
	RequireTransportSecurity() bool
	Auth(ctx context.Context) error
}

type authentication struct {
	User     string
	Password string
}

func NewAuthentication(user, password string) Authentication {
	return &authentication{
		User:     user,
		Password: password,
	}
}

// GetRequestMetadata 返回地认证信息
func (auth *authentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"user": auth.User, "password": auth.Password}, nil
}

// RequireTransportSecurity 返回true需要证书认证
func (auth *authentication) RequireTransportSecurity() bool {
	return true
}

// Auth 进行权限判断
func (auth *authentication) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}
	var appId, appKey string
	if val, ok := md["user"]; ok {
		appId = val[0]
	}
	if val, ok := md["password"]; ok {
		appKey = val[0]
	}
	if appId != auth.User || appKey != auth.Password {
		return fmt.Errorf("invaild token")
	}
	return nil
}
