package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageResourceEnvironment provides operations to manage the collection of accessReview entities.
type AccessPackageResourceEnvironment struct {
    Entity
    // Read-only. Required.
    accessPackageResources []AccessPackageResourceable
    // Connection information of an environment used to connect to a resource.
    connectionInfo ConnectionInfoable
    // The display name of the user that created this object.
    createdBy *string
    // The date and time that this object was created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description of this object.
    description *string
    // The display name of this object.
    displayName *string
    // Determines whether this is default environment or not. It is set to true for all static origin systems, such as Azure AD groups and Azure AD Applications.
    isDefaultEnvironment *bool
    // The display name of the entity that last modified this object.
    modifiedBy *string
    // The date and time that this object was last modified. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique identifier of this environment in the origin system.
    originId *string
    // The type of the resource in the origin system, that is, SharePointOnline. Requires $filter (eq).
    originSystem *string
}
// NewAccessPackageResourceEnvironment instantiates a new accessPackageResourceEnvironment and sets the default values.
func NewAccessPackageResourceEnvironment()(*AccessPackageResourceEnvironment) {
    m := &AccessPackageResourceEnvironment{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.accessPackageResourceEnvironment";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAccessPackageResourceEnvironmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageResourceEnvironmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageResourceEnvironment(), nil
}
// GetAccessPackageResources gets the accessPackageResources property value. Read-only. Required.
func (m *AccessPackageResourceEnvironment) GetAccessPackageResources()([]AccessPackageResourceable) {
    return m.accessPackageResources
}
// GetConnectionInfo gets the connectionInfo property value. Connection information of an environment used to connect to a resource.
func (m *AccessPackageResourceEnvironment) GetConnectionInfo()(ConnectionInfoable) {
    return m.connectionInfo
}
// GetCreatedBy gets the createdBy property value. The display name of the user that created this object.
func (m *AccessPackageResourceEnvironment) GetCreatedBy()(*string) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time that this object was created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AccessPackageResourceEnvironment) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description of this object.
func (m *AccessPackageResourceEnvironment) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of this object.
func (m *AccessPackageResourceEnvironment) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageResourceEnvironment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageResources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessPackageResourceFromDiscriminatorValue , m.SetAccessPackageResources)
    res["connectionInfo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateConnectionInfoFromDiscriminatorValue , m.SetConnectionInfo)
    res["createdBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCreatedBy)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["isDefaultEnvironment"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsDefaultEnvironment)
    res["modifiedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetModifiedBy)
    res["modifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetModifiedDateTime)
    res["originId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOriginId)
    res["originSystem"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOriginSystem)
    return res
}
// GetIsDefaultEnvironment gets the isDefaultEnvironment property value. Determines whether this is default environment or not. It is set to true for all static origin systems, such as Azure AD groups and Azure AD Applications.
func (m *AccessPackageResourceEnvironment) GetIsDefaultEnvironment()(*bool) {
    return m.isDefaultEnvironment
}
// GetModifiedBy gets the modifiedBy property value. The display name of the entity that last modified this object.
func (m *AccessPackageResourceEnvironment) GetModifiedBy()(*string) {
    return m.modifiedBy
}
// GetModifiedDateTime gets the modifiedDateTime property value. The date and time that this object was last modified. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AccessPackageResourceEnvironment) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.modifiedDateTime
}
// GetOriginId gets the originId property value. The unique identifier of this environment in the origin system.
func (m *AccessPackageResourceEnvironment) GetOriginId()(*string) {
    return m.originId
}
// GetOriginSystem gets the originSystem property value. The type of the resource in the origin system, that is, SharePointOnline. Requires $filter (eq).
func (m *AccessPackageResourceEnvironment) GetOriginSystem()(*string) {
    return m.originSystem
}
// Serialize serializes information the current object
func (m *AccessPackageResourceEnvironment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessPackageResources() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAccessPackageResources())
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
// SetAccessPackageResources sets the accessPackageResources property value. Read-only. Required.
func (m *AccessPackageResourceEnvironment) SetAccessPackageResources(value []AccessPackageResourceable)() {
    m.accessPackageResources = value
}
// SetConnectionInfo sets the connectionInfo property value. Connection information of an environment used to connect to a resource.
func (m *AccessPackageResourceEnvironment) SetConnectionInfo(value ConnectionInfoable)() {
    m.connectionInfo = value
}
// SetCreatedBy sets the createdBy property value. The display name of the user that created this object.
func (m *AccessPackageResourceEnvironment) SetCreatedBy(value *string)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time that this object was created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AccessPackageResourceEnvironment) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description of this object.
func (m *AccessPackageResourceEnvironment) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of this object.
func (m *AccessPackageResourceEnvironment) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsDefaultEnvironment sets the isDefaultEnvironment property value. Determines whether this is default environment or not. It is set to true for all static origin systems, such as Azure AD groups and Azure AD Applications.
func (m *AccessPackageResourceEnvironment) SetIsDefaultEnvironment(value *bool)() {
    m.isDefaultEnvironment = value
}
// SetModifiedBy sets the modifiedBy property value. The display name of the entity that last modified this object.
func (m *AccessPackageResourceEnvironment) SetModifiedBy(value *string)() {
    m.modifiedBy = value
}
// SetModifiedDateTime sets the modifiedDateTime property value. The date and time that this object was last modified. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AccessPackageResourceEnvironment) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// SetOriginId sets the originId property value. The unique identifier of this environment in the origin system.
func (m *AccessPackageResourceEnvironment) SetOriginId(value *string)() {
    m.originId = value
}
// SetOriginSystem sets the originSystem property value. The type of the resource in the origin system, that is, SharePointOnline. Requires $filter (eq).
func (m *AccessPackageResourceEnvironment) SetOriginSystem(value *string)() {
    m.originSystem = value
}
