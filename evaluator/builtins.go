package evaluator

import (
	"fmt"
	"monkey/object"
	"os"
)

var builtins = map[string]*object.Builtin{
	"exit": {
		Fn: func(args ...object.Object) object.Object {
			fmt.Println("Exiting REPL")
			os.Exit(0)
			return NULL // apparently not unreachable code
		},
	},
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Inspect()))}
			default:
				return newError("argument to `len` not supported, got %s", arg.Type())
			}
		},
	},
}
