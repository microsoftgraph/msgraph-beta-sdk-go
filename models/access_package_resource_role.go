package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageResourceRole provides operations to manage the collection of accessReview entities.
type AccessPackageResourceRole struct {
    Entity
    // The accessPackageResource property
    accessPackageResource AccessPackageResourceable
    // A description for the resource role.
    description *string
    // The display name of the resource role such as the role defined by the application.
    displayName *string
    // The unique identifier of the resource role in the origin system. For a SharePoint Online site, the originId will be the sequence number of the role in the site.
    originId *string
    // The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
    originSystem *string
}
// NewAccessPackageResourceRole instantiates a new accessPackageResourceRole and sets the default values.
func NewAccessPackageResourceRole()(*AccessPackageResourceRole) {
    m := &AccessPackageResourceRole{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.accessPackageResourceRole";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAccessPackageResourceRoleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageResourceRoleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageResourceRole(), nil
}
// GetAccessPackageResource gets the accessPackageResource property value. The accessPackageResource property
func (m *AccessPackageResourceRole) GetAccessPackageResource()(AccessPackageResourceable) {
    return m.accessPackageResource
}
// GetDescription gets the description property value. A description for the resource role.
func (m *AccessPackageResourceRole) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the resource role such as the role defined by the application.
func (m *AccessPackageResourceRole) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageResourceRole) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageResource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccessPackageResourceFromDiscriminatorValue , m.SetAccessPackageResource)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["originId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOriginId)
    res["originSystem"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOriginSystem)
    return res
}
// GetOriginId gets the originId property value. The unique identifier of the resource role in the origin system. For a SharePoint Online site, the originId will be the sequence number of the role in the site.
func (m *AccessPackageResourceRole) GetOriginId()(*string) {
    return m.originId
}
// GetOriginSystem gets the originSystem property value. The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
func (m *AccessPackageResourceRole) GetOriginSystem()(*string) {
    return m.originSystem
}
// Serialize serializes information the current object
func (m *AccessPackageResourceRole) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAccessPackageResource sets the accessPackageResource property value. The accessPackageResource property
func (m *AccessPackageResourceRole) SetAccessPackageResource(value AccessPackageResourceable)() {
    m.accessPackageResource = value
}
// SetDescription sets the description property value. A description for the resource role.
func (m *AccessPackageResourceRole) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the resource role such as the role defined by the application.
func (m *AccessPackageResourceRole) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOriginId sets the originId property value. The unique identifier of the resource role in the origin system. For a SharePoint Online site, the originId will be the sequence number of the role in the site.
func (m *AccessPackageResourceRole) SetOriginId(value *string)() {
    m.originId = value
}
// SetOriginSystem sets the originSystem property value. The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
func (m *AccessPackageResourceRole) SetOriginSystem(value *string)() {
    m.originSystem = value
}
