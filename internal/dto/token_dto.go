package dto

type TokenDB struct {
	Key         string
	Qtd         int
	FirstMoment int64
	LastMoment  int64
}

type TokenRequest struct {
	IP            string
	Token         string
	CurrentMoment int64
}
