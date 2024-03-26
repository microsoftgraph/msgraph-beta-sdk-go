package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type Product struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewProduct instantiates a new Product and sets the default values.
func NewProduct()(*Product) {
    m := &Product{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateProductFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProductFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProduct(), nil
}
// GetEditions gets the editions property value. Represents an edition of a particular Windows product.
// returns a []Editionable when successful
func (m *Product) GetEditions()([]Editionable) {
    val, err := m.GetBackingStore().Get("editions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Editionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Product) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["editions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEditionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Editionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Editionable)
                }
            }
            m.SetEditions(res)
        }
        return nil
    }
    res["friendlyNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetFriendlyNames(res)
        }
        return nil
    }
    res["groupName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupName(val)
        }
        return nil
    }
    res["knownIssues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKnownIssueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KnownIssueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(KnownIssueable)
                }
            }
            m.SetKnownIssues(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["revisions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProductRevisionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProductRevisionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ProductRevisionable)
                }
            }
            m.SetRevisions(res)
        }
        return nil
    }
    return res
}
// GetFriendlyNames gets the friendlyNames property value. The friendly names of the product. For example, Version 22H2 (OS build 22621). Read-only.
// returns a []string when successful
func (m *Product) GetFriendlyNames()([]string) {
    val, err := m.GetBackingStore().Get("friendlyNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetGroupName gets the groupName property value. The name of the product group. For example, Windows 11. Read-only.
// returns a *string when successful
func (m *Product) GetGroupName()(*string) {
    val, err := m.GetBackingStore().Get("groupName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetKnownIssues gets the knownIssues property value. Represents a known issue related to a Windows product.
// returns a []KnownIssueable when successful
func (m *Product) GetKnownIssues()([]KnownIssueable) {
    val, err := m.GetBackingStore().Get("knownIssues")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KnownIssueable)
    }
    return nil
}
// GetName gets the name property value. The name of the product. For example, Windows 11, version 22H2. Read-only.
// returns a *string when successful
func (m *Product) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRevisions gets the revisions property value. Represents a product revision.
// returns a []ProductRevisionable when successful
func (m *Product) GetRevisions()([]ProductRevisionable) {
    val, err := m.GetBackingStore().Get("revisions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ProductRevisionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Product) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEditions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEditions()))
        for i, v := range m.GetEditions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("editions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFriendlyNames() != nil {
        err = writer.WriteCollectionOfStringValues("friendlyNames", m.GetFriendlyNames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupName", m.GetGroupName())
        if err != nil {
            return err
        }
    }
    if m.GetKnownIssues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKnownIssues()))
        for i, v := range m.GetKnownIssues() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("knownIssues", cast)
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
    if m.GetRevisions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRevisions()))
        for i, v := range m.GetRevisions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("revisions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEditions sets the editions property value. Represents an edition of a particular Windows product.
func (m *Product) SetEditions(value []Editionable)() {
    err := m.GetBackingStore().Set("editions", value)
    if err != nil {
        panic(err)
    }
}
// SetFriendlyNames sets the friendlyNames property value. The friendly names of the product. For example, Version 22H2 (OS build 22621). Read-only.
func (m *Product) SetFriendlyNames(value []string)() {
    err := m.GetBackingStore().Set("friendlyNames", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupName sets the groupName property value. The name of the product group. For example, Windows 11. Read-only.
func (m *Product) SetGroupName(value *string)() {
    err := m.GetBackingStore().Set("groupName", value)
    if err != nil {
        panic(err)
    }
}
// SetKnownIssues sets the knownIssues property value. Represents a known issue related to a Windows product.
func (m *Product) SetKnownIssues(value []KnownIssueable)() {
    err := m.GetBackingStore().Set("knownIssues", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The name of the product. For example, Windows 11, version 22H2. Read-only.
func (m *Product) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetRevisions sets the revisions property value. Represents a product revision.
func (m *Product) SetRevisions(value []ProductRevisionable)() {
    err := m.GetBackingStore().Set("revisions", value)
    if err != nil {
        panic(err)
    }
}
type Productable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEditions()([]Editionable)
    GetFriendlyNames()([]string)
    GetGroupName()(*string)
    GetKnownIssues()([]KnownIssueable)
    GetName()(*string)
    GetRevisions()([]ProductRevisionable)
    SetEditions(value []Editionable)()
    SetFriendlyNames(value []string)()
    SetGroupName(value *string)()
    SetKnownIssues(value []KnownIssueable)()
    SetName(value *string)()
    SetRevisions(value []ProductRevisionable)()
}
