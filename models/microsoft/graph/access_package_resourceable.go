package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageResourceable 
type AccessPackageResourceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessPackageResourceEnvironment()(AccessPackageResourceEnvironmentable)
    GetAccessPackageResourceRoles()([]AccessPackageResourceRoleable)
    GetAccessPackageResourceScopes()([]AccessPackageResourceScopeable)
    GetAddedBy()(*string)
    GetAddedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAttributes()([]AccessPackageResourceAttributeable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsPendingOnboarding()(*bool)
    GetOriginId()(*string)
    GetOriginSystem()(*string)
    GetResourceType()(*string)
    GetUrl()(*string)
    SetAccessPackageResourceEnvironment(value AccessPackageResourceEnvironmentable)()
    SetAccessPackageResourceRoles(value []AccessPackageResourceRoleable)()
    SetAccessPackageResourceScopes(value []AccessPackageResourceScopeable)()
    SetAddedBy(value *string)()
    SetAddedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAttributes(value []AccessPackageResourceAttributeable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsPendingOnboarding(value *bool)()
    SetOriginId(value *string)()
    SetOriginSystem(value *string)()
    SetResourceType(value *string)()
    SetUrl(value *string)()
}
