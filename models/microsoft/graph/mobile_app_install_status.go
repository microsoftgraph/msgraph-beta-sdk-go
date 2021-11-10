package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MobileAppInstallStatus struct {
    Entity
    // The navigation link to the mobile app.
    app *MobileApp;
    // Device ID
    deviceId *string;
    // Device name
    deviceName *string;
    // Human readable version of the application
    displayVersion *string;
    // The error code for install or uninstall failures.
    errorCode *int32;
    // The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
    installState *ResultantAppState;
    // The install state detail of the app. Possible values are: noAdditionalDetails, dependencyFailedToInstall, dependencyWithRequirementsNotMet, dependencyPendingReboot, dependencyWithAutoInstallDisabled, supersededAppUninstallFailed, supersededAppUninstallPendingReboot, removingSupersededApps, iosAppStoreUpdateFailedToInstall, vppAppHasUpdateAvailable, userRejectedUpdate, uninstallPendingReboot, supersedingAppsDetected, supersededAppsDetected, seeInstallErrorCode, autoInstallDisabled, managedAppNoLongerPresent, userRejectedInstall, userIsNotLoggedIntoAppStore, untargetedSupersedingAppsDetected, appRemovedBySupersedence, seeUninstallErrorCode, pendingReboot, installingDependencies, contentDownloaded, supersedingAppsNotApplicable, powerShellScriptRequirementNotMet, registryRequirementNotMet, fileSystemRequirementNotMet, platformNotApplicable, minimumCpuSpeedNotMet, minimumLogicalProcessorCountNotMet, minimumPhysicalMemoryNotMet, minimumOsVersionNotMet, minimumDiskSpaceNotMet, processorArchitectureNotApplicable.
    installStateDetail *ResultantAppStateDetail;
    // Last sync date time
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
    mobileAppInstallStatusValue *ResultantAppState;
    // OS Description
    osDescription *string;
    // OS Version
    osVersion *string;
    // Device User Name
    userName *string;
    // User Principal Name
    userPrincipalName *string;
}
// Instantiates a new mobileAppInstallStatus and sets the default values.
func NewMobileAppInstallStatus()(*MobileAppInstallStatus) {
    m := &MobileAppInstallStatus{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the app property value. The navigation link to the mobile app.
func (m *MobileAppInstallStatus) GetApp()(*MobileApp) {
    if m == nil {
        return nil
    } else {
        return m.app
    }
}
// Gets the deviceId property value. Device ID
func (m *MobileAppInstallStatus) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the deviceName property value. Device name
func (m *MobileAppInstallStatus) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the displayVersion property value. Human readable version of the application
func (m *MobileAppInstallStatus) GetDisplayVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayVersion
    }
}
// Gets the errorCode property value. The error code for install or uninstall failures.
func (m *MobileAppInstallStatus) GetErrorCode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// Gets the installState property value. The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
func (m *MobileAppInstallStatus) GetInstallState()(*ResultantAppState) {
    if m == nil {
        return nil
    } else {
        return m.installState
    }
}
// Gets the installStateDetail property value. The install state detail of the app. Possible values are: noAdditionalDetails, dependencyFailedToInstall, dependencyWithRequirementsNotMet, dependencyPendingReboot, dependencyWithAutoInstallDisabled, supersededAppUninstallFailed, supersededAppUninstallPendingReboot, removingSupersededApps, iosAppStoreUpdateFailedToInstall, vppAppHasUpdateAvailable, userRejectedUpdate, uninstallPendingReboot, supersedingAppsDetected, supersededAppsDetected, seeInstallErrorCode, autoInstallDisabled, managedAppNoLongerPresent, userRejectedInstall, userIsNotLoggedIntoAppStore, untargetedSupersedingAppsDetected, appRemovedBySupersedence, seeUninstallErrorCode, pendingReboot, installingDependencies, contentDownloaded, supersedingAppsNotApplicable, powerShellScriptRequirementNotMet, registryRequirementNotMet, fileSystemRequirementNotMet, platformNotApplicable, minimumCpuSpeedNotMet, minimumLogicalProcessorCountNotMet, minimumPhysicalMemoryNotMet, minimumOsVersionNotMet, minimumDiskSpaceNotMet, processorArchitectureNotApplicable.
func (m *MobileAppInstallStatus) GetInstallStateDetail()(*ResultantAppStateDetail) {
    if m == nil {
        return nil
    } else {
        return m.installStateDetail
    }
}
// Gets the lastSyncDateTime property value. Last sync date time
func (m *MobileAppInstallStatus) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// Gets the mobileAppInstallStatusValue property value. The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
func (m *MobileAppInstallStatus) GetMobileAppInstallStatusValue()(*ResultantAppState) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppInstallStatusValue
    }
}
// Gets the osDescription property value. OS Description
func (m *MobileAppInstallStatus) GetOsDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osDescription
    }
}
// Gets the osVersion property value. OS Version
func (m *MobileAppInstallStatus) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// Gets the userName property value. Device User Name
func (m *MobileAppInstallStatus) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// Gets the userPrincipalName property value. User Principal Name
func (m *MobileAppInstallStatus) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *MobileAppInstallStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["app"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileApp() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApp(val.(*MobileApp))
        }
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["displayVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayVersion(val)
        }
        return nil
    }
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["installState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseResultantAppState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ResultantAppState)
            m.SetInstallState(&cast)
        }
        return nil
    }
    res["installStateDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseResultantAppStateDetail)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ResultantAppStateDetail)
            m.SetInstallStateDetail(&cast)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["mobileAppInstallStatusValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseResultantAppState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ResultantAppState)
            m.SetMobileAppInstallStatusValue(&cast)
        }
        return nil
    }
    res["osDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsDescription(val)
        }
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *MobileAppInstallStatus) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MobileAppInstallStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("app", m.GetApp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
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
        err = writer.WriteInt32Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    if m.GetInstallState() != nil {
        cast := m.GetInstallState().String()
        err = writer.WriteStringValue("installState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetInstallStateDetail() != nil {
        cast := m.GetInstallStateDetail().String()
        err = writer.WriteStringValue("installStateDetail", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppInstallStatusValue() != nil {
        cast := m.GetMobileAppInstallStatusValue().String()
        err = writer.WriteStringValue("mobileAppInstallStatusValue", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osDescription", m.GetOsDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the app property value. The navigation link to the mobile app.
// Parameters:
//  - value : Value to set for the app property.
func (m *MobileAppInstallStatus) SetApp(value *MobileApp)() {
    m.app = value
}
// Sets the deviceId property value. Device ID
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *MobileAppInstallStatus) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the deviceName property value. Device name
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *MobileAppInstallStatus) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the displayVersion property value. Human readable version of the application
// Parameters:
//  - value : Value to set for the displayVersion property.
func (m *MobileAppInstallStatus) SetDisplayVersion(value *string)() {
    m.displayVersion = value
}
// Sets the errorCode property value. The error code for install or uninstall failures.
// Parameters:
//  - value : Value to set for the errorCode property.
func (m *MobileAppInstallStatus) SetErrorCode(value *int32)() {
    m.errorCode = value
}
// Sets the installState property value. The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
// Parameters:
//  - value : Value to set for the installState property.
func (m *MobileAppInstallStatus) SetInstallState(value *ResultantAppState)() {
    m.installState = value
}
// Sets the installStateDetail property value. The install state detail of the app. Possible values are: noAdditionalDetails, dependencyFailedToInstall, dependencyWithRequirementsNotMet, dependencyPendingReboot, dependencyWithAutoInstallDisabled, supersededAppUninstallFailed, supersededAppUninstallPendingReboot, removingSupersededApps, iosAppStoreUpdateFailedToInstall, vppAppHasUpdateAvailable, userRejectedUpdate, uninstallPendingReboot, supersedingAppsDetected, supersededAppsDetected, seeInstallErrorCode, autoInstallDisabled, managedAppNoLongerPresent, userRejectedInstall, userIsNotLoggedIntoAppStore, untargetedSupersedingAppsDetected, appRemovedBySupersedence, seeUninstallErrorCode, pendingReboot, installingDependencies, contentDownloaded, supersedingAppsNotApplicable, powerShellScriptRequirementNotMet, registryRequirementNotMet, fileSystemRequirementNotMet, platformNotApplicable, minimumCpuSpeedNotMet, minimumLogicalProcessorCountNotMet, minimumPhysicalMemoryNotMet, minimumOsVersionNotMet, minimumDiskSpaceNotMet, processorArchitectureNotApplicable.
// Parameters:
//  - value : Value to set for the installStateDetail property.
func (m *MobileAppInstallStatus) SetInstallStateDetail(value *ResultantAppStateDetail)() {
    m.installStateDetail = value
}
// Sets the lastSyncDateTime property value. Last sync date time
// Parameters:
//  - value : Value to set for the lastSyncDateTime property.
func (m *MobileAppInstallStatus) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// Sets the mobileAppInstallStatusValue property value. The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
// Parameters:
//  - value : Value to set for the mobileAppInstallStatusValue property.
func (m *MobileAppInstallStatus) SetMobileAppInstallStatusValue(value *ResultantAppState)() {
    m.mobileAppInstallStatusValue = value
}
// Sets the osDescription property value. OS Description
// Parameters:
//  - value : Value to set for the osDescription property.
func (m *MobileAppInstallStatus) SetOsDescription(value *string)() {
    m.osDescription = value
}
// Sets the osVersion property value. OS Version
// Parameters:
//  - value : Value to set for the osVersion property.
func (m *MobileAppInstallStatus) SetOsVersion(value *string)() {
    m.osVersion = value
}
// Sets the userName property value. Device User Name
// Parameters:
//  - value : Value to set for the userName property.
func (m *MobileAppInstallStatus) SetUserName(value *string)() {
    m.userName = value
}
// Sets the userPrincipalName property value. User Principal Name
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *MobileAppInstallStatus) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
