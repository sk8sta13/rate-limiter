package dto

type IPDB struct {
	Key         string
	Qtd         int
	FirstMoment int64
	LastMoment  int64
}

type IPRequest struct {
	IP            string
	CurrentMoment int64
}
