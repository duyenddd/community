// Package manifest provides structures and primitives to define apps.
package manifest

// Manifest is a structure to define a starlark applet for Tidbyt in Go.
type Manifest struct {
	// Name is the name of the applet. Ex. "Fuzzy Clock"
	Name string `json:"name"`
	// Summary is the short form of what this applet does. Ex. "Human readable
	// time".
	Summary string `json:"summary"`
	// Desc is the long form of what this applet does. Ex. "Display the time in
	// a groovy, human-readable way."
	Desc string `json:"desc"`
	// Author is the person or organization who contributed this applet. Ex,
	// "Max Timkovich"
	Author string `json:"author"`
	// Source is the starlark source code for this applet using the go `embed`
	// module.
	Source []byte `json:"-"`
}

func (m Manifest) Validate() error {
	err := ValidateName(m.Name)
	if err != nil {
		return err
	}

	err = ValidateSummary(m.Summary)
	if err != nil {
		return err
	}

	err = ValidateDesc(m.Desc)
	if err != nil {
		return err
	}

	err = ValidateAuthor(m.Author)
	if err != nil {
		return err
	}

	return nil
}
