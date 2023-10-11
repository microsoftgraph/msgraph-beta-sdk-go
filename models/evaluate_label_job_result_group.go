package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// EvaluateLabelJobResultGroup 
type EvaluateLabelJobResultGroup struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewEvaluateLabelJobResultGroup instantiates a new evaluateLabelJobResultGroup and sets the default values.
func NewEvaluateLabelJobResultGroup()(*EvaluateLabelJobResultGroup) {
    m := &EvaluateLabelJobResultGroup{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEvaluateLabelJobResultGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEvaluateLabelJobResultGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEvaluateLabelJobResultGroup(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateLabelJobResultGroup) GetAdditionalData()(map[string]any) {
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
// GetAutomatic gets the automatic property value. The automatic property
func (m *EvaluateLabelJobResultGroup) GetAutomatic()(EvaluateLabelJobResultable) {
    val, err := m.GetBackingStore().Get("automatic")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EvaluateLabelJobResultable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *EvaluateLabelJobResultGroup) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EvaluateLabelJobResultGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["automatic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEvaluateLabelJobResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomatic(val.(EvaluateLabelJobResultable))
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
    res["recommended"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEvaluateLabelJobResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommended(val.(EvaluateLabelJobResultable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EvaluateLabelJobResultGroup) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRecommended gets the recommended property value. The recommended property
func (m *EvaluateLabelJobResultGroup) GetRecommended()(EvaluateLabelJobResultable) {
    val, err := m.GetBackingStore().Get("recommended")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EvaluateLabelJobResultable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EvaluateLabelJobResultGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("automatic", m.GetAutomatic())
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
        err := writer.WriteObjectValue("recommended", m.GetRecommended())
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
func (m *EvaluateLabelJobResultGroup) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAutomatic sets the automatic property value. The automatic property
func (m *EvaluateLabelJobResultGroup) SetAutomatic(value EvaluateLabelJobResultable)() {
    err := m.GetBackingStore().Set("automatic", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *EvaluateLabelJobResultGroup) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EvaluateLabelJobResultGroup) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRecommended sets the recommended property value. The recommended property
func (m *EvaluateLabelJobResultGroup) SetRecommended(value EvaluateLabelJobResultable)() {
    err := m.GetBackingStore().Set("recommended", value)
    if err != nil {
        panic(err)
    }
}
// EvaluateLabelJobResultGroupable 
type EvaluateLabelJobResultGroupable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutomatic()(EvaluateLabelJobResultable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetRecommended()(EvaluateLabelJobResultable)
    SetAutomatic(value EvaluateLabelJobResultable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetRecommended(value EvaluateLabelJobResultable)()
}
