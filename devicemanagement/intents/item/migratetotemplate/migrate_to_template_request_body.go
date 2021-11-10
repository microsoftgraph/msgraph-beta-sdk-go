package migratetotemplate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MigrateToTemplateRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    newTemplateId *string;
    // 
    preserveCustomValues *bool;
}
// Instantiates a new migrateToTemplateRequestBody and sets the default values.
func NewMigrateToTemplateRequestBody()(*MigrateToTemplateRequestBody) {
    m := &MigrateToTemplateRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MigrateToTemplateRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the newTemplateId property value. 
func (m *MigrateToTemplateRequestBody) GetNewTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.newTemplateId
    }
}
// Gets the preserveCustomValues property value. 
func (m *MigrateToTemplateRequestBody) GetPreserveCustomValues()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.preserveCustomValues
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *MigrateToTemplateRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the newTemplateId property value. 
// Parameters:
//  - value : Value to set for the newTemplateId property.
func (m *MigrateToTemplateRequestBody) SetNewTemplateId(value *string)() {
    m.newTemplateId = value
}
// Sets the preserveCustomValues property value. 
// Parameters:
//  - value : Value to set for the preserveCustomValues property.
func (m *MigrateToTemplateRequestBody) SetPreserveCustomValues(value *bool)() {
    m.preserveCustomValues = value
}
