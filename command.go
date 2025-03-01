package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *cmdFlags {
	cf := cmdFlags{}

	flag.StringVar(&cf.Add, "Add", "", "Add a new Todo")
	flag.StringVar(&cf.Edit, "Edit", "", "Edit existing Todo")
	flag.IntVar(&cf.Del, "Delete", -1, "Delete a Todo")
	flag.IntVar(&cf.Toggle, "Toggle", -1, "Toggle Todo")
	flag.BoolVar(&cf.List, "List", false, "List All Todos")

	//parse all todos

	flag.Parse()
	return &cf
}

func (cf *cmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid args , Please follow id:title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error")
			os.Exit(1)
		}
		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("Something Went Wrong")
	}

}
