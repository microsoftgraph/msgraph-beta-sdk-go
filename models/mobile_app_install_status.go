package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppInstallStatus contains properties for the installation state of a mobile app for a device.
type MobileAppInstallStatus struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The navigation link to the mobile app.
    app MobileAppable
    // Device ID
    deviceId *string
    // Device name
    deviceName *string
    // Human readable version of the application
    displayVersion *string
    // The error code for install or uninstall failures.
    errorCode *int32
    // The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
    installState *ResultantAppState
    // The install state detail of the app. Possible values are: noAdditionalDetails, dependencyFailedToInstall, dependencyWithRequirementsNotMet, dependencyPendingReboot, dependencyWithAutoInstallDisabled, supersededAppUninstallFailed, supersededAppUninstallPendingReboot, removingSupersededApps, iosAppStoreUpdateFailedToInstall, vppAppHasUpdateAvailable, userRejectedUpdate, uninstallPendingReboot, supersedingAppsDetected, supersededAppsDetected, seeInstallErrorCode, autoInstallDisabled, managedAppNoLongerPresent, userRejectedInstall, userIsNotLoggedIntoAppStore, untargetedSupersedingAppsDetected, appRemovedBySupersedence, seeUninstallErrorCode, pendingReboot, installingDependencies, contentDownloaded, supersedingAppsNotApplicable, powerShellScriptRequirementNotMet, registryRequirementNotMet, fileSystemRequirementNotMet, platformNotApplicable, minimumCpuSpeedNotMet, minimumLogicalProcessorCountNotMet, minimumPhysicalMemoryNotMet, minimumOsVersionNotMet, minimumDiskSpaceNotMet, processorArchitectureNotApplicable.
    installStateDetail *ResultantAppStateDetail
    // Last sync date time
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
    mobileAppInstallStatusValue *ResultantAppState
    // OS Description
    osDescription *string
    // OS Version
    osVersion *string
    // Device User Name
    userName *string
    // User Principal Name
    userPrincipalName *string
}
// NewMobileAppInstallStatus instantiates a new mobileAppInstallStatus and sets the default values.
func NewMobileAppInstallStatus()(*MobileAppInstallStatus) {
    m := &MobileAppInstallStatus{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMobileAppInstallStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppInstallStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppInstallStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppInstallStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApp gets the app property value. The navigation link to the mobile app.
func (m *MobileAppInstallStatus) GetApp()(MobileAppable) {
    if m == nil {
        return nil
    } else {
        return m.app
    }
}
// GetDeviceId gets the deviceId property value. Device ID
func (m *MobileAppInstallStatus) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetDeviceName gets the deviceName property value. Device name
func (m *MobileAppInstallStatus) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetDisplayVersion gets the displayVersion property value. Human readable version of the application
func (m *MobileAppInstallStatus) GetDisplayVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayVersion
    }
}
// GetErrorCode gets the errorCode property value. The error code for install or uninstall failures.
func (m *MobileAppInstallStatus) GetErrorCode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppInstallStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["app"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMobileAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApp(val.(MobileAppable))
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
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
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["installState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseResultantAppState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallState(val.(*ResultantAppState))
        }
        return nil
    }
    res["installStateDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseResultantAppStateDetail)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallStateDetail(val.(*ResultantAppStateDetail))
        }
        return nil
    }
    res["lastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["mobileAppInstallStatusValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseResultantAppState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileAppInstallStatusValue(val.(*ResultantAppState))
        }
        return nil
    }
    res["osDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsDescription(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetInstallState gets the installState property value. The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
func (m *MobileAppInstallStatus) GetInstallState()(*ResultantAppState) {
    if m == nil {
        return nil
    } else {
        return m.installState
    }
}
// GetInstallStateDetail gets the installStateDetail property value. The install state detail of the app. Possible values are: noAdditionalDetails, dependencyFailedToInstall, dependencyWithRequirementsNotMet, dependencyPendingReboot, dependencyWithAutoInstallDisabled, supersededAppUninstallFailed, supersededAppUninstallPendingReboot, removingSupersededApps, iosAppStoreUpdateFailedToInstall, vppAppHasUpdateAvailable, userRejectedUpdate, uninstallPendingReboot, supersedingAppsDetected, supersededAppsDetected, seeInstallErrorCode, autoInstallDisabled, managedAppNoLongerPresent, userRejectedInstall, userIsNotLoggedIntoAppStore, untargetedSupersedingAppsDetected, appRemovedBySupersedence, seeUninstallErrorCode, pendingReboot, installingDependencies, contentDownloaded, supersedingAppsNotApplicable, powerShellScriptRequirementNotMet, registryRequirementNotMet, fileSystemRequirementNotMet, platformNotApplicable, minimumCpuSpeedNotMet, minimumLogicalProcessorCountNotMet, minimumPhysicalMemoryNotMet, minimumOsVersionNotMet, minimumDiskSpaceNotMet, processorArchitectureNotApplicable.
func (m *MobileAppInstallStatus) GetInstallStateDetail()(*ResultantAppStateDetail) {
    if m == nil {
        return nil
    } else {
        return m.installStateDetail
    }
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. Last sync date time
func (m *MobileAppInstallStatus) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// GetMobileAppInstallStatusValue gets the mobileAppInstallStatusValue property value. The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
func (m *MobileAppInstallStatus) GetMobileAppInstallStatusValue()(*ResultantAppState) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppInstallStatusValue
    }
}
// GetOsDescription gets the osDescription property value. OS Description
func (m *MobileAppInstallStatus) GetOsDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osDescription
    }
}
// GetOsVersion gets the osVersion property value. OS Version
func (m *MobileAppInstallStatus) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// GetUserName gets the userName property value. Device User Name
func (m *MobileAppInstallStatus) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name
func (m *MobileAppInstallStatus) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Serialize serializes information the current object
func (m *MobileAppInstallStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := (*m.GetInstallState()).String()
        err = writer.WriteStringValue("installState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetInstallStateDetail() != nil {
        cast := (*m.GetInstallStateDetail()).String()
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
        cast := (*m.GetMobileAppInstallStatusValue()).String()
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppInstallStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApp sets the app property value. The navigation link to the mobile app.
func (m *MobileAppInstallStatus) SetApp(value MobileAppable)() {
    if m != nil {
        m.app = value
    }
}
// SetDeviceId sets the deviceId property value. Device ID
func (m *MobileAppInstallStatus) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetDeviceName sets the deviceName property value. Device name
func (m *MobileAppInstallStatus) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetDisplayVersion sets the displayVersion property value. Human readable version of the application
func (m *MobileAppInstallStatus) SetDisplayVersion(value *string)() {
    if m != nil {
        m.displayVersion = value
    }
}
// SetErrorCode sets the errorCode property value. The error code for install or uninstall failures.
func (m *MobileAppInstallStatus) SetErrorCode(value *int32)() {
    if m != nil {
        m.errorCode = value
    }
}
// SetInstallState sets the installState property value. The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
func (m *MobileAppInstallStatus) SetInstallState(value *ResultantAppState)() {
    if m != nil {
        m.installState = value
    }
}
// SetInstallStateDetail sets the installStateDetail property value. The install state detail of the app. Possible values are: noAdditionalDetails, dependencyFailedToInstall, dependencyWithRequirementsNotMet, dependencyPendingReboot, dependencyWithAutoInstallDisabled, supersededAppUninstallFailed, supersededAppUninstallPendingReboot, removingSupersededApps, iosAppStoreUpdateFailedToInstall, vppAppHasUpdateAvailable, userRejectedUpdate, uninstallPendingReboot, supersedingAppsDetected, supersededAppsDetected, seeInstallErrorCode, autoInstallDisabled, managedAppNoLongerPresent, userRejectedInstall, userIsNotLoggedIntoAppStore, untargetedSupersedingAppsDetected, appRemovedBySupersedence, seeUninstallErrorCode, pendingReboot, installingDependencies, contentDownloaded, supersedingAppsNotApplicable, powerShellScriptRequirementNotMet, registryRequirementNotMet, fileSystemRequirementNotMet, platformNotApplicable, minimumCpuSpeedNotMet, minimumLogicalProcessorCountNotMet, minimumPhysicalMemoryNotMet, minimumOsVersionNotMet, minimumDiskSpaceNotMet, processorArchitectureNotApplicable.
func (m *MobileAppInstallStatus) SetInstallStateDetail(value *ResultantAppStateDetail)() {
    if m != nil {
        m.installStateDetail = value
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. Last sync date time
func (m *MobileAppInstallStatus) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSyncDateTime = value
    }
}
// SetMobileAppInstallStatusValue sets the mobileAppInstallStatusValue property value. The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
func (m *MobileAppInstallStatus) SetMobileAppInstallStatusValue(value *ResultantAppState)() {
    if m != nil {
        m.mobileAppInstallStatusValue = value
    }
}
// SetOsDescription sets the osDescription property value. OS Description
func (m *MobileAppInstallStatus) SetOsDescription(value *string)() {
    if m != nil {
        m.osDescription = value
    }
}
// SetOsVersion sets the osVersion property value. OS Version
func (m *MobileAppInstallStatus) SetOsVersion(value *string)() {
    if m != nil {
        m.osVersion = value
    }
}
// SetUserName sets the userName property value. Device User Name
func (m *MobileAppInstallStatus) SetUserName(value *string)() {
    if m != nil {
        m.userName = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name
func (m *MobileAppInstallStatus) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
