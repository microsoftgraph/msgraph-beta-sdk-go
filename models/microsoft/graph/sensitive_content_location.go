package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SensitiveContentLocation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    confidence *int32;
    // 
    evidences []SensitiveContentEvidence;
    // 
    idMatch *string;
    // 
    length *int32;
    // 
    offset *int32;
}
// Instantiates a new sensitiveContentLocation and sets the default values.
func NewSensitiveContentLocation()(*SensitiveContentLocation) {
    m := &SensitiveContentLocation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SensitiveContentLocation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the confidence property value. 
func (m *SensitiveContentLocation) GetConfidence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.confidence
    }
}
// Gets the evidences property value. 
func (m *SensitiveContentLocation) GetEvidences()([]SensitiveContentEvidence) {
    if m == nil {
        return nil
    } else {
        return m.evidences
    }
}
// Gets the idMatch property value. 
func (m *SensitiveContentLocation) GetIdMatch()(*string) {
    if m == nil {
        return nil
    } else {
        return m.idMatch
    }
}
// Gets the length property value. 
func (m *SensitiveContentLocation) GetLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.length
    }
}
// Gets the offset property value. 
func (m *SensitiveContentLocation) GetOffset()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.offset
    }
}
// The deserialization information for the current model
func (m *SensitiveContentLocation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["confidence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetConfidence(val)
        return nil
    }
    res["evidences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSensitiveContentEvidence() })
        if err != nil {
            return err
        }
        res := make([]SensitiveContentEvidence, len(val))
        for i, v := range val {
            res[i] = *(v.(*SensitiveContentEvidence))
        }
        m.SetEvidences(res)
        return nil
    }
    res["idMatch"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIdMatch(val)
        return nil
    }
    res["length"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetLength(val)
        return nil
    }
    res["offset"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetOffset(val)
        return nil
    }
    return res
}
func (m *SensitiveContentLocation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SensitiveContentLocation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("confidence", m.GetConfidence())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEvidences()))
        for i, v := range m.GetEvidences() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("evidences", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("idMatch", m.GetIdMatch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("length", m.GetLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("offset", m.GetOffset())
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
func (m *SensitiveContentLocation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the confidence property value. 
// Parameters:
//  - value : Value to set for the confidence property.
func (m *SensitiveContentLocation) SetConfidence(value *int32)() {
    m.confidence = value
}
// Sets the evidences property value. 
// Parameters:
//  - value : Value to set for the evidences property.
func (m *SensitiveContentLocation) SetEvidences(value []SensitiveContentEvidence)() {
    m.evidences = value
}
// Sets the idMatch property value. 
// Parameters:
//  - value : Value to set for the idMatch property.
func (m *SensitiveContentLocation) SetIdMatch(value *string)() {
    m.idMatch = value
}
// Sets the length property value. 
// Parameters:
//  - value : Value to set for the length property.
func (m *SensitiveContentLocation) SetLength(value *int32)() {
    m.length = value
}
// Sets the offset property value. 
// Parameters:
//  - value : Value to set for the offset property.
func (m *SensitiveContentLocation) SetOffset(value *int32)() {
    m.offset = value
}
