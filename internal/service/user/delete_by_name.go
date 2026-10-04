package user

import (
	"context"
	"errors"
	"net/http"
	dto "studyhub/internal/dto/user"
)

func(s *userService) DeleteByName(ctx context.Context, req dto.DeleteByName) (int, error){
	//check if user exists
    checkUser, err:= s.repo.GetUserByName(ctx, req.UserName)
	if err != nil{
		return http.StatusInternalServerError, errors.New("failed to find the user due to server error")
	}

	if checkUser == nil{
        return http.StatusBadRequest, errors.New("there is no user with this name")
	}
	//delete user
    err = s.repo.DeleteByName(ctx, req.UserName)
	if err != nil{
		return http.StatusInternalServerError, errors.New("failed to find the user due to server error")
	}
	//return results
	return http.StatusOK, err
}