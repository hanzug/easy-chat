package svc

import (
	"github.com/hanzug/easy-chat/apps/social/api/internal/config"
	"github.com/hanzug/easy-chat/apps/social/rpc/socialclient"
	"github.com/hanzug/easy-chat/apps/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	socialclient.Social
	userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Social: socialclient.NewSocial(zrpc.MustNewClient(c.SocialRpc)),
		User:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
