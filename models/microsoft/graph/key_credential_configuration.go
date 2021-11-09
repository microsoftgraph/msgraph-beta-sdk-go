package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type KeyCredentialConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    maxLifetime *string;
    // Timestamp when the policy is enforced for all apps created on or after the specified date. For existing applications, the enforcement date would be back dated. To apply to all applications regardless of their creation date, this property would be null. Nullable.
    restrictForAppsCreatedAfterDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The type of restriction being applied. Possible values are asymmetricKeyLifetime, unknownFutureValue. Each value of restrictionType can be used only once per policy.
    restrictionType *AppKeyCredentialRestrictionType;
}
// Instantiates a new keyCredentialConfiguration and sets the default values.
func NewKeyCredentialConfiguration()(*KeyCredentialConfiguration) {
    m := &KeyCredentialConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KeyCredentialConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the maxLifetime property value. 
func (m *KeyCredentialConfiguration) GetMaxLifetime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maxLifetime
    }
}
// Gets the restrictForAppsCreatedAfterDateTime property value. Timestamp when the policy is enforced for all apps created on or after the specified date. For existing applications, the enforcement date would be back dated. To apply to all applications regardless of their creation date, this property would be null. Nullable.
func (m *KeyCredentialConfiguration) GetRestrictForAppsCreatedAfterDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.restrictForAppsCreatedAfterDateTime
    }
}
// Gets the restrictionType property value. The type of restriction being applied. Possible values are asymmetricKeyLifetime, unknownFutureValue. Each value of restrictionType can be used only once per policy.
func (m *KeyCredentialConfiguration) GetRestrictionType()(*AppKeyCredentialRestrictionType) {
    if m == nil {
        return nil
    } else {
        return m.restrictionType
    }
}
// The deserialization information for the current model
func (m *KeyCredentialConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["maxLifetime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMaxLifetime(val)
        return nil
    }
    res["restrictForAppsCreatedAfterDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetRestrictForAppsCreatedAfterDateTime(val)
        return nil
    }
    res["restrictionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppKeyCredentialRestrictionType)
        if err != nil {
            return err
        }
        cast := val.(AppKeyCredentialRestrictionType)
        m.SetRestrictionType(&cast)
        return nil
    }
    return res
}
func (m *KeyCredentialConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *KeyCredentialConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("maxLifetime", m.GetMaxLifetime())
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
        cast := m.GetRestrictionType().String()
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *KeyCredentialConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the maxLifetime property value. 
// Parameters:
//  - value : Value to set for the maxLifetime property.
func (m *KeyCredentialConfiguration) SetMaxLifetime(value *string)() {
    m.maxLifetime = value
}
// Sets the restrictForAppsCreatedAfterDateTime property value. Timestamp when the policy is enforced for all apps created on or after the specified date. For existing applications, the enforcement date would be back dated. To apply to all applications regardless of their creation date, this property would be null. Nullable.
// Parameters:
//  - value : Value to set for the restrictForAppsCreatedAfterDateTime property.
func (m *KeyCredentialConfiguration) SetRestrictForAppsCreatedAfterDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.restrictForAppsCreatedAfterDateTime = value
}
// Sets the restrictionType property value. The type of restriction being applied. Possible values are asymmetricKeyLifetime, unknownFutureValue. Each value of restrictionType can be used only once per policy.
// Parameters:
//  - value : Value to set for the restrictionType property.
func (m *KeyCredentialConfiguration) SetRestrictionType(value *AppKeyCredentialRestrictionType)() {
    m.restrictionType = value
}
