package user_follow

type UnFollowResponse struct {
}
type UnFollowRequest struct {
	FollowerId  uint
	FollowingId uint
}

type FollowResponse struct {
}
type FollowRequest struct {
	FollowerId  uint
	FollowingId uint
}
