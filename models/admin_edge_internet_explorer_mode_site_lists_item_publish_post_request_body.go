package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody provides operations to call the publish method.
type AdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The revision property
    revision *string
    // The sharedCookies property
    sharedCookies []BrowserSharedCookieable
    // The sites property
    sites []BrowserSiteable
}
// NewAdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody instantiates a new AdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody and sets the default values.
func NewAdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody()(*AdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody) {
    m := &AdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfObjectValues(CreateBrowserSharedCookieFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BrowserSharedCookieable, len(val))
            for i, v := range val {
                res[i] = v.(BrowserSharedCookieable)
            }
            m.SetSharedCookies(res)
        }
        return nil
    }
    res["sites"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBrowserSiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BrowserSiteable, len(val))
            for i, v := range val {
                res[i] = v.(BrowserSiteable)
            }
            m.SetSites(res)
        }
        return nil
    }
    return res
}
// GetRevision gets the revision property value. The revision property
func (m *AdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody) GetRevision()(*string) {
    return m.revision
}
// GetSharedCookies gets the sharedCookies property value. The sharedCookies property
func (m *AdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody) GetSharedCookies()([]BrowserSharedCookieable) {
    return m.sharedCookies
}
// GetSites gets the sites property value. The sites property
func (m *AdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody) GetSites()([]BrowserSiteable) {
    return m.sites
}
// Serialize serializes information the current object
func (m *AdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetRevision sets the revision property value. The revision property
func (m *AdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody) SetRevision(value *string)() {
    m.revision = value
}
// SetSharedCookies sets the sharedCookies property value. The sharedCookies property
func (m *AdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody) SetSharedCookies(value []BrowserSharedCookieable)() {
    m.sharedCookies = value
}
// SetSites sets the sites property value. The sites property
func (m *AdminEdgeInternetExplorerModeSiteListsItemPublishPostRequestBody) SetSites(value []BrowserSiteable)() {
    m.sites = value
}
