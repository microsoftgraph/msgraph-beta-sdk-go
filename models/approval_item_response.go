package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApprovalItemResponse struct {
    Entity
}
// NewApprovalItemResponse instantiates a new ApprovalItemResponse and sets the default values.
func NewApprovalItemResponse()(*ApprovalItemResponse) {
    m := &ApprovalItemResponse{
        Entity: *NewEntity(),
    }
    return m
}
// CreateApprovalItemResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApprovalItemResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApprovalItemResponse(), nil
}
// GetComments gets the comments property value. The comment made by the approver.
// returns a *string when successful
func (m *ApprovalItemResponse) GetComments()(*string) {
    val, err := m.GetBackingStore().Get("comments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. The identity set of the approver.
// returns a ApprovalIdentitySetable when successful
func (m *ApprovalItemResponse) GetCreatedBy()(ApprovalIdentitySetable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ApprovalIdentitySetable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Creation date and time of the response.
// returns a *Time when successful
func (m *ApprovalItemResponse) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
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
func (m *ApprovalItemResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComments(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApprovalIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(ApprovalIdentitySetable))
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
    res["owners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApprovalIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApprovalIdentitySetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ApprovalIdentitySetable)
                }
            }
            m.SetOwners(res)
        }
        return nil
    }
    res["response"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponse(val)
        }
        return nil
    }
    return res
}
// GetOwners gets the owners property value. The identity set of the principal who owns the approval item.
// returns a []ApprovalIdentitySetable when successful
func (m *ApprovalItemResponse) GetOwners()([]ApprovalIdentitySetable) {
    val, err := m.GetBackingStore().Get("owners")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ApprovalIdentitySetable)
    }
    return nil
}
// GetResponse gets the response property value. Approver response based on the response options. The default response options are 'Approved' and 'Rejected'. The approval item creator can also define custom response options during approval item creation.
// returns a *string when successful
func (m *ApprovalItemResponse) GetResponse()(*string) {
    val, err := m.GetBackingStore().Get("response")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ApprovalItemResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("comments", m.GetComments())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("response", m.GetResponse())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetComments sets the comments property value. The comment made by the approver.
func (m *ApprovalItemResponse) SetComments(value *string)() {
    err := m.GetBackingStore().Set("comments", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. The identity set of the approver.
func (m *ApprovalItemResponse) SetCreatedBy(value ApprovalIdentitySetable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Creation date and time of the response.
func (m *ApprovalItemResponse) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOwners sets the owners property value. The identity set of the principal who owns the approval item.
func (m *ApprovalItemResponse) SetOwners(value []ApprovalIdentitySetable)() {
    err := m.GetBackingStore().Set("owners", value)
    if err != nil {
        panic(err)
    }
}
// SetResponse sets the response property value. Approver response based on the response options. The default response options are 'Approved' and 'Rejected'. The approval item creator can also define custom response options during approval item creation.
func (m *ApprovalItemResponse) SetResponse(value *string)() {
    err := m.GetBackingStore().Set("response", value)
    if err != nil {
        panic(err)
    }
}
type ApprovalItemResponseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComments()(*string)
    GetCreatedBy()(ApprovalIdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOwners()([]ApprovalIdentitySetable)
    GetResponse()(*string)
    SetComments(value *string)()
    SetCreatedBy(value ApprovalIdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOwners(value []ApprovalIdentitySetable)()
    SetResponse(value *string)()
}
