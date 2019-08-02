package core

// ReferenceType ...
type ReferenceType int

const (
	// ReferenceString ...
	ReferenceString ReferenceType = iota
	// ReferenceMajor ...
	ReferenceMajor
)

// ValueType ...
type ValueType int

const (
	// ValueBoolean ...
	ValueBoolean ValueType = iota
	// ValueInteger ...
	ValueInteger
	// ValueFloat ...
	ValueFloat
)

func (v ValueType) String() string {
	if v == 0 {
		return "boolean"
	}
	if v == 1 {
		return "int"
	}
	if v == 2 {
		return "float"
	}
	return "undefined"
}

// MarshalJSON ...
func (v ValueType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + v.String() + `"`), nil
}

// StatusType ...
type StatusType int

const (
	// StatusComplete ...
	StatusComplete StatusType = iota
	// StatusPending ...
	StatusPending
	// StatusError ...
	StatusError
)

func (v StatusType) String() string {
	if v == 0 {
		return "complete"
	}
	if v == 1 {
		return "pending"
	}
	if v == 2 {
		return "error"
	}
	return "undefined"
}

// MarshalJSON ...
func (v StatusType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + v.String() + `"`), nil
}

type MeasureCode int

const (
	// Measureinitialpopulation ...
	Measureinitialpopulation MeasureCode = iota
	// Measurenumerator ...
	Measurenumerator
	// Measurenumeratorexclusion ...
	Measurenumeratorexclusion
	// Measuredenominator ...
	Measuredenominator
	// Measuredenominatorexclusion ...
	Measuredenominatorexclusion
	// Measuredenominatorexception ...
	Measuredenominatorexception
	// Measuremeasurepopulation ...
	Measuremeasurepopulation
	// Measuremeasurepopulationexclusion ...
	Measuremeasurepopulationexclusion
	// Measuremeasureobservation...
	Measuremeasureobservation
)

func (v MeasureCode) String() string {
	if v == 0 {
		return "initial - population "
	}
	if v == 1 {
		return "numerator "
	}
	if v == 2 {
		return "numerator - exclusion "
	}
	if v == 3 {
		return "denominator "
	}
	if v == 4 {
		return "denominator - exclusion "
	}
	if v == 5 {
		return "denominator - exception "
	}
	if v == 6 {
		return "measure - population "
	}
	if v == 7 {
		return "measure - population - exclusion "
	}
	if v == 8 {
		return "measure - observation"
	}
	if v == 9 {
		return "complete"
	}
	return "undefined"
}

// MarshalJSON ...
func (v MeasureCode) MarshalJSON() ([]byte, error) {
	return []byte(`"` + v.String() + `"`), nil
}
