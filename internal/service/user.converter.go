package service

func toUserModel(req CreateUserRequest) UserModel {
	return UserModel{
		Username:       req.Username,
		HashedPassword: req.Password,
		FirstName:      req.FirstName,
		Lastname:       req.LastName,
	}
}
