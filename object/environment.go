package object

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{outer: nil, store: s}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	s := make(map[string]Object)
	return &Environment{outer: outer, store: s}
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) { // _, ok pattern
	obj, ok := e.store[name] // why not return e.store[name]? OOOHH because checking for parents
	if !ok && e.outer != nil {
		return e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

