package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttributeSet provides operations to manage the collection of activityStatistics entities.
type AttributeSet struct {
    Entity
    // Description of the attribute set. Can be up to 128 characters long and include Unicode characters. Can be changed later.
    description *string
    // Maximum number of custom security attributes that can be defined in this attribute set. Default value is null. If not specified, the administrator can add up to the maximum of 500 active attributes per tenant. Can be changed later.
    maxAttributesPerSet *int32
}
// NewAttributeSet instantiates a new attributeSet and sets the default values.
func NewAttributeSet()(*AttributeSet) {
    m := &AttributeSet{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAttributeSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttributeSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttributeSet(), nil
}
// GetDescription gets the description property value. Description of the attribute set. Can be up to 128 characters long and include Unicode characters. Can be changed later.
func (m *AttributeSet) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["maxAttributesPerSet"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetMaxAttributesPerSet)
    return res
}
// GetMaxAttributesPerSet gets the maxAttributesPerSet property value. Maximum number of custom security attributes that can be defined in this attribute set. Default value is null. If not specified, the administrator can add up to the maximum of 500 active attributes per tenant. Can be changed later.
func (m *AttributeSet) GetMaxAttributesPerSet()(*int32) {
    return m.maxAttributesPerSet
}
// Serialize serializes information the current object
func (m *AttributeSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("maxAttributesPerSet", m.GetMaxAttributesPerSet())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Description of the attribute set. Can be up to 128 characters long and include Unicode characters. Can be changed later.
func (m *AttributeSet) SetDescription(value *string)() {
    m.description = value
}
// SetMaxAttributesPerSet sets the maxAttributesPerSet property value. Maximum number of custom security attributes that can be defined in this attribute set. Default value is null. If not specified, the administrator can add up to the maximum of 500 active attributes per tenant. Can be changed later.
func (m *AttributeSet) SetMaxAttributesPerSet(value *int32)() {
    m.maxAttributesPerSet = value
}
