package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// UserProcessingResultable 
type UserProcessingResultable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFailedTasksCount()(*int32)
    GetProcessingStatus()(*LifecycleWorkflowProcessingStatus)
    GetScheduledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubject()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)
    GetTaskProcessingResults()([]TaskProcessingResultable)
    GetTotalTasksCount()(*int32)
    GetTotalUnprocessedTasksCount()(*int32)
    GetWorkflowExecutionType()(*WorkflowExecutionType)
    GetWorkflowVersion()(*int32)
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFailedTasksCount(value *int32)()
    SetProcessingStatus(value *LifecycleWorkflowProcessingStatus)()
    SetScheduledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubject(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)()
    SetTaskProcessingResults(value []TaskProcessingResultable)()
    SetTotalTasksCount(value *int32)()
    SetTotalUnprocessedTasksCount(value *int32)()
    SetWorkflowExecutionType(value *WorkflowExecutionType)()
    SetWorkflowVersion(value *int32)()
}
