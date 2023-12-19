package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// MicrosoftManagedDesktop 
type MicrosoftManagedDesktop struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMicrosoftManagedDesktop instantiates a new microsoftManagedDesktop and sets the default values.
func NewMicrosoftManagedDesktop()(*MicrosoftManagedDesktop) {
    m := &MicrosoftManagedDesktop{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMicrosoftManagedDesktopFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftManagedDesktopFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftManagedDesktop(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftManagedDesktop) GetAdditionalData()(map[string]any) {
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
func (m *MicrosoftManagedDesktop) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftManagedDesktop) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managedType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftManagedDesktop_managedType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedType(val.(*MicrosoftManagedDesktop_managedType))
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
    res["profile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfile(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftManagedDesktop_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*MicrosoftManagedDesktop_type))
        }
        return nil
    }
    return res
}
// GetManagedType gets the managedType property value. The managedType property
func (m *MicrosoftManagedDesktop) GetManagedType()(*MicrosoftManagedDesktop_managedType) {
    val, err := m.GetBackingStore().Get("managedType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MicrosoftManagedDesktop_managedType)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MicrosoftManagedDesktop) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProfile gets the profile property value. The name of the Microsoft Managed Desktop profile that the Windows 365 Cloud PC is associated with.
func (m *MicrosoftManagedDesktop) GetProfile()(*string) {
    val, err := m.GetBackingStore().Get("profile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTypeEscaped gets the type property value. Indicates whether the provisioning policy enables Microsoft Managed Desktop. It indicates the type of plan under which the device is managed if the provisioning policy is enabled. Possible values are: notManaged, premiumManaged, standardManaged, starterManaged, unknownFutureValue.
func (m *MicrosoftManagedDesktop) GetTypeEscaped()(*MicrosoftManagedDesktop_type) {
    val, err := m.GetBackingStore().Get("typeEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MicrosoftManagedDesktop_type)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MicrosoftManagedDesktop) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetManagedType() != nil {
        cast := (*m.GetManagedType()).String()
        err := writer.WriteStringValue("managedType", &cast)
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
        err := writer.WriteStringValue("profile", m.GetProfile())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *MicrosoftManagedDesktop) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *MicrosoftManagedDesktop) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetManagedType sets the managedType property value. The managedType property
func (m *MicrosoftManagedDesktop) SetManagedType(value *MicrosoftManagedDesktop_managedType)() {
    err := m.GetBackingStore().Set("managedType", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MicrosoftManagedDesktop) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetProfile sets the profile property value. The name of the Microsoft Managed Desktop profile that the Windows 365 Cloud PC is associated with.
func (m *MicrosoftManagedDesktop) SetProfile(value *string)() {
    err := m.GetBackingStore().Set("profile", value)
    if err != nil {
        panic(err)
    }
}
// SetTypeEscaped sets the type property value. Indicates whether the provisioning policy enables Microsoft Managed Desktop. It indicates the type of plan under which the device is managed if the provisioning policy is enabled. Possible values are: notManaged, premiumManaged, standardManaged, starterManaged, unknownFutureValue.
func (m *MicrosoftManagedDesktop) SetTypeEscaped(value *MicrosoftManagedDesktop_type)() {
    err := m.GetBackingStore().Set("typeEscaped", value)
    if err != nil {
        panic(err)
    }
}
// MicrosoftManagedDesktopable 
type MicrosoftManagedDesktopable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetManagedType()(*MicrosoftManagedDesktop_managedType)
    GetOdataType()(*string)
    GetProfile()(*string)
    GetTypeEscaped()(*MicrosoftManagedDesktop_type)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetManagedType(value *MicrosoftManagedDesktop_managedType)()
    SetOdataType(value *string)()
    SetProfile(value *string)()
    SetTypeEscaped(value *MicrosoftManagedDesktop_type)()
}
