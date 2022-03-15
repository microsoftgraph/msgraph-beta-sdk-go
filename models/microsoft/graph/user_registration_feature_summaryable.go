package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserRegistrationFeatureSummaryable 
type UserRegistrationFeatureSummaryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetTotalUserCount()(*int32)
    GetUserRegistrationFeatureCounts()([]UserRegistrationFeatureCountable)
    GetUserRoles()(*IncludedUserRoles)
    GetUserTypes()(*IncludedUserTypes)
    SetTotalUserCount(value *int32)()
    SetUserRegistrationFeatureCounts(value []UserRegistrationFeatureCountable)()
    SetUserRoles(value *IncludedUserRoles)()
    SetUserTypes(value *IncludedUserTypes)()
}
