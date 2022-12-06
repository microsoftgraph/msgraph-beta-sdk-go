package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkDeviceHealth 
type TeamworkDeviceHealth struct {
    Entity
    // The connection property
    connection TeamworkConnectionable
    // Identity of the user who created the device health document.
    createdBy IdentitySetable
    // The UTC date and time when the device health document was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Health details about the device hardware.
    hardwareHealth TeamworkHardwareHealthable
    // Identity of the user who last modified the device health details.
    lastModifiedBy IdentitySetable
    // The UTC date and time when the device health detail was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The login status of Microsoft Teams, Skype for Business, and Exchange.
    loginStatus TeamworkLoginStatusable
    // Health details about all peripherals (for example, speaker and microphone) attached to a device.
    peripheralsHealth TeamworkPeripheralsHealthable
    // Software updates available for the device.
    softwareUpdateHealth TeamworkSoftwareUpdateHealthable
}
// NewTeamworkDeviceHealth instantiates a new teamworkDeviceHealth and sets the default values.
func NewTeamworkDeviceHealth()(*TeamworkDeviceHealth) {
    m := &TeamworkDeviceHealth{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamworkDeviceHealthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkDeviceHealthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkDeviceHealth(), nil
}
// GetConnection gets the connection property value. The connection property
func (m *TeamworkDeviceHealth) GetConnection()(TeamworkConnectionable) {
    return m.connection
}
// GetCreatedBy gets the createdBy property value. Identity of the user who created the device health document.
func (m *TeamworkDeviceHealth) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The UTC date and time when the device health document was created.
func (m *TeamworkDeviceHealth) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDeviceHealth) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["connection"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkConnectionFromDiscriminatorValue , m.SetConnection)
    res["createdBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentitySetFromDiscriminatorValue , m.SetCreatedBy)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["hardwareHealth"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkHardwareHealthFromDiscriminatorValue , m.SetHardwareHealth)
    res["lastModifiedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentitySetFromDiscriminatorValue , m.SetLastModifiedBy)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["loginStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkLoginStatusFromDiscriminatorValue , m.SetLoginStatus)
    res["peripheralsHealth"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkPeripheralsHealthFromDiscriminatorValue , m.SetPeripheralsHealth)
    res["softwareUpdateHealth"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkSoftwareUpdateHealthFromDiscriminatorValue , m.SetSoftwareUpdateHealth)
    return res
}
// GetHardwareHealth gets the hardwareHealth property value. Health details about the device hardware.
func (m *TeamworkDeviceHealth) GetHardwareHealth()(TeamworkHardwareHealthable) {
    return m.hardwareHealth
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user who last modified the device health details.
func (m *TeamworkDeviceHealth) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The UTC date and time when the device health detail was last modified.
func (m *TeamworkDeviceHealth) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLoginStatus gets the loginStatus property value. The login status of Microsoft Teams, Skype for Business, and Exchange.
func (m *TeamworkDeviceHealth) GetLoginStatus()(TeamworkLoginStatusable) {
    return m.loginStatus
}
// GetPeripheralsHealth gets the peripheralsHealth property value. Health details about all peripherals (for example, speaker and microphone) attached to a device.
func (m *TeamworkDeviceHealth) GetPeripheralsHealth()(TeamworkPeripheralsHealthable) {
    return m.peripheralsHealth
}
// GetSoftwareUpdateHealth gets the softwareUpdateHealth property value. Software updates available for the device.
func (m *TeamworkDeviceHealth) GetSoftwareUpdateHealth()(TeamworkSoftwareUpdateHealthable) {
    return m.softwareUpdateHealth
}
// Serialize serializes information the current object
func (m *TeamworkDeviceHealth) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("connection", m.GetConnection())
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
        err = writer.WriteObjectValue("hardwareHealth", m.GetHardwareHealth())
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
        err = writer.WriteObjectValue("loginStatus", m.GetLoginStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("peripheralsHealth", m.GetPeripheralsHealth())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("softwareUpdateHealth", m.GetSoftwareUpdateHealth())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnection sets the connection property value. The connection property
func (m *TeamworkDeviceHealth) SetConnection(value TeamworkConnectionable)() {
    m.connection = value
}
// SetCreatedBy sets the createdBy property value. Identity of the user who created the device health document.
func (m *TeamworkDeviceHealth) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The UTC date and time when the device health document was created.
func (m *TeamworkDeviceHealth) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetHardwareHealth sets the hardwareHealth property value. Health details about the device hardware.
func (m *TeamworkDeviceHealth) SetHardwareHealth(value TeamworkHardwareHealthable)() {
    m.hardwareHealth = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user who last modified the device health details.
func (m *TeamworkDeviceHealth) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The UTC date and time when the device health detail was last modified.
func (m *TeamworkDeviceHealth) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLoginStatus sets the loginStatus property value. The login status of Microsoft Teams, Skype for Business, and Exchange.
func (m *TeamworkDeviceHealth) SetLoginStatus(value TeamworkLoginStatusable)() {
    m.loginStatus = value
}
// SetPeripheralsHealth sets the peripheralsHealth property value. Health details about all peripherals (for example, speaker and microphone) attached to a device.
func (m *TeamworkDeviceHealth) SetPeripheralsHealth(value TeamworkPeripheralsHealthable)() {
    m.peripheralsHealth = value
}
// SetSoftwareUpdateHealth sets the softwareUpdateHealth property value. Software updates available for the device.
func (m *TeamworkDeviceHealth) SetSoftwareUpdateHealth(value TeamworkSoftwareUpdateHealthable)() {
    m.softwareUpdateHealth = value
}
