package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type TriggerAndScopeBasedConditions struct {
    WorkflowExecutionConditions
}
// NewTriggerAndScopeBasedConditions instantiates a new TriggerAndScopeBasedConditions and sets the default values.
func NewTriggerAndScopeBasedConditions()(*TriggerAndScopeBasedConditions) {
    m := &TriggerAndScopeBasedConditions{
        WorkflowExecutionConditions: *NewWorkflowExecutionConditions(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.triggerAndScopeBasedConditions"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTriggerAndScopeBasedConditionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTriggerAndScopeBasedConditionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTriggerAndScopeBasedConditions(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TriggerAndScopeBasedConditions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WorkflowExecutionConditions.GetFieldDeserializers()
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSubjectSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubjectSetable))
        }
        return nil
    }
    res["trigger"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkflowExecutionTriggerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrigger(val.(WorkflowExecutionTriggerable))
        }
        return nil
    }
    return res
}
// GetScope gets the scope property value. Defines who the workflow runs for.
// returns a SubjectSetable when successful
func (m *TriggerAndScopeBasedConditions) GetScope()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubjectSetable) {
    val, err := m.GetBackingStore().Get("scope")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubjectSetable)
    }
    return nil
}
// GetTrigger gets the trigger property value. What triggers a workflow to run.
// returns a WorkflowExecutionTriggerable when successful
func (m *TriggerAndScopeBasedConditions) GetTrigger()(WorkflowExecutionTriggerable) {
    val, err := m.GetBackingStore().Get("trigger")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkflowExecutionTriggerable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TriggerAndScopeBasedConditions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WorkflowExecutionConditions.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("trigger", m.GetTrigger())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetScope sets the scope property value. Defines who the workflow runs for.
func (m *TriggerAndScopeBasedConditions) SetScope(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubjectSetable)() {
    err := m.GetBackingStore().Set("scope", value)
    if err != nil {
        panic(err)
    }
}
// SetTrigger sets the trigger property value. What triggers a workflow to run.
func (m *TriggerAndScopeBasedConditions) SetTrigger(value WorkflowExecutionTriggerable)() {
    err := m.GetBackingStore().Set("trigger", value)
    if err != nil {
        panic(err)
    }
}
type TriggerAndScopeBasedConditionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WorkflowExecutionConditionsable
    GetScope()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubjectSetable)
    GetTrigger()(WorkflowExecutionTriggerable)
    SetScope(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubjectSetable)()
    SetTrigger(value WorkflowExecutionTriggerable)()
}
