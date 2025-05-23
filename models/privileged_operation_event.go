// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PrivilegedOperationEvent struct {
    Entity
}
// NewPrivilegedOperationEvent instantiates a new PrivilegedOperationEvent and sets the default values.
func NewPrivilegedOperationEvent()(*PrivilegedOperationEvent) {
    m := &PrivilegedOperationEvent{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedOperationEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePrivilegedOperationEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedOperationEvent(), nil
}
// GetAdditionalInformation gets the additionalInformation property value. The additionalInformation property
// returns a *string when successful
func (m *PrivilegedOperationEvent) GetAdditionalInformation()(*string) {
    val, err := m.GetBackingStore().Get("additionalInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreationDateTime gets the creationDateTime property value. The creationDateTime property
// returns a *Time when successful
func (m *PrivilegedOperationEvent) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("creationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetExpirationDateTime gets the expirationDateTime property value. The expirationDateTime property
// returns a *Time when successful
func (m *PrivilegedOperationEvent) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PrivilegedOperationEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalInformation(val)
        }
        return nil
    }
    res["creationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationDateTime(val)
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
    res["referenceKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceKey(val)
        }
        return nil
    }
    res["referenceSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceSystem(val)
        }
        return nil
    }
    res["requestorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestorId(val)
        }
        return nil
    }
    res["requestorName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestorName(val)
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
    res["roleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleId(val)
        }
        return nil
    }
    res["roleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleName(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userMail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserMail(val)
        }
        return nil
    }
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    return res
}
// GetReferenceKey gets the referenceKey property value. The referenceKey property
// returns a *string when successful
func (m *PrivilegedOperationEvent) GetReferenceKey()(*string) {
    val, err := m.GetBackingStore().Get("referenceKey")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReferenceSystem gets the referenceSystem property value. The referenceSystem property
// returns a *string when successful
func (m *PrivilegedOperationEvent) GetReferenceSystem()(*string) {
    val, err := m.GetBackingStore().Get("referenceSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestorId gets the requestorId property value. The requestorId property
// returns a *string when successful
func (m *PrivilegedOperationEvent) GetRequestorId()(*string) {
    val, err := m.GetBackingStore().Get("requestorId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestorName gets the requestorName property value. The requestorName property
// returns a *string when successful
func (m *PrivilegedOperationEvent) GetRequestorName()(*string) {
    val, err := m.GetBackingStore().Get("requestorName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestType gets the requestType property value. The requestType property
// returns a *string when successful
func (m *PrivilegedOperationEvent) GetRequestType()(*string) {
    val, err := m.GetBackingStore().Get("requestType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleId gets the roleId property value. The roleId property
// returns a *string when successful
func (m *PrivilegedOperationEvent) GetRoleId()(*string) {
    val, err := m.GetBackingStore().Get("roleId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleName gets the roleName property value. The roleName property
// returns a *string when successful
func (m *PrivilegedOperationEvent) GetRoleName()(*string) {
    val, err := m.GetBackingStore().Get("roleName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantId gets the tenantId property value. The tenantId property
// returns a *string when successful
func (m *PrivilegedOperationEvent) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserId gets the userId property value. The userId property
// returns a *string when successful
func (m *PrivilegedOperationEvent) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserMail gets the userMail property value. The userMail property
// returns a *string when successful
func (m *PrivilegedOperationEvent) GetUserMail()(*string) {
    val, err := m.GetBackingStore().Get("userMail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserName gets the userName property value. The userName property
// returns a *string when successful
func (m *PrivilegedOperationEvent) GetUserName()(*string) {
    val, err := m.GetBackingStore().Get("userName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PrivilegedOperationEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("additionalInformation", m.GetAdditionalInformation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("creationDateTime", m.GetCreationDateTime())
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
        err = writer.WriteStringValue("referenceKey", m.GetReferenceKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("referenceSystem", m.GetReferenceSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestorId", m.GetRequestorId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestorName", m.GetRequestorName())
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
    {
        err = writer.WriteStringValue("roleId", m.GetRoleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleName", m.GetRoleName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userMail", m.GetUserMail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalInformation sets the additionalInformation property value. The additionalInformation property
func (m *PrivilegedOperationEvent) SetAdditionalInformation(value *string)() {
    err := m.GetBackingStore().Set("additionalInformation", value)
    if err != nil {
        panic(err)
    }
}
// SetCreationDateTime sets the creationDateTime property value. The creationDateTime property
func (m *PrivilegedOperationEvent) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("creationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The expirationDateTime property
func (m *PrivilegedOperationEvent) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceKey sets the referenceKey property value. The referenceKey property
func (m *PrivilegedOperationEvent) SetReferenceKey(value *string)() {
    err := m.GetBackingStore().Set("referenceKey", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceSystem sets the referenceSystem property value. The referenceSystem property
func (m *PrivilegedOperationEvent) SetReferenceSystem(value *string)() {
    err := m.GetBackingStore().Set("referenceSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestorId sets the requestorId property value. The requestorId property
func (m *PrivilegedOperationEvent) SetRequestorId(value *string)() {
    err := m.GetBackingStore().Set("requestorId", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestorName sets the requestorName property value. The requestorName property
func (m *PrivilegedOperationEvent) SetRequestorName(value *string)() {
    err := m.GetBackingStore().Set("requestorName", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestType sets the requestType property value. The requestType property
func (m *PrivilegedOperationEvent) SetRequestType(value *string)() {
    err := m.GetBackingStore().Set("requestType", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleId sets the roleId property value. The roleId property
func (m *PrivilegedOperationEvent) SetRoleId(value *string)() {
    err := m.GetBackingStore().Set("roleId", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleName sets the roleName property value. The roleName property
func (m *PrivilegedOperationEvent) SetRoleName(value *string)() {
    err := m.GetBackingStore().Set("roleName", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *PrivilegedOperationEvent) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. The userId property
func (m *PrivilegedOperationEvent) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserMail sets the userMail property value. The userMail property
func (m *PrivilegedOperationEvent) SetUserMail(value *string)() {
    err := m.GetBackingStore().Set("userMail", value)
    if err != nil {
        panic(err)
    }
}
// SetUserName sets the userName property value. The userName property
func (m *PrivilegedOperationEvent) SetUserName(value *string)() {
    err := m.GetBackingStore().Set("userName", value)
    if err != nil {
        panic(err)
    }
}
type PrivilegedOperationEventable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalInformation()(*string)
    GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReferenceKey()(*string)
    GetReferenceSystem()(*string)
    GetRequestorId()(*string)
    GetRequestorName()(*string)
    GetRequestType()(*string)
    GetRoleId()(*string)
    GetRoleName()(*string)
    GetTenantId()(*string)
    GetUserId()(*string)
    GetUserMail()(*string)
    GetUserName()(*string)
    SetAdditionalInformation(value *string)()
    SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReferenceKey(value *string)()
    SetReferenceSystem(value *string)()
    SetRequestorId(value *string)()
    SetRequestorName(value *string)()
    SetRequestType(value *string)()
    SetRoleId(value *string)()
    SetRoleName(value *string)()
    SetTenantId(value *string)()
    SetUserId(value *string)()
    SetUserMail(value *string)()
    SetUserName(value *string)()
}
