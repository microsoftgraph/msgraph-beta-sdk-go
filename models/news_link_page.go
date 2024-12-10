package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NewsLinkPage struct {
    BaseSitePage
}
// NewNewsLinkPage instantiates a new NewsLinkPage and sets the default values.
func NewNewsLinkPage()(*NewsLinkPage) {
    m := &NewsLinkPage{
        BaseSitePage: *NewBaseSitePage(),
    }
    return m
}
// CreateNewsLinkPageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNewsLinkPageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNewsLinkPage(), nil
}
// GetBannerImageWebUrl gets the bannerImageWebUrl property value. A link to the banner image for the newsLinkPage.
// returns a *string when successful
func (m *NewsLinkPage) GetBannerImageWebUrl()(*string) {
    val, err := m.GetBackingStore().Get("bannerImageWebUrl")
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
func (m *NewsLinkPage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseSitePage.GetFieldDeserializers()
    res["bannerImageWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBannerImageWebUrl(val)
        }
        return nil
    }
    res["newsSharepointIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharepointIdsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewsSharepointIds(val.(SharepointIdsable))
        }
        return nil
    }
    res["newsWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewsWebUrl(val)
        }
        return nil
    }
    return res
}
// GetNewsSharepointIds gets the newsSharepointIds property value. The SharePoint IDs of the referenced news article if it's recognized as a SharePoint resource. Read-only.
// returns a SharepointIdsable when successful
func (m *NewsLinkPage) GetNewsSharepointIds()(SharepointIdsable) {
    val, err := m.GetBackingStore().Get("newsSharepointIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharepointIdsable)
    }
    return nil
}
// GetNewsWebUrl gets the newsWebUrl property value. The URL of the news article referenced by the newsLinkPage. It can be an external link.
// returns a *string when successful
func (m *NewsLinkPage) GetNewsWebUrl()(*string) {
    val, err := m.GetBackingStore().Get("newsWebUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *NewsLinkPage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseSitePage.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("bannerImageWebUrl", m.GetBannerImageWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("newsSharepointIds", m.GetNewsSharepointIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("newsWebUrl", m.GetNewsWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBannerImageWebUrl sets the bannerImageWebUrl property value. A link to the banner image for the newsLinkPage.
func (m *NewsLinkPage) SetBannerImageWebUrl(value *string)() {
    err := m.GetBackingStore().Set("bannerImageWebUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetNewsSharepointIds sets the newsSharepointIds property value. The SharePoint IDs of the referenced news article if it's recognized as a SharePoint resource. Read-only.
func (m *NewsLinkPage) SetNewsSharepointIds(value SharepointIdsable)() {
    err := m.GetBackingStore().Set("newsSharepointIds", value)
    if err != nil {
        panic(err)
    }
}
// SetNewsWebUrl sets the newsWebUrl property value. The URL of the news article referenced by the newsLinkPage. It can be an external link.
func (m *NewsLinkPage) SetNewsWebUrl(value *string)() {
    err := m.GetBackingStore().Set("newsWebUrl", value)
    if err != nil {
        panic(err)
    }
}
type NewsLinkPageable interface {
    BaseSitePageable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBannerImageWebUrl()(*string)
    GetNewsSharepointIds()(SharepointIdsable)
    GetNewsWebUrl()(*string)
    SetBannerImageWebUrl(value *string)()
    SetNewsSharepointIds(value SharepointIdsable)()
    SetNewsWebUrl(value *string)()
}
