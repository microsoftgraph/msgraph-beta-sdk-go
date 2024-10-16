package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppInstallStatus contains properties for the installation state of a mobile app for a device. This will be deprecated starting May, 2023 (Intune Release 2305).
type MobileAppInstallStatus struct {
    Entity
}
// NewMobileAppInstallStatus instantiates a new MobileAppInstallStatus and sets the default values.
func NewMobileAppInstallStatus()(*MobileAppInstallStatus) {
    m := &MobileAppInstallStatus{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileAppInstallStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMobileAppInstallStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppInstallStatus(), nil
}
// GetApp gets the app property value. The navigation link to the mobile app.
// returns a MobileAppable when successful
func (m *MobileAppInstallStatus) GetApp()(MobileAppable) {
    val, err := m.GetBackingStore().Get("app")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MobileAppable)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. Device ID
// returns a *string when successful
func (m *MobileAppInstallStatus) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. Device name
// returns a *string when successful
func (m *MobileAppInstallStatus) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayVersion gets the displayVersion property value. Human readable version of the application
// returns a *string when successful
func (m *MobileAppInstallStatus) GetDisplayVersion()(*string) {
    val, err := m.GetBackingStore().Get("displayVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetErrorCode gets the errorCode property value. The error code for install or uninstall failures.
// returns a *int32 when successful
func (m *MobileAppInstallStatus) GetErrorCode()(*int32) {
    val, err := m.GetBackingStore().Get("errorCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
// GetInstallState gets the installState property value. A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
// returns a *ResultantAppState when successful
func (m *MobileAppInstallStatus) GetInstallState()(*ResultantAppState) {
    val, err := m.GetBackingStore().Get("installState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ResultantAppState)
    }
    return nil
}
// GetInstallStateDetail gets the installStateDetail property value. Enum indicating additional details regarding why an application has a particular install state.
// returns a *ResultantAppStateDetail when successful
func (m *MobileAppInstallStatus) GetInstallStateDetail()(*ResultantAppStateDetail) {
    val, err := m.GetBackingStore().Get("installStateDetail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ResultantAppStateDetail)
    }
    return nil
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. Last sync date time
// returns a *Time when successful
func (m *MobileAppInstallStatus) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastSyncDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMobileAppInstallStatusValue gets the mobileAppInstallStatusValue property value. A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
// returns a *ResultantAppState when successful
func (m *MobileAppInstallStatus) GetMobileAppInstallStatusValue()(*ResultantAppState) {
    val, err := m.GetBackingStore().Get("mobileAppInstallStatusValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ResultantAppState)
    }
    return nil
}
// GetOsDescription gets the osDescription property value. OS Description
// returns a *string when successful
func (m *MobileAppInstallStatus) GetOsDescription()(*string) {
    val, err := m.GetBackingStore().Get("osDescription")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. OS Version
// returns a *string when successful
func (m *MobileAppInstallStatus) GetOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserName gets the userName property value. Device User Name
// returns a *string when successful
func (m *MobileAppInstallStatus) GetUserName()(*string) {
    val, err := m.GetBackingStore().Get("userName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name
// returns a *string when successful
func (m *MobileAppInstallStatus) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    err := m.GetBackingStore().Set("app", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. Device ID
func (m *MobileAppInstallStatus) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. Device name
func (m *MobileAppInstallStatus) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayVersion sets the displayVersion property value. Human readable version of the application
func (m *MobileAppInstallStatus) SetDisplayVersion(value *string)() {
    err := m.GetBackingStore().Set("displayVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorCode sets the errorCode property value. The error code for install or uninstall failures.
func (m *MobileAppInstallStatus) SetErrorCode(value *int32)() {
    err := m.GetBackingStore().Set("errorCode", value)
    if err != nil {
        panic(err)
    }
}
// SetInstallState sets the installState property value. A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
func (m *MobileAppInstallStatus) SetInstallState(value *ResultantAppState)() {
    err := m.GetBackingStore().Set("installState", value)
    if err != nil {
        panic(err)
    }
}
// SetInstallStateDetail sets the installStateDetail property value. Enum indicating additional details regarding why an application has a particular install state.
func (m *MobileAppInstallStatus) SetInstallStateDetail(value *ResultantAppStateDetail)() {
    err := m.GetBackingStore().Set("installStateDetail", value)
    if err != nil {
        panic(err)
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. Last sync date time
func (m *MobileAppInstallStatus) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastSyncDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMobileAppInstallStatusValue sets the mobileAppInstallStatusValue property value. A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
func (m *MobileAppInstallStatus) SetMobileAppInstallStatusValue(value *ResultantAppState)() {
    err := m.GetBackingStore().Set("mobileAppInstallStatusValue", value)
    if err != nil {
        panic(err)
    }
}
// SetOsDescription sets the osDescription property value. OS Description
func (m *MobileAppInstallStatus) SetOsDescription(value *string)() {
    err := m.GetBackingStore().Set("osDescription", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. OS Version
func (m *MobileAppInstallStatus) SetOsVersion(value *string)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetUserName sets the userName property value. Device User Name
func (m *MobileAppInstallStatus) SetUserName(value *string)() {
    err := m.GetBackingStore().Set("userName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name
func (m *MobileAppInstallStatus) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
type MobileAppInstallStatusable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApp()(MobileAppable)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetDisplayVersion()(*string)
    GetErrorCode()(*int32)
    GetInstallState()(*ResultantAppState)
    GetInstallStateDetail()(*ResultantAppStateDetail)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMobileAppInstallStatusValue()(*ResultantAppState)
    GetOsDescription()(*string)
    GetOsVersion()(*string)
    GetUserName()(*string)
    GetUserPrincipalName()(*string)
    SetApp(value MobileAppable)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetDisplayVersion(value *string)()
    SetErrorCode(value *int32)()
    SetInstallState(value *ResultantAppState)()
    SetInstallStateDetail(value *ResultantAppStateDetail)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMobileAppInstallStatusValue(value *ResultantAppState)()
    SetOsDescription(value *string)()
    SetOsVersion(value *string)()
    SetUserName(value *string)()
    SetUserPrincipalName(value *string)()
}
