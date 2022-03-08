package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssignmentFilterSupportedPropertyable 
type AssignmentFilterSupportedPropertyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDataType()(*string)
    GetIsCollection()(*bool)
    GetName()(*string)
    GetPropertyRegexConstraint()(*string)
    GetSupportedOperators()([]AssignmentFilterOperator)
    GetSupportedValues()([]string)
    SetDataType(value *string)()
    SetIsCollection(value *bool)()
    SetName(value *string)()
    SetPropertyRegexConstraint(value *string)()
    SetSupportedOperators(value []AssignmentFilterOperator)()
    SetSupportedValues(value []string)()
}
