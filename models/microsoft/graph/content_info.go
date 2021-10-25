package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ContentInfo struct {
    additionalData map[string]interface{};
    format *ContentFormat;
    identifier *string;
    metadata []KeyValuePair;
    state *ContentState;
}
func NewContentInfo()(*ContentInfo) {
    m := &ContentInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ContentInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ContentInfo) GetFormat()(*ContentFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
func (m *ContentInfo) GetIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identifier
    }
}
func (m *ContentInfo) GetMetadata()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
func (m *ContentInfo) GetState()(*ContentState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *ContentInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseContentFormat)
        if err != nil {
            return err
        }
        cast := val.(ContentFormat)
        m.SetFormat(&cast)
        return nil
    }
    res["identifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIdentifier(val)
        return nil
    }
    res["metadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        res := make([]KeyValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValuePair))
        }
        m.SetMetadata(res)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseContentState)
        if err != nil {
            return err
        }
        cast := val.(ContentState)
        m.SetState(&cast)
        return nil
    }
    return res
}
func (m *ContentInfo) IsNil()(bool) {
    return m == nil
}
func (m *ContentInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetFormat() != nil {
        cast := m.GetFormat().String()
        err := writer.WriteStringValue("format", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identifier", m.GetIdentifier())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMetadata()))
        for i, v := range m.GetMetadata() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("metadata", cast)
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *ContentInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ContentInfo) SetFormat(value *ContentFormat)() {
    m.format = value
}
func (m *ContentInfo) SetIdentifier(value *string)() {
    m.identifier = value
}
func (m *ContentInfo) SetMetadata(value []KeyValuePair)() {
    m.metadata = value
}
func (m *ContentInfo) SetState(value *ContentState)() {
    m.state = value
}
