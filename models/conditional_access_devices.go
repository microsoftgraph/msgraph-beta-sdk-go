package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ConditionalAccessDevices 
type ConditionalAccessDevices struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewConditionalAccessDevices instantiates a new conditionalAccessDevices and sets the default values.
func NewConditionalAccessDevices()(*ConditionalAccessDevices) {
    m := &ConditionalAccessDevices{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConditionalAccessDevicesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessDevicesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessDevices(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessDevices) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ConditionalAccessDevices) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDeviceFilter gets the deviceFilter property value. Filter that defines the dynamic-device-syntax rule to include/exclude devices. A filter can use device properties (such as extension attributes) to include/exclude them. Cannot be set if includeDevices or excludeDevices is set.
func (m *ConditionalAccessDevices) GetDeviceFilter()(ConditionalAccessFilterable) {
    val, err := m.GetBackingStore().Get("deviceFilter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ConditionalAccessFilterable)
    }
    return nil
}
// GetExcludeDevices gets the excludeDevices property value. States excluded from the scope of the policy. Possible values: Compliant, DomainJoined. Cannot be set if deviceFIlter is set.
func (m *ConditionalAccessDevices) GetExcludeDevices()([]string) {
    val, err := m.GetBackingStore().Get("excludeDevices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetExcludeDeviceStates gets the excludeDeviceStates property value. The excludeDeviceStates property
func (m *ConditionalAccessDevices) GetExcludeDeviceStates()([]string) {
    val, err := m.GetBackingStore().Get("excludeDeviceStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessDevices) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceFilter(val.(ConditionalAccessFilterable))
        }
        return nil
    }
    res["excludeDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetExcludeDevices(res)
        }
        return nil
    }
    res["excludeDeviceStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetExcludeDeviceStates(res)
        }
        return nil
    }
    res["includeDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetIncludeDevices(res)
        }
        return nil
    }
    res["includeDeviceStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetIncludeDeviceStates(res)
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
    return res
}
// GetIncludeDevices gets the includeDevices property value. States in the scope of the policy. All is the only allowed value. Cannot be set if deviceFIlter is set.
func (m *ConditionalAccessDevices) GetIncludeDevices()([]string) {
    val, err := m.GetBackingStore().Get("includeDevices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetIncludeDeviceStates gets the includeDeviceStates property value. The includeDeviceStates property
func (m *ConditionalAccessDevices) GetIncludeDeviceStates()([]string) {
    val, err := m.GetBackingStore().Get("includeDeviceStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ConditionalAccessDevices) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ConditionalAccessDevices) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceFilter", m.GetDeviceFilter())
        if err != nil {
            return err
        }
    }
    if m.GetExcludeDevices() != nil {
        err := writer.WriteCollectionOfStringValues("excludeDevices", m.GetExcludeDevices())
        if err != nil {
            return err
        }
    }
    if m.GetExcludeDeviceStates() != nil {
        err := writer.WriteCollectionOfStringValues("excludeDeviceStates", m.GetExcludeDeviceStates())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeDevices() != nil {
        err := writer.WriteCollectionOfStringValues("includeDevices", m.GetIncludeDevices())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeDeviceStates() != nil {
        err := writer.WriteCollectionOfStringValues("includeDeviceStates", m.GetIncludeDeviceStates())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessDevices) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ConditionalAccessDevices) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDeviceFilter sets the deviceFilter property value. Filter that defines the dynamic-device-syntax rule to include/exclude devices. A filter can use device properties (such as extension attributes) to include/exclude them. Cannot be set if includeDevices or excludeDevices is set.
func (m *ConditionalAccessDevices) SetDeviceFilter(value ConditionalAccessFilterable)() {
    err := m.GetBackingStore().Set("deviceFilter", value)
    if err != nil {
        panic(err)
    }
}
// SetExcludeDevices sets the excludeDevices property value. States excluded from the scope of the policy. Possible values: Compliant, DomainJoined. Cannot be set if deviceFIlter is set.
func (m *ConditionalAccessDevices) SetExcludeDevices(value []string)() {
    err := m.GetBackingStore().Set("excludeDevices", value)
    if err != nil {
        panic(err)
    }
}
// SetExcludeDeviceStates sets the excludeDeviceStates property value. The excludeDeviceStates property
func (m *ConditionalAccessDevices) SetExcludeDeviceStates(value []string)() {
    err := m.GetBackingStore().Set("excludeDeviceStates", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludeDevices sets the includeDevices property value. States in the scope of the policy. All is the only allowed value. Cannot be set if deviceFIlter is set.
func (m *ConditionalAccessDevices) SetIncludeDevices(value []string)() {
    err := m.GetBackingStore().Set("includeDevices", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludeDeviceStates sets the includeDeviceStates property value. The includeDeviceStates property
func (m *ConditionalAccessDevices) SetIncludeDeviceStates(value []string)() {
    err := m.GetBackingStore().Set("includeDeviceStates", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessDevices) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// ConditionalAccessDevicesable 
type ConditionalAccessDevicesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDeviceFilter()(ConditionalAccessFilterable)
    GetExcludeDevices()([]string)
    GetExcludeDeviceStates()([]string)
    GetIncludeDevices()([]string)
    GetIncludeDeviceStates()([]string)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDeviceFilter(value ConditionalAccessFilterable)()
    SetExcludeDevices(value []string)()
    SetExcludeDeviceStates(value []string)()
    SetIncludeDevices(value []string)()
    SetIncludeDeviceStates(value []string)()
    SetOdataType(value *string)()
}
