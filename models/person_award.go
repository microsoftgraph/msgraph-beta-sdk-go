package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonAward 
type PersonAward struct {
    ItemFacet
    // Descpription of the award or honor.
    description *string
    // Name of the award or honor.
    displayName *string
    // The date that the award or honor was granted.
    issuedDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // Authority which granted the award or honor.
    issuingAuthority *string
    // URL referencing a thumbnail of the award or honor.
    thumbnailUrl *string
    // URL referencing the award or honor.
    webUrl *string
}
// NewPersonAward instantiates a new PersonAward and sets the default values.
func NewPersonAward()(*PersonAward) {
    m := &PersonAward{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.personAward";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePersonAwardFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonAwardFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonAward(), nil
}
// GetDescription gets the description property value. Descpription of the award or honor.
func (m *PersonAward) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Name of the award or honor.
func (m *PersonAward) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonAward) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["issuedDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetIssuedDate)
    res["issuingAuthority"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetIssuingAuthority)
    res["thumbnailUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetThumbnailUrl)
    res["webUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWebUrl)
    return res
}
// GetIssuedDate gets the issuedDate property value. The date that the award or honor was granted.
func (m *PersonAward) GetIssuedDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.issuedDate
}
// GetIssuingAuthority gets the issuingAuthority property value. Authority which granted the award or honor.
func (m *PersonAward) GetIssuingAuthority()(*string) {
    return m.issuingAuthority
}
// GetThumbnailUrl gets the thumbnailUrl property value. URL referencing a thumbnail of the award or honor.
func (m *PersonAward) GetThumbnailUrl()(*string) {
    return m.thumbnailUrl
}
// GetWebUrl gets the webUrl property value. URL referencing the award or honor.
func (m *PersonAward) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *PersonAward) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("issuedDate", m.GetIssuedDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuingAuthority", m.GetIssuingAuthority())
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
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Descpription of the award or honor.
func (m *PersonAward) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Name of the award or honor.
func (m *PersonAward) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIssuedDate sets the issuedDate property value. The date that the award or honor was granted.
func (m *PersonAward) SetIssuedDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.issuedDate = value
}
// SetIssuingAuthority sets the issuingAuthority property value. Authority which granted the award or honor.
func (m *PersonAward) SetIssuingAuthority(value *string)() {
    m.issuingAuthority = value
}
// SetThumbnailUrl sets the thumbnailUrl property value. URL referencing a thumbnail of the award or honor.
func (m *PersonAward) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// SetWebUrl sets the webUrl property value. URL referencing the award or honor.
func (m *PersonAward) SetWebUrl(value *string)() {
    m.webUrl = value
}
