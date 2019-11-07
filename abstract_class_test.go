package go_abstract_class

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDummy(t *testing.T) {
	// Not working because AbstractClass does not implements all methods of Interface
	//assert.Equal(t, "", Method1(AbstractClass{}))

	concreteType := ConcreteType{}

	assert.Equal(t, "Abstract method 1", Method1(concreteType))
	assert.Equal(t, "Concrete method 2", Method2(concreteType))
	assert.Equal(t, "Concrete method 3", Method3(concreteType))
}
