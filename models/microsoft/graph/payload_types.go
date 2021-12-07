package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PayloadTypes 
type PayloadTypes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The notification content of a raw user notification that will be delivered to and consumed by the app client on all supported platforms (Windows, iOS, Android or WebPush) receiving this notification. At least one of Payload.RawContent or Payload.VisualContent needs to be valid for a POST Notification request.
    rawContent *string;
    // The visual content of a visual user notification, which will be consumed by the notification platform on each supported platform (Windows, iOS and Android only) and rendered for the user. At least one of Payload.RawContent or Payload.VisualContent needs to be valid for a POST Notification request.
    visualContent *VisualProperties;
}
// NewPayloadTypes instantiates a new payloadTypes and sets the default values.
func NewPayloadTypes()(*PayloadTypes) {
    m := &PayloadTypes{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PayloadTypes) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetRawContent gets the rawContent property value. The notification content of a raw user notification that will be delivered to and consumed by the app client on all supported platforms (Windows, iOS, Android or WebPush) receiving this notification. At least one of Payload.RawContent or Payload.VisualContent needs to be valid for a POST Notification request.
func (m *PayloadTypes) GetRawContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rawContent
    }
}
// GetVisualContent gets the visualContent property value. The visual content of a visual user notification, which will be consumed by the notification platform on each supported platform (Windows, iOS and Android only) and rendered for the user. At least one of Payload.RawContent or Payload.VisualContent needs to be valid for a POST Notification request.
func (m *PayloadTypes) GetVisualContent()(*VisualProperties) {
    if m == nil {
        return nil
    } else {
        return m.visualContent
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PayloadTypes) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["rawContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRawContent(val)
        }
        return nil
    }
    res["visualContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVisualProperties() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisualContent(val.(*VisualProperties))
        }
        return nil
    }
    return res
}
func (m *PayloadTypes) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PayloadTypes) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRawContent sets the rawContent property value. The notification content of a raw user notification that will be delivered to and consumed by the app client on all supported platforms (Windows, iOS, Android or WebPush) receiving this notification. At least one of Payload.RawContent or Payload.VisualContent needs to be valid for a POST Notification request.
func (m *PayloadTypes) SetRawContent(value *string)() {
    if m != nil {
        m.rawContent = value
    }
}
// SetVisualContent sets the visualContent property value. The visual content of a visual user notification, which will be consumed by the notification platform on each supported platform (Windows, iOS and Android only) and rendered for the user. At least one of Payload.RawContent or Payload.VisualContent needs to be valid for a POST Notification request.
func (m *PayloadTypes) SetVisualContent(value *VisualProperties)() {
    if m != nil {
        m.visualContent = value
    }
}
