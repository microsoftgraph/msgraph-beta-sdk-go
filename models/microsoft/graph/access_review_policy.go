package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewPolicy 
type AccessReviewPolicy struct {
    Entity
    // Description for this policy. Read-only.
    description *string;
    // Display name for this policy. Read-only.
    displayName *string;
    // If true, group owners can create and manage access reviews on groups they own.
    isGroupOwnerManagementEnabled *bool;
}
// NewAccessReviewPolicy instantiates a new accessReviewPolicy and sets the default values.
func NewAccessReviewPolicy()(*AccessReviewPolicy) {
    m := &AccessReviewPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// GetDescription gets the description property value. Description for this policy. Read-only.
func (m *AccessReviewPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Display name for this policy. Read-only.
func (m *AccessReviewPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIsGroupOwnerManagementEnabled gets the isGroupOwnerManagementEnabled property value. If true, group owners can create and manage access reviews on groups they own.
func (m *AccessReviewPolicy) GetIsGroupOwnerManagementEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isGroupOwnerManagementEnabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
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
    res["isGroupOwnerManagementEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGroupOwnerManagementEnabled(val)
        }
        return nil
    }
    return res
}
func (m *AccessReviewPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessReviewPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
    {
        err = writer.WriteBoolValue("isGroupOwnerManagementEnabled", m.GetIsGroupOwnerManagementEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Description for this policy. Read-only.
func (m *AccessReviewPolicy) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Display name for this policy. Read-only.
func (m *AccessReviewPolicy) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsGroupOwnerManagementEnabled sets the isGroupOwnerManagementEnabled property value. If true, group owners can create and manage access reviews on groups they own.
func (m *AccessReviewPolicy) SetIsGroupOwnerManagementEnabled(value *bool)() {
    if m != nil {
        m.isGroupOwnerManagementEnabled = value
    }
}
