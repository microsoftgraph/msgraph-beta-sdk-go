package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// TaskProcessingResult provides operations to manage the collection of accessReview entities.
type TaskProcessingResult struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The date time when taskProcessingResult execution ended. Value is null if task execution is still in progress.
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date time when the taskProcessingResult was created. Supports $filter(lt, gt) and orderBy.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Describes why the taskProcessingResult has failed.
    failureReason *string
    // The processingStatus property
    processingStatus *LifecycleWorkflowProcessingStatus
    // The date time when taskProcessingResult execution started. Value is null if task execution has not yet started. Supports $filter(lt, gt) and orderBy.
    startedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The subject property
    subject ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable
    // The task property
    task Taskable
}
// NewTaskProcessingResult instantiates a new taskProcessingResult and sets the default values.
func NewTaskProcessingResult()(*TaskProcessingResult) {
    m := &TaskProcessingResult{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.taskProcessingResult";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTaskProcessingResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTaskProcessingResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTaskProcessingResult(), nil
}
// GetCompletedDateTime gets the completedDateTime property value. The date time when taskProcessingResult execution ended. Value is null if task execution is still in progress.
func (m *TaskProcessingResult) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDateTime
}
// GetCreatedDateTime gets the createdDateTime property value. The date time when the taskProcessingResult was created. Supports $filter(lt, gt) and orderBy.
func (m *TaskProcessingResult) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFailureReason gets the failureReason property value. Describes why the taskProcessingResult has failed.
func (m *TaskProcessingResult) GetFailureReason()(*string) {
    return m.failureReason
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TaskProcessingResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["completedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCompletedDateTime)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["failureReason"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFailureReason)
    res["processingStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseLifecycleWorkflowProcessingStatus , m.SetProcessingStatus)
    res["startedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetStartedDateTime)
    res["subject"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue , m.SetSubject)
    res["task"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTaskFromDiscriminatorValue , m.SetTask)
    return res
}
// GetProcessingStatus gets the processingStatus property value. The processingStatus property
func (m *TaskProcessingResult) GetProcessingStatus()(*LifecycleWorkflowProcessingStatus) {
    return m.processingStatus
}
// GetStartedDateTime gets the startedDateTime property value. The date time when taskProcessingResult execution started. Value is null if task execution has not yet started. Supports $filter(lt, gt) and orderBy.
func (m *TaskProcessingResult) GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startedDateTime
}
// GetSubject gets the subject property value. The subject property
func (m *TaskProcessingResult) GetSubject()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable) {
    return m.subject
}
// GetTask gets the task property value. The task property
func (m *TaskProcessingResult) GetTask()(Taskable) {
    return m.task
}
// Serialize serializes information the current object
func (m *TaskProcessingResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("failureReason", m.GetFailureReason())
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
    {
        err = writer.WriteObjectValue("task", m.GetTask())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletedDateTime sets the completedDateTime property value. The date time when taskProcessingResult execution ended. Value is null if task execution is still in progress.
func (m *TaskProcessingResult) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date time when the taskProcessingResult was created. Supports $filter(lt, gt) and orderBy.
func (m *TaskProcessingResult) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetFailureReason sets the failureReason property value. Describes why the taskProcessingResult has failed.
func (m *TaskProcessingResult) SetFailureReason(value *string)() {
    m.failureReason = value
}
// SetProcessingStatus sets the processingStatus property value. The processingStatus property
func (m *TaskProcessingResult) SetProcessingStatus(value *LifecycleWorkflowProcessingStatus)() {
    m.processingStatus = value
}
// SetStartedDateTime sets the startedDateTime property value. The date time when taskProcessingResult execution started. Value is null if task execution has not yet started. Supports $filter(lt, gt) and orderBy.
func (m *TaskProcessingResult) SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startedDateTime = value
}
// SetSubject sets the subject property value. The subject property
func (m *TaskProcessingResult) SetSubject(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)() {
    m.subject = value
}
// SetTask sets the task property value. The task property
func (m *TaskProcessingResult) SetTask(value Taskable)() {
    m.task = value
}
