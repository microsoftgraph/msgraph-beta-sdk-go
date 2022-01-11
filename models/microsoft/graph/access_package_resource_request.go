package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageResourceRequest 
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
    // Read-only. Nullable. Supports $expand.
    requestor *AccessPackageSubject;
    // The outcome of whether the service was able to add the resource to the catalog.  The value is Delivered if the resource was added or removed. Read-Only.
    requestState *string;
    // Read-only.
    requestStatus *string;
    // Use AdminAdd to add a resource, if the caller is an administrator or resource owner, or AdminRemove to remove a resource.
    requestType *string;
}
// NewAccessPackageResourceRequest instantiates a new accessPackageResourceRequest and sets the default values.
func NewAccessPackageResourceRequest()(*AccessPackageResourceRequest) {
    m := &AccessPackageResourceRequest{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccessPackageResource gets the accessPackageResource property value. Nullable.
func (m *AccessPackageResourceRequest) GetAccessPackageResource()(*AccessPackageResource) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResource
    }
}
// GetCatalogId gets the catalogId property value. The unique ID of the access package catalog.
func (m *AccessPackageResourceRequest) GetCatalogId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.catalogId
    }
}
// GetExecuteImmediately gets the executeImmediately property value. 
func (m *AccessPackageResourceRequest) GetExecuteImmediately()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.executeImmediately
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageResourceRequest) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetIsValidationOnly gets the isValidationOnly property value. If set, does not add the resource.
func (m *AccessPackageResourceRequest) GetIsValidationOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isValidationOnly
    }
}
// GetJustification gets the justification property value. The requestor's justification for adding or removing the resource.
func (m *AccessPackageResourceRequest) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
// GetRequestor gets the requestor property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackageResourceRequest) GetRequestor()(*AccessPackageSubject) {
    if m == nil {
        return nil
    } else {
        return m.requestor
    }
}
// GetRequestState gets the requestState property value. The outcome of whether the service was able to add the resource to the catalog.  The value is Delivered if the resource was added or removed. Read-Only.
func (m *AccessPackageResourceRequest) GetRequestState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestState
    }
}
// GetRequestStatus gets the requestStatus property value. Read-only.
func (m *AccessPackageResourceRequest) GetRequestStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestStatus
    }
}
// GetRequestType gets the requestType property value. Use AdminAdd to add a resource, if the caller is an administrator or resource owner, or AdminRemove to remove a resource.
func (m *AccessPackageResourceRequest) GetRequestType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageResourceRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageResource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResource() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageResource(val.(*AccessPackageResource))
        }
        return nil
    }
    res["catalogId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogId(val)
        }
        return nil
    }
    res["executeImmediately"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExecuteImmediately(val)
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
    res["isValidationOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValidationOnly(val)
        }
        return nil
    }
    res["justification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
        }
        return nil
    }
    res["requestor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageSubject() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestor(val.(*AccessPackageSubject))
        }
        return nil
    }
    res["requestState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestState(val)
        }
        return nil
    }
    res["requestStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestStatus(val)
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
    return res
}
func (m *AccessPackageResourceRequest) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAccessPackageResource sets the accessPackageResource property value. Nullable.
func (m *AccessPackageResourceRequest) SetAccessPackageResource(value *AccessPackageResource)() {
    if m != nil {
        m.accessPackageResource = value
    }
}
// SetCatalogId sets the catalogId property value. The unique ID of the access package catalog.
func (m *AccessPackageResourceRequest) SetCatalogId(value *string)() {
    if m != nil {
        m.catalogId = value
    }
}
// SetExecuteImmediately sets the executeImmediately property value. 
func (m *AccessPackageResourceRequest) SetExecuteImmediately(value *bool)() {
    if m != nil {
        m.executeImmediately = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageResourceRequest) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetIsValidationOnly sets the isValidationOnly property value. If set, does not add the resource.
func (m *AccessPackageResourceRequest) SetIsValidationOnly(value *bool)() {
    if m != nil {
        m.isValidationOnly = value
    }
}
// SetJustification sets the justification property value. The requestor's justification for adding or removing the resource.
func (m *AccessPackageResourceRequest) SetJustification(value *string)() {
    if m != nil {
        m.justification = value
    }
}
// SetRequestor sets the requestor property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackageResourceRequest) SetRequestor(value *AccessPackageSubject)() {
    if m != nil {
        m.requestor = value
    }
}
// SetRequestState sets the requestState property value. The outcome of whether the service was able to add the resource to the catalog.  The value is Delivered if the resource was added or removed. Read-Only.
func (m *AccessPackageResourceRequest) SetRequestState(value *string)() {
    if m != nil {
        m.requestState = value
    }
}
// SetRequestStatus sets the requestStatus property value. Read-only.
func (m *AccessPackageResourceRequest) SetRequestStatus(value *string)() {
    if m != nil {
        m.requestStatus = value
    }
}
// SetRequestType sets the requestType property value. Use AdminAdd to add a resource, if the caller is an administrator or resource owner, or AdminRemove to remove a resource.
func (m *AccessPackageResourceRequest) SetRequestType(value *string)() {
    if m != nil {
        m.requestType = value
    }
}
