package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfficeSuiteAppable 
type OfficeSuiteAppable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    MobileAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutoAcceptEula()(*bool)
    GetExcludedApps()(ExcludedAppsable)
    GetInstallProgressDisplayLevel()(*OfficeSuiteInstallProgressDisplayLevel)
    GetLocalesToInstall()([]string)
    GetOfficeConfigurationXml()([]byte)
    GetOfficePlatformArchitecture()(*WindowsArchitecture)
    GetProductIds()([]string)
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
    SetProductIds(value []string)()
    SetShouldUninstallOlderVersionsOfOffice(value *bool)()
    SetTargetVersion(value *string)()
    SetUpdateChannel(value *OfficeUpdateChannel)()
    SetUpdateVersion(value *string)()
    SetUseSharedComputerActivation(value *bool)()
}
