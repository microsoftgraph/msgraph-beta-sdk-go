package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type CustomClaimConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCustomClaimConfiguration instantiates a new CustomClaimConfiguration and sets the default values.
func NewCustomClaimConfiguration()(*CustomClaimConfiguration) {
    m := &CustomClaimConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomClaimConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomClaimConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomClaimConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CustomClaimConfiguration) GetAdditionalData()(map[string]any) {
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
// GetAttribute gets the attribute property value. The attribute on which we source this property.
// returns a CustomClaimAttributeBaseable when successful
func (m *CustomClaimConfiguration) GetAttribute()(CustomClaimAttributeBaseable) {
    val, err := m.GetBackingStore().Get("attribute")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CustomClaimAttributeBaseable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *CustomClaimConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCondition gets the condition property value. The condition, if any, associated with this configuration.
// returns a CustomClaimConditionBaseable when successful
func (m *CustomClaimConfiguration) GetCondition()(CustomClaimConditionBaseable) {
    val, err := m.GetBackingStore().Get("condition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CustomClaimConditionBaseable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomClaimConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attribute"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomClaimAttributeBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttribute(val.(CustomClaimAttributeBaseable))
        }
        return nil
    }
    res["condition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomClaimConditionBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCondition(val.(CustomClaimConditionBaseable))
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
    res["transformations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomClaimTransformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomClaimTransformationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CustomClaimTransformationable)
                }
            }
            m.SetTransformations(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CustomClaimConfiguration) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTransformations gets the transformations property value. An ordered list of transformations that are applied in sequence.
// returns a []CustomClaimTransformationable when successful
func (m *CustomClaimConfiguration) GetTransformations()([]CustomClaimTransformationable) {
    val, err := m.GetBackingStore().Get("transformations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomClaimTransformationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomClaimConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attribute", m.GetAttribute())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("condition", m.GetCondition())
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
    if m.GetTransformations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTransformations()))
        for i, v := range m.GetTransformations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("transformations", cast)
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
func (m *CustomClaimConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAttribute sets the attribute property value. The attribute on which we source this property.
func (m *CustomClaimConfiguration) SetAttribute(value CustomClaimAttributeBaseable)() {
    err := m.GetBackingStore().Set("attribute", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CustomClaimConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCondition sets the condition property value. The condition, if any, associated with this configuration.
func (m *CustomClaimConfiguration) SetCondition(value CustomClaimConditionBaseable)() {
    err := m.GetBackingStore().Set("condition", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CustomClaimConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTransformations sets the transformations property value. An ordered list of transformations that are applied in sequence.
func (m *CustomClaimConfiguration) SetTransformations(value []CustomClaimTransformationable)() {
    err := m.GetBackingStore().Set("transformations", value)
    if err != nil {
        panic(err)
    }
}
type CustomClaimConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttribute()(CustomClaimAttributeBaseable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCondition()(CustomClaimConditionBaseable)
    GetOdataType()(*string)
    GetTransformations()([]CustomClaimTransformationable)
    SetAttribute(value CustomClaimAttributeBaseable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCondition(value CustomClaimConditionBaseable)()
    SetOdataType(value *string)()
    SetTransformations(value []CustomClaimTransformationable)()
}
