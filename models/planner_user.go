package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerUser 
type PlannerUser struct {
    PlannerDelta
}
// NewPlannerUser instantiates a new PlannerUser and sets the default values.
func NewPlannerUser()(*PlannerUser) {
    m := &PlannerUser{
        PlannerDelta: *NewPlannerDelta(),
    }
    return m
}
// CreatePlannerUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerUser(), nil
}
// GetAll gets the all property value. The all property
func (m *PlannerUser) GetAll()([]PlannerDeltaable) {
    val, err := m.GetBackingStore().Get("all")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PlannerDeltaable)
    }
    return nil
}
// GetFavoritePlanReferences gets the favoritePlanReferences property value. A collection that contains the references to the plans that the user has marked as favorites.
func (m *PlannerUser) GetFavoritePlanReferences()(PlannerFavoritePlanReferenceCollectionable) {
    val, err := m.GetBackingStore().Get("favoritePlanReferences")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerFavoritePlanReferenceCollectionable)
    }
    return nil
}
// GetFavoritePlans gets the favoritePlans property value. Read-only. Nullable. Returns the plannerPlans that the user marked as favorites.
func (m *PlannerUser) GetFavoritePlans()([]PlannerPlanable) {
    val, err := m.GetBackingStore().Get("favoritePlans")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PlannerPlanable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerDelta.GetFieldDeserializers()
    res["all"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerDeltaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerDeltaable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerDeltaable)
            }
            m.SetAll(res)
        }
        return nil
    }
    res["favoritePlanReferences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerFavoritePlanReferenceCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFavoritePlanReferences(val.(PlannerFavoritePlanReferenceCollectionable))
        }
        return nil
    }
    res["favoritePlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerPlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerPlanable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerPlanable)
            }
            m.SetFavoritePlans(res)
        }
        return nil
    }
    res["plans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["recentPlanReferences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerRecentPlanReferenceCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecentPlanReferences(val.(PlannerRecentPlanReferenceCollectionable))
        }
        return nil
    }
    res["recentPlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerPlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerPlanable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerPlanable)
            }
            m.SetRecentPlans(res)
        }
        return nil
    }
    res["rosterPlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerPlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerPlanable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerPlanable)
            }
            m.SetRosterPlans(res)
        }
        return nil
    }
    res["tasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerTaskable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerTaskable)
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
// GetPlans gets the plans property value. The plans property
func (m *PlannerUser) GetPlans()([]PlannerPlanable) {
    val, err := m.GetBackingStore().Get("plans")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PlannerPlanable)
    }
    return nil
}
// GetRecentPlanReferences gets the recentPlanReferences property value. A collection that contains references to the plans that were viewed recently by the user in apps that support recent plans.
func (m *PlannerUser) GetRecentPlanReferences()(PlannerRecentPlanReferenceCollectionable) {
    val, err := m.GetBackingStore().Get("recentPlanReferences")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerRecentPlanReferenceCollectionable)
    }
    return nil
}
// GetRecentPlans gets the recentPlans property value. Read-only. Nullable. Returns the plannerPlans that have been recently viewed by the user in apps that support recent plans.
func (m *PlannerUser) GetRecentPlans()([]PlannerPlanable) {
    val, err := m.GetBackingStore().Get("recentPlans")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PlannerPlanable)
    }
    return nil
}
// GetRosterPlans gets the rosterPlans property value. Read-only. Nullable. Returns the plannerPlans contained by the plannerRosters the user is a member.
func (m *PlannerUser) GetRosterPlans()([]PlannerPlanable) {
    val, err := m.GetBackingStore().Get("rosterPlans")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PlannerPlanable)
    }
    return nil
}
// GetTasks gets the tasks property value. Read-only. Nullable. Returns the plannerTasks assigned to the user.
func (m *PlannerUser) GetTasks()([]PlannerTaskable) {
    val, err := m.GetBackingStore().Get("tasks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PlannerTaskable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PlannerUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerDelta.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAll() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAll()))
        for i, v := range m.GetAll() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFavoritePlans()))
        for i, v := range m.GetFavoritePlans() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("favoritePlans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPlans()))
        for i, v := range m.GetPlans() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecentPlans()))
        for i, v := range m.GetRecentPlans() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("recentPlans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRosterPlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRosterPlans()))
        for i, v := range m.GetRosterPlans() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("rosterPlans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTasks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAll sets the all property value. The all property
func (m *PlannerUser) SetAll(value []PlannerDeltaable)() {
    err := m.GetBackingStore().Set("all", value)
    if err != nil {
        panic(err)
    }
}
// SetFavoritePlanReferences sets the favoritePlanReferences property value. A collection that contains the references to the plans that the user has marked as favorites.
func (m *PlannerUser) SetFavoritePlanReferences(value PlannerFavoritePlanReferenceCollectionable)() {
    err := m.GetBackingStore().Set("favoritePlanReferences", value)
    if err != nil {
        panic(err)
    }
}
// SetFavoritePlans sets the favoritePlans property value. Read-only. Nullable. Returns the plannerPlans that the user marked as favorites.
func (m *PlannerUser) SetFavoritePlans(value []PlannerPlanable)() {
    err := m.GetBackingStore().Set("favoritePlans", value)
    if err != nil {
        panic(err)
    }
}
// SetPlans sets the plans property value. The plans property
func (m *PlannerUser) SetPlans(value []PlannerPlanable)() {
    err := m.GetBackingStore().Set("plans", value)
    if err != nil {
        panic(err)
    }
}
// SetRecentPlanReferences sets the recentPlanReferences property value. A collection that contains references to the plans that were viewed recently by the user in apps that support recent plans.
func (m *PlannerUser) SetRecentPlanReferences(value PlannerRecentPlanReferenceCollectionable)() {
    err := m.GetBackingStore().Set("recentPlanReferences", value)
    if err != nil {
        panic(err)
    }
}
// SetRecentPlans sets the recentPlans property value. Read-only. Nullable. Returns the plannerPlans that have been recently viewed by the user in apps that support recent plans.
func (m *PlannerUser) SetRecentPlans(value []PlannerPlanable)() {
    err := m.GetBackingStore().Set("recentPlans", value)
    if err != nil {
        panic(err)
    }
}
// SetRosterPlans sets the rosterPlans property value. Read-only. Nullable. Returns the plannerPlans contained by the plannerRosters the user is a member.
func (m *PlannerUser) SetRosterPlans(value []PlannerPlanable)() {
    err := m.GetBackingStore().Set("rosterPlans", value)
    if err != nil {
        panic(err)
    }
}
// SetTasks sets the tasks property value. Read-only. Nullable. Returns the plannerTasks assigned to the user.
func (m *PlannerUser) SetTasks(value []PlannerTaskable)() {
    err := m.GetBackingStore().Set("tasks", value)
    if err != nil {
        panic(err)
    }
}
// PlannerUserable 
type PlannerUserable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerDeltaable
    GetAll()([]PlannerDeltaable)
    GetFavoritePlanReferences()(PlannerFavoritePlanReferenceCollectionable)
    GetFavoritePlans()([]PlannerPlanable)
    GetPlans()([]PlannerPlanable)
    GetRecentPlanReferences()(PlannerRecentPlanReferenceCollectionable)
    GetRecentPlans()([]PlannerPlanable)
    GetRosterPlans()([]PlannerPlanable)
    GetTasks()([]PlannerTaskable)
    SetAll(value []PlannerDeltaable)()
    SetFavoritePlanReferences(value PlannerFavoritePlanReferenceCollectionable)()
    SetFavoritePlans(value []PlannerPlanable)()
    SetPlans(value []PlannerPlanable)()
    SetRecentPlanReferences(value PlannerRecentPlanReferenceCollectionable)()
    SetRecentPlans(value []PlannerPlanable)()
    SetRosterPlans(value []PlannerPlanable)()
    SetTasks(value []PlannerTaskable)()
}
