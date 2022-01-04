package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConnectionQuota 
type ConnectionQuota struct {
    Entity
    // 
    itemsRemaining *int64;
}
// NewConnectionQuota instantiates a new connectionQuota and sets the default values.
func NewConnectionQuota()(*ConnectionQuota) {
    m := &ConnectionQuota{
        Entity: *NewEntity(),
    }
    return m
}
// GetItemsRemaining gets the itemsRemaining property value. 
func (m *ConnectionQuota) GetItemsRemaining()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.itemsRemaining
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConnectionQuota) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["itemsRemaining"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemsRemaining(val)
        }
        return nil
    }
    return res
}
func (m *ConnectionQuota) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ConnectionQuota) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("itemsRemaining", m.GetItemsRemaining())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemsRemaining sets the itemsRemaining property value. 
func (m *ConnectionQuota) SetItemsRemaining(value *int64)() {
    if m != nil {
        m.itemsRemaining = value
    }
}
