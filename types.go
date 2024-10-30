package gerrit

type ProjectState string

const (
	Active   ProjectState = "ACTIVE"
	ReadOnly ProjectState = "READ_ONLY"
	Hidden   ProjectState = "HIDDEN"
)

type ProjectType string

const (
	All         ProjectType = "ALL"
	Code        ProjectType = "CODE"
	Permissions ProjectType = "PERMISSIONS"
)

type AdditionalField string

const (
	Details   AdditionalField = "DETAILS"
	AllEmails AdditionalField = "ALL_EMAILS"
)
