package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CredentialUsageSummary 
type CredentialUsageSummary struct {
    Entity
    // Represents the authentication method that the user used. Possible values are:email, mobileSMS, mobileCall, officePhone, securityQuestion (only used for self-service password reset), appNotification, appCode, alternateMobileCall (supported only in registration), fido, appPassword, unknownFutureValue.
    authMethod *UsageAuthMethod
    // Provides the count of failed resets or registration data.
    failureActivityCount *int64
    // Defines the feature to report. Possible values are: registration, reset, unknownFutureValue.
    feature *FeatureType
    // Provides the count of successful registrations or resets.
    successfulActivityCount *int64
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
// GetAuthMethod gets the authMethod property value. Represents the authentication method that the user used. Possible values are:email, mobileSMS, mobileCall, officePhone, securityQuestion (only used for self-service password reset), appNotification, appCode, alternateMobileCall (supported only in registration), fido, appPassword, unknownFutureValue.
func (m *CredentialUsageSummary) GetAuthMethod()(*UsageAuthMethod) {
    if m == nil {
        return nil
    } else {
        return m.authMethod
    }
}
// GetFailureActivityCount gets the failureActivityCount property value. Provides the count of failed resets or registration data.
func (m *CredentialUsageSummary) GetFailureActivityCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.failureActivityCount
    }
}
// GetFeature gets the feature property value. Defines the feature to report. Possible values are: registration, reset, unknownFutureValue.
func (m *CredentialUsageSummary) GetFeature()(*FeatureType) {
    if m == nil {
        return nil
    } else {
        return m.feature
    }
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
    if m == nil {
        return nil
    } else {
        return m.successfulActivityCount
    }
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
// SetAuthMethod sets the authMethod property value. Represents the authentication method that the user used. Possible values are:email, mobileSMS, mobileCall, officePhone, securityQuestion (only used for self-service password reset), appNotification, appCode, alternateMobileCall (supported only in registration), fido, appPassword, unknownFutureValue.
func (m *CredentialUsageSummary) SetAuthMethod(value *UsageAuthMethod)() {
    if m != nil {
        m.authMethod = value
    }
}
// SetFailureActivityCount sets the failureActivityCount property value. Provides the count of failed resets or registration data.
func (m *CredentialUsageSummary) SetFailureActivityCount(value *int64)() {
    if m != nil {
        m.failureActivityCount = value
    }
}
// SetFeature sets the feature property value. Defines the feature to report. Possible values are: registration, reset, unknownFutureValue.
func (m *CredentialUsageSummary) SetFeature(value *FeatureType)() {
    if m != nil {
        m.feature = value
    }
}
// SetSuccessfulActivityCount sets the successfulActivityCount property value. Provides the count of successful registrations or resets.
func (m *CredentialUsageSummary) SetSuccessfulActivityCount(value *int64)() {
    if m != nil {
        m.successfulActivityCount = value
    }
}
