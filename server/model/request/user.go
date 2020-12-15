package request

//User register structrue
type CommitMessage struct {
	Owner      string   `json:"owner"`
	Expression []string `json:"expression"`
}
