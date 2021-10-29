package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AttackSimulationRoot struct {
    Entity
    // Represent attack simulation and training campaign of a tenant.
    simulations []Simulation;
}
// Instantiates a new attackSimulationRoot and sets the default values.
func NewAttackSimulationRoot()(*AttackSimulationRoot) {
    m := &AttackSimulationRoot{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the simulations property value. Represent attack simulation and training campaign of a tenant.
func (m *AttackSimulationRoot) GetSimulations()([]Simulation) {
    if m == nil {
        return nil
    } else {
        return m.simulations
    }
}
// The deserialization information for the current model
func (m *AttackSimulationRoot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["simulations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSimulation() })
        if err != nil {
            return err
        }
        res := make([]Simulation, len(val))
        for i, v := range val {
            res[i] = *(v.(*Simulation))
        }
        m.SetSimulations(res)
        return nil
    }
    return res
}
func (m *AttackSimulationRoot) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AttackSimulationRoot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSimulations()))
        for i, v := range m.GetSimulations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("simulations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the simulations property value. Represent attack simulation and training campaign of a tenant.
// Parameters:
//  - value : Value to set for the simulations property.
func (m *AttackSimulationRoot) SetSimulations(value []Simulation)() {
    m.simulations = value
}
