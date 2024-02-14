package networkaccess

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type CrossTenantAccess struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCrossTenantAccess instantiates a new CrossTenantAccess and sets the default values.
func NewCrossTenantAccess()(*CrossTenantAccess) {
    m := &CrossTenantAccess{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCrossTenantAccessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCrossTenantAccessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantAccess(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CrossTenantAccess) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *CrossTenantAccess) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDeviceCount gets the deviceCount property value. The number of devices that accessed the external tenant.
// returns a *int64 when successful
func (m *CrossTenantAccess) GetDeviceCount()(*int64) {
    val, err := m.GetBackingStore().Get("deviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CrossTenantAccess) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
        }
        return nil
    }
    res["lastAccessDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastAccessDateTime(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["resourceTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceTenantId(val)
        }
        return nil
    }
    res["resourceTenantName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceTenantName(val)
        }
        return nil
    }
    res["resourceTenantPrimaryDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceTenantPrimaryDomain(val)
        }
        return nil
    }
    res["usageStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUsageStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsageStatus(val.(*UsageStatus))
        }
        return nil
    }
    res["userCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserCount(val)
        }
        return nil
    }
    return res
}
// GetLastAccessDateTime gets the lastAccessDateTime property value. The timestamp of the most recent access to the external tenant.
// returns a *Time when successful
func (m *CrossTenantAccess) GetLastAccessDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastAccessDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CrossTenantAccess) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResourceTenantId gets the resourceTenantId property value. The tenant ID of the external tenant.
// returns a *string when successful
func (m *CrossTenantAccess) GetResourceTenantId()(*string) {
    val, err := m.GetBackingStore().Get("resourceTenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResourceTenantName gets the resourceTenantName property value. The name of the external tenant.
// returns a *string when successful
func (m *CrossTenantAccess) GetResourceTenantName()(*string) {
    val, err := m.GetBackingStore().Get("resourceTenantName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResourceTenantPrimaryDomain gets the resourceTenantPrimaryDomain property value. The domain of the external tenant.
// returns a *string when successful
func (m *CrossTenantAccess) GetResourceTenantPrimaryDomain()(*string) {
    val, err := m.GetBackingStore().Get("resourceTenantPrimaryDomain")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUsageStatus gets the usageStatus property value. The usageStatus property
// returns a *UsageStatus when successful
func (m *CrossTenantAccess) GetUsageStatus()(*UsageStatus) {
    val, err := m.GetBackingStore().Get("usageStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UsageStatus)
    }
    return nil
}
// GetUserCount gets the userCount property value. The number of users that accessed the external tenant.
// returns a *int64 when successful
func (m *CrossTenantAccess) GetUserCount()(*int64) {
    val, err := m.GetBackingStore().Get("userCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CrossTenantAccess) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("deviceCount", m.GetDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastAccessDateTime", m.GetLastAccessDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceTenantId", m.GetResourceTenantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceTenantName", m.GetResourceTenantName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceTenantPrimaryDomain", m.GetResourceTenantPrimaryDomain())
        if err != nil {
            return err
        }
    }
    if m.GetUsageStatus() != nil {
        cast := (*m.GetUsageStatus()).String()
        err := writer.WriteStringValue("usageStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("userCount", m.GetUserCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CrossTenantAccess) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CrossTenantAccess) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDeviceCount sets the deviceCount property value. The number of devices that accessed the external tenant.
func (m *CrossTenantAccess) SetDeviceCount(value *int64)() {
    err := m.GetBackingStore().Set("deviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetLastAccessDateTime sets the lastAccessDateTime property value. The timestamp of the most recent access to the external tenant.
func (m *CrossTenantAccess) SetLastAccessDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastAccessDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CrossTenantAccess) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceTenantId sets the resourceTenantId property value. The tenant ID of the external tenant.
func (m *CrossTenantAccess) SetResourceTenantId(value *string)() {
    err := m.GetBackingStore().Set("resourceTenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceTenantName sets the resourceTenantName property value. The name of the external tenant.
func (m *CrossTenantAccess) SetResourceTenantName(value *string)() {
    err := m.GetBackingStore().Set("resourceTenantName", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceTenantPrimaryDomain sets the resourceTenantPrimaryDomain property value. The domain of the external tenant.
func (m *CrossTenantAccess) SetResourceTenantPrimaryDomain(value *string)() {
    err := m.GetBackingStore().Set("resourceTenantPrimaryDomain", value)
    if err != nil {
        panic(err)
    }
}
// SetUsageStatus sets the usageStatus property value. The usageStatus property
func (m *CrossTenantAccess) SetUsageStatus(value *UsageStatus)() {
    err := m.GetBackingStore().Set("usageStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetUserCount sets the userCount property value. The number of users that accessed the external tenant.
func (m *CrossTenantAccess) SetUserCount(value *int64)() {
    err := m.GetBackingStore().Set("userCount", value)
    if err != nil {
        panic(err)
    }
}
type CrossTenantAccessable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDeviceCount()(*int64)
    GetLastAccessDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetResourceTenantId()(*string)
    GetResourceTenantName()(*string)
    GetResourceTenantPrimaryDomain()(*string)
    GetUsageStatus()(*UsageStatus)
    GetUserCount()(*int64)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDeviceCount(value *int64)()
    SetLastAccessDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetResourceTenantId(value *string)()
    SetResourceTenantName(value *string)()
    SetResourceTenantPrimaryDomain(value *string)()
    SetUsageStatus(value *UsageStatus)()
    SetUserCount(value *int64)()
}
