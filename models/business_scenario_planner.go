package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessScenarioPlanner 
type BusinessScenarioPlanner struct {
    Entity
    // The planConfiguration property
    planConfiguration PlannerPlanConfigurationable
    // The taskConfiguration property
    taskConfiguration PlannerTaskConfigurationable
    // The tasks property
    tasks []BusinessScenarioTaskable
}
// NewBusinessScenarioPlanner instantiates a new businessScenarioPlanner and sets the default values.
func NewBusinessScenarioPlanner()(*BusinessScenarioPlanner) {
    m := &BusinessScenarioPlanner{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBusinessScenarioPlannerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBusinessScenarioPlannerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBusinessScenarioPlanner(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BusinessScenarioPlanner) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["planConfiguration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePlannerPlanConfigurationFromDiscriminatorValue , m.SetPlanConfiguration)
    res["taskConfiguration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePlannerTaskConfigurationFromDiscriminatorValue , m.SetTaskConfiguration)
    res["tasks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBusinessScenarioTaskFromDiscriminatorValue , m.SetTasks)
    return res
}
// GetPlanConfiguration gets the planConfiguration property value. The planConfiguration property
func (m *BusinessScenarioPlanner) GetPlanConfiguration()(PlannerPlanConfigurationable) {
    return m.planConfiguration
}
// GetTaskConfiguration gets the taskConfiguration property value. The taskConfiguration property
func (m *BusinessScenarioPlanner) GetTaskConfiguration()(PlannerTaskConfigurationable) {
    return m.taskConfiguration
}
// GetTasks gets the tasks property value. The tasks property
func (m *BusinessScenarioPlanner) GetTasks()([]BusinessScenarioTaskable) {
    return m.tasks
}
// Serialize serializes information the current object
func (m *BusinessScenarioPlanner) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("planConfiguration", m.GetPlanConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("taskConfiguration", m.GetTaskConfiguration())
        if err != nil {
            return err
        }
    }
    if m.GetTasks() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTasks())
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPlanConfiguration sets the planConfiguration property value. The planConfiguration property
func (m *BusinessScenarioPlanner) SetPlanConfiguration(value PlannerPlanConfigurationable)() {
    m.planConfiguration = value
}
// SetTaskConfiguration sets the taskConfiguration property value. The taskConfiguration property
func (m *BusinessScenarioPlanner) SetTaskConfiguration(value PlannerTaskConfigurationable)() {
    m.taskConfiguration = value
}
// SetTasks sets the tasks property value. The tasks property
func (m *BusinessScenarioPlanner) SetTasks(value []BusinessScenarioTaskable)() {
    m.tasks = value
}
