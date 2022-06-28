package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationDropdownList 
type GroupPolicyPresentationDropdownList struct {
    GroupPolicyPresentation
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Localized string value identifying the default choice of the list of items.
    defaultItem GroupPolicyPresentationDropdownListItemable
    // Represents a set of localized display names and their associated values.
    items []GroupPolicyPresentationDropdownListItemable
    // Requirement to enter a value in the parameter box. The default value is false.
    required *bool
}
// NewGroupPolicyPresentationDropdownList instantiates a new GroupPolicyPresentationDropdownList and sets the default values.
func NewGroupPolicyPresentationDropdownList()(*GroupPolicyPresentationDropdownList) {
    m := &GroupPolicyPresentationDropdownList{
        GroupPolicyPresentation: *NewGroupPolicyPresentation(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGroupPolicyPresentationDropdownListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationDropdownListFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationDropdownList(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupPolicyPresentationDropdownList) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefaultItem gets the defaultItem property value. Localized string value identifying the default choice of the list of items.
func (m *GroupPolicyPresentationDropdownList) GetDefaultItem()(GroupPolicyPresentationDropdownListItemable) {
    if m == nil {
        return nil
    } else {
        return m.defaultItem
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationDropdownList) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyPresentation.GetFieldDeserializers()
    res["defaultItem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupPolicyPresentationDropdownListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultItem(val.(GroupPolicyPresentationDropdownListItemable))
        }
        return nil
    }
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyPresentationDropdownListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyPresentationDropdownListItemable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyPresentationDropdownListItemable)
            }
            m.SetItems(res)
        }
        return nil
    }
    res["required"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequired(val)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. Represents a set of localized display names and their associated values.
func (m *GroupPolicyPresentationDropdownList) GetItems()([]GroupPolicyPresentationDropdownListItemable) {
    if m == nil {
        return nil
    } else {
        return m.items
    }
}
// GetRequired gets the required property value. Requirement to enter a value in the parameter box. The default value is false.
func (m *GroupPolicyPresentationDropdownList) GetRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.required
    }
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationDropdownList) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyPresentation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("defaultItem", m.GetDefaultItem())
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("required", m.GetRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupPolicyPresentationDropdownList) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDefaultItem sets the defaultItem property value. Localized string value identifying the default choice of the list of items.
func (m *GroupPolicyPresentationDropdownList) SetDefaultItem(value GroupPolicyPresentationDropdownListItemable)() {
    if m != nil {
        m.defaultItem = value
    }
}
// SetItems sets the items property value. Represents a set of localized display names and their associated values.
func (m *GroupPolicyPresentationDropdownList) SetItems(value []GroupPolicyPresentationDropdownListItemable)() {
    if m != nil {
        m.items = value
    }
}
// SetRequired sets the required property value. Requirement to enter a value in the parameter box. The default value is false.
func (m *GroupPolicyPresentationDropdownList) SetRequired(value *bool)() {
    if m != nil {
        m.required = value
    }
}
