package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SecurityBaselineContributingPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of the policy
    displayName *string;
    // Unique identifier of the policy
    sourceId *string;
    // Authoring source of the policy. Possible values are: deviceConfiguration, deviceIntent.
    sourceType *SecurityBaselinePolicySourceType;
}
// Instantiates a new securityBaselineContributingPolicy and sets the default values.
func NewSecurityBaselineContributingPolicy()(*SecurityBaselineContributingPolicy) {
    m := &SecurityBaselineContributingPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityBaselineContributingPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the displayName property value. Name of the policy
func (m *SecurityBaselineContributingPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the sourceId property value. Unique identifier of the policy
func (m *SecurityBaselineContributingPolicy) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
// Gets the sourceType property value. Authoring source of the policy. Possible values are: deviceConfiguration, deviceIntent.
func (m *SecurityBaselineContributingPolicy) GetSourceType()(*SecurityBaselinePolicySourceType) {
    if m == nil {
        return nil
    } else {
        return m.sourceType
    }
}
// The deserialization information for the current model
func (m *SecurityBaselineContributingPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["sourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceId(val)
        }
        return nil
    }
    res["sourceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityBaselinePolicySourceType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(SecurityBaselinePolicySourceType)
            m.SetSourceType(&cast)
        }
        return nil
    }
    return res
}
func (m *SecurityBaselineContributingPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SecurityBaselineContributingPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceId", m.GetSourceId())
        if err != nil {
            return err
        }
    }
    if m.GetSourceType() != nil {
        cast := m.GetSourceType().String()
        err := writer.WriteStringValue("sourceType", &cast)
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
func (m *SecurityBaselineContributingPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the displayName property value. Name of the policy
// Parameters:
//  - value : Value to set for the displayName property.
func (m *SecurityBaselineContributingPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the sourceId property value. Unique identifier of the policy
// Parameters:
//  - value : Value to set for the sourceId property.
func (m *SecurityBaselineContributingPolicy) SetSourceId(value *string)() {
    m.sourceId = value
}
// Sets the sourceType property value. Authoring source of the policy. Possible values are: deviceConfiguration, deviceIntent.
// Parameters:
//  - value : Value to set for the sourceType property.
func (m *SecurityBaselineContributingPolicy) SetSourceType(value *SecurityBaselinePolicySourceType)() {
    m.sourceType = value
}
