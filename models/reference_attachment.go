package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReferenceAttachment 
type ReferenceAttachment struct {
    Attachment
}
// NewReferenceAttachment instantiates a new ReferenceAttachment and sets the default values.
func NewReferenceAttachment()(*ReferenceAttachment) {
    m := &ReferenceAttachment{
        Attachment: *NewAttachment(),
    }
    odataTypeValue := "#microsoft.graph.referenceAttachment"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateReferenceAttachmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReferenceAttachmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReferenceAttachment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReferenceAttachment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Attachment.GetFieldDeserializers()
    res["isFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFolder(val)
        }
        return nil
    }
    res["permission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseReferenceAttachmentPermission)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermission(val.(*ReferenceAttachmentPermission))
        }
        return nil
    }
    res["previewUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewUrl(val)
        }
        return nil
    }
    res["providerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseReferenceAttachmentProvider)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderType(val.(*ReferenceAttachmentProvider))
        }
        return nil
    }
    res["sourceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceUrl(val)
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
// GetIsFolder gets the isFolder property value. Specifies whether the attachment is a link to a folder. Must set this to true if sourceUrl is a link to a folder. Optional.
func (m *ReferenceAttachment) GetIsFolder()(*bool) {
    val, err := m.GetBackingStore().Get("isFolder")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPermission gets the permission property value. Specifies the permissions granted for the attachment by the type of provider in providerType. Possible values are: other, view, edit, anonymousView, anonymousEdit, organizationView, organizationEdit. Optional.
func (m *ReferenceAttachment) GetPermission()(*ReferenceAttachmentPermission) {
    val, err := m.GetBackingStore().Get("permission")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ReferenceAttachmentPermission)
    }
    return nil
}
// GetPreviewUrl gets the previewUrl property value. Applies to only a reference attachment of an image - URL to get a preview image. Use thumbnailUrl and previewUrl only when sourceUrl identifies an image file. Optional.
func (m *ReferenceAttachment) GetPreviewUrl()(*string) {
    val, err := m.GetBackingStore().Get("previewUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProviderType gets the providerType property value. The type of provider that supports an attachment of this contentType. Possible values are: other, oneDriveBusiness, oneDriveConsumer, dropbox. Optional.
func (m *ReferenceAttachment) GetProviderType()(*ReferenceAttachmentProvider) {
    val, err := m.GetBackingStore().Get("providerType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ReferenceAttachmentProvider)
    }
    return nil
}
// GetSourceUrl gets the sourceUrl property value. URL to get the attachment content. If this is a URL to a folder, then for the folder to be displayed correctly in Outlook or Outlook on the web, set isFolder to true. Required.
func (m *ReferenceAttachment) GetSourceUrl()(*string) {
    val, err := m.GetBackingStore().Get("sourceUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetThumbnailUrl gets the thumbnailUrl property value. Applies to only a reference attachment of an image - URL to get a thumbnail image. Use thumbnailUrl and previewUrl only when sourceUrl identifies an image file. Optional.
func (m *ReferenceAttachment) GetThumbnailUrl()(*string) {
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
func (m *ReferenceAttachment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Attachment.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isFolder", m.GetIsFolder())
        if err != nil {
            return err
        }
    }
    if m.GetPermission() != nil {
        cast := (*m.GetPermission()).String()
        err = writer.WriteStringValue("permission", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("previewUrl", m.GetPreviewUrl())
        if err != nil {
            return err
        }
    }
    if m.GetProviderType() != nil {
        cast := (*m.GetProviderType()).String()
        err = writer.WriteStringValue("providerType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceUrl", m.GetSourceUrl())
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
// SetIsFolder sets the isFolder property value. Specifies whether the attachment is a link to a folder. Must set this to true if sourceUrl is a link to a folder. Optional.
func (m *ReferenceAttachment) SetIsFolder(value *bool)() {
    err := m.GetBackingStore().Set("isFolder", value)
    if err != nil {
        panic(err)
    }
}
// SetPermission sets the permission property value. Specifies the permissions granted for the attachment by the type of provider in providerType. Possible values are: other, view, edit, anonymousView, anonymousEdit, organizationView, organizationEdit. Optional.
func (m *ReferenceAttachment) SetPermission(value *ReferenceAttachmentPermission)() {
    err := m.GetBackingStore().Set("permission", value)
    if err != nil {
        panic(err)
    }
}
// SetPreviewUrl sets the previewUrl property value. Applies to only a reference attachment of an image - URL to get a preview image. Use thumbnailUrl and previewUrl only when sourceUrl identifies an image file. Optional.
func (m *ReferenceAttachment) SetPreviewUrl(value *string)() {
    err := m.GetBackingStore().Set("previewUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetProviderType sets the providerType property value. The type of provider that supports an attachment of this contentType. Possible values are: other, oneDriveBusiness, oneDriveConsumer, dropbox. Optional.
func (m *ReferenceAttachment) SetProviderType(value *ReferenceAttachmentProvider)() {
    err := m.GetBackingStore().Set("providerType", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceUrl sets the sourceUrl property value. URL to get the attachment content. If this is a URL to a folder, then for the folder to be displayed correctly in Outlook or Outlook on the web, set isFolder to true. Required.
func (m *ReferenceAttachment) SetSourceUrl(value *string)() {
    err := m.GetBackingStore().Set("sourceUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetThumbnailUrl sets the thumbnailUrl property value. Applies to only a reference attachment of an image - URL to get a thumbnail image. Use thumbnailUrl and previewUrl only when sourceUrl identifies an image file. Optional.
func (m *ReferenceAttachment) SetThumbnailUrl(value *string)() {
    err := m.GetBackingStore().Set("thumbnailUrl", value)
    if err != nil {
        panic(err)
    }
}
// ReferenceAttachmentable 
type ReferenceAttachmentable interface {
    Attachmentable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsFolder()(*bool)
    GetPermission()(*ReferenceAttachmentPermission)
    GetPreviewUrl()(*string)
    GetProviderType()(*ReferenceAttachmentProvider)
    GetSourceUrl()(*string)
    GetThumbnailUrl()(*string)
    SetIsFolder(value *bool)()
    SetPermission(value *ReferenceAttachmentPermission)()
    SetPreviewUrl(value *string)()
    SetProviderType(value *ReferenceAttachmentProvider)()
    SetSourceUrl(value *string)()
    SetThumbnailUrl(value *string)()
}
