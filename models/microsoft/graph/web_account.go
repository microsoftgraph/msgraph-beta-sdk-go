package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new webAccount and sets the default values.
func NewWebAccount()(*WebAccount) {
    m := &WebAccount{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// Gets the description property value. Contains the description the user has provided for the account on the service being referenced.
func (m *WebAccount) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the service property value. 
func (m *WebAccount) GetService()(*ServiceInformation) {
    if m == nil {
        return nil
    } else {
        return m.service
    }
}
// Gets the statusMessage property value. Contains a status message from the cloud service if provided or synchronized.
func (m *WebAccount) GetStatusMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.statusMessage
    }
}
// Gets the thumbnailUrl property value. 
func (m *WebAccount) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
// Gets the userId property value. The user name  displayed for the webaccount.
func (m *WebAccount) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Gets the webUrl property value. Contains a link to the user's profile on the cloud service if one exists.
func (m *WebAccount) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the description property value. Contains the description the user has provided for the account on the service being referenced.
// Parameters:
//  - value : Value to set for the description property.
func (m *WebAccount) SetDescription(value *string)() {
    m.description = value
}
// Sets the service property value. 
// Parameters:
//  - value : Value to set for the service property.
func (m *WebAccount) SetService(value *ServiceInformation)() {
    m.service = value
}
// Sets the statusMessage property value. Contains a status message from the cloud service if provided or synchronized.
// Parameters:
//  - value : Value to set for the statusMessage property.
func (m *WebAccount) SetStatusMessage(value *string)() {
    m.statusMessage = value
}
// Sets the thumbnailUrl property value. 
// Parameters:
//  - value : Value to set for the thumbnailUrl property.
func (m *WebAccount) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// Sets the userId property value. The user name  displayed for the webaccount.
// Parameters:
//  - value : Value to set for the userId property.
func (m *WebAccount) SetUserId(value *string)() {
    m.userId = value
}
// Sets the webUrl property value. Contains a link to the user's profile on the cloud service if one exists.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *WebAccount) SetWebUrl(value *string)() {
    m.webUrl = value
}
