package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProjectParticipation 
type ProjectParticipation struct {
    ItemFacet
    // Contains categories a user has associated with the project (for example, digital transformation, oil rig).
    categories []string;
    // Contains detailed information about the client the project was for.
    client *CompanyDetail;
    // Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
    collaborationTags []string;
    // Lists people that also worked on the project.
    colleagues []RelatedPerson;
    // Contains detail about the user's role on the project.
    detail *PositionDetail;
    // Contains a friendly name for the project.
    displayName *string;
    // The Person or people who sponsored the project.
    sponsors []RelatedPerson;
    // 
    thumbnailUrl *string;
}
// NewProjectParticipation instantiates a new projectParticipation and sets the default values.
func NewProjectParticipation()(*ProjectParticipation) {
    m := &ProjectParticipation{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// GetCategories gets the categories property value. Contains categories a user has associated with the project (for example, digital transformation, oil rig).
func (m *ProjectParticipation) GetCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// GetClient gets the client property value. Contains detailed information about the client the project was for.
func (m *ProjectParticipation) GetClient()(*CompanyDetail) {
    if m == nil {
        return nil
    } else {
        return m.client
    }
}
// GetCollaborationTags gets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
func (m *ProjectParticipation) GetCollaborationTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.collaborationTags
    }
}
// GetColleagues gets the colleagues property value. Lists people that also worked on the project.
func (m *ProjectParticipation) GetColleagues()([]RelatedPerson) {
    if m == nil {
        return nil
    } else {
        return m.colleagues
    }
}
// GetDetail gets the detail property value. Contains detail about the user's role on the project.
func (m *ProjectParticipation) GetDetail()(*PositionDetail) {
    if m == nil {
        return nil
    } else {
        return m.detail
    }
}
// GetDisplayName gets the displayName property value. Contains a friendly name for the project.
func (m *ProjectParticipation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetSponsors gets the sponsors property value. The Person or people who sponsored the project.
func (m *ProjectParticipation) GetSponsors()([]RelatedPerson) {
    if m == nil {
        return nil
    } else {
        return m.sponsors
    }
}
// GetThumbnailUrl gets the thumbnailUrl property value. 
func (m *ProjectParticipation) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProjectParticipation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCategories(res)
        }
        return nil
    }
    res["client"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCompanyDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClient(val.(*CompanyDetail))
        }
        return nil
    }
    res["collaborationTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCollaborationTags(res)
        }
        return nil
    }
    res["colleagues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRelatedPerson() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RelatedPerson, len(val))
            for i, v := range val {
                res[i] = *(v.(*RelatedPerson))
            }
            m.SetColleagues(res)
        }
        return nil
    }
    res["detail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPositionDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(*PositionDetail))
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["sponsors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRelatedPerson() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RelatedPerson, len(val))
            for i, v := range val {
                res[i] = *(v.(*RelatedPerson))
            }
            m.SetSponsors(res)
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
    return res
}
func (m *ProjectParticipation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ProjectParticipation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategories() != nil {
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
    if m.GetCollaborationTags() != nil {
        err = writer.WriteCollectionOfStringValues("collaborationTags", m.GetCollaborationTags())
        if err != nil {
            return err
        }
    }
    if m.GetColleagues() != nil {
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
    if m.GetSponsors() != nil {
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
// SetCategories sets the categories property value. Contains categories a user has associated with the project (for example, digital transformation, oil rig).
func (m *ProjectParticipation) SetCategories(value []string)() {
    if m != nil {
        m.categories = value
    }
}
// SetClient sets the client property value. Contains detailed information about the client the project was for.
func (m *ProjectParticipation) SetClient(value *CompanyDetail)() {
    if m != nil {
        m.client = value
    }
}
// SetCollaborationTags sets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
func (m *ProjectParticipation) SetCollaborationTags(value []string)() {
    if m != nil {
        m.collaborationTags = value
    }
}
// SetColleagues sets the colleagues property value. Lists people that also worked on the project.
func (m *ProjectParticipation) SetColleagues(value []RelatedPerson)() {
    if m != nil {
        m.colleagues = value
    }
}
// SetDetail sets the detail property value. Contains detail about the user's role on the project.
func (m *ProjectParticipation) SetDetail(value *PositionDetail)() {
    if m != nil {
        m.detail = value
    }
}
// SetDisplayName sets the displayName property value. Contains a friendly name for the project.
func (m *ProjectParticipation) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetSponsors sets the sponsors property value. The Person or people who sponsored the project.
func (m *ProjectParticipation) SetSponsors(value []RelatedPerson)() {
    if m != nil {
        m.sponsors = value
    }
}
// SetThumbnailUrl sets the thumbnailUrl property value. 
func (m *ProjectParticipation) SetThumbnailUrl(value *string)() {
    if m != nil {
        m.thumbnailUrl = value
    }
}
