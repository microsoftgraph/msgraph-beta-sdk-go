package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ContentClassification provides operations to call the classifyExactMatches method.
type ContentClassification struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    confidence *int32;
    // 
    matches []MatchLocationable;
    // 
    sensitiveTypeId *string;
    // 
    uniqueCount *int32;
}
// NewContentClassification instantiates a new contentClassification and sets the default values.
func NewContentClassification()(*ContentClassification) {
    m := &ContentClassification{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateContentClassificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContentClassificationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewContentClassification(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContentClassification) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConfidence gets the confidence property value. 
func (m *ContentClassification) GetConfidence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.confidence
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
        val, err := n.GetCollectionOfObjectValues(CreateMatchLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MatchLocationable, len(val))
            for i, v := range val {
                res[i] = v.(MatchLocationable)
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
// GetMatches gets the matches property value. 
func (m *ContentClassification) GetMatches()([]MatchLocationable) {
    if m == nil {
        return nil
    } else {
        return m.matches
    }
}
// GetSensitiveTypeId gets the sensitiveTypeId property value. 
func (m *ContentClassification) GetSensitiveTypeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeId
    }
}
// GetUniqueCount gets the uniqueCount property value. 
func (m *ContentClassification) GetUniqueCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.uniqueCount
    }
}
func (m *ContentClassification) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ContentClassification) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("confidence", m.GetConfidence())
        if err != nil {
            return err
        }
    }
    if m.GetMatches() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMatches()))
        for i, v := range m.GetMatches() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContentClassification) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConfidence sets the confidence property value. 
func (m *ContentClassification) SetConfidence(value *int32)() {
    if m != nil {
        m.confidence = value
    }
}
// SetMatches sets the matches property value. 
func (m *ContentClassification) SetMatches(value []MatchLocationable)() {
    if m != nil {
        m.matches = value
    }
}
// SetSensitiveTypeId sets the sensitiveTypeId property value. 
func (m *ContentClassification) SetSensitiveTypeId(value *string)() {
    if m != nil {
        m.sensitiveTypeId = value
    }
}
// SetUniqueCount sets the uniqueCount property value. 
func (m *ContentClassification) SetUniqueCount(value *int32)() {
    if m != nil {
        m.uniqueCount = value
    }
}
