package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RemoteActionAudit struct {
    Entity
    // The action name. Possible values are: unknown, factoryReset, removeCompanyData, resetPasscode, remoteLock, enableLostMode, disableLostMode, locateDevice, rebootNow, recoverPasscode, cleanWindowsDevice, logoutSharedAppleDeviceActiveUser, quickScan, fullScan, windowsDefenderUpdateSignatures, factoryResetKeepEnrollmentData, updateDeviceAccount, automaticRedeployment, shutDown, rotateBitLockerKeys, rotateFileVaultKey, getFileVaultKey, setDeviceName, activateDeviceEsim.
    action *RemoteAction;
    // Action state. Possible values are: none, pending, canceled, active, done, failed, notSupported.
    actionState *ActionState;
    // Intune device name.
    deviceDisplayName *string;
    // IMEI of the device.
    deviceIMEI *string;
    // Upn of the device owner.
    deviceOwnerUserPrincipalName *string;
    // User who initiated the device action, format is UPN.
    initiatedByUserPrincipalName *string;
    // Action target.
    managedDeviceId *string;
    // Time when the action was issued, given in UTC.
    requestDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // [deprecated] Please use InitiatedByUserPrincipalName instead.
    userName *string;
}
// Instantiates a new remoteActionAudit and sets the default values.
func NewRemoteActionAudit()(*RemoteActionAudit) {
    m := &RemoteActionAudit{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the action property value. The action name. Possible values are: unknown, factoryReset, removeCompanyData, resetPasscode, remoteLock, enableLostMode, disableLostMode, locateDevice, rebootNow, recoverPasscode, cleanWindowsDevice, logoutSharedAppleDeviceActiveUser, quickScan, fullScan, windowsDefenderUpdateSignatures, factoryResetKeepEnrollmentData, updateDeviceAccount, automaticRedeployment, shutDown, rotateBitLockerKeys, rotateFileVaultKey, getFileVaultKey, setDeviceName, activateDeviceEsim.
func (m *RemoteActionAudit) GetAction()(*RemoteAction) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// Gets the actionState property value. Action state. Possible values are: none, pending, canceled, active, done, failed, notSupported.
func (m *RemoteActionAudit) GetActionState()(*ActionState) {
    if m == nil {
        return nil
    } else {
        return m.actionState
    }
}
// Gets the deviceDisplayName property value. Intune device name.
func (m *RemoteActionAudit) GetDeviceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceDisplayName
    }
}
// Gets the deviceIMEI property value. IMEI of the device.
func (m *RemoteActionAudit) GetDeviceIMEI()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceIMEI
    }
}
// Gets the deviceOwnerUserPrincipalName property value. Upn of the device owner.
func (m *RemoteActionAudit) GetDeviceOwnerUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceOwnerUserPrincipalName
    }
}
// Gets the initiatedByUserPrincipalName property value. User who initiated the device action, format is UPN.
func (m *RemoteActionAudit) GetInitiatedByUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.initiatedByUserPrincipalName
    }
}
// Gets the managedDeviceId property value. Action target.
func (m *RemoteActionAudit) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
// Gets the requestDateTime property value. Time when the action was issued, given in UTC.
func (m *RemoteActionAudit) GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.requestDateTime
    }
}
// Gets the userName property value. [deprecated] Please use InitiatedByUserPrincipalName instead.
func (m *RemoteActionAudit) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// The deserialization information for the current model
func (m *RemoteActionAudit) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRemoteAction)
        if err != nil {
            return err
        }
        cast := val.(RemoteAction)
        m.SetAction(&cast)
        return nil
    }
    res["actionState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseActionState)
        if err != nil {
            return err
        }
        cast := val.(ActionState)
        m.SetActionState(&cast)
        return nil
    }
    res["deviceDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceDisplayName(val)
        return nil
    }
    res["deviceIMEI"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceIMEI(val)
        return nil
    }
    res["deviceOwnerUserPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceOwnerUserPrincipalName(val)
        return nil
    }
    res["initiatedByUserPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInitiatedByUserPrincipalName(val)
        return nil
    }
    res["managedDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedDeviceId(val)
        return nil
    }
    res["requestDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetRequestDateTime(val)
        return nil
    }
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserName(val)
        return nil
    }
    return res
}
func (m *RemoteActionAudit) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RemoteActionAudit) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAction() != nil {
        cast := m.GetAction().String()
        err = writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetActionState() != nil {
        cast := m.GetActionState().String()
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
// Sets the action property value. The action name. Possible values are: unknown, factoryReset, removeCompanyData, resetPasscode, remoteLock, enableLostMode, disableLostMode, locateDevice, rebootNow, recoverPasscode, cleanWindowsDevice, logoutSharedAppleDeviceActiveUser, quickScan, fullScan, windowsDefenderUpdateSignatures, factoryResetKeepEnrollmentData, updateDeviceAccount, automaticRedeployment, shutDown, rotateBitLockerKeys, rotateFileVaultKey, getFileVaultKey, setDeviceName, activateDeviceEsim.
// Parameters:
//  - value : Value to set for the action property.
func (m *RemoteActionAudit) SetAction(value *RemoteAction)() {
    m.action = value
}
// Sets the actionState property value. Action state. Possible values are: none, pending, canceled, active, done, failed, notSupported.
// Parameters:
//  - value : Value to set for the actionState property.
func (m *RemoteActionAudit) SetActionState(value *ActionState)() {
    m.actionState = value
}
// Sets the deviceDisplayName property value. Intune device name.
// Parameters:
//  - value : Value to set for the deviceDisplayName property.
func (m *RemoteActionAudit) SetDeviceDisplayName(value *string)() {
    m.deviceDisplayName = value
}
// Sets the deviceIMEI property value. IMEI of the device.
// Parameters:
//  - value : Value to set for the deviceIMEI property.
func (m *RemoteActionAudit) SetDeviceIMEI(value *string)() {
    m.deviceIMEI = value
}
// Sets the deviceOwnerUserPrincipalName property value. Upn of the device owner.
// Parameters:
//  - value : Value to set for the deviceOwnerUserPrincipalName property.
func (m *RemoteActionAudit) SetDeviceOwnerUserPrincipalName(value *string)() {
    m.deviceOwnerUserPrincipalName = value
}
// Sets the initiatedByUserPrincipalName property value. User who initiated the device action, format is UPN.
// Parameters:
//  - value : Value to set for the initiatedByUserPrincipalName property.
func (m *RemoteActionAudit) SetInitiatedByUserPrincipalName(value *string)() {
    m.initiatedByUserPrincipalName = value
}
// Sets the managedDeviceId property value. Action target.
// Parameters:
//  - value : Value to set for the managedDeviceId property.
func (m *RemoteActionAudit) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// Sets the requestDateTime property value. Time when the action was issued, given in UTC.
// Parameters:
//  - value : Value to set for the requestDateTime property.
func (m *RemoteActionAudit) SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.requestDateTime = value
}
// Sets the userName property value. [deprecated] Please use InitiatedByUserPrincipalName instead.
// Parameters:
//  - value : Value to set for the userName property.
func (m *RemoteActionAudit) SetUserName(value *string)() {
    m.userName = value
}
