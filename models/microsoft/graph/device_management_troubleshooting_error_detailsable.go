package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementTroubleshootingErrorDetailsable 
type DeviceManagementTroubleshootingErrorDetailsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetContext()(*string)
    GetFailure()(*string)
    GetFailureDetails()(*string)
    GetRemediation()(*string)
    GetResources()([]DeviceManagementTroubleshootingErrorResourceable)
    SetContext(value *string)()
    SetFailure(value *string)()
    SetFailureDetails(value *string)()
    SetRemediation(value *string)()
    SetResources(value []DeviceManagementTroubleshootingErrorResourceable)()
}
