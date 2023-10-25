package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AwsIdentityAccessManagementKeyUsageFinding 
type AwsIdentityAccessManagementKeyUsageFinding struct {
    Finding
}
// NewAwsIdentityAccessManagementKeyUsageFinding instantiates a new awsIdentityAccessManagementKeyUsageFinding and sets the default values.
func NewAwsIdentityAccessManagementKeyUsageFinding()(*AwsIdentityAccessManagementKeyUsageFinding) {
    m := &AwsIdentityAccessManagementKeyUsageFinding{
        Finding: *NewFinding(),
    }
    return m
}
// CreateAwsIdentityAccessManagementKeyUsageFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAwsIdentityAccessManagementKeyUsageFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsIdentityAccessManagementKeyUsageFinding(), nil
}
// GetAccessKey gets the accessKey property value. The accessKey property
func (m *AwsIdentityAccessManagementKeyUsageFinding) GetAccessKey()(AwsAccessKeyable) {
    val, err := m.GetBackingStore().Get("accessKey")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AwsAccessKeyable)
    }
    return nil
}
// GetActionSummary gets the actionSummary property value. The actionSummary property
func (m *AwsIdentityAccessManagementKeyUsageFinding) GetActionSummary()(ActionSummaryable) {
    val, err := m.GetBackingStore().Get("actionSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ActionSummaryable)
    }
    return nil
}
// GetAwsAccessKeyDetails gets the awsAccessKeyDetails property value. The awsAccessKeyDetails property
func (m *AwsIdentityAccessManagementKeyUsageFinding) GetAwsAccessKeyDetails()(AwsAccessKeyDetailsable) {
    val, err := m.GetBackingStore().Get("awsAccessKeyDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AwsAccessKeyDetailsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AwsIdentityAccessManagementKeyUsageFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Finding.GetFieldDeserializers()
    res["accessKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAwsAccessKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessKey(val.(AwsAccessKeyable))
        }
        return nil
    }
    res["actionSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateActionSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionSummary(val.(ActionSummaryable))
        }
        return nil
    }
    res["awsAccessKeyDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAwsAccessKeyDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAwsAccessKeyDetails(val.(AwsAccessKeyDetailsable))
        }
        return nil
    }
    res["permissionsCreepIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePermissionsCreepIndexFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionsCreepIndex(val.(PermissionsCreepIndexable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIamStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*IamStatus))
        }
        return nil
    }
    return res
}
// GetPermissionsCreepIndex gets the permissionsCreepIndex property value. The permissionsCreepIndex property
func (m *AwsIdentityAccessManagementKeyUsageFinding) GetPermissionsCreepIndex()(PermissionsCreepIndexable) {
    val, err := m.GetBackingStore().Get("permissionsCreepIndex")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PermissionsCreepIndexable)
    }
    return nil
}
// GetStatus gets the status property value. The status property
func (m *AwsIdentityAccessManagementKeyUsageFinding) GetStatus()(*IamStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IamStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AwsIdentityAccessManagementKeyUsageFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Finding.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessKey", m.GetAccessKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("actionSummary", m.GetActionSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("awsAccessKeyDetails", m.GetAwsAccessKeyDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("permissionsCreepIndex", m.GetPermissionsCreepIndex())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessKey sets the accessKey property value. The accessKey property
func (m *AwsIdentityAccessManagementKeyUsageFinding) SetAccessKey(value AwsAccessKeyable)() {
    err := m.GetBackingStore().Set("accessKey", value)
    if err != nil {
        panic(err)
    }
}
// SetActionSummary sets the actionSummary property value. The actionSummary property
func (m *AwsIdentityAccessManagementKeyUsageFinding) SetActionSummary(value ActionSummaryable)() {
    err := m.GetBackingStore().Set("actionSummary", value)
    if err != nil {
        panic(err)
    }
}
// SetAwsAccessKeyDetails sets the awsAccessKeyDetails property value. The awsAccessKeyDetails property
func (m *AwsIdentityAccessManagementKeyUsageFinding) SetAwsAccessKeyDetails(value AwsAccessKeyDetailsable)() {
    err := m.GetBackingStore().Set("awsAccessKeyDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetPermissionsCreepIndex sets the permissionsCreepIndex property value. The permissionsCreepIndex property
func (m *AwsIdentityAccessManagementKeyUsageFinding) SetPermissionsCreepIndex(value PermissionsCreepIndexable)() {
    err := m.GetBackingStore().Set("permissionsCreepIndex", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *AwsIdentityAccessManagementKeyUsageFinding) SetStatus(value *IamStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// AwsIdentityAccessManagementKeyUsageFindingable 
type AwsIdentityAccessManagementKeyUsageFindingable interface {
    Findingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessKey()(AwsAccessKeyable)
    GetActionSummary()(ActionSummaryable)
    GetAwsAccessKeyDetails()(AwsAccessKeyDetailsable)
    GetPermissionsCreepIndex()(PermissionsCreepIndexable)
    GetStatus()(*IamStatus)
    SetAccessKey(value AwsAccessKeyable)()
    SetActionSummary(value ActionSummaryable)()
    SetAwsAccessKeyDetails(value AwsAccessKeyDetailsable)()
    SetPermissionsCreepIndex(value PermissionsCreepIndexable)()
    SetStatus(value *IamStatus)()
}
