package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// EnrollmentTimeDeviceMembershipTarget the EnrollmentTimeDeviceMembershipTarget entity represents the targets that devices will become members of when enrolled with the associated profile. The only device membership targets supported at this time is static security groups.
type EnrollmentTimeDeviceMembershipTarget struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewEnrollmentTimeDeviceMembershipTarget instantiates a new EnrollmentTimeDeviceMembershipTarget and sets the default values.
func NewEnrollmentTimeDeviceMembershipTarget()(*EnrollmentTimeDeviceMembershipTarget) {
    m := &EnrollmentTimeDeviceMembershipTarget{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnrollmentTimeDeviceMembershipTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnrollmentTimeDeviceMembershipTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnrollmentTimeDeviceMembershipTarget(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnrollmentTimeDeviceMembershipTarget) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *EnrollmentTimeDeviceMembershipTarget) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnrollmentTimeDeviceMembershipTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["targetId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetId(val)
        }
        return nil
    }
    res["targetType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentTimeDeviceMembershipTargetType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetType(val.(*EnrollmentTimeDeviceMembershipTargetType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *EnrollmentTimeDeviceMembershipTarget) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTargetId gets the targetId property value. The unique identifiers of the targets that devices will become members of when enrolled with the asociated profile.
// returns a *string when successful
func (m *EnrollmentTimeDeviceMembershipTarget) GetTargetId()(*string) {
    val, err := m.GetBackingStore().Get("targetId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTargetType gets the targetType property value. Represents the type of the targets that devices will become members of when enrolled with the associated profile. Possible values are staticSecurityGroup.
// returns a *EnrollmentTimeDeviceMembershipTargetType when successful
func (m *EnrollmentTimeDeviceMembershipTarget) GetTargetType()(*EnrollmentTimeDeviceMembershipTargetType) {
    val, err := m.GetBackingStore().Get("targetType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EnrollmentTimeDeviceMembershipTargetType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EnrollmentTimeDeviceMembershipTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetId", m.GetTargetId())
        if err != nil {
            return err
        }
    }
    if m.GetTargetType() != nil {
        cast := (*m.GetTargetType()).String()
        err := writer.WriteStringValue("targetType", &cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnrollmentTimeDeviceMembershipTarget) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *EnrollmentTimeDeviceMembershipTarget) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EnrollmentTimeDeviceMembershipTarget) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetId sets the targetId property value. The unique identifiers of the targets that devices will become members of when enrolled with the asociated profile.
func (m *EnrollmentTimeDeviceMembershipTarget) SetTargetId(value *string)() {
    err := m.GetBackingStore().Set("targetId", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetType sets the targetType property value. Represents the type of the targets that devices will become members of when enrolled with the associated profile. Possible values are staticSecurityGroup.
func (m *EnrollmentTimeDeviceMembershipTarget) SetTargetType(value *EnrollmentTimeDeviceMembershipTargetType)() {
    err := m.GetBackingStore().Set("targetType", value)
    if err != nil {
        panic(err)
    }
}
type EnrollmentTimeDeviceMembershipTargetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetTargetId()(*string)
    GetTargetType()(*EnrollmentTimeDeviceMembershipTargetType)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetTargetId(value *string)()
    SetTargetType(value *EnrollmentTimeDeviceMembershipTargetType)()
}
