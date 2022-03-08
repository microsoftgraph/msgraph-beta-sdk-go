package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CustomSecurityAttributeDefinitionable 
type CustomSecurityAttributeDefinitionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowedValues()([]AllowedValueable)
    GetAttributeSet()(*string)
    GetDescription()(*string)
    GetIsCollection()(*bool)
    GetIsSearchable()(*bool)
    GetName()(*string)
    GetStatus()(*string)
    GetType()(*string)
    GetUsePreDefinedValuesOnly()(*bool)
    SetAllowedValues(value []AllowedValueable)()
    SetAttributeSet(value *string)()
    SetDescription(value *string)()
    SetIsCollection(value *bool)()
    SetIsSearchable(value *bool)()
    SetName(value *string)()
    SetStatus(value *string)()
    SetType(value *string)()
    SetUsePreDefinedValuesOnly(value *bool)()
}
