package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PlannerUser struct {
    PlannerDelta
    // 
    all []PlannerDelta;
    // A collection containing the references to the plans that the user has marked as favorites.
    favoritePlanReferences *PlannerFavoritePlanReferenceCollection;
    // Read-only. Nullable. Returns the plannerPlans that the user marked as favorites.
    favoritePlans []PlannerPlan;
    // Read-only. Nullable. Returns the plannerTasks assigned to the user.
    plans []PlannerPlan;
    // A collection containing references to the plans that were viewed recently by the user in apps that support recent plans.
    recentPlanReferences *PlannerRecentPlanReferenceCollection;
    // Read-only. Nullable. Returns the plannerPlans that have been recently viewed by the user in apps that support recent plans.
    recentPlans []PlannerPlan;
    // Read-only. Nullable. Returns the plannerPlans contained by the plannerRosters the user is a member.
    rosterPlans []PlannerPlan;
    // Read-only. Nullable. Returns the plannerPlans shared with the user.
    tasks []PlannerTask;
}
// Instantiates a new plannerUser and sets the default values.
func NewPlannerUser()(*PlannerUser) {
    m := &PlannerUser{
        PlannerDelta: *NewPlannerDelta(),
    }
    return m
}
// Gets the all property value. 
func (m *PlannerUser) GetAll()([]PlannerDelta) {
    if m == nil {
        return nil
    } else {
        return m.all
    }
}
// Gets the favoritePlanReferences property value. A collection containing the references to the plans that the user has marked as favorites.
func (m *PlannerUser) GetFavoritePlanReferences()(*PlannerFavoritePlanReferenceCollection) {
    if m == nil {
        return nil
    } else {
        return m.favoritePlanReferences
    }
}
// Gets the favoritePlans property value. Read-only. Nullable. Returns the plannerPlans that the user marked as favorites.
func (m *PlannerUser) GetFavoritePlans()([]PlannerPlan) {
    if m == nil {
        return nil
    } else {
        return m.favoritePlans
    }
}
// Gets the plans property value. Read-only. Nullable. Returns the plannerTasks assigned to the user.
func (m *PlannerUser) GetPlans()([]PlannerPlan) {
    if m == nil {
        return nil
    } else {
        return m.plans
    }
}
// Gets the recentPlanReferences property value. A collection containing references to the plans that were viewed recently by the user in apps that support recent plans.
func (m *PlannerUser) GetRecentPlanReferences()(*PlannerRecentPlanReferenceCollection) {
    if m == nil {
        return nil
    } else {
        return m.recentPlanReferences
    }
}
// Gets the recentPlans property value. Read-only. Nullable. Returns the plannerPlans that have been recently viewed by the user in apps that support recent plans.
func (m *PlannerUser) GetRecentPlans()([]PlannerPlan) {
    if m == nil {
        return nil
    } else {
        return m.recentPlans
    }
}
// Gets the rosterPlans property value. Read-only. Nullable. Returns the plannerPlans contained by the plannerRosters the user is a member.
func (m *PlannerUser) GetRosterPlans()([]PlannerPlan) {
    if m == nil {
        return nil
    } else {
        return m.rosterPlans
    }
}
// Gets the tasks property value. Read-only. Nullable. Returns the plannerPlans shared with the user.
func (m *PlannerUser) GetTasks()([]PlannerTask) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
// The deserialization information for the current model
func (m *PlannerUser) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PlannerDelta.GetFieldDeserializers()
    res["all"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerDelta() })
        if err != nil {
            return err
        }
        res := make([]PlannerDelta, len(val))
        for i, v := range val {
            res[i] = *(v.(*PlannerDelta))
        }
        m.SetAll(res)
        return nil
    }
    res["favoritePlanReferences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerFavoritePlanReferenceCollection() })
        if err != nil {
            return err
        }
        m.SetFavoritePlanReferences(val.(*PlannerFavoritePlanReferenceCollection))
        return nil
    }
    res["favoritePlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerPlan() })
        if err != nil {
            return err
        }
        res := make([]PlannerPlan, len(val))
        for i, v := range val {
            res[i] = *(v.(*PlannerPlan))
        }
        m.SetFavoritePlans(res)
        return nil
    }
    res["plans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerPlan() })
        if err != nil {
            return err
        }
        res := make([]PlannerPlan, len(val))
        for i, v := range val {
            res[i] = *(v.(*PlannerPlan))
        }
        m.SetPlans(res)
        return nil
    }
    res["recentPlanReferences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerRecentPlanReferenceCollection() })
        if err != nil {
            return err
        }
        m.SetRecentPlanReferences(val.(*PlannerRecentPlanReferenceCollection))
        return nil
    }
    res["recentPlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerPlan() })
        if err != nil {
            return err
        }
        res := make([]PlannerPlan, len(val))
        for i, v := range val {
            res[i] = *(v.(*PlannerPlan))
        }
        m.SetRecentPlans(res)
        return nil
    }
    res["rosterPlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerPlan() })
        if err != nil {
            return err
        }
        res := make([]PlannerPlan, len(val))
        for i, v := range val {
            res[i] = *(v.(*PlannerPlan))
        }
        m.SetRosterPlans(res)
        return nil
    }
    res["tasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerTask() })
        if err != nil {
            return err
        }
        res := make([]PlannerTask, len(val))
        for i, v := range val {
            res[i] = *(v.(*PlannerTask))
        }
        m.SetTasks(res)
        return nil
    }
    return res
}
func (m *PlannerUser) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PlannerUser) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PlannerDelta.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAll()))
        for i, v := range m.GetAll() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("all", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("favoritePlanReferences", m.GetFavoritePlanReferences())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFavoritePlans()))
        for i, v := range m.GetFavoritePlans() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("favoritePlans", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPlans()))
        for i, v := range m.GetPlans() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("plans", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("recentPlanReferences", m.GetRecentPlanReferences())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRecentPlans()))
        for i, v := range m.GetRecentPlans() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("recentPlans", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRosterPlans()))
        for i, v := range m.GetRosterPlans() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("rosterPlans", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the all property value. 
// Parameters:
//  - value : Value to set for the all property.
func (m *PlannerUser) SetAll(value []PlannerDelta)() {
    m.all = value
}
// Sets the favoritePlanReferences property value. A collection containing the references to the plans that the user has marked as favorites.
// Parameters:
//  - value : Value to set for the favoritePlanReferences property.
func (m *PlannerUser) SetFavoritePlanReferences(value *PlannerFavoritePlanReferenceCollection)() {
    m.favoritePlanReferences = value
}
// Sets the favoritePlans property value. Read-only. Nullable. Returns the plannerPlans that the user marked as favorites.
// Parameters:
//  - value : Value to set for the favoritePlans property.
func (m *PlannerUser) SetFavoritePlans(value []PlannerPlan)() {
    m.favoritePlans = value
}
// Sets the plans property value. Read-only. Nullable. Returns the plannerTasks assigned to the user.
// Parameters:
//  - value : Value to set for the plans property.
func (m *PlannerUser) SetPlans(value []PlannerPlan)() {
    m.plans = value
}
// Sets the recentPlanReferences property value. A collection containing references to the plans that were viewed recently by the user in apps that support recent plans.
// Parameters:
//  - value : Value to set for the recentPlanReferences property.
func (m *PlannerUser) SetRecentPlanReferences(value *PlannerRecentPlanReferenceCollection)() {
    m.recentPlanReferences = value
}
// Sets the recentPlans property value. Read-only. Nullable. Returns the plannerPlans that have been recently viewed by the user in apps that support recent plans.
// Parameters:
//  - value : Value to set for the recentPlans property.
func (m *PlannerUser) SetRecentPlans(value []PlannerPlan)() {
    m.recentPlans = value
}
// Sets the rosterPlans property value. Read-only. Nullable. Returns the plannerPlans contained by the plannerRosters the user is a member.
// Parameters:
//  - value : Value to set for the rosterPlans property.
func (m *PlannerUser) SetRosterPlans(value []PlannerPlan)() {
    m.rosterPlans = value
}
// Sets the tasks property value. Read-only. Nullable. Returns the plannerPlans shared with the user.
// Parameters:
//  - value : Value to set for the tasks property.
func (m *PlannerUser) SetTasks(value []PlannerTask)() {
    m.tasks = value
}
