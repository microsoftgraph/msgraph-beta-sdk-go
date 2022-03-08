package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CustomExtensionHandlerInstanceable 
type CustomExtensionHandlerInstanceable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCustomExtensionId()(*string)
    GetExternalCorrelationId()(*string)
    GetStage()(*AccessPackageCustomExtensionStage)
    GetStatus()(*AccessPackageCustomExtensionHandlerStatus)
    SetCustomExtensionId(value *string)()
    SetExternalCorrelationId(value *string)()
    SetStage(value *AccessPackageCustomExtensionStage)()
    SetStatus(value *AccessPackageCustomExtensionHandlerStatus)()
}
