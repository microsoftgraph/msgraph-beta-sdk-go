package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ProjectParticipation struct {
    ItemFacet
    categories []string;
    client *CompanyDetail;
    collaborationTags []string;
    colleagues []RelatedPerson;
    detail *PositionDetail;
    displayName *string;
    sponsors []RelatedPerson;
    thumbnailUrl *string;
}
func NewProjectParticipation()(*ProjectParticipation) {
    m := &ProjectParticipation{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
func (m *ProjectParticipation) GetCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
func (m *ProjectParticipation) GetClient()(*CompanyDetail) {
    if m == nil {
        return nil
    } else {
        return m.client
    }
}
func (m *ProjectParticipation) GetCollaborationTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.collaborationTags
    }
}
func (m *ProjectParticipation) GetColleagues()([]RelatedPerson) {
    if m == nil {
        return nil
    } else {
        return m.colleagues
    }
}
func (m *ProjectParticipation) GetDetail()(*PositionDetail) {
    if m == nil {
        return nil
    } else {
        return m.detail
    }
}
func (m *ProjectParticipation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *ProjectParticipation) GetSponsors()([]RelatedPerson) {
    if m == nil {
        return nil
    } else {
        return m.sponsors
    }
}
func (m *ProjectParticipation) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
func (m *ProjectParticipation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetCategories(res)
        return nil
    }
    res["client"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCompanyDetail() })
        if err != nil {
            return err
        }
        m.SetClient(val.(*CompanyDetail))
        return nil
    }
    res["collaborationTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetCollaborationTags(res)
        return nil
    }
    res["colleagues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRelatedPerson() })
        if err != nil {
            return err
        }
        res := make([]RelatedPerson, len(val))
        for i, v := range val {
            res[i] = *(v.(*RelatedPerson))
        }
        m.SetColleagues(res)
        return nil
    }
    res["detail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPositionDetail() })
        if err != nil {
            return err
        }
        m.SetDetail(val.(*PositionDetail))
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["sponsors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRelatedPerson() })
        if err != nil {
            return err
        }
        res := make([]RelatedPerson, len(val))
        for i, v := range val {
            res[i] = *(v.(*RelatedPerson))
        }
        m.SetSponsors(res)
        return nil
    }
    res["thumbnailUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetThumbnailUrl(val)
        return nil
    }
    return res
}
func (m *ProjectParticipation) IsNil()(bool) {
    return m == nil
}
func (m *ProjectParticipation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("categories", m.GetCategories())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("client", m.GetClient())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("collaborationTags", m.GetCollaborationTags())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetColleagues()))
        for i, v := range m.GetColleagues() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("colleagues", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("detail", m.GetDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSponsors()))
        for i, v := range m.GetSponsors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("sponsors", cast)
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
    return nil
}
func (m *ProjectParticipation) SetCategories(value []string)() {
    m.categories = value
}
func (m *ProjectParticipation) SetClient(value *CompanyDetail)() {
    m.client = value
}
func (m *ProjectParticipation) SetCollaborationTags(value []string)() {
    m.collaborationTags = value
}
func (m *ProjectParticipation) SetColleagues(value []RelatedPerson)() {
    m.colleagues = value
}
func (m *ProjectParticipation) SetDetail(value *PositionDetail)() {
    m.detail = value
}
func (m *ProjectParticipation) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *ProjectParticipation) SetSponsors(value []RelatedPerson)() {
    m.sponsors = value
}
func (m *ProjectParticipation) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
