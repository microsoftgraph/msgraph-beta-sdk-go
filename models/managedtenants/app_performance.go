package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type AppPerformance struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewAppPerformance instantiates a new AppPerformance and sets the default values.
func NewAppPerformance()(*AppPerformance) {
    m := &AppPerformance{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateAppPerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppPerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppPerformance(), nil
}
// GetAppFriendlyName gets the appFriendlyName property value. The appFriendlyName property
// returns a *string when successful
func (m *AppPerformance) GetAppFriendlyName()(*string) {
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
func (m *AppPerformance) GetAppName()(*string) {
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
func (m *AppPerformance) GetAppPublisher()(*string) {
    val, err := m.GetBackingStore().Get("appPublisher")
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
func (m *AppPerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["meanTimeToFailureInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeanTimeToFailureInMinutes(val)
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
    res["totalActiveDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalActiveDeviceCount(val)
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
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The lastUpdatedDateTime property
// returns a *Time when successful
func (m *AppPerformance) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastUpdatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMeanTimeToFailureInMinutes gets the meanTimeToFailureInMinutes property value. The meanTimeToFailureInMinutes property
// returns a *int32 when successful
func (m *AppPerformance) GetMeanTimeToFailureInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("meanTimeToFailureInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTenantDisplayName gets the tenantDisplayName property value. The tenantDisplayName property
// returns a *string when successful
func (m *AppPerformance) GetTenantDisplayName()(*string) {
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
func (m *AppPerformance) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTotalActiveDeviceCount gets the totalActiveDeviceCount property value. The totalActiveDeviceCount property
// returns a *int32 when successful
func (m *AppPerformance) GetTotalActiveDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalActiveDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalAppCrashCount gets the totalAppCrashCount property value. The totalAppCrashCount property
// returns a *int32 when successful
func (m *AppPerformance) GetTotalAppCrashCount()(*int32) {
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
func (m *AppPerformance) GetTotalAppFreezeCount()(*int32) {
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
func (m *AppPerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("meanTimeToFailureInMinutes", m.GetMeanTimeToFailureInMinutes())
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
        err = writer.WriteInt32Value("totalActiveDeviceCount", m.GetTotalActiveDeviceCount())
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
func (m *AppPerformance) SetAppFriendlyName(value *string)() {
    err := m.GetBackingStore().Set("appFriendlyName", value)
    if err != nil {
        panic(err)
    }
}
// SetAppName sets the appName property value. The appName property
func (m *AppPerformance) SetAppName(value *string)() {
    err := m.GetBackingStore().Set("appName", value)
    if err != nil {
        panic(err)
    }
}
// SetAppPublisher sets the appPublisher property value. The appPublisher property
func (m *AppPerformance) SetAppPublisher(value *string)() {
    err := m.GetBackingStore().Set("appPublisher", value)
    if err != nil {
        panic(err)
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The lastUpdatedDateTime property
func (m *AppPerformance) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastUpdatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMeanTimeToFailureInMinutes sets the meanTimeToFailureInMinutes property value. The meanTimeToFailureInMinutes property
func (m *AppPerformance) SetMeanTimeToFailureInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("meanTimeToFailureInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantDisplayName sets the tenantDisplayName property value. The tenantDisplayName property
func (m *AppPerformance) SetTenantDisplayName(value *string)() {
    err := m.GetBackingStore().Set("tenantDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *AppPerformance) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalActiveDeviceCount sets the totalActiveDeviceCount property value. The totalActiveDeviceCount property
func (m *AppPerformance) SetTotalActiveDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("totalActiveDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalAppCrashCount sets the totalAppCrashCount property value. The totalAppCrashCount property
func (m *AppPerformance) SetTotalAppCrashCount(value *int32)() {
    err := m.GetBackingStore().Set("totalAppCrashCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalAppFreezeCount sets the totalAppFreezeCount property value. The totalAppFreezeCount property
func (m *AppPerformance) SetTotalAppFreezeCount(value *int32)() {
    err := m.GetBackingStore().Set("totalAppFreezeCount", value)
    if err != nil {
        panic(err)
    }
}
type AppPerformanceable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppFriendlyName()(*string)
    GetAppName()(*string)
    GetAppPublisher()(*string)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMeanTimeToFailureInMinutes()(*int32)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    GetTotalActiveDeviceCount()(*int32)
    GetTotalAppCrashCount()(*int32)
    GetTotalAppFreezeCount()(*int32)
    SetAppFriendlyName(value *string)()
    SetAppName(value *string)()
    SetAppPublisher(value *string)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMeanTimeToFailureInMinutes(value *int32)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
    SetTotalActiveDeviceCount(value *int32)()
    SetTotalAppCrashCount(value *int32)()
    SetTotalAppFreezeCount(value *int32)()
}
