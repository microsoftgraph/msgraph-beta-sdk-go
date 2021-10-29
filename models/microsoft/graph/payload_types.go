package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PayloadTypes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The notification content of a raw user notification that will be delivered to and consumed by the app client on all supported platforms (Windows, iOS, Android or WebPush) receiving this notification. At least one of Payload.RawContent or Payload.VisualContent needs to be valid for a POST Notification request.
    rawContent *string;
    // The visual content of a visual user notification, which will be consumed by the notification platform on each supported platform (Windows, iOS and Android only) and rendered for the user. At least one of Payload.RawContent or Payload.VisualContent needs to be valid for a POST Notification request.
    visualContent *VisualProperties;
}
// Instantiates a new payloadTypes and sets the default values.
func NewPayloadTypes()(*PayloadTypes) {
    m := &PayloadTypes{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PayloadTypes) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the rawContent property value. The notification content of a raw user notification that will be delivered to and consumed by the app client on all supported platforms (Windows, iOS, Android or WebPush) receiving this notification. At least one of Payload.RawContent or Payload.VisualContent needs to be valid for a POST Notification request.
func (m *PayloadTypes) GetRawContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rawContent
    }
}
// Gets the visualContent property value. The visual content of a visual user notification, which will be consumed by the notification platform on each supported platform (Windows, iOS and Android only) and rendered for the user. At least one of Payload.RawContent or Payload.VisualContent needs to be valid for a POST Notification request.
func (m *PayloadTypes) GetVisualContent()(*VisualProperties) {
    if m == nil {
        return nil
    } else {
        return m.visualContent
    }
}
// The deserialization information for the current model
func (m *PayloadTypes) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["rawContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRawContent(val)
        return nil
    }
    res["visualContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVisualProperties() })
        if err != nil {
            return err
        }
        m.SetVisualContent(val.(*VisualProperties))
        return nil
    }
    return res
}
func (m *PayloadTypes) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PayloadTypes) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("rawContent", m.GetRawContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("visualContent", m.GetVisualContent())
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
func (m *PayloadTypes) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the rawContent property value. The notification content of a raw user notification that will be delivered to and consumed by the app client on all supported platforms (Windows, iOS, Android or WebPush) receiving this notification. At least one of Payload.RawContent or Payload.VisualContent needs to be valid for a POST Notification request.
// Parameters:
//  - value : Value to set for the rawContent property.
func (m *PayloadTypes) SetRawContent(value *string)() {
    m.rawContent = value
}
// Sets the visualContent property value. The visual content of a visual user notification, which will be consumed by the notification platform on each supported platform (Windows, iOS and Android only) and rendered for the user. At least one of Payload.RawContent or Payload.VisualContent needs to be valid for a POST Notification request.
// Parameters:
//  - value : Value to set for the visualContent property.
func (m *PayloadTypes) SetVisualContent(value *VisualProperties)() {
    m.visualContent = value
}
