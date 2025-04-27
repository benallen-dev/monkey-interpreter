package object

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

type Environment struct {
	store map[string]Object
	// I'll bet you this will contain a reference to a parent scope
}

func (e *Environment) Get(name string) (Object, bool) { // _, ok pattern
	obj, ok := e.store[name] // why not return e.store[name]?
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

