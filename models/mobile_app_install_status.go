package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppInstallStatus contains properties for the installation state of a mobile app for a device.
type MobileAppInstallStatus struct {
    Entity
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
    // A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
    installState *ResultantAppState
    // Enum indicating additional details regarding why an application has a particular install state.
    installStateDetail *ResultantAppStateDetail
    // Last sync date time
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
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
    return m
}
// CreateMobileAppInstallStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppInstallStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppInstallStatus(), nil
}
// GetApp gets the app property value. The navigation link to the mobile app.
func (m *MobileAppInstallStatus) GetApp()(MobileAppable) {
    return m.app
}
// GetDeviceId gets the deviceId property value. Device ID
func (m *MobileAppInstallStatus) GetDeviceId()(*string) {
    return m.deviceId
}
// GetDeviceName gets the deviceName property value. Device name
func (m *MobileAppInstallStatus) GetDeviceName()(*string) {
    return m.deviceName
}
// GetDisplayVersion gets the displayVersion property value. Human readable version of the application
func (m *MobileAppInstallStatus) GetDisplayVersion()(*string) {
    return m.displayVersion
}
// GetErrorCode gets the errorCode property value. The error code for install or uninstall failures.
func (m *MobileAppInstallStatus) GetErrorCode()(*int32) {
    return m.errorCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppInstallStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["app"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateMobileAppFromDiscriminatorValue , m.SetApp)
    res["deviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceId)
    res["deviceName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceName)
    res["displayVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayVersion)
    res["errorCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetErrorCode)
    res["installState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseResultantAppState , m.SetInstallState)
    res["installStateDetail"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseResultantAppStateDetail , m.SetInstallStateDetail)
    res["lastSyncDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastSyncDateTime)
    res["mobileAppInstallStatusValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseResultantAppState , m.SetMobileAppInstallStatusValue)
    res["osDescription"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOsDescription)
    res["osVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOsVersion)
    res["userName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserName)
    res["userPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserPrincipalName)
    return res
}
// GetInstallState gets the installState property value. A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
func (m *MobileAppInstallStatus) GetInstallState()(*ResultantAppState) {
    return m.installState
}
// GetInstallStateDetail gets the installStateDetail property value. Enum indicating additional details regarding why an application has a particular install state.
func (m *MobileAppInstallStatus) GetInstallStateDetail()(*ResultantAppStateDetail) {
    return m.installStateDetail
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. Last sync date time
func (m *MobileAppInstallStatus) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetMobileAppInstallStatusValue gets the mobileAppInstallStatusValue property value. A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
func (m *MobileAppInstallStatus) GetMobileAppInstallStatusValue()(*ResultantAppState) {
    return m.mobileAppInstallStatusValue
}
// GetOsDescription gets the osDescription property value. OS Description
func (m *MobileAppInstallStatus) GetOsDescription()(*string) {
    return m.osDescription
}
// GetOsVersion gets the osVersion property value. OS Version
func (m *MobileAppInstallStatus) GetOsVersion()(*string) {
    return m.osVersion
}
// GetUserName gets the userName property value. Device User Name
func (m *MobileAppInstallStatus) GetUserName()(*string) {
    return m.userName
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name
func (m *MobileAppInstallStatus) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
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
    return nil
}
// SetApp sets the app property value. The navigation link to the mobile app.
func (m *MobileAppInstallStatus) SetApp(value MobileAppable)() {
    m.app = value
}
// SetDeviceId sets the deviceId property value. Device ID
func (m *MobileAppInstallStatus) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetDeviceName sets the deviceName property value. Device name
func (m *MobileAppInstallStatus) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetDisplayVersion sets the displayVersion property value. Human readable version of the application
func (m *MobileAppInstallStatus) SetDisplayVersion(value *string)() {
    m.displayVersion = value
}
// SetErrorCode sets the errorCode property value. The error code for install or uninstall failures.
func (m *MobileAppInstallStatus) SetErrorCode(value *int32)() {
    m.errorCode = value
}
// SetInstallState sets the installState property value. A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
func (m *MobileAppInstallStatus) SetInstallState(value *ResultantAppState)() {
    m.installState = value
}
// SetInstallStateDetail sets the installStateDetail property value. Enum indicating additional details regarding why an application has a particular install state.
func (m *MobileAppInstallStatus) SetInstallStateDetail(value *ResultantAppStateDetail)() {
    m.installStateDetail = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. Last sync date time
func (m *MobileAppInstallStatus) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetMobileAppInstallStatusValue sets the mobileAppInstallStatusValue property value. A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
func (m *MobileAppInstallStatus) SetMobileAppInstallStatusValue(value *ResultantAppState)() {
    m.mobileAppInstallStatusValue = value
}
// SetOsDescription sets the osDescription property value. OS Description
func (m *MobileAppInstallStatus) SetOsDescription(value *string)() {
    m.osDescription = value
}
// SetOsVersion sets the osVersion property value. OS Version
func (m *MobileAppInstallStatus) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetUserName sets the userName property value. Device User Name
func (m *MobileAppInstallStatus) SetUserName(value *string)() {
    m.userName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name
func (m *MobileAppInstallStatus) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
