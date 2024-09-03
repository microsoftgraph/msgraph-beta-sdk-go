package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApprovalItem struct {
    Entity
}
// NewApprovalItem instantiates a new ApprovalItem and sets the default values.
func NewApprovalItem()(*ApprovalItem) {
    m := &ApprovalItem{
        Entity: *NewEntity(),
    }
    return m
}
// CreateApprovalItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApprovalItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApprovalItem(), nil
}
// GetAllowCancel gets the allowCancel property value. Indicates whether the approval item can be canceled.
// returns a *bool when successful
func (m *ApprovalItem) GetAllowCancel()(*bool) {
    val, err := m.GetBackingStore().Get("allowCancel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowEmailNotification gets the allowEmailNotification property value. Indicates whether email notification is enabled.
// returns a *bool when successful
func (m *ApprovalItem) GetAllowEmailNotification()(*bool) {
    val, err := m.GetBackingStore().Get("allowEmailNotification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetApprovalType gets the approvalType property value. The workflow type of the approval item. The possible values are: basic, basicAwaitAll, custom, customAwaitAll. Required.
// returns a *ApprovalItemType when successful
func (m *ApprovalItem) GetApprovalType()(*ApprovalItemType) {
    val, err := m.GetBackingStore().Get("approvalType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ApprovalItemType)
    }
    return nil
}
// GetApprovers gets the approvers property value. The identity of the principals to whom the approval item was initially assigned. Required.
// returns a []ApprovalIdentitySetable when successful
func (m *ApprovalItem) GetApprovers()([]ApprovalIdentitySetable) {
    val, err := m.GetBackingStore().Get("approvers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ApprovalIdentitySetable)
    }
    return nil
}
// GetCompletedDateTime gets the completedDateTime property value. Approval request completion date and time. Read-only.
// returns a *Time when successful
func (m *ApprovalItem) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("completedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Creation date and time of the approval request. Read-only.
// returns a *Time when successful
func (m *ApprovalItem) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDescription gets the description property value. The description of the approval request.
// returns a *string when successful
func (m *ApprovalItem) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName of the approval request. Required.
// returns a *string when successful
func (m *ApprovalItem) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApprovalItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowCancel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowCancel(val)
        }
        return nil
    }
    res["allowEmailNotification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowEmailNotification(val)
        }
        return nil
    }
    res["approvalType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApprovalItemType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalType(val.(*ApprovalItemType))
        }
        return nil
    }
    res["approvers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetApprovers(res)
        }
        return nil
    }
    res["completedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApprovalIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(ApprovalIdentitySetable))
        }
        return nil
    }
    res["requests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApprovalItemRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApprovalItemRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ApprovalItemRequestable)
                }
            }
            m.SetRequests(res)
        }
        return nil
    }
    res["responsePrompts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetResponsePrompts(res)
        }
        return nil
    }
    res["responses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApprovalItemResponseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApprovalItemResponseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ApprovalItemResponseable)
                }
            }
            m.SetResponses(res)
        }
        return nil
    }
    res["result"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApprovalItemState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ApprovalItemState))
        }
        return nil
    }
    res["viewPoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApprovalItemViewPointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewPoint(val.(ApprovalItemViewPointable))
        }
        return nil
    }
    return res
}
// GetOwner gets the owner property value. The identity set of the principal who owns the approval item. Only provide a value for this property when creating an approval item on behalf of the principal. If the owner field isn't provided, the user information from the user context is used.
// returns a ApprovalIdentitySetable when successful
func (m *ApprovalItem) GetOwner()(ApprovalIdentitySetable) {
    val, err := m.GetBackingStore().Get("owner")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ApprovalIdentitySetable)
    }
    return nil
}
// GetRequests gets the requests property value. A collection of requests created for each approver on the approval item.
// returns a []ApprovalItemRequestable when successful
func (m *ApprovalItem) GetRequests()([]ApprovalItemRequestable) {
    val, err := m.GetBackingStore().Get("requests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ApprovalItemRequestable)
    }
    return nil
}
// GetResponsePrompts gets the responsePrompts property value. Approval response prompts. Only provide a value for this property when creating a custom approval item. For custom approval items, supply two response prompt strings. The default response prompts are 'Approve' and 'Reject'.
// returns a []string when successful
func (m *ApprovalItem) GetResponsePrompts()([]string) {
    val, err := m.GetBackingStore().Get("responsePrompts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetResponses gets the responses property value. A collection of responses created for the approval item.
// returns a []ApprovalItemResponseable when successful
func (m *ApprovalItem) GetResponses()([]ApprovalItemResponseable) {
    val, err := m.GetBackingStore().Get("responses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ApprovalItemResponseable)
    }
    return nil
}
// GetResult gets the result property value. The result field is only populated once the approval item is in its final state. The result of the approval item is based on the approvalType. For basic approval items, the result is either 'Approved' or 'Rejected'. For custom approval items, the result could either be a single response or multiple responses separated by a semi-colon. Read-only.
// returns a *string when successful
func (m *ApprovalItem) GetResult()(*string) {
    val, err := m.GetBackingStore().Get("result")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetState gets the state property value. The approval item state. The possible values are: canceled, created, pending, completed. Read-only.
// returns a *ApprovalItemState when successful
func (m *ApprovalItem) GetState()(*ApprovalItemState) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ApprovalItemState)
    }
    return nil
}
// GetViewPoint gets the viewPoint property value. Represents user viewpoints data on the ApprovalItem. The data includes the users roles regarding the approval item. Read-only.
// returns a ApprovalItemViewPointable when successful
func (m *ApprovalItem) GetViewPoint()(ApprovalItemViewPointable) {
    val, err := m.GetBackingStore().Get("viewPoint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ApprovalItemViewPointable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ApprovalItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowEmailNotification", m.GetAllowEmailNotification())
        if err != nil {
            return err
        }
    }
    if m.GetApprovalType() != nil {
        cast := (*m.GetApprovalType()).String()
        err = writer.WriteStringValue("approvalType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetApprovers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApprovers()))
        for i, v := range m.GetApprovers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("approvers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRequests()))
        for i, v := range m.GetRequests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("requests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResponsePrompts() != nil {
        err = writer.WriteCollectionOfStringValues("responsePrompts", m.GetResponsePrompts())
        if err != nil {
            return err
        }
    }
    if m.GetResponses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResponses()))
        for i, v := range m.GetResponses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("responses", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowCancel sets the allowCancel property value. Indicates whether the approval item can be canceled.
func (m *ApprovalItem) SetAllowCancel(value *bool)() {
    err := m.GetBackingStore().Set("allowCancel", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowEmailNotification sets the allowEmailNotification property value. Indicates whether email notification is enabled.
func (m *ApprovalItem) SetAllowEmailNotification(value *bool)() {
    err := m.GetBackingStore().Set("allowEmailNotification", value)
    if err != nil {
        panic(err)
    }
}
// SetApprovalType sets the approvalType property value. The workflow type of the approval item. The possible values are: basic, basicAwaitAll, custom, customAwaitAll. Required.
func (m *ApprovalItem) SetApprovalType(value *ApprovalItemType)() {
    err := m.GetBackingStore().Set("approvalType", value)
    if err != nil {
        panic(err)
    }
}
// SetApprovers sets the approvers property value. The identity of the principals to whom the approval item was initially assigned. Required.
func (m *ApprovalItem) SetApprovers(value []ApprovalIdentitySetable)() {
    err := m.GetBackingStore().Set("approvers", value)
    if err != nil {
        panic(err)
    }
}
// SetCompletedDateTime sets the completedDateTime property value. Approval request completion date and time. Read-only.
func (m *ApprovalItem) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("completedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Creation date and time of the approval request. Read-only.
func (m *ApprovalItem) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description of the approval request.
func (m *ApprovalItem) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName of the approval request. Required.
func (m *ApprovalItem) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetOwner sets the owner property value. The identity set of the principal who owns the approval item. Only provide a value for this property when creating an approval item on behalf of the principal. If the owner field isn't provided, the user information from the user context is used.
func (m *ApprovalItem) SetOwner(value ApprovalIdentitySetable)() {
    err := m.GetBackingStore().Set("owner", value)
    if err != nil {
        panic(err)
    }
}
// SetRequests sets the requests property value. A collection of requests created for each approver on the approval item.
func (m *ApprovalItem) SetRequests(value []ApprovalItemRequestable)() {
    err := m.GetBackingStore().Set("requests", value)
    if err != nil {
        panic(err)
    }
}
// SetResponsePrompts sets the responsePrompts property value. Approval response prompts. Only provide a value for this property when creating a custom approval item. For custom approval items, supply two response prompt strings. The default response prompts are 'Approve' and 'Reject'.
func (m *ApprovalItem) SetResponsePrompts(value []string)() {
    err := m.GetBackingStore().Set("responsePrompts", value)
    if err != nil {
        panic(err)
    }
}
// SetResponses sets the responses property value. A collection of responses created for the approval item.
func (m *ApprovalItem) SetResponses(value []ApprovalItemResponseable)() {
    err := m.GetBackingStore().Set("responses", value)
    if err != nil {
        panic(err)
    }
}
// SetResult sets the result property value. The result field is only populated once the approval item is in its final state. The result of the approval item is based on the approvalType. For basic approval items, the result is either 'Approved' or 'Rejected'. For custom approval items, the result could either be a single response or multiple responses separated by a semi-colon. Read-only.
func (m *ApprovalItem) SetResult(value *string)() {
    err := m.GetBackingStore().Set("result", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. The approval item state. The possible values are: canceled, created, pending, completed. Read-only.
func (m *ApprovalItem) SetState(value *ApprovalItemState)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// SetViewPoint sets the viewPoint property value. Represents user viewpoints data on the ApprovalItem. The data includes the users roles regarding the approval item. Read-only.
func (m *ApprovalItem) SetViewPoint(value ApprovalItemViewPointable)() {
    err := m.GetBackingStore().Set("viewPoint", value)
    if err != nil {
        panic(err)
    }
}
type ApprovalItemable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowCancel()(*bool)
    GetAllowEmailNotification()(*bool)
    GetApprovalType()(*ApprovalItemType)
    GetApprovers()([]ApprovalIdentitySetable)
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetOwner()(ApprovalIdentitySetable)
    GetRequests()([]ApprovalItemRequestable)
    GetResponsePrompts()([]string)
    GetResponses()([]ApprovalItemResponseable)
    GetResult()(*string)
    GetState()(*ApprovalItemState)
    GetViewPoint()(ApprovalItemViewPointable)
    SetAllowCancel(value *bool)()
    SetAllowEmailNotification(value *bool)()
    SetApprovalType(value *ApprovalItemType)()
    SetApprovers(value []ApprovalIdentitySetable)()
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetOwner(value ApprovalIdentitySetable)()
    SetRequests(value []ApprovalItemRequestable)()
    SetResponsePrompts(value []string)()
    SetResponses(value []ApprovalItemResponseable)()
    SetResult(value *string)()
    SetState(value *ApprovalItemState)()
    SetViewPoint(value ApprovalItemViewPointable)()
}
