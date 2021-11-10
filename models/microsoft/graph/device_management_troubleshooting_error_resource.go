package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementTroubleshootingErrorResource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The link to the web resource. Can contain any of the following formatters: {{UPN}}, {{DeviceGUID}}, {{UserGUID}}
    link *string;
    // Not yet documented
    text *string;
}
// Instantiates a new deviceManagementTroubleshootingErrorResource and sets the default values.
func NewDeviceManagementTroubleshootingErrorResource()(*DeviceManagementTroubleshootingErrorResource) {
    m := &DeviceManagementTroubleshootingErrorResource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementTroubleshootingErrorResource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the link property value. The link to the web resource. Can contain any of the following formatters: {{UPN}}, {{DeviceGUID}}, {{UserGUID}}
func (m *DeviceManagementTroubleshootingErrorResource) GetLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.link
    }
}
// Gets the text property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorResource) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DeviceManagementTroubleshootingErrorResource) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the link property value. The link to the web resource. Can contain any of the following formatters: {{UPN}}, {{DeviceGUID}}, {{UserGUID}}
// Parameters:
//  - value : Value to set for the link property.
func (m *DeviceManagementTroubleshootingErrorResource) SetLink(value *string)() {
    m.link = value
}
// Sets the text property value. Not yet documented
// Parameters:
//  - value : Value to set for the text property.
func (m *DeviceManagementTroubleshootingErrorResource) SetText(value *string)() {
    m.text = value
}
