package app

type Mode string

const (
	ModeProduction  Mode = "production"
	ModeStaging     Mode = "staging"
	ModeDevelopment Mode = "development"
)
