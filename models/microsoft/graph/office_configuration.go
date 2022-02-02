package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OfficeConfiguration 
type OfficeConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // List of office Client configuration.
    clientConfigurations []OfficeClientConfiguration;
    // List of office Client check-in status.
    tenantCheckinStatuses []OfficeClientCheckinStatus;
    // Entity that describes tenant check-in statues
    tenantUserCheckinSummary *OfficeUserCheckinSummary;
}
// NewOfficeConfiguration instantiates a new OfficeConfiguration and sets the default values.
func NewOfficeConfiguration()(*OfficeConfiguration) {
    m := &OfficeConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OfficeConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClientConfigurations gets the clientConfigurations property value. List of office Client configuration.
func (m *OfficeConfiguration) GetClientConfigurations()([]OfficeClientConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.clientConfigurations
    }
}
// GetTenantCheckinStatuses gets the tenantCheckinStatuses property value. List of office Client check-in status.
func (m *OfficeConfiguration) GetTenantCheckinStatuses()([]OfficeClientCheckinStatus) {
    if m == nil {
        return nil
    } else {
        return m.tenantCheckinStatuses
    }
}
// GetTenantUserCheckinSummary gets the tenantUserCheckinSummary property value. Entity that describes tenant check-in statues
func (m *OfficeConfiguration) GetTenantUserCheckinSummary()(*OfficeUserCheckinSummary) {
    if m == nil {
        return nil
    } else {
        return m.tenantUserCheckinSummary
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["clientConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfficeClientConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OfficeClientConfiguration, len(val))
            for i, v := range val {
                res[i] = *(v.(*OfficeClientConfiguration))
            }
            m.SetClientConfigurations(res)
        }
        return nil
    }
    res["tenantCheckinStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfficeClientCheckinStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OfficeClientCheckinStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*OfficeClientCheckinStatus))
            }
            m.SetTenantCheckinStatuses(res)
        }
        return nil
    }
    res["tenantUserCheckinSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfficeUserCheckinSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantUserCheckinSummary(val.(*OfficeUserCheckinSummary))
        }
        return nil
    }
    return res
}
func (m *OfficeConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OfficeConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetClientConfigurations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClientConfigurations()))
        for i, v := range m.GetClientConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("clientConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantCheckinStatuses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTenantCheckinStatuses()))
        for i, v := range m.GetTenantCheckinStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("tenantCheckinStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("tenantUserCheckinSummary", m.GetTenantUserCheckinSummary())
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
func (m *OfficeConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetClientConfigurations sets the clientConfigurations property value. List of office Client configuration.
func (m *OfficeConfiguration) SetClientConfigurations(value []OfficeClientConfiguration)() {
    if m != nil {
        m.clientConfigurations = value
    }
}
// SetTenantCheckinStatuses sets the tenantCheckinStatuses property value. List of office Client check-in status.
func (m *OfficeConfiguration) SetTenantCheckinStatuses(value []OfficeClientCheckinStatus)() {
    if m != nil {
        m.tenantCheckinStatuses = value
    }
}
// SetTenantUserCheckinSummary sets the tenantUserCheckinSummary property value. Entity that describes tenant check-in statues
func (m *OfficeConfiguration) SetTenantUserCheckinSummary(value *OfficeUserCheckinSummary)() {
    if m != nil {
        m.tenantUserCheckinSummary = value
    }
}
