package twofer

// ShareWith returns the text that we should say when giving the extra cookie to another person
func ShareWith(name string) string {
	if name != "" {
        return "One for " + name + ", one for me."
    }

    return "One for you, one for me."
}
