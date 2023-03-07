package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationDecimalTextBox 
type GroupPolicyPresentationDecimalTextBox struct {
    GroupPolicyUploadedPresentation
}
// NewGroupPolicyPresentationDecimalTextBox instantiates a new GroupPolicyPresentationDecimalTextBox and sets the default values.
func NewGroupPolicyPresentationDecimalTextBox()(*GroupPolicyPresentationDecimalTextBox) {
    m := &GroupPolicyPresentationDecimalTextBox{
        GroupPolicyUploadedPresentation: *NewGroupPolicyUploadedPresentation(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyPresentationDecimalTextBox"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGroupPolicyPresentationDecimalTextBoxFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationDecimalTextBoxFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationDecimalTextBox(), nil
}
// GetDefaultValue gets the defaultValue property value. An unsigned integer that specifies the initial value for the decimal text box. The default value is 1.
func (m *GroupPolicyPresentationDecimalTextBox) GetDefaultValue()(*int64) {
    val, err := m.GetBackingStore().Get("defaultValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationDecimalTextBox) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyUploadedPresentation.GetFieldDeserializers()
    res["defaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val)
        }
        return nil
    }
    res["maxValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxValue(val)
        }
        return nil
    }
    res["minValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinValue(val)
        }
        return nil
    }
    res["required"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequired(val)
        }
        return nil
    }
    res["spin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpin(val)
        }
        return nil
    }
    res["spinStep"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpinStep(val)
        }
        return nil
    }
    return res
}
// GetMaxValue gets the maxValue property value. An unsigned integer that specifies the maximum allowed value. The default value is 9999.
func (m *GroupPolicyPresentationDecimalTextBox) GetMaxValue()(*int64) {
    val, err := m.GetBackingStore().Get("maxValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetMinValue gets the minValue property value. An unsigned integer that specifies the minimum allowed value. The default value is 0.
func (m *GroupPolicyPresentationDecimalTextBox) GetMinValue()(*int64) {
    val, err := m.GetBackingStore().Get("minValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetRequired gets the required property value. Requirement to enter a value in the parameter box. The default value is false.
func (m *GroupPolicyPresentationDecimalTextBox) GetRequired()(*bool) {
    val, err := m.GetBackingStore().Get("required")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSpin gets the spin property value. If true, create a spin control; otherwise, create a text box for numeric entry. The default value is true.
func (m *GroupPolicyPresentationDecimalTextBox) GetSpin()(*bool) {
    val, err := m.GetBackingStore().Get("spin")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSpinStep gets the spinStep property value. An unsigned integer that specifies the increment of change for the spin control. The default value is 1.
func (m *GroupPolicyPresentationDecimalTextBox) GetSpinStep()(*int64) {
    val, err := m.GetBackingStore().Get("spinStep")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationDecimalTextBox) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyUploadedPresentation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("maxValue", m.GetMaxValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("minValue", m.GetMinValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("required", m.GetRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("spin", m.GetSpin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("spinStep", m.GetSpinStep())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultValue sets the defaultValue property value. An unsigned integer that specifies the initial value for the decimal text box. The default value is 1.
func (m *GroupPolicyPresentationDecimalTextBox) SetDefaultValue(value *int64)() {
    err := m.GetBackingStore().Set("defaultValue", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxValue sets the maxValue property value. An unsigned integer that specifies the maximum allowed value. The default value is 9999.
func (m *GroupPolicyPresentationDecimalTextBox) SetMaxValue(value *int64)() {
    err := m.GetBackingStore().Set("maxValue", value)
    if err != nil {
        panic(err)
    }
}
// SetMinValue sets the minValue property value. An unsigned integer that specifies the minimum allowed value. The default value is 0.
func (m *GroupPolicyPresentationDecimalTextBox) SetMinValue(value *int64)() {
    err := m.GetBackingStore().Set("minValue", value)
    if err != nil {
        panic(err)
    }
}
// SetRequired sets the required property value. Requirement to enter a value in the parameter box. The default value is false.
func (m *GroupPolicyPresentationDecimalTextBox) SetRequired(value *bool)() {
    err := m.GetBackingStore().Set("required", value)
    if err != nil {
        panic(err)
    }
}
// SetSpin sets the spin property value. If true, create a spin control; otherwise, create a text box for numeric entry. The default value is true.
func (m *GroupPolicyPresentationDecimalTextBox) SetSpin(value *bool)() {
    err := m.GetBackingStore().Set("spin", value)
    if err != nil {
        panic(err)
    }
}
// SetSpinStep sets the spinStep property value. An unsigned integer that specifies the increment of change for the spin control. The default value is 1.
func (m *GroupPolicyPresentationDecimalTextBox) SetSpinStep(value *int64)() {
    err := m.GetBackingStore().Set("spinStep", value)
    if err != nil {
        panic(err)
    }
}
// GroupPolicyPresentationDecimalTextBoxable 
type GroupPolicyPresentationDecimalTextBoxable interface {
    GroupPolicyUploadedPresentationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultValue()(*int64)
    GetMaxValue()(*int64)
    GetMinValue()(*int64)
    GetRequired()(*bool)
    GetSpin()(*bool)
    GetSpinStep()(*int64)
    SetDefaultValue(value *int64)()
    SetMaxValue(value *int64)()
    SetMinValue(value *int64)()
    SetRequired(value *bool)()
    SetSpin(value *bool)()
    SetSpinStep(value *int64)()
}
