package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OutlookTaskGroup struct {
    Entity
    changeKey *string;
    groupKey *string;
    isDefaultGroup *bool;
    name *string;
    taskFolders []OutlookTaskFolder;
}
func NewOutlookTaskGroup()(*OutlookTaskGroup) {
    m := &OutlookTaskGroup{
        Entity: *NewEntity(),
    }
    return m
}
func (m *OutlookTaskGroup) GetChangeKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.changeKey
    }
}
func (m *OutlookTaskGroup) GetGroupKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupKey
    }
}
func (m *OutlookTaskGroup) GetIsDefaultGroup()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefaultGroup
    }
}
func (m *OutlookTaskGroup) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *OutlookTaskGroup) GetTaskFolders()([]OutlookTaskFolder) {
    if m == nil {
        return nil
    } else {
        return m.taskFolders
    }
}
func (m *OutlookTaskGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["changeKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetChangeKey(val)
        return nil
    }
    res["groupKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGroupKey(val)
        return nil
    }
    res["isDefaultGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefaultGroup(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["taskFolders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutlookTaskFolder() })
        if err != nil {
            return err
        }
        res := make([]OutlookTaskFolder, len(val))
        for i, v := range val {
            res[i] = *(v.(*OutlookTaskFolder))
        }
        m.SetTaskFolders(res)
        return nil
    }
    return res
}
func (m *OutlookTaskGroup) IsNil()(bool) {
    return m == nil
}
func (m *OutlookTaskGroup) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("changeKey", m.GetChangeKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupKey", m.GetGroupKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefaultGroup", m.GetIsDefaultGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTaskFolders()))
        for i, v := range m.GetTaskFolders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("taskFolders", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *OutlookTaskGroup) SetChangeKey(value *string)() {
    m.changeKey = value
}
func (m *OutlookTaskGroup) SetGroupKey(value *string)() {
    m.groupKey = value
}
func (m *OutlookTaskGroup) SetIsDefaultGroup(value *bool)() {
    m.isDefaultGroup = value
}
func (m *OutlookTaskGroup) SetName(value *string)() {
    m.name = value
}
func (m *OutlookTaskGroup) SetTaskFolders(value []OutlookTaskFolder)() {
    m.taskFolders = value
}
