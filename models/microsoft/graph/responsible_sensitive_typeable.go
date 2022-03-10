package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ResponsibleSensitiveTypeable 
type ResponsibleSensitiveTypeable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDescription()(*string)
    GetId()(*string)
    GetName()(*string)
    GetPublisherName()(*string)
    GetRulePackageId()(*string)
    GetRulePackageType()(*string)
    SetDescription(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetPublisherName(value *string)()
    SetRulePackageId(value *string)()
    SetRulePackageType(value *string)()
}
