package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProjectParticipation provides operations to manage the compliance singleton.
type ProjectParticipation struct {
    ItemFacet
    // Contains categories a user has associated with the project (for example, digital transformation, oil rig).
    categories []string
    // Contains detailed information about the client the project was for.
    client CompanyDetailable
    // Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
    collaborationTags []string
    // Lists people that also worked on the project.
    colleagues []RelatedPersonable
    // Contains detail about the user's role on the project.
    detail PositionDetailable
    // Contains a friendly name for the project.
    displayName *string
    // The Person or people who sponsored the project.
    sponsors []RelatedPersonable
    // The thumbnailUrl property
    thumbnailUrl *string
}
// NewProjectParticipation instantiates a new projectParticipation and sets the default values.
func NewProjectParticipation()(*ProjectParticipation) {
    m := &ProjectParticipation{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// CreateProjectParticipationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProjectParticipationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProjectParticipation(), nil
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
func (m *ProjectParticipation) GetClient()(CompanyDetailable) {
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
func (m *ProjectParticipation) GetColleagues()([]RelatedPersonable) {
    if m == nil {
        return nil
    } else {
        return m.colleagues
    }
}
// GetDetail gets the detail property value. Contains detail about the user's role on the project.
func (m *ProjectParticipation) GetDetail()(PositionDetailable) {
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
// GetFieldDeserializers the deserialization information for the current model
func (m *ProjectParticipation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["client"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCompanyDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClient(val.(CompanyDetailable))
        }
        return nil
    }
    res["collaborationTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["colleagues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRelatedPersonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RelatedPersonable, len(val))
            for i, v := range val {
                res[i] = v.(RelatedPersonable)
            }
            m.SetColleagues(res)
        }
        return nil
    }
    res["detail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePositionDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(PositionDetailable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["sponsors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRelatedPersonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RelatedPersonable, len(val))
            for i, v := range val {
                res[i] = v.(RelatedPersonable)
            }
            m.SetSponsors(res)
        }
        return nil
    }
    res["thumbnailUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetSponsors gets the sponsors property value. The Person or people who sponsored the project.
func (m *ProjectParticipation) GetSponsors()([]RelatedPersonable) {
    if m == nil {
        return nil
    } else {
        return m.sponsors
    }
}
// GetThumbnailUrl gets the thumbnailUrl property value. The thumbnailUrl property
func (m *ProjectParticipation) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
// Serialize serializes information the current object
func (m *ProjectParticipation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetColleagues()))
        for i, v := range m.GetColleagues() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSponsors()))
        for i, v := range m.GetSponsors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
func (m *ProjectParticipation) SetClient(value CompanyDetailable)() {
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
func (m *ProjectParticipation) SetColleagues(value []RelatedPersonable)() {
    if m != nil {
        m.colleagues = value
    }
}
// SetDetail sets the detail property value. Contains detail about the user's role on the project.
func (m *ProjectParticipation) SetDetail(value PositionDetailable)() {
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
func (m *ProjectParticipation) SetSponsors(value []RelatedPersonable)() {
    if m != nil {
        m.sponsors = value
    }
}
// SetThumbnailUrl sets the thumbnailUrl property value. The thumbnailUrl property
func (m *ProjectParticipation) SetThumbnailUrl(value *string)() {
    if m != nil {
        m.thumbnailUrl = value
    }
}
