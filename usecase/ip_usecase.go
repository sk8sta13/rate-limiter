package usecase

import (
	"context"

	"github.com/sk8sta13/rate-limiter/config"
	"github.com/sk8sta13/rate-limiter/internal/dto"
	"github.com/sk8sta13/rate-limiter/internal/entity"
	"github.com/sk8sta13/rate-limiter/internal/infra/database/repository"
)

type IPUseCase struct {
	Ctx        context.Context
	Repository *repository.IpRepository
	Limit      *config.Ip
	IPDB       dto.IPDB
}

func NewIPUseCase(ctx context.Context, limit *config.Ip, repository *repository.IpRepository) *IPUseCase {
	return &IPUseCase{
		Ctx:        ctx,
		Limit:      limit,
		Repository: repository,
	}
}

func (i *IPUseCase) Execute(ip *dto.IPRequest) error {
	i.Repository.GetData(i.Ctx, ip.IP, &i.IPDB)

	i.checkFirstRequest(ip)

	if i.isBloqued(ip) {
		return entity.ErrIPExceededMaxRequests
	}

	i.registerAccess(ip)

	return nil
}

func (i *IPUseCase) checkFirstRequest(ip *dto.IPRequest) {
	diff := ip.CurrentMoment - i.IPDB.FirstMoment
	if diff > int64(i.Limit.BloquedSeconds) {
		i.Repository.DelData(i.Ctx, ip.IP)
		i.IPDB = dto.IPDB{}
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

	i.Repository.SetData(i.Ctx, &ipdb)
}
