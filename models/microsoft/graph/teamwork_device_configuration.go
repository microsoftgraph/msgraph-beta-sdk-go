package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkDeviceConfiguration 
type TeamworkDeviceConfiguration struct {
    Entity
    // The camera configuration. Applicable only for Microsoft Teams Rooms-enabled devices.
    cameraConfiguration *TeamworkCameraConfiguration;
    // Identity of the user who created the device configuration document.
    createdBy *IdentitySet;
    // The UTC date and time when the device configuration document was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The display configuration.
    displayConfiguration *TeamworkDisplayConfiguration;
    // The hardware configuration. Applicable only for Teams Rooms-enabled devices.
    hardwareConfiguration *TeamworkHardwareConfiguration;
    // Identity of the user who last modified the device configuration.
    lastModifiedBy *IdentitySet;
    // The UTC date and time when the device configuration was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The microphone configuration. Applicable only for Teams Rooms-enabled devices.
    microphoneConfiguration *TeamworkMicrophoneConfiguration;
    // Information related to software versions for the device, such as firmware, operating system, Teams client, and admin agent.
    softwareVersions *TeamworkDeviceSoftwareVersions;
    // The speaker configuration. Applicable only for Teams Rooms-enabled devices.
    speakerConfiguration *TeamworkSpeakerConfiguration;
    // The system configuration. Not applicable for Teams Rooms-enabled devices.
    systemConfiguration *TeamworkSystemConfiguration;
    // The Teams client configuration. Applicable only for Teams Rooms-enabled devices.
    teamsClientConfiguration *TeamworkTeamsClientConfiguration;
}
// NewTeamworkDeviceConfiguration instantiates a new teamworkDeviceConfiguration and sets the default values.
func NewTeamworkDeviceConfiguration()(*TeamworkDeviceConfiguration) {
    m := &TeamworkDeviceConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// GetCameraConfiguration gets the cameraConfiguration property value. The camera configuration. Applicable only for Microsoft Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) GetCameraConfiguration()(*TeamworkCameraConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.cameraConfiguration
    }
}
// GetCreatedBy gets the createdBy property value. Identity of the user who created the device configuration document.
func (m *TeamworkDeviceConfiguration) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The UTC date and time when the device configuration document was created.
func (m *TeamworkDeviceConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDisplayConfiguration gets the displayConfiguration property value. The display configuration.
func (m *TeamworkDeviceConfiguration) GetDisplayConfiguration()(*TeamworkDisplayConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.displayConfiguration
    }
}
// GetHardwareConfiguration gets the hardwareConfiguration property value. The hardware configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) GetHardwareConfiguration()(*TeamworkHardwareConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.hardwareConfiguration
    }
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user who last modified the device configuration.
func (m *TeamworkDeviceConfiguration) GetLastModifiedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The UTC date and time when the device configuration was last modified.
func (m *TeamworkDeviceConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetMicrophoneConfiguration gets the microphoneConfiguration property value. The microphone configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) GetMicrophoneConfiguration()(*TeamworkMicrophoneConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.microphoneConfiguration
    }
}
// GetSoftwareVersions gets the softwareVersions property value. Information related to software versions for the device, such as firmware, operating system, Teams client, and admin agent.
func (m *TeamworkDeviceConfiguration) GetSoftwareVersions()(*TeamworkDeviceSoftwareVersions) {
    if m == nil {
        return nil
    } else {
        return m.softwareVersions
    }
}
// GetSpeakerConfiguration gets the speakerConfiguration property value. The speaker configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) GetSpeakerConfiguration()(*TeamworkSpeakerConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.speakerConfiguration
    }
}
// GetSystemConfiguration gets the systemConfiguration property value. The system configuration. Not applicable for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) GetSystemConfiguration()(*TeamworkSystemConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.systemConfiguration
    }
}
// GetTeamsClientConfiguration gets the teamsClientConfiguration property value. The Teams client configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) GetTeamsClientConfiguration()(*TeamworkTeamsClientConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.teamsClientConfiguration
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDeviceConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cameraConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkCameraConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCameraConfiguration(val.(*TeamworkCameraConfiguration))
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["displayConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkDisplayConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayConfiguration(val.(*TeamworkDisplayConfiguration))
        }
        return nil
    }
    res["hardwareConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkHardwareConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHardwareConfiguration(val.(*TeamworkHardwareConfiguration))
        }
        return nil
    }
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["microphoneConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkMicrophoneConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrophoneConfiguration(val.(*TeamworkMicrophoneConfiguration))
        }
        return nil
    }
    res["softwareVersions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkDeviceSoftwareVersions() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareVersions(val.(*TeamworkDeviceSoftwareVersions))
        }
        return nil
    }
    res["speakerConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkSpeakerConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeakerConfiguration(val.(*TeamworkSpeakerConfiguration))
        }
        return nil
    }
    res["systemConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkSystemConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemConfiguration(val.(*TeamworkSystemConfiguration))
        }
        return nil
    }
    res["teamsClientConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkTeamsClientConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsClientConfiguration(val.(*TeamworkTeamsClientConfiguration))
        }
        return nil
    }
    return res
}
func (m *TeamworkDeviceConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkDeviceConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("cameraConfiguration", m.GetCameraConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("displayConfiguration", m.GetDisplayConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("hardwareConfiguration", m.GetHardwareConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("microphoneConfiguration", m.GetMicrophoneConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("softwareVersions", m.GetSoftwareVersions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("speakerConfiguration", m.GetSpeakerConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("systemConfiguration", m.GetSystemConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teamsClientConfiguration", m.GetTeamsClientConfiguration())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCameraConfiguration sets the cameraConfiguration property value. The camera configuration. Applicable only for Microsoft Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) SetCameraConfiguration(value *TeamworkCameraConfiguration)() {
    if m != nil {
        m.cameraConfiguration = value
    }
}
// SetCreatedBy sets the createdBy property value. Identity of the user who created the device configuration document.
func (m *TeamworkDeviceConfiguration) SetCreatedBy(value *IdentitySet)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The UTC date and time when the device configuration document was created.
func (m *TeamworkDeviceConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDisplayConfiguration sets the displayConfiguration property value. The display configuration.
func (m *TeamworkDeviceConfiguration) SetDisplayConfiguration(value *TeamworkDisplayConfiguration)() {
    if m != nil {
        m.displayConfiguration = value
    }
}
// SetHardwareConfiguration sets the hardwareConfiguration property value. The hardware configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) SetHardwareConfiguration(value *TeamworkHardwareConfiguration)() {
    if m != nil {
        m.hardwareConfiguration = value
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user who last modified the device configuration.
func (m *TeamworkDeviceConfiguration) SetLastModifiedBy(value *IdentitySet)() {
    if m != nil {
        m.lastModifiedBy = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The UTC date and time when the device configuration was last modified.
func (m *TeamworkDeviceConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetMicrophoneConfiguration sets the microphoneConfiguration property value. The microphone configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) SetMicrophoneConfiguration(value *TeamworkMicrophoneConfiguration)() {
    if m != nil {
        m.microphoneConfiguration = value
    }
}
// SetSoftwareVersions sets the softwareVersions property value. Information related to software versions for the device, such as firmware, operating system, Teams client, and admin agent.
func (m *TeamworkDeviceConfiguration) SetSoftwareVersions(value *TeamworkDeviceSoftwareVersions)() {
    if m != nil {
        m.softwareVersions = value
    }
}
// SetSpeakerConfiguration sets the speakerConfiguration property value. The speaker configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) SetSpeakerConfiguration(value *TeamworkSpeakerConfiguration)() {
    if m != nil {
        m.speakerConfiguration = value
    }
}
// SetSystemConfiguration sets the systemConfiguration property value. The system configuration. Not applicable for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) SetSystemConfiguration(value *TeamworkSystemConfiguration)() {
    if m != nil {
        m.systemConfiguration = value
    }
}
// SetTeamsClientConfiguration sets the teamsClientConfiguration property value. The Teams client configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) SetTeamsClientConfiguration(value *TeamworkTeamsClientConfiguration)() {
    if m != nil {
        m.teamsClientConfiguration = value
    }
}
