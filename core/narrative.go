package core

import "time"

type ReferenceType int

const (
	ReferenceAssignment ReferenceType = iota
	ReferenceMajor
)

type ReferenceResource struct {
	Reference     string        // C? Literal reference, Relative, internal or absolute URL
	ReferenceType ReferenceType // Type the reference refers to (e.g. "Patient")
	Display       string        // Text alternative for the resource
}

type Narrative struct {
	References []ReferenceResource // Store evidences
	Desc       string
	Status     bool
	Childs     []Narrative
}

type EvaluateResult struct {
	EvaluatedDate time.Time
	Narrative     Narrative
}
