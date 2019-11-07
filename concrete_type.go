package go_abstract_class

// Concrete type composed of AbstractClass and implementing Method3 to implements Interface interface.
type ConcreteType struct {
	AbstractClass
}

// Overriding Method2 of AbstractClass
func (a ConcreteType) Method2() string {
	return "Concrete method 2"
}

// Implements Method3
func (a ConcreteType) Method3() string {
	return "Concrete method 3"
}
