package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EntitlementManagementSettingsable 
type EntitlementManagementSettingsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDaysUntilExternalUserDeletedAfterBlocked()(*int32)
    GetExternalUserLifecycleAction()(*string)
    SetDaysUntilExternalUserDeletedAfterBlocked(value *int32)()
    SetExternalUserLifecycleAction(value *string)()
}
