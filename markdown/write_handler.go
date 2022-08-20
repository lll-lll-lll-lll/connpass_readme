package markdown

// net/httpを参考にした. https://pkg.go.dev/net/http

type WriteHandlerFunc func(content interface{}, repeat int)

func (w WriteHandlerFunc) WriteFunc(content interface{}, repeat int) {
	w(content, repeat)
}

type WriteHandler interface {
	WriteFunc(content interface{}, repeat int)
}

func (md *MarkDown) WriteHandle(content interface{}, repeat int, write WriteHandler) {
	write.WriteFunc(content, repeat)
}

func (md *MarkDown) WriteHandleFunc(content interface{}, repeat int, write func(content interface{}, repeat int)) {
	md.WriteHandle(content, repeat, WriteHandlerFunc(write))
}

// var defaultMarkDown MarkDown

// var DefaultMarkDown = &defaultMarkDown

// func WriteHandleFunc(content interface{}, repeat int, write func(content interface{}, repeat int)) {
// 	DefaultMarkDown.WriteHandleFunc(content, repeat, write)
// }

// func Handle(content interface{}, repeat int, write WriteHandler) {
// 	defaultMarkDown.WriteHandle(content, repeat, write)
// }