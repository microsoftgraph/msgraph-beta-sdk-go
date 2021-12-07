package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ResponsibleSensitiveType 
type ResponsibleSensitiveType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    description *string;
    // 
    id *string;
    // 
    name *string;
    // 
    publisherName *string;
    // 
    rulePackageId *string;
    // 
    rulePackageType *string;
}
// NewResponsibleSensitiveType instantiates a new responsibleSensitiveType and sets the default values.
func NewResponsibleSensitiveType()(*ResponsibleSensitiveType) {
    m := &ResponsibleSensitiveType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResponsibleSensitiveType) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. 
func (m *ResponsibleSensitiveType) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetId gets the id property value. 
func (m *ResponsibleSensitiveType) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetName gets the name property value. 
func (m *ResponsibleSensitiveType) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetPublisherName gets the publisherName property value. 
func (m *ResponsibleSensitiveType) GetPublisherName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisherName
    }
}
// GetRulePackageId gets the rulePackageId property value. 
func (m *ResponsibleSensitiveType) GetRulePackageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rulePackageId
    }
}
// GetRulePackageType gets the rulePackageType property value. 
func (m *ResponsibleSensitiveType) GetRulePackageType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rulePackageType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResponsibleSensitiveType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["publisherName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisherName(val)
        }
        return nil
    }
    res["rulePackageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRulePackageId(val)
        }
        return nil
    }
    res["rulePackageType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRulePackageType(val)
        }
        return nil
    }
    return res
}
func (m *ResponsibleSensitiveType) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ResponsibleSensitiveType) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publisherName", m.GetPublisherName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("rulePackageId", m.GetRulePackageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("rulePackageType", m.GetRulePackageType())
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
func (m *ResponsibleSensitiveType) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDescription sets the description property value. 
func (m *ResponsibleSensitiveType) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetId sets the id property value. 
func (m *ResponsibleSensitiveType) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetName sets the name property value. 
func (m *ResponsibleSensitiveType) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetPublisherName sets the publisherName property value. 
func (m *ResponsibleSensitiveType) SetPublisherName(value *string)() {
    if m != nil {
        m.publisherName = value
    }
}
// SetRulePackageId sets the rulePackageId property value. 
func (m *ResponsibleSensitiveType) SetRulePackageId(value *string)() {
    if m != nil {
        m.rulePackageId = value
    }
}
// SetRulePackageType sets the rulePackageType property value. 
func (m *ResponsibleSensitiveType) SetRulePackageType(value *string)() {
    if m != nil {
        m.rulePackageType = value
    }
}
