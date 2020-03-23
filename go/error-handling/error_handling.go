package erratum

// Use opens a resource, calls Frob(input) on the result resource and then closes that resource (in all cases).
func Use(o ResourceOpener, input string) (err error) {
	var resource Resource

	if resource, err = o(); err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}
		return
	}

	defer func() {
		if r := recover(); r != nil {
			if f, ok := r.(FrobError); ok {
				resource.Defrob(f.defrobTag)
			}
			err = r.(error)
		}
		resource.Close()
	}()

	resource.Frob(input)
	return
}
