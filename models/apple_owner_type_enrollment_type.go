package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type AppleOwnerTypeEnrollmentType struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAppleOwnerTypeEnrollmentType instantiates a new AppleOwnerTypeEnrollmentType and sets the default values.
func NewAppleOwnerTypeEnrollmentType()(*AppleOwnerTypeEnrollmentType) {
    m := &AppleOwnerTypeEnrollmentType{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAppleOwnerTypeEnrollmentTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppleOwnerTypeEnrollmentTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppleOwnerTypeEnrollmentType(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AppleOwnerTypeEnrollmentType) GetAdditionalData()(map[string]any) {
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
func (m *AppleOwnerTypeEnrollmentType) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEnrollmentType gets the enrollmentType property value. The enrollmentType property
// returns a *AppleUserInitiatedEnrollmentType when successful
func (m *AppleOwnerTypeEnrollmentType) GetEnrollmentType()(*AppleUserInitiatedEnrollmentType) {
    val, err := m.GetBackingStore().Get("enrollmentType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppleUserInitiatedEnrollmentType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AppleOwnerTypeEnrollmentType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enrollmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppleUserInitiatedEnrollmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentType(val.(*AppleUserInitiatedEnrollmentType))
        }
        return nil
    }
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
    res["ownerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerType(val.(*ManagedDeviceOwnerType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AppleOwnerTypeEnrollmentType) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOwnerType gets the ownerType property value. Owner type of device.
// returns a *ManagedDeviceOwnerType when successful
func (m *AppleOwnerTypeEnrollmentType) GetOwnerType()(*ManagedDeviceOwnerType) {
    val, err := m.GetBackingStore().Get("ownerType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedDeviceOwnerType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AppleOwnerTypeEnrollmentType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnrollmentType() != nil {
        cast := (*m.GetEnrollmentType()).String()
        err := writer.WriteStringValue("enrollmentType", &cast)
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
    if m.GetOwnerType() != nil {
        cast := (*m.GetOwnerType()).String()
        err := writer.WriteStringValue("ownerType", &cast)
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
func (m *AppleOwnerTypeEnrollmentType) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AppleOwnerTypeEnrollmentType) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEnrollmentType sets the enrollmentType property value. The enrollmentType property
func (m *AppleOwnerTypeEnrollmentType) SetEnrollmentType(value *AppleUserInitiatedEnrollmentType)() {
    err := m.GetBackingStore().Set("enrollmentType", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppleOwnerTypeEnrollmentType) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOwnerType sets the ownerType property value. Owner type of device.
func (m *AppleOwnerTypeEnrollmentType) SetOwnerType(value *ManagedDeviceOwnerType)() {
    err := m.GetBackingStore().Set("ownerType", value)
    if err != nil {
        panic(err)
    }
}
type AppleOwnerTypeEnrollmentTypeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEnrollmentType()(*AppleUserInitiatedEnrollmentType)
    GetOdataType()(*string)
    GetOwnerType()(*ManagedDeviceOwnerType)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEnrollmentType(value *AppleUserInitiatedEnrollmentType)()
    SetOdataType(value *string)()
    SetOwnerType(value *ManagedDeviceOwnerType)()
}
