package impi

const (

	// ImportGroupVerificationSchemeSingle allows for a single, sorted group
	ImportGroupVerificationSchemeSingle = ImportGroupVerificationScheme(iota)

	// ImportGroupVerificationSchemeStdNonStd allows for up to two groups in the following order:
	// - standard imports
	// - non-standard imports
	ImportGroupVerificationSchemeStdNonStd

	// ImportGroupVerificationSchemeStdLocalThirdParty allows for up to three groups in the following order:
	// - standard imports
	// - local imports (where local prefix is specified in verification options)
	// - non-standard imports
	ImportGroupVerificationSchemeStdLocalThirdParty

	// ImportGroupVerificationSchemeStdThirdPartyLocal allows for up to three groups in the following order:
	// - standard imports
	// - non-standard imports
	// - local imports (where local prefix is specified in verification options)
	ImportGroupVerificationSchemeStdThirdPartyLocal
)
