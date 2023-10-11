package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// BulkManagedDeviceActionResult 
type BulkManagedDeviceActionResult struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewBulkManagedDeviceActionResult instantiates a new bulkManagedDeviceActionResult and sets the default values.
func NewBulkManagedDeviceActionResult()(*BulkManagedDeviceActionResult) {
    m := &BulkManagedDeviceActionResult{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBulkManagedDeviceActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBulkManagedDeviceActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBulkManagedDeviceActionResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BulkManagedDeviceActionResult) GetAdditionalData()(map[string]any) {
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
func (m *BulkManagedDeviceActionResult) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFailedDeviceIds gets the failedDeviceIds property value. Failed devices
func (m *BulkManagedDeviceActionResult) GetFailedDeviceIds()([]string) {
    val, err := m.GetBackingStore().Get("failedDeviceIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BulkManagedDeviceActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["failedDeviceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetFailedDeviceIds(res)
        }
        return nil
    }
    res["notFoundDeviceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetNotFoundDeviceIds(res)
        }
        return nil
    }
    res["notSupportedDeviceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetNotSupportedDeviceIds(res)
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
    res["successfulDeviceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSuccessfulDeviceIds(res)
        }
        return nil
    }
    return res
}
// GetNotFoundDeviceIds gets the notFoundDeviceIds property value. Not found devices
func (m *BulkManagedDeviceActionResult) GetNotFoundDeviceIds()([]string) {
    val, err := m.GetBackingStore().Get("notFoundDeviceIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetNotSupportedDeviceIds gets the notSupportedDeviceIds property value. Not supported devices
func (m *BulkManagedDeviceActionResult) GetNotSupportedDeviceIds()([]string) {
    val, err := m.GetBackingStore().Get("notSupportedDeviceIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BulkManagedDeviceActionResult) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSuccessfulDeviceIds gets the successfulDeviceIds property value. Successful devices
func (m *BulkManagedDeviceActionResult) GetSuccessfulDeviceIds()([]string) {
    val, err := m.GetBackingStore().Get("successfulDeviceIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *BulkManagedDeviceActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFailedDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("failedDeviceIds", m.GetFailedDeviceIds())
        if err != nil {
            return err
        }
    }
    if m.GetNotFoundDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("notFoundDeviceIds", m.GetNotFoundDeviceIds())
        if err != nil {
            return err
        }
    }
    if m.GetNotSupportedDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("notSupportedDeviceIds", m.GetNotSupportedDeviceIds())
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
    if m.GetSuccessfulDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("successfulDeviceIds", m.GetSuccessfulDeviceIds())
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
func (m *BulkManagedDeviceActionResult) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *BulkManagedDeviceActionResult) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFailedDeviceIds sets the failedDeviceIds property value. Failed devices
func (m *BulkManagedDeviceActionResult) SetFailedDeviceIds(value []string)() {
    err := m.GetBackingStore().Set("failedDeviceIds", value)
    if err != nil {
        panic(err)
    }
}
// SetNotFoundDeviceIds sets the notFoundDeviceIds property value. Not found devices
func (m *BulkManagedDeviceActionResult) SetNotFoundDeviceIds(value []string)() {
    err := m.GetBackingStore().Set("notFoundDeviceIds", value)
    if err != nil {
        panic(err)
    }
}
// SetNotSupportedDeviceIds sets the notSupportedDeviceIds property value. Not supported devices
func (m *BulkManagedDeviceActionResult) SetNotSupportedDeviceIds(value []string)() {
    err := m.GetBackingStore().Set("notSupportedDeviceIds", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BulkManagedDeviceActionResult) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessfulDeviceIds sets the successfulDeviceIds property value. Successful devices
func (m *BulkManagedDeviceActionResult) SetSuccessfulDeviceIds(value []string)() {
    err := m.GetBackingStore().Set("successfulDeviceIds", value)
    if err != nil {
        panic(err)
    }
}
// BulkManagedDeviceActionResultable 
type BulkManagedDeviceActionResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFailedDeviceIds()([]string)
    GetNotFoundDeviceIds()([]string)
    GetNotSupportedDeviceIds()([]string)
    GetOdataType()(*string)
    GetSuccessfulDeviceIds()([]string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFailedDeviceIds(value []string)()
    SetNotFoundDeviceIds(value []string)()
    SetNotSupportedDeviceIds(value []string)()
    SetOdataType(value *string)()
    SetSuccessfulDeviceIds(value []string)()
}
