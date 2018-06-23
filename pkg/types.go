package impi

// ImportGroupVerificationScheme specifies what to check when inspecting import groups
type ImportGroupVerificationScheme int

// VerifyOptions specifies how to perform verification
type VerifyOptions struct {
	SkipTests   bool
	Scheme      ImportGroupVerificationScheme
	LocalPrefix string
	SkipPaths   []string
}

// VerificationError holds an error and a file path on which the error occurred
type VerificationError struct {
	error
	FilePath string
}

// ErrorReporter receives error reports as they are detected by the workers
type ErrorReporter interface {
	Report(VerificationError)
}
