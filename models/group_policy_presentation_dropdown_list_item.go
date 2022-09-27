package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationDropdownListItem 
type GroupPolicyPresentationDropdownListItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Localized display name for the drop-down list item.
    displayName *string
    // The OdataType property
    odataType *string
    // Associated value for the drop-down list item
    value *string
}
// NewGroupPolicyPresentationDropdownListItem instantiates a new groupPolicyPresentationDropdownListItem and sets the default values.
func NewGroupPolicyPresentationDropdownListItem()(*GroupPolicyPresentationDropdownListItem) {
    m := &GroupPolicyPresentationDropdownListItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.groupPolicyPresentationDropdownListItem";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupPolicyPresentationDropdownListItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationDropdownListItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationDropdownListItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupPolicyPresentationDropdownListItem) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Localized display name for the drop-down list item.
func (m *GroupPolicyPresentationDropdownListItem) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationDropdownListItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetValue)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *GroupPolicyPresentationDropdownListItem) GetOdataType()(*string) {
    return m.odataType
}
// GetValue gets the value property value. Associated value for the drop-down list item
func (m *GroupPolicyPresentationDropdownListItem) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationDropdownListItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupPolicyPresentationDropdownListItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Localized display name for the drop-down list item.
func (m *GroupPolicyPresentationDropdownListItem) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GroupPolicyPresentationDropdownListItem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValue sets the value property value. Associated value for the drop-down list item
func (m *GroupPolicyPresentationDropdownListItem) SetValue(value *string)() {
    m.value = value
}
