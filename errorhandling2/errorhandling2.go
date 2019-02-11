package main

type MyError struct {
    Inner      error
    Message    string
    StackTrace string
    Misc       map[string]interface{}
}

func wrapError(err error, messagef string, msgArgs ...interface{}) MyError {
    return MyError{
        Inner:      err, 1
        Message:    fmt.Sprintf(messagef, msgArgs...),
        StackTrace: string(debug.Stack()), 2
        Misc:       make(map[string]interface{}), 3
    }
}

func (err MyError) Error() string {
    return err.Message
}

func main() {
	
}