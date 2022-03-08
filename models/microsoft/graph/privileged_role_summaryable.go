package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrivilegedRoleSummaryable 
type PrivilegedRoleSummaryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetElevatedCount()(*int32)
    GetManagedCount()(*int32)
    GetMfaEnabled()(*bool)
    GetStatus()(*RoleSummaryStatus)
    GetUsersCount()(*int32)
    SetElevatedCount(value *int32)()
    SetManagedCount(value *int32)()
    SetMfaEnabled(value *bool)()
    SetStatus(value *RoleSummaryStatus)()
    SetUsersCount(value *int32)()
}
