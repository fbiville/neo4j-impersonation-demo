package io

func MaybePanic(err error) {
	if err != nil {
		panic(err)
	}
}

func MaybePanicFn(f func() error) {
	if err := f(); err != nil {
		panic(err)
	}
}
