package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationValueLongDecimal the entity represents an unsigned long value of a long decimal text box presentation on a policy definition.
type GroupPolicyPresentationValueLongDecimal struct {
    GroupPolicyPresentationValue
    // The OdataType property
    OdataType *string
}
// NewGroupPolicyPresentationValueLongDecimal instantiates a new groupPolicyPresentationValueLongDecimal and sets the default values.
func NewGroupPolicyPresentationValueLongDecimal()(*GroupPolicyPresentationValueLongDecimal) {
    m := &GroupPolicyPresentationValueLongDecimal{
        GroupPolicyPresentationValue: *NewGroupPolicyPresentationValue(),
    }
    return m
}
// CreateGroupPolicyPresentationValueLongDecimalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationValueLongDecimalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationValueLongDecimal(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationValueLongDecimal) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyPresentationValue.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. An unsigned long value for the associated presentation.
func (m *GroupPolicyPresentationValueLongDecimal) GetValue()(*int64) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationValueLongDecimal) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetValue sets the value property value. An unsigned long value for the associated presentation.
func (m *GroupPolicyPresentationValueLongDecimal) SetValue(value *int64)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// GroupPolicyPresentationValueLongDecimalable 
type GroupPolicyPresentationValueLongDecimalable interface {
    GroupPolicyPresentationValueable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()(*int64)
    SetValue(value *int64)()
}
