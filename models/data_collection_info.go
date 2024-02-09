package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DataCollectionInfo struct {
    Entity
}
// NewDataCollectionInfo instantiates a new DataCollectionInfo and sets the default values.
func NewDataCollectionInfo()(*DataCollectionInfo) {
    m := &DataCollectionInfo{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDataCollectionInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDataCollectionInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataCollectionInfo(), nil
}
// GetEntitlements gets the entitlements property value. The entitlements property
// returns a EntitlementsDataCollectionInfoable when successful
func (m *DataCollectionInfo) GetEntitlements()(EntitlementsDataCollectionInfoable) {
    val, err := m.GetBackingStore().Get("entitlements")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EntitlementsDataCollectionInfoable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DataCollectionInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["entitlements"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntitlementsDataCollectionInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntitlements(val.(EntitlementsDataCollectionInfoable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DataCollectionInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("entitlements", m.GetEntitlements())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEntitlements sets the entitlements property value. The entitlements property
func (m *DataCollectionInfo) SetEntitlements(value EntitlementsDataCollectionInfoable)() {
    err := m.GetBackingStore().Set("entitlements", value)
    if err != nil {
        panic(err)
    }
}
type DataCollectionInfoable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEntitlements()(EntitlementsDataCollectionInfoable)
    SetEntitlements(value EntitlementsDataCollectionInfoable)()
}
