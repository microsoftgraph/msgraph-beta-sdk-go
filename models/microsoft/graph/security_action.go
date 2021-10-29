package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SecurityAction struct {
    Entity
    // Reason for invoking this action.
    actionReason *string;
    // The Application ID of the calling application that submitted (POST) the action. The appId should be extracted from the auth token and not entered manually by the calling application.
    appId *string;
    // Azure tenant ID of the entity to determine which tenant the entity belongs to (multi-tenancy support). The azureTenantId should be extracted from the auth token and not entered manually by the calling application.
    azureTenantId *string;
    // 
    clientContext *string;
    // Timestamp when the action was completed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Timestamp when the action is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Error info when the action fails.
    errorInfo *ResultInfo;
    // Timestamp when this action was last updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Action name.
    name *string;
    // Collection of parameters (key-value pairs) necessary to invoke the action, e.g. URL or fileHash to block, etc.). Required
    parameters []KeyValuePair;
    // Collection of securityActionState to keep the history of an action.
    states []SecurityActionState;
    // Status of the action. Possible values are: NotStarted, Running, Completed, Failed.
    status *OperationStatus;
    // The user principal name of the signed-in user that submitted  (POST) the action. The user should be extracted from the auth token and not entered manually by the calling application.
    user *string;
    // Complex Type containing details about the Security product/service vendor, provider, and sub-provider (e.g. vendor=Microsoft; provider=Windows Defender ATP; sub-provider=AppLocker).
    vendorInformation *SecurityVendorInformation;
}
// Instantiates a new securityAction and sets the default values.
func NewSecurityAction()(*SecurityAction) {
    m := &SecurityAction{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the actionReason property value. Reason for invoking this action.
func (m *SecurityAction) GetActionReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionReason
    }
}
// Gets the appId property value. The Application ID of the calling application that submitted (POST) the action. The appId should be extracted from the auth token and not entered manually by the calling application.
func (m *SecurityAction) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// Gets the azureTenantId property value. Azure tenant ID of the entity to determine which tenant the entity belongs to (multi-tenancy support). The azureTenantId should be extracted from the auth token and not entered manually by the calling application.
func (m *SecurityAction) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
// Gets the clientContext property value. 
func (m *SecurityAction) GetClientContext()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientContext
    }
}
// Gets the completedDateTime property value. Timestamp when the action was completed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *SecurityAction) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.completedDateTime
    }
}
// Gets the createdDateTime property value. Timestamp when the action is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *SecurityAction) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the errorInfo property value. Error info when the action fails.
func (m *SecurityAction) GetErrorInfo()(*ResultInfo) {
    if m == nil {
        return nil
    } else {
        return m.errorInfo
    }
}
// Gets the lastActionDateTime property value. Timestamp when this action was last updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *SecurityAction) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastActionDateTime
    }
}
// Gets the name property value. Action name.
func (m *SecurityAction) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the parameters property value. Collection of parameters (key-value pairs) necessary to invoke the action, e.g. URL or fileHash to block, etc.). Required
func (m *SecurityAction) GetParameters()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.parameters
    }
}
// Gets the states property value. Collection of securityActionState to keep the history of an action.
func (m *SecurityAction) GetStates()([]SecurityActionState) {
    if m == nil {
        return nil
    } else {
        return m.states
    }
}
// Gets the status property value. Status of the action. Possible values are: NotStarted, Running, Completed, Failed.
func (m *SecurityAction) GetStatus()(*OperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the user property value. The user principal name of the signed-in user that submitted  (POST) the action. The user should be extracted from the auth token and not entered manually by the calling application.
func (m *SecurityAction) GetUser()(*string) {
    if m == nil {
        return nil
    } else {
        return m.user
    }
}
// Gets the vendorInformation property value. Complex Type containing details about the Security product/service vendor, provider, and sub-provider (e.g. vendor=Microsoft; provider=Windows Defender ATP; sub-provider=AppLocker).
func (m *SecurityAction) GetVendorInformation()(*SecurityVendorInformation) {
    if m == nil {
        return nil
    } else {
        return m.vendorInformation
    }
}
// The deserialization information for the current model
func (m *SecurityAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActionReason(val)
        return nil
    }
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppId(val)
        return nil
    }
    res["azureTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureTenantId(val)
        return nil
    }
    res["clientContext"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClientContext(val)
        return nil
    }
    res["completedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCompletedDateTime(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["errorInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResultInfo() })
        if err != nil {
            return err
        }
        m.SetErrorInfo(val.(*ResultInfo))
        return nil
    }
    res["lastActionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastActionDateTime(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["parameters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        res := make([]KeyValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValuePair))
        }
        m.SetParameters(res)
        return nil
    }
    res["states"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityActionState() })
        if err != nil {
            return err
        }
        res := make([]SecurityActionState, len(val))
        for i, v := range val {
            res[i] = *(v.(*SecurityActionState))
        }
        m.SetStates(res)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperationStatus)
        if err != nil {
            return err
        }
        cast := val.(OperationStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["user"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUser(val)
        return nil
    }
    res["vendorInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityVendorInformation() })
        if err != nil {
            return err
        }
        m.SetVendorInformation(val.(*SecurityVendorInformation))
        return nil
    }
    return res
}
func (m *SecurityAction) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SecurityAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("actionReason", m.GetActionReason())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureTenantId", m.GetAzureTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientContext", m.GetClientContext())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("errorInfo", m.GetErrorInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastActionDateTime", m.GetLastActionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetParameters()))
        for i, v := range m.GetParameters() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("parameters", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetStates()))
        for i, v := range m.GetStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("states", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("user", m.GetUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("vendorInformation", m.GetVendorInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the actionReason property value. Reason for invoking this action.
// Parameters:
//  - value : Value to set for the actionReason property.
func (m *SecurityAction) SetActionReason(value *string)() {
    m.actionReason = value
}
// Sets the appId property value. The Application ID of the calling application that submitted (POST) the action. The appId should be extracted from the auth token and not entered manually by the calling application.
// Parameters:
//  - value : Value to set for the appId property.
func (m *SecurityAction) SetAppId(value *string)() {
    m.appId = value
}
// Sets the azureTenantId property value. Azure tenant ID of the entity to determine which tenant the entity belongs to (multi-tenancy support). The azureTenantId should be extracted from the auth token and not entered manually by the calling application.
// Parameters:
//  - value : Value to set for the azureTenantId property.
func (m *SecurityAction) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// Sets the clientContext property value. 
// Parameters:
//  - value : Value to set for the clientContext property.
func (m *SecurityAction) SetClientContext(value *string)() {
    m.clientContext = value
}
// Sets the completedDateTime property value. Timestamp when the action was completed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the completedDateTime property.
func (m *SecurityAction) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// Sets the createdDateTime property value. Timestamp when the action is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *SecurityAction) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the errorInfo property value. Error info when the action fails.
// Parameters:
//  - value : Value to set for the errorInfo property.
func (m *SecurityAction) SetErrorInfo(value *ResultInfo)() {
    m.errorInfo = value
}
// Sets the lastActionDateTime property value. Timestamp when this action was last updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the lastActionDateTime property.
func (m *SecurityAction) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// Sets the name property value. Action name.
// Parameters:
//  - value : Value to set for the name property.
func (m *SecurityAction) SetName(value *string)() {
    m.name = value
}
// Sets the parameters property value. Collection of parameters (key-value pairs) necessary to invoke the action, e.g. URL or fileHash to block, etc.). Required
// Parameters:
//  - value : Value to set for the parameters property.
func (m *SecurityAction) SetParameters(value []KeyValuePair)() {
    m.parameters = value
}
// Sets the states property value. Collection of securityActionState to keep the history of an action.
// Parameters:
//  - value : Value to set for the states property.
func (m *SecurityAction) SetStates(value []SecurityActionState)() {
    m.states = value
}
// Sets the status property value. Status of the action. Possible values are: NotStarted, Running, Completed, Failed.
// Parameters:
//  - value : Value to set for the status property.
func (m *SecurityAction) SetStatus(value *OperationStatus)() {
    m.status = value
}
// Sets the user property value. The user principal name of the signed-in user that submitted  (POST) the action. The user should be extracted from the auth token and not entered manually by the calling application.
// Parameters:
//  - value : Value to set for the user property.
func (m *SecurityAction) SetUser(value *string)() {
    m.user = value
}
// Sets the vendorInformation property value. Complex Type containing details about the Security product/service vendor, provider, and sub-provider (e.g. vendor=Microsoft; provider=Windows Defender ATP; sub-provider=AppLocker).
// Parameters:
//  - value : Value to set for the vendorInformation property.
func (m *SecurityAction) SetVendorInformation(value *SecurityVendorInformation)() {
    m.vendorInformation = value
}
