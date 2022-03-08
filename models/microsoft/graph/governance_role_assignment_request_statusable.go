package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernanceRoleAssignmentRequestStatusable 
type GovernanceRoleAssignmentRequestStatusable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetStatus()(*string)
    GetStatusDetails()([]KeyValueable)
    GetSubStatus()(*string)
    SetStatus(value *string)()
    SetStatusDetails(value []KeyValueable)()
    SetSubStatus(value *string)()
}
