package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NoEntitlementsDataCollection 
type NoEntitlementsDataCollection struct {
    EntitlementsDataCollectionInfo
}
// NewNoEntitlementsDataCollection instantiates a new noEntitlementsDataCollection and sets the default values.
func NewNoEntitlementsDataCollection()(*NoEntitlementsDataCollection) {
    m := &NoEntitlementsDataCollection{
        EntitlementsDataCollectionInfo: *NewEntitlementsDataCollectionInfo(),
    }
    odataTypeValue := "#microsoft.graph.noEntitlementsDataCollection"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateNoEntitlementsDataCollectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNoEntitlementsDataCollectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNoEntitlementsDataCollection(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NoEntitlementsDataCollection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EntitlementsDataCollectionInfo.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *NoEntitlementsDataCollection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EntitlementsDataCollectionInfo.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// NoEntitlementsDataCollectionable 
type NoEntitlementsDataCollectionable interface {
    EntitlementsDataCollectionInfoable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
