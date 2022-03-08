package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobilityManagementPolicyable 
type MobilityManagementPolicyable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppliesTo()(*PolicyScope)
    GetComplianceUrl()(*string)
    GetDescription()(*string)
    GetDiscoveryUrl()(*string)
    GetDisplayName()(*string)
    GetIncludedGroups()([]Groupable)
    GetIsValid()(*bool)
    GetTermsOfUseUrl()(*string)
    SetAppliesTo(value *PolicyScope)()
    SetComplianceUrl(value *string)()
    SetDescription(value *string)()
    SetDiscoveryUrl(value *string)()
    SetDisplayName(value *string)()
    SetIncludedGroups(value []Groupable)()
    SetIsValid(value *bool)()
    SetTermsOfUseUrl(value *string)()
}
