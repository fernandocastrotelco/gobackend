package simplemath

import "fmt"

type SemanticVersion struct {
	major, minor, patch int
}

func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

func (sv SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

func (sv SemanticVersion) IncrementMajor() SemanticVersion {
	sv.major += 1
	return sv
}

func (sv SemanticVersion) IncrementMinor() SemanticVersion {
	sv.minor += 1
	return sv
}

func (sv SemanticVersion) IncrementPatch() SemanticVersion {
	sv.patch += 1
	return sv
}

