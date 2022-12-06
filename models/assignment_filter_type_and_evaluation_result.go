package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignmentFilterTypeAndEvaluationResult represents the filter type and evalaution result of the filter.
type AssignmentFilterTypeAndEvaluationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Represents type of the assignment filter.
    assignmentFilterType *DeviceAndAppManagementAssignmentFilterType
    // Supported evaluation results for filter.
    evaluationResult *AssignmentFilterEvaluationResult
    // The OdataType property
    odataType *string
}
// NewAssignmentFilterTypeAndEvaluationResult instantiates a new assignmentFilterTypeAndEvaluationResult and sets the default values.
func NewAssignmentFilterTypeAndEvaluationResult()(*AssignmentFilterTypeAndEvaluationResult) {
    m := &AssignmentFilterTypeAndEvaluationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAssignmentFilterTypeAndEvaluationResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignmentFilterTypeAndEvaluationResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentFilterTypeAndEvaluationResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterTypeAndEvaluationResult) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignmentFilterType gets the assignmentFilterType property value. Represents type of the assignment filter.
func (m *AssignmentFilterTypeAndEvaluationResult) GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType) {
    return m.assignmentFilterType
}
// GetEvaluationResult gets the evaluationResult property value. Supported evaluation results for filter.
func (m *AssignmentFilterTypeAndEvaluationResult) GetEvaluationResult()(*AssignmentFilterEvaluationResult) {
    return m.evaluationResult
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFilterTypeAndEvaluationResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentFilterType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceAndAppManagementAssignmentFilterType , m.SetAssignmentFilterType)
    res["evaluationResult"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAssignmentFilterEvaluationResult , m.SetEvaluationResult)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AssignmentFilterTypeAndEvaluationResult) GetOdataType()(*string) {
    return m.odataType
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterTypeAndEvaluationResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignmentFilterType sets the assignmentFilterType property value. Represents type of the assignment filter.
func (m *AssignmentFilterTypeAndEvaluationResult) SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)() {
    m.assignmentFilterType = value
}
// SetEvaluationResult sets the evaluationResult property value. Supported evaluation results for filter.
func (m *AssignmentFilterTypeAndEvaluationResult) SetEvaluationResult(value *AssignmentFilterEvaluationResult)() {
    m.evaluationResult = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AssignmentFilterTypeAndEvaluationResult) SetOdataType(value *string)() {
    m.odataType = value
}
