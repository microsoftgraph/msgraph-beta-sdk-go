package windowsupdates

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DriverUpdateCatalogEntry 
type DriverUpdateCatalogEntry struct {
    SoftwareUpdateCatalogEntry
    // The description property
    description *string
    // The driverClass property
    driverClass *string
    // The manufacturer property
    manufacturer *string
    // The provider property
    provider *string
    // The setupInformationFile property
    setupInformationFile *string
    // The version property
    version *string
    // The versionDateTime property
    versionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewDriverUpdateCatalogEntry instantiates a new DriverUpdateCatalogEntry and sets the default values.
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
// GetDescription gets the description property value. The description property
func (m *DriverUpdateCatalogEntry) GetDescription()(*string) {
    return m.description
}
// GetDriverClass gets the driverClass property value. The driverClass property
func (m *DriverUpdateCatalogEntry) GetDriverClass()(*string) {
    return m.driverClass
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
// GetManufacturer gets the manufacturer property value. The manufacturer property
func (m *DriverUpdateCatalogEntry) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetProvider gets the provider property value. The provider property
func (m *DriverUpdateCatalogEntry) GetProvider()(*string) {
    return m.provider
}
// GetSetupInformationFile gets the setupInformationFile property value. The setupInformationFile property
func (m *DriverUpdateCatalogEntry) GetSetupInformationFile()(*string) {
    return m.setupInformationFile
}
// GetVersion gets the version property value. The version property
func (m *DriverUpdateCatalogEntry) GetVersion()(*string) {
    return m.version
}
// GetVersionDateTime gets the versionDateTime property value. The versionDateTime property
func (m *DriverUpdateCatalogEntry) GetVersionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.versionDateTime
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
// SetDescription sets the description property value. The description property
func (m *DriverUpdateCatalogEntry) SetDescription(value *string)() {
    m.description = value
}
// SetDriverClass sets the driverClass property value. The driverClass property
func (m *DriverUpdateCatalogEntry) SetDriverClass(value *string)() {
    m.driverClass = value
}
// SetManufacturer sets the manufacturer property value. The manufacturer property
func (m *DriverUpdateCatalogEntry) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetProvider sets the provider property value. The provider property
func (m *DriverUpdateCatalogEntry) SetProvider(value *string)() {
    m.provider = value
}
// SetSetupInformationFile sets the setupInformationFile property value. The setupInformationFile property
func (m *DriverUpdateCatalogEntry) SetSetupInformationFile(value *string)() {
    m.setupInformationFile = value
}
// SetVersion sets the version property value. The version property
func (m *DriverUpdateCatalogEntry) SetVersion(value *string)() {
    m.version = value
}
// SetVersionDateTime sets the versionDateTime property value. The versionDateTime property
func (m *DriverUpdateCatalogEntry) SetVersionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.versionDateTime = value
}
