package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Employee struck {
	ID 		int
	Name 	string
	Age		int
	Division string
}

var employees = []Employee {
	{ID: 1, Name: "Billy", Age: 17, Division: "IT"}
}

var PORT = ":3070"