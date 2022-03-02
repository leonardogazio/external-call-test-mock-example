package fooserviceapi

type Mock struct{ fooClient }

func (*Mock) GetFoo() (FooBar, error) {
	return FooBar{Foo: "mocked bar response"}, nil
}
