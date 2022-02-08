package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlannerUser 
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
// NewPlannerUser instantiates a new plannerUser and sets the default values.
func NewPlannerUser()(*PlannerUser) {
    m := &PlannerUser{
        PlannerDelta: *NewPlannerDelta(),
    }
    return m
}
// GetAll gets the all property value. 
func (m *PlannerUser) GetAll()([]PlannerDelta) {
    if m == nil {
        return nil
    } else {
        return m.all
    }
}
// GetFavoritePlanReferences gets the favoritePlanReferences property value. A collection containing the references to the plans that the user has marked as favorites.
func (m *PlannerUser) GetFavoritePlanReferences()(*PlannerFavoritePlanReferenceCollection) {
    if m == nil {
        return nil
    } else {
        return m.favoritePlanReferences
    }
}
// GetFavoritePlans gets the favoritePlans property value. Read-only. Nullable. Returns the plannerPlans that the user marked as favorites.
func (m *PlannerUser) GetFavoritePlans()([]PlannerPlan) {
    if m == nil {
        return nil
    } else {
        return m.favoritePlans
    }
}
// GetPlans gets the plans property value. Read-only. Nullable. Returns the plannerTasks assigned to the user.
func (m *PlannerUser) GetPlans()([]PlannerPlan) {
    if m == nil {
        return nil
    } else {
        return m.plans
    }
}
// GetRecentPlanReferences gets the recentPlanReferences property value. A collection containing references to the plans that were viewed recently by the user in apps that support recent plans.
func (m *PlannerUser) GetRecentPlanReferences()(*PlannerRecentPlanReferenceCollection) {
    if m == nil {
        return nil
    } else {
        return m.recentPlanReferences
    }
}
// GetRecentPlans gets the recentPlans property value. Read-only. Nullable. Returns the plannerPlans that have been recently viewed by the user in apps that support recent plans.
func (m *PlannerUser) GetRecentPlans()([]PlannerPlan) {
    if m == nil {
        return nil
    } else {
        return m.recentPlans
    }
}
// GetRosterPlans gets the rosterPlans property value. Read-only. Nullable. Returns the plannerPlans contained by the plannerRosters the user is a member.
func (m *PlannerUser) GetRosterPlans()([]PlannerPlan) {
    if m == nil {
        return nil
    } else {
        return m.rosterPlans
    }
}
// GetTasks gets the tasks property value. Read-only. Nullable. Returns the plannerPlans shared with the user.
func (m *PlannerUser) GetTasks()([]PlannerTask) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerUser) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PlannerDelta.GetFieldDeserializers()
    res["all"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerDelta() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerDelta, len(val))
            for i, v := range val {
                res[i] = *(v.(*PlannerDelta))
            }
            m.SetAll(res)
        }
        return nil
    }
    res["favoritePlanReferences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerFavoritePlanReferenceCollection() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFavoritePlanReferences(val.(*PlannerFavoritePlanReferenceCollection))
        }
        return nil
    }
    res["favoritePlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerPlan() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerPlan, len(val))
            for i, v := range val {
                res[i] = *(v.(*PlannerPlan))
            }
            m.SetFavoritePlans(res)
        }
        return nil
    }
    res["plans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerPlan() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerPlan, len(val))
            for i, v := range val {
                res[i] = *(v.(*PlannerPlan))
            }
            m.SetPlans(res)
        }
        return nil
    }
    res["recentPlanReferences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerRecentPlanReferenceCollection() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecentPlanReferences(val.(*PlannerRecentPlanReferenceCollection))
        }
        return nil
    }
    res["recentPlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerPlan() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerPlan, len(val))
            for i, v := range val {
                res[i] = *(v.(*PlannerPlan))
            }
            m.SetRecentPlans(res)
        }
        return nil
    }
    res["rosterPlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerPlan() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerPlan, len(val))
            for i, v := range val {
                res[i] = *(v.(*PlannerPlan))
            }
            m.SetRosterPlans(res)
        }
        return nil
    }
    res["tasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerTask() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerTask, len(val))
            for i, v := range val {
                res[i] = *(v.(*PlannerTask))
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
func (m *PlannerUser) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PlannerUser) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PlannerDelta.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAll() != nil {
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
    if m.GetFavoritePlans() != nil {
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
    if m.GetPlans() != nil {
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
    if m.GetRecentPlans() != nil {
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
    if m.GetRosterPlans() != nil {
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
    if m.GetTasks() != nil {
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
// SetAll sets the all property value. 
func (m *PlannerUser) SetAll(value []PlannerDelta)() {
    if m != nil {
        m.all = value
    }
}
// SetFavoritePlanReferences sets the favoritePlanReferences property value. A collection containing the references to the plans that the user has marked as favorites.
func (m *PlannerUser) SetFavoritePlanReferences(value *PlannerFavoritePlanReferenceCollection)() {
    if m != nil {
        m.favoritePlanReferences = value
    }
}
// SetFavoritePlans sets the favoritePlans property value. Read-only. Nullable. Returns the plannerPlans that the user marked as favorites.
func (m *PlannerUser) SetFavoritePlans(value []PlannerPlan)() {
    if m != nil {
        m.favoritePlans = value
    }
}
// SetPlans sets the plans property value. Read-only. Nullable. Returns the plannerTasks assigned to the user.
func (m *PlannerUser) SetPlans(value []PlannerPlan)() {
    if m != nil {
        m.plans = value
    }
}
// SetRecentPlanReferences sets the recentPlanReferences property value. A collection containing references to the plans that were viewed recently by the user in apps that support recent plans.
func (m *PlannerUser) SetRecentPlanReferences(value *PlannerRecentPlanReferenceCollection)() {
    if m != nil {
        m.recentPlanReferences = value
    }
}
// SetRecentPlans sets the recentPlans property value. Read-only. Nullable. Returns the plannerPlans that have been recently viewed by the user in apps that support recent plans.
func (m *PlannerUser) SetRecentPlans(value []PlannerPlan)() {
    if m != nil {
        m.recentPlans = value
    }
}
// SetRosterPlans sets the rosterPlans property value. Read-only. Nullable. Returns the plannerPlans contained by the plannerRosters the user is a member.
func (m *PlannerUser) SetRosterPlans(value []PlannerPlan)() {
    if m != nil {
        m.rosterPlans = value
    }
}
// SetTasks sets the tasks property value. Read-only. Nullable. Returns the plannerPlans shared with the user.
func (m *PlannerUser) SetTasks(value []PlannerTask)() {
    if m != nil {
        m.tasks = value
    }
}
