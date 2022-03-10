package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CustomExtensionHandlerable 
type CustomExtensionHandlerable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCustomExtension()(CustomAccessPackageWorkflowExtensionable)
    GetStage()(*AccessPackageCustomExtensionStage)
    SetCustomExtension(value CustomAccessPackageWorkflowExtensionable)()
    SetStage(value *AccessPackageCustomExtensionStage)()
}
