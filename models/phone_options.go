package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type PhoneOptions struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPhoneOptions instantiates a new PhoneOptions and sets the default values.
func NewPhoneOptions()(*PhoneOptions) {
    m := &PhoneOptions{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePhoneOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePhoneOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPhoneOptions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PhoneOptions) GetAdditionalData()(map[string]any) {
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
func (m *PhoneOptions) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDefaultRegions gets the defaultRegions property value. A read-only, Microsoft-defined list of regions that already enable MFA. For more information, see the following list of countries.
// returns a []int32 when successful
func (m *PhoneOptions) GetDefaultRegions()([]int32) {
    val, err := m.GetBackingStore().Get("defaultRegions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]int32)
    }
    return nil
}
// GetExcludeRegions gets the excludeRegions property value. A numbers-only set representing the region telecom codes to prevent or disable the telephony service. Validates against current International Subscriber Dialing (ISD) country codes where the maximum code length is 4. Values must be non-null.
// returns a []int32 when successful
func (m *PhoneOptions) GetExcludeRegions()([]int32) {
    val, err := m.GetBackingStore().Get("excludeRegions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PhoneOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultRegions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetDefaultRegions(res)
        }
        return nil
    }
    res["excludeRegions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetExcludeRegions(res)
        }
        return nil
    }
    res["includeAdditionalRegions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetIncludeAdditionalRegions(res)
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
// GetIncludeAdditionalRegions gets the includeAdditionalRegions property value. A numbers-only set representing the country codes that can be manually added to enable telephony service in those regions, in addition to the list of countries that are already enabled. For more information about regions that require opt in, see Regions that need to opt in for MFA telephony verification. Validates against current International Subscriber Dialing (ISD) country codes where the maximum code length is 4. Values must be positive integers and can't overlap with 'excludeRegions'.
// returns a []int32 when successful
func (m *PhoneOptions) GetIncludeAdditionalRegions()([]int32) {
    val, err := m.GetBackingStore().Get("includeAdditionalRegions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *PhoneOptions) GetOdataType()(*string) {
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
func (m *PhoneOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDefaultRegions() != nil {
        err := writer.WriteCollectionOfInt32Values("defaultRegions", m.GetDefaultRegions())
        if err != nil {
            return err
        }
    }
    if m.GetExcludeRegions() != nil {
        err := writer.WriteCollectionOfInt32Values("excludeRegions", m.GetExcludeRegions())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeAdditionalRegions() != nil {
        err := writer.WriteCollectionOfInt32Values("includeAdditionalRegions", m.GetIncludeAdditionalRegions())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PhoneOptions) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *PhoneOptions) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDefaultRegions sets the defaultRegions property value. A read-only, Microsoft-defined list of regions that already enable MFA. For more information, see the following list of countries.
func (m *PhoneOptions) SetDefaultRegions(value []int32)() {
    err := m.GetBackingStore().Set("defaultRegions", value)
    if err != nil {
        panic(err)
    }
}
// SetExcludeRegions sets the excludeRegions property value. A numbers-only set representing the region telecom codes to prevent or disable the telephony service. Validates against current International Subscriber Dialing (ISD) country codes where the maximum code length is 4. Values must be non-null.
func (m *PhoneOptions) SetExcludeRegions(value []int32)() {
    err := m.GetBackingStore().Set("excludeRegions", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludeAdditionalRegions sets the includeAdditionalRegions property value. A numbers-only set representing the country codes that can be manually added to enable telephony service in those regions, in addition to the list of countries that are already enabled. For more information about regions that require opt in, see Regions that need to opt in for MFA telephony verification. Validates against current International Subscriber Dialing (ISD) country codes where the maximum code length is 4. Values must be positive integers and can't overlap with 'excludeRegions'.
func (m *PhoneOptions) SetIncludeAdditionalRegions(value []int32)() {
    err := m.GetBackingStore().Set("includeAdditionalRegions", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PhoneOptions) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type PhoneOptionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDefaultRegions()([]int32)
    GetExcludeRegions()([]int32)
    GetIncludeAdditionalRegions()([]int32)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDefaultRegions(value []int32)()
    SetExcludeRegions(value []int32)()
    SetIncludeAdditionalRegions(value []int32)()
    SetOdataType(value *string)()
}
