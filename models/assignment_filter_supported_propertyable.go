package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignmentFilterSupportedPropertyable 
type AssignmentFilterSupportedPropertyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDataType()(*string)
    GetIsCollection()(*bool)
    GetName()(*string)
    GetOdataType()(*string)
    GetPropertyRegexConstraint()(*string)
    GetSupportedOperators()([]AssignmentFilterOperator)
    GetSupportedValues()([]string)
    SetDataType(value *string)()
    SetIsCollection(value *bool)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetPropertyRegexConstraint(value *string)()
    SetSupportedOperators(value []AssignmentFilterOperator)()
    SetSupportedValues(value []string)()
}
