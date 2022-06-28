package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationDecimalTextBox 
type GroupPolicyPresentationDecimalTextBox struct {
    GroupPolicyPresentation
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
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
        GroupPolicyPresentation: *NewGroupPolicyPresentation(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGroupPolicyPresentationDecimalTextBoxFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationDecimalTextBoxFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationDecimalTextBox(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupPolicyPresentationDecimalTextBox) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefaultValue gets the defaultValue property value. An unsigned integer that specifies the initial value for the decimal text box. The default value is 1.
func (m *GroupPolicyPresentationDecimalTextBox) GetDefaultValue()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.defaultValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationDecimalTextBox) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyPresentation.GetFieldDeserializers()
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
    if m == nil {
        return nil
    } else {
        return m.maxValue
    }
}
// GetMinValue gets the minValue property value. An unsigned integer that specifies the minimum allowed value. The default value is 0.
func (m *GroupPolicyPresentationDecimalTextBox) GetMinValue()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.minValue
    }
}
// GetRequired gets the required property value. Requirement to enter a value in the parameter box. The default value is false.
func (m *GroupPolicyPresentationDecimalTextBox) GetRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.required
    }
}
// GetSpin gets the spin property value. If true, create a spin control; otherwise, create a text box for numeric entry. The default value is true.
func (m *GroupPolicyPresentationDecimalTextBox) GetSpin()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.spin
    }
}
// GetSpinStep gets the spinStep property value. An unsigned integer that specifies the increment of change for the spin control. The default value is 1.
func (m *GroupPolicyPresentationDecimalTextBox) GetSpinStep()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.spinStep
    }
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationDecimalTextBox) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyPresentation.Serialize(writer)
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupPolicyPresentationDecimalTextBox) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDefaultValue sets the defaultValue property value. An unsigned integer that specifies the initial value for the decimal text box. The default value is 1.
func (m *GroupPolicyPresentationDecimalTextBox) SetDefaultValue(value *int64)() {
    if m != nil {
        m.defaultValue = value
    }
}
// SetMaxValue sets the maxValue property value. An unsigned integer that specifies the maximum allowed value. The default value is 9999.
func (m *GroupPolicyPresentationDecimalTextBox) SetMaxValue(value *int64)() {
    if m != nil {
        m.maxValue = value
    }
}
// SetMinValue sets the minValue property value. An unsigned integer that specifies the minimum allowed value. The default value is 0.
func (m *GroupPolicyPresentationDecimalTextBox) SetMinValue(value *int64)() {
    if m != nil {
        m.minValue = value
    }
}
// SetRequired sets the required property value. Requirement to enter a value in the parameter box. The default value is false.
func (m *GroupPolicyPresentationDecimalTextBox) SetRequired(value *bool)() {
    if m != nil {
        m.required = value
    }
}
// SetSpin sets the spin property value. If true, create a spin control; otherwise, create a text box for numeric entry. The default value is true.
func (m *GroupPolicyPresentationDecimalTextBox) SetSpin(value *bool)() {
    if m != nil {
        m.spin = value
    }
}
// SetSpinStep sets the spinStep property value. An unsigned integer that specifies the increment of change for the spin control. The default value is 1.
func (m *GroupPolicyPresentationDecimalTextBox) SetSpinStep(value *int64)() {
    if m != nil {
        m.spinStep = value
    }
}
