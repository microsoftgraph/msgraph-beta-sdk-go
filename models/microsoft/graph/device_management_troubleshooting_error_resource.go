package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementTroubleshootingErrorResource 
type DeviceManagementTroubleshootingErrorResource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The link to the web resource. Can contain any of the following formatters: {{UPN}}, {{DeviceGUID}}, {{UserGUID}}
    link *string;
    // Not yet documented
    text *string;
}
// NewDeviceManagementTroubleshootingErrorResource instantiates a new deviceManagementTroubleshootingErrorResource and sets the default values.
func NewDeviceManagementTroubleshootingErrorResource()(*DeviceManagementTroubleshootingErrorResource) {
    m := &DeviceManagementTroubleshootingErrorResource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementTroubleshootingErrorResource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetLink gets the link property value. The link to the web resource. Can contain any of the following formatters: {{UPN}}, {{DeviceGUID}}, {{UserGUID}}
func (m *DeviceManagementTroubleshootingErrorResource) GetLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.link
    }
}
// GetText gets the text property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorResource) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementTroubleshootingErrorResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["link"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLink(val)
        }
        return nil
    }
    res["text"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementTroubleshootingErrorResource) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementTroubleshootingErrorResource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("link", m.GetLink())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("text", m.GetText())
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
func (m *DeviceManagementTroubleshootingErrorResource) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLink sets the link property value. The link to the web resource. Can contain any of the following formatters: {{UPN}}, {{DeviceGUID}}, {{UserGUID}}
func (m *DeviceManagementTroubleshootingErrorResource) SetLink(value *string)() {
    if m != nil {
        m.link = value
    }
}
// SetText sets the text property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorResource) SetText(value *string)() {
    if m != nil {
        m.text = value
    }
}
