package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServicePrincipalCreationConditionSetable 
type ServicePrincipalCreationConditionSetable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplicationIds()([]string)
    GetApplicationPublisherIds()([]string)
    GetApplicationsFromVerifiedPublisherOnly()(*bool)
    GetApplicationTenantIds()([]string)
    GetCertifiedApplicationsOnly()(*bool)
    SetApplicationIds(value []string)()
    SetApplicationPublisherIds(value []string)()
    SetApplicationsFromVerifiedPublisherOnly(value *bool)()
    SetApplicationTenantIds(value []string)()
    SetCertifiedApplicationsOnly(value *bool)()
}
