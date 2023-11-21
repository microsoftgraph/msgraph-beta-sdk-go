package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OpenShift 
type OpenShift struct {
    ChangeTrackedEntity
}
// NewOpenShift instantiates a new openShift and sets the default values.
func NewOpenShift()(*OpenShift) {
    m := &OpenShift{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    odataTypeValue := "#microsoft.graph.openShift"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOpenShiftFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOpenShiftFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOpenShift(), nil
}
// GetDraftOpenShift gets the draftOpenShift property value. An unpublished open shift.
func (m *OpenShift) GetDraftOpenShift()(OpenShiftItemable) {
    val, err := m.GetBackingStore().Get("draftOpenShift")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OpenShiftItemable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OpenShift) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["draftOpenShift"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOpenShiftItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDraftOpenShift(val.(OpenShiftItemable))
        }
        return nil
    }
    res["isStagedForDeletion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsStagedForDeletion(val)
        }
        return nil
    }
    res["schedulingGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedulingGroupId(val)
        }
        return nil
    }
    res["schedulingGroupName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedulingGroupName(val)
        }
        return nil
    }
    res["sharedOpenShift"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOpenShiftItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedOpenShift(val.(OpenShiftItemable))
        }
        return nil
    }
    res["teamId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamId(val)
        }
        return nil
    }
    res["teamName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamName(val)
        }
        return nil
    }
    return res
}
// GetIsStagedForDeletion gets the isStagedForDeletion property value. The isStagedForDeletion property
func (m *OpenShift) GetIsStagedForDeletion()(*bool) {
    val, err := m.GetBackingStore().Get("isStagedForDeletion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSchedulingGroupId gets the schedulingGroupId property value. ID for the scheduling group that the open shift belongs to.
func (m *OpenShift) GetSchedulingGroupId()(*string) {
    val, err := m.GetBackingStore().Get("schedulingGroupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSchedulingGroupName gets the schedulingGroupName property value. The schedulingGroupName property
func (m *OpenShift) GetSchedulingGroupName()(*string) {
    val, err := m.GetBackingStore().Get("schedulingGroupName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSharedOpenShift gets the sharedOpenShift property value. A published open shift.
func (m *OpenShift) GetSharedOpenShift()(OpenShiftItemable) {
    val, err := m.GetBackingStore().Get("sharedOpenShift")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OpenShiftItemable)
    }
    return nil
}
// GetTeamId gets the teamId property value. The teamId property
func (m *OpenShift) GetTeamId()(*string) {
    val, err := m.GetBackingStore().Get("teamId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTeamName gets the teamName property value. The teamName property
func (m *OpenShift) GetTeamName()(*string) {
    val, err := m.GetBackingStore().Get("teamName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OpenShift) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("draftOpenShift", m.GetDraftOpenShift())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isStagedForDeletion", m.GetIsStagedForDeletion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("schedulingGroupId", m.GetSchedulingGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharedOpenShift", m.GetSharedOpenShift())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDraftOpenShift sets the draftOpenShift property value. An unpublished open shift.
func (m *OpenShift) SetDraftOpenShift(value OpenShiftItemable)() {
    err := m.GetBackingStore().Set("draftOpenShift", value)
    if err != nil {
        panic(err)
    }
}
// SetIsStagedForDeletion sets the isStagedForDeletion property value. The isStagedForDeletion property
func (m *OpenShift) SetIsStagedForDeletion(value *bool)() {
    err := m.GetBackingStore().Set("isStagedForDeletion", value)
    if err != nil {
        panic(err)
    }
}
// SetSchedulingGroupId sets the schedulingGroupId property value. ID for the scheduling group that the open shift belongs to.
func (m *OpenShift) SetSchedulingGroupId(value *string)() {
    err := m.GetBackingStore().Set("schedulingGroupId", value)
    if err != nil {
        panic(err)
    }
}
// SetSchedulingGroupName sets the schedulingGroupName property value. The schedulingGroupName property
func (m *OpenShift) SetSchedulingGroupName(value *string)() {
    err := m.GetBackingStore().Set("schedulingGroupName", value)
    if err != nil {
        panic(err)
    }
}
// SetSharedOpenShift sets the sharedOpenShift property value. A published open shift.
func (m *OpenShift) SetSharedOpenShift(value OpenShiftItemable)() {
    err := m.GetBackingStore().Set("sharedOpenShift", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamId sets the teamId property value. The teamId property
func (m *OpenShift) SetTeamId(value *string)() {
    err := m.GetBackingStore().Set("teamId", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamName sets the teamName property value. The teamName property
func (m *OpenShift) SetTeamName(value *string)() {
    err := m.GetBackingStore().Set("teamName", value)
    if err != nil {
        panic(err)
    }
}
// OpenShiftable 
type OpenShiftable interface {
    ChangeTrackedEntityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDraftOpenShift()(OpenShiftItemable)
    GetIsStagedForDeletion()(*bool)
    GetSchedulingGroupId()(*string)
    GetSchedulingGroupName()(*string)
    GetSharedOpenShift()(OpenShiftItemable)
    GetTeamId()(*string)
    GetTeamName()(*string)
    SetDraftOpenShift(value OpenShiftItemable)()
    SetIsStagedForDeletion(value *bool)()
    SetSchedulingGroupId(value *string)()
    SetSchedulingGroupName(value *string)()
    SetSharedOpenShift(value OpenShiftItemable)()
    SetTeamId(value *string)()
    SetTeamName(value *string)()
}
