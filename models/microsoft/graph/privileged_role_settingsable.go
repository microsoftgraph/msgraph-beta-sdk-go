package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrivilegedRoleSettingsable 
type PrivilegedRoleSettingsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApprovalOnElevation()(*bool)
    GetApproverIds()([]string)
    GetElevationDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetIsMfaOnElevationConfigurable()(*bool)
    GetLastGlobalAdmin()(*bool)
    GetMaxElavationDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetMfaOnElevation()(*bool)
    GetMinElevationDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetNotificationToUserOnElevation()(*bool)
    GetTicketingInfoOnElevation()(*bool)
    SetApprovalOnElevation(value *bool)()
    SetApproverIds(value []string)()
    SetElevationDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetIsMfaOnElevationConfigurable(value *bool)()
    SetLastGlobalAdmin(value *bool)()
    SetMaxElavationDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetMfaOnElevation(value *bool)()
    SetMinElevationDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetNotificationToUserOnElevation(value *bool)()
    SetTicketingInfoOnElevation(value *bool)()
}
