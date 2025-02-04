package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MailboxFolder struct {
    Entity
}
// NewMailboxFolder instantiates a new MailboxFolder and sets the default values.
func NewMailboxFolder()(*MailboxFolder) {
    m := &MailboxFolder{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMailboxFolderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMailboxFolderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMailboxFolder(), nil
}
// GetChildFolderCount gets the childFolderCount property value. The number of immediate child folders in the current folder.
// returns a *int32 when successful
func (m *MailboxFolder) GetChildFolderCount()(*int32) {
    val, err := m.GetBackingStore().Get("childFolderCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetChildFolders gets the childFolders property value. The collection of child folders in this folder.
// returns a []MailboxFolderable when successful
func (m *MailboxFolder) GetChildFolders()([]MailboxFolderable) {
    val, err := m.GetBackingStore().Get("childFolders")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MailboxFolderable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name of the folder.
// returns a *string when successful
func (m *MailboxFolder) GetDisplayName()(*string) {
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
func (m *MailboxFolder) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["childFolderCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChildFolderCount(val)
        }
        return nil
    }
    res["childFolders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMailboxFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MailboxFolderable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MailboxFolderable)
                }
            }
            m.SetChildFolders(res)
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
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMailboxItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MailboxItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MailboxItemable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    res["multiValueExtendedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MultiValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MultiValueLegacyExtendedPropertyable)
                }
            }
            m.SetMultiValueExtendedProperties(res)
        }
        return nil
    }
    res["parentFolderId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentFolderId(val)
        }
        return nil
    }
    res["parentMailboxUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentMailboxUrl(val)
        }
        return nil
    }
    res["singleValueExtendedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SingleValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SingleValueLegacyExtendedPropertyable)
                }
            }
            m.SetSingleValueExtendedProperties(res)
        }
        return nil
    }
    res["totalItemCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalItemCount(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. The collection of items in this folder.
// returns a []MailboxItemable when successful
func (m *MailboxFolder) GetItems()([]MailboxItemable) {
    val, err := m.GetBackingStore().Get("items")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MailboxItemable)
    }
    return nil
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the mailboxFolder.
// returns a []MultiValueLegacyExtendedPropertyable when successful
func (m *MailboxFolder) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    val, err := m.GetBackingStore().Get("multiValueExtendedProperties")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MultiValueLegacyExtendedPropertyable)
    }
    return nil
}
// GetParentFolderId gets the parentFolderId property value. The unique identifier for the parent folder of this folder.
// returns a *string when successful
func (m *MailboxFolder) GetParentFolderId()(*string) {
    val, err := m.GetBackingStore().Get("parentFolderId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetParentMailboxUrl gets the parentMailboxUrl property value. The routing link to the actual underlying mailbox where the folder physically resides. The folder can be accessed using GET {parentMailboxUrl}/folders/{id}, which treats the entire URL as an opaque string.  This method is especially important when auto-expanding archiving is enabled for a user's in-place archive mailbox. The user's archive content can span across multiple mailboxes in such scenarios.
// returns a *string when successful
func (m *MailboxFolder) GetParentMailboxUrl()(*string) {
    val, err := m.GetBackingStore().Get("parentMailboxUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the mailboxFolder.
// returns a []SingleValueLegacyExtendedPropertyable when successful
func (m *MailboxFolder) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    val, err := m.GetBackingStore().Get("singleValueExtendedProperties")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SingleValueLegacyExtendedPropertyable)
    }
    return nil
}
// GetTotalItemCount gets the totalItemCount property value. The number of items in the folder.
// returns a *int32 when successful
func (m *MailboxFolder) GetTotalItemCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalItemCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTypeEscaped gets the type property value. Describes the folder class type.
// returns a *string when successful
func (m *MailboxFolder) GetTypeEscaped()(*string) {
    val, err := m.GetBackingStore().Get("typeEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MailboxFolder) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("childFolderCount", m.GetChildFolderCount())
        if err != nil {
            return err
        }
    }
    if m.GetChildFolders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChildFolders()))
        for i, v := range m.GetChildFolders() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("childFolders", cast)
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
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMultiValueExtendedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentFolderId", m.GetParentFolderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentMailboxUrl", m.GetParentMailboxUrl())
        if err != nil {
            return err
        }
    }
    if m.GetSingleValueExtendedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalItemCount", m.GetTotalItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChildFolderCount sets the childFolderCount property value. The number of immediate child folders in the current folder.
func (m *MailboxFolder) SetChildFolderCount(value *int32)() {
    err := m.GetBackingStore().Set("childFolderCount", value)
    if err != nil {
        panic(err)
    }
}
// SetChildFolders sets the childFolders property value. The collection of child folders in this folder.
func (m *MailboxFolder) SetChildFolders(value []MailboxFolderable)() {
    err := m.GetBackingStore().Set("childFolders", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name of the folder.
func (m *MailboxFolder) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetItems sets the items property value. The collection of items in this folder.
func (m *MailboxFolder) SetItems(value []MailboxItemable)() {
    err := m.GetBackingStore().Set("items", value)
    if err != nil {
        panic(err)
    }
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the mailboxFolder.
func (m *MailboxFolder) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    err := m.GetBackingStore().Set("multiValueExtendedProperties", value)
    if err != nil {
        panic(err)
    }
}
// SetParentFolderId sets the parentFolderId property value. The unique identifier for the parent folder of this folder.
func (m *MailboxFolder) SetParentFolderId(value *string)() {
    err := m.GetBackingStore().Set("parentFolderId", value)
    if err != nil {
        panic(err)
    }
}
// SetParentMailboxUrl sets the parentMailboxUrl property value. The routing link to the actual underlying mailbox where the folder physically resides. The folder can be accessed using GET {parentMailboxUrl}/folders/{id}, which treats the entire URL as an opaque string.  This method is especially important when auto-expanding archiving is enabled for a user's in-place archive mailbox. The user's archive content can span across multiple mailboxes in such scenarios.
func (m *MailboxFolder) SetParentMailboxUrl(value *string)() {
    err := m.GetBackingStore().Set("parentMailboxUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the mailboxFolder.
func (m *MailboxFolder) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    err := m.GetBackingStore().Set("singleValueExtendedProperties", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalItemCount sets the totalItemCount property value. The number of items in the folder.
func (m *MailboxFolder) SetTotalItemCount(value *int32)() {
    err := m.GetBackingStore().Set("totalItemCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTypeEscaped sets the type property value. Describes the folder class type.
func (m *MailboxFolder) SetTypeEscaped(value *string)() {
    err := m.GetBackingStore().Set("typeEscaped", value)
    if err != nil {
        panic(err)
    }
}
type MailboxFolderable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChildFolderCount()(*int32)
    GetChildFolders()([]MailboxFolderable)
    GetDisplayName()(*string)
    GetItems()([]MailboxItemable)
    GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable)
    GetParentFolderId()(*string)
    GetParentMailboxUrl()(*string)
    GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable)
    GetTotalItemCount()(*int32)
    GetTypeEscaped()(*string)
    SetChildFolderCount(value *int32)()
    SetChildFolders(value []MailboxFolderable)()
    SetDisplayName(value *string)()
    SetItems(value []MailboxItemable)()
    SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)()
    SetParentFolderId(value *string)()
    SetParentMailboxUrl(value *string)()
    SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)()
    SetTotalItemCount(value *int32)()
    SetTypeEscaped(value *string)()
}
