package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationTextBox represents an ADMX textBox element and an ADMX text element.
type GroupPolicyPresentationTextBox struct {
    GroupPolicyUploadedPresentation
}
// NewGroupPolicyPresentationTextBox instantiates a new GroupPolicyPresentationTextBox and sets the default values.
func NewGroupPolicyPresentationTextBox()(*GroupPolicyPresentationTextBox) {
    m := &GroupPolicyPresentationTextBox{
        GroupPolicyUploadedPresentation: *NewGroupPolicyUploadedPresentation(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyPresentationTextBox"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGroupPolicyPresentationTextBoxFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupPolicyPresentationTextBoxFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationTextBox(), nil
}
// GetDefaultValue gets the defaultValue property value. Localized default string displayed in the text box. The default value is empty.
// returns a *string when successful
func (m *GroupPolicyPresentationTextBox) GetDefaultValue()(*string) {
    val, err := m.GetBackingStore().Get("defaultValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupPolicyPresentationTextBox) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyUploadedPresentation.GetFieldDeserializers()
    res["defaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val)
        }
        return nil
    }
    res["maxLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxLength(val)
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
    return res
}
// GetMaxLength gets the maxLength property value. An unsigned integer that specifies the maximum number of text characters. Default value is 1023.
// returns a *int64 when successful
func (m *GroupPolicyPresentationTextBox) GetMaxLength()(*int64) {
    val, err := m.GetBackingStore().Get("maxLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetRequired gets the required property value. Requirement to enter a value in the text box. Default value is false.
// returns a *bool when successful
func (m *GroupPolicyPresentationTextBox) GetRequired()(*bool) {
    val, err := m.GetBackingStore().Get("required")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationTextBox) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyUploadedPresentation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("maxLength", m.GetMaxLength())
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
    return nil
}
// SetDefaultValue sets the defaultValue property value. Localized default string displayed in the text box. The default value is empty.
func (m *GroupPolicyPresentationTextBox) SetDefaultValue(value *string)() {
    err := m.GetBackingStore().Set("defaultValue", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxLength sets the maxLength property value. An unsigned integer that specifies the maximum number of text characters. Default value is 1023.
func (m *GroupPolicyPresentationTextBox) SetMaxLength(value *int64)() {
    err := m.GetBackingStore().Set("maxLength", value)
    if err != nil {
        panic(err)
    }
}
// SetRequired sets the required property value. Requirement to enter a value in the text box. Default value is false.
func (m *GroupPolicyPresentationTextBox) SetRequired(value *bool)() {
    err := m.GetBackingStore().Set("required", value)
    if err != nil {
        panic(err)
    }
}
type GroupPolicyPresentationTextBoxable interface {
    GroupPolicyUploadedPresentationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultValue()(*string)
    GetMaxLength()(*int64)
    GetRequired()(*bool)
    SetDefaultValue(value *string)()
    SetMaxLength(value *int64)()
    SetRequired(value *bool)()
}
