package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// UserExperienceAnalyticsDeviceScopeSummary the user experience analytics tenant level information for all the device scope configurations
type UserExperienceAnalyticsDeviceScopeSummary struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsDeviceScopeSummary instantiates a new userExperienceAnalyticsDeviceScopeSummary and sets the default values.
func NewUserExperienceAnalyticsDeviceScopeSummary()(*UserExperienceAnalyticsDeviceScopeSummary) {
    m := &UserExperienceAnalyticsDeviceScopeSummary{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserExperienceAnalyticsDeviceScopeSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsDeviceScopeSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsDeviceScopeSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsDeviceScopeSummary) GetAdditionalData()(map[string]any) {
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
func (m *UserExperienceAnalyticsDeviceScopeSummary) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCompletedDeviceScopeIds gets the completedDeviceScopeIds property value. A collection of the user experience analytics device scope Unique Identifiers that are enabled and finished recalculating the report metric.
func (m *UserExperienceAnalyticsDeviceScopeSummary) GetCompletedDeviceScopeIds()([]string) {
    val, err := m.GetBackingStore().Get("completedDeviceScopeIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceScopeSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["completedDeviceScopeIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCompletedDeviceScopeIds(res)
        }
        return nil
    }
    res["insufficientDataDeviceScopeIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetInsufficientDataDeviceScopeIds(res)
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
    res["totalDeviceScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalDeviceScopes(val)
        }
        return nil
    }
    res["totalDeviceScopesEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalDeviceScopesEnabled(val)
        }
        return nil
    }
    return res
}
// GetInsufficientDataDeviceScopeIds gets the insufficientDataDeviceScopeIds property value. A collection of user experience analytics device scope Unique Identitfiers that are enabled but there is insufficient data to calculate results.
func (m *UserExperienceAnalyticsDeviceScopeSummary) GetInsufficientDataDeviceScopeIds()([]string) {
    val, err := m.GetBackingStore().Get("insufficientDataDeviceScopeIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsDeviceScopeSummary) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTotalDeviceScopes gets the totalDeviceScopes property value. The total number of user experience analytics device scopes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsDeviceScopeSummary) GetTotalDeviceScopes()(*int32) {
    val, err := m.GetBackingStore().Get("totalDeviceScopes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalDeviceScopesEnabled gets the totalDeviceScopesEnabled property value. The total number of user experience analytics device scopes that are enabled. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsDeviceScopeSummary) GetTotalDeviceScopesEnabled()(*int32) {
    val, err := m.GetBackingStore().Get("totalDeviceScopesEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsDeviceScopeSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCompletedDeviceScopeIds() != nil {
        err := writer.WriteCollectionOfStringValues("completedDeviceScopeIds", m.GetCompletedDeviceScopeIds())
        if err != nil {
            return err
        }
    }
    if m.GetInsufficientDataDeviceScopeIds() != nil {
        err := writer.WriteCollectionOfStringValues("insufficientDataDeviceScopeIds", m.GetInsufficientDataDeviceScopeIds())
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
        err := writer.WriteInt32Value("totalDeviceScopes", m.GetTotalDeviceScopes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalDeviceScopesEnabled", m.GetTotalDeviceScopesEnabled())
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
func (m *UserExperienceAnalyticsDeviceScopeSummary) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *UserExperienceAnalyticsDeviceScopeSummary) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCompletedDeviceScopeIds sets the completedDeviceScopeIds property value. A collection of the user experience analytics device scope Unique Identifiers that are enabled and finished recalculating the report metric.
func (m *UserExperienceAnalyticsDeviceScopeSummary) SetCompletedDeviceScopeIds(value []string)() {
    err := m.GetBackingStore().Set("completedDeviceScopeIds", value)
    if err != nil {
        panic(err)
    }
}
// SetInsufficientDataDeviceScopeIds sets the insufficientDataDeviceScopeIds property value. A collection of user experience analytics device scope Unique Identitfiers that are enabled but there is insufficient data to calculate results.
func (m *UserExperienceAnalyticsDeviceScopeSummary) SetInsufficientDataDeviceScopeIds(value []string)() {
    err := m.GetBackingStore().Set("insufficientDataDeviceScopeIds", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsDeviceScopeSummary) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalDeviceScopes sets the totalDeviceScopes property value. The total number of user experience analytics device scopes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsDeviceScopeSummary) SetTotalDeviceScopes(value *int32)() {
    err := m.GetBackingStore().Set("totalDeviceScopes", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalDeviceScopesEnabled sets the totalDeviceScopesEnabled property value. The total number of user experience analytics device scopes that are enabled. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsDeviceScopeSummary) SetTotalDeviceScopesEnabled(value *int32)() {
    err := m.GetBackingStore().Set("totalDeviceScopesEnabled", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsDeviceScopeSummaryable 
type UserExperienceAnalyticsDeviceScopeSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCompletedDeviceScopeIds()([]string)
    GetInsufficientDataDeviceScopeIds()([]string)
    GetOdataType()(*string)
    GetTotalDeviceScopes()(*int32)
    GetTotalDeviceScopesEnabled()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCompletedDeviceScopeIds(value []string)()
    SetInsufficientDataDeviceScopeIds(value []string)()
    SetOdataType(value *string)()
    SetTotalDeviceScopes(value *int32)()
    SetTotalDeviceScopesEnabled(value *int32)()
}
