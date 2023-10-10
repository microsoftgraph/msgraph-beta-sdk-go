package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// EvaluateLabelJobResult 
type EvaluateLabelJobResult struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewEvaluateLabelJobResult instantiates a new evaluateLabelJobResult and sets the default values.
func NewEvaluateLabelJobResult()(*EvaluateLabelJobResult) {
    m := &EvaluateLabelJobResult{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEvaluateLabelJobResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEvaluateLabelJobResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEvaluateLabelJobResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateLabelJobResult) GetAdditionalData()(map[string]any) {
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
func (m *EvaluateLabelJobResult) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EvaluateLabelJobResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["responsiblePolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResponsiblePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponsiblePolicy(val.(ResponsiblePolicyable))
        }
        return nil
    }
    res["responsibleSensitiveTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateResponsibleSensitiveTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResponsibleSensitiveTypeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ResponsibleSensitiveTypeable)
                }
            }
            m.SetResponsibleSensitiveTypes(res)
        }
        return nil
    }
    res["sensitivityLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMatchingLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivityLabel(val.(MatchingLabelable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EvaluateLabelJobResult) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResponsiblePolicy gets the responsiblePolicy property value. The responsiblePolicy property
func (m *EvaluateLabelJobResult) GetResponsiblePolicy()(ResponsiblePolicyable) {
    val, err := m.GetBackingStore().Get("responsiblePolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ResponsiblePolicyable)
    }
    return nil
}
// GetResponsibleSensitiveTypes gets the responsibleSensitiveTypes property value. The responsibleSensitiveTypes property
func (m *EvaluateLabelJobResult) GetResponsibleSensitiveTypes()([]ResponsibleSensitiveTypeable) {
    val, err := m.GetBackingStore().Get("responsibleSensitiveTypes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ResponsibleSensitiveTypeable)
    }
    return nil
}
// GetSensitivityLabel gets the sensitivityLabel property value. The sensitivityLabel property
func (m *EvaluateLabelJobResult) GetSensitivityLabel()(MatchingLabelable) {
    val, err := m.GetBackingStore().Get("sensitivityLabel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MatchingLabelable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EvaluateLabelJobResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("responsiblePolicy", m.GetResponsiblePolicy())
        if err != nil {
            return err
        }
    }
    if m.GetResponsibleSensitiveTypes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResponsibleSensitiveTypes()))
        for i, v := range m.GetResponsibleSensitiveTypes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("responsibleSensitiveTypes", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sensitivityLabel", m.GetSensitivityLabel())
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
func (m *EvaluateLabelJobResult) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *EvaluateLabelJobResult) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EvaluateLabelJobResult) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetResponsiblePolicy sets the responsiblePolicy property value. The responsiblePolicy property
func (m *EvaluateLabelJobResult) SetResponsiblePolicy(value ResponsiblePolicyable)() {
    err := m.GetBackingStore().Set("responsiblePolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetResponsibleSensitiveTypes sets the responsibleSensitiveTypes property value. The responsibleSensitiveTypes property
func (m *EvaluateLabelJobResult) SetResponsibleSensitiveTypes(value []ResponsibleSensitiveTypeable)() {
    err := m.GetBackingStore().Set("responsibleSensitiveTypes", value)
    if err != nil {
        panic(err)
    }
}
// SetSensitivityLabel sets the sensitivityLabel property value. The sensitivityLabel property
func (m *EvaluateLabelJobResult) SetSensitivityLabel(value MatchingLabelable)() {
    err := m.GetBackingStore().Set("sensitivityLabel", value)
    if err != nil {
        panic(err)
    }
}
// EvaluateLabelJobResultable 
type EvaluateLabelJobResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetResponsiblePolicy()(ResponsiblePolicyable)
    GetResponsibleSensitiveTypes()([]ResponsibleSensitiveTypeable)
    GetSensitivityLabel()(MatchingLabelable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetResponsiblePolicy(value ResponsiblePolicyable)()
    SetResponsibleSensitiveTypes(value []ResponsibleSensitiveTypeable)()
    SetSensitivityLabel(value MatchingLabelable)()
}
