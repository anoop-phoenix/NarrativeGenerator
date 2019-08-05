package core

import (
	"fmt"
	"time"
)

// Subject ...
type Subject struct {
	Name        string
	Assignments []float64
}

// Student ...
type Student struct {
	Name     string
	Subjects []Subject
	Ielts    bool
	Jlpt     bool
}

func average(assignment []float64) float64 {
	total := 0.0
	for _, v := range assignment {
		total += v
	}
	return total / float64(len(assignment))
}

func isOverFive(assignments []float64) bool {
	for _, assignment := range assignments {
		if assignment < 5 {
			return false
		}
	}
	return true
}

func evaluateAssignments(assignments []float64) EvaluationProcess {
	// Check average mark
	evaluationProcess := EvaluationProcess{
		Chapter: "All items are over 5 and the average is $gte 8.5",
		ComposedValue: ResultValue{
			ValueType: ValueBoolean,
			Value:     true,
		},
	}

	averageMark := average(assignments)

	if (averageMark < 8.5) && (isOverFive(assignments) == false) {
		evaluationProcess.ComposedValue.Value = false
	}

	evaluationProcess.References = []ReferenceResource{
		ReferenceResource{
			Display: fmt.Sprintf("The assignments are %v", assignments),
		},
		ReferenceResource{
			Display: fmt.Sprintf("The average is %.1f", averageMark),
		},
	}
	return evaluationProcess
}

func (s Student) evaluateIelts() EvaluationProcess {
	evaluationProcess := EvaluationProcess{
		Chapter: "The student has the IELTS certificate",
		ComposedValue: ResultValue{
			ValueType: ValueBoolean,
			Value:     false,
		},
	}

	if s.Ielts == false {
		evaluationProcess.References = append(evaluationProcess.References,
			ReferenceResource{
				Display: "The student does not have the IELTS certificate",
			})
	} else {
		evaluationProcess.ComposedValue.Value = true
		evaluationProcess.References = append(evaluationProcess.References,
			ReferenceResource{
				Display: "The student have the IELTS certificate",
			})
	}

	return evaluationProcess
}

func (s Student) evaluateJlpt() EvaluationProcess {
	evaluationProcess := EvaluationProcess{
		Chapter: "The student has the JLPT certificate",
		ComposedValue: ResultValue{
			ValueType: ValueBoolean,
			Value:     false,
		},
	}

	if s.Jlpt == false {
		evaluationProcess.References = append(evaluationProcess.References,
			ReferenceResource{
				Display: "The student does not have the JLPT certificate",
			})
	} else {
		evaluationProcess.ComposedValue.Value = true
		evaluationProcess.References = append(evaluationProcess.References,
			ReferenceResource{
				Display: "The student have the JLPT certificate",
			})
	}

	return evaluationProcess
}

func evaluateSubjects(subjects []Subject) EvaluationProcess {
	evaluationProcess := EvaluationProcess{
		Chapter: "Every subject's assignments are over 5 and the average is $gte 8.5",
		ComposedValue: ResultValue{
			ValueType: ValueBoolean,
			Value:     true,
		},
	}

	for _, subject := range subjects {
		// Check each assignment
		assignmentsEvaluationProcess := evaluateAssignments(subject.Assignments)
		evaluationProcess.Childs = append(evaluationProcess.Childs, assignmentsEvaluationProcess)
		if assignmentsEvaluationProcess.ComposedValue.Value == false {
			evaluationProcess.ComposedValue.Value = false
			return evaluationProcess
		}
	}
	return evaluationProcess
}

// Evaluate check if the student is capable for Good or Very Good honor
func (s Student) Evaluate() EvaluationResult {
	evaluationResult := EvaluationResult{
		EvaluationDate: time.Now(),
		Result: Result{
			Good: ResultValue{
				ValueType: ValueBoolean,
				Value:     false,
			},
			VeryGood: ResultValue{
				ValueType: ValueBoolean,
				Value:     false,
			},
		},
		MeasureID:           "2019-ABX-ESAS",
		MeasureDesc:         "Evaluate-student-and-average-score",
		EvaluationProcesses: EvaluationProcesses{},
	}

	// Evaluate Subjects
	subjectsEvaluationProcess := evaluateSubjects(s.Subjects)
	evaluationResult.EvaluationProcesses = append(
		evaluationResult.EvaluationProcesses,
		subjectsEvaluationProcess,
	)

	if subjectsEvaluationProcess.ComposedValue.Value == false {
		// Terminate process
		return evaluationResult
	}

	// Evaluate Ielts
	ieltsEvaluationProcess := s.evaluateIelts()
	evaluationResult.EvaluationProcesses = append(
		evaluationResult.EvaluationProcesses,
		ieltsEvaluationProcess,
	)
	if ieltsEvaluationProcess.ComposedValue.Value == false {
		// Terminate process
		return evaluationResult
	}
	// The student is now qualified for the Good student hornor
	evaluationResult.Result.Good.Value = true

	// Evaluate JLPT
	jlptEvaluationProcess := s.evaluateJlpt()
	evaluationResult.EvaluationProcesses = append(
		evaluationResult.EvaluationProcesses,
		jlptEvaluationProcess,
	)
	if jlptEvaluationProcess.ComposedValue.Value == false {
		// Terminate process
		return evaluationResult
	}
	evaluationResult.Result.VeryGood.Value = true

	return evaluationResult
}

// Average returns the overall score of the student
func (s Student) Average() EvaluationResult {
	evaluationResult := EvaluationResult{
		EvaluationDate:      time.Now(),
		MeasureID:           "2019-ABX-AS",
		MeasureDesc:         "Average-score",
		EvaluationProcesses: EvaluationProcesses{},
	}

	sm1 := float64(0)
	for _, subject := range s.Subjects {
		sm2 := float64(0)
		for _, assignment := range subject.Assignments {
			sm2 = sm2 + assignment
		}
		subjectAverage := sm2 / float64(len(subject.Assignments))
		evaluationResult.EvaluationProcesses = append(evaluationResult.EvaluationProcesses, EvaluationProcess{
			Chapter: fmt.Sprintf("The average of subject %s", subject.Name),
			ComposedValue: ResultValue{
				ValueType: ValueFloat,
				Value:     subjectAverage,
			},
			References: []ReferenceResource{
				ReferenceResource{
					Display: fmt.Sprintf("Assignment %v", subject.Assignments),
				},
			},
		})

		sm1 = sm1 + subjectAverage
	}
	average := sm1 / float64(len(s.Subjects))
	evaluationResult.Result.Mean = ResultValue{
		ValueType: ValueFloat,
		Value:     average,
	}

	return evaluationResult
}
