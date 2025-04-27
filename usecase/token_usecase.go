package usecase

import (
	"context"
	"fmt"

	"github.com/sk8sta13/rate-limiter/config"
	"github.com/sk8sta13/rate-limiter/internal/dto"
	"github.com/sk8sta13/rate-limiter/internal/entity"
	"github.com/sk8sta13/rate-limiter/internal/infra/database/repository"
)

type TokenUseCase struct {
	Ctx        context.Context
	Token      *config.Token
	Repository *repository.TokenRepository
	TokenDB    dto.TokenDB
}

func NewTokenUseCase(ctx context.Context, repository *repository.TokenRepository) *TokenUseCase {
	return &TokenUseCase{
		Ctx:        ctx,
		Repository: repository,
	}
}

func (t *TokenUseCase) Execute(tokenRequest *dto.TokenRequest, limtsTokens *[]config.Token) error {
	tokenValid := t.getTokenByValue(*limtsTokens, tokenRequest.Token)
	if !tokenValid {
		return entity.ErrTokenUnauthorized
	}

	dbKey := t.getKeyDb(tokenRequest.IP)

	t.Repository.GetData(t.Ctx, dbKey, &t.TokenDB)

	t.checkFirstRequest(tokenRequest)

	if t.isBloqued(tokenRequest) {
		return entity.ErrTokenExceededMaxRequest
	}

	t.registerAccess(tokenRequest)

	return nil
}

func (t *TokenUseCase) checkFirstRequest(tokenRequest *dto.TokenRequest) {
	diff := tokenRequest.CurrentMoment - t.TokenDB.FirstMoment
	if diff > int64(t.Token.BloquedSeconds) {
		t.Repository.DelData(t.Ctx, t.getKeyDb(tokenRequest.IP))
		t.TokenDB = dto.TokenDB{}
	}
}

func (t *TokenUseCase) isBloqued(tokenRequest *dto.TokenRequest) bool {
	if t.TokenDB.LastMoment == 0 && t.TokenDB.LastMoment == 0 && t.TokenDB.Qtd == 0 {
		return false
	}

	currentQtd := t.TokenDB.Qtd + 1
	diff := tokenRequest.CurrentMoment - t.TokenDB.FirstMoment

	return currentQtd > t.Token.MaxRequests && diff <= int64(t.Token.MaxRequestsInSeconds)
}

func (t *TokenUseCase) registerAccess(tokenRequest *dto.TokenRequest) {
	var tokendb dto.TokenDB
	tokendb.Key = t.getKeyDb(tokenRequest.IP)
	tokendb.Qtd = t.TokenDB.Qtd + 1
	tokendb.LastMoment = tokenRequest.CurrentMoment
	tokendb.FirstMoment = tokenRequest.CurrentMoment

	if t.TokenDB.FirstMoment != 0 {
		tokendb.FirstMoment = t.TokenDB.FirstMoment
	}

	t.Repository.SetData(t.Ctx, &tokendb)
}

func (t *TokenUseCase) getTokenByValue(tokens []config.Token, search string) bool {
	for _, tk := range tokens {
		if tk.Token == search {
			t.Token = &tk
			return true
		}
	}
	return false
}

func (t *TokenUseCase) getKeyDb(ip string) string {
	return fmt.Sprintf("%s_%s", t.Token, ip)
}
