package importresourceactions

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ImportResourceActionsRequestBody 
type ImportResourceActionsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    format *string;
    // 
    overwriteResourceNamespace *bool;
    // 
    value *string;
}
// NewImportResourceActionsRequestBody instantiates a new importResourceActionsRequestBody and sets the default values.
func NewImportResourceActionsRequestBody()(*ImportResourceActionsRequestBody) {
    m := &ImportResourceActionsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImportResourceActionsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFormat gets the format property value. 
func (m *ImportResourceActionsRequestBody) GetFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetOverwriteResourceNamespace gets the overwriteResourceNamespace property value. 
func (m *ImportResourceActionsRequestBody) GetOverwriteResourceNamespace()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.overwriteResourceNamespace
    }
}
// GetValue gets the value property value. 
func (m *ImportResourceActionsRequestBody) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImportResourceActionsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val)
        }
        return nil
    }
    res["overwriteResourceNamespace"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverwriteResourceNamespace(val)
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
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
func (m *ImportResourceActionsRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ImportResourceActionsRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("overwriteResourceNamespace", m.GetOverwriteResourceNamespace())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImportResourceActionsRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFormat sets the format property value. 
func (m *ImportResourceActionsRequestBody) SetFormat(value *string)() {
    if m != nil {
        m.format = value
    }
}
// SetOverwriteResourceNamespace sets the overwriteResourceNamespace property value. 
func (m *ImportResourceActionsRequestBody) SetOverwriteResourceNamespace(value *bool)() {
    if m != nil {
        m.overwriteResourceNamespace = value
    }
}
// SetValue sets the value property value. 
func (m *ImportResourceActionsRequestBody) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}
