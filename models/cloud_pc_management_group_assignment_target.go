package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcManagementGroupAssignmentTarget struct {
    CloudPcManagementAssignmentTarget
}
// NewCloudPcManagementGroupAssignmentTarget instantiates a new CloudPcManagementGroupAssignmentTarget and sets the default values.
func NewCloudPcManagementGroupAssignmentTarget()(*CloudPcManagementGroupAssignmentTarget) {
    m := &CloudPcManagementGroupAssignmentTarget{
        CloudPcManagementAssignmentTarget: *NewCloudPcManagementAssignmentTarget(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcManagementGroupAssignmentTarget"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCloudPcManagementGroupAssignmentTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcManagementGroupAssignmentTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcManagementGroupAssignmentTarget(), nil
}
// GetAllotmentDisplayName gets the allotmentDisplayName property value. The allotmentDisplayName property
// returns a *string when successful
func (m *CloudPcManagementGroupAssignmentTarget) GetAllotmentDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("allotmentDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAllotmentLicensesCount gets the allotmentLicensesCount property value. The allotmentLicensesCount property
// returns a *int32 when successful
func (m *CloudPcManagementGroupAssignmentTarget) GetAllotmentLicensesCount()(*int32) {
    val, err := m.GetBackingStore().Get("allotmentLicensesCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcManagementGroupAssignmentTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CloudPcManagementAssignmentTarget.GetFieldDeserializers()
    res["allotmentDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllotmentDisplayName(val)
        }
        return nil
    }
    res["allotmentLicensesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllotmentLicensesCount(val)
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
    res["servicePlanId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanId(val)
        }
        return nil
    }
    return res
}
// GetGroupId gets the groupId property value. The ID of the target group for the assignment.
// returns a *string when successful
func (m *CloudPcManagementGroupAssignmentTarget) GetGroupId()(*string) {
    val, err := m.GetBackingStore().Get("groupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetServicePlanId gets the servicePlanId property value. The unique identifier for the service plan that indicates which size of the Cloud PC to provision for the user. Use a null value, when the provisioningType is dedicated.
// returns a *string when successful
func (m *CloudPcManagementGroupAssignmentTarget) GetServicePlanId()(*string) {
    val, err := m.GetBackingStore().Get("servicePlanId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcManagementGroupAssignmentTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CloudPcManagementAssignmentTarget.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("allotmentDisplayName", m.GetAllotmentDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("allotmentLicensesCount", m.GetAllotmentLicensesCount())
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
        err = writer.WriteStringValue("servicePlanId", m.GetServicePlanId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllotmentDisplayName sets the allotmentDisplayName property value. The allotmentDisplayName property
func (m *CloudPcManagementGroupAssignmentTarget) SetAllotmentDisplayName(value *string)() {
    err := m.GetBackingStore().Set("allotmentDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetAllotmentLicensesCount sets the allotmentLicensesCount property value. The allotmentLicensesCount property
func (m *CloudPcManagementGroupAssignmentTarget) SetAllotmentLicensesCount(value *int32)() {
    err := m.GetBackingStore().Set("allotmentLicensesCount", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupId sets the groupId property value. The ID of the target group for the assignment.
func (m *CloudPcManagementGroupAssignmentTarget) SetGroupId(value *string)() {
    err := m.GetBackingStore().Set("groupId", value)
    if err != nil {
        panic(err)
    }
}
// SetServicePlanId sets the servicePlanId property value. The unique identifier for the service plan that indicates which size of the Cloud PC to provision for the user. Use a null value, when the provisioningType is dedicated.
func (m *CloudPcManagementGroupAssignmentTarget) SetServicePlanId(value *string)() {
    err := m.GetBackingStore().Set("servicePlanId", value)
    if err != nil {
        panic(err)
    }
}
type CloudPcManagementGroupAssignmentTargetable interface {
    CloudPcManagementAssignmentTargetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllotmentDisplayName()(*string)
    GetAllotmentLicensesCount()(*int32)
    GetGroupId()(*string)
    GetServicePlanId()(*string)
    SetAllotmentDisplayName(value *string)()
    SetAllotmentLicensesCount(value *int32)()
    SetGroupId(value *string)()
    SetServicePlanId(value *string)()
}
