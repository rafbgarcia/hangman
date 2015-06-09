package api

const (
    devError = iota
    userError
)

type Error struct {
    text string
    kind int
}

func NoError() Error {
    return Error{}
}

func DevError(err error) Error {
    text := ""

    if err != nil {
        text = err.Error()
    }

    return Error{
        text: text,
        kind: devError,
    }
}

func UserError(text string) Error {
    return Error{
        text: text,
        kind: userError,
    }
}

func (e *Error) Text() string {
    return e.text
}

func (e *Error) IsUserError() bool {
    return e.kind == userError
}

func (e *Error) Any() bool {
    return len(e.text) > 0
}