package client

import (
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/client/interface"
	config "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/config"
)

type userClient struct {

}

func NewUserClient(cfg config.Config) services.UserClient {

	return &userClient{
		
	}

}

func (u *userClient) SampleRequest(request string) (string,error) {

	return "",nil
	
}


