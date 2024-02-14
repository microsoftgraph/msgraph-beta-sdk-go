package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AssignedComputeInstanceDetails struct {
    Entity
}
// NewAssignedComputeInstanceDetails instantiates a new AssignedComputeInstanceDetails and sets the default values.
func NewAssignedComputeInstanceDetails()(*AssignedComputeInstanceDetails) {
    m := &AssignedComputeInstanceDetails{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAssignedComputeInstanceDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAssignedComputeInstanceDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignedComputeInstanceDetails(), nil
}
// GetAccessedStorageBuckets gets the accessedStorageBuckets property value. Represents a set of S3 buckets accessed by this EC2 instance.
// returns a []AuthorizationSystemResourceable when successful
func (m *AssignedComputeInstanceDetails) GetAccessedStorageBuckets()([]AuthorizationSystemResourceable) {
    val, err := m.GetBackingStore().Get("accessedStorageBuckets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthorizationSystemResourceable)
    }
    return nil
}
// GetAssignedComputeInstance gets the assignedComputeInstance property value. assigned EC2 instance.
// returns a AuthorizationSystemResourceable when successful
func (m *AssignedComputeInstanceDetails) GetAssignedComputeInstance()(AuthorizationSystemResourceable) {
    val, err := m.GetBackingStore().Get("assignedComputeInstance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthorizationSystemResourceable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AssignedComputeInstanceDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessedStorageBuckets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthorizationSystemResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthorizationSystemResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthorizationSystemResourceable)
                }
            }
            m.SetAccessedStorageBuckets(res)
        }
        return nil
    }
    res["assignedComputeInstance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationSystemResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedComputeInstance(val.(AuthorizationSystemResourceable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AssignedComputeInstanceDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessedStorageBuckets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessedStorageBuckets()))
        for i, v := range m.GetAccessedStorageBuckets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("accessedStorageBuckets", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assignedComputeInstance", m.GetAssignedComputeInstance())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessedStorageBuckets sets the accessedStorageBuckets property value. Represents a set of S3 buckets accessed by this EC2 instance.
func (m *AssignedComputeInstanceDetails) SetAccessedStorageBuckets(value []AuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("accessedStorageBuckets", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignedComputeInstance sets the assignedComputeInstance property value. assigned EC2 instance.
func (m *AssignedComputeInstanceDetails) SetAssignedComputeInstance(value AuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("assignedComputeInstance", value)
    if err != nil {
        panic(err)
    }
}
type AssignedComputeInstanceDetailsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessedStorageBuckets()([]AuthorizationSystemResourceable)
    GetAssignedComputeInstance()(AuthorizationSystemResourceable)
    SetAccessedStorageBuckets(value []AuthorizationSystemResourceable)()
    SetAssignedComputeInstance(value AuthorizationSystemResourceable)()
}
