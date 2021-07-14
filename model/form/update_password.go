package form

type UpdatePasswordForm struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
