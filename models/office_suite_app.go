package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfficeSuiteApp 
type OfficeSuiteApp struct {
    MobileApp
}
// NewOfficeSuiteApp instantiates a new OfficeSuiteApp and sets the default values.
func NewOfficeSuiteApp()(*OfficeSuiteApp) {
    m := &OfficeSuiteApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.officeSuiteApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOfficeSuiteAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOfficeSuiteAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOfficeSuiteApp(), nil
}
// GetAutoAcceptEula gets the autoAcceptEula property value. The value to accept the EULA automatically on the enduser's device.
func (m *OfficeSuiteApp) GetAutoAcceptEula()(*bool) {
    val, err := m.GetBackingStore().Get("autoAcceptEula")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetExcludedApps gets the excludedApps property value. The property to represent the apps which are excluded from the selected Office365 Product Id.
func (m *OfficeSuiteApp) GetExcludedApps()(ExcludedAppsable) {
    val, err := m.GetBackingStore().Get("excludedApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ExcludedAppsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeSuiteApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["autoAcceptEula"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoAcceptEula(val)
        }
        return nil
    }
    res["excludedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExcludedAppsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludedApps(val.(ExcludedAppsable))
        }
        return nil
    }
    res["installProgressDisplayLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOfficeSuiteInstallProgressDisplayLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallProgressDisplayLevel(val.(*OfficeSuiteInstallProgressDisplayLevel))
        }
        return nil
    }
    res["localesToInstall"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetLocalesToInstall(res)
        }
        return nil
    }
    res["officeConfigurationXml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeConfigurationXml(val)
        }
        return nil
    }
    res["officePlatformArchitecture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsArchitecture)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficePlatformArchitecture(val.(*WindowsArchitecture))
        }
        return nil
    }
    res["officeSuiteAppDefaultFileFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOfficeSuiteDefaultFileFormatType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeSuiteAppDefaultFileFormat(val.(*OfficeSuiteDefaultFileFormatType))
        }
        return nil
    }
    res["productIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseOfficeProductId)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OfficeProductId, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*OfficeProductId))
                }
            }
            m.SetProductIds(res)
        }
        return nil
    }
    res["shouldUninstallOlderVersionsOfOffice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShouldUninstallOlderVersionsOfOffice(val)
        }
        return nil
    }
    res["targetVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetVersion(val)
        }
        return nil
    }
    res["updateChannel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOfficeUpdateChannel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateChannel(val.(*OfficeUpdateChannel))
        }
        return nil
    }
    res["updateVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateVersion(val)
        }
        return nil
    }
    res["useSharedComputerActivation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseSharedComputerActivation(val)
        }
        return nil
    }
    return res
}
// GetInstallProgressDisplayLevel gets the installProgressDisplayLevel property value. The Enum to specify the level of display for the Installation Progress Setup UI on the Device.
func (m *OfficeSuiteApp) GetInstallProgressDisplayLevel()(*OfficeSuiteInstallProgressDisplayLevel) {
    val, err := m.GetBackingStore().Get("installProgressDisplayLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OfficeSuiteInstallProgressDisplayLevel)
    }
    return nil
}
// GetLocalesToInstall gets the localesToInstall property value. The property to represent the locales which are installed when the apps from Office365 is installed. It uses standard RFC 6033. Ref: https://technet.microsoft.com/library/cc179219(v=office.16).aspx
func (m *OfficeSuiteApp) GetLocalesToInstall()([]string) {
    val, err := m.GetBackingStore().Get("localesToInstall")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetOfficeConfigurationXml gets the officeConfigurationXml property value. The property to represent the XML configuration file that can be specified for Office ProPlus Apps. Takes precedence over all other properties. When present, the XML configuration file will be used to create the app.
func (m *OfficeSuiteApp) GetOfficeConfigurationXml()([]byte) {
    val, err := m.GetBackingStore().Get("officeConfigurationXml")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetOfficePlatformArchitecture gets the officePlatformArchitecture property value. Contains properties for Windows architecture.
func (m *OfficeSuiteApp) GetOfficePlatformArchitecture()(*WindowsArchitecture) {
    val, err := m.GetBackingStore().Get("officePlatformArchitecture")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsArchitecture)
    }
    return nil
}
// GetOfficeSuiteAppDefaultFileFormat gets the officeSuiteAppDefaultFileFormat property value. Describes the OfficeSuiteApp file format types that can be selected.
func (m *OfficeSuiteApp) GetOfficeSuiteAppDefaultFileFormat()(*OfficeSuiteDefaultFileFormatType) {
    val, err := m.GetBackingStore().Get("officeSuiteAppDefaultFileFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OfficeSuiteDefaultFileFormatType)
    }
    return nil
}
// GetProductIds gets the productIds property value. The Product Ids that represent the Office365 Suite SKU.
func (m *OfficeSuiteApp) GetProductIds()([]OfficeProductId) {
    val, err := m.GetBackingStore().Get("productIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OfficeProductId)
    }
    return nil
}
// GetShouldUninstallOlderVersionsOfOffice gets the shouldUninstallOlderVersionsOfOffice property value. The property to determine whether to uninstall existing Office MSI if an Office365 app suite is deployed to the device or not.
func (m *OfficeSuiteApp) GetShouldUninstallOlderVersionsOfOffice()(*bool) {
    val, err := m.GetBackingStore().Get("shouldUninstallOlderVersionsOfOffice")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTargetVersion gets the targetVersion property value. The property to represent the specific target version for the Office365 app suite that should be remained deployed on the devices.
func (m *OfficeSuiteApp) GetTargetVersion()(*string) {
    val, err := m.GetBackingStore().Get("targetVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUpdateChannel gets the updateChannel property value. The Enum to specify the Office365 Updates Channel.
func (m *OfficeSuiteApp) GetUpdateChannel()(*OfficeUpdateChannel) {
    val, err := m.GetBackingStore().Get("updateChannel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OfficeUpdateChannel)
    }
    return nil
}
// GetUpdateVersion gets the updateVersion property value. The property to represent the update version in which the specific target version is available for the Office365 app suite.
func (m *OfficeSuiteApp) GetUpdateVersion()(*string) {
    val, err := m.GetBackingStore().Get("updateVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUseSharedComputerActivation gets the useSharedComputerActivation property value. The property to represent that whether the shared computer activation is used not for Office365 app suite.
func (m *OfficeSuiteApp) GetUseSharedComputerActivation()(*bool) {
    val, err := m.GetBackingStore().Get("useSharedComputerActivation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OfficeSuiteApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("autoAcceptEula", m.GetAutoAcceptEula())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("excludedApps", m.GetExcludedApps())
        if err != nil {
            return err
        }
    }
    if m.GetInstallProgressDisplayLevel() != nil {
        cast := (*m.GetInstallProgressDisplayLevel()).String()
        err = writer.WriteStringValue("installProgressDisplayLevel", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLocalesToInstall() != nil {
        err = writer.WriteCollectionOfStringValues("localesToInstall", m.GetLocalesToInstall())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("officeConfigurationXml", m.GetOfficeConfigurationXml())
        if err != nil {
            return err
        }
    }
    if m.GetOfficePlatformArchitecture() != nil {
        cast := (*m.GetOfficePlatformArchitecture()).String()
        err = writer.WriteStringValue("officePlatformArchitecture", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOfficeSuiteAppDefaultFileFormat() != nil {
        cast := (*m.GetOfficeSuiteAppDefaultFileFormat()).String()
        err = writer.WriteStringValue("officeSuiteAppDefaultFileFormat", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetProductIds() != nil {
        err = writer.WriteCollectionOfStringValues("productIds", SerializeOfficeProductId(m.GetProductIds()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("shouldUninstallOlderVersionsOfOffice", m.GetShouldUninstallOlderVersionsOfOffice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetVersion", m.GetTargetVersion())
        if err != nil {
            return err
        }
    }
    if m.GetUpdateChannel() != nil {
        cast := (*m.GetUpdateChannel()).String()
        err = writer.WriteStringValue("updateChannel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("updateVersion", m.GetUpdateVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useSharedComputerActivation", m.GetUseSharedComputerActivation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAutoAcceptEula sets the autoAcceptEula property value. The value to accept the EULA automatically on the enduser's device.
func (m *OfficeSuiteApp) SetAutoAcceptEula(value *bool)() {
    err := m.GetBackingStore().Set("autoAcceptEula", value)
    if err != nil {
        panic(err)
    }
}
// SetExcludedApps sets the excludedApps property value. The property to represent the apps which are excluded from the selected Office365 Product Id.
func (m *OfficeSuiteApp) SetExcludedApps(value ExcludedAppsable)() {
    err := m.GetBackingStore().Set("excludedApps", value)
    if err != nil {
        panic(err)
    }
}
// SetInstallProgressDisplayLevel sets the installProgressDisplayLevel property value. The Enum to specify the level of display for the Installation Progress Setup UI on the Device.
func (m *OfficeSuiteApp) SetInstallProgressDisplayLevel(value *OfficeSuiteInstallProgressDisplayLevel)() {
    err := m.GetBackingStore().Set("installProgressDisplayLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalesToInstall sets the localesToInstall property value. The property to represent the locales which are installed when the apps from Office365 is installed. It uses standard RFC 6033. Ref: https://technet.microsoft.com/library/cc179219(v=office.16).aspx
func (m *OfficeSuiteApp) SetLocalesToInstall(value []string)() {
    err := m.GetBackingStore().Set("localesToInstall", value)
    if err != nil {
        panic(err)
    }
}
// SetOfficeConfigurationXml sets the officeConfigurationXml property value. The property to represent the XML configuration file that can be specified for Office ProPlus Apps. Takes precedence over all other properties. When present, the XML configuration file will be used to create the app.
func (m *OfficeSuiteApp) SetOfficeConfigurationXml(value []byte)() {
    err := m.GetBackingStore().Set("officeConfigurationXml", value)
    if err != nil {
        panic(err)
    }
}
// SetOfficePlatformArchitecture sets the officePlatformArchitecture property value. Contains properties for Windows architecture.
func (m *OfficeSuiteApp) SetOfficePlatformArchitecture(value *WindowsArchitecture)() {
    err := m.GetBackingStore().Set("officePlatformArchitecture", value)
    if err != nil {
        panic(err)
    }
}
// SetOfficeSuiteAppDefaultFileFormat sets the officeSuiteAppDefaultFileFormat property value. Describes the OfficeSuiteApp file format types that can be selected.
func (m *OfficeSuiteApp) SetOfficeSuiteAppDefaultFileFormat(value *OfficeSuiteDefaultFileFormatType)() {
    err := m.GetBackingStore().Set("officeSuiteAppDefaultFileFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetProductIds sets the productIds property value. The Product Ids that represent the Office365 Suite SKU.
func (m *OfficeSuiteApp) SetProductIds(value []OfficeProductId)() {
    err := m.GetBackingStore().Set("productIds", value)
    if err != nil {
        panic(err)
    }
}
// SetShouldUninstallOlderVersionsOfOffice sets the shouldUninstallOlderVersionsOfOffice property value. The property to determine whether to uninstall existing Office MSI if an Office365 app suite is deployed to the device or not.
func (m *OfficeSuiteApp) SetShouldUninstallOlderVersionsOfOffice(value *bool)() {
    err := m.GetBackingStore().Set("shouldUninstallOlderVersionsOfOffice", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetVersion sets the targetVersion property value. The property to represent the specific target version for the Office365 app suite that should be remained deployed on the devices.
func (m *OfficeSuiteApp) SetTargetVersion(value *string)() {
    err := m.GetBackingStore().Set("targetVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdateChannel sets the updateChannel property value. The Enum to specify the Office365 Updates Channel.
func (m *OfficeSuiteApp) SetUpdateChannel(value *OfficeUpdateChannel)() {
    err := m.GetBackingStore().Set("updateChannel", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdateVersion sets the updateVersion property value. The property to represent the update version in which the specific target version is available for the Office365 app suite.
func (m *OfficeSuiteApp) SetUpdateVersion(value *string)() {
    err := m.GetBackingStore().Set("updateVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetUseSharedComputerActivation sets the useSharedComputerActivation property value. The property to represent that whether the shared computer activation is used not for Office365 app suite.
func (m *OfficeSuiteApp) SetUseSharedComputerActivation(value *bool)() {
    err := m.GetBackingStore().Set("useSharedComputerActivation", value)
    if err != nil {
        panic(err)
    }
}
// OfficeSuiteAppable 
type OfficeSuiteAppable interface {
    MobileAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutoAcceptEula()(*bool)
    GetExcludedApps()(ExcludedAppsable)
    GetInstallProgressDisplayLevel()(*OfficeSuiteInstallProgressDisplayLevel)
    GetLocalesToInstall()([]string)
    GetOfficeConfigurationXml()([]byte)
    GetOfficePlatformArchitecture()(*WindowsArchitecture)
    GetOfficeSuiteAppDefaultFileFormat()(*OfficeSuiteDefaultFileFormatType)
    GetProductIds()([]OfficeProductId)
    GetShouldUninstallOlderVersionsOfOffice()(*bool)
    GetTargetVersion()(*string)
    GetUpdateChannel()(*OfficeUpdateChannel)
    GetUpdateVersion()(*string)
    GetUseSharedComputerActivation()(*bool)
    SetAutoAcceptEula(value *bool)()
    SetExcludedApps(value ExcludedAppsable)()
    SetInstallProgressDisplayLevel(value *OfficeSuiteInstallProgressDisplayLevel)()
    SetLocalesToInstall(value []string)()
    SetOfficeConfigurationXml(value []byte)()
    SetOfficePlatformArchitecture(value *WindowsArchitecture)()
    SetOfficeSuiteAppDefaultFileFormat(value *OfficeSuiteDefaultFileFormatType)()
    SetProductIds(value []OfficeProductId)()
    SetShouldUninstallOlderVersionsOfOffice(value *bool)()
    SetTargetVersion(value *string)()
    SetUpdateChannel(value *OfficeUpdateChannel)()
    SetUpdateVersion(value *string)()
    SetUseSharedComputerActivation(value *bool)()
}
