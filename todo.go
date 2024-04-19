package main

var (
	TODOs         []TODO
	TODOSequencer uint32
)

func init() {
	CreateTODO("Trocar agua dos gatos")
	CreateTODO("Comprar ração")
	CreateTODO("Estudar Go")
	CreateTODO("Estudar Fiber")
	CreateTODO("Estudar React")
	CreateTODO("Estudar Docker")
	CreateTODO("Estudar Kubernetes")
	CreateTODO("Estudar CI/CD")
	CreateTODO("Estudar TDD")
	CreateTODO("Estudar Clean Code")
	CreateTODO("Estudar Design Patterns")
	CreateTODO("Estudar SOLID")
	CreateTODO("Estudar DDD")
	CreateTODO("Limpar merda dos gatos")
}

type TODO struct {
	ID    uint32
	Title string
	Done  bool
}

func NewTODO() TODO {
	return TODO{}
}

func CreateTODO(title string) {
	todo := NewTODO()
	TODOSequencer++
	todo.ID = TODOSequencer
	todo.Title = title
	todo.Done = false
	TODOs = append(TODOs, todo)
}

func DeleteTODO(id uint32) {
	for i, todo := range TODOs {
		if todo.ID == id {
			TODOs = append(TODOs[:i], TODOs[i+1:]...)
			break
		}
	}
}
