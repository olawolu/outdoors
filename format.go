package outdoors

// Formatter is implemented by any value that has a Format() method
type Formatter interface {
	Format() interface{}
}

// Format takes in an object and checks if it implements Formatter
func Format(o interface{}) interface{} {
	if object, isFormatter := o.(Formatter); isFormatter {
		return object.Format()
	}
	return o
}