package core

import "time"

// ReferenceType ...
type ReferenceType int

const (
	// ReferenceString ...
	ReferenceString ReferenceType = iota
	// ReferenceMajor ...
	ReferenceMajor
)

// ReferenceResource ...
type ReferenceResource struct {
	Reference     string        // C? Literal reference, Relative, internal or absolute URL
	ReferenceType ReferenceType // Type the reference refers to (e.g. "Patient")
	Display       string        // Text alternative for the resource
}

// Narrative ...
type Narrative struct {
	References []ReferenceResource // Store evidences
	Desc       string
	Status     bool
	Childs     []Narrative
}

// EvaluateResult ...
type EvaluateResult struct {
	EvaluatedDate time.Time
	Narrative     Narrative
}
