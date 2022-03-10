package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkDeviceSoftwareVersionsable 
type TeamworkDeviceSoftwareVersionsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAdminAgentSoftwareVersion()(*string)
    GetFirmwareSoftwareVersion()(*string)
    GetOperatingSystemSoftwareVersion()(*string)
    GetPartnerAgentSoftwareVersion()(*string)
    GetTeamsClientSoftwareVersion()(*string)
    SetAdminAgentSoftwareVersion(value *string)()
    SetFirmwareSoftwareVersion(value *string)()
    SetOperatingSystemSoftwareVersion(value *string)()
    SetPartnerAgentSoftwareVersion(value *string)()
    SetTeamsClientSoftwareVersion(value *string)()
}
