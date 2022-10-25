package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Run provides operations to manage the collection of accessReview entities.
type Run struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The date time that the run completed. Value is null if the workflow hasn't completed. Optional.
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of tasks that failed in the run execution. Required.
    failedTasksCount *int32
    // The number of users that failed in the run execution. Required.
    failedUsersCount *int32
    // The datetime that the run was last updated. Optional.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The processingStatus property
    processingStatus *LifecycleWorkflowProcessingStatus
    // The date time that the run is scheduled to be executed for a workflow. Required.
    scheduledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date time that the run execution started. Optional.
    startedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of successfully completed users in the run. Required.
    successfulUsersCount *int32
    // The related taskProcessingResults.
    taskProcessingResults []TaskProcessingResultable
    // The totalTasksCount property
    totalTasksCount *int32
    // The total number of unprocessed tasks in the run execution. Required.
    totalUnprocessedTasksCount *int32
    // The total number of users in the workflow execution. Required.
    totalUsersCount *int32
    // The associated individual user execution.
    userProcessingResults []UserProcessingResultable
    // The workflowExecutionType property
    workflowExecutionType *WorkflowExecutionType
}
// NewRun instantiates a new run and sets the default values.
func NewRun()(*Run) {
    m := &Run{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.run";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRunFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRunFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRun(), nil
}
// GetCompletedDateTime gets the completedDateTime property value. The date time that the run completed. Value is null if the workflow hasn't completed. Optional.
func (m *Run) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDateTime
}
// GetFailedTasksCount gets the failedTasksCount property value. The number of tasks that failed in the run execution. Required.
func (m *Run) GetFailedTasksCount()(*int32) {
    return m.failedTasksCount
}
// GetFailedUsersCount gets the failedUsersCount property value. The number of users that failed in the run execution. Required.
func (m *Run) GetFailedUsersCount()(*int32) {
    return m.failedUsersCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Run) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["completedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCompletedDateTime)
    res["failedTasksCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetFailedTasksCount)
    res["failedUsersCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetFailedUsersCount)
    res["lastUpdatedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastUpdatedDateTime)
    res["processingStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseLifecycleWorkflowProcessingStatus , m.SetProcessingStatus)
    res["scheduledDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetScheduledDateTime)
    res["startedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetStartedDateTime)
    res["successfulUsersCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSuccessfulUsersCount)
    res["taskProcessingResults"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTaskProcessingResultFromDiscriminatorValue , m.SetTaskProcessingResults)
    res["totalTasksCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalTasksCount)
    res["totalUnprocessedTasksCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalUnprocessedTasksCount)
    res["totalUsersCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalUsersCount)
    res["userProcessingResults"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserProcessingResultFromDiscriminatorValue , m.SetUserProcessingResults)
    res["workflowExecutionType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWorkflowExecutionType , m.SetWorkflowExecutionType)
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The datetime that the run was last updated. Optional.
func (m *Run) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetProcessingStatus gets the processingStatus property value. The processingStatus property
func (m *Run) GetProcessingStatus()(*LifecycleWorkflowProcessingStatus) {
    return m.processingStatus
}
// GetScheduledDateTime gets the scheduledDateTime property value. The date time that the run is scheduled to be executed for a workflow. Required.
func (m *Run) GetScheduledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.scheduledDateTime
}
// GetStartedDateTime gets the startedDateTime property value. The date time that the run execution started. Optional.
func (m *Run) GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startedDateTime
}
// GetSuccessfulUsersCount gets the successfulUsersCount property value. The number of successfully completed users in the run. Required.
func (m *Run) GetSuccessfulUsersCount()(*int32) {
    return m.successfulUsersCount
}
// GetTaskProcessingResults gets the taskProcessingResults property value. The related taskProcessingResults.
func (m *Run) GetTaskProcessingResults()([]TaskProcessingResultable) {
    return m.taskProcessingResults
}
// GetTotalTasksCount gets the totalTasksCount property value. The totalTasksCount property
func (m *Run) GetTotalTasksCount()(*int32) {
    return m.totalTasksCount
}
// GetTotalUnprocessedTasksCount gets the totalUnprocessedTasksCount property value. The total number of unprocessed tasks in the run execution. Required.
func (m *Run) GetTotalUnprocessedTasksCount()(*int32) {
    return m.totalUnprocessedTasksCount
}
// GetTotalUsersCount gets the totalUsersCount property value. The total number of users in the workflow execution. Required.
func (m *Run) GetTotalUsersCount()(*int32) {
    return m.totalUsersCount
}
// GetUserProcessingResults gets the userProcessingResults property value. The associated individual user execution.
func (m *Run) GetUserProcessingResults()([]UserProcessingResultable) {
    return m.userProcessingResults
}
// GetWorkflowExecutionType gets the workflowExecutionType property value. The workflowExecutionType property
func (m *Run) GetWorkflowExecutionType()(*WorkflowExecutionType) {
    return m.workflowExecutionType
}
// Serialize serializes information the current object
func (m *Run) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("failedTasksCount", m.GetFailedTasksCount())
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
        err = writer.WriteTimeValue("scheduledDateTime", m.GetScheduledDateTime())
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
    if m.GetTaskProcessingResults() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTaskProcessingResults())
        err = writer.WriteCollectionOfObjectValues("taskProcessingResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalTasksCount", m.GetTotalTasksCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalUnprocessedTasksCount", m.GetTotalUnprocessedTasksCount())
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
    if m.GetUserProcessingResults() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserProcessingResults())
        err = writer.WriteCollectionOfObjectValues("userProcessingResults", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkflowExecutionType() != nil {
        cast := (*m.GetWorkflowExecutionType()).String()
        err = writer.WriteStringValue("workflowExecutionType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletedDateTime sets the completedDateTime property value. The date time that the run completed. Value is null if the workflow hasn't completed. Optional.
func (m *Run) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// SetFailedTasksCount sets the failedTasksCount property value. The number of tasks that failed in the run execution. Required.
func (m *Run) SetFailedTasksCount(value *int32)() {
    m.failedTasksCount = value
}
// SetFailedUsersCount sets the failedUsersCount property value. The number of users that failed in the run execution. Required.
func (m *Run) SetFailedUsersCount(value *int32)() {
    m.failedUsersCount = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The datetime that the run was last updated. Optional.
func (m *Run) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetProcessingStatus sets the processingStatus property value. The processingStatus property
func (m *Run) SetProcessingStatus(value *LifecycleWorkflowProcessingStatus)() {
    m.processingStatus = value
}
// SetScheduledDateTime sets the scheduledDateTime property value. The date time that the run is scheduled to be executed for a workflow. Required.
func (m *Run) SetScheduledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.scheduledDateTime = value
}
// SetStartedDateTime sets the startedDateTime property value. The date time that the run execution started. Optional.
func (m *Run) SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startedDateTime = value
}
// SetSuccessfulUsersCount sets the successfulUsersCount property value. The number of successfully completed users in the run. Required.
func (m *Run) SetSuccessfulUsersCount(value *int32)() {
    m.successfulUsersCount = value
}
// SetTaskProcessingResults sets the taskProcessingResults property value. The related taskProcessingResults.
func (m *Run) SetTaskProcessingResults(value []TaskProcessingResultable)() {
    m.taskProcessingResults = value
}
// SetTotalTasksCount sets the totalTasksCount property value. The totalTasksCount property
func (m *Run) SetTotalTasksCount(value *int32)() {
    m.totalTasksCount = value
}
// SetTotalUnprocessedTasksCount sets the totalUnprocessedTasksCount property value. The total number of unprocessed tasks in the run execution. Required.
func (m *Run) SetTotalUnprocessedTasksCount(value *int32)() {
    m.totalUnprocessedTasksCount = value
}
// SetTotalUsersCount sets the totalUsersCount property value. The total number of users in the workflow execution. Required.
func (m *Run) SetTotalUsersCount(value *int32)() {
    m.totalUsersCount = value
}
// SetUserProcessingResults sets the userProcessingResults property value. The associated individual user execution.
func (m *Run) SetUserProcessingResults(value []UserProcessingResultable)() {
    m.userProcessingResults = value
}
// SetWorkflowExecutionType sets the workflowExecutionType property value. The workflowExecutionType property
func (m *Run) SetWorkflowExecutionType(value *WorkflowExecutionType)() {
    m.workflowExecutionType = value
}
