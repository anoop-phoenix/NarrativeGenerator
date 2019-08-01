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
}

func average(assignment []float64) float64 {
	total := 0.0
	for _, v := range assignment {
		total += v
	}
	return total / float64(len(assignment))
}

func evaluateSubjects(subjects []Subject) Narrative {
	narrative := Narrative{
		Desc:   "Every subject's assignments are over 5 and the average is $gte 8.5",
		Status: true,
	}
	for _, subject := range subjects {
		// Check each assignment
		assignmentsNarrative := evaluateAssignments(subject.Assignments)
		narrative.Childs = append(narrative.Childs, assignmentsNarrative)
		if assignmentsNarrative.Status == false {
			narrative.Status = false
			return narrative
		}
	}
	return narrative
}

func evaluateAssignments(assignments []float64) Narrative {
	// Check average mark
	narrative := Narrative{
		Desc:   "All items are over 5 and the average is $gte 8.5",
		Status: true,
	}

	// Check the average
	averageNarrative := Narrative{
		Desc:   "The average is $gte 8.5",
		Status: true,
	}

	averageMark := average(assignments)
	if averageMark < 8.5 {
		averageNarrative.Status = false
	}
	averageNarrative.References = []ReferenceResource{
		ReferenceResource{
			Display: fmt.Sprintf("The assignments are %v", assignments),
		},
		ReferenceResource{
			Display: fmt.Sprintf("The average is %.1f", averageMark),
		},
	}
	narrative.Childs = append(narrative.Childs, averageNarrative)
	if averageNarrative.Status == false {
		narrative.Status = false
		return narrative
	}

	// Check the assignments
	assignmentsNarrative := Narrative{
		Desc:   "All assignments is over 5",
		Status: true,
	}

	for _, assignment := range assignments {
		if assignment < 5 {
			assignmentsNarrative.Status = false
			break
		}
	}
	assignmentsNarrative.References = []ReferenceResource{
		ReferenceResource{
			Display: fmt.Sprintf("The assignments are %v", assignments),
		},
	}
	narrative.Childs = append(narrative.Childs, assignmentsNarrative)
	if assignmentsNarrative.Status == false {
		narrative.Status = false
		return narrative
	}
	return narrative
}

func evaluateIelts(s Student) Narrative {
	narrative := Narrative{
		Desc: "The student has the IELTS certificate",
	}
	ref := ReferenceResource{}
	if s.Ielts {
		narrative.Status = true
		ref.Display = "The student has the IELTS certificate"
	} else {
		narrative.Status = false
		ref.Display = "The student doesn't have the IELTS certificate"
	}

	narrative.References = []ReferenceResource{
		ref,
	}
	return narrative
}

func isVeryGoodStudent(s Student) Narrative {
	narrative := Narrative{
		Desc:   fmt.Sprintf("%s is a very good student", s.Name),
		Status: true,
	}
	// Evaluate Subjects
	subjectsNarrative := evaluateSubjects(s.Subjects)
	narrative.Childs = append(narrative.Childs, subjectsNarrative)
	if subjectsNarrative.Status == false {
		narrative.Status = false
		return narrative
	}
	// Evaluate Ielts
	ieltsNarrative := evaluateIelts(s)
	narrative.Childs = append(narrative.Childs, ieltsNarrative)
	if ieltsNarrative.Status == false {
		narrative.Status = false
		return narrative
	}

	return narrative
}

// Evaluate ...
func (s Student) Evaluate() EvaluateResult {
	evaluateResult := EvaluateResult{
		EvaluatedDate: time.Now(),
	}

	evaluateResult.Narrative = isVeryGoodStudent(s)
	return evaluateResult
}
