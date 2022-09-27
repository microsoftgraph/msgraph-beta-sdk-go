package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CartToClassAssociation cartToClassAssociation for associating device carts with classrooms.
type CartToClassAssociation struct {
    Entity
    // Identifiers of classrooms to be associated with device carts.
    classroomIds []string
    // DateTime the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Admin provided description of the CartToClassAssociation.
    description *string
    // Identifiers of device carts to be associated with classes.
    deviceCartIds []string
    // Admin provided name of the device configuration.
    displayName *string
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Version of the CartToClassAssociation.
    version *int32
}
// NewCartToClassAssociation instantiates a new cartToClassAssociation and sets the default values.
func NewCartToClassAssociation()(*CartToClassAssociation) {
    m := &CartToClassAssociation{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.cartToClassAssociation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCartToClassAssociationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCartToClassAssociationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCartToClassAssociation(), nil
}
// GetClassroomIds gets the classroomIds property value. Identifiers of classrooms to be associated with device carts.
func (m *CartToClassAssociation) GetClassroomIds()([]string) {
    return m.classroomIds
}
// GetCreatedDateTime gets the createdDateTime property value. DateTime the object was created.
func (m *CartToClassAssociation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Admin provided description of the CartToClassAssociation.
func (m *CartToClassAssociation) GetDescription()(*string) {
    return m.description
}
// GetDeviceCartIds gets the deviceCartIds property value. Identifiers of device carts to be associated with classes.
func (m *CartToClassAssociation) GetDeviceCartIds()([]string) {
    return m.deviceCartIds
}
// GetDisplayName gets the displayName property value. Admin provided name of the device configuration.
func (m *CartToClassAssociation) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CartToClassAssociation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classroomIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetClassroomIds)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["deviceCartIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetDeviceCartIds)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["version"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetVersion)
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *CartToClassAssociation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetVersion gets the version property value. Version of the CartToClassAssociation.
func (m *CartToClassAssociation) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *CartToClassAssociation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassroomIds() != nil {
        err = writer.WriteCollectionOfStringValues("classroomIds", m.GetClassroomIds())
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
    if m.GetDeviceCartIds() != nil {
        err = writer.WriteCollectionOfStringValues("deviceCartIds", m.GetDeviceCartIds())
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassroomIds sets the classroomIds property value. Identifiers of classrooms to be associated with device carts.
func (m *CartToClassAssociation) SetClassroomIds(value []string)() {
    m.classroomIds = value
}
// SetCreatedDateTime sets the createdDateTime property value. DateTime the object was created.
func (m *CartToClassAssociation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Admin provided description of the CartToClassAssociation.
func (m *CartToClassAssociation) SetDescription(value *string)() {
    m.description = value
}
// SetDeviceCartIds sets the deviceCartIds property value. Identifiers of device carts to be associated with classes.
func (m *CartToClassAssociation) SetDeviceCartIds(value []string)() {
    m.deviceCartIds = value
}
// SetDisplayName sets the displayName property value. Admin provided name of the device configuration.
func (m *CartToClassAssociation) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *CartToClassAssociation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetVersion sets the version property value. Version of the CartToClassAssociation.
func (m *CartToClassAssociation) SetVersion(value *int32)() {
    m.version = value
}
