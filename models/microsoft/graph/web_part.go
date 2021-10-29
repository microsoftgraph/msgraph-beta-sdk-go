package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WebPart struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The required properties for the webPart (varies by webPart)
    data *SitePageData;
    // A unique identifier specifying the webPart type. Read-only.
    type_escaped *string;
}
// Instantiates a new webPart and sets the default values.
func NewWebPart()(*WebPart) {
    m := &WebPart{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WebPart) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the data property value. The required properties for the webPart (varies by webPart)
func (m *WebPart) GetData()(*SitePageData) {
    if m == nil {
        return nil
    } else {
        return m.data
    }
}
// Gets the type_escaped property value. A unique identifier specifying the webPart type. Read-only.
func (m *WebPart) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
func (m *WebPart) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["data"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSitePageData() })
        if err != nil {
            return err
        }
        m.SetData(val.(*SitePageData))
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
        return nil
    }
    return res
}
func (m *WebPart) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WebPart) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("data", m.GetData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type_escaped", m.GetType_escaped())
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
func (m *WebPart) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the data property value. The required properties for the webPart (varies by webPart)
// Parameters:
//  - value : Value to set for the data property.
func (m *WebPart) SetData(value *SitePageData)() {
    m.data = value
}
// Sets the type_escaped property value. A unique identifier specifying the webPart type. Read-only.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *WebPart) SetType_escaped(value *string)() {
    m.type_escaped = value
}
