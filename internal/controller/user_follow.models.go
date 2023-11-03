package controller

type FollowRequest struct {
	FollowerId  uint
	FollowingId uint `uri:"id"`
}

type FollowResponse struct {
}

type UnFollowRequest struct {
	FollowerId  uint
	FollowingId uint `uri:"id"`
}

type UnFollowResponse struct {
}
