package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApprovalSolution struct {
    Entity
}
// NewApprovalSolution instantiates a new ApprovalSolution and sets the default values.
func NewApprovalSolution()(*ApprovalSolution) {
    m := &ApprovalSolution{
        Entity: *NewEntity(),
    }
    return m
}
// CreateApprovalSolutionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApprovalSolutionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApprovalSolution(), nil
}
// GetApprovalItems gets the approvalItems property value. A collection of approval items.
// returns a []ApprovalItemable when successful
func (m *ApprovalSolution) GetApprovalItems()([]ApprovalItemable) {
    val, err := m.GetBackingStore().Get("approvalItems")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ApprovalItemable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApprovalSolution) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["approvalItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApprovalItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApprovalItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ApprovalItemable)
                }
            }
            m.SetApprovalItems(res)
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApprovalOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApprovalOperationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ApprovalOperationable)
                }
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["provisioningStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisionState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningStatus(val.(*ProvisionState))
        }
        return nil
    }
    return res
}
// GetOperations gets the operations property value. The operations property
// returns a []ApprovalOperationable when successful
func (m *ApprovalSolution) GetOperations()([]ApprovalOperationable) {
    val, err := m.GetBackingStore().Get("operations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ApprovalOperationable)
    }
    return nil
}
// GetProvisioningStatus gets the provisioningStatus property value. The approval provisioning status for a tenant on an environment. The possible values are: notProvisioned, provisioningInProgress, provisioningFailed, provisioningCompleted, unknownFutureValue.
// returns a *ProvisionState when successful
func (m *ApprovalSolution) GetProvisioningStatus()(*ProvisionState) {
    val, err := m.GetBackingStore().Get("provisioningStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ProvisionState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ApprovalSolution) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApprovalItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApprovalItems()))
        for i, v := range m.GetApprovalItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("approvalItems", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProvisioningStatus() != nil {
        cast := (*m.GetProvisioningStatus()).String()
        err = writer.WriteStringValue("provisioningStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApprovalItems sets the approvalItems property value. A collection of approval items.
func (m *ApprovalSolution) SetApprovalItems(value []ApprovalItemable)() {
    err := m.GetBackingStore().Set("approvalItems", value)
    if err != nil {
        panic(err)
    }
}
// SetOperations sets the operations property value. The operations property
func (m *ApprovalSolution) SetOperations(value []ApprovalOperationable)() {
    err := m.GetBackingStore().Set("operations", value)
    if err != nil {
        panic(err)
    }
}
// SetProvisioningStatus sets the provisioningStatus property value. The approval provisioning status for a tenant on an environment. The possible values are: notProvisioned, provisioningInProgress, provisioningFailed, provisioningCompleted, unknownFutureValue.
func (m *ApprovalSolution) SetProvisioningStatus(value *ProvisionState)() {
    err := m.GetBackingStore().Set("provisioningStatus", value)
    if err != nil {
        panic(err)
    }
}
type ApprovalSolutionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApprovalItems()([]ApprovalItemable)
    GetOperations()([]ApprovalOperationable)
    GetProvisioningStatus()(*ProvisionState)
    SetApprovalItems(value []ApprovalItemable)()
    SetOperations(value []ApprovalOperationable)()
    SetProvisioningStatus(value *ProvisionState)()
}
