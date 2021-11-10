package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ContentClassification struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    confidence *int32;
    // 
    matches []MatchLocation;
    // 
    sensitiveTypeId *string;
    // 
    uniqueCount *int32;
}
// Instantiates a new contentClassification and sets the default values.
func NewContentClassification()(*ContentClassification) {
    m := &ContentClassification{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContentClassification) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the confidence property value. 
func (m *ContentClassification) GetConfidence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.confidence
    }
}
// Gets the matches property value. 
func (m *ContentClassification) GetMatches()([]MatchLocation) {
    if m == nil {
        return nil
    } else {
        return m.matches
    }
}
// Gets the sensitiveTypeId property value. 
func (m *ContentClassification) GetSensitiveTypeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeId
    }
}
// Gets the uniqueCount property value. 
func (m *ContentClassification) GetUniqueCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.uniqueCount
    }
}
// The deserialization information for the current model
func (m *ContentClassification) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["confidence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfidence(val)
        }
        return nil
    }
    res["matches"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMatchLocation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MatchLocation, len(val))
            for i, v := range val {
                res[i] = *(v.(*MatchLocation))
            }
            m.SetMatches(res)
        }
        return nil
    }
    res["sensitiveTypeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitiveTypeId(val)
        }
        return nil
    }
    res["uniqueCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueCount(val)
        }
        return nil
    }
    return res
}
func (m *ContentClassification) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ContentClassification) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("confidence", m.GetConfidence())
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
        err := writer.WriteCollectionOfObjectValues("matches", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sensitiveTypeId", m.GetSensitiveTypeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("uniqueCount", m.GetUniqueCount())
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
func (m *ContentClassification) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the confidence property value. 
// Parameters:
//  - value : Value to set for the confidence property.
func (m *ContentClassification) SetConfidence(value *int32)() {
    m.confidence = value
}
// Sets the matches property value. 
// Parameters:
//  - value : Value to set for the matches property.
func (m *ContentClassification) SetMatches(value []MatchLocation)() {
    m.matches = value
}
// Sets the sensitiveTypeId property value. 
// Parameters:
//  - value : Value to set for the sensitiveTypeId property.
func (m *ContentClassification) SetSensitiveTypeId(value *string)() {
    m.sensitiveTypeId = value
}
// Sets the uniqueCount property value. 
// Parameters:
//  - value : Value to set for the uniqueCount property.
func (m *ContentClassification) SetUniqueCount(value *int32)() {
    m.uniqueCount = value
}
