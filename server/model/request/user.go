package request

//User register structrue
type CommitMessage struct {
	Owner      string `json:"owner"`
	Expression []Item `json:"expression"`
}

type Item struct {
	Exp string `json:"exp"`
}

//ItemToString 将[]Item变为[]string
func ItemToString(e *[]Item) []string {
	exps := make([]string, len(*e))
	for index, elem := range *e {
		exps[index] = elem.Exp
	}
	return exps
}
