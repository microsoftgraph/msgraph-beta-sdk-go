package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcAuditResource 
type CloudPcAuditResource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The resource entity display name.
    displayName *string;
    // A list of modified properties.
    modifiedProperties []CloudPcAuditPropertyable;
    // The ID of the audit resource.
    resourceId *string;
    // The type of the audit resource.
    type_escaped *string;
}
// NewCloudPcAuditResource instantiates a new cloudPcAuditResource and sets the default values.
func NewCloudPcAuditResource()(*CloudPcAuditResource) {
    m := &CloudPcAuditResource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcAuditResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcAuditResourceFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCloudPcAuditResource(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcAuditResource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. The resource entity display name.
func (m *CloudPcAuditResource) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcAuditResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["modifiedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcAuditPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcAuditPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcAuditPropertyable)
            }
            m.SetModifiedProperties(res)
        }
        return nil
    }
    res["resourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetModifiedProperties gets the modifiedProperties property value. A list of modified properties.
func (m *CloudPcAuditResource) GetModifiedProperties()([]CloudPcAuditPropertyable) {
    if m == nil {
        return nil
    } else {
        return m.modifiedProperties
    }
}
// GetResourceId gets the resourceId property value. The ID of the audit resource.
func (m *CloudPcAuditResource) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// GetType gets the type property value. The type of the audit resource.
func (m *CloudPcAuditResource) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *CloudPcAuditResource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetModifiedProperties() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetModifiedProperties()))
        for i, v := range m.GetModifiedProperties() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *CloudPcAuditResource) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. The resource entity display name.
func (m *CloudPcAuditResource) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetModifiedProperties sets the modifiedProperties property value. A list of modified properties.
func (m *CloudPcAuditResource) SetModifiedProperties(value []CloudPcAuditPropertyable)() {
    if m != nil {
        m.modifiedProperties = value
    }
}
// SetResourceId sets the resourceId property value. The ID of the audit resource.
func (m *CloudPcAuditResource) SetResourceId(value *string)() {
    if m != nil {
        m.resourceId = value
    }
}
// SetType sets the type property value. The type of the audit resource.
func (m *CloudPcAuditResource) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
