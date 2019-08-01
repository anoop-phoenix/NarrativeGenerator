package core

import (
	"fmt"
	"time"
)

type Assignment float64

type Subject struct {
	Name        string
	Assignments []Assignment
}

type Student struct {
	Name     string
	Subjects []Subject
	Ielts    bool
}

func average(assignment []Assignment) float64 {
	total := 0.0
	for _, v := range assignment {
		total += float64(v)
	}
	return total / float64(len(assignment))
}

func evaluate_subjects(subjects []Subject) Narrative {
	narrative := Narrative{
		Desc:   "Every subject's assignments are over 5 and the average is $gte 8.5",
		Status: true,
	}
	for _, subject := range subjects {
		// Check each assignment
		assignments_narrative := evaluate_assignments(subject.Assignments)
		narrative.Childs = append(narrative.Childs, assignments_narrative)
		if assignments_narrative.Status == false {
			narrative.Status = false
			return narrative
		}
	}
	return narrative
}

func evaluate_assignments(assignments []Assignment) Narrative {
	// Check average mark
	narrative := Narrative{
		Desc:   "All items are over 5 and the average is $gte 8.5",
		Status: true,
	}

	// Check the average
	average_narrative := Narrative{
		Desc:   "The average is $gte 8.5",
		Status: true,
	}

	average_mark := average(assignments)
	if average_mark < 8.5 {
		average_narrative.Status = false
	}
	narrative.Childs = append(narrative.Childs, average_narrative)
	if average_narrative.Status == false {
		narrative.Status = false
		return narrative
	}

	// Check the assignments
	assignments_narrative := Narrative{
		Desc:   "All assignments is over 5",
		Status: true,
	}

	for _, assignment := range assignments {
		if assignment < 5 {
			assignments_narrative.Status = false
			break
		}
	}
	narrative.Childs = append(narrative.Childs, assignments_narrative)
	if assignments_narrative.Status == false {
		narrative.Status = false
		return narrative
	}
	return narrative
}

func evaluate_ielts(s Student) Narrative {
	narrative := Narrative{
		Desc: "The student has the IELTS certificate",
	}
	if s.Ielts {
		narrative.Status = true
	} else {
		narrative.Status = false
	}
	return narrative
}

func is_very_good_student(s Student) Narrative {
	narrative := Narrative{
		Desc:   fmt.Sprintf("%s is a very good student", s.Name),
		Status: true,
	}
	// Evaluate Subjects
	subjects_narrative := evaluate_subjects(s.Subjects)
	narrative.Childs = append(narrative.Childs, subjects_narrative)
	if subjects_narrative.Status == false {
		narrative.Status = false
		return narrative
	}
	// Evaluate Ielts
	ielts_narrative := evaluate_ielts(s)
	narrative.Childs = append(narrative.Childs, ielts_narrative)
	if ielts_narrative.Status == false {
		narrative.Status = false
		return narrative
	}

	return narrative
}

func (s Student) Evaluate() EvaluateResult {
	evaluateResult := EvaluateResult{
		EvaluatedDate: time.Now(),
	}

	evaluateResult.Narrative = is_very_good_student(s)
	return evaluateResult
}
