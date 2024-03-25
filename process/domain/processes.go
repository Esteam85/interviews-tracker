package domain

type Processes []Process

func (p Processes) ToProcessesAsPrimitives() []ProcessAsPrimitives {
	var processesAsPrimitives []ProcessAsPrimitives

	for i := range p {
		processesAsPrimitives = append(processesAsPrimitives, *p[i].ToPrimitives())
	}
	return processesAsPrimitives
}
