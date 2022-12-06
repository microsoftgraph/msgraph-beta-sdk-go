package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InternetExplorerMode 
type InternetExplorerMode struct {
    Entity
    // A collection of site lists to support Internet Explorer mode.
    siteLists []BrowserSiteListable
}
// NewInternetExplorerMode instantiates a new internetExplorerMode and sets the default values.
func NewInternetExplorerMode()(*InternetExplorerMode) {
    m := &InternetExplorerMode{
        Entity: *NewEntity(),
    }
    return m
}
// CreateInternetExplorerModeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInternetExplorerModeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInternetExplorerMode(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InternetExplorerMode) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["siteLists"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBrowserSiteListFromDiscriminatorValue , m.SetSiteLists)
    return res
}
// GetSiteLists gets the siteLists property value. A collection of site lists to support Internet Explorer mode.
func (m *InternetExplorerMode) GetSiteLists()([]BrowserSiteListable) {
    return m.siteLists
}
// Serialize serializes information the current object
func (m *InternetExplorerMode) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSiteLists() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSiteLists())
        err = writer.WriteCollectionOfObjectValues("siteLists", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSiteLists sets the siteLists property value. A collection of site lists to support Internet Explorer mode.
func (m *InternetExplorerMode) SetSiteLists(value []BrowserSiteListable)() {
    m.siteLists = value
}
