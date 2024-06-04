package musicalinfohandler

import musicalinfo "github.com/donaderoyan/talentgrowth-be/controllers/user/musicalinfo"

type handler struct {
	service musicalinfo.Service
}

func NewMusicalInfohandler(service musicalinfo.Service) *handler {
	return &handler{service: service}
}
