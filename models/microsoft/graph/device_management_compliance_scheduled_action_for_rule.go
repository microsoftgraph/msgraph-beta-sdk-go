package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementComplianceScheduledActionForRule provides operations to manage the deviceManagement singleton.
type DeviceManagementComplianceScheduledActionForRule struct {
    Entity
    // Name of the rule which this scheduled action applies to.
    ruleName *string;
    // The list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
    scheduledActionConfigurations []DeviceManagementComplianceActionItemable;
}
// NewDeviceManagementComplianceScheduledActionForRule instantiates a new deviceManagementComplianceScheduledActionForRule and sets the default values.
func NewDeviceManagementComplianceScheduledActionForRule()(*DeviceManagementComplianceScheduledActionForRule) {
    m := &DeviceManagementComplianceScheduledActionForRule{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementComplianceScheduledActionForRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementComplianceScheduledActionForRuleFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceManagementComplianceScheduledActionForRule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementComplianceScheduledActionForRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["ruleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleName(val)
        }
        return nil
    }
    res["scheduledActionConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementComplianceActionItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementComplianceActionItemable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementComplianceActionItemable)
            }
            m.SetScheduledActionConfigurations(res)
        }
        return nil
    }
    return res
}
// GetRuleName gets the ruleName property value. Name of the rule which this scheduled action applies to.
func (m *DeviceManagementComplianceScheduledActionForRule) GetRuleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ruleName
    }
}
// GetScheduledActionConfigurations gets the scheduledActionConfigurations property value. The list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
func (m *DeviceManagementComplianceScheduledActionForRule) GetScheduledActionConfigurations()([]DeviceManagementComplianceActionItemable) {
    if m == nil {
        return nil
    } else {
        return m.scheduledActionConfigurations
    }
}
func (m *DeviceManagementComplianceScheduledActionForRule) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementComplianceScheduledActionForRule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("ruleName", m.GetRuleName())
        if err != nil {
            return err
        }
    }
    if m.GetScheduledActionConfigurations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetScheduledActionConfigurations()))
        for i, v := range m.GetScheduledActionConfigurations() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("scheduledActionConfigurations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRuleName sets the ruleName property value. Name of the rule which this scheduled action applies to.
func (m *DeviceManagementComplianceScheduledActionForRule) SetRuleName(value *string)() {
    if m != nil {
        m.ruleName = value
    }
}
// SetScheduledActionConfigurations sets the scheduledActionConfigurations property value. The list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
func (m *DeviceManagementComplianceScheduledActionForRule) SetScheduledActionConfigurations(value []DeviceManagementComplianceActionItemable)() {
    if m != nil {
        m.scheduledActionConfigurations = value
    }
}
