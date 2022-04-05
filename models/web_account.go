package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebAccount 
type WebAccount struct {
    ItemFacet
    // Contains the description the user has provided for the account on the service being referenced.
    description *string;
    // The service property
    service ServiceInformationable;
    // Contains a status message from the cloud service if provided or synchronized.
    statusMessage *string;
    // The thumbnailUrl property
    thumbnailUrl *string;
    // The user name  displayed for the webaccount.
    userId *string;
    // Contains a link to the user's profile on the cloud service if one exists.
    webUrl *string;
}
// NewWebAccount instantiates a new webAccount and sets the default values.
func NewWebAccount()(*WebAccount) {
    m := &WebAccount{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// CreateWebAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebAccount(), nil
}
// GetDescription gets the description property value. Contains the description the user has provided for the account on the service being referenced.
func (m *WebAccount) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebAccount) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["service"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServiceInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val.(ServiceInformationable))
        }
        return nil
    }
    res["statusMessage"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusMessage(val)
        }
        return nil
    }
    res["thumbnailUrl"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailUrl(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["webUrl"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetService gets the service property value. The service property
func (m *WebAccount) GetService()(ServiceInformationable) {
    if m == nil {
        return nil
    } else {
        return m.service
    }
}
// GetStatusMessage gets the statusMessage property value. Contains a status message from the cloud service if provided or synchronized.
func (m *WebAccount) GetStatusMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.statusMessage
    }
}
// GetThumbnailUrl gets the thumbnailUrl property value. The thumbnailUrl property
func (m *WebAccount) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
// GetUserId gets the userId property value. The user name  displayed for the webaccount.
func (m *WebAccount) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetWebUrl gets the webUrl property value. Contains a link to the user's profile on the cloud service if one exists.
func (m *WebAccount) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// Serialize serializes information the current object
func (m *WebAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("statusMessage", m.GetStatusMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("thumbnailUrl", m.GetThumbnailUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Contains the description the user has provided for the account on the service being referenced.
func (m *WebAccount) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetService sets the service property value. The service property
func (m *WebAccount) SetService(value ServiceInformationable)() {
    if m != nil {
        m.service = value
    }
}
// SetStatusMessage sets the statusMessage property value. Contains a status message from the cloud service if provided or synchronized.
func (m *WebAccount) SetStatusMessage(value *string)() {
    if m != nil {
        m.statusMessage = value
    }
}
// SetThumbnailUrl sets the thumbnailUrl property value. The thumbnailUrl property
func (m *WebAccount) SetThumbnailUrl(value *string)() {
    if m != nil {
        m.thumbnailUrl = value
    }
}
// SetUserId sets the userId property value. The user name  displayed for the webaccount.
func (m *WebAccount) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
// SetWebUrl sets the webUrl property value. Contains a link to the user's profile on the cloud service if one exists.
func (m *WebAccount) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
