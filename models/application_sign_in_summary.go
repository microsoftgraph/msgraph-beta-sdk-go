package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplicationSignInSummary 
type ApplicationSignInSummary struct {
    Entity
    // Name of the application that the user signed into.
    appDisplayName *string
    // Count of failed sign-ins made by the application.
    failedSignInCount *int64
    // Count of successful sign-ins made by the application.
    successfulSignInCount *int64
    // Percentage of successful sign-ins made by the application.
    successPercentage *float64
}
// NewApplicationSignInSummary instantiates a new ApplicationSignInSummary and sets the default values.
func NewApplicationSignInSummary()(*ApplicationSignInSummary) {
    m := &ApplicationSignInSummary{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.applicationSignInSummary";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateApplicationSignInSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicationSignInSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplicationSignInSummary(), nil
}
// GetAppDisplayName gets the appDisplayName property value. Name of the application that the user signed into.
func (m *ApplicationSignInSummary) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetFailedSignInCount gets the failedSignInCount property value. Count of failed sign-ins made by the application.
func (m *ApplicationSignInSummary) GetFailedSignInCount()(*int64) {
    return m.failedSignInCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicationSignInSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppDisplayName)
    res["failedSignInCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetFailedSignInCount)
    res["successfulSignInCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetSuccessfulSignInCount)
    res["successPercentage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetSuccessPercentage)
    return res
}
// GetSuccessfulSignInCount gets the successfulSignInCount property value. Count of successful sign-ins made by the application.
func (m *ApplicationSignInSummary) GetSuccessfulSignInCount()(*int64) {
    return m.successfulSignInCount
}
// GetSuccessPercentage gets the successPercentage property value. Percentage of successful sign-ins made by the application.
func (m *ApplicationSignInSummary) GetSuccessPercentage()(*float64) {
    return m.successPercentage
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
        err = writer.WriteFloat64Value("successPercentage", m.GetSuccessPercentage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppDisplayName sets the appDisplayName property value. Name of the application that the user signed into.
func (m *ApplicationSignInSummary) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetFailedSignInCount sets the failedSignInCount property value. Count of failed sign-ins made by the application.
func (m *ApplicationSignInSummary) SetFailedSignInCount(value *int64)() {
    m.failedSignInCount = value
}
// SetSuccessfulSignInCount sets the successfulSignInCount property value. Count of successful sign-ins made by the application.
func (m *ApplicationSignInSummary) SetSuccessfulSignInCount(value *int64)() {
    m.successfulSignInCount = value
}
// SetSuccessPercentage sets the successPercentage property value. Percentage of successful sign-ins made by the application.
func (m *ApplicationSignInSummary) SetSuccessPercentage(value *float64)() {
    m.successPercentage = value
}
