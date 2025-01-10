package constants

type ADMIN_ROLE string

const (
	SUPER_ADMIN ADMIN_ROLE = "super_admin"
	ADMIN       ADMIN_ROLE = "admin"
	MODERATOR   ADMIN_ROLE = "moderator"
	WRITER      ADMIN_ROLE = "writer"
)

type ADMIN_STATUS string

const (
	ACTIVE  ADMIN_STATUS = "active"
	PASSICE ADMIN_STATUS = "passive"
)
