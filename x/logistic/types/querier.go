package types

import "strings"

// Query endpoints supported by the logistic querier
const (
	QueryDeals = "deals"
	QueryDeal  = "deal"
)

type QueryResDeals []string

// implement fmt.Stringer
func (n QueryResDeals) String() string {
	return strings.Join(n[:], "\n")
}
