package message

// Identifier is an interface that can uniquely identify itself.
type Identifier interface {
	ID() string
	SetID(string)
	Name() string
	Fingerprint() string
}

// SimpleID is a simple Identifier implementation used for testing.
type SimpleID string

// ID returns the ID as a string.
func (i SimpleID) ID() string {
	return string(i)
}

// SetID is a no-op
func (i SimpleID) SetID(s string) {
	// no-op
}

// Fingerprint returns an empty string
func (i SimpleID) Fingerprint() string {
	return ""
}

// Name returns the ID
func (i SimpleID) Name() string {
	return i.ID()
}
