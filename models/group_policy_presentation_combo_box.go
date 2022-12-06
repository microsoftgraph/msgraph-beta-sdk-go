package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationComboBox 
type GroupPolicyPresentationComboBox struct {
    GroupPolicyUploadedPresentation
    // Localized default string displayed in the combo box. The default value is empty.
    defaultValue *string
    // An unsigned integer that specifies the maximum number of text characters for the parameter. The default value is 1023.
    maxLength *int64
    // Specifies whether a value must be specified for the parameter. The default value is false.
    required *bool
    // Localized strings listed in the drop-down list of the combo box. The default value is empty.
    suggestions []string
}
// NewGroupPolicyPresentationComboBox instantiates a new GroupPolicyPresentationComboBox and sets the default values.
func NewGroupPolicyPresentationComboBox()(*GroupPolicyPresentationComboBox) {
    m := &GroupPolicyPresentationComboBox{
        GroupPolicyUploadedPresentation: *NewGroupPolicyUploadedPresentation(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyPresentationComboBox";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupPolicyPresentationComboBoxFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationComboBoxFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationComboBox(), nil
}
// GetDefaultValue gets the defaultValue property value. Localized default string displayed in the combo box. The default value is empty.
func (m *GroupPolicyPresentationComboBox) GetDefaultValue()(*string) {
    return m.defaultValue
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationComboBox) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyUploadedPresentation.GetFieldDeserializers()
    res["defaultValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefaultValue)
    res["maxLength"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetMaxLength)
    res["required"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRequired)
    res["suggestions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetSuggestions)
    return res
}
// GetMaxLength gets the maxLength property value. An unsigned integer that specifies the maximum number of text characters for the parameter. The default value is 1023.
func (m *GroupPolicyPresentationComboBox) GetMaxLength()(*int64) {
    return m.maxLength
}
// GetRequired gets the required property value. Specifies whether a value must be specified for the parameter. The default value is false.
func (m *GroupPolicyPresentationComboBox) GetRequired()(*bool) {
    return m.required
}
// GetSuggestions gets the suggestions property value. Localized strings listed in the drop-down list of the combo box. The default value is empty.
func (m *GroupPolicyPresentationComboBox) GetSuggestions()([]string) {
    return m.suggestions
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
    m.defaultValue = value
}
// SetMaxLength sets the maxLength property value. An unsigned integer that specifies the maximum number of text characters for the parameter. The default value is 1023.
func (m *GroupPolicyPresentationComboBox) SetMaxLength(value *int64)() {
    m.maxLength = value
}
// SetRequired sets the required property value. Specifies whether a value must be specified for the parameter. The default value is false.
func (m *GroupPolicyPresentationComboBox) SetRequired(value *bool)() {
    m.required = value
}
// SetSuggestions sets the suggestions property value. Localized strings listed in the drop-down list of the combo box. The default value is empty.
func (m *GroupPolicyPresentationComboBox) SetSuggestions(value []string)() {
    m.suggestions = value
}
