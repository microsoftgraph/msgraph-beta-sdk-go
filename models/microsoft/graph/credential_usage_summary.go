package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CredentialUsageSummary 
type CredentialUsageSummary struct {
    Entity
    // Represents the authentication method that the user used. Possible values are:email, mobileSMS, mobileCall, officePhone, securityQuestion (only used for self-service password reset), appNotification, appCode, alternateMobileCall (supported only in registration), fido, appPassword, unknownFutureValue.
    authMethod *UsageAuthMethod;
    // Provides the count of failed resets or registration data.
    failureActivityCount *int64;
    // Defines the feature to report. Possible values are: registration, reset, unknownFutureValue.
    feature *FeatureType;
    // Provides the count of successful registrations or resets.
    successfulActivityCount *int64;
}
// NewCredentialUsageSummary instantiates a new CredentialUsageSummary and sets the default values.
func NewCredentialUsageSummary()(*CredentialUsageSummary) {
    m := &CredentialUsageSummary{
        Entity: *NewEntity(),
    }
    return m
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
// GetSuccessfulActivityCount gets the successfulActivityCount property value. Provides the count of successful registrations or resets.
func (m *CredentialUsageSummary) GetSuccessfulActivityCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.successfulActivityCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CredentialUsageSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUsageAuthMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthMethod(val.(*UsageAuthMethod))
        }
        return nil
    }
    res["failureActivityCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureActivityCount(val)
        }
        return nil
    }
    res["feature"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseFeatureType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeature(val.(*FeatureType))
        }
        return nil
    }
    res["successfulActivityCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *CredentialUsageSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CredentialUsageSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
