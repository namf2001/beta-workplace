package model

// Provider represents an authentication provider
type Provider string

const (
	// ProviderCredentials is for email/password registration
	ProviderCredentials Provider = "credentials"
	// ProviderGoogle is for Google OAuth
	ProviderGoogle Provider = "google"
	// ProviderGitHub is for GitHub OAuth
	ProviderGitHub Provider = "github"
	// ProviderDiscord is for Discord OAuth
	ProviderDiscord Provider = "discord"
	// ProviderMicrosoft is for Microsoft OAuth
	ProviderMicrosoft Provider = "microsoft"
)

// String converts to string value
func (p Provider) String() string {
	return string(p)
}

// IsValid checks if Provider is valid
func (p Provider) IsValid() bool {
	return p == ProviderCredentials ||
		p == ProviderGoogle ||
		p == ProviderGitHub ||
		p == ProviderDiscord ||
		p == ProviderMicrosoft
}

// IsOAuth returns true if the provider is an OAuth provider (not credentials)
func (p Provider) IsOAuth() bool {
	return p != ProviderCredentials
}
