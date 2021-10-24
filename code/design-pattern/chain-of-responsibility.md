# Chain of Responsibility Pattern

## Go

```go
package main

import (
	"fmt"
)

type Job struct {
	Name string
}

func (j *Job) String() string {
	return fmt.Sprintf("Job [%s]", j.Name)
}

type Step interface {
	Execute(job *Job)
}

type BuildStep struct {
	next Step
}

func (s *BuildStep) Execute(job *Job) {
	fmt.Printf("%s: building ...\n", job)

	if s.next != nil {
		s.next.Execute(job)
	}
}

func NewBuildStep(next Step) Step {
	return &BuildStep{
		next: next,
	}
}

type TestStep struct {
	next Step
}

func (s *TestStep) Execute(job *Job) {
	fmt.Printf("%s: testing ...\n", job)

	if s.next != nil {
		s.next.Execute(job)
	}
}

func NewTestStep(next Step) Step {
	return &TestStep{
		next: next,
	}
}

type PublishStep struct {
	next Step
}

func (s *PublishStep) Execute(job *Job) {
	fmt.Printf("%s: publishing ...\n", job)

	if s.next != nil {
		s.next.Execute(job)
	}
}

func NewPublishStep(next Step) Step {
	return &PublishStep{
		next: next,
	}
}

func main() {
	job := &Job{Name: "CI/CD-001"}
	step := NewBuildStep(NewTestStep(NewPublishStep(nil)))
	step.Execute(job)
}
```

```
Job [CI/CD-001]: building ...
Job [CI/CD-001]: testing ...
Job [CI/CD-001]: publishing ...
```
