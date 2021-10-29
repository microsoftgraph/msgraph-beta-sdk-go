package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new projectParticipation and sets the default values.
func NewProjectParticipation()(*ProjectParticipation) {
    m := &ProjectParticipation{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// Gets the categories property value. Contains categories a user has associated with the project (for example, digital transformation, oil rig).
func (m *ProjectParticipation) GetCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// Gets the client property value. Contains detailed information about the client the project was for.
func (m *ProjectParticipation) GetClient()(*CompanyDetail) {
    if m == nil {
        return nil
    } else {
        return m.client
    }
}
// Gets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
func (m *ProjectParticipation) GetCollaborationTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.collaborationTags
    }
}
// Gets the colleagues property value. Lists people that also worked on the project.
func (m *ProjectParticipation) GetColleagues()([]RelatedPerson) {
    if m == nil {
        return nil
    } else {
        return m.colleagues
    }
}
// Gets the detail property value. Contains detail about the user's role on the project.
func (m *ProjectParticipation) GetDetail()(*PositionDetail) {
    if m == nil {
        return nil
    } else {
        return m.detail
    }
}
// Gets the displayName property value. Contains a friendly name for the project.
func (m *ProjectParticipation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the sponsors property value. The Person or people who sponsored the project.
func (m *ProjectParticipation) GetSponsors()([]RelatedPerson) {
    if m == nil {
        return nil
    } else {
        return m.sponsors
    }
}
// Gets the thumbnailUrl property value. 
func (m *ProjectParticipation) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the categories property value. Contains categories a user has associated with the project (for example, digital transformation, oil rig).
// Parameters:
//  - value : Value to set for the categories property.
func (m *ProjectParticipation) SetCategories(value []string)() {
    m.categories = value
}
// Sets the client property value. Contains detailed information about the client the project was for.
// Parameters:
//  - value : Value to set for the client property.
func (m *ProjectParticipation) SetClient(value *CompanyDetail)() {
    m.client = value
}
// Sets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
// Parameters:
//  - value : Value to set for the collaborationTags property.
func (m *ProjectParticipation) SetCollaborationTags(value []string)() {
    m.collaborationTags = value
}
// Sets the colleagues property value. Lists people that also worked on the project.
// Parameters:
//  - value : Value to set for the colleagues property.
func (m *ProjectParticipation) SetColleagues(value []RelatedPerson)() {
    m.colleagues = value
}
// Sets the detail property value. Contains detail about the user's role on the project.
// Parameters:
//  - value : Value to set for the detail property.
func (m *ProjectParticipation) SetDetail(value *PositionDetail)() {
    m.detail = value
}
// Sets the displayName property value. Contains a friendly name for the project.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ProjectParticipation) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the sponsors property value. The Person or people who sponsored the project.
// Parameters:
//  - value : Value to set for the sponsors property.
func (m *ProjectParticipation) SetSponsors(value []RelatedPerson)() {
    m.sponsors = value
}
// Sets the thumbnailUrl property value. 
// Parameters:
//  - value : Value to set for the thumbnailUrl property.
func (m *ProjectParticipation) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
