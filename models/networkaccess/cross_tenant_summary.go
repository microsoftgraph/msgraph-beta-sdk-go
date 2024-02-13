package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type CrossTenantSummary struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCrossTenantSummary instantiates a new CrossTenantSummary and sets the default values.
func NewCrossTenantSummary()(*CrossTenantSummary) {
    m := &CrossTenantSummary{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCrossTenantSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCrossTenantSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantSummary(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CrossTenantSummary) GetAdditionalData()(map[string]any) {
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
// GetAuthTransactionCount gets the authTransactionCount property value. The total number of authentication sessions between startDateTime and endDateTime.
// returns a *int32 when successful
func (m *CrossTenantSummary) GetAuthTransactionCount()(*int32) {
    val, err := m.GetBackingStore().Get("authTransactionCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *CrossTenantSummary) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDeviceCount gets the deviceCount property value. The number of unique devices that performed cross-tenant access.
// returns a *int32 when successful
func (m *CrossTenantSummary) GetDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("deviceCount")
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
func (m *CrossTenantSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authTransactionCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthTransactionCount(val)
        }
        return nil
    }
    res["deviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
        }
        return nil
    }
    res["newTenantCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewTenantCount(val)
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
    res["rarelyUsedTenantCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRarelyUsedTenantCount(val)
        }
        return nil
    }
    res["tenantCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantCount(val)
        }
        return nil
    }
    res["userCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
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
// GetNewTenantCount gets the newTenantCount property value. The number of unique tenants that were accessed between endDateTime and discoveryPivotDateTime, but weren't accessed between discoveryPivotDateTime and startDateTime.
// returns a *int32 when successful
func (m *CrossTenantSummary) GetNewTenantCount()(*int32) {
    val, err := m.GetBackingStore().Get("newTenantCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CrossTenantSummary) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRarelyUsedTenantCount gets the rarelyUsedTenantCount property value. The rarelyUsedTenantCount property
// returns a *int32 when successful
func (m *CrossTenantSummary) GetRarelyUsedTenantCount()(*int32) {
    val, err := m.GetBackingStore().Get("rarelyUsedTenantCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTenantCount gets the tenantCount property value. The number of unique tenants that were accessed, not including the device's tenant.
// returns a *int32 when successful
func (m *CrossTenantSummary) GetTenantCount()(*int32) {
    val, err := m.GetBackingStore().Get("tenantCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUserCount gets the userCount property value. The number of unique users that performed cross-tenant access.
// returns a *int32 when successful
func (m *CrossTenantSummary) GetUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("userCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CrossTenantSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("authTransactionCount", m.GetAuthTransactionCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("deviceCount", m.GetDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("newTenantCount", m.GetNewTenantCount())
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
        err := writer.WriteInt32Value("rarelyUsedTenantCount", m.GetRarelyUsedTenantCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("tenantCount", m.GetTenantCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("userCount", m.GetUserCount())
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
func (m *CrossTenantSummary) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthTransactionCount sets the authTransactionCount property value. The total number of authentication sessions between startDateTime and endDateTime.
func (m *CrossTenantSummary) SetAuthTransactionCount(value *int32)() {
    err := m.GetBackingStore().Set("authTransactionCount", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CrossTenantSummary) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDeviceCount sets the deviceCount property value. The number of unique devices that performed cross-tenant access.
func (m *CrossTenantSummary) SetDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("deviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNewTenantCount sets the newTenantCount property value. The number of unique tenants that were accessed between endDateTime and discoveryPivotDateTime, but weren't accessed between discoveryPivotDateTime and startDateTime.
func (m *CrossTenantSummary) SetNewTenantCount(value *int32)() {
    err := m.GetBackingStore().Set("newTenantCount", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CrossTenantSummary) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRarelyUsedTenantCount sets the rarelyUsedTenantCount property value. The rarelyUsedTenantCount property
func (m *CrossTenantSummary) SetRarelyUsedTenantCount(value *int32)() {
    err := m.GetBackingStore().Set("rarelyUsedTenantCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantCount sets the tenantCount property value. The number of unique tenants that were accessed, not including the device's tenant.
func (m *CrossTenantSummary) SetTenantCount(value *int32)() {
    err := m.GetBackingStore().Set("tenantCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUserCount sets the userCount property value. The number of unique users that performed cross-tenant access.
func (m *CrossTenantSummary) SetUserCount(value *int32)() {
    err := m.GetBackingStore().Set("userCount", value)
    if err != nil {
        panic(err)
    }
}
type CrossTenantSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthTransactionCount()(*int32)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDeviceCount()(*int32)
    GetNewTenantCount()(*int32)
    GetOdataType()(*string)
    GetRarelyUsedTenantCount()(*int32)
    GetTenantCount()(*int32)
    GetUserCount()(*int32)
    SetAuthTransactionCount(value *int32)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDeviceCount(value *int32)()
    SetNewTenantCount(value *int32)()
    SetOdataType(value *string)()
    SetRarelyUsedTenantCount(value *int32)()
    SetTenantCount(value *int32)()
    SetUserCount(value *int32)()
}
