package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CustomTaskExtensionCalloutData 
type CustomTaskExtensionCalloutData struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionData
    // The subject property
    subject ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable
    // The task property
    task Taskable
    // The taskProcessingresult property
    taskProcessingresult TaskProcessingResultable
    // The workflow property
    workflow Workflowable
}
// NewCustomTaskExtensionCalloutData instantiates a new CustomTaskExtensionCalloutData and sets the default values.
func NewCustomTaskExtensionCalloutData()(*CustomTaskExtensionCalloutData) {
    m := &CustomTaskExtensionCalloutData{
        CustomExtensionData: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewCustomExtensionData(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.customTaskExtensionCalloutData"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCustomTaskExtensionCalloutDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomTaskExtensionCalloutDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomTaskExtensionCalloutData(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomTaskExtensionCalloutData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomExtensionData.GetFieldDeserializers()
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable))
        }
        return nil
    }
    res["task"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTask(val.(Taskable))
        }
        return nil
    }
    res["taskProcessingresult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTaskProcessingResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaskProcessingresult(val.(TaskProcessingResultable))
        }
        return nil
    }
    res["workflow"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkflowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkflow(val.(Workflowable))
        }
        return nil
    }
    return res
}
// GetSubject gets the subject property value. The subject property
func (m *CustomTaskExtensionCalloutData) GetSubject()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable) {
    return m.subject
}
// GetTask gets the task property value. The task property
func (m *CustomTaskExtensionCalloutData) GetTask()(Taskable) {
    return m.task
}
// GetTaskProcessingresult gets the taskProcessingresult property value. The taskProcessingresult property
func (m *CustomTaskExtensionCalloutData) GetTaskProcessingresult()(TaskProcessingResultable) {
    return m.taskProcessingresult
}
// GetWorkflow gets the workflow property value. The workflow property
func (m *CustomTaskExtensionCalloutData) GetWorkflow()(Workflowable) {
    return m.workflow
}
// Serialize serializes information the current object
func (m *CustomTaskExtensionCalloutData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomExtensionData.Serialize(writer)
    if err != nil {
        return err
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
    {
        err = writer.WriteObjectValue("taskProcessingresult", m.GetTaskProcessingresult())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("workflow", m.GetWorkflow())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSubject sets the subject property value. The subject property
func (m *CustomTaskExtensionCalloutData) SetSubject(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)() {
    m.subject = value
}
// SetTask sets the task property value. The task property
func (m *CustomTaskExtensionCalloutData) SetTask(value Taskable)() {
    m.task = value
}
// SetTaskProcessingresult sets the taskProcessingresult property value. The taskProcessingresult property
func (m *CustomTaskExtensionCalloutData) SetTaskProcessingresult(value TaskProcessingResultable)() {
    m.taskProcessingresult = value
}
// SetWorkflow sets the workflow property value. The workflow property
func (m *CustomTaskExtensionCalloutData) SetWorkflow(value Workflowable)() {
    m.workflow = value
}
