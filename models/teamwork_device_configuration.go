package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TeamworkDeviceConfiguration struct {
    Entity
}
// NewTeamworkDeviceConfiguration instantiates a new TeamworkDeviceConfiguration and sets the default values.
func NewTeamworkDeviceConfiguration()(*TeamworkDeviceConfiguration) {
    m := &TeamworkDeviceConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamworkDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamworkDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkDeviceConfiguration(), nil
}
// GetCameraConfiguration gets the cameraConfiguration property value. The camera configuration. Applicable only for Microsoft Teams Rooms-enabled devices.
// returns a TeamworkCameraConfigurationable when successful
func (m *TeamworkDeviceConfiguration) GetCameraConfiguration()(TeamworkCameraConfigurationable) {
    val, err := m.GetBackingStore().Get("cameraConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkCameraConfigurationable)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. Identity of the user who created the device configuration document.
// returns a IdentitySetable when successful
func (m *TeamworkDeviceConfiguration) GetCreatedBy()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The UTC date and time when the device configuration document was created.
// returns a *Time when successful
func (m *TeamworkDeviceConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDisplayConfiguration gets the displayConfiguration property value. The display configuration.
// returns a TeamworkDisplayConfigurationable when successful
func (m *TeamworkDeviceConfiguration) GetDisplayConfiguration()(TeamworkDisplayConfigurationable) {
    val, err := m.GetBackingStore().Get("displayConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkDisplayConfigurationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamworkDeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cameraConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkCameraConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCameraConfiguration(val.(TeamworkCameraConfigurationable))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["displayConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkDisplayConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayConfiguration(val.(TeamworkDisplayConfigurationable))
        }
        return nil
    }
    res["hardwareConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkHardwareConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHardwareConfiguration(val.(TeamworkHardwareConfigurationable))
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["microphoneConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkMicrophoneConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrophoneConfiguration(val.(TeamworkMicrophoneConfigurationable))
        }
        return nil
    }
    res["softwareVersions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkDeviceSoftwareVersionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareVersions(val.(TeamworkDeviceSoftwareVersionsable))
        }
        return nil
    }
    res["speakerConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkSpeakerConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeakerConfiguration(val.(TeamworkSpeakerConfigurationable))
        }
        return nil
    }
    res["systemConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkSystemConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemConfiguration(val.(TeamworkSystemConfigurationable))
        }
        return nil
    }
    res["teamsClientConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkTeamsClientConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsClientConfiguration(val.(TeamworkTeamsClientConfigurationable))
        }
        return nil
    }
    return res
}
// GetHardwareConfiguration gets the hardwareConfiguration property value. The hardware configuration. Applicable only for Teams Rooms-enabled devices.
// returns a TeamworkHardwareConfigurationable when successful
func (m *TeamworkDeviceConfiguration) GetHardwareConfiguration()(TeamworkHardwareConfigurationable) {
    val, err := m.GetBackingStore().Get("hardwareConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkHardwareConfigurationable)
    }
    return nil
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user who last modified the device configuration.
// returns a IdentitySetable when successful
func (m *TeamworkDeviceConfiguration) GetLastModifiedBy()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("lastModifiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The UTC date and time when the device configuration was last modified.
// returns a *Time when successful
func (m *TeamworkDeviceConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMicrophoneConfiguration gets the microphoneConfiguration property value. The microphone configuration. Applicable only for Teams Rooms-enabled devices.
// returns a TeamworkMicrophoneConfigurationable when successful
func (m *TeamworkDeviceConfiguration) GetMicrophoneConfiguration()(TeamworkMicrophoneConfigurationable) {
    val, err := m.GetBackingStore().Get("microphoneConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkMicrophoneConfigurationable)
    }
    return nil
}
// GetSoftwareVersions gets the softwareVersions property value. Information related to software versions for the device, such as firmware, operating system, Teams client, and admin agent.
// returns a TeamworkDeviceSoftwareVersionsable when successful
func (m *TeamworkDeviceConfiguration) GetSoftwareVersions()(TeamworkDeviceSoftwareVersionsable) {
    val, err := m.GetBackingStore().Get("softwareVersions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkDeviceSoftwareVersionsable)
    }
    return nil
}
// GetSpeakerConfiguration gets the speakerConfiguration property value. The speaker configuration. Applicable only for Teams Rooms-enabled devices.
// returns a TeamworkSpeakerConfigurationable when successful
func (m *TeamworkDeviceConfiguration) GetSpeakerConfiguration()(TeamworkSpeakerConfigurationable) {
    val, err := m.GetBackingStore().Get("speakerConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkSpeakerConfigurationable)
    }
    return nil
}
// GetSystemConfiguration gets the systemConfiguration property value. The system configuration. Not applicable for Teams Rooms-enabled devices.
// returns a TeamworkSystemConfigurationable when successful
func (m *TeamworkDeviceConfiguration) GetSystemConfiguration()(TeamworkSystemConfigurationable) {
    val, err := m.GetBackingStore().Get("systemConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkSystemConfigurationable)
    }
    return nil
}
// GetTeamsClientConfiguration gets the teamsClientConfiguration property value. The Teams client configuration. Applicable only for Teams Rooms-enabled devices.
// returns a TeamworkTeamsClientConfigurationable when successful
func (m *TeamworkDeviceConfiguration) GetTeamsClientConfiguration()(TeamworkTeamsClientConfigurationable) {
    val, err := m.GetBackingStore().Get("teamsClientConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkTeamsClientConfigurationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamworkDeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TeamworkDeviceConfiguration) SetCameraConfiguration(value TeamworkCameraConfigurationable)() {
    err := m.GetBackingStore().Set("cameraConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. Identity of the user who created the device configuration document.
func (m *TeamworkDeviceConfiguration) SetCreatedBy(value IdentitySetable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The UTC date and time when the device configuration document was created.
func (m *TeamworkDeviceConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayConfiguration sets the displayConfiguration property value. The display configuration.
func (m *TeamworkDeviceConfiguration) SetDisplayConfiguration(value TeamworkDisplayConfigurationable)() {
    err := m.GetBackingStore().Set("displayConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetHardwareConfiguration sets the hardwareConfiguration property value. The hardware configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) SetHardwareConfiguration(value TeamworkHardwareConfigurationable)() {
    err := m.GetBackingStore().Set("hardwareConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user who last modified the device configuration.
func (m *TeamworkDeviceConfiguration) SetLastModifiedBy(value IdentitySetable)() {
    err := m.GetBackingStore().Set("lastModifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The UTC date and time when the device configuration was last modified.
func (m *TeamworkDeviceConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrophoneConfiguration sets the microphoneConfiguration property value. The microphone configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) SetMicrophoneConfiguration(value TeamworkMicrophoneConfigurationable)() {
    err := m.GetBackingStore().Set("microphoneConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetSoftwareVersions sets the softwareVersions property value. Information related to software versions for the device, such as firmware, operating system, Teams client, and admin agent.
func (m *TeamworkDeviceConfiguration) SetSoftwareVersions(value TeamworkDeviceSoftwareVersionsable)() {
    err := m.GetBackingStore().Set("softwareVersions", value)
    if err != nil {
        panic(err)
    }
}
// SetSpeakerConfiguration sets the speakerConfiguration property value. The speaker configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) SetSpeakerConfiguration(value TeamworkSpeakerConfigurationable)() {
    err := m.GetBackingStore().Set("speakerConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemConfiguration sets the systemConfiguration property value. The system configuration. Not applicable for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) SetSystemConfiguration(value TeamworkSystemConfigurationable)() {
    err := m.GetBackingStore().Set("systemConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamsClientConfiguration sets the teamsClientConfiguration property value. The Teams client configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) SetTeamsClientConfiguration(value TeamworkTeamsClientConfigurationable)() {
    err := m.GetBackingStore().Set("teamsClientConfiguration", value)
    if err != nil {
        panic(err)
    }
}
type TeamworkDeviceConfigurationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCameraConfiguration()(TeamworkCameraConfigurationable)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayConfiguration()(TeamworkDisplayConfigurationable)
    GetHardwareConfiguration()(TeamworkHardwareConfigurationable)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMicrophoneConfiguration()(TeamworkMicrophoneConfigurationable)
    GetSoftwareVersions()(TeamworkDeviceSoftwareVersionsable)
    GetSpeakerConfiguration()(TeamworkSpeakerConfigurationable)
    GetSystemConfiguration()(TeamworkSystemConfigurationable)
    GetTeamsClientConfiguration()(TeamworkTeamsClientConfigurationable)
    SetCameraConfiguration(value TeamworkCameraConfigurationable)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayConfiguration(value TeamworkDisplayConfigurationable)()
    SetHardwareConfiguration(value TeamworkHardwareConfigurationable)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMicrophoneConfiguration(value TeamworkMicrophoneConfigurationable)()
    SetSoftwareVersions(value TeamworkDeviceSoftwareVersionsable)()
    SetSpeakerConfiguration(value TeamworkSpeakerConfigurationable)()
    SetSystemConfiguration(value TeamworkSystemConfigurationable)()
    SetTeamsClientConfiguration(value TeamworkTeamsClientConfigurationable)()
}
