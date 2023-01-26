package user

var (
	ErrEmailAlreadyExists    error = &errEmailAlreadyExists{}
	ErrNicknameAlreadyExists error = &errNicknameAlreadyExists{}
	ErrUserNotFound          error = &errUserNotFound{}
	ErrNoRecordDrink         error = &errNoRecordDrink{}
)

type errEmailAlreadyExists struct{}

func (e *errEmailAlreadyExists) Error() string {
	return "email already exists"
}

type errNicknameAlreadyExists struct{}

func (e *errNicknameAlreadyExists) Error() string {
	return "nickname already exists"
}

type errUserNotFound struct{}

func (e *errUserNotFound) Error() string {
	return "user not found"
}

type errNoRecordDrink struct{}

func (e *errNoRecordDrink) Error() string {
	return "no record of drinking water"
}
