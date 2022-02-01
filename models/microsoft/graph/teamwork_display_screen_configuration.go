package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkDisplayScreenConfiguration 
type TeamworkDisplayScreenConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The brightness level on the device (0-100). Not applicable for Microsoft Teams Rooms devices.
    backlightBrightness *int32;
    // Timeout for backlight (30-3600 secs). Not applicable for Teams Rooms devices.
    backlightTimeout *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // True if high contrast mode is enabled. Not applicable for Teams Rooms devices.
    isHighContrastEnabled *bool;
    // True if screensaver is enabled. Not applicable for Teams Rooms devices.
    isScreensaverEnabled *bool;
    // Screensaver timeout from 30 to 3600 secs. Not applicable for Teams Rooms devices.
    screensaverTimeout *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
}
// NewTeamworkDisplayScreenConfiguration instantiates a new teamworkDisplayScreenConfiguration and sets the default values.
func NewTeamworkDisplayScreenConfiguration()(*TeamworkDisplayScreenConfiguration) {
    m := &TeamworkDisplayScreenConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkDisplayScreenConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBacklightBrightness gets the backlightBrightness property value. The brightness level on the device (0-100). Not applicable for Microsoft Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) GetBacklightBrightness()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.backlightBrightness
    }
}
// GetBacklightTimeout gets the backlightTimeout property value. Timeout for backlight (30-3600 secs). Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) GetBacklightTimeout()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.backlightTimeout
    }
}
// GetIsHighContrastEnabled gets the isHighContrastEnabled property value. True if high contrast mode is enabled. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) GetIsHighContrastEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHighContrastEnabled
    }
}
// GetIsScreensaverEnabled gets the isScreensaverEnabled property value. True if screensaver is enabled. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) GetIsScreensaverEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isScreensaverEnabled
    }
}
// GetScreensaverTimeout gets the screensaverTimeout property value. Screensaver timeout from 30 to 3600 secs. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) GetScreensaverTimeout()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.screensaverTimeout
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDisplayScreenConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["backlightBrightness"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBacklightBrightness(val)
        }
        return nil
    }
    res["backlightTimeout"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBacklightTimeout(val)
        }
        return nil
    }
    res["isHighContrastEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHighContrastEnabled(val)
        }
        return nil
    }
    res["isScreensaverEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsScreensaverEnabled(val)
        }
        return nil
    }
    res["screensaverTimeout"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScreensaverTimeout(val)
        }
        return nil
    }
    return res
}
func (m *TeamworkDisplayScreenConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkDisplayScreenConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("backlightBrightness", m.GetBacklightBrightness())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("backlightTimeout", m.GetBacklightTimeout())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isHighContrastEnabled", m.GetIsHighContrastEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isScreensaverEnabled", m.GetIsScreensaverEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("screensaverTimeout", m.GetScreensaverTimeout())
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
func (m *TeamworkDisplayScreenConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBacklightBrightness sets the backlightBrightness property value. The brightness level on the device (0-100). Not applicable for Microsoft Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) SetBacklightBrightness(value *int32)() {
    if m != nil {
        m.backlightBrightness = value
    }
}
// SetBacklightTimeout sets the backlightTimeout property value. Timeout for backlight (30-3600 secs). Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) SetBacklightTimeout(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.backlightTimeout = value
    }
}
// SetIsHighContrastEnabled sets the isHighContrastEnabled property value. True if high contrast mode is enabled. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) SetIsHighContrastEnabled(value *bool)() {
    if m != nil {
        m.isHighContrastEnabled = value
    }
}
// SetIsScreensaverEnabled sets the isScreensaverEnabled property value. True if screensaver is enabled. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) SetIsScreensaverEnabled(value *bool)() {
    if m != nil {
        m.isScreensaverEnabled = value
    }
}
// SetScreensaverTimeout sets the screensaverTimeout property value. Screensaver timeout from 30 to 3600 secs. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) SetScreensaverTimeout(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.screensaverTimeout = value
    }
}
