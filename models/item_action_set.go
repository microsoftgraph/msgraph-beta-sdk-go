package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActionSet 
type ItemActionSet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A comment was added to the item.
    comment CommentActionable
    // An item was created.
    create CreateActionable
    // An item was deleted.
    delete DeleteActionable
    // An item was edited.
    edit EditActionable
    // A user was mentioned in the item.
    mention MentionActionable
    // An item was moved.
    move MoveActionable
    // An item was renamed.
    rename RenameActionable
    // An item was restored.
    restore RestoreActionable
    // An item was shared.
    share ShareActionable
    // An item was versioned.
    version VersionActionable
}
// NewItemActionSet instantiates a new itemActionSet and sets the default values.
func NewItemActionSet()(*ItemActionSet) {
    m := &ItemActionSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemActionSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionSet(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionSet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetComment gets the comment property value. A comment was added to the item.
func (m *ItemActionSet) GetComment()(CommentActionable) {
    if m == nil {
        return nil
    } else {
        return m.comment
    }
}
// GetCreate gets the create property value. An item was created.
func (m *ItemActionSet) GetCreate()(CreateActionable) {
    if m == nil {
        return nil
    } else {
        return m.create
    }
}
// GetDelete gets the delete property value. An item was deleted.
func (m *ItemActionSet) GetDelete()(DeleteActionable) {
    if m == nil {
        return nil
    } else {
        return m.delete
    }
}
// GetEdit gets the edit property value. An item was edited.
func (m *ItemActionSet) GetEdit()(EditActionable) {
    if m == nil {
        return nil
    } else {
        return m.edit
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActionSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCommentActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val.(CommentActionable))
        }
        return nil
    }
    res["create"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCreateActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreate(val.(CreateActionable))
        }
        return nil
    }
    res["delete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeleteActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDelete(val.(DeleteActionable))
        }
        return nil
    }
    res["edit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEditActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdit(val.(EditActionable))
        }
        return nil
    }
    res["mention"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMentionActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMention(val.(MentionActionable))
        }
        return nil
    }
    res["move"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMoveActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMove(val.(MoveActionable))
        }
        return nil
    }
    res["rename"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRenameActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRename(val.(RenameActionable))
        }
        return nil
    }
    res["restore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRestoreActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestore(val.(RestoreActionable))
        }
        return nil
    }
    res["share"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateShareActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShare(val.(ShareActionable))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVersionActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val.(VersionActionable))
        }
        return nil
    }
    return res
}
// GetMention gets the mention property value. A user was mentioned in the item.
func (m *ItemActionSet) GetMention()(MentionActionable) {
    if m == nil {
        return nil
    } else {
        return m.mention
    }
}
// GetMove gets the move property value. An item was moved.
func (m *ItemActionSet) GetMove()(MoveActionable) {
    if m == nil {
        return nil
    } else {
        return m.move
    }
}
// GetRename gets the rename property value. An item was renamed.
func (m *ItemActionSet) GetRename()(RenameActionable) {
    if m == nil {
        return nil
    } else {
        return m.rename
    }
}
// GetRestore gets the restore property value. An item was restored.
func (m *ItemActionSet) GetRestore()(RestoreActionable) {
    if m == nil {
        return nil
    } else {
        return m.restore
    }
}
// GetShare gets the share property value. An item was shared.
func (m *ItemActionSet) GetShare()(ShareActionable) {
    if m == nil {
        return nil
    } else {
        return m.share
    }
}
// GetVersion gets the version property value. An item was versioned.
func (m *ItemActionSet) GetVersion()(VersionActionable) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// Serialize serializes information the current object
func (m *ItemActionSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("create", m.GetCreate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("delete", m.GetDelete())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("edit", m.GetEdit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("mention", m.GetMention())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("move", m.GetMove())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("rename", m.GetRename())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("restore", m.GetRestore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("share", m.GetShare())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionSet) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetComment sets the comment property value. A comment was added to the item.
func (m *ItemActionSet) SetComment(value CommentActionable)() {
    if m != nil {
        m.comment = value
    }
}
// SetCreate sets the create property value. An item was created.
func (m *ItemActionSet) SetCreate(value CreateActionable)() {
    if m != nil {
        m.create = value
    }
}
// SetDelete sets the delete property value. An item was deleted.
func (m *ItemActionSet) SetDelete(value DeleteActionable)() {
    if m != nil {
        m.delete = value
    }
}
// SetEdit sets the edit property value. An item was edited.
func (m *ItemActionSet) SetEdit(value EditActionable)() {
    if m != nil {
        m.edit = value
    }
}
// SetMention sets the mention property value. A user was mentioned in the item.
func (m *ItemActionSet) SetMention(value MentionActionable)() {
    if m != nil {
        m.mention = value
    }
}
// SetMove sets the move property value. An item was moved.
func (m *ItemActionSet) SetMove(value MoveActionable)() {
    if m != nil {
        m.move = value
    }
}
// SetRename sets the rename property value. An item was renamed.
func (m *ItemActionSet) SetRename(value RenameActionable)() {
    if m != nil {
        m.rename = value
    }
}
// SetRestore sets the restore property value. An item was restored.
func (m *ItemActionSet) SetRestore(value RestoreActionable)() {
    if m != nil {
        m.restore = value
    }
}
// SetShare sets the share property value. An item was shared.
func (m *ItemActionSet) SetShare(value ShareActionable)() {
    if m != nil {
        m.share = value
    }
}
// SetVersion sets the version property value. An item was versioned.
func (m *ItemActionSet) SetVersion(value VersionActionable)() {
    if m != nil {
        m.version = value
    }
}
