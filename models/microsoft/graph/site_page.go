package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SitePage 
type SitePage struct {
    BaseItem
    // The content type of the page.
    contentType *ContentTypeInfo;
    // 
    pageLayoutType *string;
    // 
    publishingState *PublicationFacet;
    // 
    title *string;
    // 
    webParts []WebPart;
}
// NewSitePage instantiates a new sitePage and sets the default values.
func NewSitePage()(*SitePage) {
    m := &SitePage{
        BaseItem: *NewBaseItem(),
    }
    return m
}
// GetContentType gets the contentType property value. The content type of the page.
func (m *SitePage) GetContentType()(*ContentTypeInfo) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
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
func (m *SitePage) GetPublishingState()(*PublicationFacet) {
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
func (m *SitePage) GetWebParts()([]WebPart) {
    if m == nil {
        return nil
    } else {
        return m.webParts
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SitePage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentTypeInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val.(*ContentTypeInfo))
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPublicationFacet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishingState(val.(*PublicationFacet))
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
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWebPart() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WebPart, len(val))
            for i, v := range val {
                res[i] = *(v.(*WebPart))
            }
            m.SetWebParts(res)
        }
        return nil
    }
    return res
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
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("webParts", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentType sets the contentType property value. The content type of the page.
func (m *SitePage) SetContentType(value *ContentTypeInfo)() {
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
func (m *SitePage) SetPublishingState(value *PublicationFacet)() {
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
func (m *SitePage) SetWebParts(value []WebPart)() {
    if m != nil {
        m.webParts = value
    }
}
