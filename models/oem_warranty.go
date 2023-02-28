package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// OemWarranty oEM Warranty information for a given device
type OemWarranty struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOemWarranty instantiates a new oemWarranty and sets the default values.
func NewOemWarranty()(*OemWarranty) {
    m := &OemWarranty{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOemWarrantyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOemWarrantyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOemWarranty(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OemWarranty) GetAdditionalData()(map[string]any) {
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
// GetAdditionalWarranties gets the additionalWarranties property value. List of additional warranty offers. This collection can contain a maximum of 100 elements.
func (m *OemWarranty) GetAdditionalWarranties()([]WarrantyOfferable) {
    val, err := m.GetBackingStore().Get("additionalWarranties")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WarrantyOfferable)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *OemWarranty) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBaseWarranties gets the baseWarranties property value. List of base warranty offers. This collection can contain a maximum of 100 elements.
func (m *OemWarranty) GetBaseWarranties()([]WarrantyOfferable) {
    val, err := m.GetBackingStore().Get("baseWarranties")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WarrantyOfferable)
    }
    return nil
}
// GetDeviceConfigurationUrl gets the deviceConfigurationUrl property value. Device configuration page URL
func (m *OemWarranty) GetDeviceConfigurationUrl()(*string) {
    val, err := m.GetBackingStore().Get("deviceConfigurationUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceWarrantyUrl gets the deviceWarrantyUrl property value. Device warranty page URL
func (m *OemWarranty) GetDeviceWarrantyUrl()(*string) {
    val, err := m.GetBackingStore().Get("deviceWarrantyUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OemWarranty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["additionalWarranties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWarrantyOfferFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WarrantyOfferable, len(val))
            for i, v := range val {
                res[i] = v.(WarrantyOfferable)
            }
            m.SetAdditionalWarranties(res)
        }
        return nil
    }
    res["baseWarranties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWarrantyOfferFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WarrantyOfferable, len(val))
            for i, v := range val {
                res[i] = v.(WarrantyOfferable)
            }
            m.SetBaseWarranties(res)
        }
        return nil
    }
    res["deviceConfigurationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfigurationUrl(val)
        }
        return nil
    }
    res["deviceWarrantyUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceWarrantyUrl(val)
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OemWarranty) GetOdataType()(*string) {
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
func (m *OemWarranty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAdditionalWarranties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdditionalWarranties()))
        for i, v := range m.GetAdditionalWarranties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("additionalWarranties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBaseWarranties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBaseWarranties()))
        for i, v := range m.GetBaseWarranties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("baseWarranties", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceConfigurationUrl", m.GetDeviceConfigurationUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceWarrantyUrl", m.GetDeviceWarrantyUrl())
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
func (m *OemWarranty) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalWarranties sets the additionalWarranties property value. List of additional warranty offers. This collection can contain a maximum of 100 elements.
func (m *OemWarranty) SetAdditionalWarranties(value []WarrantyOfferable)() {
    err := m.GetBackingStore().Set("additionalWarranties", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *OemWarranty) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBaseWarranties sets the baseWarranties property value. List of base warranty offers. This collection can contain a maximum of 100 elements.
func (m *OemWarranty) SetBaseWarranties(value []WarrantyOfferable)() {
    err := m.GetBackingStore().Set("baseWarranties", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceConfigurationUrl sets the deviceConfigurationUrl property value. Device configuration page URL
func (m *OemWarranty) SetDeviceConfigurationUrl(value *string)() {
    err := m.GetBackingStore().Set("deviceConfigurationUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceWarrantyUrl sets the deviceWarrantyUrl property value. Device warranty page URL
func (m *OemWarranty) SetDeviceWarrantyUrl(value *string)() {
    err := m.GetBackingStore().Set("deviceWarrantyUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OemWarranty) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// OemWarrantyable 
type OemWarrantyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalWarranties()([]WarrantyOfferable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBaseWarranties()([]WarrantyOfferable)
    GetDeviceConfigurationUrl()(*string)
    GetDeviceWarrantyUrl()(*string)
    GetOdataType()(*string)
    SetAdditionalWarranties(value []WarrantyOfferable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBaseWarranties(value []WarrantyOfferable)()
    SetDeviceConfigurationUrl(value *string)()
    SetDeviceWarrantyUrl(value *string)()
    SetOdataType(value *string)()
}
