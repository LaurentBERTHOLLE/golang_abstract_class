package go_abstract_class

type Interface interface {
	Method1() string
	Method2() string
	Method3() string
}

func Method1(i Interface) string {
	return i.Method1()
}

func Method2(i Interface) string {
	return i.Method2()
}

func Method3(i Interface) string {
	return i.Method3()
}
