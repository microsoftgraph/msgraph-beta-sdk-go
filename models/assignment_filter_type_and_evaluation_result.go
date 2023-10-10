package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AssignmentFilterTypeAndEvaluationResult represents the filter type and evalaution result of the filter.
type AssignmentFilterTypeAndEvaluationResult struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAssignmentFilterTypeAndEvaluationResult instantiates a new assignmentFilterTypeAndEvaluationResult and sets the default values.
func NewAssignmentFilterTypeAndEvaluationResult()(*AssignmentFilterTypeAndEvaluationResult) {
    m := &AssignmentFilterTypeAndEvaluationResult{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAssignmentFilterTypeAndEvaluationResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignmentFilterTypeAndEvaluationResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentFilterTypeAndEvaluationResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterTypeAndEvaluationResult) GetAdditionalData()(map[string]any) {
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
// GetAssignmentFilterType gets the assignmentFilterType property value. Represents type of the assignment filter.
func (m *AssignmentFilterTypeAndEvaluationResult) GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType) {
    val, err := m.GetBackingStore().Get("assignmentFilterType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceAndAppManagementAssignmentFilterType)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *AssignmentFilterTypeAndEvaluationResult) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEvaluationResult gets the evaluationResult property value. Supported evaluation results for filter.
func (m *AssignmentFilterTypeAndEvaluationResult) GetEvaluationResult()(*AssignmentFilterEvaluationResult) {
    val, err := m.GetBackingStore().Get("evaluationResult")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AssignmentFilterEvaluationResult)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFilterTypeAndEvaluationResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentFilterType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAndAppManagementAssignmentFilterType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterType(val.(*DeviceAndAppManagementAssignmentFilterType))
        }
        return nil
    }
    res["evaluationResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAssignmentFilterEvaluationResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluationResult(val.(*AssignmentFilterEvaluationResult))
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AssignmentFilterTypeAndEvaluationResult) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AssignmentFilterTypeAndEvaluationResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignmentFilterType() != nil {
        cast := (*m.GetAssignmentFilterType()).String()
        err := writer.WriteStringValue("assignmentFilterType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvaluationResult() != nil {
        cast := (*m.GetEvaluationResult()).String()
        err := writer.WriteStringValue("evaluationResult", &cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterTypeAndEvaluationResult) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentFilterType sets the assignmentFilterType property value. Represents type of the assignment filter.
func (m *AssignmentFilterTypeAndEvaluationResult) SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)() {
    err := m.GetBackingStore().Set("assignmentFilterType", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AssignmentFilterTypeAndEvaluationResult) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEvaluationResult sets the evaluationResult property value. Supported evaluation results for filter.
func (m *AssignmentFilterTypeAndEvaluationResult) SetEvaluationResult(value *AssignmentFilterEvaluationResult)() {
    err := m.GetBackingStore().Set("evaluationResult", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AssignmentFilterTypeAndEvaluationResult) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// AssignmentFilterTypeAndEvaluationResultable 
type AssignmentFilterTypeAndEvaluationResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEvaluationResult()(*AssignmentFilterEvaluationResult)
    GetOdataType()(*string)
    SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEvaluationResult(value *AssignmentFilterEvaluationResult)()
    SetOdataType(value *string)()
}
