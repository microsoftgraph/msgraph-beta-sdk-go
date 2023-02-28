package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobApp 
type Win32LobApp struct {
    MobileLobApp
}
// NewWin32LobApp instantiates a new Win32LobApp and sets the default values.
func NewWin32LobApp()(*Win32LobApp) {
    m := &Win32LobApp{
        MobileLobApp: *NewMobileLobApp(),
    }
    odataTypeValue := "#microsoft.graph.win32LobApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWin32LobAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobApp(), nil
}
// GetAllowAvailableUninstall gets the allowAvailableUninstall property value. When TRUE, indicates that uninstall is supported from the company portal for the Windows app (Win32) with an Available assignment. When FALSE, indicates that uninstall is not supported for the Windows app (Win32) with an Available assignment. Default value is FALSE.
func (m *Win32LobApp) GetAllowAvailableUninstall()(*bool) {
    val, err := m.GetBackingStore().Get("allowAvailableUninstall")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetApplicableArchitectures gets the applicableArchitectures property value. Contains properties for Windows architecture.
func (m *Win32LobApp) GetApplicableArchitectures()(*WindowsArchitecture) {
    val, err := m.GetBackingStore().Get("applicableArchitectures")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsArchitecture)
    }
    return nil
}
// GetDetectionRules gets the detectionRules property value. The detection rules to detect Win32 Line of Business (LoB) app.
func (m *Win32LobApp) GetDetectionRules()([]Win32LobAppDetectionable) {
    val, err := m.GetBackingStore().Get("detectionRules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Win32LobAppDetectionable)
    }
    return nil
}
// GetDisplayVersion gets the displayVersion property value. The version displayed in the UX for this app.
func (m *Win32LobApp) GetDisplayVersion()(*string) {
    val, err := m.GetBackingStore().Get("displayVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
    res["allowAvailableUninstall"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAvailableUninstall(val)
        }
        return nil
    }
    res["applicableArchitectures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsArchitecture)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableArchitectures(val.(*WindowsArchitecture))
        }
        return nil
    }
    res["detectionRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWin32LobAppDetectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Win32LobAppDetectionable, len(val))
            for i, v := range val {
                res[i] = v.(Win32LobAppDetectionable)
            }
            m.SetDetectionRules(res)
        }
        return nil
    }
    res["displayVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayVersion(val)
        }
        return nil
    }
    res["installCommandLine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallCommandLine(val)
        }
        return nil
    }
    res["installExperience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWin32LobAppInstallExperienceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallExperience(val.(Win32LobAppInstallExperienceable))
        }
        return nil
    }
    res["minimumCpuSpeedInMHz"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumCpuSpeedInMHz(val)
        }
        return nil
    }
    res["minimumFreeDiskSpaceInMB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumFreeDiskSpaceInMB(val)
        }
        return nil
    }
    res["minimumMemoryInMB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumMemoryInMB(val)
        }
        return nil
    }
    res["minimumNumberOfProcessors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumNumberOfProcessors(val)
        }
        return nil
    }
    res["minimumSupportedOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsMinimumOperatingSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedOperatingSystem(val.(WindowsMinimumOperatingSystemable))
        }
        return nil
    }
    res["minimumSupportedWindowsRelease"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedWindowsRelease(val)
        }
        return nil
    }
    res["msiInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWin32LobAppMsiInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMsiInformation(val.(Win32LobAppMsiInformationable))
        }
        return nil
    }
    res["requirementRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWin32LobAppRequirementFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Win32LobAppRequirementable, len(val))
            for i, v := range val {
                res[i] = v.(Win32LobAppRequirementable)
            }
            m.SetRequirementRules(res)
        }
        return nil
    }
    res["returnCodes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWin32LobAppReturnCodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Win32LobAppReturnCodeable, len(val))
            for i, v := range val {
                res[i] = v.(Win32LobAppReturnCodeable)
            }
            m.SetReturnCodes(res)
        }
        return nil
    }
    res["rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWin32LobAppRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Win32LobAppRuleable, len(val))
            for i, v := range val {
                res[i] = v.(Win32LobAppRuleable)
            }
            m.SetRules(res)
        }
        return nil
    }
    res["setupFilePath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetupFilePath(val)
        }
        return nil
    }
    res["uninstallCommandLine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUninstallCommandLine(val)
        }
        return nil
    }
    return res
}
// GetInstallCommandLine gets the installCommandLine property value. The command line to install this app
func (m *Win32LobApp) GetInstallCommandLine()(*string) {
    val, err := m.GetBackingStore().Get("installCommandLine")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetInstallExperience gets the installExperience property value. The install experience for this app.
func (m *Win32LobApp) GetInstallExperience()(Win32LobAppInstallExperienceable) {
    val, err := m.GetBackingStore().Get("installExperience")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Win32LobAppInstallExperienceable)
    }
    return nil
}
// GetMinimumCpuSpeedInMHz gets the minimumCpuSpeedInMHz property value. The value for the minimum CPU speed which is required to install this app.
func (m *Win32LobApp) GetMinimumCpuSpeedInMHz()(*int32) {
    val, err := m.GetBackingStore().Get("minimumCpuSpeedInMHz")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMinimumFreeDiskSpaceInMB gets the minimumFreeDiskSpaceInMB property value. The value for the minimum free disk space which is required to install this app.
func (m *Win32LobApp) GetMinimumFreeDiskSpaceInMB()(*int32) {
    val, err := m.GetBackingStore().Get("minimumFreeDiskSpaceInMB")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMinimumMemoryInMB gets the minimumMemoryInMB property value. The value for the minimum physical memory which is required to install this app.
func (m *Win32LobApp) GetMinimumMemoryInMB()(*int32) {
    val, err := m.GetBackingStore().Get("minimumMemoryInMB")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMinimumNumberOfProcessors gets the minimumNumberOfProcessors property value. The value for the minimum number of processors which is required to install this app.
func (m *Win32LobApp) GetMinimumNumberOfProcessors()(*int32) {
    val, err := m.GetBackingStore().Get("minimumNumberOfProcessors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *Win32LobApp) GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable) {
    val, err := m.GetBackingStore().Get("minimumSupportedOperatingSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsMinimumOperatingSystemable)
    }
    return nil
}
// GetMinimumSupportedWindowsRelease gets the minimumSupportedWindowsRelease property value. The value for the minimum supported windows release.
func (m *Win32LobApp) GetMinimumSupportedWindowsRelease()(*string) {
    val, err := m.GetBackingStore().Get("minimumSupportedWindowsRelease")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMsiInformation gets the msiInformation property value. The MSI details if this Win32 app is an MSI app.
func (m *Win32LobApp) GetMsiInformation()(Win32LobAppMsiInformationable) {
    val, err := m.GetBackingStore().Get("msiInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Win32LobAppMsiInformationable)
    }
    return nil
}
// GetRequirementRules gets the requirementRules property value. The requirement rules to detect Win32 Line of Business (LoB) app.
func (m *Win32LobApp) GetRequirementRules()([]Win32LobAppRequirementable) {
    val, err := m.GetBackingStore().Get("requirementRules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Win32LobAppRequirementable)
    }
    return nil
}
// GetReturnCodes gets the returnCodes property value. The return codes for post installation behavior.
func (m *Win32LobApp) GetReturnCodes()([]Win32LobAppReturnCodeable) {
    val, err := m.GetBackingStore().Get("returnCodes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Win32LobAppReturnCodeable)
    }
    return nil
}
// GetRules gets the rules property value. The detection and requirement rules for this app.
func (m *Win32LobApp) GetRules()([]Win32LobAppRuleable) {
    val, err := m.GetBackingStore().Get("rules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Win32LobAppRuleable)
    }
    return nil
}
// GetSetupFilePath gets the setupFilePath property value. The relative path of the setup file in the encrypted Win32LobApp package.
func (m *Win32LobApp) GetSetupFilePath()(*string) {
    val, err := m.GetBackingStore().Get("setupFilePath")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUninstallCommandLine gets the uninstallCommandLine property value. The command line to uninstall this app
func (m *Win32LobApp) GetUninstallCommandLine()(*string) {
    val, err := m.GetBackingStore().Get("uninstallCommandLine")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Win32LobApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileLobApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowAvailableUninstall", m.GetAllowAvailableUninstall())
        if err != nil {
            return err
        }
    }
    if m.GetApplicableArchitectures() != nil {
        cast := (*m.GetApplicableArchitectures()).String()
        err = writer.WriteStringValue("applicableArchitectures", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDetectionRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDetectionRules()))
        for i, v := range m.GetDetectionRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("detectionRules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayVersion", m.GetDisplayVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("installCommandLine", m.GetInstallCommandLine())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("installExperience", m.GetInstallExperience())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumCpuSpeedInMHz", m.GetMinimumCpuSpeedInMHz())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumFreeDiskSpaceInMB", m.GetMinimumFreeDiskSpaceInMB())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumMemoryInMB", m.GetMinimumMemoryInMB())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumNumberOfProcessors", m.GetMinimumNumberOfProcessors())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minimumSupportedOperatingSystem", m.GetMinimumSupportedOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumSupportedWindowsRelease", m.GetMinimumSupportedWindowsRelease())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("msiInformation", m.GetMsiInformation())
        if err != nil {
            return err
        }
    }
    if m.GetRequirementRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRequirementRules()))
        for i, v := range m.GetRequirementRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("requirementRules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetReturnCodes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReturnCodes()))
        for i, v := range m.GetReturnCodes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("returnCodes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("rules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("setupFilePath", m.GetSetupFilePath())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uninstallCommandLine", m.GetUninstallCommandLine())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowAvailableUninstall sets the allowAvailableUninstall property value. When TRUE, indicates that uninstall is supported from the company portal for the Windows app (Win32) with an Available assignment. When FALSE, indicates that uninstall is not supported for the Windows app (Win32) with an Available assignment. Default value is FALSE.
func (m *Win32LobApp) SetAllowAvailableUninstall(value *bool)() {
    err := m.GetBackingStore().Set("allowAvailableUninstall", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicableArchitectures sets the applicableArchitectures property value. Contains properties for Windows architecture.
func (m *Win32LobApp) SetApplicableArchitectures(value *WindowsArchitecture)() {
    err := m.GetBackingStore().Set("applicableArchitectures", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectionRules sets the detectionRules property value. The detection rules to detect Win32 Line of Business (LoB) app.
func (m *Win32LobApp) SetDetectionRules(value []Win32LobAppDetectionable)() {
    err := m.GetBackingStore().Set("detectionRules", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayVersion sets the displayVersion property value. The version displayed in the UX for this app.
func (m *Win32LobApp) SetDisplayVersion(value *string)() {
    err := m.GetBackingStore().Set("displayVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetInstallCommandLine sets the installCommandLine property value. The command line to install this app
func (m *Win32LobApp) SetInstallCommandLine(value *string)() {
    err := m.GetBackingStore().Set("installCommandLine", value)
    if err != nil {
        panic(err)
    }
}
// SetInstallExperience sets the installExperience property value. The install experience for this app.
func (m *Win32LobApp) SetInstallExperience(value Win32LobAppInstallExperienceable)() {
    err := m.GetBackingStore().Set("installExperience", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumCpuSpeedInMHz sets the minimumCpuSpeedInMHz property value. The value for the minimum CPU speed which is required to install this app.
func (m *Win32LobApp) SetMinimumCpuSpeedInMHz(value *int32)() {
    err := m.GetBackingStore().Set("minimumCpuSpeedInMHz", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumFreeDiskSpaceInMB sets the minimumFreeDiskSpaceInMB property value. The value for the minimum free disk space which is required to install this app.
func (m *Win32LobApp) SetMinimumFreeDiskSpaceInMB(value *int32)() {
    err := m.GetBackingStore().Set("minimumFreeDiskSpaceInMB", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumMemoryInMB sets the minimumMemoryInMB property value. The value for the minimum physical memory which is required to install this app.
func (m *Win32LobApp) SetMinimumMemoryInMB(value *int32)() {
    err := m.GetBackingStore().Set("minimumMemoryInMB", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumNumberOfProcessors sets the minimumNumberOfProcessors property value. The value for the minimum number of processors which is required to install this app.
func (m *Win32LobApp) SetMinimumNumberOfProcessors(value *int32)() {
    err := m.GetBackingStore().Set("minimumNumberOfProcessors", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *Win32LobApp) SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)() {
    err := m.GetBackingStore().Set("minimumSupportedOperatingSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumSupportedWindowsRelease sets the minimumSupportedWindowsRelease property value. The value for the minimum supported windows release.
func (m *Win32LobApp) SetMinimumSupportedWindowsRelease(value *string)() {
    err := m.GetBackingStore().Set("minimumSupportedWindowsRelease", value)
    if err != nil {
        panic(err)
    }
}
// SetMsiInformation sets the msiInformation property value. The MSI details if this Win32 app is an MSI app.
func (m *Win32LobApp) SetMsiInformation(value Win32LobAppMsiInformationable)() {
    err := m.GetBackingStore().Set("msiInformation", value)
    if err != nil {
        panic(err)
    }
}
// SetRequirementRules sets the requirementRules property value. The requirement rules to detect Win32 Line of Business (LoB) app.
func (m *Win32LobApp) SetRequirementRules(value []Win32LobAppRequirementable)() {
    err := m.GetBackingStore().Set("requirementRules", value)
    if err != nil {
        panic(err)
    }
}
// SetReturnCodes sets the returnCodes property value. The return codes for post installation behavior.
func (m *Win32LobApp) SetReturnCodes(value []Win32LobAppReturnCodeable)() {
    err := m.GetBackingStore().Set("returnCodes", value)
    if err != nil {
        panic(err)
    }
}
// SetRules sets the rules property value. The detection and requirement rules for this app.
func (m *Win32LobApp) SetRules(value []Win32LobAppRuleable)() {
    err := m.GetBackingStore().Set("rules", value)
    if err != nil {
        panic(err)
    }
}
// SetSetupFilePath sets the setupFilePath property value. The relative path of the setup file in the encrypted Win32LobApp package.
func (m *Win32LobApp) SetSetupFilePath(value *string)() {
    err := m.GetBackingStore().Set("setupFilePath", value)
    if err != nil {
        panic(err)
    }
}
// SetUninstallCommandLine sets the uninstallCommandLine property value. The command line to uninstall this app
func (m *Win32LobApp) SetUninstallCommandLine(value *string)() {
    err := m.GetBackingStore().Set("uninstallCommandLine", value)
    if err != nil {
        panic(err)
    }
}
// Win32LobAppable 
type Win32LobAppable interface {
    MobileLobAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowAvailableUninstall()(*bool)
    GetApplicableArchitectures()(*WindowsArchitecture)
    GetDetectionRules()([]Win32LobAppDetectionable)
    GetDisplayVersion()(*string)
    GetInstallCommandLine()(*string)
    GetInstallExperience()(Win32LobAppInstallExperienceable)
    GetMinimumCpuSpeedInMHz()(*int32)
    GetMinimumFreeDiskSpaceInMB()(*int32)
    GetMinimumMemoryInMB()(*int32)
    GetMinimumNumberOfProcessors()(*int32)
    GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable)
    GetMinimumSupportedWindowsRelease()(*string)
    GetMsiInformation()(Win32LobAppMsiInformationable)
    GetRequirementRules()([]Win32LobAppRequirementable)
    GetReturnCodes()([]Win32LobAppReturnCodeable)
    GetRules()([]Win32LobAppRuleable)
    GetSetupFilePath()(*string)
    GetUninstallCommandLine()(*string)
    SetAllowAvailableUninstall(value *bool)()
    SetApplicableArchitectures(value *WindowsArchitecture)()
    SetDetectionRules(value []Win32LobAppDetectionable)()
    SetDisplayVersion(value *string)()
    SetInstallCommandLine(value *string)()
    SetInstallExperience(value Win32LobAppInstallExperienceable)()
    SetMinimumCpuSpeedInMHz(value *int32)()
    SetMinimumFreeDiskSpaceInMB(value *int32)()
    SetMinimumMemoryInMB(value *int32)()
    SetMinimumNumberOfProcessors(value *int32)()
    SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)()
    SetMinimumSupportedWindowsRelease(value *string)()
    SetMsiInformation(value Win32LobAppMsiInformationable)()
    SetRequirementRules(value []Win32LobAppRequirementable)()
    SetReturnCodes(value []Win32LobAppReturnCodeable)()
    SetRules(value []Win32LobAppRuleable)()
    SetSetupFilePath(value *string)()
    SetUninstallCommandLine(value *string)()
}
