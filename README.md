# errutil

Package errutil provides simple error utilities predominantly around behavioral errors and panicked errors

### Using Error as a constant
```go
const myErr errutil.StringErr = "ruh oh"
....

if err == "ruh oh" {
	// since errutil.StringErr is a string you can compare with other strings
	// and save the error as a constant
	return "this is fine"
}
```
### Adding Common Behavior To An Error
```go
func myFunc() error {
    const myErr errutil.StringErr = "ruh oh"
    return errutil.Wrap(myErr, errutil.WithNotFound(), errutil.WithConflict())
}

func myFunc2() error {
    return errutil.New("ruh oh", errutil.WithEasyTags(
		"example", "yes",
		"component", "main",
		))
}

func main(){
	err := myFunc()
	if err != nil && errutil.IsNotFound(err) || errutil.IsConflict(err) {
	    fmt.Println("not found or conflict behavior detected .. handling")	
   }

   err = myFunc2()
	if errutil.IsTaggable(err) {
		fmt.Printf("error tags are: %+v", errutil.GetTags(err))
	}
}
```

### Handling Panicked Errors
If you find your code is like the following with multiple error returns consider using the ExpectedPanicAsError function

**Code with pass through error returns**
```go
func myFunc() error {
	if err := a1(); err != nil {
		return err
    }

    if err := a2(); err != nil {
        return err
    }

    if err := a3(); err != nil {
        return err
    }
}
```

**Code with panics**
```go
func myFunc() (e error) {
	defer errutil.ExpectedPanicAsError(&e)
	a1()
	a2()
	a3();
}

func a1(){
	panic(errors.New("notice we are panicking with an error interface"))
}

func a2(){
    panic(errors.New("if a panic is done without an error interface the panic will be re-raised"))
}

func a3(){
    panic(errors.New("if a panic is done without an error interface the panic will be re-raised"))
}
```

