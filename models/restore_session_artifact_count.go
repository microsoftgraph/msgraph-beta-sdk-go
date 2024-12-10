package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type RestoreSessionArtifactCount struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewRestoreSessionArtifactCount instantiates a new RestoreSessionArtifactCount and sets the default values.
func NewRestoreSessionArtifactCount()(*RestoreSessionArtifactCount) {
    m := &RestoreSessionArtifactCount{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRestoreSessionArtifactCountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRestoreSessionArtifactCountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRestoreSessionArtifactCount(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RestoreSessionArtifactCount) GetAdditionalData()(map[string]any) {
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
func (m *RestoreSessionArtifactCount) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCompleted gets the completed property value. The number of artifacts whose restoration completed.
// returns a *int32 when successful
func (m *RestoreSessionArtifactCount) GetCompleted()(*int32) {
    val, err := m.GetBackingStore().Get("completed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFailed gets the failed property value. The number of artifacts whose restoration failed.
// returns a *int32 when successful
func (m *RestoreSessionArtifactCount) GetFailed()(*int32) {
    val, err := m.GetBackingStore().Get("failed")
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
func (m *RestoreSessionArtifactCount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["completed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompleted(val)
        }
        return nil
    }
    res["failed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailed(val)
        }
        return nil
    }
    res["inProgress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInProgress(val)
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
    res["total"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    return res
}
// GetInProgress gets the inProgress property value. The number of artifacts whose restoration is in progress.
// returns a *int32 when successful
func (m *RestoreSessionArtifactCount) GetInProgress()(*int32) {
    val, err := m.GetBackingStore().Get("inProgress")
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
func (m *RestoreSessionArtifactCount) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTotal gets the total property value. The number of artifacts present in the restore session.
// returns a *int32 when successful
func (m *RestoreSessionArtifactCount) GetTotal()(*int32) {
    val, err := m.GetBackingStore().Get("total")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RestoreSessionArtifactCount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("completed", m.GetCompleted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("failed", m.GetFailed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("inProgress", m.GetInProgress())
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
        err := writer.WriteInt32Value("total", m.GetTotal())
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
func (m *RestoreSessionArtifactCount) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *RestoreSessionArtifactCount) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCompleted sets the completed property value. The number of artifacts whose restoration completed.
func (m *RestoreSessionArtifactCount) SetCompleted(value *int32)() {
    err := m.GetBackingStore().Set("completed", value)
    if err != nil {
        panic(err)
    }
}
// SetFailed sets the failed property value. The number of artifacts whose restoration failed.
func (m *RestoreSessionArtifactCount) SetFailed(value *int32)() {
    err := m.GetBackingStore().Set("failed", value)
    if err != nil {
        panic(err)
    }
}
// SetInProgress sets the inProgress property value. The number of artifacts whose restoration is in progress.
func (m *RestoreSessionArtifactCount) SetInProgress(value *int32)() {
    err := m.GetBackingStore().Set("inProgress", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RestoreSessionArtifactCount) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTotal sets the total property value. The number of artifacts present in the restore session.
func (m *RestoreSessionArtifactCount) SetTotal(value *int32)() {
    err := m.GetBackingStore().Set("total", value)
    if err != nil {
        panic(err)
    }
}
type RestoreSessionArtifactCountable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCompleted()(*int32)
    GetFailed()(*int32)
    GetInProgress()(*int32)
    GetOdataType()(*string)
    GetTotal()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCompleted(value *int32)()
    SetFailed(value *int32)()
    SetInProgress(value *int32)()
    SetOdataType(value *string)()
    SetTotal(value *int32)()
}
