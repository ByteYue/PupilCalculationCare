package request

//User register structrue
type CommitMessage struct {
	Owner      string `json:"owner"`
	Expression []Item `json:"expression"`
}

type Item struct {
	Exp string `json:"exp"`
}
