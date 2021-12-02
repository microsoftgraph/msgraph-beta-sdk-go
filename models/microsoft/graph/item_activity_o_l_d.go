package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemActivityOLD 
type ItemActivityOLD struct {
    Entity
    // 
    action *ItemActionSet;
    // 
    actor *IdentitySet;
    // 
    driveItem *DriveItem;
    // 
    listItem *ListItem;
    // 
    times *ItemActivityTimeSet;
}
// NewItemActivityOLD instantiates a new itemActivityOLD and sets the default values.
func NewItemActivityOLD()(*ItemActivityOLD) {
    m := &ItemActivityOLD{
        Entity: *NewEntity(),
    }
    return m
}
// GetAction gets the action property value. 
func (m *ItemActivityOLD) GetAction()(*ItemActionSet) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// GetActor gets the actor property value. 
func (m *ItemActivityOLD) GetActor()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.actor
    }
}
// GetDriveItem gets the driveItem property value. 
func (m *ItemActivityOLD) GetDriveItem()(*DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.driveItem
    }
}
// GetListItem gets the listItem property value. 
func (m *ItemActivityOLD) GetListItem()(*ListItem) {
    if m == nil {
        return nil
    } else {
        return m.listItem
    }
}
// GetTimes gets the times property value. 
func (m *ItemActivityOLD) GetTimes()(*ItemActivityTimeSet) {
    if m == nil {
        return nil
    } else {
        return m.times
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActivityOLD) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActionSet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*ItemActionSet))
        }
        return nil
    }
    res["actor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActor(val.(*IdentitySet))
        }
        return nil
    }
    res["driveItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriveItem(val.(*DriveItem))
        }
        return nil
    }
    res["listItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewListItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetListItem(val.(*ListItem))
        }
        return nil
    }
    res["times"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActivityTimeSet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimes(val.(*ItemActivityTimeSet))
        }
        return nil
    }
    return res
}
func (m *ItemActivityOLD) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAction sets the action property value. 
func (m *ItemActivityOLD) SetAction(value *ItemActionSet)() {
    if m != nil {
        m.action = value
    }
}
// SetActor sets the actor property value. 
func (m *ItemActivityOLD) SetActor(value *IdentitySet)() {
    if m != nil {
        m.actor = value
    }
}
// SetDriveItem sets the driveItem property value. 
func (m *ItemActivityOLD) SetDriveItem(value *DriveItem)() {
    if m != nil {
        m.driveItem = value
    }
}
// SetListItem sets the listItem property value. 
func (m *ItemActivityOLD) SetListItem(value *ListItem)() {
    if m != nil {
        m.listItem = value
    }
}
// SetTimes sets the times property value. 
func (m *ItemActivityOLD) SetTimes(value *ItemActivityTimeSet)() {
    if m != nil {
        m.times = value
    }
}
