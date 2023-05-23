package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DeviceHealthStatus 
type DeviceHealthStatus struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewDeviceHealthStatus instantiates a new DeviceHealthStatus and sets the default values.
func NewDeviceHealthStatus()(*DeviceHealthStatus) {
    m := &DeviceHealthStatus{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateDeviceHealthStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthStatus(), nil
}
// GetBlueScreenCount gets the blueScreenCount property value. The blueScreenCount property
func (m *DeviceHealthStatus) GetBlueScreenCount()(*int32) {
    val, err := m.GetBackingStore().Get("blueScreenCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetBootTotalDurationInSeconds gets the bootTotalDurationInSeconds property value. The bootTotalDurationInSeconds property
func (m *DeviceHealthStatus) GetBootTotalDurationInSeconds()(*float64) {
    val, err := m.GetBackingStore().Get("bootTotalDurationInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The deviceId property
func (m *DeviceHealthStatus) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceMake gets the deviceMake property value. The deviceMake property
func (m *DeviceHealthStatus) GetDeviceMake()(*string) {
    val, err := m.GetBackingStore().Get("deviceMake")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceModel gets the deviceModel property value. The deviceModel property
func (m *DeviceHealthStatus) GetDeviceModel()(*string) {
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
func (m *DeviceHealthStatus) GetDeviceName()(*string) {
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
func (m *DeviceHealthStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["blueScreenCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlueScreenCount(val)
        }
        return nil
    }
    res["bootTotalDurationInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBootTotalDurationInSeconds(val)
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
    res["deviceMake"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceMake(val)
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
    res["primaryDiskType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryDiskType(val)
        }
        return nil
    }
    res["restartCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartCount(val)
        }
        return nil
    }
    res["startupPerformanceScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupPerformanceScore(val)
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
    res["topProcesses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopProcesses(val)
        }
        return nil
    }
    return res
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
func (m *DeviceHealthStatus) GetHealthStatus()(*string) {
    val, err := m.GetBackingStore().Get("healthStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The lastUpdatedDateTime property
func (m *DeviceHealthStatus) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastUpdatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. The osVersion property
func (m *DeviceHealthStatus) GetOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrimaryDiskType gets the primaryDiskType property value. The primaryDiskType property
func (m *DeviceHealthStatus) GetPrimaryDiskType()(*string) {
    val, err := m.GetBackingStore().Get("primaryDiskType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRestartCount gets the restartCount property value. The restartCount property
func (m *DeviceHealthStatus) GetRestartCount()(*int32) {
    val, err := m.GetBackingStore().Get("restartCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetStartupPerformanceScore gets the startupPerformanceScore property value. The startupPerformanceScore property
func (m *DeviceHealthStatus) GetStartupPerformanceScore()(*float64) {
    val, err := m.GetBackingStore().Get("startupPerformanceScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetTenantDisplayName gets the tenantDisplayName property value. The tenantDisplayName property
func (m *DeviceHealthStatus) GetTenantDisplayName()(*string) {
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
func (m *DeviceHealthStatus) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTopProcesses gets the topProcesses property value. The topProcesses property
func (m *DeviceHealthStatus) GetTopProcesses()(*string) {
    val, err := m.GetBackingStore().Get("topProcesses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceHealthStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("blueScreenCount", m.GetBlueScreenCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("bootTotalDurationInSeconds", m.GetBootTotalDurationInSeconds())
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
        err = writer.WriteStringValue("deviceMake", m.GetDeviceMake())
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
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
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
        err = writer.WriteStringValue("primaryDiskType", m.GetPrimaryDiskType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("restartCount", m.GetRestartCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("startupPerformanceScore", m.GetStartupPerformanceScore())
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
        err = writer.WriteStringValue("topProcesses", m.GetTopProcesses())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBlueScreenCount sets the blueScreenCount property value. The blueScreenCount property
func (m *DeviceHealthStatus) SetBlueScreenCount(value *int32)() {
    err := m.GetBackingStore().Set("blueScreenCount", value)
    if err != nil {
        panic(err)
    }
}
// SetBootTotalDurationInSeconds sets the bootTotalDurationInSeconds property value. The bootTotalDurationInSeconds property
func (m *DeviceHealthStatus) SetBootTotalDurationInSeconds(value *float64)() {
    err := m.GetBackingStore().Set("bootTotalDurationInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The deviceId property
func (m *DeviceHealthStatus) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceMake sets the deviceMake property value. The deviceMake property
func (m *DeviceHealthStatus) SetDeviceMake(value *string)() {
    err := m.GetBackingStore().Set("deviceMake", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceModel sets the deviceModel property value. The deviceModel property
func (m *DeviceHealthStatus) SetDeviceModel(value *string)() {
    err := m.GetBackingStore().Set("deviceModel", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The deviceName property
func (m *DeviceHealthStatus) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *DeviceHealthStatus) SetHealthStatus(value *string)() {
    err := m.GetBackingStore().Set("healthStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The lastUpdatedDateTime property
func (m *DeviceHealthStatus) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastUpdatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. The osVersion property
func (m *DeviceHealthStatus) SetOsVersion(value *string)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetPrimaryDiskType sets the primaryDiskType property value. The primaryDiskType property
func (m *DeviceHealthStatus) SetPrimaryDiskType(value *string)() {
    err := m.GetBackingStore().Set("primaryDiskType", value)
    if err != nil {
        panic(err)
    }
}
// SetRestartCount sets the restartCount property value. The restartCount property
func (m *DeviceHealthStatus) SetRestartCount(value *int32)() {
    err := m.GetBackingStore().Set("restartCount", value)
    if err != nil {
        panic(err)
    }
}
// SetStartupPerformanceScore sets the startupPerformanceScore property value. The startupPerformanceScore property
func (m *DeviceHealthStatus) SetStartupPerformanceScore(value *float64)() {
    err := m.GetBackingStore().Set("startupPerformanceScore", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantDisplayName sets the tenantDisplayName property value. The tenantDisplayName property
func (m *DeviceHealthStatus) SetTenantDisplayName(value *string)() {
    err := m.GetBackingStore().Set("tenantDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *DeviceHealthStatus) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetTopProcesses sets the topProcesses property value. The topProcesses property
func (m *DeviceHealthStatus) SetTopProcesses(value *string)() {
    err := m.GetBackingStore().Set("topProcesses", value)
    if err != nil {
        panic(err)
    }
}
// DeviceHealthStatusable 
type DeviceHealthStatusable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBlueScreenCount()(*int32)
    GetBootTotalDurationInSeconds()(*float64)
    GetDeviceId()(*string)
    GetDeviceMake()(*string)
    GetDeviceModel()(*string)
    GetDeviceName()(*string)
    GetHealthStatus()(*string)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOsVersion()(*string)
    GetPrimaryDiskType()(*string)
    GetRestartCount()(*int32)
    GetStartupPerformanceScore()(*float64)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    GetTopProcesses()(*string)
    SetBlueScreenCount(value *int32)()
    SetBootTotalDurationInSeconds(value *float64)()
    SetDeviceId(value *string)()
    SetDeviceMake(value *string)()
    SetDeviceModel(value *string)()
    SetDeviceName(value *string)()
    SetHealthStatus(value *string)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOsVersion(value *string)()
    SetPrimaryDiskType(value *string)()
    SetRestartCount(value *int32)()
    SetStartupPerformanceScore(value *float64)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
    SetTopProcesses(value *string)()
}
