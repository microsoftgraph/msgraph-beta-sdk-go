package ediscovery

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type SiteSource struct {
    DataSource
}
// NewSiteSource instantiates a new SiteSource and sets the default values.
func NewSiteSource()(*SiteSource) {
    m := &SiteSource{
        DataSource: *NewDataSource(),
    }
    odataTypeValue := "#microsoft.graph.ediscovery.siteSource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSiteSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSiteSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSiteSource(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SiteSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DataSource.GetFieldDeserializers()
    res["site"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSite(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable))
        }
        return nil
    }
    return res
}
// GetSite gets the site property value. The site property
// returns a Siteable when successful
func (m *SiteSource) GetSite()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable) {
    val, err := m.GetBackingStore().Get("site")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SiteSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DataSource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("site", m.GetSite())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSite sets the site property value. The site property
func (m *SiteSource) SetSite(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable)() {
    err := m.GetBackingStore().Set("site", value)
    if err != nil {
        panic(err)
    }
}
type SiteSourceable interface {
    DataSourceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSite()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable)
    SetSite(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable)()
}
