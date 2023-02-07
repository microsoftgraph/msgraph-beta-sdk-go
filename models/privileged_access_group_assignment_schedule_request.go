package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccessGroupAssignmentScheduleRequest 
type PrivilegedAccessGroupAssignmentScheduleRequest struct {
    PrivilegedAccessScheduleRequest
    // The accessId property
    accessId *PrivilegedAccessGroupRelationships
    // The activatedUsing property
    activatedUsing PrivilegedAccessGroupEligibilityScheduleable
    // The group property
    group Groupable
    // The groupId property
    groupId *string
    // The principal property
    principal DirectoryObjectable
    // The principalId property
    principalId *string
    // The targetSchedule property
    targetSchedule PrivilegedAccessGroupEligibilityScheduleable
    // The targetScheduleId property
    targetScheduleId *string
}
// NewPrivilegedAccessGroupAssignmentScheduleRequest instantiates a new PrivilegedAccessGroupAssignmentScheduleRequest and sets the default values.
func NewPrivilegedAccessGroupAssignmentScheduleRequest()(*PrivilegedAccessGroupAssignmentScheduleRequest) {
    m := &PrivilegedAccessGroupAssignmentScheduleRequest{
        PrivilegedAccessScheduleRequest: *NewPrivilegedAccessScheduleRequest(),
    }
    odataTypeValue := "#microsoft.graph.privilegedAccessGroupAssignmentScheduleRequest"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePrivilegedAccessGroupAssignmentScheduleRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedAccessGroupAssignmentScheduleRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccessGroupAssignmentScheduleRequest(), nil
}
// GetAccessId gets the accessId property value. The accessId property
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetAccessId()(*PrivilegedAccessGroupRelationships) {
    return m.accessId
}
// GetActivatedUsing gets the activatedUsing property value. The activatedUsing property
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetActivatedUsing()(PrivilegedAccessGroupEligibilityScheduleable) {
    return m.activatedUsing
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PrivilegedAccessScheduleRequest.GetFieldDeserializers()
    res["accessId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrivilegedAccessGroupRelationships)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessId(val.(*PrivilegedAccessGroupRelationships))
        }
        return nil
    }
    res["activatedUsing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivatedUsing(val.(PrivilegedAccessGroupEligibilityScheduleable))
        }
        return nil
    }
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(Groupable))
        }
        return nil
    }
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    res["principal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipal(val.(DirectoryObjectable))
        }
        return nil
    }
    res["principalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalId(val)
        }
        return nil
    }
    res["targetSchedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetSchedule(val.(PrivilegedAccessGroupEligibilityScheduleable))
        }
        return nil
    }
    res["targetScheduleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetScheduleId(val)
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group property
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetGroup()(Groupable) {
    return m.group
}
// GetGroupId gets the groupId property value. The groupId property
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetGroupId()(*string) {
    return m.groupId
}
// GetPrincipal gets the principal property value. The principal property
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetPrincipal()(DirectoryObjectable) {
    return m.principal
}
// GetPrincipalId gets the principalId property value. The principalId property
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetPrincipalId()(*string) {
    return m.principalId
}
// GetTargetSchedule gets the targetSchedule property value. The targetSchedule property
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetTargetSchedule()(PrivilegedAccessGroupEligibilityScheduleable) {
    return m.targetSchedule
}
// GetTargetScheduleId gets the targetScheduleId property value. The targetScheduleId property
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetTargetScheduleId()(*string) {
    return m.targetScheduleId
}
// Serialize serializes information the current object
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PrivilegedAccessScheduleRequest.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessId() != nil {
        cast := (*m.GetAccessId()).String()
        err = writer.WriteStringValue("accessId", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("activatedUsing", m.GetActivatedUsing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("principal", m.GetPrincipal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalId", m.GetPrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("targetSchedule", m.GetTargetSchedule())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetScheduleId", m.GetTargetScheduleId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessId sets the accessId property value. The accessId property
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) SetAccessId(value *PrivilegedAccessGroupRelationships)() {
    m.accessId = value
}
// SetActivatedUsing sets the activatedUsing property value. The activatedUsing property
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) SetActivatedUsing(value PrivilegedAccessGroupEligibilityScheduleable)() {
    m.activatedUsing = value
}
// SetGroup sets the group property value. The group property
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) SetGroup(value Groupable)() {
    m.group = value
}
// SetGroupId sets the groupId property value. The groupId property
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) SetGroupId(value *string)() {
    m.groupId = value
}
// SetPrincipal sets the principal property value. The principal property
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) SetPrincipal(value DirectoryObjectable)() {
    m.principal = value
}
// SetPrincipalId sets the principalId property value. The principalId property
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) SetPrincipalId(value *string)() {
    m.principalId = value
}
// SetTargetSchedule sets the targetSchedule property value. The targetSchedule property
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) SetTargetSchedule(value PrivilegedAccessGroupEligibilityScheduleable)() {
    m.targetSchedule = value
}
// SetTargetScheduleId sets the targetScheduleId property value. The targetScheduleId property
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) SetTargetScheduleId(value *string)() {
    m.targetScheduleId = value
}
