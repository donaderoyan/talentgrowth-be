package profilehandler

import profile "github.com/donaderoyan/talentgrowth-be/controllers/user/profile"

type handler struct {
	service profile.Service
}

func NewHandlerProfile(service profile.Service) *handler {
	return &handler{service: service}
}
