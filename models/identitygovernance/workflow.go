package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Workflow 
type Workflow struct {
    WorkflowBase
    // The deletedDateTime property
    deletedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The executionScope property
    executionScope []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable
    // The id property
    id *string
    // The isEnabled property
    isEnabled *bool
    // The isSchedulingEnabled property
    isSchedulingEnabled *bool
    // The nextScheduleRunDateTime property
    nextScheduleRunDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The runs property
    runs []Runable
    // The taskReports property
    taskReports []TaskReportable
    // The userProcessingResults property
    userProcessingResults []UserProcessingResultable
    // The version property
    version *int32
    // The versions property
    versions []WorkflowVersionable
}
// NewWorkflow instantiates a new Workflow and sets the default values.
func NewWorkflow()(*Workflow) {
    m := &Workflow{
        WorkflowBase: *NewWorkflowBase(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.workflow";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWorkflowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkflowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkflow(), nil
}
// GetDeletedDateTime gets the deletedDateTime property value. The deletedDateTime property
func (m *Workflow) GetDeletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deletedDateTime
}
// GetExecutionScope gets the executionScope property value. The executionScope property
func (m *Workflow) GetExecutionScope()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable) {
    return m.executionScope
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Workflow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WorkflowBase.GetFieldDeserializers()
    res["deletedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedDateTime(val)
        }
        return nil
    }
    res["executionScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)
            }
            m.SetExecutionScope(res)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["isSchedulingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSchedulingEnabled(val)
        }
        return nil
    }
    res["nextScheduleRunDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextScheduleRunDateTime(val)
        }
        return nil
    }
    res["runs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRunFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Runable, len(val))
            for i, v := range val {
                res[i] = v.(Runable)
            }
            m.SetRuns(res)
        }
        return nil
    }
    res["taskReports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTaskReportFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TaskReportable, len(val))
            for i, v := range val {
                res[i] = v.(TaskReportable)
            }
            m.SetTaskReports(res)
        }
        return nil
    }
    res["userProcessingResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserProcessingResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserProcessingResultable, len(val))
            for i, v := range val {
                res[i] = v.(UserProcessingResultable)
            }
            m.SetUserProcessingResults(res)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    res["versions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkflowVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkflowVersionable, len(val))
            for i, v := range val {
                res[i] = v.(WorkflowVersionable)
            }
            m.SetVersions(res)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *Workflow) GetId()(*string) {
    return m.id
}
// GetIsEnabled gets the isEnabled property value. The isEnabled property
func (m *Workflow) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetIsSchedulingEnabled gets the isSchedulingEnabled property value. The isSchedulingEnabled property
func (m *Workflow) GetIsSchedulingEnabled()(*bool) {
    return m.isSchedulingEnabled
}
// GetNextScheduleRunDateTime gets the nextScheduleRunDateTime property value. The nextScheduleRunDateTime property
func (m *Workflow) GetNextScheduleRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.nextScheduleRunDateTime
}
// GetRuns gets the runs property value. The runs property
func (m *Workflow) GetRuns()([]Runable) {
    return m.runs
}
// GetTaskReports gets the taskReports property value. The taskReports property
func (m *Workflow) GetTaskReports()([]TaskReportable) {
    return m.taskReports
}
// GetUserProcessingResults gets the userProcessingResults property value. The userProcessingResults property
func (m *Workflow) GetUserProcessingResults()([]UserProcessingResultable) {
    return m.userProcessingResults
}
// GetVersion gets the version property value. The version property
func (m *Workflow) GetVersion()(*int32) {
    return m.version
}
// GetVersions gets the versions property value. The versions property
func (m *Workflow) GetVersions()([]WorkflowVersionable) {
    return m.versions
}
// Serialize serializes information the current object
func (m *Workflow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WorkflowBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("deletedDateTime", m.GetDeletedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetExecutionScope() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExecutionScope()))
        for i, v := range m.GetExecutionScope() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("executionScope", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSchedulingEnabled", m.GetIsSchedulingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("nextScheduleRunDateTime", m.GetNextScheduleRunDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRuns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRuns()))
        for i, v := range m.GetRuns() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("runs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTaskReports() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaskReports()))
        for i, v := range m.GetTaskReports() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("taskReports", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserProcessingResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserProcessingResults()))
        for i, v := range m.GetUserProcessingResults() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userProcessingResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    if m.GetVersions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVersions()))
        for i, v := range m.GetVersions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("versions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeletedDateTime sets the deletedDateTime property value. The deletedDateTime property
func (m *Workflow) SetDeletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deletedDateTime = value
}
// SetExecutionScope sets the executionScope property value. The executionScope property
func (m *Workflow) SetExecutionScope(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)() {
    m.executionScope = value
}
// SetId sets the id property value. The id property
func (m *Workflow) SetId(value *string)() {
    m.id = value
}
// SetIsEnabled sets the isEnabled property value. The isEnabled property
func (m *Workflow) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetIsSchedulingEnabled sets the isSchedulingEnabled property value. The isSchedulingEnabled property
func (m *Workflow) SetIsSchedulingEnabled(value *bool)() {
    m.isSchedulingEnabled = value
}
// SetNextScheduleRunDateTime sets the nextScheduleRunDateTime property value. The nextScheduleRunDateTime property
func (m *Workflow) SetNextScheduleRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.nextScheduleRunDateTime = value
}
// SetRuns sets the runs property value. The runs property
func (m *Workflow) SetRuns(value []Runable)() {
    m.runs = value
}
// SetTaskReports sets the taskReports property value. The taskReports property
func (m *Workflow) SetTaskReports(value []TaskReportable)() {
    m.taskReports = value
}
// SetUserProcessingResults sets the userProcessingResults property value. The userProcessingResults property
func (m *Workflow) SetUserProcessingResults(value []UserProcessingResultable)() {
    m.userProcessingResults = value
}
// SetVersion sets the version property value. The version property
func (m *Workflow) SetVersion(value *int32)() {
    m.version = value
}
// SetVersions sets the versions property value. The versions property
func (m *Workflow) SetVersions(value []WorkflowVersionable)() {
    m.versions = value
}
