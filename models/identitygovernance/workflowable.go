package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Workflowable 
type Workflowable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WorkflowBaseable
    GetDeletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExecutionScope()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)
    GetId()(*string)
    GetIsEnabled()(*bool)
    GetIsSchedulingEnabled()(*bool)
    GetNextScheduleRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRuns()([]Runable)
    GetTaskReports()([]TaskReportable)
    GetUserProcessingResults()([]UserProcessingResultable)
    GetVersion()(*int32)
    GetVersions()([]WorkflowVersionable)
    SetDeletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExecutionScope(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)()
    SetId(value *string)()
    SetIsEnabled(value *bool)()
    SetIsSchedulingEnabled(value *bool)()
    SetNextScheduleRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRuns(value []Runable)()
    SetTaskReports(value []TaskReportable)()
    SetUserProcessingResults(value []UserProcessingResultable)()
    SetVersion(value *int32)()
    SetVersions(value []WorkflowVersionable)()
}
