package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ItemActivityOLD struct {
    Entity
    action *ItemActionSet;
    actor *IdentitySet;
    driveItem *DriveItem;
    listItem *ListItem;
    times *ItemActivityTimeSet;
}
func NewItemActivityOLD()(*ItemActivityOLD) {
    m := &ItemActivityOLD{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ItemActivityOLD) GetAction()(*ItemActionSet) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
func (m *ItemActivityOLD) GetActor()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.actor
    }
}
func (m *ItemActivityOLD) GetDriveItem()(*DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.driveItem
    }
}
func (m *ItemActivityOLD) GetListItem()(*ListItem) {
    if m == nil {
        return nil
    } else {
        return m.listItem
    }
}
func (m *ItemActivityOLD) GetTimes()(*ItemActivityTimeSet) {
    if m == nil {
        return nil
    } else {
        return m.times
    }
}
func (m *ItemActivityOLD) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActionSet() })
        if err != nil {
            return err
        }
        m.SetAction(val.(*ItemActionSet))
        return nil
    }
    res["actor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetActor(val.(*IdentitySet))
        return nil
    }
    res["driveItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItem() })
        if err != nil {
            return err
        }
        m.SetDriveItem(val.(*DriveItem))
        return nil
    }
    res["listItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewListItem() })
        if err != nil {
            return err
        }
        m.SetListItem(val.(*ListItem))
        return nil
    }
    res["times"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActivityTimeSet() })
        if err != nil {
            return err
        }
        m.SetTimes(val.(*ItemActivityTimeSet))
        return nil
    }
    return res
}
func (m *ItemActivityOLD) IsNil()(bool) {
    return m == nil
}
func (m *ItemActivityOLD) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("action", m.GetAction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("actor", m.GetActor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("driveItem", m.GetDriveItem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("listItem", m.GetListItem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("times", m.GetTimes())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ItemActivityOLD) SetAction(value *ItemActionSet)() {
    m.action = value
}
func (m *ItemActivityOLD) SetActor(value *IdentitySet)() {
    m.actor = value
}
func (m *ItemActivityOLD) SetDriveItem(value *DriveItem)() {
    m.driveItem = value
}
func (m *ItemActivityOLD) SetListItem(value *ListItem)() {
    m.listItem = value
}
func (m *ItemActivityOLD) SetTimes(value *ItemActivityTimeSet)() {
    m.times = value
}
