package stateful

type ExecutionContext struct {
	fc *FunctionContext
}

func NewExecutionContext(fc *FunctionContext) ExecutionContext {
	return ExecutionContext{
		fc: fc,
	}
}

// Create a new executionState that shares context, but not state
func (e ExecutionContext) Create() ExecutionState {
	return ExecutionState{
		Funcs: NewFunctions(e.fc),
	}
}

// ExecutionState is auxiliary struct for data/context that needs to be passed
// to evaluation functions
type ExecutionState struct {
	Funcs   Funcs
	context ExecutionContext
}

func (e ExecutionState) CreateFromContext() ExecutionState {
	return e.context.Create()
}

func (e ExecutionState) ResetAll() {
	// Reset the functions
	for _, f := range e.Funcs {
		f.Reset()
	}
}
