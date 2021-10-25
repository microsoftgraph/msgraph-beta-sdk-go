package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ItemActionSet struct {
    additionalData map[string]interface{};
    comment *CommentAction;
    create *CreateAction;
    delete *DeleteAction;
    edit *EditAction;
    mention *MentionAction;
    move *MoveAction;
    rename *RenameAction;
    restore *RestoreAction;
    share *ShareAction;
    version *VersionAction;
}
func NewItemActionSet()(*ItemActionSet) {
    m := &ItemActionSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ItemActionSet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ItemActionSet) GetComment()(*CommentAction) {
    if m == nil {
        return nil
    } else {
        return m.comment
    }
}
func (m *ItemActionSet) GetCreate()(*CreateAction) {
    if m == nil {
        return nil
    } else {
        return m.create
    }
}
func (m *ItemActionSet) GetDelete()(*DeleteAction) {
    if m == nil {
        return nil
    } else {
        return m.delete
    }
}
func (m *ItemActionSet) GetEdit()(*EditAction) {
    if m == nil {
        return nil
    } else {
        return m.edit
    }
}
func (m *ItemActionSet) GetMention()(*MentionAction) {
    if m == nil {
        return nil
    } else {
        return m.mention
    }
}
func (m *ItemActionSet) GetMove()(*MoveAction) {
    if m == nil {
        return nil
    } else {
        return m.move
    }
}
func (m *ItemActionSet) GetRename()(*RenameAction) {
    if m == nil {
        return nil
    } else {
        return m.rename
    }
}
func (m *ItemActionSet) GetRestore()(*RestoreAction) {
    if m == nil {
        return nil
    } else {
        return m.restore
    }
}
func (m *ItemActionSet) GetShare()(*ShareAction) {
    if m == nil {
        return nil
    } else {
        return m.share
    }
}
func (m *ItemActionSet) GetVersion()(*VersionAction) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
func (m *ItemActionSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["comment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCommentAction() })
        if err != nil {
            return err
        }
        m.SetComment(val.(*CommentAction))
        return nil
    }
    res["create"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCreateAction() })
        if err != nil {
            return err
        }
        m.SetCreate(val.(*CreateAction))
        return nil
    }
    res["delete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeleteAction() })
        if err != nil {
            return err
        }
        m.SetDelete(val.(*DeleteAction))
        return nil
    }
    res["edit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEditAction() })
        if err != nil {
            return err
        }
        m.SetEdit(val.(*EditAction))
        return nil
    }
    res["mention"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMentionAction() })
        if err != nil {
            return err
        }
        m.SetMention(val.(*MentionAction))
        return nil
    }
    res["move"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMoveAction() })
        if err != nil {
            return err
        }
        m.SetMove(val.(*MoveAction))
        return nil
    }
    res["rename"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRenameAction() })
        if err != nil {
            return err
        }
        m.SetRename(val.(*RenameAction))
        return nil
    }
    res["restore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRestoreAction() })
        if err != nil {
            return err
        }
        m.SetRestore(val.(*RestoreAction))
        return nil
    }
    res["share"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShareAction() })
        if err != nil {
            return err
        }
        m.SetShare(val.(*ShareAction))
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVersionAction() })
        if err != nil {
            return err
        }
        m.SetVersion(val.(*VersionAction))
        return nil
    }
    return res
}
func (m *ItemActionSet) IsNil()(bool) {
    return m == nil
}
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
func (m *ItemActionSet) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ItemActionSet) SetComment(value *CommentAction)() {
    m.comment = value
}
func (m *ItemActionSet) SetCreate(value *CreateAction)() {
    m.create = value
}
func (m *ItemActionSet) SetDelete(value *DeleteAction)() {
    m.delete = value
}
func (m *ItemActionSet) SetEdit(value *EditAction)() {
    m.edit = value
}
func (m *ItemActionSet) SetMention(value *MentionAction)() {
    m.mention = value
}
func (m *ItemActionSet) SetMove(value *MoveAction)() {
    m.move = value
}
func (m *ItemActionSet) SetRename(value *RenameAction)() {
    m.rename = value
}
func (m *ItemActionSet) SetRestore(value *RestoreAction)() {
    m.restore = value
}
func (m *ItemActionSet) SetShare(value *ShareAction)() {
    m.share = value
}
func (m *ItemActionSet) SetVersion(value *VersionAction)() {
    m.version = value
}
