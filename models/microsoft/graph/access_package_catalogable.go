package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageCatalogable 
type AccessPackageCatalogable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessPackageResourceRoles()([]AccessPackageResourceRoleable)
    GetAccessPackageResources()([]AccessPackageResourceable)
    GetAccessPackageResourceScopes()([]AccessPackageResourceScopeable)
    GetAccessPackages()([]AccessPackageable)
    GetCatalogStatus()(*string)
    GetCatalogType()(*string)
    GetCreatedBy()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomAccessPackageWorkflowExtensions()([]CustomAccessPackageWorkflowExtensionable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsExternallyVisible()(*bool)
    GetModifiedBy()(*string)
    GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccessPackageResourceRoles(value []AccessPackageResourceRoleable)()
    SetAccessPackageResources(value []AccessPackageResourceable)()
    SetAccessPackageResourceScopes(value []AccessPackageResourceScopeable)()
    SetAccessPackages(value []AccessPackageable)()
    SetCatalogStatus(value *string)()
    SetCatalogType(value *string)()
    SetCreatedBy(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomAccessPackageWorkflowExtensions(value []CustomAccessPackageWorkflowExtensionable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsExternallyVisible(value *bool)()
    SetModifiedBy(value *string)()
    SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
