package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TenantSetupInfoable 
type TenantSetupInfoable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDefaultRolesSettings()(PrivilegedRoleSettingsable)
    GetFirstTimeSetup()(*bool)
    GetRelevantRolesSettings()([]string)
    GetSetupStatus()(*SetupStatus)
    GetSkipSetup()(*bool)
    GetUserRolesActions()(*string)
    SetDefaultRolesSettings(value PrivilegedRoleSettingsable)()
    SetFirstTimeSetup(value *bool)()
    SetRelevantRolesSettings(value []string)()
    SetSetupStatus(value *SetupStatus)()
    SetSkipSetup(value *bool)()
    SetUserRolesActions(value *string)()
}
