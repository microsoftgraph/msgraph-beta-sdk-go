package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SensitiveType struct {
    Entity
    // 
    classificationMethod *ClassificationMethod;
    // 
    description *string;
    // 
    name *string;
    // 
    publisherName *string;
    // 
    rulePackageId *string;
    // 
    rulePackageType *string;
    // 
    scope *SensitiveTypeScope;
    // 
    sensitiveTypeSource *SensitiveTypeSource;
    // 
    state *string;
}
// Instantiates a new sensitiveType and sets the default values.
func NewSensitiveType()(*SensitiveType) {
    m := &SensitiveType{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the classificationMethod property value. 
func (m *SensitiveType) GetClassificationMethod()(*ClassificationMethod) {
    if m == nil {
        return nil
    } else {
        return m.classificationMethod
    }
}
// Gets the description property value. 
func (m *SensitiveType) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the name property value. 
func (m *SensitiveType) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the publisherName property value. 
func (m *SensitiveType) GetPublisherName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisherName
    }
}
// Gets the rulePackageId property value. 
func (m *SensitiveType) GetRulePackageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rulePackageId
    }
}
// Gets the rulePackageType property value. 
func (m *SensitiveType) GetRulePackageType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rulePackageType
    }
}
// Gets the scope property value. 
func (m *SensitiveType) GetScope()(*SensitiveTypeScope) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// Gets the sensitiveTypeSource property value. 
func (m *SensitiveType) GetSensitiveTypeSource()(*SensitiveTypeSource) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeSource
    }
}
// Gets the state property value. 
func (m *SensitiveType) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// The deserialization information for the current model
func (m *SensitiveType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classificationMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseClassificationMethod)
        if err != nil {
            return err
        }
        cast := val.(ClassificationMethod)
        m.SetClassificationMethod(&cast)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
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
    res["scope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitiveTypeScope)
        if err != nil {
            return err
        }
        cast := val.(SensitiveTypeScope)
        m.SetScope(&cast)
        return nil
    }
    res["sensitiveTypeSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitiveTypeSource)
        if err != nil {
            return err
        }
        cast := val.(SensitiveTypeSource)
        m.SetSensitiveTypeSource(&cast)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetState(val)
        return nil
    }
    return res
}
func (m *SensitiveType) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SensitiveType) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassificationMethod() != nil {
        cast := m.GetClassificationMethod().String()
        err = writer.WriteStringValue("classificationMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisherName", m.GetPublisherName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("rulePackageId", m.GetRulePackageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("rulePackageType", m.GetRulePackageType())
        if err != nil {
            return err
        }
    }
    if m.GetScope() != nil {
        cast := m.GetScope().String()
        err = writer.WriteStringValue("scope", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensitiveTypeSource() != nil {
        cast := m.GetSensitiveTypeSource().String()
        err = writer.WriteStringValue("sensitiveTypeSource", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the classificationMethod property value. 
// Parameters:
//  - value : Value to set for the classificationMethod property.
func (m *SensitiveType) SetClassificationMethod(value *ClassificationMethod)() {
    m.classificationMethod = value
}
// Sets the description property value. 
// Parameters:
//  - value : Value to set for the description property.
func (m *SensitiveType) SetDescription(value *string)() {
    m.description = value
}
// Sets the name property value. 
// Parameters:
//  - value : Value to set for the name property.
func (m *SensitiveType) SetName(value *string)() {
    m.name = value
}
// Sets the publisherName property value. 
// Parameters:
//  - value : Value to set for the publisherName property.
func (m *SensitiveType) SetPublisherName(value *string)() {
    m.publisherName = value
}
// Sets the rulePackageId property value. 
// Parameters:
//  - value : Value to set for the rulePackageId property.
func (m *SensitiveType) SetRulePackageId(value *string)() {
    m.rulePackageId = value
}
// Sets the rulePackageType property value. 
// Parameters:
//  - value : Value to set for the rulePackageType property.
func (m *SensitiveType) SetRulePackageType(value *string)() {
    m.rulePackageType = value
}
// Sets the scope property value. 
// Parameters:
//  - value : Value to set for the scope property.
func (m *SensitiveType) SetScope(value *SensitiveTypeScope)() {
    m.scope = value
}
// Sets the sensitiveTypeSource property value. 
// Parameters:
//  - value : Value to set for the sensitiveTypeSource property.
func (m *SensitiveType) SetSensitiveTypeSource(value *SensitiveTypeSource)() {
    m.sensitiveTypeSource = value
}
// Sets the state property value. 
// Parameters:
//  - value : Value to set for the state property.
func (m *SensitiveType) SetState(value *string)() {
    m.state = value
}
