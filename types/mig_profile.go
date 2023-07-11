package types

import (
	nvdev "gitlab.com/nvidia/cloud-native/go-nvlib/pkg/nvlib/device"
)

const (
	AttributeMediaExtensions = nvdev.AttributeMediaExtensions
)

// MigProfile represents a specific MIG profile.
// Examples include "1g.5gb", "2g.10gb", "1c.2g.10gb", or "1c.1g.5gb+me", etc.
type MigProfile struct {
	nvdev.MigProfileInfo
}

// AssertValidMigProfileFormat checks if the string is in the proper format to represent a MIG profile
func AssertValidMigProfileFormat(profile string) error {
	return nvdevAssertValidMigProfileFormat(profile)
}

// ParseMigProfile converts a string representation of a MigProfile into an object.
func ParseMigProfile(profile string) (*MigProfile, error) {
	mp, err := nvdevParseMigProfile(profile)
	if err != nil {
		return nil, err
	}
	return &MigProfile{mp.GetInfo()}, nil
}

// HasAttribute checks if the MigProfile has the specified attribute associated with it.
func (m MigProfile) HasAttribute(attr string) bool {
	for _, a := range m.Attributes {
		if a == attr {
			return true
		}
	}
	return false
}
