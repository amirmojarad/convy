package service

type GetRefreshTokenRequest struct {
	Id     uint
	UserId uint
}

type GetRefreshTokenResponse struct {
}

type SetRefreshTokenResponse struct {
	Id uint
}

type SetRefreshTokenRequest struct {
	UserId       uint
	RefreshToken string
}
