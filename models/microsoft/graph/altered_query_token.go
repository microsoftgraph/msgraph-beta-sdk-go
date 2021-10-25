package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AlteredQueryToken struct {
    additionalData map[string]interface{};
    length *int32;
    offset *int32;
    suggestion *string;
}
func NewAlteredQueryToken()(*AlteredQueryToken) {
    m := &AlteredQueryToken{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AlteredQueryToken) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AlteredQueryToken) GetLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.length
    }
}
func (m *AlteredQueryToken) GetOffset()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.offset
    }
}
func (m *AlteredQueryToken) GetSuggestion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.suggestion
    }
}
func (m *AlteredQueryToken) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["suggestion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSuggestion(val)
        return nil
    }
    return res
}
func (m *AlteredQueryToken) IsNil()(bool) {
    return m == nil
}
func (m *AlteredQueryToken) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("suggestion", m.GetSuggestion())
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
func (m *AlteredQueryToken) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AlteredQueryToken) SetLength(value *int32)() {
    m.length = value
}
func (m *AlteredQueryToken) SetOffset(value *int32)() {
    m.offset = value
}
func (m *AlteredQueryToken) SetSuggestion(value *string)() {
    m.suggestion = value
}
