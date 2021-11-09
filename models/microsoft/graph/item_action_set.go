package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ItemActionSet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A comment was added to the item.
    comment *CommentAction;
    // An item was created.
    create *CreateAction;
    // An item was deleted.
    delete *DeleteAction;
    // An item was edited.
    edit *EditAction;
    // A user was mentioned in the item.
    mention *MentionAction;
    // An item was moved.
    move *MoveAction;
    // An item was renamed.
    rename *RenameAction;
    // An item was restored.
    restore *RestoreAction;
    // An item was shared.
    share *ShareAction;
    // An item was versioned.
    version *VersionAction;
}
// Instantiates a new itemActionSet and sets the default values.
func NewItemActionSet()(*ItemActionSet) {
    m := &ItemActionSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionSet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the comment property value. A comment was added to the item.
func (m *ItemActionSet) GetComment()(*CommentAction) {
    if m == nil {
        return nil
    } else {
        return m.comment
    }
}
// Gets the create property value. An item was created.
func (m *ItemActionSet) GetCreate()(*CreateAction) {
    if m == nil {
        return nil
    } else {
        return m.create
    }
}
// Gets the delete property value. An item was deleted.
func (m *ItemActionSet) GetDelete()(*DeleteAction) {
    if m == nil {
        return nil
    } else {
        return m.delete
    }
}
// Gets the edit property value. An item was edited.
func (m *ItemActionSet) GetEdit()(*EditAction) {
    if m == nil {
        return nil
    } else {
        return m.edit
    }
}
// Gets the mention property value. A user was mentioned in the item.
func (m *ItemActionSet) GetMention()(*MentionAction) {
    if m == nil {
        return nil
    } else {
        return m.mention
    }
}
// Gets the move property value. An item was moved.
func (m *ItemActionSet) GetMove()(*MoveAction) {
    if m == nil {
        return nil
    } else {
        return m.move
    }
}
// Gets the rename property value. An item was renamed.
func (m *ItemActionSet) GetRename()(*RenameAction) {
    if m == nil {
        return nil
    } else {
        return m.rename
    }
}
// Gets the restore property value. An item was restored.
func (m *ItemActionSet) GetRestore()(*RestoreAction) {
    if m == nil {
        return nil
    } else {
        return m.restore
    }
}
// Gets the share property value. An item was shared.
func (m *ItemActionSet) GetShare()(*ShareAction) {
    if m == nil {
        return nil
    } else {
        return m.share
    }
}
// Gets the version property value. An item was versioned.
func (m *ItemActionSet) GetVersion()(*VersionAction) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *ItemActionSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["comment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCommentAction() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val.(*CommentAction))
        }
        return nil
    }
    res["create"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCreateAction() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreate(val.(*CreateAction))
        }
        return nil
    }
    res["delete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeleteAction() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDelete(val.(*DeleteAction))
        }
        return nil
    }
    res["edit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEditAction() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdit(val.(*EditAction))
        }
        return nil
    }
    res["mention"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMentionAction() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMention(val.(*MentionAction))
        }
        return nil
    }
    res["move"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMoveAction() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMove(val.(*MoveAction))
        }
        return nil
    }
    res["rename"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRenameAction() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRename(val.(*RenameAction))
        }
        return nil
    }
    res["restore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRestoreAction() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestore(val.(*RestoreAction))
        }
        return nil
    }
    res["share"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShareAction() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShare(val.(*ShareAction))
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVersionAction() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val.(*VersionAction))
        }
        return nil
    }
    return res
}
func (m *ItemActionSet) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ItemActionSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ItemActionSet) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the comment property value. A comment was added to the item.
// Parameters:
//  - value : Value to set for the comment property.
func (m *ItemActionSet) SetComment(value *CommentAction)() {
    m.comment = value
}
// Sets the create property value. An item was created.
// Parameters:
//  - value : Value to set for the create property.
func (m *ItemActionSet) SetCreate(value *CreateAction)() {
    m.create = value
}
// Sets the delete property value. An item was deleted.
// Parameters:
//  - value : Value to set for the delete property.
func (m *ItemActionSet) SetDelete(value *DeleteAction)() {
    m.delete = value
}
// Sets the edit property value. An item was edited.
// Parameters:
//  - value : Value to set for the edit property.
func (m *ItemActionSet) SetEdit(value *EditAction)() {
    m.edit = value
}
// Sets the mention property value. A user was mentioned in the item.
// Parameters:
//  - value : Value to set for the mention property.
func (m *ItemActionSet) SetMention(value *MentionAction)() {
    m.mention = value
}
// Sets the move property value. An item was moved.
// Parameters:
//  - value : Value to set for the move property.
func (m *ItemActionSet) SetMove(value *MoveAction)() {
    m.move = value
}
// Sets the rename property value. An item was renamed.
// Parameters:
//  - value : Value to set for the rename property.
func (m *ItemActionSet) SetRename(value *RenameAction)() {
    m.rename = value
}
// Sets the restore property value. An item was restored.
// Parameters:
//  - value : Value to set for the restore property.
func (m *ItemActionSet) SetRestore(value *RestoreAction)() {
    m.restore = value
}
// Sets the share property value. An item was shared.
// Parameters:
//  - value : Value to set for the share property.
func (m *ItemActionSet) SetShare(value *ShareAction)() {
    m.share = value
}
// Sets the version property value. An item was versioned.
// Parameters:
//  - value : Value to set for the version property.
func (m *ItemActionSet) SetVersion(value *VersionAction)() {
    m.version = value
}
