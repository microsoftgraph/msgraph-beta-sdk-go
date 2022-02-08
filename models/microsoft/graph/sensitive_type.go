package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SensitiveType 
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
// NewSensitiveType instantiates a new sensitiveType and sets the default values.
func NewSensitiveType()(*SensitiveType) {
    m := &SensitiveType{
        Entity: *NewEntity(),
    }
    return m
}
// GetClassificationMethod gets the classificationMethod property value. 
func (m *SensitiveType) GetClassificationMethod()(*ClassificationMethod) {
    if m == nil {
        return nil
    } else {
        return m.classificationMethod
    }
}
// GetDescription gets the description property value. 
func (m *SensitiveType) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetName gets the name property value. 
func (m *SensitiveType) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetPublisherName gets the publisherName property value. 
func (m *SensitiveType) GetPublisherName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisherName
    }
}
// GetRulePackageId gets the rulePackageId property value. 
func (m *SensitiveType) GetRulePackageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rulePackageId
    }
}
// GetRulePackageType gets the rulePackageType property value. 
func (m *SensitiveType) GetRulePackageType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rulePackageType
    }
}
// GetScope gets the scope property value. 
func (m *SensitiveType) GetScope()(*SensitiveTypeScope) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// GetSensitiveTypeSource gets the sensitiveTypeSource property value. 
func (m *SensitiveType) GetSensitiveTypeSource()(*SensitiveTypeSource) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeSource
    }
}
// GetState gets the state property value. 
func (m *SensitiveType) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitiveType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classificationMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseClassificationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassificationMethod(val.(*ClassificationMethod))
        }
        return nil
    }
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
    res["scope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitiveTypeScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(*SensitiveTypeScope))
        }
        return nil
    }
    res["sensitiveTypeSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitiveTypeSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitiveTypeSource(val.(*SensitiveTypeSource))
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    return res
}
func (m *SensitiveType) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SensitiveType) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassificationMethod() != nil {
        cast := (*m.GetClassificationMethod()).String()
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
        cast := (*m.GetScope()).String()
        err = writer.WriteStringValue("scope", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensitiveTypeSource() != nil {
        cast := (*m.GetSensitiveTypeSource()).String()
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
// SetClassificationMethod sets the classificationMethod property value. 
func (m *SensitiveType) SetClassificationMethod(value *ClassificationMethod)() {
    if m != nil {
        m.classificationMethod = value
    }
}
// SetDescription sets the description property value. 
func (m *SensitiveType) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetName sets the name property value. 
func (m *SensitiveType) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetPublisherName sets the publisherName property value. 
func (m *SensitiveType) SetPublisherName(value *string)() {
    if m != nil {
        m.publisherName = value
    }
}
// SetRulePackageId sets the rulePackageId property value. 
func (m *SensitiveType) SetRulePackageId(value *string)() {
    if m != nil {
        m.rulePackageId = value
    }
}
// SetRulePackageType sets the rulePackageType property value. 
func (m *SensitiveType) SetRulePackageType(value *string)() {
    if m != nil {
        m.rulePackageType = value
    }
}
// SetScope sets the scope property value. 
func (m *SensitiveType) SetScope(value *SensitiveTypeScope)() {
    if m != nil {
        m.scope = value
    }
}
// SetSensitiveTypeSource sets the sensitiveTypeSource property value. 
func (m *SensitiveType) SetSensitiveTypeSource(value *SensitiveTypeSource)() {
    if m != nil {
        m.sensitiveTypeSource = value
    }
}
// SetState sets the state property value. 
func (m *SensitiveType) SetState(value *string)() {
    if m != nil {
        m.state = value
    }
}
