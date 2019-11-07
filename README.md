# Golang Abstract Class

## Introduction
If you are coming from a Java world like me, you are probably used to using abstract classes. <br/>
However, in Golang, the notion of class is quite different and creating an abstract class can be painful when you do not know how to do it.
What you actually need to do is to change your mindset and adopt golang's philosophy which favors composition over inheritance.

## Composition over inheritance
In Golang, you do not have abstract classes, you only have interfaces and types.
For a type to implements an interface, it needs to implements all its methods.
If it is not the case, you will get an error at compile time when passing it to methods requiring this specific interface.<br/>
In addition, when you call a method from a struct, if the struct does not implement it, it can delegates the call to a field implementing this method.<br/>
These two mechanisms actually allow you to create a kind of "abstract class" in go.

## Implementation
You can create an interface.
```go
type Interface interface {
	Method1() string
	Method2() string
	Method3() string
}
```
Your "abstract class" will be a normal struct in Golang.
```go
type AbstractClass struct {
}
```
It can implement some methods of the interface but it is important not to implement all of them to benefit of the compile check.
```go
// Implements Method1
func (ab AbstractClass) Method1() string {
	return "Abstract method 1"
}

// Implements Method2
func (ab AbstractClass) Method2() string {
	return "Abstract method 2"
}
```
You can then create another struct which will be composed of the "abstract class" and have to implement the missing methods of the interface.
```go
type ConcreteType struct {
	AbstractClass
}

// Implements Method3
func (a ConcreteType) Method3() string {
	return "Concrete method 3"
}
```
This way, you will not be able to pass the "abstract class" to methods requiring the interface.
```go
func Function(i Interface) string {
}

func TestCannotPassAbstractClass(t *testing.T) {
	// Not working because AbstractClass does not implements all methods of Interface
	//assert.Equal(t, "", Function(AbstractClass{}))
}
```
The calls to the methods already implemented by the "abstract class" will be delegated thanks to golang composition mechanism.
```go
func Function1(i Interface) {
	return i.Method1()
}

func TestMethod1(t *testing.T) {
	assert.Equal(t, "Abstract method 1", Function1(ConcreteType{}))
}
```
And the calls to the methods implemented by the concrete type will be called directly.
```go
func Function3(i Interface) string {
	return i.Method3()
}

func TestMethod3(t *testing.T) {
	assert.Equal(t, "Concrete method 3", Function3(ConcreteType{}))
}
```
It means also that if the concrete type implements some methods of the interface the "abstract class" already implemented, it then "overrides" them.
```go
func Function2(i Interface) string {
	return i.Method2()
}

func TestMethod2(t *testing.T) {
	assert.Equal(t, "Concrete method 2", Function2(ConcreteType{}))
}
```