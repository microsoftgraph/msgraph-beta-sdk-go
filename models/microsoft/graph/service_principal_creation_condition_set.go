package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServicePrincipalCreationConditionSet 
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
// NewServicePrincipalCreationConditionSet instantiates a new servicePrincipalCreationConditionSet and sets the default values.
func NewServicePrincipalCreationConditionSet()(*ServicePrincipalCreationConditionSet) {
    m := &ServicePrincipalCreationConditionSet{
        Entity: *NewEntity(),
    }
    return m
}
// GetApplicationIds gets the applicationIds property value. 
func (m *ServicePrincipalCreationConditionSet) GetApplicationIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.applicationIds
    }
}
// GetApplicationPublisherIds gets the applicationPublisherIds property value. 
func (m *ServicePrincipalCreationConditionSet) GetApplicationPublisherIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.applicationPublisherIds
    }
}
// GetApplicationsFromVerifiedPublisherOnly gets the applicationsFromVerifiedPublisherOnly property value. 
func (m *ServicePrincipalCreationConditionSet) GetApplicationsFromVerifiedPublisherOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applicationsFromVerifiedPublisherOnly
    }
}
// GetApplicationTenantIds gets the applicationTenantIds property value. 
func (m *ServicePrincipalCreationConditionSet) GetApplicationTenantIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.applicationTenantIds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalCreationConditionSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetApplicationIds(res)
        }
        return nil
    }
    res["applicationPublisherIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetApplicationPublisherIds(res)
        }
        return nil
    }
    res["applicationsFromVerifiedPublisherOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationsFromVerifiedPublisherOnly(val)
        }
        return nil
    }
    res["applicationTenantIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetApplicationTenantIds(res)
        }
        return nil
    }
    return res
}
func (m *ServicePrincipalCreationConditionSet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ServicePrincipalCreationConditionSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicationIds() != nil {
        err = writer.WriteCollectionOfStringValues("applicationIds", m.GetApplicationIds())
        if err != nil {
            return err
        }
    }
    if m.GetApplicationPublisherIds() != nil {
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
    if m.GetApplicationTenantIds() != nil {
        err = writer.WriteCollectionOfStringValues("applicationTenantIds", m.GetApplicationTenantIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationIds sets the applicationIds property value. 
func (m *ServicePrincipalCreationConditionSet) SetApplicationIds(value []string)() {
    if m != nil {
        m.applicationIds = value
    }
}
// SetApplicationPublisherIds sets the applicationPublisherIds property value. 
func (m *ServicePrincipalCreationConditionSet) SetApplicationPublisherIds(value []string)() {
    if m != nil {
        m.applicationPublisherIds = value
    }
}
// SetApplicationsFromVerifiedPublisherOnly sets the applicationsFromVerifiedPublisherOnly property value. 
func (m *ServicePrincipalCreationConditionSet) SetApplicationsFromVerifiedPublisherOnly(value *bool)() {
    if m != nil {
        m.applicationsFromVerifiedPublisherOnly = value
    }
}
// SetApplicationTenantIds sets the applicationTenantIds property value. 
func (m *ServicePrincipalCreationConditionSet) SetApplicationTenantIds(value []string)() {
    if m != nil {
        m.applicationTenantIds = value
    }
}
