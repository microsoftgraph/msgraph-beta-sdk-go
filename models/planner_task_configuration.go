package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PlannerTaskConfiguration struct {
    Entity
}
// NewPlannerTaskConfiguration instantiates a new PlannerTaskConfiguration and sets the default values.
func NewPlannerTaskConfiguration()(*PlannerTaskConfiguration) {
    m := &PlannerTaskConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerTaskConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePlannerTaskConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerTaskConfiguration(), nil
}
// GetEditPolicy gets the editPolicy property value. Policy configuration for tasks created for the businessScenario when they're being changed outside of the scenario.
// returns a PlannerTaskPolicyable when successful
func (m *PlannerTaskConfiguration) GetEditPolicy()(PlannerTaskPolicyable) {
    val, err := m.GetBackingStore().Get("editPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerTaskPolicyable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PlannerTaskConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["editPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerTaskPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEditPolicy(val.(PlannerTaskPolicyable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PlannerTaskConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("editPolicy", m.GetEditPolicy())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEditPolicy sets the editPolicy property value. Policy configuration for tasks created for the businessScenario when they're being changed outside of the scenario.
func (m *PlannerTaskConfiguration) SetEditPolicy(value PlannerTaskPolicyable)() {
    err := m.GetBackingStore().Set("editPolicy", value)
    if err != nil {
        panic(err)
    }
}
type PlannerTaskConfigurationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEditPolicy()(PlannerTaskPolicyable)
    SetEditPolicy(value PlannerTaskPolicyable)()
}
