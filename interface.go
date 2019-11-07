package go_abstract_class

type Interface interface {
	Method1() string
	Method2() string
	Method3() string
}

func Function(i Interface) {
}

func Function1(i Interface) string {
	return i.Method1()
}

func Function2(i Interface) string {
	return i.Method2()
}

func Function3(i Interface) string {
	return i.Method3()
}
