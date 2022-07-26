package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RemoteActionAudit 
type RemoteActionAudit struct {
    Entity
    // Remote actions Intune supports.
    action *RemoteAction
    // The actionState property
    actionState *ActionState
    // Intune device name.
    deviceDisplayName *string
    // IMEI of the device.
    deviceIMEI *string
    // Upn of the device owner.
    deviceOwnerUserPrincipalName *string
    // User who initiated the device action, format is UPN.
    initiatedByUserPrincipalName *string
    // Action target.
    managedDeviceId *string
    // Time when the action was issued, given in UTC.
    requestDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // [deprecated] Please use InitiatedByUserPrincipalName instead.
    userName *string
}
// NewRemoteActionAudit instantiates a new RemoteActionAudit and sets the default values.
func NewRemoteActionAudit()(*RemoteActionAudit) {
    m := &RemoteActionAudit{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.remoteActionAudit";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRemoteActionAuditFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRemoteActionAuditFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoteActionAudit(), nil
}
// GetAction gets the action property value. Remote actions Intune supports.
func (m *RemoteActionAudit) GetAction()(*RemoteAction) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// GetActionState gets the actionState property value. The actionState property
func (m *RemoteActionAudit) GetActionState()(*ActionState) {
    if m == nil {
        return nil
    } else {
        return m.actionState
    }
}
// GetDeviceDisplayName gets the deviceDisplayName property value. Intune device name.
func (m *RemoteActionAudit) GetDeviceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceDisplayName
    }
}
// GetDeviceIMEI gets the deviceIMEI property value. IMEI of the device.
func (m *RemoteActionAudit) GetDeviceIMEI()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceIMEI
    }
}
// GetDeviceOwnerUserPrincipalName gets the deviceOwnerUserPrincipalName property value. Upn of the device owner.
func (m *RemoteActionAudit) GetDeviceOwnerUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceOwnerUserPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
func (m *RemoteActionAudit) GetInitiatedByUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.initiatedByUserPrincipalName
    }
}
// GetManagedDeviceId gets the managedDeviceId property value. Action target.
func (m *RemoteActionAudit) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
// GetRequestDateTime gets the requestDateTime property value. Time when the action was issued, given in UTC.
func (m *RemoteActionAudit) GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.requestDateTime
    }
}
// GetUserName gets the userName property value. [deprecated] Please use InitiatedByUserPrincipalName instead.
func (m *RemoteActionAudit) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
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
    if m != nil {
        m.action = value
    }
}
// SetActionState sets the actionState property value. The actionState property
func (m *RemoteActionAudit) SetActionState(value *ActionState)() {
    if m != nil {
        m.actionState = value
    }
}
// SetDeviceDisplayName sets the deviceDisplayName property value. Intune device name.
func (m *RemoteActionAudit) SetDeviceDisplayName(value *string)() {
    if m != nil {
        m.deviceDisplayName = value
    }
}
// SetDeviceIMEI sets the deviceIMEI property value. IMEI of the device.
func (m *RemoteActionAudit) SetDeviceIMEI(value *string)() {
    if m != nil {
        m.deviceIMEI = value
    }
}
// SetDeviceOwnerUserPrincipalName sets the deviceOwnerUserPrincipalName property value. Upn of the device owner.
func (m *RemoteActionAudit) SetDeviceOwnerUserPrincipalName(value *string)() {
    if m != nil {
        m.deviceOwnerUserPrincipalName = value
    }
}
// SetInitiatedByUserPrincipalName sets the initiatedByUserPrincipalName property value. User who initiated the device action, format is UPN.
func (m *RemoteActionAudit) SetInitiatedByUserPrincipalName(value *string)() {
    if m != nil {
        m.initiatedByUserPrincipalName = value
    }
}
// SetManagedDeviceId sets the managedDeviceId property value. Action target.
func (m *RemoteActionAudit) SetManagedDeviceId(value *string)() {
    if m != nil {
        m.managedDeviceId = value
    }
}
// SetRequestDateTime sets the requestDateTime property value. Time when the action was issued, given in UTC.
func (m *RemoteActionAudit) SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.requestDateTime = value
    }
}
// SetUserName sets the userName property value. [deprecated] Please use InitiatedByUserPrincipalName instead.
func (m *RemoteActionAudit) SetUserName(value *string)() {
    if m != nil {
        m.userName = value
    }
}
