package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SecurityBaselineContributingPolicy 
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
// NewSecurityBaselineContributingPolicy instantiates a new securityBaselineContributingPolicy and sets the default values.
func NewSecurityBaselineContributingPolicy()(*SecurityBaselineContributingPolicy) {
    m := &SecurityBaselineContributingPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityBaselineContributingPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. Name of the policy
func (m *SecurityBaselineContributingPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetSourceId gets the sourceId property value. Unique identifier of the policy
func (m *SecurityBaselineContributingPolicy) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
// GetSourceType gets the sourceType property value. Authoring source of the policy. Possible values are: deviceConfiguration, deviceIntent.
func (m *SecurityBaselineContributingPolicy) GetSourceType()(*SecurityBaselinePolicySourceType) {
    if m == nil {
        return nil
    } else {
        return m.sourceType
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
            m.SetSourceType(val.(*SecurityBaselinePolicySourceType))
        }
        return nil
    }
    return res
}
func (m *SecurityBaselineContributingPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetSourceType()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityBaselineContributingPolicy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. Name of the policy
func (m *SecurityBaselineContributingPolicy) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetSourceId sets the sourceId property value. Unique identifier of the policy
func (m *SecurityBaselineContributingPolicy) SetSourceId(value *string)() {
    if m != nil {
        m.sourceId = value
    }
}
// SetSourceType sets the sourceType property value. Authoring source of the policy. Possible values are: deviceConfiguration, deviceIntent.
func (m *SecurityBaselineContributingPolicy) SetSourceType(value *SecurityBaselinePolicySourceType)() {
    if m != nil {
        m.sourceType = value
    }
}
