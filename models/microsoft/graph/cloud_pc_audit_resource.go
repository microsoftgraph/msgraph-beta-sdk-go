package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CloudPcAuditResource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The resource entity display name.
    displayName *string;
    // A list of modified properties.
    modifiedProperties []CloudPcAuditProperty;
    // The ID of the audit resource.
    resourceId *string;
    // The type of the audit resource.
    type_escaped *string;
}
// Instantiates a new cloudPcAuditResource and sets the default values.
func NewCloudPcAuditResource()(*CloudPcAuditResource) {
    m := &CloudPcAuditResource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcAuditResource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the displayName property value. The resource entity display name.
func (m *CloudPcAuditResource) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the modifiedProperties property value. A list of modified properties.
func (m *CloudPcAuditResource) GetModifiedProperties()([]CloudPcAuditProperty) {
    if m == nil {
        return nil
    } else {
        return m.modifiedProperties
    }
}
// Gets the resourceId property value. The ID of the audit resource.
func (m *CloudPcAuditResource) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// Gets the type_escaped property value. The type of the audit resource.
func (m *CloudPcAuditResource) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
func (m *CloudPcAuditResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["modifiedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcAuditProperty() })
        if err != nil {
            return err
        }
        res := make([]CloudPcAuditProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*CloudPcAuditProperty))
        }
        m.SetModifiedProperties(res)
        return nil
    }
    res["resourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceId(val)
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
func (m *CloudPcAuditResource) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CloudPcAuditResource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetModifiedProperties()))
        for i, v := range m.GetModifiedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("modifiedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceId", m.GetResourceId())
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
func (m *CloudPcAuditResource) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the displayName property value. The resource entity display name.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *CloudPcAuditResource) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the modifiedProperties property value. A list of modified properties.
// Parameters:
//  - value : Value to set for the modifiedProperties property.
func (m *CloudPcAuditResource) SetModifiedProperties(value []CloudPcAuditProperty)() {
    m.modifiedProperties = value
}
// Sets the resourceId property value. The ID of the audit resource.
// Parameters:
//  - value : Value to set for the resourceId property.
func (m *CloudPcAuditResource) SetResourceId(value *string)() {
    m.resourceId = value
}
// Sets the type_escaped property value. The type of the audit resource.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *CloudPcAuditResource) SetType_escaped(value *string)() {
    m.type_escaped = value
}
