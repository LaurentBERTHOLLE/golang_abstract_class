package go_abstract_class

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMethods(t *testing.T) {
	// Not working because AbstractClass does not implement all methods of Interface
	//assert.Equal(t, "", Function(AbstractClass{}))

	assert.Equal(t, "Abstract method 1", Function1(ConcreteType{}))
	assert.Equal(t, "Concrete method 2", Function2(ConcreteType{}))
	assert.Equal(t, "Concrete method 3", Function3(ConcreteType{}))
}
