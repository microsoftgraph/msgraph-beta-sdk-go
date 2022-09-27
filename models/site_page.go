package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SitePage provides operations to manage the collection of accessReviewDecision entities.
type SitePage struct {
    BaseItem
    // The content type of the page.
    contentType ContentTypeInfoable
    // The pageLayoutType property
    pageLayoutType *string
    // The publishingState property
    publishingState PublicationFacetable
    // The title property
    title *string
    // The webParts property
    webParts []WebPartable
}
// NewSitePage instantiates a new sitePage and sets the default values.
func NewSitePage()(*SitePage) {
    m := &SitePage{
        BaseItem: *NewBaseItem(),
    }
    odataTypeValue := "#microsoft.graph.sitePage";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSitePageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSitePageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSitePage(), nil
}
// GetContentType gets the contentType property value. The content type of the page.
func (m *SitePage) GetContentType()(ContentTypeInfoable) {
    return m.contentType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SitePage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["contentType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateContentTypeInfoFromDiscriminatorValue , m.SetContentType)
    res["pageLayoutType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPageLayoutType)
    res["publishingState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePublicationFacetFromDiscriminatorValue , m.SetPublishingState)
    res["title"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTitle)
    res["webParts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWebPartFromDiscriminatorValue , m.SetWebParts)
    return res
}
// GetPageLayoutType gets the pageLayoutType property value. The pageLayoutType property
func (m *SitePage) GetPageLayoutType()(*string) {
    return m.pageLayoutType
}
// GetPublishingState gets the publishingState property value. The publishingState property
func (m *SitePage) GetPublishingState()(PublicationFacetable) {
    return m.publishingState
}
// GetTitle gets the title property value. The title property
func (m *SitePage) GetTitle()(*string) {
    return m.title
}
// GetWebParts gets the webParts property value. The webParts property
func (m *SitePage) GetWebParts()([]WebPartable) {
    return m.webParts
}
// Serialize serializes information the current object
func (m *SitePage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("pageLayoutType", m.GetPageLayoutType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("publishingState", m.GetPublishingState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    if m.GetWebParts() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWebParts())
        err = writer.WriteCollectionOfObjectValues("webParts", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentType sets the contentType property value. The content type of the page.
func (m *SitePage) SetContentType(value ContentTypeInfoable)() {
    m.contentType = value
}
// SetPageLayoutType sets the pageLayoutType property value. The pageLayoutType property
func (m *SitePage) SetPageLayoutType(value *string)() {
    m.pageLayoutType = value
}
// SetPublishingState sets the publishingState property value. The publishingState property
func (m *SitePage) SetPublishingState(value PublicationFacetable)() {
    m.publishingState = value
}
// SetTitle sets the title property value. The title property
func (m *SitePage) SetTitle(value *string)() {
    m.title = value
}
// SetWebParts sets the webParts property value. The webParts property
func (m *SitePage) SetWebParts(value []WebPartable)() {
    m.webParts = value
}
