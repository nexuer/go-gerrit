package gerrit

import (
	"errors"
	"time"
)

const MetaConfigRef = "refs/meta/config"
const HeadRef = "HEAD"

// RevisionKind describes the change kind.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#revision-info
type RevisionKind string

const (
	Rework                 RevisionKind = "REWORK"
	TrivialRebase          RevisionKind = "TRIVIAL_REBASE"
	MergeFirstParentUpdate RevisionKind = "MERGE_FIRST_PARENT_UPDATE"
	NoCodeChange           RevisionKind = "NO_CODE_CHANGE"
	NoChange               RevisionKind = "NO_CHANGE"
)

type ProjectState string

const (
	Active   ProjectState = "ACTIVE"
	ReadOnly ProjectState = "READ_ONLY"
	Hidden   ProjectState = "HIDDEN"
)

type PermissionAction string

const (
	Allow       PermissionAction = "ALLOW"
	Deny        PermissionAction = "DENY"
	Block       PermissionAction = "BLOCK"
	Interactive PermissionAction = "INTERACTIVE"
	Batch       PermissionAction = "BATCH"
)

type ProjectType string

const (
	All         ProjectType = "ALL"
	Code        ProjectType = "CODE"
	Permissions ProjectType = "PERMISSIONS"
)

// Gerrit's timestamp layout is like time.RFC3339Nano, but with a space instead
// of the "T", without a timezone (it's always in UTC), and always includes nanoseconds.
// See https://gerrit-review.googlesource.com/Documentation/rest-api.html#timestamp.
const timeLayout = "2006-01-02 15:04:05.000000000"

// Timestamp represents an instant in time with nanosecond precision, in UTC time zone.
// It encodes to and from JSON in Gerrit's timestamp format.
// All exported methods of time.Time can be called on Timestamp.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api.html#timestamp
type Timestamp struct {
	// Time is an instant in time. Its time zone must be UTC.
	time.Time
}

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in Gerrit's timestamp format.
// An error is returned if t.Time time zone is not UTC.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	if t.Location() != time.UTC {
		return nil, errors.New("Timestamp.MarshalJSON: time zone must be UTC")
	}
	if y := t.Year(); y < 0 || 9999 < y {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#issuecomment-66073163 for more discussion.
		return nil, errors.New("Timestamp.MarshalJSON: year outside of range [0,9999]")
	}
	b := make([]byte, 0, len(timeLayout)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, timeLayout)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in Gerrit's timestamp format.
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	// Ignore null, like in the main JSON package.
	if string(b) == "null" {
		return nil
	}
	var err error
	t.Time, err = time.Parse(`"`+timeLayout+`"`, string(b))
	return err
}
