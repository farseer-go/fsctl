package parse

import "strings"

const CommentPrefix = "//"

type Annotation struct {
	Cmd  string
	Args []string
}

func GetAnnotation(comment string) *Annotation {
	// 去除前缀：//
	comment = strings.TrimPrefix(comment, CommentPrefix)
	// 去除空格
	comment = strings.TrimPrefix(comment, " ")
	// 约定前缀为@标记
	if strings.HasPrefix(comment, "@") {
		comments := strings.Split(strings.TrimPrefix(comment, "@"), " ")
		return &Annotation{
			Cmd:  strings.ToLower(comments[0]),
			Args: comments[1:],
		}
	}
	return nil
}

func (receiver *Annotation) IsArea() bool {
	return receiver.Cmd == "area"
}

func (receiver *Annotation) IsFilter() bool {
	return receiver.Cmd == "filter"
}

func (receiver *Annotation) IsDi() bool {
	return receiver.Cmd == "di" && len(receiver.Args) == 2
}

func (receiver *Annotation) IsMessage() bool {
	return receiver.Cmd == "message" && len(receiver.Args) == 1
}

func (receiver *Annotation) IsApi() bool {
	return (receiver.Cmd == "get" ||
		receiver.Cmd == "post" ||
		receiver.Cmd == "put" ||
		receiver.Cmd == "delete") && len(receiver.Args) == 1
}
