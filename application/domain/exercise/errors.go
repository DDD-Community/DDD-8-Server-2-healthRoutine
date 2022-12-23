package exercise

var (
	ErrNotMatchUserId error = &errNotMatchUserId{}
)

type errNotMatchUserId struct{}

func (e *errNotMatchUserId) Error() string {
	return "user id doesn't match"
}
