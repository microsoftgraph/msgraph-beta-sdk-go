package changedeploymentstatus

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChangeDeploymentStatusRequestBodyable 
type ChangeDeploymentStatusRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetManagementTemplateStepId()(*string)
    GetStatus()(*string)
    GetTenantId()(*string)
    SetManagementTemplateStepId(value *string)()
    SetStatus(value *string)()
    SetTenantId(value *string)()
}
