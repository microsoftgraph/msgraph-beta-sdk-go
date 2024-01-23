package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OperationApprovalRequest the OperationApprovalRequest entity encompasses the operation an admin wishes to perform and is requesting approval to complete. It contains the detail of the operation one wishes to perform, user metadata of the requestor, and a justification for the change. It allows for several operations for both the requestor and the potential approver to either approve, deny, or cancel the request and a response justification to provide information for the decision.
type OperationApprovalRequest struct {
    Entity
}
// NewOperationApprovalRequest instantiates a new operationApprovalRequest and sets the default values.
func NewOperationApprovalRequest()(*OperationApprovalRequest) {
    m := &OperationApprovalRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOperationApprovalRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOperationApprovalRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOperationApprovalRequest(), nil
}
// GetApprovalJustification gets the approvalJustification property value. Indicates the justification for approving or rejecting the request. Maximum length of justification is 1024 characters. For example: 'Approved per Change 23423 - needed for Feb 2023 application baseline updates.' Read-only. This property is read-only.
func (m *OperationApprovalRequest) GetApprovalJustification()(*string) {
    val, err := m.GetBackingStore().Get("approvalJustification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetApprover gets the approver property value. The identity of the approver as an Identity Set. Optionally contains the application ID, the device ID and the User ID. See information about this type here: https://learn.microsoft.com/graph/api/resources/identityset?view=graph-rest-1.0. Read-only. This property is read-only.
func (m *OperationApprovalRequest) GetApprover()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("approver")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetExpirationDateTime gets the expirationDateTime property value. Indicates the DateTime when any action on the approval request is no longer permitted. The value cannot be modified and is automatically populated when the request is created using expiration offset values defined in the service controllers. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only. This property is read-only.
func (m *OperationApprovalRequest) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OperationApprovalRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["approvalJustification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalJustification(val)
        }
        return nil
    }
    res["approver"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprover(val.(IdentitySetable))
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["operationApprovalPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationApprovalPolicies(val)
        }
        return nil
    }
    res["requestDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestDateTime(val)
        }
        return nil
    }
    res["requestJustification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestJustification(val)
        }
        return nil
    }
    res["requestor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestor(val.(IdentitySetable))
        }
        return nil
    }
    res["requiredOperationApprovalPolicyTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseOperationApprovalPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OperationApprovalPolicyType, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*OperationApprovalPolicyType))
                }
            }
            m.SetRequiredOperationApprovalPolicyTypes(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperationApprovalRequestStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*OperationApprovalRequestStatus))
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Indicates the last DateTime that the request was modified. The value cannot be modified and is automatically populated whenever values in the request are updated. For example, when the 'status' property changes from needsApproval to approved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only. This property is read-only.
func (m *OperationApprovalRequest) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOperationApprovalPolicies gets the operationApprovalPolicies property value. The operational approval policies used in the request. Indicates the policy and platform combinations that are required for this request to be approved or rejected. Read-only. This property is read-only.
func (m *OperationApprovalRequest) GetOperationApprovalPolicies()(*string) {
    val, err := m.GetBackingStore().Get("operationApprovalPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestDateTime gets the requestDateTime property value. Indicates the DateTime that the request was made. The value cannot be modified and is automatically populated when the request is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only. This property is read-only.
func (m *OperationApprovalRequest) GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("requestDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRequestJustification gets the requestJustification property value. Indicates the justification for creating the request. Maximum length of justification is 1024 characters. For example: 'Needed for Feb 2023 application baseline updates.' Read-only. This property is read-only.
func (m *OperationApprovalRequest) GetRequestJustification()(*string) {
    val, err := m.GetBackingStore().Get("requestJustification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestor gets the requestor property value. The identity of the requestor as an Identity Set. Optionally contains the application ID, the device ID and the User ID. See information about this type here: https://learn.microsoft.com/graph/api/resources/identityset?view=graph-rest-1.0. Read-only. This property is read-only.
func (m *OperationApprovalRequest) GetRequestor()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("requestor")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetRequiredOperationApprovalPolicyTypes gets the requiredOperationApprovalPolicyTypes property value. Indicates the approval policy types required by the request in order for the request to be approved or rejected. Read-only. This property is read-only.
func (m *OperationApprovalRequest) GetRequiredOperationApprovalPolicyTypes()([]OperationApprovalPolicyType) {
    val, err := m.GetBackingStore().Get("requiredOperationApprovalPolicyTypes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OperationApprovalPolicyType)
    }
    return nil
}
// GetStatus gets the status property value. Indicates the status of the Approval Request. The status of a request will change when an action is successfully performed on it, such as when it is `approved` or `rejected`, or when the request's expiration DateTime passes and the result is `expired`.
func (m *OperationApprovalRequest) GetStatus()(*OperationApprovalRequestStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OperationApprovalRequestStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OperationApprovalRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
// SetApprovalJustification sets the approvalJustification property value. Indicates the justification for approving or rejecting the request. Maximum length of justification is 1024 characters. For example: 'Approved per Change 23423 - needed for Feb 2023 application baseline updates.' Read-only. This property is read-only.
func (m *OperationApprovalRequest) SetApprovalJustification(value *string)() {
    err := m.GetBackingStore().Set("approvalJustification", value)
    if err != nil {
        panic(err)
    }
}
// SetApprover sets the approver property value. The identity of the approver as an Identity Set. Optionally contains the application ID, the device ID and the User ID. See information about this type here: https://learn.microsoft.com/graph/api/resources/identityset?view=graph-rest-1.0. Read-only. This property is read-only.
func (m *OperationApprovalRequest) SetApprover(value IdentitySetable)() {
    err := m.GetBackingStore().Set("approver", value)
    if err != nil {
        panic(err)
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. Indicates the DateTime when any action on the approval request is no longer permitted. The value cannot be modified and is automatically populated when the request is created using expiration offset values defined in the service controllers. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only. This property is read-only.
func (m *OperationApprovalRequest) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Indicates the last DateTime that the request was modified. The value cannot be modified and is automatically populated whenever values in the request are updated. For example, when the 'status' property changes from needsApproval to approved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only. This property is read-only.
func (m *OperationApprovalRequest) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOperationApprovalPolicies sets the operationApprovalPolicies property value. The operational approval policies used in the request. Indicates the policy and platform combinations that are required for this request to be approved or rejected. Read-only. This property is read-only.
func (m *OperationApprovalRequest) SetOperationApprovalPolicies(value *string)() {
    err := m.GetBackingStore().Set("operationApprovalPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestDateTime sets the requestDateTime property value. Indicates the DateTime that the request was made. The value cannot be modified and is automatically populated when the request is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only. This property is read-only.
func (m *OperationApprovalRequest) SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("requestDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestJustification sets the requestJustification property value. Indicates the justification for creating the request. Maximum length of justification is 1024 characters. For example: 'Needed for Feb 2023 application baseline updates.' Read-only. This property is read-only.
func (m *OperationApprovalRequest) SetRequestJustification(value *string)() {
    err := m.GetBackingStore().Set("requestJustification", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestor sets the requestor property value. The identity of the requestor as an Identity Set. Optionally contains the application ID, the device ID and the User ID. See information about this type here: https://learn.microsoft.com/graph/api/resources/identityset?view=graph-rest-1.0. Read-only. This property is read-only.
func (m *OperationApprovalRequest) SetRequestor(value IdentitySetable)() {
    err := m.GetBackingStore().Set("requestor", value)
    if err != nil {
        panic(err)
    }
}
// SetRequiredOperationApprovalPolicyTypes sets the requiredOperationApprovalPolicyTypes property value. Indicates the approval policy types required by the request in order for the request to be approved or rejected. Read-only. This property is read-only.
func (m *OperationApprovalRequest) SetRequiredOperationApprovalPolicyTypes(value []OperationApprovalPolicyType)() {
    err := m.GetBackingStore().Set("requiredOperationApprovalPolicyTypes", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Indicates the status of the Approval Request. The status of a request will change when an action is successfully performed on it, such as when it is `approved` or `rejected`, or when the request's expiration DateTime passes and the result is `expired`.
func (m *OperationApprovalRequest) SetStatus(value *OperationApprovalRequestStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// OperationApprovalRequestable 
type OperationApprovalRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApprovalJustification()(*string)
    GetApprover()(IdentitySetable)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOperationApprovalPolicies()(*string)
    GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRequestJustification()(*string)
    GetRequestor()(IdentitySetable)
    GetRequiredOperationApprovalPolicyTypes()([]OperationApprovalPolicyType)
    GetStatus()(*OperationApprovalRequestStatus)
    SetApprovalJustification(value *string)()
    SetApprover(value IdentitySetable)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOperationApprovalPolicies(value *string)()
    SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRequestJustification(value *string)()
    SetRequestor(value IdentitySetable)()
    SetRequiredOperationApprovalPolicyTypes(value []OperationApprovalPolicyType)()
    SetStatus(value *OperationApprovalRequestStatus)()
}
