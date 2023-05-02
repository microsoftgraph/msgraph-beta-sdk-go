package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttackSimulationRoot 
type AttackSimulationRoot struct {
    Entity
}
// NewAttackSimulationRoot instantiates a new AttackSimulationRoot and sets the default values.
func NewAttackSimulationRoot()(*AttackSimulationRoot) {
    m := &AttackSimulationRoot{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAttackSimulationRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttackSimulationRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttackSimulationRoot(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttackSimulationRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttackSimulationOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttackSimulationOperationable, len(val))
            for i, v := range val {
                res[i] = v.(AttackSimulationOperationable)
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["payloads"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePayloadFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Payloadable, len(val))
            for i, v := range val {
                res[i] = v.(Payloadable)
            }
            m.SetPayloads(res)
        }
        return nil
    }
    res["simulationAutomations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSimulationAutomationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SimulationAutomationable, len(val))
            for i, v := range val {
                res[i] = v.(SimulationAutomationable)
            }
            m.SetSimulationAutomations(res)
        }
        return nil
    }
    res["simulations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSimulationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Simulationable, len(val))
            for i, v := range val {
                res[i] = v.(Simulationable)
            }
            m.SetSimulations(res)
        }
        return nil
    }
    return res
}
// GetOperations gets the operations property value. Represents an attack simulation training operation.
func (m *AttackSimulationRoot) GetOperations()([]AttackSimulationOperationable) {
    val, err := m.GetBackingStore().Get("operations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AttackSimulationOperationable)
    }
    return nil
}
// GetPayloads gets the payloads property value. Represents an attack simulation training campaign payload in a tenant.
func (m *AttackSimulationRoot) GetPayloads()([]Payloadable) {
    val, err := m.GetBackingStore().Get("payloads")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Payloadable)
    }
    return nil
}
// GetSimulationAutomations gets the simulationAutomations property value. Represents simulation automation created to run on a tenant.
func (m *AttackSimulationRoot) GetSimulationAutomations()([]SimulationAutomationable) {
    val, err := m.GetBackingStore().Get("simulationAutomations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SimulationAutomationable)
    }
    return nil
}
// GetSimulations gets the simulations property value. Represents an attack simulation training campaign in a tenant.
func (m *AttackSimulationRoot) GetSimulations()([]Simulationable) {
    val, err := m.GetBackingStore().Get("simulations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Simulationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AttackSimulationRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPayloads() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPayloads()))
        for i, v := range m.GetPayloads() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("payloads", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSimulationAutomations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSimulationAutomations()))
        for i, v := range m.GetSimulationAutomations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("simulationAutomations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSimulations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSimulations()))
        for i, v := range m.GetSimulations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("simulations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOperations sets the operations property value. Represents an attack simulation training operation.
func (m *AttackSimulationRoot) SetOperations(value []AttackSimulationOperationable)() {
    err := m.GetBackingStore().Set("operations", value)
    if err != nil {
        panic(err)
    }
}
// SetPayloads sets the payloads property value. Represents an attack simulation training campaign payload in a tenant.
func (m *AttackSimulationRoot) SetPayloads(value []Payloadable)() {
    err := m.GetBackingStore().Set("payloads", value)
    if err != nil {
        panic(err)
    }
}
// SetSimulationAutomations sets the simulationAutomations property value. Represents simulation automation created to run on a tenant.
func (m *AttackSimulationRoot) SetSimulationAutomations(value []SimulationAutomationable)() {
    err := m.GetBackingStore().Set("simulationAutomations", value)
    if err != nil {
        panic(err)
    }
}
// SetSimulations sets the simulations property value. Represents an attack simulation training campaign in a tenant.
func (m *AttackSimulationRoot) SetSimulations(value []Simulationable)() {
    err := m.GetBackingStore().Set("simulations", value)
    if err != nil {
        panic(err)
    }
}
// AttackSimulationRootable 
type AttackSimulationRootable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOperations()([]AttackSimulationOperationable)
    GetPayloads()([]Payloadable)
    GetSimulationAutomations()([]SimulationAutomationable)
    GetSimulations()([]Simulationable)
    SetOperations(value []AttackSimulationOperationable)()
    SetPayloads(value []Payloadable)()
    SetSimulationAutomations(value []SimulationAutomationable)()
    SetSimulations(value []Simulationable)()
}
