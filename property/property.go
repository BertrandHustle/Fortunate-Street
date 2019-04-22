package property

type Property struct {
	Name       string
	BaseCost   int
	Cost       int
	Group      string
	Investment int
	Owner      string
}

// calculate total cost of property
// TODO: this should have a stock market multiplier
func CalcCost(prop Property) int {
	return prop.BaseCost + prop.Investment
}
