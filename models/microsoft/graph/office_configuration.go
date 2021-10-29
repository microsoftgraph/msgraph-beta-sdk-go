package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new OfficeConfiguration and sets the default values.
func NewOfficeConfiguration()(*OfficeConfiguration) {
    m := &OfficeConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OfficeConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the clientConfigurations property value. List of office Client configuration.
func (m *OfficeConfiguration) GetClientConfigurations()([]OfficeClientConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.clientConfigurations
    }
}
// Gets the tenantCheckinStatuses property value. List of office Client check-in status.
func (m *OfficeConfiguration) GetTenantCheckinStatuses()([]OfficeClientCheckinStatus) {
    if m == nil {
        return nil
    } else {
        return m.tenantCheckinStatuses
    }
}
// Gets the tenantUserCheckinSummary property value. Entity that describes tenant check-in statues
func (m *OfficeConfiguration) GetTenantUserCheckinSummary()(*OfficeUserCheckinSummary) {
    if m == nil {
        return nil
    } else {
        return m.tenantUserCheckinSummary
    }
}
// The deserialization information for the current model
func (m *OfficeConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["clientConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfficeClientConfiguration() })
        if err != nil {
            return err
        }
        res := make([]OfficeClientConfiguration, len(val))
        for i, v := range val {
            res[i] = *(v.(*OfficeClientConfiguration))
        }
        m.SetClientConfigurations(res)
        return nil
    }
    res["tenantCheckinStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfficeClientCheckinStatus() })
        if err != nil {
            return err
        }
        res := make([]OfficeClientCheckinStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*OfficeClientCheckinStatus))
        }
        m.SetTenantCheckinStatuses(res)
        return nil
    }
    res["tenantUserCheckinSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfficeUserCheckinSummary() })
        if err != nil {
            return err
        }
        m.SetTenantUserCheckinSummary(val.(*OfficeUserCheckinSummary))
        return nil
    }
    return res
}
func (m *OfficeConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OfficeConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
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
    {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *OfficeConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the clientConfigurations property value. List of office Client configuration.
// Parameters:
//  - value : Value to set for the clientConfigurations property.
func (m *OfficeConfiguration) SetClientConfigurations(value []OfficeClientConfiguration)() {
    m.clientConfigurations = value
}
// Sets the tenantCheckinStatuses property value. List of office Client check-in status.
// Parameters:
//  - value : Value to set for the tenantCheckinStatuses property.
func (m *OfficeConfiguration) SetTenantCheckinStatuses(value []OfficeClientCheckinStatus)() {
    m.tenantCheckinStatuses = value
}
// Sets the tenantUserCheckinSummary property value. Entity that describes tenant check-in statues
// Parameters:
//  - value : Value to set for the tenantUserCheckinSummary property.
func (m *OfficeConfiguration) SetTenantUserCheckinSummary(value *OfficeUserCheckinSummary)() {
    m.tenantUserCheckinSummary = value
}
