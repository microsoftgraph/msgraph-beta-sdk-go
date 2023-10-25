package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegeEscalation 
type PrivilegeEscalation struct {
    Entity
}
// NewPrivilegeEscalation instantiates a new privilegeEscalation and sets the default values.
func NewPrivilegeEscalation()(*PrivilegeEscalation) {
    m := &PrivilegeEscalation{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegeEscalationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegeEscalationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegeEscalation(), nil
}
// GetActions gets the actions property value. The actions property
func (m *PrivilegeEscalation) GetActions()([]AuthorizationSystemTypeActionable) {
    val, err := m.GetBackingStore().Get("actions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthorizationSystemTypeActionable)
    }
    return nil
}
// GetDescription gets the description property value. The description property
func (m *PrivilegeEscalation) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *PrivilegeEscalation) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegeEscalation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthorizationSystemTypeActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthorizationSystemTypeActionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthorizationSystemTypeActionable)
                }
            }
            m.SetActions(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthorizationSystemResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthorizationSystemResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthorizationSystemResourceable)
                }
            }
            m.SetResources(res)
        }
        return nil
    }
    return res
}
// GetResources gets the resources property value. The resources property
func (m *PrivilegeEscalation) GetResources()([]AuthorizationSystemResourceable) {
    val, err := m.GetBackingStore().Get("resources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthorizationSystemResourceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PrivilegeEscalation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActions()))
        for i, v := range m.GetActions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("actions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActions sets the actions property value. The actions property
func (m *PrivilegeEscalation) SetActions(value []AuthorizationSystemTypeActionable)() {
    err := m.GetBackingStore().Set("actions", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *PrivilegeEscalation) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *PrivilegeEscalation) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetResources sets the resources property value. The resources property
func (m *PrivilegeEscalation) SetResources(value []AuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("resources", value)
    if err != nil {
        panic(err)
    }
}
// PrivilegeEscalationable 
type PrivilegeEscalationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]AuthorizationSystemTypeActionable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetResources()([]AuthorizationSystemResourceable)
    SetActions(value []AuthorizationSystemTypeActionable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetResources(value []AuthorizationSystemResourceable)()
}
