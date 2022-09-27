package windowsupdates

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Catalog 
type Catalog struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Lists the content that you can approve for deployment. Read-only.
    entries []CatalogEntryable
}
// NewCatalog instantiates a new catalog and sets the default values.
func NewCatalog()(*Catalog) {
    m := &Catalog{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.catalog";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCatalogFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCatalogFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCatalog(), nil
}
// GetEntries gets the entries property value. Lists the content that you can approve for deployment. Read-only.
func (m *Catalog) GetEntries()([]CatalogEntryable) {
    return m.entries
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Catalog) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["entries"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCatalogEntryFromDiscriminatorValue , m.SetEntries)
    return res
}
// Serialize serializes information the current object
func (m *Catalog) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEntries() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEntries())
        err = writer.WriteCollectionOfObjectValues("entries", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEntries sets the entries property value. Lists the content that you can approve for deployment. Read-only.
func (m *Catalog) SetEntries(value []CatalogEntryable)() {
    m.entries = value
}
