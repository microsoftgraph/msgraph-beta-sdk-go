package activateservice

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ActivateServiceRequestBodyable 
type ActivateServiceRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetService()(*string)
    GetServicePlanId()(*string)
    GetSkuId()(*string)
    SetService(value *string)()
    SetServicePlanId(value *string)()
    SetSkuId(value *string)()
}
