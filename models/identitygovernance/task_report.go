package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// TaskReport provides operations to manage the collection of accessReviewDecision entities.
type TaskReport struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The date time that the associated run completed. Value is null if the run has not completed.
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of users in the run execution for which the associated task failed.
    failedUsersCount *int32
    // The date and time that the task report was last updated.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The processingStatus property
    processingStatus *LifecycleWorkflowProcessingStatus
    // The unique identifier of the associated run.
    runId *string
    // The date time that the associated run started. Value is null if the run has not started.
    startedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of users in the run execution for which the associated task succeeded.
    successfulUsersCount *int32
    // The task property
    task Taskable
    // The taskDefinition property
    taskDefinition TaskDefinitionable
    // The related lifecycle workflow taskProcessingResults.
    taskProcessingResults []TaskProcessingResultable
    // The total number of users in the run execution for which the associated task was scheduled to execute.
    totalUsersCount *int32
    // The number of users in the run execution for which the associated task is queued, in progress, or canceled.
    unprocessedUsersCount *int32
}
// NewTaskReport instantiates a new taskReport and sets the default values.
func NewTaskReport()(*TaskReport) {
    m := &TaskReport{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.taskReport";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTaskReportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTaskReportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTaskReport(), nil
}
// GetCompletedDateTime gets the completedDateTime property value. The date time that the associated run completed. Value is null if the run has not completed.
func (m *TaskReport) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDateTime
}
// GetFailedUsersCount gets the failedUsersCount property value. The number of users in the run execution for which the associated task failed.
func (m *TaskReport) GetFailedUsersCount()(*int32) {
    return m.failedUsersCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TaskReport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["completedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCompletedDateTime)
    res["failedUsersCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetFailedUsersCount)
    res["lastUpdatedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastUpdatedDateTime)
    res["processingStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseLifecycleWorkflowProcessingStatus , m.SetProcessingStatus)
    res["runId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRunId)
    res["startedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetStartedDateTime)
    res["successfulUsersCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSuccessfulUsersCount)
    res["task"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTaskFromDiscriminatorValue , m.SetTask)
    res["taskDefinition"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTaskDefinitionFromDiscriminatorValue , m.SetTaskDefinition)
    res["taskProcessingResults"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTaskProcessingResultFromDiscriminatorValue , m.SetTaskProcessingResults)
    res["totalUsersCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalUsersCount)
    res["unprocessedUsersCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetUnprocessedUsersCount)
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The date and time that the task report was last updated.
func (m *TaskReport) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetProcessingStatus gets the processingStatus property value. The processingStatus property
func (m *TaskReport) GetProcessingStatus()(*LifecycleWorkflowProcessingStatus) {
    return m.processingStatus
}
// GetRunId gets the runId property value. The unique identifier of the associated run.
func (m *TaskReport) GetRunId()(*string) {
    return m.runId
}
// GetStartedDateTime gets the startedDateTime property value. The date time that the associated run started. Value is null if the run has not started.
func (m *TaskReport) GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startedDateTime
}
// GetSuccessfulUsersCount gets the successfulUsersCount property value. The number of users in the run execution for which the associated task succeeded.
func (m *TaskReport) GetSuccessfulUsersCount()(*int32) {
    return m.successfulUsersCount
}
// GetTask gets the task property value. The task property
func (m *TaskReport) GetTask()(Taskable) {
    return m.task
}
// GetTaskDefinition gets the taskDefinition property value. The taskDefinition property
func (m *TaskReport) GetTaskDefinition()(TaskDefinitionable) {
    return m.taskDefinition
}
// GetTaskProcessingResults gets the taskProcessingResults property value. The related lifecycle workflow taskProcessingResults.
func (m *TaskReport) GetTaskProcessingResults()([]TaskProcessingResultable) {
    return m.taskProcessingResults
}
// GetTotalUsersCount gets the totalUsersCount property value. The total number of users in the run execution for which the associated task was scheduled to execute.
func (m *TaskReport) GetTotalUsersCount()(*int32) {
    return m.totalUsersCount
}
// GetUnprocessedUsersCount gets the unprocessedUsersCount property value. The number of users in the run execution for which the associated task is queued, in progress, or canceled.
func (m *TaskReport) GetUnprocessedUsersCount()(*int32) {
    return m.unprocessedUsersCount
}
// Serialize serializes information the current object
func (m *TaskReport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedUsersCount", m.GetFailedUsersCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetProcessingStatus() != nil {
        cast := (*m.GetProcessingStatus()).String()
        err = writer.WriteStringValue("processingStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("runId", m.GetRunId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startedDateTime", m.GetStartedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("successfulUsersCount", m.GetSuccessfulUsersCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("task", m.GetTask())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("taskDefinition", m.GetTaskDefinition())
        if err != nil {
            return err
        }
    }
    if m.GetTaskProcessingResults() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTaskProcessingResults())
        err = writer.WriteCollectionOfObjectValues("taskProcessingResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalUsersCount", m.GetTotalUsersCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unprocessedUsersCount", m.GetUnprocessedUsersCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletedDateTime sets the completedDateTime property value. The date time that the associated run completed. Value is null if the run has not completed.
func (m *TaskReport) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// SetFailedUsersCount sets the failedUsersCount property value. The number of users in the run execution for which the associated task failed.
func (m *TaskReport) SetFailedUsersCount(value *int32)() {
    m.failedUsersCount = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The date and time that the task report was last updated.
func (m *TaskReport) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetProcessingStatus sets the processingStatus property value. The processingStatus property
func (m *TaskReport) SetProcessingStatus(value *LifecycleWorkflowProcessingStatus)() {
    m.processingStatus = value
}
// SetRunId sets the runId property value. The unique identifier of the associated run.
func (m *TaskReport) SetRunId(value *string)() {
    m.runId = value
}
// SetStartedDateTime sets the startedDateTime property value. The date time that the associated run started. Value is null if the run has not started.
func (m *TaskReport) SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startedDateTime = value
}
// SetSuccessfulUsersCount sets the successfulUsersCount property value. The number of users in the run execution for which the associated task succeeded.
func (m *TaskReport) SetSuccessfulUsersCount(value *int32)() {
    m.successfulUsersCount = value
}
// SetTask sets the task property value. The task property
func (m *TaskReport) SetTask(value Taskable)() {
    m.task = value
}
// SetTaskDefinition sets the taskDefinition property value. The taskDefinition property
func (m *TaskReport) SetTaskDefinition(value TaskDefinitionable)() {
    m.taskDefinition = value
}
// SetTaskProcessingResults sets the taskProcessingResults property value. The related lifecycle workflow taskProcessingResults.
func (m *TaskReport) SetTaskProcessingResults(value []TaskProcessingResultable)() {
    m.taskProcessingResults = value
}
// SetTotalUsersCount sets the totalUsersCount property value. The total number of users in the run execution for which the associated task was scheduled to execute.
func (m *TaskReport) SetTotalUsersCount(value *int32)() {
    m.totalUsersCount = value
}
// SetUnprocessedUsersCount sets the unprocessedUsersCount property value. The number of users in the run execution for which the associated task is queued, in progress, or canceled.
func (m *TaskReport) SetUnprocessedUsersCount(value *int32)() {
    m.unprocessedUsersCount = value
}
