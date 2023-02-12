package utilx

import "io.github.tuanictu97/api/pkg/errors"

var (
	ErrInvalidUsernameOrPassword = errors.BadRequest(errors.ErrBadRequestID, "Invalid username or password")
	ErrUserFreezed               = errors.Forbidden(errors.ErrForbiddenID, "User is freezed")
	ErrPermissionDenied          = errors.Forbidden("com.perm.denied", "Permission denied")
	ErrInvalidToken              = errors.Unauthorized("com.invalid.token", "Invalid token")
)
