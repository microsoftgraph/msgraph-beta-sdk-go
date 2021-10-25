package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AccessReviewPolicy struct {
    Entity
    description *string;
    displayName *string;
    isGroupOwnerManagementEnabled *bool;
}
func NewAccessReviewPolicy()(*AccessReviewPolicy) {
    m := &AccessReviewPolicy{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AccessReviewPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *AccessReviewPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *AccessReviewPolicy) GetIsGroupOwnerManagementEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isGroupOwnerManagementEnabled
    }
}
func (m *AccessReviewPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["isGroupOwnerManagementEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsGroupOwnerManagementEnabled(val)
        return nil
    }
    return res
}
func (m *AccessReviewPolicy) IsNil()(bool) {
    return m == nil
}
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
func (m *AccessReviewPolicy) SetDescription(value *string)() {
    m.description = value
}
func (m *AccessReviewPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *AccessReviewPolicy) SetIsGroupOwnerManagementEnabled(value *bool)() {
    m.isGroupOwnerManagementEnabled = value
}
