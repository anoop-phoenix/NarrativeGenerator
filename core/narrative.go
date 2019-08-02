package core

import "time"

// ReferenceResource ...
type ReferenceResource struct {
	Reference     string        // C? Literal reference, Relative, internal or absolute URL
	ReferenceType ReferenceType // Type the reference refers to (e.g. "Patient")
	Display       string        // Text alternative for the resource
}

// EvaluationProcess ...
type EvaluationProcess struct {
	FunctionID string // Id of the executed function
	Chapter    string
	References []ReferenceResource // Store evidences
	ValueType  ValueType
	Value      interface{}
	Childs     []EvaluationProcess
}

// Result ...
type Result struct {
	FunctionID          string      // Id of the executed function
	Code                MeasureCode // Meaning of the measure
	ValueType           ValueType
	Value               interface{}
	EvaluationProcesses []EvaluationProcess
}

// [] remove list in result root level
// 	[] list -> object
// 	[] items -> object attr
// [] narrative need adapt continous evaluation. i.e: numerator depends on denominator

// EvaluationResult ...
type EvaluationResult struct {
	EvaluatedDate time.Time
	Status        StatusType
	Measure       string
	Results       []Result
}
