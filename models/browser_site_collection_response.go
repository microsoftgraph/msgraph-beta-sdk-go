package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BrowserSiteCollectionResponse provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type BrowserSiteCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []BrowserSiteable
}
// NewBrowserSiteCollectionResponse instantiates a new BrowserSiteCollectionResponse and sets the default values.
func NewBrowserSiteCollectionResponse()(*BrowserSiteCollectionResponse) {
    m := &BrowserSiteCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateBrowserSiteCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBrowserSiteCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBrowserSiteCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BrowserSiteCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBrowserSiteFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *BrowserSiteCollectionResponse) GetValue()([]BrowserSiteable) {
    return m.value
}
// Serialize serializes information the current object
func (m *BrowserSiteCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValue())
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *BrowserSiteCollectionResponse) SetValue(value []BrowserSiteable)() {
    m.value = value
}
