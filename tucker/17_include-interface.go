package main

// 포함된 인터페이스
// Read(), Close()
type Reader interface {
	Read() (n int, err error)
	Close() error
}

// Read(), Close()
// Write(), Close()
type Writer interface {
	Write() (n int, err error)
	Close() error
}

// Reader, Writer인터페이스의 메서드 집합을 모두 포함한 ReadWriter인터페이스
// Read(), Close()
// Write(), Close()
// Read(), Write(), Close()
type ReadWriter interface {
	Reader
	Writer
}

func main() {
	// Read(), Close() - Reader / ReadWriter
	// Write(), Close() - Writer / ReadWriter
	// Read(), Write(), Close() - Reader / Writer /ReadWriter
	// Read(), Write() - 사용불가능

}
