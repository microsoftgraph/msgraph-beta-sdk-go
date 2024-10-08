package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// UserExperienceAnalyticsImpactingProcess the user experience analytics top impacting process entity.
type UserExperienceAnalyticsImpactingProcess struct {
    Entity
}
// UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue instantiates a new UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue and sets the default values.
func NewUserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue()(*UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue) {
    m := &UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue()
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
func (m *UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue) GetString()(*string) {
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
func (m *UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValueable interface {
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
// NewUserExperienceAnalyticsImpactingProcess instantiates a new UserExperienceAnalyticsImpactingProcess and sets the default values.
func NewUserExperienceAnalyticsImpactingProcess()(*UserExperienceAnalyticsImpactingProcess) {
    m := &UserExperienceAnalyticsImpactingProcess{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsImpactingProcessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsImpactingProcessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsImpactingProcess(), nil
}
// GetCategory gets the category property value. The category of impacting process.
// returns a *string when successful
func (m *UserExperienceAnalyticsImpactingProcess) GetCategory()(*string) {
    val, err := m.GetBackingStore().Get("category")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. The description of process.
// returns a *string when successful
func (m *UserExperienceAnalyticsImpactingProcess) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The unique identifier of the impacted device.
// returns a *string when successful
func (m *UserExperienceAnalyticsImpactingProcess) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
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
func (m *UserExperienceAnalyticsImpactingProcess) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["impactValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImpactValue(val.(*UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValue))
        }
        return nil
    }
    res["processName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessName(val)
        }
        return nil
    }
    res["publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    return res
}
// GetImpactValue gets the impactValue property value. The impact value of the process. Valid values 0 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValueable when successful
func (m *UserExperienceAnalyticsImpactingProcess) GetImpactValue()(UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValueable) {
    val, err := m.GetBackingStore().Get("impactValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValueable)
    }
    return nil
}
// GetProcessName gets the processName property value. The process name.
// returns a *string when successful
func (m *UserExperienceAnalyticsImpactingProcess) GetProcessName()(*string) {
    val, err := m.GetBackingStore().Get("processName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPublisher gets the publisher property value. The publisher of the process.
// returns a *string when successful
func (m *UserExperienceAnalyticsImpactingProcess) GetPublisher()(*string) {
    val, err := m.GetBackingStore().Get("publisher")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsImpactingProcess) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteObjectValue("impactValue", m.GetImpactValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("processName", m.GetProcessName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategory sets the category property value. The category of impacting process.
func (m *UserExperienceAnalyticsImpactingProcess) SetCategory(value *string)() {
    err := m.GetBackingStore().Set("category", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description of process.
func (m *UserExperienceAnalyticsImpactingProcess) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The unique identifier of the impacted device.
func (m *UserExperienceAnalyticsImpactingProcess) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetImpactValue sets the impactValue property value. The impact value of the process. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsImpactingProcess) SetImpactValue(value UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValueable)() {
    err := m.GetBackingStore().Set("impactValue", value)
    if err != nil {
        panic(err)
    }
}
// SetProcessName sets the processName property value. The process name.
func (m *UserExperienceAnalyticsImpactingProcess) SetProcessName(value *string)() {
    err := m.GetBackingStore().Set("processName", value)
    if err != nil {
        panic(err)
    }
}
// SetPublisher sets the publisher property value. The publisher of the process.
func (m *UserExperienceAnalyticsImpactingProcess) SetPublisher(value *string)() {
    err := m.GetBackingStore().Set("publisher", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsImpactingProcessable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(*string)
    GetDescription()(*string)
    GetDeviceId()(*string)
    GetImpactValue()(UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValueable)
    GetProcessName()(*string)
    GetPublisher()(*string)
    SetCategory(value *string)()
    SetDescription(value *string)()
    SetDeviceId(value *string)()
    SetImpactValue(value UserExperienceAnalyticsImpactingProcess_UserExperienceAnalyticsImpactingProcess_impactValueable)()
    SetProcessName(value *string)()
    SetPublisher(value *string)()
}
