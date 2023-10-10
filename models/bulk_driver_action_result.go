package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// BulkDriverActionResult a complex type to represent the result of bulk driver action.
type BulkDriverActionResult struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewBulkDriverActionResult instantiates a new bulkDriverActionResult and sets the default values.
func NewBulkDriverActionResult()(*BulkDriverActionResult) {
    m := &BulkDriverActionResult{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBulkDriverActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBulkDriverActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBulkDriverActionResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BulkDriverActionResult) GetAdditionalData()(map[string]any) {
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
func (m *BulkDriverActionResult) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFailedDriverIds gets the failedDriverIds property value. List of driver Ids where the action is failed.
func (m *BulkDriverActionResult) GetFailedDriverIds()([]string) {
    val, err := m.GetBackingStore().Get("failedDriverIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BulkDriverActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["failedDriverIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetFailedDriverIds(res)
        }
        return nil
    }
    res["notFoundDriverIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetNotFoundDriverIds(res)
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
    res["successfulDriverIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSuccessfulDriverIds(res)
        }
        return nil
    }
    return res
}
// GetNotFoundDriverIds gets the notFoundDriverIds property value. List of driver Ids that are not found.
func (m *BulkDriverActionResult) GetNotFoundDriverIds()([]string) {
    val, err := m.GetBackingStore().Get("notFoundDriverIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BulkDriverActionResult) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSuccessfulDriverIds gets the successfulDriverIds property value. List of driver Ids where the action is successful.
func (m *BulkDriverActionResult) GetSuccessfulDriverIds()([]string) {
    val, err := m.GetBackingStore().Get("successfulDriverIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *BulkDriverActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFailedDriverIds() != nil {
        err := writer.WriteCollectionOfStringValues("failedDriverIds", m.GetFailedDriverIds())
        if err != nil {
            return err
        }
    }
    if m.GetNotFoundDriverIds() != nil {
        err := writer.WriteCollectionOfStringValues("notFoundDriverIds", m.GetNotFoundDriverIds())
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
    if m.GetSuccessfulDriverIds() != nil {
        err := writer.WriteCollectionOfStringValues("successfulDriverIds", m.GetSuccessfulDriverIds())
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
func (m *BulkDriverActionResult) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *BulkDriverActionResult) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFailedDriverIds sets the failedDriverIds property value. List of driver Ids where the action is failed.
func (m *BulkDriverActionResult) SetFailedDriverIds(value []string)() {
    err := m.GetBackingStore().Set("failedDriverIds", value)
    if err != nil {
        panic(err)
    }
}
// SetNotFoundDriverIds sets the notFoundDriverIds property value. List of driver Ids that are not found.
func (m *BulkDriverActionResult) SetNotFoundDriverIds(value []string)() {
    err := m.GetBackingStore().Set("notFoundDriverIds", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BulkDriverActionResult) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessfulDriverIds sets the successfulDriverIds property value. List of driver Ids where the action is successful.
func (m *BulkDriverActionResult) SetSuccessfulDriverIds(value []string)() {
    err := m.GetBackingStore().Set("successfulDriverIds", value)
    if err != nil {
        panic(err)
    }
}
// BulkDriverActionResultable 
type BulkDriverActionResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFailedDriverIds()([]string)
    GetNotFoundDriverIds()([]string)
    GetOdataType()(*string)
    GetSuccessfulDriverIds()([]string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFailedDriverIds(value []string)()
    SetNotFoundDriverIds(value []string)()
    SetOdataType(value *string)()
    SetSuccessfulDriverIds(value []string)()
}
