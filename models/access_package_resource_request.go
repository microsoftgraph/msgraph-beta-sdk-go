package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageResourceRequest 
type AccessPackageResourceRequest struct {
    Entity
}
// NewAccessPackageResourceRequest instantiates a new accessPackageResourceRequest and sets the default values.
func NewAccessPackageResourceRequest()(*AccessPackageResourceRequest) {
    m := &AccessPackageResourceRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageResourceRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageResourceRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageResourceRequest(), nil
}
// GetAccessPackageResource gets the accessPackageResource property value. The accessPackageResource property
func (m *AccessPackageResourceRequest) GetAccessPackageResource()(AccessPackageResourceable) {
    val, err := m.GetBackingStore().Get("accessPackageResource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageResourceable)
    }
    return nil
}
// GetCatalogId gets the catalogId property value. The unique ID of the access package catalog.
func (m *AccessPackageResourceRequest) GetCatalogId()(*string) {
    val, err := m.GetBackingStore().Get("catalogId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExecuteImmediately gets the executeImmediately property value. The executeImmediately property
func (m *AccessPackageResourceRequest) GetExecuteImmediately()(*bool) {
    val, err := m.GetBackingStore().Get("executeImmediately")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetExpirationDateTime gets the expirationDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageResourceRequest) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
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
func (m *AccessPackageResourceRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageResource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageResource(val.(AccessPackageResourceable))
        }
        return nil
    }
    res["catalogId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogId(val)
        }
        return nil
    }
    res["executeImmediately"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExecuteImmediately(val)
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
    res["isValidationOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValidationOnly(val)
        }
        return nil
    }
    res["justification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
        }
        return nil
    }
    res["requestor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestor(val.(AccessPackageSubjectable))
        }
        return nil
    }
    res["requestState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestState(val)
        }
        return nil
    }
    res["requestStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestStatus(val)
        }
        return nil
    }
    res["requestType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestType(val)
        }
        return nil
    }
    return res
}
// GetIsValidationOnly gets the isValidationOnly property value. If set, does not add the resource.
func (m *AccessPackageResourceRequest) GetIsValidationOnly()(*bool) {
    val, err := m.GetBackingStore().Get("isValidationOnly")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetJustification gets the justification property value. The requestor's justification for adding or removing the resource.
func (m *AccessPackageResourceRequest) GetJustification()(*string) {
    val, err := m.GetBackingStore().Get("justification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestor gets the requestor property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackageResourceRequest) GetRequestor()(AccessPackageSubjectable) {
    val, err := m.GetBackingStore().Get("requestor")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageSubjectable)
    }
    return nil
}
// GetRequestState gets the requestState property value. The outcome of whether the service was able to add the resource to the catalog.  The value is Delivered if the resource was added or removed. Read-Only.
func (m *AccessPackageResourceRequest) GetRequestState()(*string) {
    val, err := m.GetBackingStore().Get("requestState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestStatus gets the requestStatus property value. The requestStatus property
func (m *AccessPackageResourceRequest) GetRequestStatus()(*string) {
    val, err := m.GetBackingStore().Get("requestStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestType gets the requestType property value. Use AdminAdd to add a resource, if the caller is an administrator or resource owner, AdminUpdate to update a resource, or AdminRemove to remove a resource.
func (m *AccessPackageResourceRequest) GetRequestType()(*string) {
    val, err := m.GetBackingStore().Get("requestType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessPackageResourceRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessPackageResource", m.GetAccessPackageResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("catalogId", m.GetCatalogId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("executeImmediately", m.GetExecuteImmediately())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isValidationOnly", m.GetIsValidationOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("requestor", m.GetRequestor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestState", m.GetRequestState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestStatus", m.GetRequestStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestType", m.GetRequestType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackageResource sets the accessPackageResource property value. The accessPackageResource property
func (m *AccessPackageResourceRequest) SetAccessPackageResource(value AccessPackageResourceable)() {
    err := m.GetBackingStore().Set("accessPackageResource", value)
    if err != nil {
        panic(err)
    }
}
// SetCatalogId sets the catalogId property value. The unique ID of the access package catalog.
func (m *AccessPackageResourceRequest) SetCatalogId(value *string)() {
    err := m.GetBackingStore().Set("catalogId", value)
    if err != nil {
        panic(err)
    }
}
// SetExecuteImmediately sets the executeImmediately property value. The executeImmediately property
func (m *AccessPackageResourceRequest) SetExecuteImmediately(value *bool)() {
    err := m.GetBackingStore().Set("executeImmediately", value)
    if err != nil {
        panic(err)
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageResourceRequest) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIsValidationOnly sets the isValidationOnly property value. If set, does not add the resource.
func (m *AccessPackageResourceRequest) SetIsValidationOnly(value *bool)() {
    err := m.GetBackingStore().Set("isValidationOnly", value)
    if err != nil {
        panic(err)
    }
}
// SetJustification sets the justification property value. The requestor's justification for adding or removing the resource.
func (m *AccessPackageResourceRequest) SetJustification(value *string)() {
    err := m.GetBackingStore().Set("justification", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestor sets the requestor property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackageResourceRequest) SetRequestor(value AccessPackageSubjectable)() {
    err := m.GetBackingStore().Set("requestor", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestState sets the requestState property value. The outcome of whether the service was able to add the resource to the catalog.  The value is Delivered if the resource was added or removed. Read-Only.
func (m *AccessPackageResourceRequest) SetRequestState(value *string)() {
    err := m.GetBackingStore().Set("requestState", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestStatus sets the requestStatus property value. The requestStatus property
func (m *AccessPackageResourceRequest) SetRequestStatus(value *string)() {
    err := m.GetBackingStore().Set("requestStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestType sets the requestType property value. Use AdminAdd to add a resource, if the caller is an administrator or resource owner, AdminUpdate to update a resource, or AdminRemove to remove a resource.
func (m *AccessPackageResourceRequest) SetRequestType(value *string)() {
    err := m.GetBackingStore().Set("requestType", value)
    if err != nil {
        panic(err)
    }
}
// AccessPackageResourceRequestable 
type AccessPackageResourceRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackageResource()(AccessPackageResourceable)
    GetCatalogId()(*string)
    GetExecuteImmediately()(*bool)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsValidationOnly()(*bool)
    GetJustification()(*string)
    GetRequestor()(AccessPackageSubjectable)
    GetRequestState()(*string)
    GetRequestStatus()(*string)
    GetRequestType()(*string)
    SetAccessPackageResource(value AccessPackageResourceable)()
    SetCatalogId(value *string)()
    SetExecuteImmediately(value *bool)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsValidationOnly(value *bool)()
    SetJustification(value *string)()
    SetRequestor(value AccessPackageSubjectable)()
    SetRequestState(value *string)()
    SetRequestStatus(value *string)()
    SetRequestType(value *string)()
}
