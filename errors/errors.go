package errors

type JWTTokenExpiredError struct {}
func (j JWTTokenExpiredError) Error() string {
	return "JWT Expired"
}

type JWTSignatureMismatch struct {}
func (j JWTSignatureMismatch) Error() string {
	return "The JWT Signatures Do Not Match"
}
