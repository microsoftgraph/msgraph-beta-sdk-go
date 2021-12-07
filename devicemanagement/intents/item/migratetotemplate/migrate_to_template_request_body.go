package migratetotemplate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MigrateToTemplateRequestBody 
type MigrateToTemplateRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    newTemplateId *string;
    // 
    preserveCustomValues *bool;
}
// NewMigrateToTemplateRequestBody instantiates a new migrateToTemplateRequestBody and sets the default values.
func NewMigrateToTemplateRequestBody()(*MigrateToTemplateRequestBody) {
    m := &MigrateToTemplateRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MigrateToTemplateRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetNewTemplateId gets the newTemplateId property value. 
func (m *MigrateToTemplateRequestBody) GetNewTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.newTemplateId
    }
}
// GetPreserveCustomValues gets the preserveCustomValues property value. 
func (m *MigrateToTemplateRequestBody) GetPreserveCustomValues()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.preserveCustomValues
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MigrateToTemplateRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["newTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewTemplateId(val)
        }
        return nil
    }
    res["preserveCustomValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreserveCustomValues(val)
        }
        return nil
    }
    return res
}
func (m *MigrateToTemplateRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MigrateToTemplateRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("newTemplateId", m.GetNewTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("preserveCustomValues", m.GetPreserveCustomValues())
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
func (m *MigrateToTemplateRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetNewTemplateId sets the newTemplateId property value. 
func (m *MigrateToTemplateRequestBody) SetNewTemplateId(value *string)() {
    if m != nil {
        m.newTemplateId = value
    }
}
// SetPreserveCustomValues sets the preserveCustomValues property value. 
func (m *MigrateToTemplateRequestBody) SetPreserveCustomValues(value *bool)() {
    if m != nil {
        m.preserveCustomValues = value
    }
}
