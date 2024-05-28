package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func (p *Person) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Age, validation.Required, validation.Min(0)),
		validation.Field(&p.Email, validation.Required, is.Email),
	)
}

// This directive is for [tinygo](https://tinygo.org/getting-started/install/).
// For go global functions should be fine
// https://stackoverflow.com/questions/67978442/go-wasm-export-functions
//
//export validatePersonJSON
func validatePersonJSON(this js.Value, args []js.Value) interface{} {
	personJSON := args[0].String()
	var person Person
	if err := json.Unmarshal([]byte(personJSON), &person); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return false
	}

	if err := person.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return false
	}

	return true
}

func main() {
	js.Global().Set("bam_validatePersonJSON", js.FuncOf(validatePersonJSON))
	select {}
}
