package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// UserProcessingResult provides operations to manage the collection of accessReview entities.
type UserProcessingResult struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The date time that the workflow execution for a user completed. Value is null if the workflow hasn't completed.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of tasks that failed in the workflow execution.
    failedTasksCount *int32
    // The processingStatus property
    processingStatus *LifecycleWorkflowProcessingStatus
    // The date time that the workflow is scheduled to be executed for a user.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    scheduledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date time that the workflow execution started. Value is null if the workflow execution has not started.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    startedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The subject property
    subject ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable
    // The associated individual task execution.
    taskProcessingResults []TaskProcessingResultable
    // The total number of tasks that in the workflow execution.
    totalTasksCount *int32
    // The total number of unprocessed tasks for the workflow.
    totalUnprocessedTasksCount *int32
    // The workflowExecutionType property
    workflowExecutionType *WorkflowExecutionType
    // The version of the workflow that was executed.
    workflowVersion *int32
}
// NewUserProcessingResult instantiates a new userProcessingResult and sets the default values.
func NewUserProcessingResult()(*UserProcessingResult) {
    m := &UserProcessingResult{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.userProcessingResult";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserProcessingResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserProcessingResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserProcessingResult(), nil
}
// GetCompletedDateTime gets the completedDateTime property value. The date time that the workflow execution for a user completed. Value is null if the workflow hasn't completed.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *UserProcessingResult) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDateTime
}
// GetFailedTasksCount gets the failedTasksCount property value. The number of tasks that failed in the workflow execution.
func (m *UserProcessingResult) GetFailedTasksCount()(*int32) {
    return m.failedTasksCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserProcessingResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["completedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCompletedDateTime)
    res["failedTasksCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetFailedTasksCount)
    res["processingStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseLifecycleWorkflowProcessingStatus , m.SetProcessingStatus)
    res["scheduledDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetScheduledDateTime)
    res["startedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetStartedDateTime)
    res["subject"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue , m.SetSubject)
    res["taskProcessingResults"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTaskProcessingResultFromDiscriminatorValue , m.SetTaskProcessingResults)
    res["totalTasksCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalTasksCount)
    res["totalUnprocessedTasksCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalUnprocessedTasksCount)
    res["workflowExecutionType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWorkflowExecutionType , m.SetWorkflowExecutionType)
    res["workflowVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetWorkflowVersion)
    return res
}
// GetProcessingStatus gets the processingStatus property value. The processingStatus property
func (m *UserProcessingResult) GetProcessingStatus()(*LifecycleWorkflowProcessingStatus) {
    return m.processingStatus
}
// GetScheduledDateTime gets the scheduledDateTime property value. The date time that the workflow is scheduled to be executed for a user.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *UserProcessingResult) GetScheduledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.scheduledDateTime
}
// GetStartedDateTime gets the startedDateTime property value. The date time that the workflow execution started. Value is null if the workflow execution has not started.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *UserProcessingResult) GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startedDateTime
}
// GetSubject gets the subject property value. The subject property
func (m *UserProcessingResult) GetSubject()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable) {
    return m.subject
}
// GetTaskProcessingResults gets the taskProcessingResults property value. The associated individual task execution.
func (m *UserProcessingResult) GetTaskProcessingResults()([]TaskProcessingResultable) {
    return m.taskProcessingResults
}
// GetTotalTasksCount gets the totalTasksCount property value. The total number of tasks that in the workflow execution.
func (m *UserProcessingResult) GetTotalTasksCount()(*int32) {
    return m.totalTasksCount
}
// GetTotalUnprocessedTasksCount gets the totalUnprocessedTasksCount property value. The total number of unprocessed tasks for the workflow.
func (m *UserProcessingResult) GetTotalUnprocessedTasksCount()(*int32) {
    return m.totalUnprocessedTasksCount
}
// GetWorkflowExecutionType gets the workflowExecutionType property value. The workflowExecutionType property
func (m *UserProcessingResult) GetWorkflowExecutionType()(*WorkflowExecutionType) {
    return m.workflowExecutionType
}
// GetWorkflowVersion gets the workflowVersion property value. The version of the workflow that was executed.
func (m *UserProcessingResult) GetWorkflowVersion()(*int32) {
    return m.workflowVersion
}
// Serialize serializes information the current object
func (m *UserProcessingResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("subject", m.GetSubject())
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
    if m.GetWorkflowExecutionType() != nil {
        cast := (*m.GetWorkflowExecutionType()).String()
        err = writer.WriteStringValue("workflowExecutionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workflowVersion", m.GetWorkflowVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletedDateTime sets the completedDateTime property value. The date time that the workflow execution for a user completed. Value is null if the workflow hasn't completed.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *UserProcessingResult) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// SetFailedTasksCount sets the failedTasksCount property value. The number of tasks that failed in the workflow execution.
func (m *UserProcessingResult) SetFailedTasksCount(value *int32)() {
    m.failedTasksCount = value
}
// SetProcessingStatus sets the processingStatus property value. The processingStatus property
func (m *UserProcessingResult) SetProcessingStatus(value *LifecycleWorkflowProcessingStatus)() {
    m.processingStatus = value
}
// SetScheduledDateTime sets the scheduledDateTime property value. The date time that the workflow is scheduled to be executed for a user.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *UserProcessingResult) SetScheduledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.scheduledDateTime = value
}
// SetStartedDateTime sets the startedDateTime property value. The date time that the workflow execution started. Value is null if the workflow execution has not started.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *UserProcessingResult) SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startedDateTime = value
}
// SetSubject sets the subject property value. The subject property
func (m *UserProcessingResult) SetSubject(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)() {
    m.subject = value
}
// SetTaskProcessingResults sets the taskProcessingResults property value. The associated individual task execution.
func (m *UserProcessingResult) SetTaskProcessingResults(value []TaskProcessingResultable)() {
    m.taskProcessingResults = value
}
// SetTotalTasksCount sets the totalTasksCount property value. The total number of tasks that in the workflow execution.
func (m *UserProcessingResult) SetTotalTasksCount(value *int32)() {
    m.totalTasksCount = value
}
// SetTotalUnprocessedTasksCount sets the totalUnprocessedTasksCount property value. The total number of unprocessed tasks for the workflow.
func (m *UserProcessingResult) SetTotalUnprocessedTasksCount(value *int32)() {
    m.totalUnprocessedTasksCount = value
}
// SetWorkflowExecutionType sets the workflowExecutionType property value. The workflowExecutionType property
func (m *UserProcessingResult) SetWorkflowExecutionType(value *WorkflowExecutionType)() {
    m.workflowExecutionType = value
}
// SetWorkflowVersion sets the workflowVersion property value. The version of the workflow that was executed.
func (m *UserProcessingResult) SetWorkflowVersion(value *int32)() {
    m.workflowVersion = value
}
