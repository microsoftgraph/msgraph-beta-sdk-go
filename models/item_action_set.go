package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ItemActionSet struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemActionSet instantiates a new ItemActionSet and sets the default values.
func NewItemActionSet()(*ItemActionSet) {
    m := &ItemActionSet{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemActionSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionSet(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemActionSet) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *ItemActionSet) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetComment gets the comment property value. A comment was added to the item.
// returns a CommentActionable when successful
func (m *ItemActionSet) GetComment()(CommentActionable) {
    val, err := m.GetBackingStore().Get("comment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CommentActionable)
    }
    return nil
}
// GetCreate gets the create property value. An item was created.
// returns a CreateActionable when successful
func (m *ItemActionSet) GetCreate()(CreateActionable) {
    val, err := m.GetBackingStore().Get("create")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CreateActionable)
    }
    return nil
}
// GetDelete gets the delete property value. An item was deleted.
// returns a DeleteActionable when successful
func (m *ItemActionSet) GetDelete()(DeleteActionable) {
    val, err := m.GetBackingStore().Get("delete")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeleteActionable)
    }
    return nil
}
// GetEdit gets the edit property value. An item was edited.
// returns a EditActionable when successful
func (m *ItemActionSet) GetEdit()(EditActionable) {
    val, err := m.GetBackingStore().Get("edit")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EditActionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
// returns a MentionActionable when successful
func (m *ItemActionSet) GetMention()(MentionActionable) {
    val, err := m.GetBackingStore().Get("mention")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MentionActionable)
    }
    return nil
}
// GetMove gets the move property value. An item was moved.
// returns a MoveActionable when successful
func (m *ItemActionSet) GetMove()(MoveActionable) {
    val, err := m.GetBackingStore().Get("move")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MoveActionable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ItemActionSet) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRename gets the rename property value. An item was renamed.
// returns a RenameActionable when successful
func (m *ItemActionSet) GetRename()(RenameActionable) {
    val, err := m.GetBackingStore().Get("rename")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RenameActionable)
    }
    return nil
}
// GetRestore gets the restore property value. An item was restored.
// returns a RestoreActionable when successful
func (m *ItemActionSet) GetRestore()(RestoreActionable) {
    val, err := m.GetBackingStore().Get("restore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RestoreActionable)
    }
    return nil
}
// GetShare gets the share property value. An item was shared.
// returns a ShareActionable when successful
func (m *ItemActionSet) GetShare()(ShareActionable) {
    val, err := m.GetBackingStore().Get("share")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ShareActionable)
    }
    return nil
}
// GetVersion gets the version property value. An item was versioned.
// returns a VersionActionable when successful
func (m *ItemActionSet) GetVersion()(VersionActionable) {
    val, err := m.GetBackingStore().Get("version")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(VersionActionable)
    }
    return nil
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionSet) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ItemActionSet) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetComment sets the comment property value. A comment was added to the item.
func (m *ItemActionSet) SetComment(value CommentActionable)() {
    err := m.GetBackingStore().Set("comment", value)
    if err != nil {
        panic(err)
    }
}
// SetCreate sets the create property value. An item was created.
func (m *ItemActionSet) SetCreate(value CreateActionable)() {
    err := m.GetBackingStore().Set("create", value)
    if err != nil {
        panic(err)
    }
}
// SetDelete sets the delete property value. An item was deleted.
func (m *ItemActionSet) SetDelete(value DeleteActionable)() {
    err := m.GetBackingStore().Set("delete", value)
    if err != nil {
        panic(err)
    }
}
// SetEdit sets the edit property value. An item was edited.
func (m *ItemActionSet) SetEdit(value EditActionable)() {
    err := m.GetBackingStore().Set("edit", value)
    if err != nil {
        panic(err)
    }
}
// SetMention sets the mention property value. A user was mentioned in the item.
func (m *ItemActionSet) SetMention(value MentionActionable)() {
    err := m.GetBackingStore().Set("mention", value)
    if err != nil {
        panic(err)
    }
}
// SetMove sets the move property value. An item was moved.
func (m *ItemActionSet) SetMove(value MoveActionable)() {
    err := m.GetBackingStore().Set("move", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ItemActionSet) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRename sets the rename property value. An item was renamed.
func (m *ItemActionSet) SetRename(value RenameActionable)() {
    err := m.GetBackingStore().Set("rename", value)
    if err != nil {
        panic(err)
    }
}
// SetRestore sets the restore property value. An item was restored.
func (m *ItemActionSet) SetRestore(value RestoreActionable)() {
    err := m.GetBackingStore().Set("restore", value)
    if err != nil {
        panic(err)
    }
}
// SetShare sets the share property value. An item was shared.
func (m *ItemActionSet) SetShare(value ShareActionable)() {
    err := m.GetBackingStore().Set("share", value)
    if err != nil {
        panic(err)
    }
}
// SetVersion sets the version property value. An item was versioned.
func (m *ItemActionSet) SetVersion(value VersionActionable)() {
    err := m.GetBackingStore().Set("version", value)
    if err != nil {
        panic(err)
    }
}
type ItemActionSetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetComment()(CommentActionable)
    GetCreate()(CreateActionable)
    GetDelete()(DeleteActionable)
    GetEdit()(EditActionable)
    GetMention()(MentionActionable)
    GetMove()(MoveActionable)
    GetOdataType()(*string)
    GetRename()(RenameActionable)
    GetRestore()(RestoreActionable)
    GetShare()(ShareActionable)
    GetVersion()(VersionActionable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetComment(value CommentActionable)()
    SetCreate(value CreateActionable)()
    SetDelete(value DeleteActionable)()
    SetEdit(value EditActionable)()
    SetMention(value MentionActionable)()
    SetMove(value MoveActionable)()
    SetOdataType(value *string)()
    SetRename(value RenameActionable)()
    SetRestore(value RestoreActionable)()
    SetShare(value ShareActionable)()
    SetVersion(value VersionActionable)()
}
