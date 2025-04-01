package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrganizationalUnit struct {
    DirectoryObject
}
// NewOrganizationalUnit instantiates a new OrganizationalUnit and sets the default values.
func NewOrganizationalUnit()(*OrganizationalUnit) {
    m := &OrganizationalUnit{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.organizationalUnit"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOrganizationalUnitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationalUnitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationalUnit(), nil
}
// GetChildren gets the children property value. The children property
// returns a []OrganizationalUnitable when successful
func (m *OrganizationalUnit) GetChildren()([]OrganizationalUnitable) {
    val, err := m.GetBackingStore().Get("children")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OrganizationalUnitable)
    }
    return nil
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *OrganizationalUnit) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *OrganizationalUnit) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationalUnit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["children"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOrganizationalUnitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OrganizationalUnitable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OrganizationalUnitable)
                }
            }
            m.SetChildren(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["organizationalUnitParent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationalUnitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalUnitParent(val.(OrganizationalUnitable))
        }
        return nil
    }
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DirectoryObjectable)
                }
            }
            m.SetResources(res)
        }
        return nil
    }
    res["transitiveChildren"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOrganizationalUnitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OrganizationalUnitable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OrganizationalUnitable)
                }
            }
            m.SetTransitiveChildren(res)
        }
        return nil
    }
    res["transitiveResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DirectoryObjectable)
                }
            }
            m.SetTransitiveResources(res)
        }
        return nil
    }
    return res
}
// GetOrganizationalUnitParent gets the organizationalUnitParent property value. The organizationalUnitParent property
// returns a OrganizationalUnitable when successful
func (m *OrganizationalUnit) GetOrganizationalUnitParent()(OrganizationalUnitable) {
    val, err := m.GetBackingStore().Get("organizationalUnitParent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OrganizationalUnitable)
    }
    return nil
}
// GetResources gets the resources property value. The resources property
// returns a []DirectoryObjectable when successful
func (m *OrganizationalUnit) GetResources()([]DirectoryObjectable) {
    val, err := m.GetBackingStore().Get("resources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DirectoryObjectable)
    }
    return nil
}
// GetTransitiveChildren gets the transitiveChildren property value. The transitiveChildren property
// returns a []OrganizationalUnitable when successful
func (m *OrganizationalUnit) GetTransitiveChildren()([]OrganizationalUnitable) {
    val, err := m.GetBackingStore().Get("transitiveChildren")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OrganizationalUnitable)
    }
    return nil
}
// GetTransitiveResources gets the transitiveResources property value. The transitiveResources property
// returns a []DirectoryObjectable when successful
func (m *OrganizationalUnit) GetTransitiveResources()([]DirectoryObjectable) {
    val, err := m.GetBackingStore().Get("transitiveResources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DirectoryObjectable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OrganizationalUnit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChildren() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChildren()))
        for i, v := range m.GetChildren() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("children", cast)
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
        err = writer.WriteObjectValue("organizationalUnitParent", m.GetOrganizationalUnitParent())
        if err != nil {
            return err
        }
    }
    if m.GetResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTransitiveChildren() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTransitiveChildren()))
        for i, v := range m.GetTransitiveChildren() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("transitiveChildren", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTransitiveResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTransitiveResources()))
        for i, v := range m.GetTransitiveResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("transitiveResources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChildren sets the children property value. The children property
func (m *OrganizationalUnit) SetChildren(value []OrganizationalUnitable)() {
    err := m.GetBackingStore().Set("children", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *OrganizationalUnit) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *OrganizationalUnit) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetOrganizationalUnitParent sets the organizationalUnitParent property value. The organizationalUnitParent property
func (m *OrganizationalUnit) SetOrganizationalUnitParent(value OrganizationalUnitable)() {
    err := m.GetBackingStore().Set("organizationalUnitParent", value)
    if err != nil {
        panic(err)
    }
}
// SetResources sets the resources property value. The resources property
func (m *OrganizationalUnit) SetResources(value []DirectoryObjectable)() {
    err := m.GetBackingStore().Set("resources", value)
    if err != nil {
        panic(err)
    }
}
// SetTransitiveChildren sets the transitiveChildren property value. The transitiveChildren property
func (m *OrganizationalUnit) SetTransitiveChildren(value []OrganizationalUnitable)() {
    err := m.GetBackingStore().Set("transitiveChildren", value)
    if err != nil {
        panic(err)
    }
}
// SetTransitiveResources sets the transitiveResources property value. The transitiveResources property
func (m *OrganizationalUnit) SetTransitiveResources(value []DirectoryObjectable)() {
    err := m.GetBackingStore().Set("transitiveResources", value)
    if err != nil {
        panic(err)
    }
}
type OrganizationalUnitable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChildren()([]OrganizationalUnitable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetOrganizationalUnitParent()(OrganizationalUnitable)
    GetResources()([]DirectoryObjectable)
    GetTransitiveChildren()([]OrganizationalUnitable)
    GetTransitiveResources()([]DirectoryObjectable)
    SetChildren(value []OrganizationalUnitable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetOrganizationalUnitParent(value OrganizationalUnitable)()
    SetResources(value []DirectoryObjectable)()
    SetTransitiveChildren(value []OrganizationalUnitable)()
    SetTransitiveResources(value []DirectoryObjectable)()
}
