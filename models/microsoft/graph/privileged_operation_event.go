package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrivilegedOperationEvent struct {
    Entity
    // Detailed human readable information for the event.
    additionalInformation *string;
    // Indicates the time when the event is created.
    creationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // This is only used when the requestType is Activate, and it indicates the expiration time for the role activation.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Incident/Request ticket number during role activation. The value is presented only if the ticket number is provided during role activation.
    referenceKey *string;
    // Incident/Request ticketing system provided during tole activation. The value is presented only if the ticket system is provided during role activation.
    referenceSystem *string;
    // The user id of the requestor who initiates the operation.
    requestorId *string;
    // The user name of the requestor who initiates the operation.
    requestorName *string;
    // The request operation type. The requestType value can be: Assign (role assignment), Activate (role activation), Unassign (remove role assignment), Deactivate (role deactivation), ScanAlertsNow (scan security alerts), DismissAlert (dismiss security alert), FixAlertItem (fix a security alert issue),  AccessReview_Review (review an Access Review), AccessReview_Create (create an Access Review) , AccessReview_Update (update an Access Review), AccessReview_Delete (delete an Access Review).
    requestType *string;
    // The id of the role that is associated with the operation.
    roleId *string;
    // The name of the role.
    roleName *string;
    // The tenant (organization) id.
    tenantId *string;
    // The id of the user that is associated with the operation.
    userId *string;
    // The user's email.
    userMail *string;
    // The user's display name.
    userName *string;
}
// Instantiates a new privilegedOperationEvent and sets the default values.
func NewPrivilegedOperationEvent()(*PrivilegedOperationEvent) {
    m := &PrivilegedOperationEvent{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the additionalInformation property value. Detailed human readable information for the event.
func (m *PrivilegedOperationEvent) GetAdditionalInformation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalInformation
    }
}
// Gets the creationDateTime property value. Indicates the time when the event is created.
func (m *PrivilegedOperationEvent) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.creationDateTime
    }
}
// Gets the expirationDateTime property value. This is only used when the requestType is Activate, and it indicates the expiration time for the role activation.
func (m *PrivilegedOperationEvent) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the referenceKey property value. Incident/Request ticket number during role activation. The value is presented only if the ticket number is provided during role activation.
func (m *PrivilegedOperationEvent) GetReferenceKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.referenceKey
    }
}
// Gets the referenceSystem property value. Incident/Request ticketing system provided during tole activation. The value is presented only if the ticket system is provided during role activation.
func (m *PrivilegedOperationEvent) GetReferenceSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.referenceSystem
    }
}
// Gets the requestorId property value. The user id of the requestor who initiates the operation.
func (m *PrivilegedOperationEvent) GetRequestorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestorId
    }
}
// Gets the requestorName property value. The user name of the requestor who initiates the operation.
func (m *PrivilegedOperationEvent) GetRequestorName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestorName
    }
}
// Gets the requestType property value. The request operation type. The requestType value can be: Assign (role assignment), Activate (role activation), Unassign (remove role assignment), Deactivate (role deactivation), ScanAlertsNow (scan security alerts), DismissAlert (dismiss security alert), FixAlertItem (fix a security alert issue),  AccessReview_Review (review an Access Review), AccessReview_Create (create an Access Review) , AccessReview_Update (update an Access Review), AccessReview_Delete (delete an Access Review).
func (m *PrivilegedOperationEvent) GetRequestType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestType
    }
}
// Gets the roleId property value. The id of the role that is associated with the operation.
func (m *PrivilegedOperationEvent) GetRoleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleId
    }
}
// Gets the roleName property value. The name of the role.
func (m *PrivilegedOperationEvent) GetRoleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleName
    }
}
// Gets the tenantId property value. The tenant (organization) id.
func (m *PrivilegedOperationEvent) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// Gets the userId property value. The id of the user that is associated with the operation.
func (m *PrivilegedOperationEvent) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Gets the userMail property value. The user's email.
func (m *PrivilegedOperationEvent) GetUserMail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userMail
    }
}
// Gets the userName property value. The user's display name.
func (m *PrivilegedOperationEvent) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// The deserialization information for the current model
func (m *PrivilegedOperationEvent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalInformation(val)
        }
        return nil
    }
    res["creationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationDateTime(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["referenceKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceKey(val)
        }
        return nil
    }
    res["referenceSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceSystem(val)
        }
        return nil
    }
    res["requestorId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestorId(val)
        }
        return nil
    }
    res["requestorName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestorName(val)
        }
        return nil
    }
    res["requestType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestType(val)
        }
        return nil
    }
    res["roleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleId(val)
        }
        return nil
    }
    res["roleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleName(val)
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userMail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserMail(val)
        }
        return nil
    }
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *PrivilegedOperationEvent) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrivilegedOperationEvent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the additionalInformation property value. Detailed human readable information for the event.
// Parameters:
//  - value : Value to set for the additionalInformation property.
func (m *PrivilegedOperationEvent) SetAdditionalInformation(value *string)() {
    m.additionalInformation = value
}
// Sets the creationDateTime property value. Indicates the time when the event is created.
// Parameters:
//  - value : Value to set for the creationDateTime property.
func (m *PrivilegedOperationEvent) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.creationDateTime = value
}
// Sets the expirationDateTime property value. This is only used when the requestType is Activate, and it indicates the expiration time for the role activation.
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *PrivilegedOperationEvent) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the referenceKey property value. Incident/Request ticket number during role activation. The value is presented only if the ticket number is provided during role activation.
// Parameters:
//  - value : Value to set for the referenceKey property.
func (m *PrivilegedOperationEvent) SetReferenceKey(value *string)() {
    m.referenceKey = value
}
// Sets the referenceSystem property value. Incident/Request ticketing system provided during tole activation. The value is presented only if the ticket system is provided during role activation.
// Parameters:
//  - value : Value to set for the referenceSystem property.
func (m *PrivilegedOperationEvent) SetReferenceSystem(value *string)() {
    m.referenceSystem = value
}
// Sets the requestorId property value. The user id of the requestor who initiates the operation.
// Parameters:
//  - value : Value to set for the requestorId property.
func (m *PrivilegedOperationEvent) SetRequestorId(value *string)() {
    m.requestorId = value
}
// Sets the requestorName property value. The user name of the requestor who initiates the operation.
// Parameters:
//  - value : Value to set for the requestorName property.
func (m *PrivilegedOperationEvent) SetRequestorName(value *string)() {
    m.requestorName = value
}
// Sets the requestType property value. The request operation type. The requestType value can be: Assign (role assignment), Activate (role activation), Unassign (remove role assignment), Deactivate (role deactivation), ScanAlertsNow (scan security alerts), DismissAlert (dismiss security alert), FixAlertItem (fix a security alert issue),  AccessReview_Review (review an Access Review), AccessReview_Create (create an Access Review) , AccessReview_Update (update an Access Review), AccessReview_Delete (delete an Access Review).
// Parameters:
//  - value : Value to set for the requestType property.
func (m *PrivilegedOperationEvent) SetRequestType(value *string)() {
    m.requestType = value
}
// Sets the roleId property value. The id of the role that is associated with the operation.
// Parameters:
//  - value : Value to set for the roleId property.
func (m *PrivilegedOperationEvent) SetRoleId(value *string)() {
    m.roleId = value
}
// Sets the roleName property value. The name of the role.
// Parameters:
//  - value : Value to set for the roleName property.
func (m *PrivilegedOperationEvent) SetRoleName(value *string)() {
    m.roleName = value
}
// Sets the tenantId property value. The tenant (organization) id.
// Parameters:
//  - value : Value to set for the tenantId property.
func (m *PrivilegedOperationEvent) SetTenantId(value *string)() {
    m.tenantId = value
}
// Sets the userId property value. The id of the user that is associated with the operation.
// Parameters:
//  - value : Value to set for the userId property.
func (m *PrivilegedOperationEvent) SetUserId(value *string)() {
    m.userId = value
}
// Sets the userMail property value. The user's email.
// Parameters:
//  - value : Value to set for the userMail property.
func (m *PrivilegedOperationEvent) SetUserMail(value *string)() {
    m.userMail = value
}
// Sets the userName property value. The user's display name.
// Parameters:
//  - value : Value to set for the userName property.
func (m *PrivilegedOperationEvent) SetUserName(value *string)() {
    m.userName = value
}
