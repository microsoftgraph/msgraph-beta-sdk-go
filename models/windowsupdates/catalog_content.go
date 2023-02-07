package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CatalogContent 
type CatalogContent struct {
    DeployableContent
    // The catalogEntry property
    catalogEntry CatalogEntryable
}
// NewCatalogContent instantiates a new CatalogContent and sets the default values.
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
    return m.catalogEntry
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
    return res
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
    return nil
}
// SetCatalogEntry sets the catalogEntry property value. The catalogEntry property
func (m *CatalogContent) SetCatalogEntry(value CatalogEntryable)() {
    m.catalogEntry = value
}
