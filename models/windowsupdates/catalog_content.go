package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CatalogContent 
type CatalogContent struct {
    DeployableContent
}
// NewCatalogContent instantiates a new catalogContent and sets the default values.
func NewCatalogContent()(*CatalogContent) {
    m := &CatalogContent{
        DeployableContent: *NewDeployableContent(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.catalogContent"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCatalogContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCatalogContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCatalogContent(), nil
}
// GetCatalogEntry gets the catalogEntry property value. The catalogEntry property
func (m *CatalogContent) GetCatalogEntry()(CatalogEntryable) {
    val, err := m.GetBackingStore().Get("catalogEntry")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CatalogEntryable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CatalogContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeployableContent.GetFieldDeserializers()
    res["catalogEntry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCatalogEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogEntry(val.(CatalogEntryable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CatalogContent) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CatalogContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeployableContent.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("catalogEntry", m.GetCatalogEntry())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCatalogEntry sets the catalogEntry property value. The catalogEntry property
func (m *CatalogContent) SetCatalogEntry(value CatalogEntryable)() {
    err := m.GetBackingStore().Set("catalogEntry", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CatalogContent) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// CatalogContentable 
type CatalogContentable interface {
    DeployableContentable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCatalogEntry()(CatalogEntryable)
    GetOdataType()(*string)
    SetCatalogEntry(value CatalogEntryable)()
    SetOdataType(value *string)()
}
