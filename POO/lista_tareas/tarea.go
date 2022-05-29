package main

import (
	"fmt"
)

//lista de tareas
type TaskList struct {
	tasks [] *Task

	//task seria el array que guarda las tareas
	// task[0] 0 1 2 
			
	//task [1] 0 1 2
}

func (tl *TaskList) appendTask (t *Task) {

	//Lo que hago es agregar un Objeto de la clase Task a mi lista de tareas
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) deleteTask (index int) {

	//Lo que hago es agregar un Objeto de la clase Task a mi lista de tareas
	tl.tasks = append(tl.tasks[:index],tl.tasks[index+1:]... )

	//Lo que hago aca es crear un nuevo slice que agarre de 0 hasta antes del index, ejemplo si elijo 2 tomaria 0 y 1. Despues le agrego los datos siguientes del index osea agregaria 3,4 y sucesivamente 
}




//tareas
type Task struct {
	name      string
	desc      string
	completed bool
}

func (t *Task) imprimir() {

	fmt.Printf("Nombre: %s\nDescripción: %s\nCompletado:%t\n",t.name,t.desc,t.completed)
}

func (t *Task) markCompleted () {

	t.completed = true
}

func main() {

	t1 := Task{
		name: "Cursos Go",
		desc: "POO este mes completar",
		completed: false,
	}

	t2 := Task{
		name: "Cursos HTML",
		desc: "HTML esta semana completar",
		completed: true,
	}


	list := TaskList{}
	list.appendTask(&t1)
	list.appendTask(&t2)

	t3 := Task{
		name: "Cursos CSS",
		desc: "CSS esta semana completar",
		completed: false,
	}

	list.appendTask(&t3)
	list.deleteTask(1)

	for i,task := range list.tasks {
		fmt.Println(i,task.name)
	}

	// t1.imprimir()
	// t2.imprimir()
}