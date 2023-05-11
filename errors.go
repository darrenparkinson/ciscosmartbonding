package ciscosmartbonding

// Err implements the error interface so we can have constant errors.
type Err string

func (e Err) Error() string {
	return string(e)
}

// Error Constants
const (
	ErrProcessingListParams = Err("ciscosmartbonding: error processing list parameters")

	ErrBadRequest    = Err("ciscosmartbonding: bad request")
	ErrUnauthorized  = Err("ciscosmartbonding: unauthorized request")
	ErrForbidden     = Err("ciscosmartbonding: forbidden")
	ErrInternalError = Err("ciscosmartbonding: internal error")
	ErrUnknown       = Err("ciscosmartbonding: unexpected error occurred")
	ErrNotFound      = Err("ciscosmartbonding: not found")

	ErrUnsupportedType = Err("ciscosmartbonding: unsupported type")
)
