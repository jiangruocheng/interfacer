package foo

type Closer interface {
	Close() error
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type ReadCloser interface {
	Reader
	Closer
}

func FooGo(rc ReadCloser) {
	rc.Read(nil)
	go func() {
		rc.Close()
	}()
}

func FooArg(rc ReadCloser) {
	rc.Read(nil)
	f := func(err error) {}
	f(rc.Close())
}

func FooGoWrong(rc ReadCloser) {
	go func() {
		rc.Close()
	}()
}

func FooArgWrong(rc ReadCloser) {
	f := func(err error) {}
	f(rc.Close())
}

func FooNestedWrong(rc ReadCloser) {
	f := func(rc ReadCloser) {
		rc.Close()
	}
	f(nil)
	b := make([]byte, 10)
	rc.Read(b)
}
