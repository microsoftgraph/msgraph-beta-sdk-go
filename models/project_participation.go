package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProjectParticipation 
type ProjectParticipation struct {
    ItemFacet
    // The OdataType property
    OdataType *string
}
// NewProjectParticipation instantiates a new projectParticipation and sets the default values.
func NewProjectParticipation()(*ProjectParticipation) {
    m := &ProjectParticipation{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.projectParticipation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateProjectParticipationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProjectParticipationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProjectParticipation(), nil
}
// GetCategories gets the categories property value. Contains categories a user has associated with the project (for example, digital transformation, oil rig).
func (m *ProjectParticipation) GetCategories()([]string) {
    val, err := m.GetBackingStore().Get("categories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetClient gets the client property value. Contains detailed information about the client the project was for.
func (m *ProjectParticipation) GetClient()(CompanyDetailable) {
    val, err := m.GetBackingStore().Get("client")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CompanyDetailable)
    }
    return nil
}
// GetCollaborationTags gets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
func (m *ProjectParticipation) GetCollaborationTags()([]string) {
    val, err := m.GetBackingStore().Get("collaborationTags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetColleagues gets the colleagues property value. Lists people that also worked on the project.
func (m *ProjectParticipation) GetColleagues()([]RelatedPersonable) {
    val, err := m.GetBackingStore().Get("colleagues")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]RelatedPersonable)
    }
    return nil
}
// GetDetail gets the detail property value. Contains detail about the user's role on the project.
func (m *ProjectParticipation) GetDetail()(PositionDetailable) {
    val, err := m.GetBackingStore().Get("detail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PositionDetailable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Contains a friendly name for the project.
func (m *ProjectParticipation) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
                if v != nil {
                    res[i] = v.(RelatedPersonable)
                }
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
                if v != nil {
                    res[i] = v.(RelatedPersonable)
                }
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
    val, err := m.GetBackingStore().Get("sponsors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]RelatedPersonable)
    }
    return nil
}
// GetThumbnailUrl gets the thumbnailUrl property value. The thumbnailUrl property
func (m *ProjectParticipation) GetThumbnailUrl()(*string) {
    val, err := m.GetBackingStore().Get("thumbnailUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    err := m.GetBackingStore().Set("categories", value)
    if err != nil {
        panic(err)
    }
}
// SetClient sets the client property value. Contains detailed information about the client the project was for.
func (m *ProjectParticipation) SetClient(value CompanyDetailable)() {
    err := m.GetBackingStore().Set("client", value)
    if err != nil {
        panic(err)
    }
}
// SetCollaborationTags sets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
func (m *ProjectParticipation) SetCollaborationTags(value []string)() {
    err := m.GetBackingStore().Set("collaborationTags", value)
    if err != nil {
        panic(err)
    }
}
// SetColleagues sets the colleagues property value. Lists people that also worked on the project.
func (m *ProjectParticipation) SetColleagues(value []RelatedPersonable)() {
    err := m.GetBackingStore().Set("colleagues", value)
    if err != nil {
        panic(err)
    }
}
// SetDetail sets the detail property value. Contains detail about the user's role on the project.
func (m *ProjectParticipation) SetDetail(value PositionDetailable)() {
    err := m.GetBackingStore().Set("detail", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Contains a friendly name for the project.
func (m *ProjectParticipation) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetSponsors sets the sponsors property value. The Person or people who sponsored the project.
func (m *ProjectParticipation) SetSponsors(value []RelatedPersonable)() {
    err := m.GetBackingStore().Set("sponsors", value)
    if err != nil {
        panic(err)
    }
}
// SetThumbnailUrl sets the thumbnailUrl property value. The thumbnailUrl property
func (m *ProjectParticipation) SetThumbnailUrl(value *string)() {
    err := m.GetBackingStore().Set("thumbnailUrl", value)
    if err != nil {
        panic(err)
    }
}
// ProjectParticipationable 
type ProjectParticipationable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategories()([]string)
    GetClient()(CompanyDetailable)
    GetCollaborationTags()([]string)
    GetColleagues()([]RelatedPersonable)
    GetDetail()(PositionDetailable)
    GetDisplayName()(*string)
    GetSponsors()([]RelatedPersonable)
    GetThumbnailUrl()(*string)
    SetCategories(value []string)()
    SetClient(value CompanyDetailable)()
    SetCollaborationTags(value []string)()
    SetColleagues(value []RelatedPersonable)()
    SetDetail(value PositionDetailable)()
    SetDisplayName(value *string)()
    SetSponsors(value []RelatedPersonable)()
    SetThumbnailUrl(value *string)()
}
