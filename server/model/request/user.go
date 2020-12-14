package request

//User register structrue
type Register struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type Login struct {
	Username string `json:"Username"`
	Password string `json:"Username"`
}
