package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemActivityOLD provides operations to manage the deviceManagement singleton.
type ItemActivityOLD struct {
    Entity
    // 
    action ItemActionSetable;
    // 
    actor IdentitySetable;
    // 
    driveItem DriveItemable;
    // 
    listItem ListItemable;
    // 
    times ItemActivityTimeSetable;
}
// NewItemActivityOLD instantiates a new itemActivityOLD and sets the default values.
func NewItemActivityOLD()(*ItemActivityOLD) {
    m := &ItemActivityOLD{
        Entity: *NewEntity(),
    }
    return m
}
// CreateItemActivityOLDFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActivityOLDFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewItemActivityOLD(), nil
}
// GetAction gets the action property value. 
func (m *ItemActivityOLD) GetAction()(ItemActionSetable) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// GetActor gets the actor property value. 
func (m *ItemActivityOLD) GetActor()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.actor
    }
}
// GetDriveItem gets the driveItem property value. 
func (m *ItemActivityOLD) GetDriveItem()(DriveItemable) {
    if m == nil {
        return nil
    } else {
        return m.driveItem
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActivityOLD) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemActionSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(ItemActionSetable))
        }
        return nil
    }
    res["actor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActor(val.(IdentitySetable))
        }
        return nil
    }
    res["driveItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriveItem(val.(DriveItemable))
        }
        return nil
    }
    res["listItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetListItem(val.(ListItemable))
        }
        return nil
    }
    res["times"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemActivityTimeSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimes(val.(ItemActivityTimeSetable))
        }
        return nil
    }
    return res
}
// GetListItem gets the listItem property value. 
func (m *ItemActivityOLD) GetListItem()(ListItemable) {
    if m == nil {
        return nil
    } else {
        return m.listItem
    }
}
// GetTimes gets the times property value. 
func (m *ItemActivityOLD) GetTimes()(ItemActivityTimeSetable) {
    if m == nil {
        return nil
    } else {
        return m.times
    }
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
func (m *ItemActivityOLD) SetAction(value ItemActionSetable)() {
    if m != nil {
        m.action = value
    }
}
// SetActor sets the actor property value. 
func (m *ItemActivityOLD) SetActor(value IdentitySetable)() {
    if m != nil {
        m.actor = value
    }
}
// SetDriveItem sets the driveItem property value. 
func (m *ItemActivityOLD) SetDriveItem(value DriveItemable)() {
    if m != nil {
        m.driveItem = value
    }
}
// SetListItem sets the listItem property value. 
func (m *ItemActivityOLD) SetListItem(value ListItemable)() {
    if m != nil {
        m.listItem = value
    }
}
// SetTimes sets the times property value. 
func (m *ItemActivityOLD) SetTimes(value ItemActivityTimeSetable)() {
    if m != nil {
        m.times = value
    }
}
