package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type CloudPcBulkActionSummary struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCloudPcBulkActionSummary instantiates a new CloudPcBulkActionSummary and sets the default values.
func NewCloudPcBulkActionSummary()(*CloudPcBulkActionSummary) {
    m := &CloudPcBulkActionSummary{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudPcBulkActionSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcBulkActionSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcBulkActionSummary(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CloudPcBulkActionSummary) GetAdditionalData()(map[string]any) {
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
func (m *CloudPcBulkActionSummary) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFailedCount gets the failedCount property value. The number of Cloud PCs where the action failed.
// returns a *int32 when successful
func (m *CloudPcBulkActionSummary) GetFailedCount()(*int32) {
    val, err := m.GetBackingStore().Get("failedCount")
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
func (m *CloudPcBulkActionSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["failedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedCount(val)
        }
        return nil
    }
    res["inProgressCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInProgressCount(val)
        }
        return nil
    }
    res["notSupportedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotSupportedCount(val)
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
    res["pendingCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingCount(val)
        }
        return nil
    }
    res["successfulCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulCount(val)
        }
        return nil
    }
    return res
}
// GetInProgressCount gets the inProgressCount property value. The number of Cloud PCs where the action is in progress.
// returns a *int32 when successful
func (m *CloudPcBulkActionSummary) GetInProgressCount()(*int32) {
    val, err := m.GetBackingStore().Get("inProgressCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNotSupportedCount gets the notSupportedCount property value. The number of Cloud PCs where the action isn't supported.
// returns a *int32 when successful
func (m *CloudPcBulkActionSummary) GetNotSupportedCount()(*int32) {
    val, err := m.GetBackingStore().Get("notSupportedCount")
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
func (m *CloudPcBulkActionSummary) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPendingCount gets the pendingCount property value. The number of Cloud PCs where the action is pending.
// returns a *int32 when successful
func (m *CloudPcBulkActionSummary) GetPendingCount()(*int32) {
    val, err := m.GetBackingStore().Get("pendingCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSuccessfulCount gets the successfulCount property value. The number of Cloud PCs where the action is successful.
// returns a *int32 when successful
func (m *CloudPcBulkActionSummary) GetSuccessfulCount()(*int32) {
    val, err := m.GetBackingStore().Get("successfulCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcBulkActionSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("failedCount", m.GetFailedCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("inProgressCount", m.GetInProgressCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("notSupportedCount", m.GetNotSupportedCount())
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
        err := writer.WriteInt32Value("pendingCount", m.GetPendingCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("successfulCount", m.GetSuccessfulCount())
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
func (m *CloudPcBulkActionSummary) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CloudPcBulkActionSummary) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFailedCount sets the failedCount property value. The number of Cloud PCs where the action failed.
func (m *CloudPcBulkActionSummary) SetFailedCount(value *int32)() {
    err := m.GetBackingStore().Set("failedCount", value)
    if err != nil {
        panic(err)
    }
}
// SetInProgressCount sets the inProgressCount property value. The number of Cloud PCs where the action is in progress.
func (m *CloudPcBulkActionSummary) SetInProgressCount(value *int32)() {
    err := m.GetBackingStore().Set("inProgressCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNotSupportedCount sets the notSupportedCount property value. The number of Cloud PCs where the action isn't supported.
func (m *CloudPcBulkActionSummary) SetNotSupportedCount(value *int32)() {
    err := m.GetBackingStore().Set("notSupportedCount", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcBulkActionSummary) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPendingCount sets the pendingCount property value. The number of Cloud PCs where the action is pending.
func (m *CloudPcBulkActionSummary) SetPendingCount(value *int32)() {
    err := m.GetBackingStore().Set("pendingCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessfulCount sets the successfulCount property value. The number of Cloud PCs where the action is successful.
func (m *CloudPcBulkActionSummary) SetSuccessfulCount(value *int32)() {
    err := m.GetBackingStore().Set("successfulCount", value)
    if err != nil {
        panic(err)
    }
}
type CloudPcBulkActionSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFailedCount()(*int32)
    GetInProgressCount()(*int32)
    GetNotSupportedCount()(*int32)
    GetOdataType()(*string)
    GetPendingCount()(*int32)
    GetSuccessfulCount()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFailedCount(value *int32)()
    SetInProgressCount(value *int32)()
    SetNotSupportedCount(value *int32)()
    SetOdataType(value *string)()
    SetPendingCount(value *int32)()
    SetSuccessfulCount(value *int32)()
}
