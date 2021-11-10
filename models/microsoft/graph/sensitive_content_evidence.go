package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new sensitiveContentEvidence and sets the default values.
func NewSensitiveContentEvidence()(*SensitiveContentEvidence) {
    m := &SensitiveContentEvidence{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SensitiveContentEvidence) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the length property value. 
func (m *SensitiveContentEvidence) GetLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.length
    }
}
// Gets the match property value. 
func (m *SensitiveContentEvidence) GetMatch()(*string) {
    if m == nil {
        return nil
    } else {
        return m.match
    }
}
// Gets the offset property value. 
func (m *SensitiveContentEvidence) GetOffset()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.offset
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SensitiveContentEvidence) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the length property value. 
// Parameters:
//  - value : Value to set for the length property.
func (m *SensitiveContentEvidence) SetLength(value *int32)() {
    m.length = value
}
// Sets the match property value. 
// Parameters:
//  - value : Value to set for the match property.
func (m *SensitiveContentEvidence) SetMatch(value *string)() {
    m.match = value
}
// Sets the offset property value. 
// Parameters:
//  - value : Value to set for the offset property.
func (m *SensitiveContentEvidence) SetOffset(value *int32)() {
    m.offset = value
}
