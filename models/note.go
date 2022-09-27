package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Note 
type Note struct {
    OutlookItem
    // The attachments property
    attachments []Attachmentable
    // The body property
    body ItemBodyable
    // The extensions property
    extensions []Extensionable
    // The hasAttachments property
    hasAttachments *bool
    // The isDeleted property
    isDeleted *bool
    // The multiValueExtendedProperties property
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable
    // The singleValueExtendedProperties property
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable
    // The subject property
    subject *string
}
// NewNote instantiates a new Note and sets the default values.
func NewNote()(*Note) {
    m := &Note{
        OutlookItem: *NewOutlookItem(),
    }
    odataTypeValue := "#microsoft.graph.note";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateNoteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNoteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNote(), nil
}
// GetAttachments gets the attachments property value. The attachments property
func (m *Note) GetAttachments()([]Attachmentable) {
    return m.attachments
}
// GetBody gets the body property value. The body property
func (m *Note) GetBody()(ItemBodyable) {
    return m.body
}
// GetExtensions gets the extensions property value. The extensions property
func (m *Note) GetExtensions()([]Extensionable) {
    return m.extensions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Note) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["attachments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAttachmentFromDiscriminatorValue , m.SetAttachments)
    res["body"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateItemBodyFromDiscriminatorValue , m.SetBody)
    res["extensions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue , m.SetExtensions)
    res["hasAttachments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetHasAttachments)
    res["isDeleted"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsDeleted)
    res["multiValueExtendedProperties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue , m.SetMultiValueExtendedProperties)
    res["singleValueExtendedProperties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue , m.SetSingleValueExtendedProperties)
    res["subject"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubject)
    return res
}
// GetHasAttachments gets the hasAttachments property value. The hasAttachments property
func (m *Note) GetHasAttachments()(*bool) {
    return m.hasAttachments
}
// GetIsDeleted gets the isDeleted property value. The isDeleted property
func (m *Note) GetIsDeleted()(*bool) {
    return m.isDeleted
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The multiValueExtendedProperties property
func (m *Note) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    return m.multiValueExtendedProperties
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The singleValueExtendedProperties property
func (m *Note) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    return m.singleValueExtendedProperties
}
// GetSubject gets the subject property value. The subject property
func (m *Note) GetSubject()(*string) {
    return m.subject
}
// Serialize serializes information the current object
func (m *Note) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OutlookItem.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAttachments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAttachments())
        err = writer.WriteCollectionOfObjectValues("attachments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExtensions())
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasAttachments", m.GetHasAttachments())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    if m.GetMultiValueExtendedProperties() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMultiValueExtendedProperties())
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSingleValueExtendedProperties() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSingleValueExtendedProperties())
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttachments sets the attachments property value. The attachments property
func (m *Note) SetAttachments(value []Attachmentable)() {
    m.attachments = value
}
// SetBody sets the body property value. The body property
func (m *Note) SetBody(value ItemBodyable)() {
    m.body = value
}
// SetExtensions sets the extensions property value. The extensions property
func (m *Note) SetExtensions(value []Extensionable)() {
    m.extensions = value
}
// SetHasAttachments sets the hasAttachments property value. The hasAttachments property
func (m *Note) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
// SetIsDeleted sets the isDeleted property value. The isDeleted property
func (m *Note) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The multiValueExtendedProperties property
func (m *Note) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    m.multiValueExtendedProperties = value
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The singleValueExtendedProperties property
func (m *Note) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    m.singleValueExtendedProperties = value
}
// SetSubject sets the subject property value. The subject property
func (m *Note) SetSubject(value *string)() {
    m.subject = value
}
