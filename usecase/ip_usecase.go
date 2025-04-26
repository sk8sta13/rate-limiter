package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/sk8sta13/rate-limiter/config"
	"github.com/sk8sta13/rate-limiter/internal/dto"
	"github.com/sk8sta13/rate-limiter/internal/entity"
	"github.com/sk8sta13/rate-limiter/internal/infra/database"
)

type IPUseCase struct {
	Ctx   *context.Context
	Limit *config.Ip
	Db    *database.IP
	IPDB  dto.IPDB
}

func NewIPUseCase(ctx *context.Context, limit *config.Ip, db *database.IP) *IPUseCase {
	return &IPUseCase{
		Ctx:   ctx,
		Limit: limit,
		Db:    db,
	}
}

func (i *IPUseCase) Execute(ip *dto.IPRequest) error {
	i.Db.GetData(*i.Ctx, ip.IP, &i.IPDB)

	i.checkFirstRequest(ip)

	if i.isBloqued(ip) {
		log.Println("IP blocked for exceeding maximum request")
		return entity.ErrIPExceededMaxRequests
	}

	i.registerAccess(ip)

	return nil
}

func (i *IPUseCase) checkFirstRequest(ip *dto.IPRequest) {
	diff := ip.CurrentMoment - i.IPDB.FirstMoment
	log.Println(fmt.Sprintf("%s: %d > %d", ip.IP, diff, i.Limit.BloquedSeconds))
	if diff > int64(i.Limit.BloquedSeconds) {
		i.Db.DelData(*i.Ctx, ip.IP)
		i.IPDB.Qtd = 0
		i.IPDB.FirstMoment = 0
		i.IPDB.LastMoment = 0
	}
}

func (i *IPUseCase) isBloqued(ip *dto.IPRequest) bool {
	if i.IPDB.LastMoment == 0 && i.IPDB.LastMoment == 0 && i.IPDB.Qtd == 0 {
		return false
	}

	currentQtd := i.IPDB.Qtd + 1
	diff := ip.CurrentMoment - i.IPDB.FirstMoment

	return currentQtd > i.Limit.MaxRequests && diff <= int64(i.Limit.MaxRequestsInSeconds)
}

func (i *IPUseCase) registerAccess(ip *dto.IPRequest) {
	var ipdb dto.IPDB
	ipdb.Key = ip.IP
	ipdb.Qtd = i.IPDB.Qtd + 1
	ipdb.LastMoment = ip.CurrentMoment
	ipdb.FirstMoment = ip.CurrentMoment

	if i.IPDB.FirstMoment != 0 {
		ipdb.FirstMoment = i.IPDB.FirstMoment
	}

	i.Db.SetData(*i.Ctx, &ipdb)
}
