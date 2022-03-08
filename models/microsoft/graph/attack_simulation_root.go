package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttackSimulationRoot provides operations to manage the attackSimulation property of the microsoft.graph.security entity.
type AttackSimulationRoot struct {
    Entity
    // 
    simulationAutomations []SimulationAutomationable;
    // Represent attack simulation and training campaign of a tenant.
    simulations []Simulationable;
}
// NewAttackSimulationRoot instantiates a new attackSimulationRoot and sets the default values.
func NewAttackSimulationRoot()(*AttackSimulationRoot) {
    m := &AttackSimulationRoot{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAttackSimulationRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttackSimulationRootFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAttackSimulationRoot(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttackSimulationRoot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["simulationAutomations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["simulations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
// GetSimulationAutomations gets the simulationAutomations property value. 
func (m *AttackSimulationRoot) GetSimulationAutomations()([]SimulationAutomationable) {
    if m == nil {
        return nil
    } else {
        return m.simulationAutomations
    }
}
// GetSimulations gets the simulations property value. Represent attack simulation and training campaign of a tenant.
func (m *AttackSimulationRoot) GetSimulations()([]Simulationable) {
    if m == nil {
        return nil
    } else {
        return m.simulations
    }
}
func (m *AttackSimulationRoot) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AttackSimulationRoot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSimulationAutomations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSimulationAutomations()))
        for i, v := range m.GetSimulationAutomations() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("simulationAutomations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSimulations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSimulations()))
        for i, v := range m.GetSimulations() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("simulations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSimulationAutomations sets the simulationAutomations property value. 
func (m *AttackSimulationRoot) SetSimulationAutomations(value []SimulationAutomationable)() {
    if m != nil {
        m.simulationAutomations = value
    }
}
// SetSimulations sets the simulations property value. Represent attack simulation and training campaign of a tenant.
func (m *AttackSimulationRoot) SetSimulations(value []Simulationable)() {
    if m != nil {
        m.simulations = value
    }
}
