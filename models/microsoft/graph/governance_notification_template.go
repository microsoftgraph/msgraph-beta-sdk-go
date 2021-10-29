package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new governanceNotificationTemplate and sets the default values.
func NewGovernanceNotificationTemplate()(*GovernanceNotificationTemplate) {
    m := &GovernanceNotificationTemplate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernanceNotificationTemplate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the culture property value. 
func (m *GovernanceNotificationTemplate) GetCulture()(*string) {
    if m == nil {
        return nil
    } else {
        return m.culture
    }
}
// Gets the id property value. 
func (m *GovernanceNotificationTemplate) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the source property value. 
func (m *GovernanceNotificationTemplate) GetSource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// Gets the type_escaped property value. 
func (m *GovernanceNotificationTemplate) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Gets the version property value. 
func (m *GovernanceNotificationTemplate) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *GovernanceNotificationTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["culture"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCulture(val)
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSource(val)
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
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    return res
}
func (m *GovernanceNotificationTemplate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        err := writer.WriteStringValue("type_escaped", m.GetType_escaped())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *GovernanceNotificationTemplate) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the culture property value. 
// Parameters:
//  - value : Value to set for the culture property.
func (m *GovernanceNotificationTemplate) SetCulture(value *string)() {
    m.culture = value
}
// Sets the id property value. 
// Parameters:
//  - value : Value to set for the id property.
func (m *GovernanceNotificationTemplate) SetId(value *string)() {
    m.id = value
}
// Sets the source property value. 
// Parameters:
//  - value : Value to set for the source property.
func (m *GovernanceNotificationTemplate) SetSource(value *string)() {
    m.source = value
}
// Sets the type_escaped property value. 
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *GovernanceNotificationTemplate) SetType_escaped(value *string)() {
    m.type_escaped = value
}
// Sets the version property value. 
// Parameters:
//  - value : Value to set for the version property.
func (m *GovernanceNotificationTemplate) SetVersion(value *string)() {
    m.version = value
}
