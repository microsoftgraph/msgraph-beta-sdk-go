package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkSoftwareUpdateHealthable 
type TeamworkSoftwareUpdateHealthable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAdminAgentSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    GetCompanyPortalSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    GetFirmwareSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    GetOperatingSystemSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    GetPartnerAgentSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    GetTeamsClientSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    SetAdminAgentSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
    SetCompanyPortalSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
    SetFirmwareSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
    SetOperatingSystemSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
    SetPartnerAgentSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
    SetTeamsClientSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
}
