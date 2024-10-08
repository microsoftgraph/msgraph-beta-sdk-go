package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ResponseTimeSecurityRequirement struct {
    SecurityRequirement
}
// ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours composed type wrapper for classes float32, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric, string
type ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours instantiates a new ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours and sets the default values.
func NewResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours()(*ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours) {
    m := &ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHoursFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHoursFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetFloat(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours) GetFloat()(*float32) {
    val, err := m.GetBackingStore().Get("float")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours) GetReferenceNumeric()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetFloat() != nil {
        err := writer.WriteFloat32Value("", m.GetFloat())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric
func (m *ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours) SetReferenceNumeric(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHoursable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)()
    SetString(value *string)()
}
// NewResponseTimeSecurityRequirement instantiates a new ResponseTimeSecurityRequirement and sets the default values.
func NewResponseTimeSecurityRequirement()(*ResponseTimeSecurityRequirement) {
    m := &ResponseTimeSecurityRequirement{
        SecurityRequirement: *NewSecurityRequirement(),
    }
    return m
}
// CreateResponseTimeSecurityRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateResponseTimeSecurityRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResponseTimeSecurityRequirement(), nil
}
// GetAverageResponseTimeInHours gets the averageResponseTimeInHours property value. The average response time for alerts from the past 30 days.
// returns a ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHoursable when successful
func (m *ResponseTimeSecurityRequirement) GetAverageResponseTimeInHours()(ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHoursable) {
    val, err := m.GetBackingStore().Get("averageResponseTimeInHours")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHoursable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ResponseTimeSecurityRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SecurityRequirement.GetFieldDeserializers()
    res["averageResponseTimeInHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHoursFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageResponseTimeInHours(val.(*ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHours))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ResponseTimeSecurityRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SecurityRequirement.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("averageResponseTimeInHours", m.GetAverageResponseTimeInHours())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAverageResponseTimeInHours sets the averageResponseTimeInHours property value. The average response time for alerts from the past 30 days.
func (m *ResponseTimeSecurityRequirement) SetAverageResponseTimeInHours(value ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHoursable)() {
    err := m.GetBackingStore().Set("averageResponseTimeInHours", value)
    if err != nil {
        panic(err)
    }
}
type ResponseTimeSecurityRequirementable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SecurityRequirementable
    GetAverageResponseTimeInHours()(ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHoursable)
    SetAverageResponseTimeInHours(value ResponseTimeSecurityRequirement_ResponseTimeSecurityRequirement_averageResponseTimeInHoursable)()
}
