package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PasswordCredentialConfiguration 
type PasswordCredentialConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The maxLifetime property
    maxLifetime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Enforces the policy for an app created on or after the enforcement date. For existing applications, the enforcement date would be backdated. To apply to all applications, this date would be null.
    restrictForAppsCreatedAfterDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The type of restriction being applied. The possible values are: passwordAddition, passwordLifetime, symmetricKeyAddition, symmetricKeyLifetime,customPasswordAddition, unknownFutureValue. Each value of restrictionType can be used only once per policy.
    restrictionType *AppCredentialRestrictionType
}
// NewPasswordCredentialConfiguration instantiates a new passwordCredentialConfiguration and sets the default values.
func NewPasswordCredentialConfiguration()(*PasswordCredentialConfiguration) {
    m := &PasswordCredentialConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePasswordCredentialConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePasswordCredentialConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPasswordCredentialConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordCredentialConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PasswordCredentialConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["maxLifetime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxLifetime(val)
        }
        return nil
    }
    res["restrictForAppsCreatedAfterDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictForAppsCreatedAfterDateTime(val)
        }
        return nil
    }
    res["restrictionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppCredentialRestrictionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictionType(val.(*AppCredentialRestrictionType))
        }
        return nil
    }
    return res
}
// GetMaxLifetime gets the maxLifetime property value. The maxLifetime property
func (m *PasswordCredentialConfiguration) GetMaxLifetime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.maxLifetime
    }
}
// GetRestrictForAppsCreatedAfterDateTime gets the restrictForAppsCreatedAfterDateTime property value. Enforces the policy for an app created on or after the enforcement date. For existing applications, the enforcement date would be backdated. To apply to all applications, this date would be null.
func (m *PasswordCredentialConfiguration) GetRestrictForAppsCreatedAfterDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.restrictForAppsCreatedAfterDateTime
    }
}
// GetRestrictionType gets the restrictionType property value. The type of restriction being applied. The possible values are: passwordAddition, passwordLifetime, symmetricKeyAddition, symmetricKeyLifetime,customPasswordAddition, unknownFutureValue. Each value of restrictionType can be used only once per policy.
func (m *PasswordCredentialConfiguration) GetRestrictionType()(*AppCredentialRestrictionType) {
    if m == nil {
        return nil
    } else {
        return m.restrictionType
    }
}
// Serialize serializes information the current object
func (m *PasswordCredentialConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteISODurationValue("maxLifetime", m.GetMaxLifetime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("restrictForAppsCreatedAfterDateTime", m.GetRestrictForAppsCreatedAfterDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRestrictionType() != nil {
        cast := (*m.GetRestrictionType()).String()
        err := writer.WriteStringValue("restrictionType", &cast)
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
func (m *PasswordCredentialConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMaxLifetime sets the maxLifetime property value. The maxLifetime property
func (m *PasswordCredentialConfiguration) SetMaxLifetime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.maxLifetime = value
    }
}
// SetRestrictForAppsCreatedAfterDateTime sets the restrictForAppsCreatedAfterDateTime property value. Enforces the policy for an app created on or after the enforcement date. For existing applications, the enforcement date would be backdated. To apply to all applications, this date would be null.
func (m *PasswordCredentialConfiguration) SetRestrictForAppsCreatedAfterDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.restrictForAppsCreatedAfterDateTime = value
    }
}
// SetRestrictionType sets the restrictionType property value. The type of restriction being applied. The possible values are: passwordAddition, passwordLifetime, symmetricKeyAddition, symmetricKeyLifetime,customPasswordAddition, unknownFutureValue. Each value of restrictionType can be used only once per policy.
func (m *PasswordCredentialConfiguration) SetRestrictionType(value *AppCredentialRestrictionType)() {
    if m != nil {
        m.restrictionType = value
    }
}
