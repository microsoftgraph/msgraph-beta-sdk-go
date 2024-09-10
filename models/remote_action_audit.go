package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RemoteActionAudit report of remote actions initiated on the devices belonging to a certain tenant.
type RemoteActionAudit struct {
    Entity
}
// NewRemoteActionAudit instantiates a new RemoteActionAudit and sets the default values.
func NewRemoteActionAudit()(*RemoteActionAudit) {
    m := &RemoteActionAudit{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRemoteActionAuditFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRemoteActionAuditFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoteActionAudit(), nil
}
// GetAction gets the action property value. Remote actions Intune supports.
// returns a *RemoteAction when successful
func (m *RemoteActionAudit) GetAction()(*RemoteAction) {
    val, err := m.GetBackingStore().Get("action")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RemoteAction)
    }
    return nil
}
// GetActionState gets the actionState property value. The actionState property
// returns a *ActionState when successful
func (m *RemoteActionAudit) GetActionState()(*ActionState) {
    val, err := m.GetBackingStore().Get("actionState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ActionState)
    }
    return nil
}
// GetBulkDeviceActionId gets the bulkDeviceActionId property value. BulkAction ID
// returns a *string when successful
func (m *RemoteActionAudit) GetBulkDeviceActionId()(*string) {
    val, err := m.GetBackingStore().Get("bulkDeviceActionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceActionCategory gets the deviceActionCategory property value. Enum type used for DeviceActionCategory
// returns a *DeviceActionCategory when successful
func (m *RemoteActionAudit) GetDeviceActionCategory()(*DeviceActionCategory) {
    val, err := m.GetBackingStore().Get("deviceActionCategory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceActionCategory)
    }
    return nil
}
// GetDeviceDisplayName gets the deviceDisplayName property value. Intune device name.
// returns a *string when successful
func (m *RemoteActionAudit) GetDeviceDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("deviceDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceIMEI gets the deviceIMEI property value. IMEI of the device.
// returns a *string when successful
func (m *RemoteActionAudit) GetDeviceIMEI()(*string) {
    val, err := m.GetBackingStore().Get("deviceIMEI")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceOwnerUserPrincipalName gets the deviceOwnerUserPrincipalName property value. Upn of the device owner.
// returns a *string when successful
func (m *RemoteActionAudit) GetDeviceOwnerUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("deviceOwnerUserPrincipalName")
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
func (m *RemoteActionAudit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRemoteAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*RemoteAction))
        }
        return nil
    }
    res["actionState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseActionState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionState(val.(*ActionState))
        }
        return nil
    }
    res["bulkDeviceActionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBulkDeviceActionId(val)
        }
        return nil
    }
    res["deviceActionCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceActionCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceActionCategory(val.(*DeviceActionCategory))
        }
        return nil
    }
    res["deviceDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDisplayName(val)
        }
        return nil
    }
    res["deviceIMEI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceIMEI(val)
        }
        return nil
    }
    res["deviceOwnerUserPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceOwnerUserPrincipalName(val)
        }
        return nil
    }
    res["initiatedByUserPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatedByUserPrincipalName(val)
        }
        return nil
    }
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
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
// GetInitiatedByUserPrincipalName gets the initiatedByUserPrincipalName property value. User who initiated the device action, format is UPN.
// returns a *string when successful
func (m *RemoteActionAudit) GetInitiatedByUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("initiatedByUserPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagedDeviceId gets the managedDeviceId property value. Action target.
// returns a *string when successful
func (m *RemoteActionAudit) GetManagedDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("managedDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestDateTime gets the requestDateTime property value. Time when the action was issued, given in UTC.
// returns a *Time when successful
func (m *RemoteActionAudit) GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("requestDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetUserName gets the userName property value. [deprecated] Please use InitiatedByUserPrincipalName instead.
// returns a *string when successful
func (m *RemoteActionAudit) GetUserName()(*string) {
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
func (m *RemoteActionAudit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err = writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetActionState() != nil {
        cast := (*m.GetActionState()).String()
        err = writer.WriteStringValue("actionState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bulkDeviceActionId", m.GetBulkDeviceActionId())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceActionCategory() != nil {
        cast := (*m.GetDeviceActionCategory()).String()
        err = writer.WriteStringValue("deviceActionCategory", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceDisplayName", m.GetDeviceDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceIMEI", m.GetDeviceIMEI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceOwnerUserPrincipalName", m.GetDeviceOwnerUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initiatedByUserPrincipalName", m.GetInitiatedByUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestDateTime", m.GetRequestDateTime())
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
// SetAction sets the action property value. Remote actions Intune supports.
func (m *RemoteActionAudit) SetAction(value *RemoteAction)() {
    err := m.GetBackingStore().Set("action", value)
    if err != nil {
        panic(err)
    }
}
// SetActionState sets the actionState property value. The actionState property
func (m *RemoteActionAudit) SetActionState(value *ActionState)() {
    err := m.GetBackingStore().Set("actionState", value)
    if err != nil {
        panic(err)
    }
}
// SetBulkDeviceActionId sets the bulkDeviceActionId property value. BulkAction ID
func (m *RemoteActionAudit) SetBulkDeviceActionId(value *string)() {
    err := m.GetBackingStore().Set("bulkDeviceActionId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceActionCategory sets the deviceActionCategory property value. Enum type used for DeviceActionCategory
func (m *RemoteActionAudit) SetDeviceActionCategory(value *DeviceActionCategory)() {
    err := m.GetBackingStore().Set("deviceActionCategory", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceDisplayName sets the deviceDisplayName property value. Intune device name.
func (m *RemoteActionAudit) SetDeviceDisplayName(value *string)() {
    err := m.GetBackingStore().Set("deviceDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceIMEI sets the deviceIMEI property value. IMEI of the device.
func (m *RemoteActionAudit) SetDeviceIMEI(value *string)() {
    err := m.GetBackingStore().Set("deviceIMEI", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceOwnerUserPrincipalName sets the deviceOwnerUserPrincipalName property value. Upn of the device owner.
func (m *RemoteActionAudit) SetDeviceOwnerUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("deviceOwnerUserPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetInitiatedByUserPrincipalName sets the initiatedByUserPrincipalName property value. User who initiated the device action, format is UPN.
func (m *RemoteActionAudit) SetInitiatedByUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("initiatedByUserPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceId sets the managedDeviceId property value. Action target.
func (m *RemoteActionAudit) SetManagedDeviceId(value *string)() {
    err := m.GetBackingStore().Set("managedDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestDateTime sets the requestDateTime property value. Time when the action was issued, given in UTC.
func (m *RemoteActionAudit) SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("requestDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetUserName sets the userName property value. [deprecated] Please use InitiatedByUserPrincipalName instead.
func (m *RemoteActionAudit) SetUserName(value *string)() {
    err := m.GetBackingStore().Set("userName", value)
    if err != nil {
        panic(err)
    }
}
type RemoteActionAuditable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*RemoteAction)
    GetActionState()(*ActionState)
    GetBulkDeviceActionId()(*string)
    GetDeviceActionCategory()(*DeviceActionCategory)
    GetDeviceDisplayName()(*string)
    GetDeviceIMEI()(*string)
    GetDeviceOwnerUserPrincipalName()(*string)
    GetInitiatedByUserPrincipalName()(*string)
    GetManagedDeviceId()(*string)
    GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUserName()(*string)
    SetAction(value *RemoteAction)()
    SetActionState(value *ActionState)()
    SetBulkDeviceActionId(value *string)()
    SetDeviceActionCategory(value *DeviceActionCategory)()
    SetDeviceDisplayName(value *string)()
    SetDeviceIMEI(value *string)()
    SetDeviceOwnerUserPrincipalName(value *string)()
    SetInitiatedByUserPrincipalName(value *string)()
    SetManagedDeviceId(value *string)()
    SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUserName(value *string)()
}
