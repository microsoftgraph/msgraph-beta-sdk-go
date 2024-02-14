package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OpenAwsSecurityGroupFinding struct {
    Finding
}
// NewOpenAwsSecurityGroupFinding instantiates a new OpenAwsSecurityGroupFinding and sets the default values.
func NewOpenAwsSecurityGroupFinding()(*OpenAwsSecurityGroupFinding) {
    m := &OpenAwsSecurityGroupFinding{
        Finding: *NewFinding(),
    }
    return m
}
// CreateOpenAwsSecurityGroupFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOpenAwsSecurityGroupFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOpenAwsSecurityGroupFinding(), nil
}
// GetAssignedComputeInstancesDetails gets the assignedComputeInstancesDetails property value. A set of AWS EC2 compute instances related to this open security group.
// returns a []AssignedComputeInstanceDetailsable when successful
func (m *OpenAwsSecurityGroupFinding) GetAssignedComputeInstancesDetails()([]AssignedComputeInstanceDetailsable) {
    val, err := m.GetBackingStore().Get("assignedComputeInstancesDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AssignedComputeInstanceDetailsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OpenAwsSecurityGroupFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Finding.GetFieldDeserializers()
    res["assignedComputeInstancesDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignedComputeInstanceDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedComputeInstanceDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AssignedComputeInstanceDetailsable)
                }
            }
            m.SetAssignedComputeInstancesDetails(res)
        }
        return nil
    }
    res["inboundPorts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInboundPortsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInboundPorts(val.(InboundPortsable))
        }
        return nil
    }
    res["securityGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAwsAuthorizationSystemResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityGroup(val.(AwsAuthorizationSystemResourceable))
        }
        return nil
    }
    res["totalStorageBucketCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalStorageBucketCount(val)
        }
        return nil
    }
    return res
}
// GetInboundPorts gets the inboundPorts property value. The inboundPorts property
// returns a InboundPortsable when successful
func (m *OpenAwsSecurityGroupFinding) GetInboundPorts()(InboundPortsable) {
    val, err := m.GetBackingStore().Get("inboundPorts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(InboundPortsable)
    }
    return nil
}
// GetSecurityGroup gets the securityGroup property value. The securityGroup property
// returns a AwsAuthorizationSystemResourceable when successful
func (m *OpenAwsSecurityGroupFinding) GetSecurityGroup()(AwsAuthorizationSystemResourceable) {
    val, err := m.GetBackingStore().Get("securityGroup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AwsAuthorizationSystemResourceable)
    }
    return nil
}
// GetTotalStorageBucketCount gets the totalStorageBucketCount property value. The number of storage buckets accessed by the assigned compute instances.
// returns a *int32 when successful
func (m *OpenAwsSecurityGroupFinding) GetTotalStorageBucketCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalStorageBucketCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OpenAwsSecurityGroupFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Finding.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignedComputeInstancesDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedComputeInstancesDetails()))
        for i, v := range m.GetAssignedComputeInstancesDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignedComputeInstancesDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("inboundPorts", m.GetInboundPorts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("securityGroup", m.GetSecurityGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalStorageBucketCount", m.GetTotalStorageBucketCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedComputeInstancesDetails sets the assignedComputeInstancesDetails property value. A set of AWS EC2 compute instances related to this open security group.
func (m *OpenAwsSecurityGroupFinding) SetAssignedComputeInstancesDetails(value []AssignedComputeInstanceDetailsable)() {
    err := m.GetBackingStore().Set("assignedComputeInstancesDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetInboundPorts sets the inboundPorts property value. The inboundPorts property
func (m *OpenAwsSecurityGroupFinding) SetInboundPorts(value InboundPortsable)() {
    err := m.GetBackingStore().Set("inboundPorts", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityGroup sets the securityGroup property value. The securityGroup property
func (m *OpenAwsSecurityGroupFinding) SetSecurityGroup(value AwsAuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("securityGroup", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalStorageBucketCount sets the totalStorageBucketCount property value. The number of storage buckets accessed by the assigned compute instances.
func (m *OpenAwsSecurityGroupFinding) SetTotalStorageBucketCount(value *int32)() {
    err := m.GetBackingStore().Set("totalStorageBucketCount", value)
    if err != nil {
        panic(err)
    }
}
type OpenAwsSecurityGroupFindingable interface {
    Findingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedComputeInstancesDetails()([]AssignedComputeInstanceDetailsable)
    GetInboundPorts()(InboundPortsable)
    GetSecurityGroup()(AwsAuthorizationSystemResourceable)
    GetTotalStorageBucketCount()(*int32)
    SetAssignedComputeInstancesDetails(value []AssignedComputeInstanceDetailsable)()
    SetInboundPorts(value InboundPortsable)()
    SetSecurityGroup(value AwsAuthorizationSystemResourceable)()
    SetTotalStorageBucketCount(value *int32)()
}
