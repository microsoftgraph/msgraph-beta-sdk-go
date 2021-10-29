package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PasswordSingleSignOnField struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Title/label override for customization.
    customizedLabel *string;
    // Label that would be used if no customizedLabel is provided. Read only.
    defaultLabel *string;
    // Id used to identity the field type. This is an internal id and possible values are param_1, param_2, param_userName, param_password.
    fieldId *string;
    // Type of the credential. The values can be text, password.
    type_escaped *string;
}
// Instantiates a new passwordSingleSignOnField and sets the default values.
func NewPasswordSingleSignOnField()(*PasswordSingleSignOnField) {
    m := &PasswordSingleSignOnField{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordSingleSignOnField) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the customizedLabel property value. Title/label override for customization.
func (m *PasswordSingleSignOnField) GetCustomizedLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customizedLabel
    }
}
// Gets the defaultLabel property value. Label that would be used if no customizedLabel is provided. Read only.
func (m *PasswordSingleSignOnField) GetDefaultLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultLabel
    }
}
// Gets the fieldId property value. Id used to identity the field type. This is an internal id and possible values are param_1, param_2, param_userName, param_password.
func (m *PasswordSingleSignOnField) GetFieldId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fieldId
    }
}
// Gets the type_escaped property value. Type of the credential. The values can be text, password.
func (m *PasswordSingleSignOnField) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
func (m *PasswordSingleSignOnField) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["customizedLabel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomizedLabel(val)
        return nil
    }
    res["defaultLabel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefaultLabel(val)
        return nil
    }
    res["fieldId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFieldId(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
        return nil
    }
    return res
}
func (m *PasswordSingleSignOnField) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PasswordSingleSignOnField) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("customizedLabel", m.GetCustomizedLabel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("defaultLabel", m.GetDefaultLabel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fieldId", m.GetFieldId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type_escaped", m.GetType_escaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *PasswordSingleSignOnField) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the customizedLabel property value. Title/label override for customization.
// Parameters:
//  - value : Value to set for the customizedLabel property.
func (m *PasswordSingleSignOnField) SetCustomizedLabel(value *string)() {
    m.customizedLabel = value
}
// Sets the defaultLabel property value. Label that would be used if no customizedLabel is provided. Read only.
// Parameters:
//  - value : Value to set for the defaultLabel property.
func (m *PasswordSingleSignOnField) SetDefaultLabel(value *string)() {
    m.defaultLabel = value
}
// Sets the fieldId property value. Id used to identity the field type. This is an internal id and possible values are param_1, param_2, param_userName, param_password.
// Parameters:
//  - value : Value to set for the fieldId property.
func (m *PasswordSingleSignOnField) SetFieldId(value *string)() {
    m.fieldId = value
}
// Sets the type_escaped property value. Type of the credential. The values can be text, password.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *PasswordSingleSignOnField) SetType_escaped(value *string)() {
    m.type_escaped = value
}
