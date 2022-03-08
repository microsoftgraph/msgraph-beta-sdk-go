package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcOrganizationSettingsable 
type CloudPcOrganizationSettingsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetOsVersion()(*CloudPcOperatingSystem)
    GetUserAccountType()(*CloudPcUserAccountType)
    SetOsVersion(value *CloudPcOperatingSystem)()
    SetUserAccountType(value *CloudPcUserAccountType)()
}
