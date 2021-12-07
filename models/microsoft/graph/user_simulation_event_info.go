package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserSimulationEventInfo 
type UserSimulationEventInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Browser information from where the simulation event was initiated by a user in an attack simulation and training campaign.
    browser *string;
    // Date and time of the simulation event by a user in an attack simulation and training campaign.
    eventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Name of the simulation event by a user in an attack simulation and training campaign.
    eventName *string;
    // IP address from where the simulation event was initiated by a user in an attack simulation and training campaign.
    ipAddress *string;
    // The operating system, platform, and device details from where the simulation event was initiated by a user in an attack simulation and training campaign.
    osPlatformDeviceDetails *string;
}
// NewUserSimulationEventInfo instantiates a new userSimulationEventInfo and sets the default values.
func NewUserSimulationEventInfo()(*UserSimulationEventInfo) {
    m := &UserSimulationEventInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserSimulationEventInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBrowser gets the browser property value. Browser information from where the simulation event was initiated by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) GetBrowser()(*string) {
    if m == nil {
        return nil
    } else {
        return m.browser
    }
}
// GetEventDateTime gets the eventDateTime property value. Date and time of the simulation event by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.eventDateTime
    }
}
// GetEventName gets the eventName property value. Name of the simulation event by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) GetEventName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventName
    }
}
// GetIpAddress gets the ipAddress property value. IP address from where the simulation event was initiated by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
// GetOsPlatformDeviceDetails gets the osPlatformDeviceDetails property value. The operating system, platform, and device details from where the simulation event was initiated by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) GetOsPlatformDeviceDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osPlatformDeviceDetails
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserSimulationEventInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["browser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowser(val)
        }
        return nil
    }
    res["eventDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDateTime(val)
        }
        return nil
    }
    res["eventName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventName(val)
        }
        return nil
    }
    res["ipAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["osPlatformDeviceDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsPlatformDeviceDetails(val)
        }
        return nil
    }
    return res
}
func (m *UserSimulationEventInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserSimulationEventInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("browser", m.GetBrowser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventName", m.GetEventName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("osPlatformDeviceDetails", m.GetOsPlatformDeviceDetails())
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
func (m *UserSimulationEventInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBrowser sets the browser property value. Browser information from where the simulation event was initiated by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) SetBrowser(value *string)() {
    if m != nil {
        m.browser = value
    }
}
// SetEventDateTime sets the eventDateTime property value. Date and time of the simulation event by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.eventDateTime = value
    }
}
// SetEventName sets the eventName property value. Name of the simulation event by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) SetEventName(value *string)() {
    if m != nil {
        m.eventName = value
    }
}
// SetIpAddress sets the ipAddress property value. IP address from where the simulation event was initiated by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) SetIpAddress(value *string)() {
    if m != nil {
        m.ipAddress = value
    }
}
// SetOsPlatformDeviceDetails sets the osPlatformDeviceDetails property value. The operating system, platform, and device details from where the simulation event was initiated by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) SetOsPlatformDeviceDetails(value *string)() {
    if m != nil {
        m.osPlatformDeviceDetails = value
    }
}
