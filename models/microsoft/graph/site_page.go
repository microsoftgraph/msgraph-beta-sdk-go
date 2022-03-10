package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SitePage provides operations to manage the compliance singleton.
type SitePage struct {
    BaseItem
    // The content type of the page.
    contentType ContentTypeInfoable;
    // 
    pageLayoutType *string;
    // 
    publishingState PublicationFacetable;
    // 
    title *string;
    // 
    webParts []WebPartable;
}
// NewSitePage instantiates a new sitePage and sets the default values.
func NewSitePage()(*SitePage) {
    m := &SitePage{
        BaseItem: *NewBaseItem(),
    }
    return m
}
// CreateSitePageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSitePageFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSitePage(), nil
}
// GetContentType gets the contentType property value. The content type of the page.
func (m *SitePage) GetContentType()(ContentTypeInfoable) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SitePage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateContentTypeInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val.(ContentTypeInfoable))
        }
        return nil
    }
    res["pageLayoutType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPageLayoutType(val)
        }
        return nil
    }
    res["publishingState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicationFacetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishingState(val.(PublicationFacetable))
        }
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["webParts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWebPartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WebPartable, len(val))
            for i, v := range val {
                res[i] = v.(WebPartable)
            }
            m.SetWebParts(res)
        }
        return nil
    }
    return res
}
// GetPageLayoutType gets the pageLayoutType property value. 
func (m *SitePage) GetPageLayoutType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.pageLayoutType
    }
}
// GetPublishingState gets the publishingState property value. 
func (m *SitePage) GetPublishingState()(PublicationFacetable) {
    if m == nil {
        return nil
    } else {
        return m.publishingState
    }
}
// GetTitle gets the title property value. 
func (m *SitePage) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// GetWebParts gets the webParts property value. 
func (m *SitePage) GetWebParts()([]WebPartable) {
    if m == nil {
        return nil
    } else {
        return m.webParts
    }
}
func (m *SitePage) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SitePage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWebParts()))
        for i, v := range m.GetWebParts() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("webParts", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentType sets the contentType property value. The content type of the page.
func (m *SitePage) SetContentType(value ContentTypeInfoable)() {
    if m != nil {
        m.contentType = value
    }
}
// SetPageLayoutType sets the pageLayoutType property value. 
func (m *SitePage) SetPageLayoutType(value *string)() {
    if m != nil {
        m.pageLayoutType = value
    }
}
// SetPublishingState sets the publishingState property value. 
func (m *SitePage) SetPublishingState(value PublicationFacetable)() {
    if m != nil {
        m.publishingState = value
    }
}
// SetTitle sets the title property value. 
func (m *SitePage) SetTitle(value *string)() {
    if m != nil {
        m.title = value
    }
}
// SetWebParts sets the webParts property value. 
func (m *SitePage) SetWebParts(value []WebPartable)() {
    if m != nil {
        m.webParts = value
    }
}
