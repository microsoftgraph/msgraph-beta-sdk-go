package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationComboBox 
type GroupPolicyPresentationComboBox struct {
    GroupPolicyUploadedPresentation
}
// NewGroupPolicyPresentationComboBox instantiates a new GroupPolicyPresentationComboBox and sets the default values.
func NewGroupPolicyPresentationComboBox()(*GroupPolicyPresentationComboBox) {
    m := &GroupPolicyPresentationComboBox{
        GroupPolicyUploadedPresentation: *NewGroupPolicyUploadedPresentation(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyPresentationComboBox"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGroupPolicyPresentationComboBoxFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationComboBoxFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationComboBox(), nil
}
// GetDefaultValue gets the defaultValue property value. Localized default string displayed in the combo box. The default value is empty.
func (m *GroupPolicyPresentationComboBox) GetDefaultValue()(*string) {
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
func (m *GroupPolicyPresentationComboBox) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["suggestions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSuggestions(res)
        }
        return nil
    }
    return res
}
// GetMaxLength gets the maxLength property value. An unsigned integer that specifies the maximum number of text characters for the parameter. The default value is 1023.
func (m *GroupPolicyPresentationComboBox) GetMaxLength()(*int64) {
    val, err := m.GetBackingStore().Get("maxLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetRequired gets the required property value. Specifies whether a value must be specified for the parameter. The default value is false.
func (m *GroupPolicyPresentationComboBox) GetRequired()(*bool) {
    val, err := m.GetBackingStore().Get("required")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSuggestions gets the suggestions property value. Localized strings listed in the drop-down list of the combo box. The default value is empty.
func (m *GroupPolicyPresentationComboBox) GetSuggestions()([]string) {
    val, err := m.GetBackingStore().Get("suggestions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationComboBox) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetSuggestions() != nil {
        err = writer.WriteCollectionOfStringValues("suggestions", m.GetSuggestions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultValue sets the defaultValue property value. Localized default string displayed in the combo box. The default value is empty.
func (m *GroupPolicyPresentationComboBox) SetDefaultValue(value *string)() {
    err := m.GetBackingStore().Set("defaultValue", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxLength sets the maxLength property value. An unsigned integer that specifies the maximum number of text characters for the parameter. The default value is 1023.
func (m *GroupPolicyPresentationComboBox) SetMaxLength(value *int64)() {
    err := m.GetBackingStore().Set("maxLength", value)
    if err != nil {
        panic(err)
    }
}
// SetRequired sets the required property value. Specifies whether a value must be specified for the parameter. The default value is false.
func (m *GroupPolicyPresentationComboBox) SetRequired(value *bool)() {
    err := m.GetBackingStore().Set("required", value)
    if err != nil {
        panic(err)
    }
}
// SetSuggestions sets the suggestions property value. Localized strings listed in the drop-down list of the combo box. The default value is empty.
func (m *GroupPolicyPresentationComboBox) SetSuggestions(value []string)() {
    err := m.GetBackingStore().Set("suggestions", value)
    if err != nil {
        panic(err)
    }
}
// GroupPolicyPresentationComboBoxable 
type GroupPolicyPresentationComboBoxable interface {
    GroupPolicyUploadedPresentationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultValue()(*string)
    GetMaxLength()(*int64)
    GetRequired()(*bool)
    GetSuggestions()([]string)
    SetDefaultValue(value *string)()
    SetMaxLength(value *int64)()
    SetRequired(value *bool)()
    SetSuggestions(value []string)()
}
