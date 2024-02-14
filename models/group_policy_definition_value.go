package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyDefinitionValue the definition value entity stores the value for a single group policy definition.
type GroupPolicyDefinitionValue struct {
    Entity
}
// NewGroupPolicyDefinitionValue instantiates a new GroupPolicyDefinitionValue and sets the default values.
func NewGroupPolicyDefinitionValue()(*GroupPolicyDefinitionValue) {
    m := &GroupPolicyDefinitionValue{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGroupPolicyDefinitionValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupPolicyDefinitionValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyDefinitionValue(), nil
}
// GetConfigurationType gets the configurationType property value. Group Policy Configuration Type
// returns a *GroupPolicyConfigurationType when successful
func (m *GroupPolicyDefinitionValue) GetConfigurationType()(*GroupPolicyConfigurationType) {
    val, err := m.GetBackingStore().Get("configurationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*GroupPolicyConfigurationType)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the object was created.
// returns a *Time when successful
func (m *GroupPolicyDefinitionValue) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDefinition gets the definition property value. The associated group policy definition with the value.
// returns a GroupPolicyDefinitionable when successful
func (m *GroupPolicyDefinitionValue) GetDefinition()(GroupPolicyDefinitionable) {
    val, err := m.GetBackingStore().Get("definition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GroupPolicyDefinitionable)
    }
    return nil
}
// GetEnabled gets the enabled property value. Enables or disables the associated group policy definition.
// returns a *bool when successful
func (m *GroupPolicyDefinitionValue) GetEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("enabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupPolicyDefinitionValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configurationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyConfigurationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationType(val.(*GroupPolicyConfigurationType))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["definition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupPolicyDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinition(val.(GroupPolicyDefinitionable))
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["presentationValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyPresentationValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyPresentationValueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupPolicyPresentationValueable)
                }
            }
            m.SetPresentationValues(res)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
// returns a *Time when successful
func (m *GroupPolicyDefinitionValue) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetPresentationValues gets the presentationValues property value. The associated group policy presentation values with the definition value.
// returns a []GroupPolicyPresentationValueable when successful
func (m *GroupPolicyDefinitionValue) GetPresentationValues()([]GroupPolicyPresentationValueable) {
    val, err := m.GetBackingStore().Get("presentationValues")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GroupPolicyPresentationValueable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GroupPolicyDefinitionValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConfigurationType() != nil {
        cast := (*m.GetConfigurationType()).String()
        err = writer.WriteStringValue("configurationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("definition", m.GetDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPresentationValues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPresentationValues()))
        for i, v := range m.GetPresentationValues() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("presentationValues", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfigurationType sets the configurationType property value. Group Policy Configuration Type
func (m *GroupPolicyDefinitionValue) SetConfigurationType(value *GroupPolicyConfigurationType)() {
    err := m.GetBackingStore().Set("configurationType", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the object was created.
func (m *GroupPolicyDefinitionValue) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDefinition sets the definition property value. The associated group policy definition with the value.
func (m *GroupPolicyDefinitionValue) SetDefinition(value GroupPolicyDefinitionable)() {
    err := m.GetBackingStore().Set("definition", value)
    if err != nil {
        panic(err)
    }
}
// SetEnabled sets the enabled property value. Enables or disables the associated group policy definition.
func (m *GroupPolicyDefinitionValue) SetEnabled(value *bool)() {
    err := m.GetBackingStore().Set("enabled", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyDefinitionValue) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPresentationValues sets the presentationValues property value. The associated group policy presentation values with the definition value.
func (m *GroupPolicyDefinitionValue) SetPresentationValues(value []GroupPolicyPresentationValueable)() {
    err := m.GetBackingStore().Set("presentationValues", value)
    if err != nil {
        panic(err)
    }
}
type GroupPolicyDefinitionValueable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfigurationType()(*GroupPolicyConfigurationType)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDefinition()(GroupPolicyDefinitionable)
    GetEnabled()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPresentationValues()([]GroupPolicyPresentationValueable)
    SetConfigurationType(value *GroupPolicyConfigurationType)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDefinition(value GroupPolicyDefinitionable)()
    SetEnabled(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPresentationValues(value []GroupPolicyPresentationValueable)()
}
