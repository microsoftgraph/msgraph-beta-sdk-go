package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// UserExperienceAnalyticsResourcePerformance the user experience analytics resource performance entity.
type UserExperienceAnalyticsResourcePerformance struct {
    Entity
}
// UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz instantiates a new UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz and sets the default values.
func NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz()(*UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz) {
    m := &UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHzFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHzFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz()
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
    if val, err := parseNode.GetEnumValue(ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz) GetReferenceNumeric()(*ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz) GetString()(*string) {
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage instantiates a new UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage and sets the default values.
func NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage()(*UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage) {
    m := &UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage()
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
    if val, err := parseNode.GetEnumValue(ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage) GetReferenceNumeric()(*ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage) GetString()(*string) {
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold instantiates a new UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold and sets the default values.
func NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold()(*UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold) {
    m := &UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThresholdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThresholdFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold()
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
    if val, err := parseNode.GetEnumValue(ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold) GetReferenceNumeric()(*ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold) GetString()(*string) {
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage instantiates a new UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage and sets the default values.
func NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage()(*UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage) {
    m := &UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage()
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
    if val, err := parseNode.GetEnumValue(ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage) GetReferenceNumeric()(*ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage) GetString()(*string) {
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold instantiates a new UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold and sets the default values.
func NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold()(*UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold) {
    m := &UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThresholdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThresholdFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold()
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
    if val, err := parseNode.GetEnumValue(ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold) GetReferenceNumeric()(*ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold) GetString()(*string) {
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB instantiates a new UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB and sets the default values.
func NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB()(*UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB) {
    m := &UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMBFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMBFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB()
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
    if val, err := parseNode.GetEnumValue(ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB) GetReferenceNumeric()(*ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB) GetString()(*string) {
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
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
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHzable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
type UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
type UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThresholdable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
type UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
type UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThresholdable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
type UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMBable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
// NewUserExperienceAnalyticsResourcePerformance instantiates a new UserExperienceAnalyticsResourcePerformance and sets the default values.
func NewUserExperienceAnalyticsResourcePerformance()(*UserExperienceAnalyticsResourcePerformance) {
    m := &UserExperienceAnalyticsResourcePerformance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsResourcePerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsResourcePerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsResourcePerformance(), nil
}
// GetAverageSpikeTimeScore gets the averageSpikeTimeScore property value. AverageSpikeTimeScore of a device or a model type. Valid values 0 to 100
// returns a *int32 when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetAverageSpikeTimeScore()(*int32) {
    val, err := m.GetBackingStore().Get("averageSpikeTimeScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCpuClockSpeedInMHz gets the cpuClockSpeedInMHz property value. The clock speed of the processor, in MHz. Valid values 0 to 1000000
// returns a UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHzable when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuClockSpeedInMHz()(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHzable) {
    val, err := m.GetBackingStore().Get("cpuClockSpeedInMHz")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHzable)
    }
    return nil
}
// GetCpuDisplayName gets the cpuDisplayName property value. The name of the processor on the device, For example, 11th Gen Intel(R) Core(TM) i7.
// returns a *string when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("cpuDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCpuSpikeTimePercentage gets the cpuSpikeTimePercentage property value. CPU spike time in percentage. Valid values 0 to 100
// returns a UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageable when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuSpikeTimePercentage()(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageable) {
    val, err := m.GetBackingStore().Get("cpuSpikeTimePercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageable)
    }
    return nil
}
// GetCpuSpikeTimePercentageThreshold gets the cpuSpikeTimePercentageThreshold property value. Threshold of cpuSpikeTimeScore. Valid values 0 to 100
// returns a UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThresholdable when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuSpikeTimePercentageThreshold()(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThresholdable) {
    val, err := m.GetBackingStore().Get("cpuSpikeTimePercentageThreshold")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThresholdable)
    }
    return nil
}
// GetCpuSpikeTimeScore gets the cpuSpikeTimeScore property value. The user experience analytics device CPU spike time score. Valid values 0 to 100
// returns a *int32 when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuSpikeTimeScore()(*int32) {
    val, err := m.GetBackingStore().Get("cpuSpikeTimeScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDeviceCount gets the deviceCount property value. User experience analytics summarized device count.
// returns a *int64 when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceCount()(*int64) {
    val, err := m.GetBackingStore().Get("deviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The id of the device.
// returns a *string when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The name of the device.
// returns a *string when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceResourcePerformanceScore gets the deviceResourcePerformanceScore property value. Resource performance score of a specific device. Valid values 0 to 100
// returns a *int32 when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceResourcePerformanceScore()(*int32) {
    val, err := m.GetBackingStore().Get("deviceResourcePerformanceScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDiskType gets the diskType property value. The diskType property
// returns a *DiskType when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetDiskType()(*DiskType) {
    val, err := m.GetBackingStore().Get("diskType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DiskType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["averageSpikeTimeScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageSpikeTimeScore(val)
        }
        return nil
    }
    res["cpuClockSpeedInMHz"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHzFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuClockSpeedInMHz(val.(*UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHz))
        }
        return nil
    }
    res["cpuDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuDisplayName(val)
        }
        return nil
    }
    res["cpuSpikeTimePercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuSpikeTimePercentage(val.(*UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentage))
        }
        return nil
    }
    res["cpuSpikeTimePercentageThreshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThresholdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuSpikeTimePercentageThreshold(val.(*UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThreshold))
        }
        return nil
    }
    res["cpuSpikeTimeScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuSpikeTimeScore(val)
        }
        return nil
    }
    res["deviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["deviceResourcePerformanceScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceResourcePerformanceScore(val)
        }
        return nil
    }
    res["diskType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDiskType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiskType(val.(*DiskType))
        }
        return nil
    }
    res["healthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatus(val.(*UserExperienceAnalyticsHealthState))
        }
        return nil
    }
    res["machineType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsMachineType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMachineType(val.(*UserExperienceAnalyticsMachineType))
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["ramSpikeTimePercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamSpikeTimePercentage(val.(*UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentage))
        }
        return nil
    }
    res["ramSpikeTimePercentageThreshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThresholdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamSpikeTimePercentageThreshold(val.(*UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThreshold))
        }
        return nil
    }
    res["ramSpikeTimeScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamSpikeTimeScore(val)
        }
        return nil
    }
    res["totalProcessorCoreCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalProcessorCoreCount(val)
        }
        return nil
    }
    res["totalRamInMB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMBFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalRamInMB(val.(*UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMB))
        }
        return nil
    }
    return res
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
// returns a *UserExperienceAnalyticsHealthState when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    val, err := m.GetBackingStore().Get("healthStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserExperienceAnalyticsHealthState)
    }
    return nil
}
// GetMachineType gets the machineType property value. Indicates if machine is physical or virtual. Possible values are: physical or virtual
// returns a *UserExperienceAnalyticsMachineType when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetMachineType()(*UserExperienceAnalyticsMachineType) {
    val, err := m.GetBackingStore().Get("machineType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserExperienceAnalyticsMachineType)
    }
    return nil
}
// GetManufacturer gets the manufacturer property value. The user experience analytics device manufacturer.
// returns a *string when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("manufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModel gets the model property value. The user experience analytics device model.
// returns a *string when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRamSpikeTimePercentage gets the ramSpikeTimePercentage property value. RAM spike time in percentage. Valid values 0 to 100
// returns a UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageable when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetRamSpikeTimePercentage()(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageable) {
    val, err := m.GetBackingStore().Get("ramSpikeTimePercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageable)
    }
    return nil
}
// GetRamSpikeTimePercentageThreshold gets the ramSpikeTimePercentageThreshold property value. Threshold of ramSpikeTimeScore. Valid values 0 to 100
// returns a UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThresholdable when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetRamSpikeTimePercentageThreshold()(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThresholdable) {
    val, err := m.GetBackingStore().Get("ramSpikeTimePercentageThreshold")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThresholdable)
    }
    return nil
}
// GetRamSpikeTimeScore gets the ramSpikeTimeScore property value. The user experience analytics device RAM spike time score. Valid values 0 to 100
// returns a *int32 when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetRamSpikeTimeScore()(*int32) {
    val, err := m.GetBackingStore().Get("ramSpikeTimeScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalProcessorCoreCount gets the totalProcessorCoreCount property value. The count of cores of the processor of device. Valid values 0 to 512
// returns a *int32 when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetTotalProcessorCoreCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalProcessorCoreCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalRamInMB gets the totalRamInMB property value. The total RAM of the device, in MB. Valid values 0 to 1000000
// returns a UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMBable when successful
func (m *UserExperienceAnalyticsResourcePerformance) GetTotalRamInMB()(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMBable) {
    val, err := m.GetBackingStore().Get("totalRamInMB")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMBable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsResourcePerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("averageSpikeTimeScore", m.GetAverageSpikeTimeScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("cpuClockSpeedInMHz", m.GetCpuClockSpeedInMHz())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("cpuDisplayName", m.GetCpuDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("cpuSpikeTimePercentage", m.GetCpuSpikeTimePercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("cpuSpikeTimePercentageThreshold", m.GetCpuSpikeTimePercentageThreshold())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("cpuSpikeTimeScore", m.GetCpuSpikeTimeScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("deviceCount", m.GetDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceResourcePerformanceScore", m.GetDeviceResourcePerformanceScore())
        if err != nil {
            return err
        }
    }
    if m.GetDiskType() != nil {
        cast := (*m.GetDiskType()).String()
        err = writer.WriteStringValue("diskType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetHealthStatus() != nil {
        cast := (*m.GetHealthStatus()).String()
        err = writer.WriteStringValue("healthStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMachineType() != nil {
        cast := (*m.GetMachineType()).String()
        err = writer.WriteStringValue("machineType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("ramSpikeTimePercentage", m.GetRamSpikeTimePercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("ramSpikeTimePercentageThreshold", m.GetRamSpikeTimePercentageThreshold())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ramSpikeTimeScore", m.GetRamSpikeTimeScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalProcessorCoreCount", m.GetTotalProcessorCoreCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("totalRamInMB", m.GetTotalRamInMB())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAverageSpikeTimeScore sets the averageSpikeTimeScore property value. AverageSpikeTimeScore of a device or a model type. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetAverageSpikeTimeScore(value *int32)() {
    err := m.GetBackingStore().Set("averageSpikeTimeScore", value)
    if err != nil {
        panic(err)
    }
}
// SetCpuClockSpeedInMHz sets the cpuClockSpeedInMHz property value. The clock speed of the processor, in MHz. Valid values 0 to 1000000
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuClockSpeedInMHz(value UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHzable)() {
    err := m.GetBackingStore().Set("cpuClockSpeedInMHz", value)
    if err != nil {
        panic(err)
    }
}
// SetCpuDisplayName sets the cpuDisplayName property value. The name of the processor on the device, For example, 11th Gen Intel(R) Core(TM) i7.
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuDisplayName(value *string)() {
    err := m.GetBackingStore().Set("cpuDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetCpuSpikeTimePercentage sets the cpuSpikeTimePercentage property value. CPU spike time in percentage. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuSpikeTimePercentage(value UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageable)() {
    err := m.GetBackingStore().Set("cpuSpikeTimePercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetCpuSpikeTimePercentageThreshold sets the cpuSpikeTimePercentageThreshold property value. Threshold of cpuSpikeTimeScore. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuSpikeTimePercentageThreshold(value UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThresholdable)() {
    err := m.GetBackingStore().Set("cpuSpikeTimePercentageThreshold", value)
    if err != nil {
        panic(err)
    }
}
// SetCpuSpikeTimeScore sets the cpuSpikeTimeScore property value. The user experience analytics device CPU spike time score. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuSpikeTimeScore(value *int32)() {
    err := m.GetBackingStore().Set("cpuSpikeTimeScore", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCount sets the deviceCount property value. User experience analytics summarized device count.
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceCount(value *int64)() {
    err := m.GetBackingStore().Set("deviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The name of the device.
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceResourcePerformanceScore sets the deviceResourcePerformanceScore property value. Resource performance score of a specific device. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceResourcePerformanceScore(value *int32)() {
    err := m.GetBackingStore().Set("deviceResourcePerformanceScore", value)
    if err != nil {
        panic(err)
    }
}
// SetDiskType sets the diskType property value. The diskType property
func (m *UserExperienceAnalyticsResourcePerformance) SetDiskType(value *DiskType)() {
    err := m.GetBackingStore().Set("diskType", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsResourcePerformance) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    err := m.GetBackingStore().Set("healthStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetMachineType sets the machineType property value. Indicates if machine is physical or virtual. Possible values are: physical or virtual
func (m *UserExperienceAnalyticsResourcePerformance) SetMachineType(value *UserExperienceAnalyticsMachineType)() {
    err := m.GetBackingStore().Set("machineType", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. The user experience analytics device manufacturer.
func (m *UserExperienceAnalyticsResourcePerformance) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsResourcePerformance) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// SetRamSpikeTimePercentage sets the ramSpikeTimePercentage property value. RAM spike time in percentage. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetRamSpikeTimePercentage(value UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageable)() {
    err := m.GetBackingStore().Set("ramSpikeTimePercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetRamSpikeTimePercentageThreshold sets the ramSpikeTimePercentageThreshold property value. Threshold of ramSpikeTimeScore. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetRamSpikeTimePercentageThreshold(value UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThresholdable)() {
    err := m.GetBackingStore().Set("ramSpikeTimePercentageThreshold", value)
    if err != nil {
        panic(err)
    }
}
// SetRamSpikeTimeScore sets the ramSpikeTimeScore property value. The user experience analytics device RAM spike time score. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetRamSpikeTimeScore(value *int32)() {
    err := m.GetBackingStore().Set("ramSpikeTimeScore", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalProcessorCoreCount sets the totalProcessorCoreCount property value. The count of cores of the processor of device. Valid values 0 to 512
func (m *UserExperienceAnalyticsResourcePerformance) SetTotalProcessorCoreCount(value *int32)() {
    err := m.GetBackingStore().Set("totalProcessorCoreCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalRamInMB sets the totalRamInMB property value. The total RAM of the device, in MB. Valid values 0 to 1000000
func (m *UserExperienceAnalyticsResourcePerformance) SetTotalRamInMB(value UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMBable)() {
    err := m.GetBackingStore().Set("totalRamInMB", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsResourcePerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAverageSpikeTimeScore()(*int32)
    GetCpuClockSpeedInMHz()(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHzable)
    GetCpuDisplayName()(*string)
    GetCpuSpikeTimePercentage()(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageable)
    GetCpuSpikeTimePercentageThreshold()(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThresholdable)
    GetCpuSpikeTimeScore()(*int32)
    GetDeviceCount()(*int64)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetDeviceResourcePerformanceScore()(*int32)
    GetDiskType()(*DiskType)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetMachineType()(*UserExperienceAnalyticsMachineType)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetRamSpikeTimePercentage()(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageable)
    GetRamSpikeTimePercentageThreshold()(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThresholdable)
    GetRamSpikeTimeScore()(*int32)
    GetTotalProcessorCoreCount()(*int32)
    GetTotalRamInMB()(UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMBable)
    SetAverageSpikeTimeScore(value *int32)()
    SetCpuClockSpeedInMHz(value UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuClockSpeedInMHzable)()
    SetCpuDisplayName(value *string)()
    SetCpuSpikeTimePercentage(value UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageable)()
    SetCpuSpikeTimePercentageThreshold(value UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_cpuSpikeTimePercentageThresholdable)()
    SetCpuSpikeTimeScore(value *int32)()
    SetDeviceCount(value *int64)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetDeviceResourcePerformanceScore(value *int32)()
    SetDiskType(value *DiskType)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetMachineType(value *UserExperienceAnalyticsMachineType)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetRamSpikeTimePercentage(value UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageable)()
    SetRamSpikeTimePercentageThreshold(value UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_ramSpikeTimePercentageThresholdable)()
    SetRamSpikeTimeScore(value *int32)()
    SetTotalProcessorCoreCount(value *int32)()
    SetTotalRamInMB(value UserExperienceAnalyticsResourcePerformance_UserExperienceAnalyticsResourcePerformance_totalRamInMBable)()
}
