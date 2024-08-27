package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApprovalItemRequest struct {
    Entity
}
// NewApprovalItemRequest instantiates a new ApprovalItemRequest and sets the default values.
func NewApprovalItemRequest()(*ApprovalItemRequest) {
    m := &ApprovalItemRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateApprovalItemRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApprovalItemRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApprovalItemRequest(), nil
}
// GetApprover gets the approver property value. The identity set of the principal assigned to this request.
// returns a ApprovalIdentitySetable when successful
func (m *ApprovalItemRequest) GetApprover()(ApprovalIdentitySetable) {
    val, err := m.GetBackingStore().Get("approver")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ApprovalIdentitySetable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Creation date and time for the request.
// returns a *Time when successful
func (m *ApprovalItemRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApprovalItemRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["approver"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApprovalIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprover(val.(ApprovalIdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["isReassigned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReassigned(val)
        }
        return nil
    }
    res["reassignedFrom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApprovalIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReassignedFrom(val.(ApprovalIdentitySetable))
        }
        return nil
    }
    return res
}
// GetIsReassigned gets the isReassigned property value. Indicates whether a request was reassigned.
// returns a *bool when successful
func (m *ApprovalItemRequest) GetIsReassigned()(*bool) {
    val, err := m.GetBackingStore().Get("isReassigned")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetReassignedFrom gets the reassignedFrom property value. The identity set of the principal who reassigned the request.
// returns a ApprovalIdentitySetable when successful
func (m *ApprovalItemRequest) GetReassignedFrom()(ApprovalIdentitySetable) {
    val, err := m.GetBackingStore().Get("reassignedFrom")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ApprovalIdentitySetable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ApprovalItemRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SetApprover sets the approver property value. The identity set of the principal assigned to this request.
func (m *ApprovalItemRequest) SetApprover(value ApprovalIdentitySetable)() {
    err := m.GetBackingStore().Set("approver", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Creation date and time for the request.
func (m *ApprovalItemRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIsReassigned sets the isReassigned property value. Indicates whether a request was reassigned.
func (m *ApprovalItemRequest) SetIsReassigned(value *bool)() {
    err := m.GetBackingStore().Set("isReassigned", value)
    if err != nil {
        panic(err)
    }
}
// SetReassignedFrom sets the reassignedFrom property value. The identity set of the principal who reassigned the request.
func (m *ApprovalItemRequest) SetReassignedFrom(value ApprovalIdentitySetable)() {
    err := m.GetBackingStore().Set("reassignedFrom", value)
    if err != nil {
        panic(err)
    }
}
type ApprovalItemRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApprover()(ApprovalIdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsReassigned()(*bool)
    GetReassignedFrom()(ApprovalIdentitySetable)
    SetApprover(value ApprovalIdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsReassigned(value *bool)()
    SetReassignedFrom(value ApprovalIdentitySetable)()
}
