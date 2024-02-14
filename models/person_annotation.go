package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PersonAnnotation struct {
    ItemFacet
}
// NewPersonAnnotation instantiates a new PersonAnnotation and sets the default values.
func NewPersonAnnotation()(*PersonAnnotation) {
    m := &PersonAnnotation{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.personAnnotation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePersonAnnotationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePersonAnnotationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonAnnotation(), nil
}
// GetDetail gets the detail property value. Contains the detail of the note itself.
// returns a ItemBodyable when successful
func (m *PersonAnnotation) GetDetail()(ItemBodyable) {
    val, err := m.GetBackingStore().Get("detail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ItemBodyable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Contains a friendly name for the note.
// returns a *string when successful
func (m *PersonAnnotation) GetDisplayName()(*string) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PersonAnnotation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["detail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(ItemBodyable))
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
// GetThumbnailUrl gets the thumbnailUrl property value. The thumbnailUrl property
// returns a *string when successful
func (m *PersonAnnotation) GetThumbnailUrl()(*string) {
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
func (m *PersonAnnotation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteStringValue("thumbnailUrl", m.GetThumbnailUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDetail sets the detail property value. Contains the detail of the note itself.
func (m *PersonAnnotation) SetDetail(value ItemBodyable)() {
    err := m.GetBackingStore().Set("detail", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Contains a friendly name for the note.
func (m *PersonAnnotation) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetThumbnailUrl sets the thumbnailUrl property value. The thumbnailUrl property
func (m *PersonAnnotation) SetThumbnailUrl(value *string)() {
    err := m.GetBackingStore().Set("thumbnailUrl", value)
    if err != nil {
        panic(err)
    }
}
type PersonAnnotationable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetail()(ItemBodyable)
    GetDisplayName()(*string)
    GetThumbnailUrl()(*string)
    SetDetail(value ItemBodyable)()
    SetDisplayName(value *string)()
    SetThumbnailUrl(value *string)()
}
