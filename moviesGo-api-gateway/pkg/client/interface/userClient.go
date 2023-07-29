package interfaces

import "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/models"

type UserClient interface {
	SampleRequest(request string) (string, error)
	SignUpRequest(userDetails models.UserDetails) (int, error)
}
