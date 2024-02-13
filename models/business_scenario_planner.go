package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BusinessScenarioPlanner struct {
    Entity
}
// NewBusinessScenarioPlanner instantiates a new BusinessScenarioPlanner and sets the default values.
func NewBusinessScenarioPlanner()(*BusinessScenarioPlanner) {
    m := &BusinessScenarioPlanner{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBusinessScenarioPlannerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBusinessScenarioPlannerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBusinessScenarioPlanner(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BusinessScenarioPlanner) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["planConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerPlanConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlanConfiguration(val.(PlannerPlanConfigurationable))
        }
        return nil
    }
    res["taskConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerTaskConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaskConfiguration(val.(PlannerTaskConfigurationable))
        }
        return nil
    }
    res["tasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBusinessScenarioTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BusinessScenarioTaskable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BusinessScenarioTaskable)
                }
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
// GetPlanConfiguration gets the planConfiguration property value. The configuration of Planner plans that will be created for the scenario.
// returns a PlannerPlanConfigurationable when successful
func (m *BusinessScenarioPlanner) GetPlanConfiguration()(PlannerPlanConfigurationable) {
    val, err := m.GetBackingStore().Get("planConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerPlanConfigurationable)
    }
    return nil
}
// GetTaskConfiguration gets the taskConfiguration property value. The configuration of Planner tasks that will be created for the scenario.
// returns a PlannerTaskConfigurationable when successful
func (m *BusinessScenarioPlanner) GetTaskConfiguration()(PlannerTaskConfigurationable) {
    val, err := m.GetBackingStore().Get("taskConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerTaskConfigurationable)
    }
    return nil
}
// GetTasks gets the tasks property value. The Planner tasks for the scenario.
// returns a []BusinessScenarioTaskable when successful
func (m *BusinessScenarioPlanner) GetTasks()([]BusinessScenarioTaskable) {
    val, err := m.GetBackingStore().Get("tasks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]BusinessScenarioTaskable)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPlanConfiguration sets the planConfiguration property value. The configuration of Planner plans that will be created for the scenario.
func (m *BusinessScenarioPlanner) SetPlanConfiguration(value PlannerPlanConfigurationable)() {
    err := m.GetBackingStore().Set("planConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetTaskConfiguration sets the taskConfiguration property value. The configuration of Planner tasks that will be created for the scenario.
func (m *BusinessScenarioPlanner) SetTaskConfiguration(value PlannerTaskConfigurationable)() {
    err := m.GetBackingStore().Set("taskConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetTasks sets the tasks property value. The Planner tasks for the scenario.
func (m *BusinessScenarioPlanner) SetTasks(value []BusinessScenarioTaskable)() {
    err := m.GetBackingStore().Set("tasks", value)
    if err != nil {
        panic(err)
    }
}
type BusinessScenarioPlannerable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPlanConfiguration()(PlannerPlanConfigurationable)
    GetTaskConfiguration()(PlannerTaskConfigurationable)
    GetTasks()([]BusinessScenarioTaskable)
    SetPlanConfiguration(value PlannerPlanConfigurationable)()
    SetTaskConfiguration(value PlannerTaskConfigurationable)()
    SetTasks(value []BusinessScenarioTaskable)()
}
