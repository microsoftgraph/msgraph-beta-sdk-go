package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SensitiveContentEvidence 
type SensitiveContentEvidence struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    length *int32;
    // 
    match *string;
    // 
    offset *int32;
}
// NewSensitiveContentEvidence instantiates a new sensitiveContentEvidence and sets the default values.
func NewSensitiveContentEvidence()(*SensitiveContentEvidence) {
    m := &SensitiveContentEvidence{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SensitiveContentEvidence) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetLength gets the length property value. 
func (m *SensitiveContentEvidence) GetLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.length
    }
}
// GetMatch gets the match property value. 
func (m *SensitiveContentEvidence) GetMatch()(*string) {
    if m == nil {
        return nil
    } else {
        return m.match
    }
}
// GetOffset gets the offset property value. 
func (m *SensitiveContentEvidence) GetOffset()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.offset
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitiveContentEvidence) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["match"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatch(val)
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
func (m *SensitiveContentEvidence) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SensitiveContentEvidence) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("length", m.GetLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("match", m.GetMatch())
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
func (m *SensitiveContentEvidence) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLength sets the length property value. 
func (m *SensitiveContentEvidence) SetLength(value *int32)() {
    if m != nil {
        m.length = value
    }
}
// SetMatch sets the match property value. 
func (m *SensitiveContentEvidence) SetMatch(value *string)() {
    if m != nil {
        m.match = value
    }
}
// SetOffset sets the offset property value. 
func (m *SensitiveContentEvidence) SetOffset(value *int32)() {
    if m != nil {
        m.offset = value
    }
}
