package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CredentialUsageSummary 
type CredentialUsageSummary struct {
    Entity
}
// NewCredentialUsageSummary instantiates a new CredentialUsageSummary and sets the default values.
func NewCredentialUsageSummary()(*CredentialUsageSummary) {
    m := &CredentialUsageSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCredentialUsageSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCredentialUsageSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCredentialUsageSummary(), nil
}
// GetAuthMethod gets the authMethod property value. The authMethod property
func (m *CredentialUsageSummary) GetAuthMethod()(*UsageAuthMethod) {
    val, err := m.GetBackingStore().Get("authMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UsageAuthMethod)
    }
    return nil
}
// GetFailureActivityCount gets the failureActivityCount property value. Provides the count of failed resets or registration data.
func (m *CredentialUsageSummary) GetFailureActivityCount()(*int64) {
    val, err := m.GetBackingStore().Get("failureActivityCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFeature gets the feature property value. The feature property
func (m *CredentialUsageSummary) GetFeature()(*FeatureType) {
    val, err := m.GetBackingStore().Get("feature")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*FeatureType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CredentialUsageSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUsageAuthMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthMethod(val.(*UsageAuthMethod))
        }
        return nil
    }
    res["failureActivityCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureActivityCount(val)
        }
        return nil
    }
    res["feature"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFeatureType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeature(val.(*FeatureType))
        }
        return nil
    }
    res["successfulActivityCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulActivityCount(val)
        }
        return nil
    }
    return res
}
// GetSuccessfulActivityCount gets the successfulActivityCount property value. Provides the count of successful registrations or resets.
func (m *CredentialUsageSummary) GetSuccessfulActivityCount()(*int64) {
    val, err := m.GetBackingStore().Get("successfulActivityCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CredentialUsageSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthMethod() != nil {
        cast := (*m.GetAuthMethod()).String()
        err = writer.WriteStringValue("authMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("failureActivityCount", m.GetFailureActivityCount())
        if err != nil {
            return err
        }
    }
    if m.GetFeature() != nil {
        cast := (*m.GetFeature()).String()
        err = writer.WriteStringValue("feature", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("successfulActivityCount", m.GetSuccessfulActivityCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthMethod sets the authMethod property value. The authMethod property
func (m *CredentialUsageSummary) SetAuthMethod(value *UsageAuthMethod)() {
    err := m.GetBackingStore().Set("authMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetFailureActivityCount sets the failureActivityCount property value. Provides the count of failed resets or registration data.
func (m *CredentialUsageSummary) SetFailureActivityCount(value *int64)() {
    err := m.GetBackingStore().Set("failureActivityCount", value)
    if err != nil {
        panic(err)
    }
}
// SetFeature sets the feature property value. The feature property
func (m *CredentialUsageSummary) SetFeature(value *FeatureType)() {
    err := m.GetBackingStore().Set("feature", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessfulActivityCount sets the successfulActivityCount property value. Provides the count of successful registrations or resets.
func (m *CredentialUsageSummary) SetSuccessfulActivityCount(value *int64)() {
    err := m.GetBackingStore().Set("successfulActivityCount", value)
    if err != nil {
        panic(err)
    }
}
// CredentialUsageSummaryable 
type CredentialUsageSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthMethod()(*UsageAuthMethod)
    GetFailureActivityCount()(*int64)
    GetFeature()(*FeatureType)
    GetSuccessfulActivityCount()(*int64)
    SetAuthMethod(value *UsageAuthMethod)()
    SetFailureActivityCount(value *int64)()
    SetFeature(value *FeatureType)()
    SetSuccessfulActivityCount(value *int64)()
}
