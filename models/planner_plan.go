package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PlannerPlan struct {
    PlannerDelta
}
// NewPlannerPlan instantiates a new PlannerPlan and sets the default values.
func NewPlannerPlan()(*PlannerPlan) {
    m := &PlannerPlan{
        PlannerDelta: *NewPlannerDelta(),
    }
    return m
}
// CreatePlannerPlanFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePlannerPlanFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerPlan(), nil
}
// GetArchivalInfo gets the archivalInfo property value. Read-only. Nullable. Contains information about who archived or unarchived the plan and why.
// returns a PlannerArchivalInfoable when successful
func (m *PlannerPlan) GetArchivalInfo()(PlannerArchivalInfoable) {
    val, err := m.GetBackingStore().Get("archivalInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerArchivalInfoable)
    }
    return nil
}
// GetBuckets gets the buckets property value. Collection of buckets in the plan. Read-only. Nullable.
// returns a []PlannerBucketable when successful
func (m *PlannerPlan) GetBuckets()([]PlannerBucketable) {
    val, err := m.GetBackingStore().Get("buckets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PlannerBucketable)
    }
    return nil
}
// GetContainer gets the container property value. Identifies the container of the plan. Either specify all properties, or specify only the url, the containerId, and type. After it's set, this property can’t be updated. It changes when a plan is moved from one container to another, using plan move to container. Required.
// returns a PlannerPlanContainerable when successful
func (m *PlannerPlan) GetContainer()(PlannerPlanContainerable) {
    val, err := m.GetBackingStore().Get("container")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerPlanContainerable)
    }
    return nil
}
// GetContexts gets the contexts property value. Read-only. Other user experiences in which this plan is used, represented as plannerPlanContext entries.
// returns a PlannerPlanContextCollectionable when successful
func (m *PlannerPlan) GetContexts()(PlannerPlanContextCollectionable) {
    val, err := m.GetBackingStore().Get("contexts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerPlanContextCollectionable)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. Read-only. The user who created the plan.
// returns a IdentitySetable when successful
func (m *PlannerPlan) GetCreatedBy()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Read-only. Date and time at which the plan is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *PlannerPlan) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCreationSource gets the creationSource property value. Contains information about the origin of the plan.
// returns a PlannerPlanCreationable when successful
func (m *PlannerPlan) GetCreationSource()(PlannerPlanCreationable) {
    val, err := m.GetBackingStore().Get("creationSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerPlanCreationable)
    }
    return nil
}
// GetDetails gets the details property value. Extra details about the plan. Read-only. Nullable.
// returns a PlannerPlanDetailsable when successful
func (m *PlannerPlan) GetDetails()(PlannerPlanDetailsable) {
    val, err := m.GetBackingStore().Get("details")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerPlanDetailsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PlannerPlan) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerDelta.GetFieldDeserializers()
    res["archivalInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerArchivalInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchivalInfo(val.(PlannerArchivalInfoable))
        }
        return nil
    }
    res["buckets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerBucketFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerBucketable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PlannerBucketable)
                }
            }
            m.SetBuckets(res)
        }
        return nil
    }
    res["container"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerPlanContainerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContainer(val.(PlannerPlanContainerable))
        }
        return nil
    }
    res["contexts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerPlanContextCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContexts(val.(PlannerPlanContextCollectionable))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["creationSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerPlanCreationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationSource(val.(PlannerPlanCreationable))
        }
        return nil
    }
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerPlanDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val.(PlannerPlanDetailsable))
        }
        return nil
    }
    res["isArchived"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsArchived(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val)
        }
        return nil
    }
    res["sharedWithContainers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerSharedWithContainerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerSharedWithContainerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PlannerSharedWithContainerable)
                }
            }
            m.SetSharedWithContainers(res)
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
                if v != nil {
                    res[i] = v.(PlannerTaskable)
                }
            }
            m.SetTasks(res)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetIsArchived gets the isArchived property value. Read-only. If set to true, the plan is archived. An archived plan is read-only.
