package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationComboBox 
type GroupPolicyPresentationComboBox struct {
    GroupPolicyPresentation
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
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
        GroupPolicyPresentation: *NewGroupPolicyPresentation(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGroupPolicyPresentationComboBoxFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationComboBoxFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationComboBox(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupPolicyPresentationComboBox) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefaultValue gets the defaultValue property value. Localized default string displayed in the combo box. The default value is empty.
func (m *GroupPolicyPresentationComboBox) GetDefaultValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationComboBox) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyPresentation.GetFieldDeserializers()
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
                res[i] = *(v.(*string))
            }
            m.SetSuggestions(res)
        }
        return nil
    }
    return res
}
// GetMaxLength gets the maxLength property value. An unsigned integer that specifies the maximum number of text characters for the parameter. The default value is 1023.
func (m *GroupPolicyPresentationComboBox) GetMaxLength()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.maxLength
    }
}
// GetRequired gets the required property value. Specifies whether a value must be specified for the parameter. The default value is false.
func (m *GroupPolicyPresentationComboBox) GetRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.required
    }
}
// GetSuggestions gets the suggestions property value. Localized strings listed in the drop-down list of the combo box. The default value is empty.
func (m *GroupPolicyPresentationComboBox) GetSuggestions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.suggestions
    }
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationComboBox) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyPresentation.Serialize(writer)
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupPolicyPresentationComboBox) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDefaultValue sets the defaultValue property value. Localized default string displayed in the combo box. The default value is empty.
func (m *GroupPolicyPresentationComboBox) SetDefaultValue(value *string)() {
    if m != nil {
        m.defaultValue = value
    }
}
// SetMaxLength sets the maxLength property value. An unsigned integer that specifies the maximum number of text characters for the parameter. The default value is 1023.
func (m *GroupPolicyPresentationComboBox) SetMaxLength(value *int64)() {
    if m != nil {
        m.maxLength = value
    }
}
// SetRequired sets the required property value. Specifies whether a value must be specified for the parameter. The default value is false.
func (m *GroupPolicyPresentationComboBox) SetRequired(value *bool)() {
    if m != nil {
        m.required = value
    }
}
// SetSuggestions sets the suggestions property value. Localized strings listed in the drop-down list of the combo box. The default value is empty.
func (m *GroupPolicyPresentationComboBox) SetSuggestions(value []string)() {
    if m != nil {
        m.suggestions = value
    }
}
