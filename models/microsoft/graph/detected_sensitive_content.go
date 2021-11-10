package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DetectedSensitiveContent struct {
    DetectedSensitiveContentBase
    // 
    classificationAttributes []ClassificationAttribute;
    // 
    classificationMethod *ClassificationMethod;
    // 
    matches []SensitiveContentLocation;
    // 
    scope *SensitiveTypeScope;
    // 
    sensitiveTypeSource *SensitiveTypeSource;
}
// Instantiates a new detectedSensitiveContent and sets the default values.
func NewDetectedSensitiveContent()(*DetectedSensitiveContent) {
    m := &DetectedSensitiveContent{
        DetectedSensitiveContentBase: *NewDetectedSensitiveContentBase(),
    }
    return m
}
// Gets the classificationAttributes property value. 
func (m *DetectedSensitiveContent) GetClassificationAttributes()([]ClassificationAttribute) {
    if m == nil {
        return nil
    } else {
        return m.classificationAttributes
    }
}
// Gets the classificationMethod property value. 
func (m *DetectedSensitiveContent) GetClassificationMethod()(*ClassificationMethod) {
    if m == nil {
        return nil
    } else {
        return m.classificationMethod
    }
}
// Gets the matches property value. 
func (m *DetectedSensitiveContent) GetMatches()([]SensitiveContentLocation) {
    if m == nil {
        return nil
    } else {
        return m.matches
    }
}
// Gets the scope property value. 
func (m *DetectedSensitiveContent) GetScope()(*SensitiveTypeScope) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// Gets the sensitiveTypeSource property value. 
func (m *DetectedSensitiveContent) GetSensitiveTypeSource()(*SensitiveTypeSource) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeSource
    }
}
// The deserialization information for the current model
func (m *DetectedSensitiveContent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DetectedSensitiveContentBase.GetFieldDeserializers()
    res["classificationAttributes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewClassificationAttribute() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ClassificationAttribute, len(val))
            for i, v := range val {
                res[i] = *(v.(*ClassificationAttribute))
            }
            m.SetClassificationAttributes(res)
        }
        return nil
    }
    res["classificationMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseClassificationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ClassificationMethod)
            m.SetClassificationMethod(&cast)
        }
        return nil
    }
    res["matches"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSensitiveContentLocation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitiveContentLocation, len(val))
            for i, v := range val {
                res[i] = *(v.(*SensitiveContentLocation))
            }
            m.SetMatches(res)
        }
        return nil
    }
    res["scope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitiveTypeScope)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(SensitiveTypeScope)
            m.SetScope(&cast)
        }
        return nil
    }
    res["sensitiveTypeSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitiveTypeSource)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(SensitiveTypeSource)
            m.SetSensitiveTypeSource(&cast)
        }
        return nil
    }
    return res
}
func (m *DetectedSensitiveContent) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DetectedSensitiveContent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DetectedSensitiveContentBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClassificationAttributes()))
        for i, v := range m.GetClassificationAttributes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("classificationAttributes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetClassificationMethod() != nil {
        cast := m.GetClassificationMethod().String()
        err = writer.WriteStringValue("classificationMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMatches()))
        for i, v := range m.GetMatches() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("matches", cast)
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
    return nil
}
// Sets the classificationAttributes property value. 
// Parameters:
//  - value : Value to set for the classificationAttributes property.
func (m *DetectedSensitiveContent) SetClassificationAttributes(value []ClassificationAttribute)() {
    m.classificationAttributes = value
}
// Sets the classificationMethod property value. 
// Parameters:
//  - value : Value to set for the classificationMethod property.
func (m *DetectedSensitiveContent) SetClassificationMethod(value *ClassificationMethod)() {
    m.classificationMethod = value
}
// Sets the matches property value. 
// Parameters:
//  - value : Value to set for the matches property.
func (m *DetectedSensitiveContent) SetMatches(value []SensitiveContentLocation)() {
    m.matches = value
}
// Sets the scope property value. 
// Parameters:
//  - value : Value to set for the scope property.
func (m *DetectedSensitiveContent) SetScope(value *SensitiveTypeScope)() {
    m.scope = value
}
// Sets the sensitiveTypeSource property value. 
// Parameters:
//  - value : Value to set for the sensitiveTypeSource property.
func (m *DetectedSensitiveContent) SetSensitiveTypeSource(value *SensitiveTypeSource)() {
    m.sensitiveTypeSource = value
}