// returns a *bool when successful
func (m *PlannerPlan) GetIsArchived()(*bool) {
    val, err := m.GetBackingStore().Get("isArchived")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOwner gets the owner property value. Use the container property instead. ID of the group that owns the plan. After it's set, this property can’t be updated. This property doesn't return a valid group ID if the container of the plan isn't a group.
// returns a *string when successful
func (m *PlannerPlan) GetOwner()(*string) {
    val, err := m.GetBackingStore().Get("owner")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSharedWithContainers gets the sharedWithContainers property value. List of containers the plan is shared with.
// returns a []PlannerSharedWithContainerable when successful
func (m *PlannerPlan) GetSharedWithContainers()([]PlannerSharedWithContainerable) {
    val, err := m.GetBackingStore().Get("sharedWithContainers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PlannerSharedWithContainerable)
    }
    return nil
}
// GetTasks gets the tasks property value. Collection of tasks in the plan. Read-only. Nullable.
// returns a []PlannerTaskable when successful
func (m *PlannerPlan) GetTasks()([]PlannerTaskable) {
    val, err := m.GetBackingStore().Get("tasks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PlannerTaskable)
    }
    return nil
}
// GetTitle gets the title property value. Required. Title of the plan.
// returns a *string when successful
func (m *PlannerPlan) GetTitle()(*string) {
    val, err := m.GetBackingStore().Get("title")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PlannerPlan) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerDelta.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("archivalInfo", m.GetArchivalInfo())
        if err != nil {
            return err
        }
    }
    if m.GetBuckets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBuckets()))
        for i, v := range m.GetBuckets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("buckets", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("container", m.GetContainer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("contexts", m.GetContexts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("creationSource", m.GetCreationSource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isArchived", m.GetIsArchived())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    if m.GetSharedWithContainers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSharedWithContainers()))
        for i, v := range m.GetSharedWithContainers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sharedWithContainers", cast)
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
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetArchivalInfo sets the archivalInfo property value. Read-only. Nullable. Contains information about who archived or unarchived the plan and why.
func (m *PlannerPlan) SetArchivalInfo(value PlannerArchivalInfoable)() {
    err := m.GetBackingStore().Set("archivalInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetBuckets sets the buckets property value. Collection of buckets in the plan. Read-only. Nullable.
func (m *PlannerPlan) SetBuckets(value []PlannerBucketable)() {
    err := m.GetBackingStore().Set("buckets", value)
    if err != nil {
        panic(err)
    }
}
// SetContainer sets the container property value. Identifies the container of the plan. Either specify all properties, or specify only the url, the containerId, and type. After it's set, this property can’t be updated. It changes when a plan is moved from one container to another, using plan move to container. Required.
func (m *PlannerPlan) SetContainer(value PlannerPlanContainerable)() {
    err := m.GetBackingStore().Set("container", value)
    if err != nil {
        panic(err)
    }
}
// SetContexts sets the contexts property value. Read-only. Other user experiences in which this plan is used, represented as plannerPlanContext entries.
func (m *PlannerPlan) SetContexts(value PlannerPlanContextCollectionable)() {
    err := m.GetBackingStore().Set("contexts", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. Read-only. The user who created the plan.
func (m *PlannerPlan) SetCreatedBy(value IdentitySetable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Read-only. Date and time at which the plan is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PlannerPlan) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCreationSource sets the creationSource property value. Contains information about the origin of the plan.
func (m *PlannerPlan) SetCreationSource(value PlannerPlanCreationable)() {
    err := m.GetBackingStore().Set("creationSource", value)
    if err != nil {
        panic(err)
    }
}
// SetDetails sets the details property value. Extra details about the plan. Read-only. Nullable.
func (m *PlannerPlan) SetDetails(value PlannerPlanDetailsable)() {
    err := m.GetBackingStore().Set("details", value)
    if err != nil {
        panic(err)
    }
}
// SetIsArchived sets the isArchived property value. Read-only. If set to true, the plan is archived. An archived plan is read-only.
func (m *PlannerPlan) SetIsArchived(value *bool)() {
    err := m.GetBackingStore().Set("isArchived", value)
    if err != nil {
        panic(err)
    }
}
// SetOwner sets the owner property value. Use the container property instead. ID of the group that owns the plan. After it's set, this property can’t be updated. This property doesn't return a valid group ID if the container of the plan isn't a group.
func (m *PlannerPlan) SetOwner(value *string)() {
    err := m.GetBackingStore().Set("owner", value)
    if err != nil {
        panic(err)
    }
}
// SetSharedWithContainers sets the sharedWithContainers property value. List of containers the plan is shared with.
func (m *PlannerPlan) SetSharedWithContainers(value []PlannerSharedWithContainerable)() {
    err := m.GetBackingStore().Set("sharedWithContainers", value)
    if err != nil {
        panic(err)
    }
}
// SetTasks sets the tasks property value. Collection of tasks in the plan. Read-only. Nullable.
func (m *PlannerPlan) SetTasks(value []PlannerTaskable)() {
    err := m.GetBackingStore().Set("tasks", value)
    if err != nil {
        panic(err)
    }
}
// SetTitle sets the title property value. Required. Title of the plan.
func (m *PlannerPlan) SetTitle(value *string)() {
    err := m.GetBackingStore().Set("title", value)
    if err != nil {
        panic(err)
    }
}
type PlannerPlanable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerDeltaable
    GetArchivalInfo()(PlannerArchivalInfoable)
    GetBuckets()([]PlannerBucketable)
    GetContainer()(PlannerPlanContainerable)
    GetContexts()(PlannerPlanContextCollectionable)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreationSource()(PlannerPlanCreationable)
    GetDetails()(PlannerPlanDetailsable)
    GetIsArchived()(*bool)
    GetOwner()(*string)
    GetSharedWithContainers()([]PlannerSharedWithContainerable)
    GetTasks()([]PlannerTaskable)
    GetTitle()(*string)
    SetArchivalInfo(value PlannerArchivalInfoable)()
    SetBuckets(value []PlannerBucketable)()
    SetContainer(value PlannerPlanContainerable)()
    SetContexts(value PlannerPlanContextCollectionable)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreationSource(value PlannerPlanCreationable)()
    SetDetails(value PlannerPlanDetailsable)()
    SetIsArchived(value *bool)()
    SetOwner(value *string)()
    SetSharedWithContainers(value []PlannerSharedWithContainerable)()
    SetTasks(value []PlannerTaskable)()
    SetTitle(value *string)()
}
