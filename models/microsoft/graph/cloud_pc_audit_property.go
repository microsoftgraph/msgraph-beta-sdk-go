package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcAuditProperty 
type CloudPcAuditProperty struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Display name.
    displayName *string;
    // New value.
    newValue *string;
    // Old value.
    oldValue *string;
}
// NewCloudPcAuditProperty instantiates a new cloudPcAuditProperty and sets the default values.
func NewCloudPcAuditProperty()(*CloudPcAuditProperty) {
    m := &CloudPcAuditProperty{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcAuditProperty) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. Display name.
func (m *CloudPcAuditProperty) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetNewValue gets the newValue property value. New value.
func (m *CloudPcAuditProperty) GetNewValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.newValue
    }
}
// GetOldValue gets the oldValue property value. Old value.
func (m *CloudPcAuditProperty) GetOldValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.oldValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcAuditProperty) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["newValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewValue(val)
        }
        return nil
    }
    res["oldValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOldValue(val)
        }
        return nil
    }
    return res
}
func (m *CloudPcAuditProperty) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CloudPcAuditProperty) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("newValue", m.GetNewValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("oldValue", m.GetOldValue())
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
func (m *CloudPcAuditProperty) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. Display name.
func (m *CloudPcAuditProperty) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetNewValue sets the newValue property value. New value.
func (m *CloudPcAuditProperty) SetNewValue(value *string)() {
    if m != nil {
        m.newValue = value
    }
}
// SetOldValue sets the oldValue property value. Old value.
func (m *CloudPcAuditProperty) SetOldValue(value *string)() {
    if m != nil {
        m.oldValue = value
    }
}
