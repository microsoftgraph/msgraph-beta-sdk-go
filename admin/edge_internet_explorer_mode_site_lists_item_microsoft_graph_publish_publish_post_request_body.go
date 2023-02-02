package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// EdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody 
type EdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The revision property
    revision *string
    // The sharedCookies property
    sharedCookies []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSharedCookieable
    // The sites property
    sites []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteable
}
// NewEdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody instantiates a new EdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody and sets the default values.
func NewEdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody()(*EdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody) {
    m := &EdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["revision"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRevision(val)
        }
        return nil
    }
    res["sharedCookies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBrowserSharedCookieFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSharedCookieable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSharedCookieable)
            }
            m.SetSharedCookies(res)
        }
        return nil
    }
    res["sites"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBrowserSiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteable)
            }
            m.SetSites(res)
        }
        return nil
    }
    return res
}
// GetRevision gets the revision property value. The revision property
func (m *EdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody) GetRevision()(*string) {
    return m.revision
}
// GetSharedCookies gets the sharedCookies property value. The sharedCookies property
func (m *EdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody) GetSharedCookies()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSharedCookieable) {
    return m.sharedCookies
}
// GetSites gets the sites property value. The sites property
func (m *EdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody) GetSites()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteable) {
    return m.sites
}
// Serialize serializes information the current object
func (m *EdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("revision", m.GetRevision())
        if err != nil {
            return err
        }
    }
    if m.GetSharedCookies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSharedCookies()))
        for i, v := range m.GetSharedCookies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("sharedCookies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSites() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSites()))
        for i, v := range m.GetSites() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("sites", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRevision sets the revision property value. The revision property
func (m *EdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody) SetRevision(value *string)() {
    m.revision = value
}
// SetSharedCookies sets the sharedCookies property value. The sharedCookies property
func (m *EdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody) SetSharedCookies(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSharedCookieable)() {
    m.sharedCookies = value
}
// SetSites sets the sites property value. The sites property
func (m *EdgeInternetExplorerModeSiteListsItemMicrosoftGraphPublishPublishPostRequestBody) SetSites(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteable)() {
    m.sites = value
}
