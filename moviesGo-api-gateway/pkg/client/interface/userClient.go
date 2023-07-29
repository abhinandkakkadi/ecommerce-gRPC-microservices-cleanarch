package interfaces



type UserClient interface {

	SampleRequest(request string) (string,error)
	
}