//https://www.youtube.com/watch?v=F4wUrj6pmSI&t=2169s

//Interface as Contract

abstract types
- describes behaviour
      io.Reader io.Writer fmt.Stringer
- they define a set of methods, without specifying the receiver

type Positiver interface {
	Positive() bool
}


//Interface as Types

//Interface Satisfaction