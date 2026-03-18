package model

// VerificationCodeType represents the purpose of a verification code
type VerificationCodeType string

const (
	VerificationCodeTypeEmailVerification VerificationCodeType = "email_verification"
	VerificationCodeTypeOrganizationJoin  VerificationCodeType = "organization_join"
)

// String converts to string value
func (t VerificationCodeType) String() string {
	return string(t)
}

// IsValid checks if VerificationCodeType is valid
func (t VerificationCodeType) IsValid() bool {
	return t == VerificationCodeTypeEmailVerification || t == VerificationCodeTypeOrganizationJoin
}

// FriendshipStatus represents the status of a friendship between two users
type FriendshipStatus string

const (
	FriendshipStatusPending  FriendshipStatus = "pending"
	FriendshipStatusAccepted FriendshipStatus = "accepted"
	FriendshipStatusBlocked  FriendshipStatus = "blocked"
)

// String converts to string value
func (s FriendshipStatus) String() string {
	return string(s)
}

// IsValid checks if FriendshipStatus is valid
func (s FriendshipStatus) IsValid() bool {
	return s == FriendshipStatusPending ||
		s == FriendshipStatusAccepted ||
		s == FriendshipStatusBlocked
}

// WorkplaceSize represents the size category of a workplace
type WorkplaceSize string

const (
	WorkplaceSizeXS WorkplaceSize = "1-10"
	WorkplaceSizeS  WorkplaceSize = "11-50"
	WorkplaceSizeM  WorkplaceSize = "51-200"
	WorkplaceSizeL  WorkplaceSize = "201-500"
	WorkplaceSizeXL WorkplaceSize = "500+"
)

// String converts to string value
func (s WorkplaceSize) String() string {
	return string(s)
}

// IsValid checks if WorkplaceSize is valid
func (s WorkplaceSize) IsValid() bool {
	return s == WorkplaceSizeXS ||
		s == WorkplaceSizeS ||
		s == WorkplaceSizeM ||
		s == WorkplaceSizeL ||
		s == WorkplaceSizeXL
}
