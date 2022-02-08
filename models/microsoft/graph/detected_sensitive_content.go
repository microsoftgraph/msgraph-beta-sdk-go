package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DetectedSensitiveContent 
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
// NewDetectedSensitiveContent instantiates a new detectedSensitiveContent and sets the default values.
func NewDetectedSensitiveContent()(*DetectedSensitiveContent) {
    m := &DetectedSensitiveContent{
        DetectedSensitiveContentBase: *NewDetectedSensitiveContentBase(),
    }
    return m
}
// GetClassificationAttributes gets the classificationAttributes property value. 
func (m *DetectedSensitiveContent) GetClassificationAttributes()([]ClassificationAttribute) {
    if m == nil {
        return nil
    } else {
        return m.classificationAttributes
    }
}
// GetClassificationMethod gets the classificationMethod property value. 
func (m *DetectedSensitiveContent) GetClassificationMethod()(*ClassificationMethod) {
    if m == nil {
        return nil
    } else {
        return m.classificationMethod
    }
}
// GetMatches gets the matches property value. 
func (m *DetectedSensitiveContent) GetMatches()([]SensitiveContentLocation) {
    if m == nil {
        return nil
    } else {
        return m.matches
    }
}
// GetScope gets the scope property value. 
func (m *DetectedSensitiveContent) GetScope()(*SensitiveTypeScope) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// GetSensitiveTypeSource gets the sensitiveTypeSource property value. 
func (m *DetectedSensitiveContent) GetSensitiveTypeSource()(*SensitiveTypeSource) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeSource
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
            m.SetClassificationMethod(val.(*ClassificationMethod))
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
    return res
}
func (m *DetectedSensitiveContent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DetectedSensitiveContent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DetectedSensitiveContentBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassificationAttributes() != nil {
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
        cast := (*m.GetClassificationMethod()).String()
        err = writer.WriteStringValue("classificationMethod", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMatches() != nil {
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
    return nil
}
// SetClassificationAttributes sets the classificationAttributes property value. 
func (m *DetectedSensitiveContent) SetClassificationAttributes(value []ClassificationAttribute)() {
    if m != nil {
        m.classificationAttributes = value
    }
}
// SetClassificationMethod sets the classificationMethod property value. 
func (m *DetectedSensitiveContent) SetClassificationMethod(value *ClassificationMethod)() {
    if m != nil {
        m.classificationMethod = value
    }
}
// SetMatches sets the matches property value. 
func (m *DetectedSensitiveContent) SetMatches(value []SensitiveContentLocation)() {
    if m != nil {
        m.matches = value
    }
}
// SetScope sets the scope property value. 
func (m *DetectedSensitiveContent) SetScope(value *SensitiveTypeScope)() {
    if m != nil {
        m.scope = value
    }
}
// SetSensitiveTypeSource sets the sensitiveTypeSource property value. 
func (m *DetectedSensitiveContent) SetSensitiveTypeSource(value *SensitiveTypeSource)() {
    if m != nil {
        m.sensitiveTypeSource = value
    }
}
