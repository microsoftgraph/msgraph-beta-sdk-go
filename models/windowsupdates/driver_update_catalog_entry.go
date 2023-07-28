package windowsupdates

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DriverUpdateCatalogEntry 
type DriverUpdateCatalogEntry struct {
    SoftwareUpdateCatalogEntry
}
// NewDriverUpdateCatalogEntry instantiates a new driverUpdateCatalogEntry and sets the default values.
func NewDriverUpdateCatalogEntry()(*DriverUpdateCatalogEntry) {
    m := &DriverUpdateCatalogEntry{
        SoftwareUpdateCatalogEntry: *NewSoftwareUpdateCatalogEntry(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.driverUpdateCatalogEntry"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDriverUpdateCatalogEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDriverUpdateCatalogEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDriverUpdateCatalogEntry(), nil
}
// GetDescription gets the description property value. The description of the content.
func (m *DriverUpdateCatalogEntry) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDriverClass gets the driverClass property value. The classification of the driver.
func (m *DriverUpdateCatalogEntry) GetDriverClass()(*string) {
    val, err := m.GetBackingStore().Get("driverClass")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DriverUpdateCatalogEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SoftwareUpdateCatalogEntry.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["driverClass"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriverClass(val)
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["provider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvider(val)
        }
        return nil
    }
    res["setupInformationFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetupInformationFile(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    res["versionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionDateTime(val)
        }
        return nil
    }
    return res
}
// GetManufacturer gets the manufacturer property value. The manufacturer of the driver.
func (m *DriverUpdateCatalogEntry) GetManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("manufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProvider gets the provider property value. The provider of the driver.
func (m *DriverUpdateCatalogEntry) GetProvider()(*string) {
    val, err := m.GetBackingStore().Get("provider")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSetupInformationFile gets the setupInformationFile property value. The setup information file of the driver.
func (m *DriverUpdateCatalogEntry) GetSetupInformationFile()(*string) {
    val, err := m.GetBackingStore().Get("setupInformationFile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVersion gets the version property value. The unique version of the content.
func (m *DriverUpdateCatalogEntry) GetVersion()(*string) {
    val, err := m.GetBackingStore().Get("version")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVersionDateTime gets the versionDateTime property value. The date and time when a new version of content was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *DriverUpdateCatalogEntry) GetVersionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("versionDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DriverUpdateCatalogEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SoftwareUpdateCatalogEntry.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("driverClass", m.GetDriverClass())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("provider", m.GetProvider())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("setupInformationFile", m.GetSetupInformationFile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("versionDateTime", m.GetVersionDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description of the content.
func (m *DriverUpdateCatalogEntry) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDriverClass sets the driverClass property value. The classification of the driver.
func (m *DriverUpdateCatalogEntry) SetDriverClass(value *string)() {
    err := m.GetBackingStore().Set("driverClass", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. The manufacturer of the driver.
func (m *DriverUpdateCatalogEntry) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetProvider sets the provider property value. The provider of the driver.
func (m *DriverUpdateCatalogEntry) SetProvider(value *string)() {
    err := m.GetBackingStore().Set("provider", value)
    if err != nil {
        panic(err)
    }
}
// SetSetupInformationFile sets the setupInformationFile property value. The setup information file of the driver.
func (m *DriverUpdateCatalogEntry) SetSetupInformationFile(value *string)() {
    err := m.GetBackingStore().Set("setupInformationFile", value)
    if err != nil {
        panic(err)
    }
}
// SetVersion sets the version property value. The unique version of the content.
func (m *DriverUpdateCatalogEntry) SetVersion(value *string)() {
    err := m.GetBackingStore().Set("version", value)
    if err != nil {
        panic(err)
    }
}
// SetVersionDateTime sets the versionDateTime property value. The date and time when a new version of content was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *DriverUpdateCatalogEntry) SetVersionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("versionDateTime", value)
    if err != nil {
        panic(err)
    }
}
// DriverUpdateCatalogEntryable 
type DriverUpdateCatalogEntryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SoftwareUpdateCatalogEntryable
    GetDescription()(*string)
    GetDriverClass()(*string)
    GetManufacturer()(*string)
    GetProvider()(*string)
    GetSetupInformationFile()(*string)
    GetVersion()(*string)
    GetVersionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetDescription(value *string)()
    SetDriverClass(value *string)()
    SetManufacturer(value *string)()
    SetProvider(value *string)()
    SetSetupInformationFile(value *string)()
    SetVersion(value *string)()
    SetVersionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
