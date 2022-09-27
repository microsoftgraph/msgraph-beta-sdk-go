package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationDecimalTextBox 
type GroupPolicyPresentationDecimalTextBox struct {
    GroupPolicyUploadedPresentation
    // An unsigned integer that specifies the initial value for the decimal text box. The default value is 1.
    defaultValue *int64
    // An unsigned integer that specifies the maximum allowed value. The default value is 9999.
    maxValue *int64
    // An unsigned integer that specifies the minimum allowed value. The default value is 0.
    minValue *int64
    // Requirement to enter a value in the parameter box. The default value is false.
    required *bool
    // If true, create a spin control; otherwise, create a text box for numeric entry. The default value is true.
    spin *bool
    // An unsigned integer that specifies the increment of change for the spin control. The default value is 1.
    spinStep *int64
}
// NewGroupPolicyPresentationDecimalTextBox instantiates a new GroupPolicyPresentationDecimalTextBox and sets the default values.
func NewGroupPolicyPresentationDecimalTextBox()(*GroupPolicyPresentationDecimalTextBox) {
    m := &GroupPolicyPresentationDecimalTextBox{
        GroupPolicyUploadedPresentation: *NewGroupPolicyUploadedPresentation(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyPresentationDecimalTextBox";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupPolicyPresentationDecimalTextBoxFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationDecimalTextBoxFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationDecimalTextBox(), nil
}
// GetDefaultValue gets the defaultValue property value. An unsigned integer that specifies the initial value for the decimal text box. The default value is 1.
func (m *GroupPolicyPresentationDecimalTextBox) GetDefaultValue()(*int64) {
    return m.defaultValue
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationDecimalTextBox) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyUploadedPresentation.GetFieldDeserializers()
    res["defaultValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetDefaultValue)
    res["maxValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetMaxValue)
    res["minValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetMinValue)
    res["required"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRequired)
    res["spin"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSpin)
    res["spinStep"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetSpinStep)
    return res
}
// GetMaxValue gets the maxValue property value. An unsigned integer that specifies the maximum allowed value. The default value is 9999.
func (m *GroupPolicyPresentationDecimalTextBox) GetMaxValue()(*int64) {
    return m.maxValue
}
// GetMinValue gets the minValue property value. An unsigned integer that specifies the minimum allowed value. The default value is 0.
func (m *GroupPolicyPresentationDecimalTextBox) GetMinValue()(*int64) {
    return m.minValue
}
// GetRequired gets the required property value. Requirement to enter a value in the parameter box. The default value is false.
func (m *GroupPolicyPresentationDecimalTextBox) GetRequired()(*bool) {
    return m.required
}
// GetSpin gets the spin property value. If true, create a spin control; otherwise, create a text box for numeric entry. The default value is true.
func (m *GroupPolicyPresentationDecimalTextBox) GetSpin()(*bool) {
    return m.spin
}
// GetSpinStep gets the spinStep property value. An unsigned integer that specifies the increment of change for the spin control. The default value is 1.
func (m *GroupPolicyPresentationDecimalTextBox) GetSpinStep()(*int64) {
    return m.spinStep
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
    m.defaultValue = value
}
// SetMaxValue sets the maxValue property value. An unsigned integer that specifies the maximum allowed value. The default value is 9999.
func (m *GroupPolicyPresentationDecimalTextBox) SetMaxValue(value *int64)() {
    m.maxValue = value
}
// SetMinValue sets the minValue property value. An unsigned integer that specifies the minimum allowed value. The default value is 0.
func (m *GroupPolicyPresentationDecimalTextBox) SetMinValue(value *int64)() {
    m.minValue = value
}
// SetRequired sets the required property value. Requirement to enter a value in the parameter box. The default value is false.
func (m *GroupPolicyPresentationDecimalTextBox) SetRequired(value *bool)() {
    m.required = value
}
// SetSpin sets the spin property value. If true, create a spin control; otherwise, create a text box for numeric entry. The default value is true.
func (m *GroupPolicyPresentationDecimalTextBox) SetSpin(value *bool)() {
    m.spin = value
}
// SetSpinStep sets the spinStep property value. An unsigned integer that specifies the increment of change for the spin control. The default value is 1.
func (m *GroupPolicyPresentationDecimalTextBox) SetSpinStep(value *int64)() {
    m.spinStep = value
}
