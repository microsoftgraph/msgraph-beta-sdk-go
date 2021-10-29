package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GovernanceNotificationPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    enabledTemplateTypes []string;
    // 
    notificationTemplates []GovernanceNotificationTemplate;
}
// Instantiates a new governanceNotificationPolicy and sets the default values.
func NewGovernanceNotificationPolicy()(*GovernanceNotificationPolicy) {
    m := &GovernanceNotificationPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernanceNotificationPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the enabledTemplateTypes property value. 
func (m *GovernanceNotificationPolicy) GetEnabledTemplateTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enabledTemplateTypes
    }
}
// Gets the notificationTemplates property value. 
func (m *GovernanceNotificationPolicy) GetNotificationTemplates()([]GovernanceNotificationTemplate) {
    if m == nil {
        return nil
    } else {
        return m.notificationTemplates
    }
}
// The deserialization information for the current model
func (m *GovernanceNotificationPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enabledTemplateTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetEnabledTemplateTypes(res)
        return nil
    }
    res["notificationTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceNotificationTemplate() })
        if err != nil {
            return err
        }
        res := make([]GovernanceNotificationTemplate, len(val))
        for i, v := range val {
            res[i] = *(v.(*GovernanceNotificationTemplate))
        }
        m.SetNotificationTemplates(res)
        return nil
    }
    return res
}
func (m *GovernanceNotificationPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GovernanceNotificationPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("enabledTemplateTypes", m.GetEnabledTemplateTypes())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNotificationTemplates()))
        for i, v := range m.GetNotificationTemplates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("notificationTemplates", cast)
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
func (m *GovernanceNotificationPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the enabledTemplateTypes property value. 
// Parameters:
//  - value : Value to set for the enabledTemplateTypes property.
func (m *GovernanceNotificationPolicy) SetEnabledTemplateTypes(value []string)() {
    m.enabledTemplateTypes = value
}
// Sets the notificationTemplates property value. 
// Parameters:
//  - value : Value to set for the notificationTemplates property.
func (m *GovernanceNotificationPolicy) SetNotificationTemplates(value []GovernanceNotificationTemplate)() {
    m.notificationTemplates = value
}
