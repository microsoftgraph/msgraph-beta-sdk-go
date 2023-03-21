package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleReferenceValue 
type RoleReferenceValue struct {
    ReferenceValue
}
// NewRoleReferenceValue instantiates a new RoleReferenceValue and sets the default values.
func NewRoleReferenceValue()(*RoleReferenceValue) {
    m := &RoleReferenceValue{
        ReferenceValue: *NewReferenceValue(),
    }
    odataTypeValue := "#microsoft.graph.industryData.roleReferenceValue"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRoleReferenceValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleReferenceValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleReferenceValue(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleReferenceValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ReferenceValue.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RoleReferenceValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ReferenceValue.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// RoleReferenceValueable 
type RoleReferenceValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ReferenceValueable
}
