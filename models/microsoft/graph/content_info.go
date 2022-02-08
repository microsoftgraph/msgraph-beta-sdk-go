package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ContentInfo 
type ContentInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Possible values are: default, email.
    format *ContentFormat;
    // Identifier used for Azure Information Protection Analytics.
    identifier *string;
    // Existing Microsoft Information Protection metadata is passed as key/value pairs, where the key is the MSIP_Label_GUID_PropName.
    metadata []KeyValuePair;
    // Possible values are: rest, motion, use.
    state *ContentState;
}
// NewContentInfo instantiates a new contentInfo and sets the default values.
func NewContentInfo()(*ContentInfo) {
    m := &ContentInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContentInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFormat gets the format property value. Possible values are: default, email.
func (m *ContentInfo) GetFormat()(*ContentFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetIdentifier gets the identifier property value. Identifier used for Azure Information Protection Analytics.
func (m *ContentInfo) GetIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identifier
    }
}
// GetMetadata gets the metadata property value. Existing Microsoft Information Protection metadata is passed as key/value pairs, where the key is the MSIP_Label_GUID_PropName.
func (m *ContentInfo) GetMetadata()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
// GetState gets the state property value. Possible values are: rest, motion, use.
func (m *ContentInfo) GetState()(*ContentState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseContentFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(*ContentFormat))
        }
        return nil
    }
    res["identifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val)
        }
        return nil
    }
    res["metadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePair, len(val))
            for i, v := range val {
                res[i] = *(v.(*KeyValuePair))
            }
            m.SetMetadata(res)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseContentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ContentState))
        }
        return nil
    }
    return res
}
func (m *ContentInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ContentInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetFormat() != nil {
        cast := (*m.GetFormat()).String()
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
    if m.GetMetadata() != nil {
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
        cast := (*m.GetState()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContentInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFormat sets the format property value. Possible values are: default, email.
func (m *ContentInfo) SetFormat(value *ContentFormat)() {
    if m != nil {
        m.format = value
    }
}
// SetIdentifier sets the identifier property value. Identifier used for Azure Information Protection Analytics.
func (m *ContentInfo) SetIdentifier(value *string)() {
    if m != nil {
        m.identifier = value
    }
}
// SetMetadata sets the metadata property value. Existing Microsoft Information Protection metadata is passed as key/value pairs, where the key is the MSIP_Label_GUID_PropName.
func (m *ContentInfo) SetMetadata(value []KeyValuePair)() {
    if m != nil {
        m.metadata = value
    }
}
// SetState sets the state property value. Possible values are: rest, motion, use.
func (m *ContentInfo) SetState(value *ContentState)() {
    if m != nil {
        m.state = value
    }
}
