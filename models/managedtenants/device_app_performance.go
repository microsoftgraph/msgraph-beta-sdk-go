package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type DeviceAppPerformance struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewDeviceAppPerformance instantiates a new DeviceAppPerformance and sets the default values.
func NewDeviceAppPerformance()(*DeviceAppPerformance) {
    m := &DeviceAppPerformance{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateDeviceAppPerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceAppPerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppPerformance(), nil
}
// GetAppFriendlyName gets the appFriendlyName property value. The appFriendlyName property
// returns a *string when successful
func (m *DeviceAppPerformance) GetAppFriendlyName()(*string) {
    val, err := m.GetBackingStore().Get("appFriendlyName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppName gets the appName property value. The appName property
// returns a *string when successful
func (m *DeviceAppPerformance) GetAppName()(*string) {
    val, err := m.GetBackingStore().Get("appName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppPublisher gets the appPublisher property value. The appPublisher property
// returns a *string when successful
func (m *DeviceAppPerformance) GetAppPublisher()(*string) {
    val, err := m.GetBackingStore().Get("appPublisher")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppVersion gets the appVersion property value. The appVersion property
// returns a *string when successful
func (m *DeviceAppPerformance) GetAppVersion()(*string) {
    val, err := m.GetBackingStore().Get("appVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The deviceId property
// returns a *string when successful
func (m *DeviceAppPerformance) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceManufacturer gets the deviceManufacturer property value. The deviceManufacturer property
// returns a *string when successful
func (m *DeviceAppPerformance) GetDeviceManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("deviceManufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceModel gets the deviceModel property value. The deviceModel property
// returns a *string when successful
func (m *DeviceAppPerformance) GetDeviceModel()(*string) {
    val, err := m.GetBackingStore().Get("deviceModel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The deviceName property
// returns a *string when successful
func (m *DeviceAppPerformance) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceAppPerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appFriendlyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppFriendlyName(val)
        }
        return nil
    }
    res["appName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppName(val)
        }
        return nil
    }
    res["appPublisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppPublisher(val)
        }
        return nil
    }
    res["appVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppVersion(val)
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
    res["deviceManufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManufacturer(val)
        }
        return nil
    }
    res["deviceModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceModel(val)
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
    res["healthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatus(val)
        }
        return nil
    }
    res["isLatestUsedVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLatestUsedVersion(val)
        }
        return nil
    }
    res["isMostUsedVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMostUsedVersion(val)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["tenantDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantDisplayName(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["totalAppCrashCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAppCrashCount(val)
        }
        return nil
    }
    res["totalAppFreezeCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAppFreezeCount(val)
        }
        return nil
    }
    return res
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
// returns a *string when successful
func (m *DeviceAppPerformance) GetHealthStatus()(*string) {
    val, err := m.GetBackingStore().Get("healthStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsLatestUsedVersion gets the isLatestUsedVersion property value. The isLatestUsedVersion property
// returns a *int32 when successful
func (m *DeviceAppPerformance) GetIsLatestUsedVersion()(*int32) {
    val, err := m.GetBackingStore().Get("isLatestUsedVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetIsMostUsedVersion gets the isMostUsedVersion property value. The isMostUsedVersion property
// returns a *int32 when successful
func (m *DeviceAppPerformance) GetIsMostUsedVersion()(*int32) {
    val, err := m.GetBackingStore().Get("isMostUsedVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The lastUpdatedDateTime property
// returns a *Time when successful
func (m *DeviceAppPerformance) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastUpdatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetTenantDisplayName gets the tenantDisplayName property value. The tenantDisplayName property
// returns a *string when successful
func (m *DeviceAppPerformance) GetTenantDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("tenantDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantId gets the tenantId property value. The tenantId property
// returns a *string when successful
func (m *DeviceAppPerformance) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTotalAppCrashCount gets the totalAppCrashCount property value. The totalAppCrashCount property
// returns a *int32 when successful
func (m *DeviceAppPerformance) GetTotalAppCrashCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalAppCrashCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalAppFreezeCount gets the totalAppFreezeCount property value. The totalAppFreezeCount property
// returns a *int32 when successful
func (m *DeviceAppPerformance) GetTotalAppFreezeCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalAppFreezeCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceAppPerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appFriendlyName", m.GetAppFriendlyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appName", m.GetAppName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appPublisher", m.GetAppPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appVersion", m.GetAppVersion())
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
        err = writer.WriteStringValue("deviceManufacturer", m.GetDeviceManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceModel", m.GetDeviceModel())
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
        err = writer.WriteStringValue("healthStatus", m.GetHealthStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("isLatestUsedVersion", m.GetIsLatestUsedVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("isMostUsedVersion", m.GetIsMostUsedVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantDisplayName", m.GetTenantDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalAppCrashCount", m.GetTotalAppCrashCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalAppFreezeCount", m.GetTotalAppFreezeCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppFriendlyName sets the appFriendlyName property value. The appFriendlyName property
func (m *DeviceAppPerformance) SetAppFriendlyName(value *string)() {
    err := m.GetBackingStore().Set("appFriendlyName", value)
    if err != nil {
        panic(err)
    }
}
// SetAppName sets the appName property value. The appName property
func (m *DeviceAppPerformance) SetAppName(value *string)() {
    err := m.GetBackingStore().Set("appName", value)
    if err != nil {
        panic(err)
    }
}
// SetAppPublisher sets the appPublisher property value. The appPublisher property
func (m *DeviceAppPerformance) SetAppPublisher(value *string)() {
    err := m.GetBackingStore().Set("appPublisher", value)
    if err != nil {
        panic(err)
    }
}
// SetAppVersion sets the appVersion property value. The appVersion property
func (m *DeviceAppPerformance) SetAppVersion(value *string)() {
    err := m.GetBackingStore().Set("appVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The deviceId property
func (m *DeviceAppPerformance) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceManufacturer sets the deviceManufacturer property value. The deviceManufacturer property
func (m *DeviceAppPerformance) SetDeviceManufacturer(value *string)() {
    err := m.GetBackingStore().Set("deviceManufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceModel sets the deviceModel property value. The deviceModel property
func (m *DeviceAppPerformance) SetDeviceModel(value *string)() {
    err := m.GetBackingStore().Set("deviceModel", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The deviceName property
func (m *DeviceAppPerformance) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *DeviceAppPerformance) SetHealthStatus(value *string)() {
    err := m.GetBackingStore().Set("healthStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetIsLatestUsedVersion sets the isLatestUsedVersion property value. The isLatestUsedVersion property
func (m *DeviceAppPerformance) SetIsLatestUsedVersion(value *int32)() {
    err := m.GetBackingStore().Set("isLatestUsedVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetIsMostUsedVersion sets the isMostUsedVersion property value. The isMostUsedVersion property
func (m *DeviceAppPerformance) SetIsMostUsedVersion(value *int32)() {
    err := m.GetBackingStore().Set("isMostUsedVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The lastUpdatedDateTime property
func (m *DeviceAppPerformance) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastUpdatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantDisplayName sets the tenantDisplayName property value. The tenantDisplayName property
func (m *DeviceAppPerformance) SetTenantDisplayName(value *string)() {
    err := m.GetBackingStore().Set("tenantDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *DeviceAppPerformance) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalAppCrashCount sets the totalAppCrashCount property value. The totalAppCrashCount property
func (m *DeviceAppPerformance) SetTotalAppCrashCount(value *int32)() {
    err := m.GetBackingStore().Set("totalAppCrashCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalAppFreezeCount sets the totalAppFreezeCount property value. The totalAppFreezeCount property
func (m *DeviceAppPerformance) SetTotalAppFreezeCount(value *int32)() {
    err := m.GetBackingStore().Set("totalAppFreezeCount", value)
    if err != nil {
        panic(err)
    }
}
type DeviceAppPerformanceable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppFriendlyName()(*string)
    GetAppName()(*string)
    GetAppPublisher()(*string)
    GetAppVersion()(*string)
    GetDeviceId()(*string)
    GetDeviceManufacturer()(*string)
    GetDeviceModel()(*string)
    GetDeviceName()(*string)
    GetHealthStatus()(*string)
    GetIsLatestUsedVersion()(*int32)
    GetIsMostUsedVersion()(*int32)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    GetTotalAppCrashCount()(*int32)
    GetTotalAppFreezeCount()(*int32)
    SetAppFriendlyName(value *string)()
    SetAppName(value *string)()
    SetAppPublisher(value *string)()
    SetAppVersion(value *string)()
    SetDeviceId(value *string)()
    SetDeviceManufacturer(value *string)()
    SetDeviceModel(value *string)()
    SetDeviceName(value *string)()
    SetHealthStatus(value *string)()
    SetIsLatestUsedVersion(value *int32)()
    SetIsMostUsedVersion(value *int32)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
    SetTotalAppCrashCount(value *int32)()
    SetTotalAppFreezeCount(value *int32)()
}
