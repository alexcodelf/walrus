package types

import "golang.org/x/exp/slices"

const (
	SubjectKindUser  = "user"
	SubjectKindGroup = "group"
)

var SubjectKinds = []string{
	SubjectKindUser,
	SubjectKindGroup,
}

func IsSubjectKind(s string) bool {
	return slices.Contains(SubjectKinds, s)
}

const SubjectDomainBuiltin = "builtin"

// SubjectUserBot is the name of the bot user.
const SubjectUserBot = "bot"
