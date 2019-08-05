package core

import (
	"time"
)

// ReferenceResource represents the resource used in a evaluation stage
type ReferenceResource struct {
	Reference     string        // Literal reference, Relative, internal or absolute URL
	ReferenceType ReferenceType // Type the reference refers to (e.g. "Patient")
	Display       string        // Text alternative for the resource
}

// EvaluationProcess represent the process of executing a function
type EvaluationProcess struct {
	FunctionID    string // Id of the executed function
	Chapter       string
	ComposedValue ResultValue
	// NOTE: Only the leaf nodes in the evaluation tree have attr References.
	// On the other hand, only intermediated nodes have attr Childs
	References []ReferenceResource // Store evidences
	Childs     []EvaluationProcess
}

// ResultValue is basically a key-value pair
type ResultValue struct {
	ValueType ValueType
	Value     interface{}
}

// Result object represent the result of a measure
type Result struct {
	Good     ResultValue // Denominator
	VeryGood ResultValue // Numerator
	Average  ResultValue // Custom 1
	Mean     ResultValue // Custom 2
	Max      ResultValue // Custom 3
}

// EvaluationResult wrap the measure result with metadata and data on the evaluation process
type EvaluationResult struct {
	EvaluationDate      time.Time
	Status              StatusType
	MeasureID           string
	MeasureDesc         string
	Result              Result
	EvaluationProcesses EvaluationProcesses
}
