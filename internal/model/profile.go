package model

// ProfileShellType is a type for profile shell.
type ProfileShellType string

// all possible values for ProfileShellType.
const (
	Bash ProfileShellType = "bash"
	Fish ProfileShellType = "fish"
	Zsh  ProfileShellType = "zsh"
)
