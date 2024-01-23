package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// OperationApprovalRequestEntityStatus the operationApprovalRequestEntityStatus complex type is used to provide the basic infortmation relating to the status of a request without revealing too much information to the calling user as it may be an object out of scope.
type OperationApprovalRequestEntityStatus struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOperationApprovalRequestEntityStatus instantiates a new operationApprovalRequestEntityStatus and sets the default values.
func NewOperationApprovalRequestEntityStatus()(*OperationApprovalRequestEntityStatus) {
    m := &OperationApprovalRequestEntityStatus{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOperationApprovalRequestEntityStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOperationApprovalRequestEntityStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOperationApprovalRequestEntityStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OperationApprovalRequestEntityStatus) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *OperationApprovalRequestEntityStatus) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEntityLocked gets the entityLocked property value. The status of the Entity connected to the OperationApprovalRequest in regard to changes, whether further requests are allowed or if the Entity is locked. When true, a lock is present on the Entity and no approval requests can be currently made for it. When false, the Entity is not locked and approval requests are allowed. Default value is false. Read-only. This property is read-only.
func (m *OperationApprovalRequestEntityStatus) GetEntityLocked()(*bool) {
    val, err := m.GetBackingStore().Get("entityLocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OperationApprovalRequestEntityStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["entityLocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntityLocked(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["requestExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestExpirationDateTime(val)
        }
        return nil
    }
    res["requestId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestId(val)
        }
        return nil
    }
    res["requestStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperationApprovalRequestStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestStatus(val.(*OperationApprovalRequestStatus))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OperationApprovalRequestEntityStatus) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestExpirationDateTime gets the requestExpirationDateTime property value. Indicates the DateTime when any action on the OperationApprovalRequest is no longer permitted. The value cannot be modified and is automatically populated when the request is created using expiration offset values defined in the service controllers. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only. This property is read-only.
func (m *OperationApprovalRequestEntityStatus) GetRequestExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("requestExpirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRequestId gets the requestId property value. The unique identifier of the OperationApprovalRequest. This property cannot be modified and is required when the entity status is created. Read-only. This property is read-only.
func (m *OperationApprovalRequestEntityStatus) GetRequestId()(*string) {
    val, err := m.GetBackingStore().Get("requestId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestStatus gets the requestStatus property value. Indicates the status of the Approval Request. The status of a request will change when an action is successfully performed on it, such as when it is `approved` or `rejected`, or when the request's expiration DateTime passes and the result is `expired`.
func (m *OperationApprovalRequestEntityStatus) GetRequestStatus()(*OperationApprovalRequestStatus) {
    val, err := m.GetBackingStore().Get("requestStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OperationApprovalRequestStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OperationApprovalRequestEntityStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetRequestStatus() != nil {
        cast := (*m.GetRequestStatus()).String()
        err := writer.WriteStringValue("requestStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OperationApprovalRequestEntityStatus) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *OperationApprovalRequestEntityStatus) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEntityLocked sets the entityLocked property value. The status of the Entity connected to the OperationApprovalRequest in regard to changes, whether further requests are allowed or if the Entity is locked. When true, a lock is present on the Entity and no approval requests can be currently made for it. When false, the Entity is not locked and approval requests are allowed. Default value is false. Read-only. This property is read-only.
func (m *OperationApprovalRequestEntityStatus) SetEntityLocked(value *bool)() {
    err := m.GetBackingStore().Set("entityLocked", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OperationApprovalRequestEntityStatus) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestExpirationDateTime sets the requestExpirationDateTime property value. Indicates the DateTime when any action on the OperationApprovalRequest is no longer permitted. The value cannot be modified and is automatically populated when the request is created using expiration offset values defined in the service controllers. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only. This property is read-only.
func (m *OperationApprovalRequestEntityStatus) SetRequestExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("requestExpirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestId sets the requestId property value. The unique identifier of the OperationApprovalRequest. This property cannot be modified and is required when the entity status is created. Read-only. This property is read-only.
func (m *OperationApprovalRequestEntityStatus) SetRequestId(value *string)() {
    err := m.GetBackingStore().Set("requestId", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestStatus sets the requestStatus property value. Indicates the status of the Approval Request. The status of a request will change when an action is successfully performed on it, such as when it is `approved` or `rejected`, or when the request's expiration DateTime passes and the result is `expired`.
func (m *OperationApprovalRequestEntityStatus) SetRequestStatus(value *OperationApprovalRequestStatus)() {
    err := m.GetBackingStore().Set("requestStatus", value)
    if err != nil {
        panic(err)
    }
}
// OperationApprovalRequestEntityStatusable 
type OperationApprovalRequestEntityStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEntityLocked()(*bool)
    GetOdataType()(*string)
    GetRequestExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRequestId()(*string)
    GetRequestStatus()(*OperationApprovalRequestStatus)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEntityLocked(value *bool)()
    SetOdataType(value *string)()
    SetRequestExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRequestId(value *string)()
    SetRequestStatus(value *OperationApprovalRequestStatus)()
}
