package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernanceNotificationPolicy provides operations to manage the collection of approvalWorkflowProvider entities.
type GovernanceNotificationPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    enabledTemplateTypes []string;
    // 
    notificationTemplates []GovernanceNotificationTemplateable;
}
// NewGovernanceNotificationPolicy instantiates a new governanceNotificationPolicy and sets the default values.
func NewGovernanceNotificationPolicy()(*GovernanceNotificationPolicy) {
    m := &GovernanceNotificationPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGovernanceNotificationPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernanceNotificationPolicyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewGovernanceNotificationPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernanceNotificationPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEnabledTemplateTypes gets the enabledTemplateTypes property value. 
func (m *GovernanceNotificationPolicy) GetEnabledTemplateTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enabledTemplateTypes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceNotificationPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enabledTemplateTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEnabledTemplateTypes(res)
        }
        return nil
    }
    res["notificationTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceNotificationTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceNotificationTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceNotificationTemplateable)
            }
            m.SetNotificationTemplates(res)
        }
        return nil
    }
    return res
}
// GetNotificationTemplates gets the notificationTemplates property value. 
func (m *GovernanceNotificationPolicy) GetNotificationTemplates()([]GovernanceNotificationTemplateable) {
    if m == nil {
        return nil
    } else {
        return m.notificationTemplates
    }
}
func (m *GovernanceNotificationPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GovernanceNotificationPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetEnabledTemplateTypes() != nil {
        err := writer.WriteCollectionOfStringValues("enabledTemplateTypes", m.GetEnabledTemplateTypes())
        if err != nil {
            return err
        }
    }
    if m.GetNotificationTemplates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNotificationTemplates()))
        for i, v := range m.GetNotificationTemplates() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernanceNotificationPolicy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEnabledTemplateTypes sets the enabledTemplateTypes property value. 
func (m *GovernanceNotificationPolicy) SetEnabledTemplateTypes(value []string)() {
    if m != nil {
        m.enabledTemplateTypes = value
    }
}
// SetNotificationTemplates sets the notificationTemplates property value. 
func (m *GovernanceNotificationPolicy) SetNotificationTemplates(value []GovernanceNotificationTemplateable)() {
    if m != nil {
        m.notificationTemplates = value
    }
}
