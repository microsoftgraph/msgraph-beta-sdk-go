package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerBucket 
type PlannerBucket struct {
    PlannerDelta
}
// NewPlannerBucket instantiates a new plannerBucket and sets the default values.
func NewPlannerBucket()(*PlannerBucket) {
    m := &PlannerBucket{
        PlannerDelta: *NewPlannerDelta(),
    }
    return m
}
// CreatePlannerBucketFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerBucketFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerBucket(), nil
}
// GetArchivalInfo gets the archivalInfo property value. The archivalInfo property
func (m *PlannerBucket) GetArchivalInfo()(PlannerArchivalInfoable) {
    val, err := m.GetBackingStore().Get("archivalInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerArchivalInfoable)
    }
    return nil
}
// GetCreationSource gets the creationSource property value. Contains information about the origin of the bucket.
func (m *PlannerBucket) GetCreationSource()(PlannerBucketCreationable) {
    val, err := m.GetBackingStore().Get("creationSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerBucketCreationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerBucket) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["creationSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerBucketCreationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationSource(val.(PlannerBucketCreationable))
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["orderHint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrderHint(val)
        }
        return nil
    }
    res["planId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlanId(val)
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
    return res
}
// GetIsArchived gets the isArchived property value. The isArchived property
func (m *PlannerBucket) GetIsArchived()(*bool) {
    val, err := m.GetBackingStore().Get("isArchived")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetName gets the name property value. Name of the bucket.
func (m *PlannerBucket) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOrderHint gets the orderHint property value. Hint used to order items of this type in a list view. For details about the supported format, see Using order hints in Planner.
func (m *PlannerBucket) GetOrderHint()(*string) {
    val, err := m.GetBackingStore().Get("orderHint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPlanId gets the planId property value. Plan ID to which the bucket belongs.
func (m *PlannerBucket) GetPlanId()(*string) {
    val, err := m.GetBackingStore().Get("planId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTasks gets the tasks property value. Read-only. Nullable. The collection of tasks in the bucket.
func (m *PlannerBucket) GetTasks()([]PlannerTaskable) {
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
func (m *PlannerBucket) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteObjectValue("creationSource", m.GetCreationSource())
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
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("orderHint", m.GetOrderHint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("planId", m.GetPlanId())
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
    return nil
}
// SetArchivalInfo sets the archivalInfo property value. The archivalInfo property
func (m *PlannerBucket) SetArchivalInfo(value PlannerArchivalInfoable)() {
    err := m.GetBackingStore().Set("archivalInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetCreationSource sets the creationSource property value. Contains information about the origin of the bucket.
func (m *PlannerBucket) SetCreationSource(value PlannerBucketCreationable)() {
    err := m.GetBackingStore().Set("creationSource", value)
    if err != nil {
        panic(err)
    }
}
// SetIsArchived sets the isArchived property value. The isArchived property
func (m *PlannerBucket) SetIsArchived(value *bool)() {
    err := m.GetBackingStore().Set("isArchived", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. Name of the bucket.
func (m *PlannerBucket) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetOrderHint sets the orderHint property value. Hint used to order items of this type in a list view. For details about the supported format, see Using order hints in Planner.
func (m *PlannerBucket) SetOrderHint(value *string)() {
    err := m.GetBackingStore().Set("orderHint", value)
    if err != nil {
        panic(err)
    }
}
// SetPlanId sets the planId property value. Plan ID to which the bucket belongs.
func (m *PlannerBucket) SetPlanId(value *string)() {
    err := m.GetBackingStore().Set("planId", value)
    if err != nil {
        panic(err)
    }
}
// SetTasks sets the tasks property value. Read-only. Nullable. The collection of tasks in the bucket.
func (m *PlannerBucket) SetTasks(value []PlannerTaskable)() {
    err := m.GetBackingStore().Set("tasks", value)
    if err != nil {
        panic(err)
    }
}
// PlannerBucketable 
type PlannerBucketable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerDeltaable
    GetArchivalInfo()(PlannerArchivalInfoable)
    GetCreationSource()(PlannerBucketCreationable)
    GetIsArchived()(*bool)
    GetName()(*string)
    GetOrderHint()(*string)
    GetPlanId()(*string)
    GetTasks()([]PlannerTaskable)
    SetArchivalInfo(value PlannerArchivalInfoable)()
    SetCreationSource(value PlannerBucketCreationable)()
    SetIsArchived(value *bool)()
    SetName(value *string)()
    SetOrderHint(value *string)()
    SetPlanId(value *string)()
    SetTasks(value []PlannerTaskable)()
}
