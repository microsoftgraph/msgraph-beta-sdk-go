package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernanceNotificationTemplate 
type GovernanceNotificationTemplate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    culture *string;
    // 
    id *string;
    // 
    source *string;
    // 
    type_escaped *string;
    // 
    version *string;
}
// NewGovernanceNotificationTemplate instantiates a new governanceNotificationTemplate and sets the default values.
func NewGovernanceNotificationTemplate()(*GovernanceNotificationTemplate) {
    m := &GovernanceNotificationTemplate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernanceNotificationTemplate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCulture gets the culture property value. 
func (m *GovernanceNotificationTemplate) GetCulture()(*string) {
    if m == nil {
        return nil
    } else {
        return m.culture
    }
}
// GetId gets the id property value. 
func (m *GovernanceNotificationTemplate) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetSource gets the source property value. 
func (m *GovernanceNotificationTemplate) GetSource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// GetType gets the type property value. 
func (m *GovernanceNotificationTemplate) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetVersion gets the version property value. 
func (m *GovernanceNotificationTemplate) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceNotificationTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["culture"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCulture(val)
        }
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val)
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
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
func (m *GovernanceNotificationTemplate) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GovernanceNotificationTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("culture", m.GetCulture())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("source", m.GetSource())
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
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *GovernanceNotificationTemplate) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCulture sets the culture property value. 
func (m *GovernanceNotificationTemplate) SetCulture(value *string)() {
    if m != nil {
        m.culture = value
    }
}
// SetId sets the id property value. 
func (m *GovernanceNotificationTemplate) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetSource sets the source property value. 
func (m *GovernanceNotificationTemplate) SetSource(value *string)() {
    if m != nil {
        m.source = value
    }
}
// SetType sets the type property value. 
func (m *GovernanceNotificationTemplate) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetVersion sets the version property value. 
func (m *GovernanceNotificationTemplate) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
