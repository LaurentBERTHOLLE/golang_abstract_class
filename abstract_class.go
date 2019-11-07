package go_abstract_class

// Abstract class implementing Method1 and Method2
type AbstractClass struct {
}

// Implements Method1
func (ab AbstractClass) Method1() string {
	return "Abstract method 1"
}

// Implements Method2
func (ab AbstractClass) Method2() string {
	return "Abstract method 2"
}
