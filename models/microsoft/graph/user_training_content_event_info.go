package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserTrainingContentEventInfo 
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
// NewUserTrainingContentEventInfo instantiates a new userTrainingContentEventInfo and sets the default values.
func NewUserTrainingContentEventInfo()(*UserTrainingContentEventInfo) {
    m := &UserTrainingContentEventInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserTrainingContentEventInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBrowser gets the browser property value. Browser of the user from where the training event was generated.
func (m *UserTrainingContentEventInfo) GetBrowser()(*string) {
    if m == nil {
        return nil
    } else {
        return m.browser
    }
}
// GetContentDateTime gets the contentDateTime property value. Date and time of the training content playback by the user.
func (m *UserTrainingContentEventInfo) GetContentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.contentDateTime
    }
}
// GetIpAddress gets the ipAddress property value. IP address of the user for the training event.
func (m *UserTrainingContentEventInfo) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
// GetOsPlatformDeviceDetails gets the osPlatformDeviceDetails property value. The operating system, platform, and device details of the user for the training event.
func (m *UserTrainingContentEventInfo) GetOsPlatformDeviceDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osPlatformDeviceDetails
    }
}
// GetPotentialScoreImpact gets the potentialScoreImpact property value. Potential improvement in security posture of the tenant after completion of the training by the user.
func (m *UserTrainingContentEventInfo) GetPotentialScoreImpact()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.potentialScoreImpact
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserTrainingContentEventInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBrowser sets the browser property value. Browser of the user from where the training event was generated.
func (m *UserTrainingContentEventInfo) SetBrowser(value *string)() {
    if m != nil {
        m.browser = value
    }
}
// SetContentDateTime sets the contentDateTime property value. Date and time of the training content playback by the user.
func (m *UserTrainingContentEventInfo) SetContentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.contentDateTime = value
    }
}
// SetIpAddress sets the ipAddress property value. IP address of the user for the training event.
func (m *UserTrainingContentEventInfo) SetIpAddress(value *string)() {
    if m != nil {
        m.ipAddress = value
    }
}
// SetOsPlatformDeviceDetails sets the osPlatformDeviceDetails property value. The operating system, platform, and device details of the user for the training event.
func (m *UserTrainingContentEventInfo) SetOsPlatformDeviceDetails(value *string)() {
    if m != nil {
        m.osPlatformDeviceDetails = value
    }
}
// SetPotentialScoreImpact sets the potentialScoreImpact property value. Potential improvement in security posture of the tenant after completion of the training by the user.
func (m *UserTrainingContentEventInfo) SetPotentialScoreImpact(value *float64)() {
    if m != nil {
        m.potentialScoreImpact = value
    }
}
