package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedEBookCategory 
type ManagedEBookCategory struct {
    Entity
    // The name of the eBook category.
    displayName *string;
    // The date and time the ManagedEBookCategory was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewManagedEBookCategory instantiates a new managedEBookCategory and sets the default values.
func NewManagedEBookCategory()(*ManagedEBookCategory) {
    m := &ManagedEBookCategory{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. The name of the eBook category.
func (m *ManagedEBookCategory) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the ManagedEBookCategory was last modified.
func (m *ManagedEBookCategory) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedEBookCategory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    return res
}
func (m *ManagedEBookCategory) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedEBookCategory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
    return nil
}
// SetDisplayName sets the displayName property value. The name of the eBook category.
func (m *ManagedEBookCategory) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the ManagedEBookCategory was last modified.
func (m *ManagedEBookCategory) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
