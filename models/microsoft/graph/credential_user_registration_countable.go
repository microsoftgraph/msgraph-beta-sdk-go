package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CredentialUserRegistrationCountable 
type CredentialUserRegistrationCountable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetTotalUserCount()(*int64)
    GetUserRegistrationCounts()([]UserRegistrationCountable)
    SetTotalUserCount(value *int64)()
    SetUserRegistrationCounts(value []UserRegistrationCountable)()
}
