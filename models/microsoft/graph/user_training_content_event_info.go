package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserTrainingContentEventInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Browser of the user from where the training event was generated.
    browser *string;
    // Date and time of the training content playback by the user.
    contentDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // IP address of the user for the training event.
    ipAddress *string;
    // The operating system, platform, and device details of the user for the training event.
    osPlatformDeviceDetails *string;
    // Potential improvement in security posture of the tenant after completion of the training by the user.
    potentialScoreImpact *float64;
}
// Instantiates a new userTrainingContentEventInfo and sets the default values.
func NewUserTrainingContentEventInfo()(*UserTrainingContentEventInfo) {
    m := &UserTrainingContentEventInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserTrainingContentEventInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the browser property value. Browser of the user from where the training event was generated.
func (m *UserTrainingContentEventInfo) GetBrowser()(*string) {
    if m == nil {
        return nil
    } else {
        return m.browser
    }
}
// Gets the contentDateTime property value. Date and time of the training content playback by the user.
func (m *UserTrainingContentEventInfo) GetContentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.contentDateTime
    }
}
// Gets the ipAddress property value. IP address of the user for the training event.
func (m *UserTrainingContentEventInfo) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
// Gets the osPlatformDeviceDetails property value. The operating system, platform, and device details of the user for the training event.
func (m *UserTrainingContentEventInfo) GetOsPlatformDeviceDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osPlatformDeviceDetails
    }
}
// Gets the potentialScoreImpact property value. Potential improvement in security posture of the tenant after completion of the training by the user.
func (m *UserTrainingContentEventInfo) GetPotentialScoreImpact()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.potentialScoreImpact
    }
}
// The deserialization information for the current model
func (m *UserTrainingContentEventInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["contentDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentDateTime(val)
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
    res["potentialScoreImpact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPotentialScoreImpact(val)
        }
        return nil
    }
    return res
}
func (m *UserTrainingContentEventInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserTrainingContentEventInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("browser", m.GetBrowser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("contentDateTime", m.GetContentDateTime())
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
        err := writer.WriteFloat64Value("potentialScoreImpact", m.GetPotentialScoreImpact())
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
func (m *UserTrainingContentEventInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the browser property value. Browser of the user from where the training event was generated.
// Parameters:
//  - value : Value to set for the browser property.
func (m *UserTrainingContentEventInfo) SetBrowser(value *string)() {
    m.browser = value
}
// Sets the contentDateTime property value. Date and time of the training content playback by the user.
// Parameters:
//  - value : Value to set for the contentDateTime property.
func (m *UserTrainingContentEventInfo) SetContentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.contentDateTime = value
}
// Sets the ipAddress property value. IP address of the user for the training event.
// Parameters:
//  - value : Value to set for the ipAddress property.
func (m *UserTrainingContentEventInfo) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// Sets the osPlatformDeviceDetails property value. The operating system, platform, and device details of the user for the training event.
// Parameters:
//  - value : Value to set for the osPlatformDeviceDetails property.
func (m *UserTrainingContentEventInfo) SetOsPlatformDeviceDetails(value *string)() {
    m.osPlatformDeviceDetails = value
}
// Sets the potentialScoreImpact property value. Potential improvement in security posture of the tenant after completion of the training by the user.
// Parameters:
//  - value : Value to set for the potentialScoreImpact property.
func (m *UserTrainingContentEventInfo) SetPotentialScoreImpact(value *float64)() {
    m.potentialScoreImpact = value
}
