package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessPackageResourceRequest struct {
    Entity
    // Nullable.
    accessPackageResource *AccessPackageResource;
    // The unique ID of the access package catalog.
    catalogId *string;
    // 
    executeImmediately *bool;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // If set, does not add the resource.
    isValidationOnly *bool;
    // The requestor's justification for adding or removing the resource.
    justification *string;
    // Read-only. Nullable.
    requestor *AccessPackageSubject;
    // The outcome of whether the service was able to add the resource to the catalog.  The value is Delivered if the resource was added or removed. Read-Only.
    requestState *string;
    // Read-only.
    requestStatus *string;
    // Use AdminAdd to add a resource, if the caller is an administrator or resource owner, or AdminRemove to remove a resource.
    requestType *string;
}
// Instantiates a new accessPackageResourceRequest and sets the default values.
func NewAccessPackageResourceRequest()(*AccessPackageResourceRequest) {
    m := &AccessPackageResourceRequest{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accessPackageResource property value. Nullable.
func (m *AccessPackageResourceRequest) GetAccessPackageResource()(*AccessPackageResource) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResource
    }
}
// Gets the catalogId property value. The unique ID of the access package catalog.
func (m *AccessPackageResourceRequest) GetCatalogId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.catalogId
    }
}
// Gets the executeImmediately property value. 
func (m *AccessPackageResourceRequest) GetExecuteImmediately()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.executeImmediately
    }
}
// Gets the expirationDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageResourceRequest) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the isValidationOnly property value. If set, does not add the resource.
func (m *AccessPackageResourceRequest) GetIsValidationOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isValidationOnly
    }
}
// Gets the justification property value. The requestor's justification for adding or removing the resource.
func (m *AccessPackageResourceRequest) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
// Gets the requestor property value. Read-only. Nullable.
func (m *AccessPackageResourceRequest) GetRequestor()(*AccessPackageSubject) {
    if m == nil {
        return nil
    } else {
        return m.requestor
    }
}
// Gets the requestState property value. The outcome of whether the service was able to add the resource to the catalog.  The value is Delivered if the resource was added or removed. Read-Only.
func (m *AccessPackageResourceRequest) GetRequestState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestState
    }
}
// Gets the requestStatus property value. Read-only.
func (m *AccessPackageResourceRequest) GetRequestStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestStatus
    }
}
// Gets the requestType property value. Use AdminAdd to add a resource, if the caller is an administrator or resource owner, or AdminRemove to remove a resource.
func (m *AccessPackageResourceRequest) GetRequestType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestType
    }
}
// The deserialization information for the current model
func (m *AccessPackageResourceRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageResource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResource() })
        if err != nil {
            return err
        }
        m.SetAccessPackageResource(val.(*AccessPackageResource))
        return nil
    }
    res["catalogId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCatalogId(val)
        return nil
    }
    res["executeImmediately"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetExecuteImmediately(val)
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpirationDateTime(val)
        return nil
    }
    res["isValidationOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsValidationOnly(val)
        return nil
    }
    res["justification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJustification(val)
        return nil
    }
    res["requestor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageSubject() })
        if err != nil {
            return err
        }
        m.SetRequestor(val.(*AccessPackageSubject))
        return nil
    }
    res["requestState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRequestState(val)
        return nil
    }
    res["requestStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRequestStatus(val)
        return nil
    }
    res["requestType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRequestType(val)
        return nil
    }
    return res
}
func (m *AccessPackageResourceRequest) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AccessPackageResourceRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the accessPackageResource property value. Nullable.
// Parameters:
//  - value : Value to set for the accessPackageResource property.
func (m *AccessPackageResourceRequest) SetAccessPackageResource(value *AccessPackageResource)() {
    m.accessPackageResource = value
}
// Sets the catalogId property value. The unique ID of the access package catalog.
// Parameters:
//  - value : Value to set for the catalogId property.
func (m *AccessPackageResourceRequest) SetCatalogId(value *string)() {
    m.catalogId = value
}
// Sets the executeImmediately property value. 
// Parameters:
//  - value : Value to set for the executeImmediately property.
func (m *AccessPackageResourceRequest) SetExecuteImmediately(value *bool)() {
    m.executeImmediately = value
}
// Sets the expirationDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *AccessPackageResourceRequest) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the isValidationOnly property value. If set, does not add the resource.
// Parameters:
//  - value : Value to set for the isValidationOnly property.
func (m *AccessPackageResourceRequest) SetIsValidationOnly(value *bool)() {
    m.isValidationOnly = value
}
// Sets the justification property value. The requestor's justification for adding or removing the resource.
// Parameters:
//  - value : Value to set for the justification property.
func (m *AccessPackageResourceRequest) SetJustification(value *string)() {
    m.justification = value
}
// Sets the requestor property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the requestor property.
func (m *AccessPackageResourceRequest) SetRequestor(value *AccessPackageSubject)() {
    m.requestor = value
}
// Sets the requestState property value. The outcome of whether the service was able to add the resource to the catalog.  The value is Delivered if the resource was added or removed. Read-Only.
// Parameters:
//  - value : Value to set for the requestState property.
func (m *AccessPackageResourceRequest) SetRequestState(value *string)() {
    m.requestState = value
}
// Sets the requestStatus property value. Read-only.
// Parameters:
//  - value : Value to set for the requestStatus property.
func (m *AccessPackageResourceRequest) SetRequestStatus(value *string)() {
    m.requestStatus = value
}
// Sets the requestType property value. Use AdminAdd to add a resource, if the caller is an administrator or resource owner, or AdminRemove to remove a resource.
// Parameters:
//  - value : Value to set for the requestType property.
func (m *AccessPackageResourceRequest) SetRequestType(value *string)() {
    m.requestType = value
}
