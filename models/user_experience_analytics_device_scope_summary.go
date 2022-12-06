package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceScopeSummary the user experience analytics tenant level information for all the device scope configurations
type UserExperienceAnalyticsDeviceScopeSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A collection of the user experience analytics device scope Unique Identifiers that are enabled and finished recalculating the report metric.
    completedDeviceScopeIds []string
    // A collection of user experience analytics device scope Unique Identitfiers that are enabled but there is insufficient data to calculate results.
    insufficientDataDeviceScopeIds []string
    // The OdataType property
    odataType *string
    // The total number of user experience analytics device scopes. Valid values -2147483648 to 2147483647
    totalDeviceScopes *int32
    // The total number of user experience analytics device scopes that are enabled. Valid values -2147483648 to 2147483647
    totalDeviceScopesEnabled *int32
}
// NewUserExperienceAnalyticsDeviceScopeSummary instantiates a new userExperienceAnalyticsDeviceScopeSummary and sets the default values.
func NewUserExperienceAnalyticsDeviceScopeSummary()(*UserExperienceAnalyticsDeviceScopeSummary) {
    m := &UserExperienceAnalyticsDeviceScopeSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserExperienceAnalyticsDeviceScopeSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsDeviceScopeSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsDeviceScopeSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsDeviceScopeSummary) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCompletedDeviceScopeIds gets the completedDeviceScopeIds property value. A collection of the user experience analytics device scope Unique Identifiers that are enabled and finished recalculating the report metric.
func (m *UserExperienceAnalyticsDeviceScopeSummary) GetCompletedDeviceScopeIds()([]string) {
    return m.completedDeviceScopeIds
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceScopeSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["completedDeviceScopeIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetCompletedDeviceScopeIds)
    res["insufficientDataDeviceScopeIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetInsufficientDataDeviceScopeIds)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["totalDeviceScopes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalDeviceScopes)
    res["totalDeviceScopesEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalDeviceScopesEnabled)
    return res
}
// GetInsufficientDataDeviceScopeIds gets the insufficientDataDeviceScopeIds property value. A collection of user experience analytics device scope Unique Identitfiers that are enabled but there is insufficient data to calculate results.
func (m *UserExperienceAnalyticsDeviceScopeSummary) GetInsufficientDataDeviceScopeIds()([]string) {
    return m.insufficientDataDeviceScopeIds
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsDeviceScopeSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetTotalDeviceScopes gets the totalDeviceScopes property value. The total number of user experience analytics device scopes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsDeviceScopeSummary) GetTotalDeviceScopes()(*int32) {
    return m.totalDeviceScopes
}
// GetTotalDeviceScopesEnabled gets the totalDeviceScopesEnabled property value. The total number of user experience analytics device scopes that are enabled. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsDeviceScopeSummary) GetTotalDeviceScopesEnabled()(*int32) {
    return m.totalDeviceScopesEnabled
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
func (m *UserExperienceAnalyticsDeviceScopeSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCompletedDeviceScopeIds sets the completedDeviceScopeIds property value. A collection of the user experience analytics device scope Unique Identifiers that are enabled and finished recalculating the report metric.
func (m *UserExperienceAnalyticsDeviceScopeSummary) SetCompletedDeviceScopeIds(value []string)() {
    m.completedDeviceScopeIds = value
}
// SetInsufficientDataDeviceScopeIds sets the insufficientDataDeviceScopeIds property value. A collection of user experience analytics device scope Unique Identitfiers that are enabled but there is insufficient data to calculate results.
func (m *UserExperienceAnalyticsDeviceScopeSummary) SetInsufficientDataDeviceScopeIds(value []string)() {
    m.insufficientDataDeviceScopeIds = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsDeviceScopeSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTotalDeviceScopes sets the totalDeviceScopes property value. The total number of user experience analytics device scopes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsDeviceScopeSummary) SetTotalDeviceScopes(value *int32)() {
    m.totalDeviceScopes = value
}
// SetTotalDeviceScopesEnabled sets the totalDeviceScopesEnabled property value. The total number of user experience analytics device scopes that are enabled. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsDeviceScopeSummary) SetTotalDeviceScopesEnabled(value *int32)() {
    m.totalDeviceScopesEnabled = value
}
