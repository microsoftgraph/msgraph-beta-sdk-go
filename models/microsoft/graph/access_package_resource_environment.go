package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessPackageResourceEnvironment struct {
    Entity
    // Read-only. Required.
    accessPackageResources []AccessPackageResource;
    // Connection information of an environment used to connect to a resource.
    connectionInfo *ConnectionInfo;
    // The display name of the user that created this object.
    createdBy *string;
    // The date and time that this object was created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The description of this accessPackageResourceEnvironment object.
    description *string;
    // The display name of this object.
    displayName *string;
    // Determines whether this is default environment or not. It is set to true for all static origin systems, such as Azure AD groups and Azure AD Applications.
    isDefaultEnvironment *bool;
    // The display name of the entity that last modified this object.
    modifiedBy *string;
    // The date and time that this object was last modified. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The unique identifier of this environment in the origin system.
    originId *string;
    // The type of the resource in the origin system such as SharePointOnline. Supports $filter.
    originSystem *string;
}
// Instantiates a new accessPackageResourceEnvironment and sets the default values.
func NewAccessPackageResourceEnvironment()(*AccessPackageResourceEnvironment) {
    m := &AccessPackageResourceEnvironment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accessPackageResources property value. Read-only. Required.
func (m *AccessPackageResourceEnvironment) GetAccessPackageResources()([]AccessPackageResource) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResources
    }
}
// Gets the connectionInfo property value. Connection information of an environment used to connect to a resource.
func (m *AccessPackageResourceEnvironment) GetConnectionInfo()(*ConnectionInfo) {
    if m == nil {
        return nil
    } else {
        return m.connectionInfo
    }
}
// Gets the createdBy property value. The display name of the user that created this object.
func (m *AccessPackageResourceEnvironment) GetCreatedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the createdDateTime property value. The date and time that this object was created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AccessPackageResourceEnvironment) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. The description of this accessPackageResourceEnvironment object.
func (m *AccessPackageResourceEnvironment) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name of this object.
func (m *AccessPackageResourceEnvironment) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the isDefaultEnvironment property value. Determines whether this is default environment or not. It is set to true for all static origin systems, such as Azure AD groups and Azure AD Applications.
func (m *AccessPackageResourceEnvironment) GetIsDefaultEnvironment()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefaultEnvironment
    }
}
// Gets the modifiedBy property value. The display name of the entity that last modified this object.
func (m *AccessPackageResourceEnvironment) GetModifiedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.modifiedBy
    }
}
// Gets the modifiedDateTime property value. The date and time that this object was last modified. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AccessPackageResourceEnvironment) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// Gets the originId property value. The unique identifier of this environment in the origin system.
func (m *AccessPackageResourceEnvironment) GetOriginId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originId
    }
}
// Gets the originSystem property value. The type of the resource in the origin system such as SharePointOnline. Supports $filter.
func (m *AccessPackageResourceEnvironment) GetOriginSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originSystem
    }
}
// The deserialization information for the current model
func (m *AccessPackageResourceEnvironment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResource, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageResource))
            }
            m.SetAccessPackageResources(res)
        }
        return nil
    }
    res["connectionInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConnectionInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionInfo(val.(*ConnectionInfo))
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["isDefaultEnvironment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultEnvironment(val)
        }
        return nil
    }
    res["modifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedBy(val)
        }
        return nil
    }
    res["modifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    res["originId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginId(val)
        }
        return nil
    }
    res["originSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginSystem(val)
        }
        return nil
    }
    return res
}
func (m *AccessPackageResourceEnvironment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AccessPackageResourceEnvironment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageResources()))
        for i, v := range m.GetAccessPackageResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("connectionInfo", m.GetConnectionInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefaultEnvironment", m.GetIsDefaultEnvironment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("modifiedBy", m.GetModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originId", m.GetOriginId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originSystem", m.GetOriginSystem())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the accessPackageResources property value. Read-only. Required.
// Parameters:
//  - value : Value to set for the accessPackageResources property.
func (m *AccessPackageResourceEnvironment) SetAccessPackageResources(value []AccessPackageResource)() {
    m.accessPackageResources = value
}
// Sets the connectionInfo property value. Connection information of an environment used to connect to a resource.
// Parameters:
//  - value : Value to set for the connectionInfo property.
func (m *AccessPackageResourceEnvironment) SetConnectionInfo(value *ConnectionInfo)() {
    m.connectionInfo = value
}
// Sets the createdBy property value. The display name of the user that created this object.
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *AccessPackageResourceEnvironment) SetCreatedBy(value *string)() {
    m.createdBy = value
}
// Sets the createdDateTime property value. The date and time that this object was created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *AccessPackageResourceEnvironment) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. The description of this accessPackageResourceEnvironment object.
// Parameters:
//  - value : Value to set for the description property.
func (m *AccessPackageResourceEnvironment) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name of this object.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AccessPackageResourceEnvironment) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the isDefaultEnvironment property value. Determines whether this is default environment or not. It is set to true for all static origin systems, such as Azure AD groups and Azure AD Applications.
// Parameters:
//  - value : Value to set for the isDefaultEnvironment property.
func (m *AccessPackageResourceEnvironment) SetIsDefaultEnvironment(value *bool)() {
    m.isDefaultEnvironment = value
}
// Sets the modifiedBy property value. The display name of the entity that last modified this object.
// Parameters:
//  - value : Value to set for the modifiedBy property.
func (m *AccessPackageResourceEnvironment) SetModifiedBy(value *string)() {
    m.modifiedBy = value
}
// Sets the modifiedDateTime property value. The date and time that this object was last modified. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the modifiedDateTime property.
func (m *AccessPackageResourceEnvironment) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// Sets the originId property value. The unique identifier of this environment in the origin system.
// Parameters:
//  - value : Value to set for the originId property.
func (m *AccessPackageResourceEnvironment) SetOriginId(value *string)() {
    m.originId = value
}
// Sets the originSystem property value. The type of the resource in the origin system such as SharePointOnline. Supports $filter.
// Parameters:
//  - value : Value to set for the originSystem property.
func (m *AccessPackageResourceEnvironment) SetOriginSystem(value *string)() {
    m.originSystem = value
}
