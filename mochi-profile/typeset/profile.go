package typeset

type ProfileType string

const (
	PROFILE_TYPE_USER              ProfileType = "user"
	PROFILE_TYPE_APPLICATION       ProfileType = "application"
	PROFILE_TYPE_VAULT             ProfileType = "vault"
	PROFILE_TYPE_APPLICATION_VAULT ProfileType = "application_vault"
	PROFILE_TYPE_EARNING_VAULT     ProfileType = "earning_vault"
)
