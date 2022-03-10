package updatepriorities

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UpdatePrioritiesRequestBodyable 
type UpdatePrioritiesRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetOfficeConfigurationPolicyIds()([]string)
    GetOfficeConfigurationPriorities()([]int32)
    SetOfficeConfigurationPolicyIds(value []string)()
    SetOfficeConfigurationPriorities(value []int32)()
}
