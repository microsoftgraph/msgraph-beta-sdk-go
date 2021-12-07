package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WebPart 
type WebPart struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The required properties for the webPart (varies by webPart)
    data *SitePageData;
    // A unique identifier specifying the webPart type. Read-only.
    type_escaped *string;
}
// NewWebPart instantiates a new webPart and sets the default values.
func NewWebPart()(*WebPart) {
    m := &WebPart{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WebPart) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetData gets the data property value. The required properties for the webPart (varies by webPart)
func (m *WebPart) GetData()(*SitePageData) {
    if m == nil {
        return nil
    } else {
        return m.data
    }
}
// GetType gets the type property value. A unique identifier specifying the webPart type. Read-only.
func (m *WebPart) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebPart) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["data"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSitePageData() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetData(val.(*SitePageData))
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
func (m *WebPart) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WebPart) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("data", m.GetData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *WebPart) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetData sets the data property value. The required properties for the webPart (varies by webPart)
func (m *WebPart) SetData(value *SitePageData)() {
    if m != nil {
        m.data = value
    }
}
// SetType sets the type property value. A unique identifier specifying the webPart type. Read-only.
func (m *WebPart) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
