package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ServicePrincipalCreationConditionSet struct {
    Entity
    // 
    applicationIds []string;
    // 
    applicationPublisherIds []string;
    // 
    applicationsFromVerifiedPublisherOnly *bool;
    // 
    applicationTenantIds []string;
}
// Instantiates a new servicePrincipalCreationConditionSet and sets the default values.
func NewServicePrincipalCreationConditionSet()(*ServicePrincipalCreationConditionSet) {
    m := &ServicePrincipalCreationConditionSet{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the applicationIds property value. 
func (m *ServicePrincipalCreationConditionSet) GetApplicationIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.applicationIds
    }
}
// Gets the applicationPublisherIds property value. 
func (m *ServicePrincipalCreationConditionSet) GetApplicationPublisherIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.applicationPublisherIds
    }
}
// Gets the applicationsFromVerifiedPublisherOnly property value. 
func (m *ServicePrincipalCreationConditionSet) GetApplicationsFromVerifiedPublisherOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applicationsFromVerifiedPublisherOnly
    }
}
// Gets the applicationTenantIds property value. 
func (m *ServicePrincipalCreationConditionSet) GetApplicationTenantIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.applicationTenantIds
    }
}
// The deserialization information for the current model
func (m *ServicePrincipalCreationConditionSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetApplicationIds(res)
        return nil
    }
    res["applicationPublisherIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetApplicationPublisherIds(res)
        return nil
    }
    res["applicationsFromVerifiedPublisherOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetApplicationsFromVerifiedPublisherOnly(val)
        return nil
    }
    res["applicationTenantIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetApplicationTenantIds(res)
        return nil
    }
    return res
}
func (m *ServicePrincipalCreationConditionSet) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ServicePrincipalCreationConditionSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("applicationIds", m.GetApplicationIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("applicationPublisherIds", m.GetApplicationPublisherIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationsFromVerifiedPublisherOnly", m.GetApplicationsFromVerifiedPublisherOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("applicationTenantIds", m.GetApplicationTenantIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the applicationIds property value. 
// Parameters:
//  - value : Value to set for the applicationIds property.
func (m *ServicePrincipalCreationConditionSet) SetApplicationIds(value []string)() {
    m.applicationIds = value
}
// Sets the applicationPublisherIds property value. 
// Parameters:
//  - value : Value to set for the applicationPublisherIds property.
func (m *ServicePrincipalCreationConditionSet) SetApplicationPublisherIds(value []string)() {
    m.applicationPublisherIds = value
}
// Sets the applicationsFromVerifiedPublisherOnly property value. 
// Parameters:
//  - value : Value to set for the applicationsFromVerifiedPublisherOnly property.
func (m *ServicePrincipalCreationConditionSet) SetApplicationsFromVerifiedPublisherOnly(value *bool)() {
    m.applicationsFromVerifiedPublisherOnly = value
}
// Sets the applicationTenantIds property value. 
// Parameters:
//  - value : Value to set for the applicationTenantIds property.
func (m *ServicePrincipalCreationConditionSet) SetApplicationTenantIds(value []string)() {
    m.applicationTenantIds = value
}
