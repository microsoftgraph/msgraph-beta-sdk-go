package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerUser 
type PlannerUser struct {
    PlannerDelta
    // The all property
    all []PlannerDeltaable
    // A collection containing the references to the plans that the user has marked as favorites.
    favoritePlanReferences PlannerFavoritePlanReferenceCollectionable
    // Read-only. Nullable. Returns the plannerPlans that the user marked as favorites.
    favoritePlans []PlannerPlanable
    // The plans property
    plans []PlannerPlanable
    // A collection containing references to the plans that were viewed recently by the user in apps that support recent plans.
    recentPlanReferences PlannerRecentPlanReferenceCollectionable
    // Read-only. Nullable. Returns the plannerPlans that have been recently viewed by the user in apps that support recent plans.
    recentPlans []PlannerPlanable
    // Read-only. Nullable. Returns the plannerPlans contained by the plannerRosters the user is a member.
    rosterPlans []PlannerPlanable
    // Read-only. Nullable. Returns the plannerTasks assigned to the user.
    tasks []PlannerTaskable
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
    return m.all
}
// GetFavoritePlanReferences gets the favoritePlanReferences property value. A collection containing the references to the plans that the user has marked as favorites.
func (m *PlannerUser) GetFavoritePlanReferences()(PlannerFavoritePlanReferenceCollectionable) {
    return m.favoritePlanReferences
}
// GetFavoritePlans gets the favoritePlans property value. Read-only. Nullable. Returns the plannerPlans that the user marked as favorites.
func (m *PlannerUser) GetFavoritePlans()([]PlannerPlanable) {
    return m.favoritePlans
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerDelta.GetFieldDeserializers()
    res["all"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePlannerDeltaFromDiscriminatorValue , m.SetAll)
    res["favoritePlanReferences"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePlannerFavoritePlanReferenceCollectionFromDiscriminatorValue , m.SetFavoritePlanReferences)
    res["favoritePlans"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePlannerPlanFromDiscriminatorValue , m.SetFavoritePlans)
    res["plans"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePlannerPlanFromDiscriminatorValue , m.SetPlans)
    res["recentPlanReferences"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePlannerRecentPlanReferenceCollectionFromDiscriminatorValue , m.SetRecentPlanReferences)
    res["recentPlans"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePlannerPlanFromDiscriminatorValue , m.SetRecentPlans)
    res["rosterPlans"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePlannerPlanFromDiscriminatorValue , m.SetRosterPlans)
    res["tasks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePlannerTaskFromDiscriminatorValue , m.SetTasks)
    return res
}
// GetPlans gets the plans property value. The plans property
func (m *PlannerUser) GetPlans()([]PlannerPlanable) {
    return m.plans
}
// GetRecentPlanReferences gets the recentPlanReferences property value. A collection containing references to the plans that were viewed recently by the user in apps that support recent plans.
func (m *PlannerUser) GetRecentPlanReferences()(PlannerRecentPlanReferenceCollectionable) {
    return m.recentPlanReferences
}
// GetRecentPlans gets the recentPlans property value. Read-only. Nullable. Returns the plannerPlans that have been recently viewed by the user in apps that support recent plans.
func (m *PlannerUser) GetRecentPlans()([]PlannerPlanable) {
    return m.recentPlans
}
// GetRosterPlans gets the rosterPlans property value. Read-only. Nullable. Returns the plannerPlans contained by the plannerRosters the user is a member.
func (m *PlannerUser) GetRosterPlans()([]PlannerPlanable) {
    return m.rosterPlans
}
// GetTasks gets the tasks property value. Read-only. Nullable. Returns the plannerTasks assigned to the user.
func (m *PlannerUser) GetTasks()([]PlannerTaskable) {
    return m.tasks
}
// Serialize serializes information the current object
func (m *PlannerUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerDelta.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAll() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAll())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetFavoritePlans())
        err = writer.WriteCollectionOfObjectValues("favoritePlans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlans() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPlans())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRecentPlans())
        err = writer.WriteCollectionOfObjectValues("recentPlans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRosterPlans() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRosterPlans())
        err = writer.WriteCollectionOfObjectValues("rosterPlans", cast)
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
// SetAll sets the all property value. The all property
func (m *PlannerUser) SetAll(value []PlannerDeltaable)() {
    m.all = value
}
// SetFavoritePlanReferences sets the favoritePlanReferences property value. A collection containing the references to the plans that the user has marked as favorites.
func (m *PlannerUser) SetFavoritePlanReferences(value PlannerFavoritePlanReferenceCollectionable)() {
    m.favoritePlanReferences = value
}
// SetFavoritePlans sets the favoritePlans property value. Read-only. Nullable. Returns the plannerPlans that the user marked as favorites.
func (m *PlannerUser) SetFavoritePlans(value []PlannerPlanable)() {
    m.favoritePlans = value
}
// SetPlans sets the plans property value. The plans property
func (m *PlannerUser) SetPlans(value []PlannerPlanable)() {
    m.plans = value
}
// SetRecentPlanReferences sets the recentPlanReferences property value. A collection containing references to the plans that were viewed recently by the user in apps that support recent plans.
func (m *PlannerUser) SetRecentPlanReferences(value PlannerRecentPlanReferenceCollectionable)() {
    m.recentPlanReferences = value
}
// SetRecentPlans sets the recentPlans property value. Read-only. Nullable. Returns the plannerPlans that have been recently viewed by the user in apps that support recent plans.
func (m *PlannerUser) SetRecentPlans(value []PlannerPlanable)() {
    m.recentPlans = value
}
// SetRosterPlans sets the rosterPlans property value. Read-only. Nullable. Returns the plannerPlans contained by the plannerRosters the user is a member.
func (m *PlannerUser) SetRosterPlans(value []PlannerPlanable)() {
    m.rosterPlans = value
}
// SetTasks sets the tasks property value. Read-only. Nullable. Returns the plannerTasks assigned to the user.
func (m *PlannerUser) SetTasks(value []PlannerTaskable)() {
    m.tasks = value
}
