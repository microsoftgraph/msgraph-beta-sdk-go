package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new userSimulationEventInfo and sets the default values.
func NewUserSimulationEventInfo()(*UserSimulationEventInfo) {
    m := &UserSimulationEventInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserSimulationEventInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the browser property value. Browser information from where the simulation event was initiated by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) GetBrowser()(*string) {
    if m == nil {
        return nil
    } else {
        return m.browser
    }
}
// Gets the eventDateTime property value. Date and time of the simulation event by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.eventDateTime
    }
}
// Gets the eventName property value. Name of the simulation event by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) GetEventName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventName
    }
}
// Gets the ipAddress property value. IP address from where the simulation event was initiated by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
// Gets the osPlatformDeviceDetails property value. The operating system, platform, and device details from where the simulation event was initiated by a user in an attack simulation and training campaign.
func (m *UserSimulationEventInfo) GetOsPlatformDeviceDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osPlatformDeviceDetails
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *UserSimulationEventInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the browser property value. Browser information from where the simulation event was initiated by a user in an attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the browser property.
func (m *UserSimulationEventInfo) SetBrowser(value *string)() {
    m.browser = value
}
// Sets the eventDateTime property value. Date and time of the simulation event by a user in an attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the eventDateTime property.
func (m *UserSimulationEventInfo) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventDateTime = value
}
// Sets the eventName property value. Name of the simulation event by a user in an attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the eventName property.
func (m *UserSimulationEventInfo) SetEventName(value *string)() {
    m.eventName = value
}
// Sets the ipAddress property value. IP address from where the simulation event was initiated by a user in an attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the ipAddress property.
func (m *UserSimulationEventInfo) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// Sets the osPlatformDeviceDetails property value. The operating system, platform, and device details from where the simulation event was initiated by a user in an attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the osPlatformDeviceDetails property.
func (m *UserSimulationEventInfo) SetOsPlatformDeviceDetails(value *string)() {
    m.osPlatformDeviceDetails = value
}
