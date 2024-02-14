package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationMultiTextBox represents an ADMX multiTextBox element and an ADMX multiText element.
type GroupPolicyPresentationMultiTextBox struct {
    GroupPolicyUploadedPresentation
}
// NewGroupPolicyPresentationMultiTextBox instantiates a new GroupPolicyPresentationMultiTextBox and sets the default values.
func NewGroupPolicyPresentationMultiTextBox()(*GroupPolicyPresentationMultiTextBox) {
    m := &GroupPolicyPresentationMultiTextBox{
        GroupPolicyUploadedPresentation: *NewGroupPolicyUploadedPresentation(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyPresentationMultiTextBox"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGroupPolicyPresentationMultiTextBoxFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupPolicyPresentationMultiTextBoxFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationMultiTextBox(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupPolicyPresentationMultiTextBox) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyUploadedPresentation.GetFieldDeserializers()
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
    res["maxStrings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxStrings(val)
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
func (m *GroupPolicyPresentationMultiTextBox) GetMaxLength()(*int64) {
    val, err := m.GetBackingStore().Get("maxLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetMaxStrings gets the maxStrings property value. An unsigned integer that specifies the maximum number of strings. Default value is 0.
// returns a *int64 when successful
func (m *GroupPolicyPresentationMultiTextBox) GetMaxStrings()(*int64) {
    val, err := m.GetBackingStore().Get("maxStrings")
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
func (m *GroupPolicyPresentationMultiTextBox) GetRequired()(*bool) {
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
func (m *GroupPolicyPresentationMultiTextBox) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyUploadedPresentation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("maxLength", m.GetMaxLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("maxStrings", m.GetMaxStrings())
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
// SetMaxLength sets the maxLength property value. An unsigned integer that specifies the maximum number of text characters. Default value is 1023.
func (m *GroupPolicyPresentationMultiTextBox) SetMaxLength(value *int64)() {
    err := m.GetBackingStore().Set("maxLength", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxStrings sets the maxStrings property value. An unsigned integer that specifies the maximum number of strings. Default value is 0.
func (m *GroupPolicyPresentationMultiTextBox) SetMaxStrings(value *int64)() {
    err := m.GetBackingStore().Set("maxStrings", value)
    if err != nil {
        panic(err)
    }
}
// SetRequired sets the required property value. Requirement to enter a value in the text box. Default value is false.
func (m *GroupPolicyPresentationMultiTextBox) SetRequired(value *bool)() {
    err := m.GetBackingStore().Set("required", value)
    if err != nil {
        panic(err)
    }
}
type GroupPolicyPresentationMultiTextBoxable interface {
    GroupPolicyUploadedPresentationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaxLength()(*int64)
    GetMaxStrings()(*int64)
    GetRequired()(*bool)
    SetMaxLength(value *int64)()
    SetMaxStrings(value *int64)()
    SetRequired(value *bool)()
}
