package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// UserExperienceAnalyticsRemoteConnection the user experience analytics remote connection entity. The report will be retired on December 31, 2024. You can start using the Cloud PC connection quality report now via https://go.microsoft.com/fwlink/?linkid=2283835.
type UserExperienceAnalyticsRemoteConnection struct {
    Entity
}
// UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage instantiates a new UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage and sets the default values.
func NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage()(*UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage) {
    m := &UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage()
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage) GetString()(*string) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime instantiates a new UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime and sets the default values.
func NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime()(*UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime) {
    m := &UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTimeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTimeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime()
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime) GetString()(*string) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime instantiates a new UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime and sets the default values.
func NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime()(*UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime) {
    m := &UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTimeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTimeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime()
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime) GetString()(*string) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime instantiates a new UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime and sets the default values.
func NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime()(*UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime) {
    m := &UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTimeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTimeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime()
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime) GetString()(*string) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime instantiates a new UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime and sets the default values.
func NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime()(*UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime) {
    m := &UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTimeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTimeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime()
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime) GetString()(*string) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime instantiates a new UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime and sets the default values.
func NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime()(*UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime) {
    m := &UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTimeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTimeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime()
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime) GetString()(*string) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentageable interface {
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
type UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTimeable interface {
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
type UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTimeable interface {
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
type UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTimeable interface {
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
type UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTimeable interface {
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
type UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTimeable interface {
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
// NewUserExperienceAnalyticsRemoteConnection instantiates a new UserExperienceAnalyticsRemoteConnection and sets the default values.
func NewUserExperienceAnalyticsRemoteConnection()(*UserExperienceAnalyticsRemoteConnection) {
    m := &UserExperienceAnalyticsRemoteConnection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsRemoteConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsRemoteConnectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsRemoteConnection(), nil
}
// GetCloudPcFailurePercentage gets the cloudPcFailurePercentage property value. The sign in failure percentage of Cloud PC Device. Valid values 0 to 100
// returns a UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentageable when successful
func (m *UserExperienceAnalyticsRemoteConnection) GetCloudPcFailurePercentage()(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentageable) {
    val, err := m.GetBackingStore().Get("cloudPcFailurePercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentageable)
    }
    return nil
}
// GetCloudPcRoundTripTime gets the cloudPcRoundTripTime property value. The round tip time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTimeable when successful
func (m *UserExperienceAnalyticsRemoteConnection) GetCloudPcRoundTripTime()(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTimeable) {
    val, err := m.GetBackingStore().Get("cloudPcRoundTripTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTimeable)
    }
    return nil
}
// GetCloudPcSignInTime gets the cloudPcSignInTime property value. The sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTimeable when successful
func (m *UserExperienceAnalyticsRemoteConnection) GetCloudPcSignInTime()(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTimeable) {
    val, err := m.GetBackingStore().Get("cloudPcSignInTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTimeable)
    }
    return nil
}
// GetCoreBootTime gets the coreBootTime property value. The core boot time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTimeable when successful
func (m *UserExperienceAnalyticsRemoteConnection) GetCoreBootTime()(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTimeable) {
    val, err := m.GetBackingStore().Get("coreBootTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTimeable)
    }
    return nil
}
// GetCoreSignInTime gets the coreSignInTime property value. The core sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTimeable when successful
func (m *UserExperienceAnalyticsRemoteConnection) GetCoreSignInTime()(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTimeable) {
    val, err := m.GetBackingStore().Get("coreSignInTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTimeable)
    }
    return nil
}
// GetDeviceCount gets the deviceCount property value. The count of remote connection. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("deviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The id of the device.
// returns a *string when successful
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceId()(*string) {
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
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsRemoteConnection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cloudPcFailurePercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcFailurePercentage(val.(*UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentage))
        }
        return nil
    }
    res["cloudPcRoundTripTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTimeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcRoundTripTime(val.(*UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTime))
        }
        return nil
    }
    res["cloudPcSignInTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTimeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcSignInTime(val.(*UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTime))
        }
        return nil
    }
    res["coreBootTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTimeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoreBootTime(val.(*UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTime))
        }
        return nil
    }
    res["coreSignInTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTimeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoreSignInTime(val.(*UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTime))
        }
        return nil
    }
    res["deviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
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
    res["remoteSignInTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTimeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteSignInTime(val.(*UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTime))
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["virtualNetwork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualNetwork(val)
        }
        return nil
    }
    return res
}
// GetManufacturer gets the manufacturer property value. The user experience analytics manufacturer.
// returns a *string when successful
func (m *UserExperienceAnalyticsRemoteConnection) GetManufacturer()(*string) {
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
func (m *UserExperienceAnalyticsRemoteConnection) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemoteSignInTime gets the remoteSignInTime property value. The remote sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTimeable when successful
func (m *UserExperienceAnalyticsRemoteConnection) GetRemoteSignInTime()(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTimeable) {
    val, err := m.GetBackingStore().Get("remoteSignInTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTimeable)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. The user experience analytics userPrincipalName.
// returns a *string when successful
func (m *UserExperienceAnalyticsRemoteConnection) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVirtualNetwork gets the virtualNetwork property value. The user experience analytics virtual network.
// returns a *string when successful
func (m *UserExperienceAnalyticsRemoteConnection) GetVirtualNetwork()(*string) {
    val, err := m.GetBackingStore().Get("virtualNetwork")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsRemoteConnection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("cloudPcFailurePercentage", m.GetCloudPcFailurePercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("cloudPcRoundTripTime", m.GetCloudPcRoundTripTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("cloudPcSignInTime", m.GetCloudPcSignInTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("coreBootTime", m.GetCoreBootTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("coreSignInTime", m.GetCoreSignInTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceCount", m.GetDeviceCount())
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
        err = writer.WriteObjectValue("remoteSignInTime", m.GetRemoteSignInTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("virtualNetwork", m.GetVirtualNetwork())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCloudPcFailurePercentage sets the cloudPcFailurePercentage property value. The sign in failure percentage of Cloud PC Device. Valid values 0 to 100
func (m *UserExperienceAnalyticsRemoteConnection) SetCloudPcFailurePercentage(value UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentageable)() {
    err := m.GetBackingStore().Set("cloudPcFailurePercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPcRoundTripTime sets the cloudPcRoundTripTime property value. The round tip time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetCloudPcRoundTripTime(value UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTimeable)() {
    err := m.GetBackingStore().Set("cloudPcRoundTripTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPcSignInTime sets the cloudPcSignInTime property value. The sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetCloudPcSignInTime(value UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTimeable)() {
    err := m.GetBackingStore().Set("cloudPcSignInTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCoreBootTime sets the coreBootTime property value. The core boot time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetCoreBootTime(value UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTimeable)() {
    err := m.GetBackingStore().Set("coreBootTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCoreSignInTime sets the coreSignInTime property value. The core sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetCoreSignInTime(value UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTimeable)() {
    err := m.GetBackingStore().Set("coreSignInTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCount sets the deviceCount property value. The count of remote connection. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("deviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The name of the device.
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. The user experience analytics manufacturer.
func (m *UserExperienceAnalyticsRemoteConnection) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsRemoteConnection) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoteSignInTime sets the remoteSignInTime property value. The remote sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetRemoteSignInTime(value UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTimeable)() {
    err := m.GetBackingStore().Set("remoteSignInTime", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user experience analytics userPrincipalName.
func (m *UserExperienceAnalyticsRemoteConnection) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetVirtualNetwork sets the virtualNetwork property value. The user experience analytics virtual network.
func (m *UserExperienceAnalyticsRemoteConnection) SetVirtualNetwork(value *string)() {
    err := m.GetBackingStore().Set("virtualNetwork", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsRemoteConnectionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCloudPcFailurePercentage()(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentageable)
    GetCloudPcRoundTripTime()(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTimeable)
    GetCloudPcSignInTime()(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTimeable)
    GetCoreBootTime()(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTimeable)
    GetCoreSignInTime()(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTimeable)
    GetDeviceCount()(*int32)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetRemoteSignInTime()(UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTimeable)
    GetUserPrincipalName()(*string)
    GetVirtualNetwork()(*string)
    SetCloudPcFailurePercentage(value UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcFailurePercentageable)()
    SetCloudPcRoundTripTime(value UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcRoundTripTimeable)()
    SetCloudPcSignInTime(value UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_cloudPcSignInTimeable)()
    SetCoreBootTime(value UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreBootTimeable)()
    SetCoreSignInTime(value UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_coreSignInTimeable)()
    SetDeviceCount(value *int32)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetRemoteSignInTime(value UserExperienceAnalyticsRemoteConnection_UserExperienceAnalyticsRemoteConnection_remoteSignInTimeable)()
    SetUserPrincipalName(value *string)()
    SetVirtualNetwork(value *string)()
}
