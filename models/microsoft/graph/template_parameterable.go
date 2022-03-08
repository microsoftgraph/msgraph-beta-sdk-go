package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// TemplateParameterable 
type TemplateParameterable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetJsonAllowedValues()(*string)
    GetJsonDefaultValue()(*string)
    GetValueType()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetJsonAllowedValues(value *string)()
    SetJsonDefaultValue(value *string)()
    SetValueType(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType)()
}
