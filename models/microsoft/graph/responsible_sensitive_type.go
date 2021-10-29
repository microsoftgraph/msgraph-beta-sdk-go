package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new responsibleSensitiveType and sets the default values.
func NewResponsibleSensitiveType()(*ResponsibleSensitiveType) {
    m := &ResponsibleSensitiveType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResponsibleSensitiveType) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the description property value. 
func (m *ResponsibleSensitiveType) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the id property value. 
func (m *ResponsibleSensitiveType) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the name property value. 
func (m *ResponsibleSensitiveType) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the publisherName property value. 
func (m *ResponsibleSensitiveType) GetPublisherName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisherName
    }
}
// Gets the rulePackageId property value. 
func (m *ResponsibleSensitiveType) GetRulePackageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rulePackageId
    }
}
// Gets the rulePackageType property value. 
func (m *ResponsibleSensitiveType) GetRulePackageType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rulePackageType
    }
}
// The deserialization information for the current model
func (m *ResponsibleSensitiveType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
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
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["publisherName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPublisherName(val)
        return nil
    }
    res["rulePackageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRulePackageId(val)
        return nil
    }
    res["rulePackageType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRulePackageType(val)
        return nil
    }
    return res
}
func (m *ResponsibleSensitiveType) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ResponsibleSensitiveType) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the description property value. 
// Parameters:
//  - value : Value to set for the description property.
func (m *ResponsibleSensitiveType) SetDescription(value *string)() {
    m.description = value
}
// Sets the id property value. 
// Parameters:
//  - value : Value to set for the id property.
func (m *ResponsibleSensitiveType) SetId(value *string)() {
    m.id = value
}
// Sets the name property value. 
// Parameters:
//  - value : Value to set for the name property.
func (m *ResponsibleSensitiveType) SetName(value *string)() {
    m.name = value
}
// Sets the publisherName property value. 
// Parameters:
//  - value : Value to set for the publisherName property.
func (m *ResponsibleSensitiveType) SetPublisherName(value *string)() {
    m.publisherName = value
}
// Sets the rulePackageId property value. 
// Parameters:
//  - value : Value to set for the rulePackageId property.
func (m *ResponsibleSensitiveType) SetRulePackageId(value *string)() {
    m.rulePackageId = value
}
// Sets the rulePackageType property value. 
// Parameters:
//  - value : Value to set for the rulePackageType property.
func (m *ResponsibleSensitiveType) SetRulePackageType(value *string)() {
    m.rulePackageType = value
}
