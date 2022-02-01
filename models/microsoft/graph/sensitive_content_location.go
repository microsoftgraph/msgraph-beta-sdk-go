package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SensitiveContentLocation 
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
// NewSensitiveContentLocation instantiates a new sensitiveContentLocation and sets the default values.
func NewSensitiveContentLocation()(*SensitiveContentLocation) {
    m := &SensitiveContentLocation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SensitiveContentLocation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConfidence gets the confidence property value. 
func (m *SensitiveContentLocation) GetConfidence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.confidence
    }
}
// GetEvidences gets the evidences property value. 
func (m *SensitiveContentLocation) GetEvidences()([]SensitiveContentEvidence) {
    if m == nil {
        return nil
    } else {
        return m.evidences
    }
}
// GetIdMatch gets the idMatch property value. 
func (m *SensitiveContentLocation) GetIdMatch()(*string) {
    if m == nil {
        return nil
    } else {
        return m.idMatch
    }
}
// GetLength gets the length property value. 
func (m *SensitiveContentLocation) GetLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.length
    }
}
// GetOffset gets the offset property value. 
func (m *SensitiveContentLocation) GetOffset()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.offset
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitiveContentLocation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["evidences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSensitiveContentEvidence() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitiveContentEvidence, len(val))
            for i, v := range val {
                res[i] = *(v.(*SensitiveContentEvidence))
            }
            m.SetEvidences(res)
        }
        return nil
    }
    res["idMatch"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdMatch(val)
        }
        return nil
    }
    res["length"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLength(val)
        }
        return nil
    }
    res["offset"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffset(val)
        }
        return nil
    }
    return res
}
func (m *SensitiveContentLocation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SensitiveContentLocation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("confidence", m.GetConfidence())
        if err != nil {
            return err
        }
    }
    if m.GetEvidences() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SensitiveContentLocation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConfidence sets the confidence property value. 
func (m *SensitiveContentLocation) SetConfidence(value *int32)() {
    if m != nil {
        m.confidence = value
    }
}
// SetEvidences sets the evidences property value. 
func (m *SensitiveContentLocation) SetEvidences(value []SensitiveContentEvidence)() {
    if m != nil {
        m.evidences = value
    }
}
// SetIdMatch sets the idMatch property value. 
func (m *SensitiveContentLocation) SetIdMatch(value *string)() {
    if m != nil {
        m.idMatch = value
    }
}
// SetLength sets the length property value. 
func (m *SensitiveContentLocation) SetLength(value *int32)() {
    if m != nil {
        m.length = value
    }
}
// SetOffset sets the offset property value. 
func (m *SensitiveContentLocation) SetOffset(value *int32)() {
    if m != nil {
        m.offset = value
    }
}
