package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RenameAction struct {
    additionalData map[string]interface{};
    newName *string;
    oldName *string;
}
func NewRenameAction()(*RenameAction) {
    m := &RenameAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RenameAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RenameAction) GetNewName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.newName
    }
}
func (m *RenameAction) GetOldName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.oldName
    }
}
func (m *RenameAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["newName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNewName(val)
        return nil
    }
    res["oldName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOldName(val)
        return nil
    }
    return res
}
func (m *RenameAction) IsNil()(bool) {
    return m == nil
}
func (m *RenameAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("newName", m.GetNewName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("oldName", m.GetOldName())
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
func (m *RenameAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RenameAction) SetNewName(value *string)() {
    m.newName = value
}
func (m *RenameAction) SetOldName(value *string)() {
    m.oldName = value
}
