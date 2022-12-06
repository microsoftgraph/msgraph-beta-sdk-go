package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationValueDecimal 
type GroupPolicyPresentationValueDecimal struct {
    GroupPolicyPresentationValue
    // An unsigned integer value for the associated presentation.
    value *int64
}
// NewGroupPolicyPresentationValueDecimal instantiates a new GroupPolicyPresentationValueDecimal and sets the default values.
func NewGroupPolicyPresentationValueDecimal()(*GroupPolicyPresentationValueDecimal) {
    m := &GroupPolicyPresentationValueDecimal{
        GroupPolicyPresentationValue: *NewGroupPolicyPresentationValue(),
    }
    return m
}
// CreateGroupPolicyPresentationValueDecimalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationValueDecimalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationValueDecimal(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationValueDecimal) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyPresentationValue.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetValue)
    return res
}
// GetValue gets the value property value. An unsigned integer value for the associated presentation.
func (m *GroupPolicyPresentationValueDecimal) GetValue()(*int64) {
    return m.value
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationValueDecimal) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyPresentationValue.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. An unsigned integer value for the associated presentation.
func (m *GroupPolicyPresentationValueDecimal) SetValue(value *int64)() {
    m.value = value
}
