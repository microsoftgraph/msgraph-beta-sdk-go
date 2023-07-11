package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkDeviceHealth 
type TeamworkDeviceHealth struct {
    Entity
    // The OdataType property
    OdataType *string
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
    val, err := m.GetBackingStore().Get("connection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkConnectionable)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. Identity of the user who created the device health document.
func (m *TeamworkDeviceHealth) GetCreatedBy()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The UTC date and time when the device health document was created.
func (m *TeamworkDeviceHealth) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDeviceHealth) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["connection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkConnectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnection(val.(TeamworkConnectionable))
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
    res["hardwareHealth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkHardwareHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHardwareHealth(val.(TeamworkHardwareHealthable))
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
    res["loginStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkLoginStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginStatus(val.(TeamworkLoginStatusable))
        }
        return nil
    }
    res["peripheralsHealth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralsHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeripheralsHealth(val.(TeamworkPeripheralsHealthable))
        }
        return nil
    }
    res["softwareUpdateHealth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkSoftwareUpdateHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareUpdateHealth(val.(TeamworkSoftwareUpdateHealthable))
        }
        return nil
    }
    return res
}
// GetHardwareHealth gets the hardwareHealth property value. Health details about the device hardware.
func (m *TeamworkDeviceHealth) GetHardwareHealth()(TeamworkHardwareHealthable) {
    val, err := m.GetBackingStore().Get("hardwareHealth")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkHardwareHealthable)
    }
    return nil
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user who last modified the device health details.
func (m *TeamworkDeviceHealth) GetLastModifiedBy()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("lastModifiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The UTC date and time when the device health detail was last modified.
func (m *TeamworkDeviceHealth) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLoginStatus gets the loginStatus property value. The login status of Microsoft Teams, Skype for Business, and Exchange.
func (m *TeamworkDeviceHealth) GetLoginStatus()(TeamworkLoginStatusable) {
    val, err := m.GetBackingStore().Get("loginStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkLoginStatusable)
    }
    return nil
}
// GetPeripheralsHealth gets the peripheralsHealth property value. Health details about all peripherals (for example, speaker and microphone) attached to a device.
func (m *TeamworkDeviceHealth) GetPeripheralsHealth()(TeamworkPeripheralsHealthable) {
    val, err := m.GetBackingStore().Get("peripheralsHealth")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkPeripheralsHealthable)
    }
    return nil
}
// GetSoftwareUpdateHealth gets the softwareUpdateHealth property value. Software updates available for the device.
func (m *TeamworkDeviceHealth) GetSoftwareUpdateHealth()(TeamworkSoftwareUpdateHealthable) {
    val, err := m.GetBackingStore().Get("softwareUpdateHealth")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkSoftwareUpdateHealthable)
    }
    return nil
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
    err := m.GetBackingStore().Set("connection", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. Identity of the user who created the device health document.
func (m *TeamworkDeviceHealth) SetCreatedBy(value IdentitySetable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The UTC date and time when the device health document was created.
func (m *TeamworkDeviceHealth) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetHardwareHealth sets the hardwareHealth property value. Health details about the device hardware.
func (m *TeamworkDeviceHealth) SetHardwareHealth(value TeamworkHardwareHealthable)() {
    err := m.GetBackingStore().Set("hardwareHealth", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user who last modified the device health details.
func (m *TeamworkDeviceHealth) SetLastModifiedBy(value IdentitySetable)() {
    err := m.GetBackingStore().Set("lastModifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The UTC date and time when the device health detail was last modified.
func (m *TeamworkDeviceHealth) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLoginStatus sets the loginStatus property value. The login status of Microsoft Teams, Skype for Business, and Exchange.
func (m *TeamworkDeviceHealth) SetLoginStatus(value TeamworkLoginStatusable)() {
    err := m.GetBackingStore().Set("loginStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetPeripheralsHealth sets the peripheralsHealth property value. Health details about all peripherals (for example, speaker and microphone) attached to a device.
func (m *TeamworkDeviceHealth) SetPeripheralsHealth(value TeamworkPeripheralsHealthable)() {
    err := m.GetBackingStore().Set("peripheralsHealth", value)
    if err != nil {
        panic(err)
    }
}
// SetSoftwareUpdateHealth sets the softwareUpdateHealth property value. Software updates available for the device.
func (m *TeamworkDeviceHealth) SetSoftwareUpdateHealth(value TeamworkSoftwareUpdateHealthable)() {
    err := m.GetBackingStore().Set("softwareUpdateHealth", value)
    if err != nil {
        panic(err)
    }
}
// TeamworkDeviceHealthable 
type TeamworkDeviceHealthable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnection()(TeamworkConnectionable)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHardwareHealth()(TeamworkHardwareHealthable)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLoginStatus()(TeamworkLoginStatusable)
    GetPeripheralsHealth()(TeamworkPeripheralsHealthable)
    GetSoftwareUpdateHealth()(TeamworkSoftwareUpdateHealthable)
    SetConnection(value TeamworkConnectionable)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHardwareHealth(value TeamworkHardwareHealthable)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLoginStatus(value TeamworkLoginStatusable)()
    SetPeripheralsHealth(value TeamworkPeripheralsHealthable)()
    SetSoftwareUpdateHealth(value TeamworkSoftwareUpdateHealthable)()
}
