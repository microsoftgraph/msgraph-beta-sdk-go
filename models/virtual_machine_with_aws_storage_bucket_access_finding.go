package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualMachineWithAwsStorageBucketAccessFinding 
type VirtualMachineWithAwsStorageBucketAccessFinding struct {
    Finding
}
// NewVirtualMachineWithAwsStorageBucketAccessFinding instantiates a new virtualMachineWithAwsStorageBucketAccessFinding and sets the default values.
func NewVirtualMachineWithAwsStorageBucketAccessFinding()(*VirtualMachineWithAwsStorageBucketAccessFinding) {
    m := &VirtualMachineWithAwsStorageBucketAccessFinding{
        Finding: *NewFinding(),
    }
    return m
}
// CreateVirtualMachineWithAwsStorageBucketAccessFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualMachineWithAwsStorageBucketAccessFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualMachineWithAwsStorageBucketAccessFinding(), nil
}
// GetAccessibleCount gets the accessibleCount property value. The total number of storage buckets that the EC2 instance can access using the role
func (m *VirtualMachineWithAwsStorageBucketAccessFinding) GetAccessibleCount()(*int32) {
    val, err := m.GetBackingStore().Get("accessibleCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetBucketCount gets the bucketCount property value. The total number of storage buckets in the authorization system that host the EC2 instance
func (m *VirtualMachineWithAwsStorageBucketAccessFinding) GetBucketCount()(*int32) {
    val, err := m.GetBackingStore().Get("bucketCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetEc2Instance gets the ec2Instance property value. The ec2Instance property
func (m *VirtualMachineWithAwsStorageBucketAccessFinding) GetEc2Instance()(AuthorizationSystemResourceable) {
    val, err := m.GetBackingStore().Get("ec2Instance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthorizationSystemResourceable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualMachineWithAwsStorageBucketAccessFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Finding.GetFieldDeserializers()
    res["accessibleCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessibleCount(val)
        }
        return nil
    }
    res["bucketCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBucketCount(val)
        }
        return nil
    }
    res["ec2Instance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationSystemResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEc2Instance(val.(AuthorizationSystemResourceable))
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
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAwsRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(AwsRoleable))
        }
        return nil
    }
    return res
}
// GetPermissionsCreepIndex gets the permissionsCreepIndex property value. The permissionsCreepIndex property
func (m *VirtualMachineWithAwsStorageBucketAccessFinding) GetPermissionsCreepIndex()(PermissionsCreepIndexable) {
    val, err := m.GetBackingStore().Get("permissionsCreepIndex")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PermissionsCreepIndexable)
    }
    return nil
}
// GetRole gets the role property value. The role property
func (m *VirtualMachineWithAwsStorageBucketAccessFinding) GetRole()(AwsRoleable) {
    val, err := m.GetBackingStore().Get("role")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AwsRoleable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VirtualMachineWithAwsStorageBucketAccessFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Finding.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("accessibleCount", m.GetAccessibleCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("bucketCount", m.GetBucketCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("ec2Instance", m.GetEc2Instance())
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
    {
        err = writer.WriteObjectValue("role", m.GetRole())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessibleCount sets the accessibleCount property value. The total number of storage buckets that the EC2 instance can access using the role
func (m *VirtualMachineWithAwsStorageBucketAccessFinding) SetAccessibleCount(value *int32)() {
    err := m.GetBackingStore().Set("accessibleCount", value)
    if err != nil {
        panic(err)
    }
}
// SetBucketCount sets the bucketCount property value. The total number of storage buckets in the authorization system that host the EC2 instance
func (m *VirtualMachineWithAwsStorageBucketAccessFinding) SetBucketCount(value *int32)() {
    err := m.GetBackingStore().Set("bucketCount", value)
    if err != nil {
        panic(err)
    }
}
// SetEc2Instance sets the ec2Instance property value. The ec2Instance property
func (m *VirtualMachineWithAwsStorageBucketAccessFinding) SetEc2Instance(value AuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("ec2Instance", value)
    if err != nil {
        panic(err)
    }
}
// SetPermissionsCreepIndex sets the permissionsCreepIndex property value. The permissionsCreepIndex property
func (m *VirtualMachineWithAwsStorageBucketAccessFinding) SetPermissionsCreepIndex(value PermissionsCreepIndexable)() {
    err := m.GetBackingStore().Set("permissionsCreepIndex", value)
    if err != nil {
        panic(err)
    }
}
// SetRole sets the role property value. The role property
func (m *VirtualMachineWithAwsStorageBucketAccessFinding) SetRole(value AwsRoleable)() {
    err := m.GetBackingStore().Set("role", value)
    if err != nil {
        panic(err)
    }
}
// VirtualMachineWithAwsStorageBucketAccessFindingable 
type VirtualMachineWithAwsStorageBucketAccessFindingable interface {
    Findingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessibleCount()(*int32)
    GetBucketCount()(*int32)
    GetEc2Instance()(AuthorizationSystemResourceable)
    GetPermissionsCreepIndex()(PermissionsCreepIndexable)
    GetRole()(AwsRoleable)
    SetAccessibleCount(value *int32)()
    SetBucketCount(value *int32)()
    SetEc2Instance(value AuthorizationSystemResourceable)()
    SetPermissionsCreepIndex(value PermissionsCreepIndexable)()
    SetRole(value AwsRoleable)()
}
