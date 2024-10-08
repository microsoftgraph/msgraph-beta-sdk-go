package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ApplicationSignInSummary struct {
    Entity
}
// ApplicationSignInSummary_ApplicationSignInSummary_successPercentage composed type wrapper for classes float64, ReferenceNumeric, string
type ApplicationSignInSummary_ApplicationSignInSummary_successPercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewApplicationSignInSummary_ApplicationSignInSummary_successPercentage instantiates a new ApplicationSignInSummary_ApplicationSignInSummary_successPercentage and sets the default values.
func NewApplicationSignInSummary_ApplicationSignInSummary_successPercentage()(*ApplicationSignInSummary_ApplicationSignInSummary_successPercentage) {
    m := &ApplicationSignInSummary_ApplicationSignInSummary_successPercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateApplicationSignInSummary_ApplicationSignInSummary_successPercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApplicationSignInSummary_ApplicationSignInSummary_successPercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewApplicationSignInSummary_ApplicationSignInSummary_successPercentage()
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
func (m *ApplicationSignInSummary_ApplicationSignInSummary_successPercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *ApplicationSignInSummary_ApplicationSignInSummary_successPercentage) GetDouble()(*float64) {
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
func (m *ApplicationSignInSummary_ApplicationSignInSummary_successPercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ApplicationSignInSummary_ApplicationSignInSummary_successPercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *ApplicationSignInSummary_ApplicationSignInSummary_successPercentage) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *ApplicationSignInSummary_ApplicationSignInSummary_successPercentage) GetString()(*string) {
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
func (m *ApplicationSignInSummary_ApplicationSignInSummary_successPercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ApplicationSignInSummary_ApplicationSignInSummary_successPercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *ApplicationSignInSummary_ApplicationSignInSummary_successPercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *ApplicationSignInSummary_ApplicationSignInSummary_successPercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *ApplicationSignInSummary_ApplicationSignInSummary_successPercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type ApplicationSignInSummary_ApplicationSignInSummary_successPercentageable interface {
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
// NewApplicationSignInSummary instantiates a new ApplicationSignInSummary and sets the default values.
func NewApplicationSignInSummary()(*ApplicationSignInSummary) {
    m := &ApplicationSignInSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateApplicationSignInSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApplicationSignInSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplicationSignInSummary(), nil
}
// GetAppDisplayName gets the appDisplayName property value. Name of the application that the user signed into.
// returns a *string when successful
func (m *ApplicationSignInSummary) GetAppDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("appDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFailedSignInCount gets the failedSignInCount property value. Count of failed sign-ins made by the application.
// returns a *int64 when successful
func (m *ApplicationSignInSummary) GetFailedSignInCount()(*int64) {
    val, err := m.GetBackingStore().Get("failedSignInCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApplicationSignInSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["failedSignInCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedSignInCount(val)
        }
        return nil
    }
    res["successfulSignInCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulSignInCount(val)
        }
        return nil
    }
    res["successPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplicationSignInSummary_ApplicationSignInSummary_successPercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessPercentage(val.(*ApplicationSignInSummary_ApplicationSignInSummary_successPercentage))
        }
        return nil
    }
    return res
}
// GetSuccessfulSignInCount gets the successfulSignInCount property value. Count of successful sign-ins made by the application.
// returns a *int64 when successful
func (m *ApplicationSignInSummary) GetSuccessfulSignInCount()(*int64) {
    val, err := m.GetBackingStore().Get("successfulSignInCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSuccessPercentage gets the successPercentage property value. Percentage of successful sign-ins made by the application.
// returns a ApplicationSignInSummary_ApplicationSignInSummary_successPercentageable when successful
func (m *ApplicationSignInSummary) GetSuccessPercentage()(ApplicationSignInSummary_ApplicationSignInSummary_successPercentageable) {
    val, err := m.GetBackingStore().Get("successPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ApplicationSignInSummary_ApplicationSignInSummary_successPercentageable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ApplicationSignInSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("failedSignInCount", m.GetFailedSignInCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("successfulSignInCount", m.GetSuccessfulSignInCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("successPercentage", m.GetSuccessPercentage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppDisplayName sets the appDisplayName property value. Name of the application that the user signed into.
func (m *ApplicationSignInSummary) SetAppDisplayName(value *string)() {
    err := m.GetBackingStore().Set("appDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetFailedSignInCount sets the failedSignInCount property value. Count of failed sign-ins made by the application.
func (m *ApplicationSignInSummary) SetFailedSignInCount(value *int64)() {
    err := m.GetBackingStore().Set("failedSignInCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessfulSignInCount sets the successfulSignInCount property value. Count of successful sign-ins made by the application.
func (m *ApplicationSignInSummary) SetSuccessfulSignInCount(value *int64)() {
    err := m.GetBackingStore().Set("successfulSignInCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessPercentage sets the successPercentage property value. Percentage of successful sign-ins made by the application.
func (m *ApplicationSignInSummary) SetSuccessPercentage(value ApplicationSignInSummary_ApplicationSignInSummary_successPercentageable)() {
    err := m.GetBackingStore().Set("successPercentage", value)
    if err != nil {
        panic(err)
    }
}
type ApplicationSignInSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppDisplayName()(*string)
    GetFailedSignInCount()(*int64)
    GetSuccessfulSignInCount()(*int64)
    GetSuccessPercentage()(ApplicationSignInSummary_ApplicationSignInSummary_successPercentageable)
    SetAppDisplayName(value *string)()
    SetFailedSignInCount(value *int64)()
    SetSuccessfulSignInCount(value *int64)()
    SetSuccessPercentage(value ApplicationSignInSummary_ApplicationSignInSummary_successPercentageable)()
}
