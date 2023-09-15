// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package semver

import (
	"fmt"
)

// Version is the major/minor/patch of the application.
type Version struct {
	Major      int
	Minor      int
	Patch      int
	PreRelease string
	Build      string
}

// String implements the Stringer interface.
func (v Version) String() string {
	// format the version per https://semver.org/ rules
	hasPreRelease, hasBuild := v.PreRelease != "", v.Build != ""
	if hasPreRelease && hasBuild {
		return fmt.Sprintf("%d.%d.%d-%s+%s", v.Major, v.Minor, v.Patch, v.PreRelease, v.Build)
	} else if hasPreRelease && !hasBuild {
		return fmt.Sprintf("%d.%d.%d-%s", v.Major, v.Minor, v.Patch, v.PreRelease)
	} else if !hasPreRelease && hasBuild {
		return fmt.Sprintf("%d.%d.%d+%s", v.Major, v.Minor, v.Patch, v.Build)
	}
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
