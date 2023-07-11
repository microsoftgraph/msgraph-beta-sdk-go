package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserMatchTargetReferenceValue 
type UserMatchTargetReferenceValue struct {
    ReferenceValue
    // The OdataType property
    OdataType *string
}
// NewUserMatchTargetReferenceValue instantiates a new userMatchTargetReferenceValue and sets the default values.
func NewUserMatchTargetReferenceValue()(*UserMatchTargetReferenceValue) {
    m := &UserMatchTargetReferenceValue{
        ReferenceValue: *NewReferenceValue(),
    }
    odataTypeValue := "#microsoft.graph.industryData.userMatchTargetReferenceValue"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUserMatchTargetReferenceValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserMatchTargetReferenceValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserMatchTargetReferenceValue(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserMatchTargetReferenceValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ReferenceValue.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *UserMatchTargetReferenceValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ReferenceValue.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// UserMatchTargetReferenceValueable 
type UserMatchTargetReferenceValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ReferenceValueable
}
