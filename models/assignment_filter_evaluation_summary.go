package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AssignmentFilterEvaluationSummary represent result summary for assignment filter evaluation
type AssignmentFilterEvaluationSummary struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAssignmentFilterEvaluationSummary instantiates a new assignmentFilterEvaluationSummary and sets the default values.
func NewAssignmentFilterEvaluationSummary()(*AssignmentFilterEvaluationSummary) {
    m := &AssignmentFilterEvaluationSummary{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAssignmentFilterEvaluationSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignmentFilterEvaluationSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentFilterEvaluationSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterEvaluationSummary) GetAdditionalData()(map[string]any) {
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
// GetAssignmentFilterDisplayName gets the assignmentFilterDisplayName property value. The admin defined name for assignment filter.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("assignmentFilterDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssignmentFilterId gets the assignmentFilterId property value. Unique identifier for the assignment filter object
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterId()(*string) {
    val, err := m.GetBackingStore().Get("assignmentFilterId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssignmentFilterLastModifiedDateTime gets the assignmentFilterLastModifiedDateTime property value. The time the assignment filter was last modified.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("assignmentFilterLastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetAssignmentFilterPlatform gets the assignmentFilterPlatform property value. Supported platform types.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterPlatform()(*DevicePlatformType) {
    val, err := m.GetBackingStore().Get("assignmentFilterPlatform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DevicePlatformType)
    }
    return nil
}
// GetAssignmentFilterType gets the assignmentFilterType property value. Represents type of the assignment filter.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType) {
    val, err := m.GetBackingStore().Get("assignmentFilterType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceAndAppManagementAssignmentFilterType)
    }
    return nil
}
// GetAssignmentFilterTypeAndEvaluationResults gets the assignmentFilterTypeAndEvaluationResults property value. A collection of filter types and their corresponding evaluation results.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterTypeAndEvaluationResults()([]AssignmentFilterTypeAndEvaluationResultable) {
    val, err := m.GetBackingStore().Get("assignmentFilterTypeAndEvaluationResults")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AssignmentFilterTypeAndEvaluationResultable)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *AssignmentFilterEvaluationSummary) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEvaluationDateTime gets the evaluationDateTime property value. The time assignment filter was evaluated.
func (m *AssignmentFilterEvaluationSummary) GetEvaluationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("evaluationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetEvaluationResult gets the evaluationResult property value. Supported evaluation results for filter.
func (m *AssignmentFilterEvaluationSummary) GetEvaluationResult()(*AssignmentFilterEvaluationResult) {
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
func (m *AssignmentFilterEvaluationSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentFilterDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterDisplayName(val)
        }
        return nil
    }
    res["assignmentFilterId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterId(val)
        }
        return nil
    }
    res["assignmentFilterLastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterLastModifiedDateTime(val)
        }
        return nil
    }
    res["assignmentFilterPlatform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterPlatform(val.(*DevicePlatformType))
        }
        return nil
    }
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
    res["assignmentFilterTypeAndEvaluationResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignmentFilterTypeAndEvaluationResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignmentFilterTypeAndEvaluationResultable, len(val))
            for i, v := range val {
                res[i] = v.(AssignmentFilterTypeAndEvaluationResultable)
            }
            m.SetAssignmentFilterTypeAndEvaluationResults(res)
        }
        return nil
    }
    res["evaluationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluationDateTime(val)
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
func (m *AssignmentFilterEvaluationSummary) GetOdataType()(*string) {
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
func (m *AssignmentFilterEvaluationSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("assignmentFilterDisplayName", m.GetAssignmentFilterDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("assignmentFilterId", m.GetAssignmentFilterId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("assignmentFilterLastModifiedDateTime", m.GetAssignmentFilterLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentFilterPlatform() != nil {
        cast := (*m.GetAssignmentFilterPlatform()).String()
        err := writer.WriteStringValue("assignmentFilterPlatform", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentFilterType() != nil {
        cast := (*m.GetAssignmentFilterType()).String()
        err := writer.WriteStringValue("assignmentFilterType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentFilterTypeAndEvaluationResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignmentFilterTypeAndEvaluationResults()))
        for i, v := range m.GetAssignmentFilterTypeAndEvaluationResults() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("assignmentFilterTypeAndEvaluationResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("evaluationDateTime", m.GetEvaluationDateTime())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterEvaluationSummary) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentFilterDisplayName sets the assignmentFilterDisplayName property value. The admin defined name for assignment filter.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterDisplayName(value *string)() {
    err := m.GetBackingStore().Set("assignmentFilterDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentFilterId sets the assignmentFilterId property value. Unique identifier for the assignment filter object
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterId(value *string)() {
    err := m.GetBackingStore().Set("assignmentFilterId", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentFilterLastModifiedDateTime sets the assignmentFilterLastModifiedDateTime property value. The time the assignment filter was last modified.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("assignmentFilterLastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentFilterPlatform sets the assignmentFilterPlatform property value. Supported platform types.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterPlatform(value *DevicePlatformType)() {
    err := m.GetBackingStore().Set("assignmentFilterPlatform", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentFilterType sets the assignmentFilterType property value. Represents type of the assignment filter.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)() {
    err := m.GetBackingStore().Set("assignmentFilterType", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentFilterTypeAndEvaluationResults sets the assignmentFilterTypeAndEvaluationResults property value. A collection of filter types and their corresponding evaluation results.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterTypeAndEvaluationResults(value []AssignmentFilterTypeAndEvaluationResultable)() {
    err := m.GetBackingStore().Set("assignmentFilterTypeAndEvaluationResults", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *AssignmentFilterEvaluationSummary) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEvaluationDateTime sets the evaluationDateTime property value. The time assignment filter was evaluated.
func (m *AssignmentFilterEvaluationSummary) SetEvaluationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("evaluationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetEvaluationResult sets the evaluationResult property value. Supported evaluation results for filter.
func (m *AssignmentFilterEvaluationSummary) SetEvaluationResult(value *AssignmentFilterEvaluationResult)() {
    err := m.GetBackingStore().Set("evaluationResult", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AssignmentFilterEvaluationSummary) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// AssignmentFilterEvaluationSummaryable 
type AssignmentFilterEvaluationSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentFilterDisplayName()(*string)
    GetAssignmentFilterId()(*string)
    GetAssignmentFilterLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAssignmentFilterPlatform()(*DevicePlatformType)
    GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType)
    GetAssignmentFilterTypeAndEvaluationResults()([]AssignmentFilterTypeAndEvaluationResultable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEvaluationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEvaluationResult()(*AssignmentFilterEvaluationResult)
    GetOdataType()(*string)
    SetAssignmentFilterDisplayName(value *string)()
    SetAssignmentFilterId(value *string)()
    SetAssignmentFilterLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAssignmentFilterPlatform(value *DevicePlatformType)()
    SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)()
    SetAssignmentFilterTypeAndEvaluationResults(value []AssignmentFilterTypeAndEvaluationResultable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEvaluationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEvaluationResult(value *AssignmentFilterEvaluationResult)()
    SetOdataType(value *string)()
}
