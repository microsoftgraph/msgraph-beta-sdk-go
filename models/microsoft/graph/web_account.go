package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WebAccount 
type WebAccount struct {
    ItemFacet
    // Contains the description the user has provided for the account on the service being referenced.
    description *string;
    // 
    service *ServiceInformation;
    // Contains a status message from the cloud service if provided or synchronized.
    statusMessage *string;
    // 
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
// GetDescription gets the description property value. Contains the description the user has provided for the account on the service being referenced.
func (m *WebAccount) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetService gets the service property value. 
func (m *WebAccount) GetService()(*ServiceInformation) {
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
// GetThumbnailUrl gets the thumbnailUrl property value. 
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
// GetFieldDeserializers the deserialization information for the current model
func (m *WebAccount) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["service"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceInformation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val.(*ServiceInformation))
        }
        return nil
    }
    res["statusMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusMessage(val)
        }
        return nil
    }
    res["thumbnailUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailUrl(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *WebAccount) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WebAccount) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetService sets the service property value. 
func (m *WebAccount) SetService(value *ServiceInformation)() {
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
// SetThumbnailUrl sets the thumbnailUrl property value. 
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
