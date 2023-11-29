package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OpenNetworkAzureSecurityGroupFinding 
type OpenNetworkAzureSecurityGroupFinding struct {
    Finding
}
// NewOpenNetworkAzureSecurityGroupFinding instantiates a new openNetworkAzureSecurityGroupFinding and sets the default values.
func NewOpenNetworkAzureSecurityGroupFinding()(*OpenNetworkAzureSecurityGroupFinding) {
    m := &OpenNetworkAzureSecurityGroupFinding{
        Finding: *NewFinding(),
    }
    return m
}
// CreateOpenNetworkAzureSecurityGroupFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOpenNetworkAzureSecurityGroupFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOpenNetworkAzureSecurityGroupFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OpenNetworkAzureSecurityGroupFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Finding.GetFieldDeserializers()
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
        val, err := n.GetObjectValue(CreateAuthorizationSystemResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityGroup(val.(AuthorizationSystemResourceable))
        }
        return nil
    }
    res["virtualMachines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVirtualMachineDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VirtualMachineDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VirtualMachineDetailsable)
                }
            }
            m.SetVirtualMachines(res)
        }
        return nil
    }
    return res
}
// GetInboundPorts gets the inboundPorts property value. The inboundPorts property
func (m *OpenNetworkAzureSecurityGroupFinding) GetInboundPorts()(InboundPortsable) {
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
func (m *OpenNetworkAzureSecurityGroupFinding) GetSecurityGroup()(AuthorizationSystemResourceable) {
    val, err := m.GetBackingStore().Get("securityGroup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthorizationSystemResourceable)
    }
    return nil
}
// GetVirtualMachines gets the virtualMachines property value. Represents a virtual machine in an authorization system.
func (m *OpenNetworkAzureSecurityGroupFinding) GetVirtualMachines()([]VirtualMachineDetailsable) {
    val, err := m.GetBackingStore().Get("virtualMachines")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VirtualMachineDetailsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OpenNetworkAzureSecurityGroupFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Finding.Serialize(writer)
    if err != nil {
        return err
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
    if m.GetVirtualMachines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVirtualMachines()))
        for i, v := range m.GetVirtualMachines() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("virtualMachines", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInboundPorts sets the inboundPorts property value. The inboundPorts property
func (m *OpenNetworkAzureSecurityGroupFinding) SetInboundPorts(value InboundPortsable)() {
    err := m.GetBackingStore().Set("inboundPorts", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityGroup sets the securityGroup property value. The securityGroup property
func (m *OpenNetworkAzureSecurityGroupFinding) SetSecurityGroup(value AuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("securityGroup", value)
    if err != nil {
        panic(err)
    }
}
// SetVirtualMachines sets the virtualMachines property value. Represents a virtual machine in an authorization system.
func (m *OpenNetworkAzureSecurityGroupFinding) SetVirtualMachines(value []VirtualMachineDetailsable)() {
    err := m.GetBackingStore().Set("virtualMachines", value)
    if err != nil {
        panic(err)
    }
}
// OpenNetworkAzureSecurityGroupFindingable 
type OpenNetworkAzureSecurityGroupFindingable interface {
    Findingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInboundPorts()(InboundPortsable)
    GetSecurityGroup()(AuthorizationSystemResourceable)
    GetVirtualMachines()([]VirtualMachineDetailsable)
    SetInboundPorts(value InboundPortsable)()
    SetSecurityGroup(value AuthorizationSystemResourceable)()
    SetVirtualMachines(value []VirtualMachineDetailsable)()
}
