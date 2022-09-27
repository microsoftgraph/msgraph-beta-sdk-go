package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TaskReportSummary 
type TaskReportSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The number of failed tasks in a report.
    failedTasks *int32
    // The OdataType property
    odataType *string
    // The total number of successful tasks in a report.
    successfulTasks *int32
    // The total number of tasks in a report.
    totalTasks *int32
    // The number of unprocessed tasks in a report.
    unprocessedTasks *int32
}
// NewTaskReportSummary instantiates a new taskReportSummary and sets the default values.
func NewTaskReportSummary()(*TaskReportSummary) {
    m := &TaskReportSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.identityGovernance.taskReportSummary";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTaskReportSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTaskReportSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTaskReportSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TaskReportSummary) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFailedTasks gets the failedTasks property value. The number of failed tasks in a report.
func (m *TaskReportSummary) GetFailedTasks()(*int32) {
    return m.failedTasks
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TaskReportSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["failedTasks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetFailedTasks)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["successfulTasks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSuccessfulTasks)
    res["totalTasks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalTasks)
    res["unprocessedTasks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetUnprocessedTasks)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TaskReportSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetSuccessfulTasks gets the successfulTasks property value. The total number of successful tasks in a report.
func (m *TaskReportSummary) GetSuccessfulTasks()(*int32) {
    return m.successfulTasks
}
// GetTotalTasks gets the totalTasks property value. The total number of tasks in a report.
func (m *TaskReportSummary) GetTotalTasks()(*int32) {
    return m.totalTasks
}
// GetUnprocessedTasks gets the unprocessedTasks property value. The number of unprocessed tasks in a report.
func (m *TaskReportSummary) GetUnprocessedTasks()(*int32) {
    return m.unprocessedTasks
}
// Serialize serializes information the current object
func (m *TaskReportSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("failedTasks", m.GetFailedTasks())
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
        err := writer.WriteInt32Value("successfulTasks", m.GetSuccessfulTasks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalTasks", m.GetTotalTasks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unprocessedTasks", m.GetUnprocessedTasks())
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
func (m *TaskReportSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFailedTasks sets the failedTasks property value. The number of failed tasks in a report.
func (m *TaskReportSummary) SetFailedTasks(value *int32)() {
    m.failedTasks = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TaskReportSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSuccessfulTasks sets the successfulTasks property value. The total number of successful tasks in a report.
func (m *TaskReportSummary) SetSuccessfulTasks(value *int32)() {
    m.successfulTasks = value
}
// SetTotalTasks sets the totalTasks property value. The total number of tasks in a report.
func (m *TaskReportSummary) SetTotalTasks(value *int32)() {
    m.totalTasks = value
}
// SetUnprocessedTasks sets the unprocessedTasks property value. The number of unprocessed tasks in a report.
func (m *TaskReportSummary) SetUnprocessedTasks(value *int32)() {
    m.unprocessedTasks = value
}
