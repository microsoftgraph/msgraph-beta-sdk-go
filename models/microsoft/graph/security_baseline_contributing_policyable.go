package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SecurityBaselineContributingPolicyable 
type SecurityBaselineContributingPolicyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetSourceId()(*string)
    GetSourceType()(*SecurityBaselinePolicySourceType)
    SetDisplayName(value *string)()
    SetSourceId(value *string)()
    SetSourceType(value *SecurityBaselinePolicySourceType)()
}
