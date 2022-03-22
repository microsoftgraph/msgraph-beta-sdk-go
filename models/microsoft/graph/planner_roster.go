package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlannerRoster 
type PlannerRoster struct {
    Entity
    // Retrieves the members of the plannerRoster.
    members []PlannerRosterMemberable;
    // Retrieves the plans contained by the plannerRoster.
    plans []PlannerPlanable;
}
// NewPlannerRoster instantiates a new plannerRoster and sets the default values.
func NewPlannerRoster()(*PlannerRoster) {
    m := &PlannerRoster{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerRosterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerRosterFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPlannerRoster(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerRoster) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerRosterMemberFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerRosterMemberable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerRosterMemberable)
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["plans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerPlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerPlanable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerPlanable)
            }
            m.SetPlans(res)
        }
        return nil
    }
    return res
}
// GetMembers gets the members property value. Retrieves the members of the plannerRoster.
func (m *PlannerRoster) GetMembers()([]PlannerRosterMemberable) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// GetPlans gets the plans property value. Retrieves the plans contained by the plannerRoster.
func (m *PlannerRoster) GetPlans()([]PlannerPlanable) {
    if m == nil {
        return nil
    } else {
        return m.plans
    }
}
// Serialize serializes information the current object
func (m *PlannerRoster) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMembers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlans() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPlans()))
        for i, v := range m.GetPlans() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("plans", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMembers sets the members property value. Retrieves the members of the plannerRoster.
func (m *PlannerRoster) SetMembers(value []PlannerRosterMemberable)() {
    if m != nil {
        m.members = value
    }
}
// SetPlans sets the plans property value. Retrieves the plans contained by the plannerRoster.
func (m *PlannerRoster) SetPlans(value []PlannerPlanable)() {
    if m != nil {
        m.plans = value
    }
}
