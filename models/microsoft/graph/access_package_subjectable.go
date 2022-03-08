package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageSubjectable 
type AccessPackageSubjectable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAltSecId()(*string)
    GetConnectedOrganization()(ConnectedOrganizationable)
    GetConnectedOrganizationId()(*string)
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetObjectId()(*string)
    GetOnPremisesSecurityIdentifier()(*string)
    GetPrincipalName()(*string)
    GetType()(*string)
    SetAltSecId(value *string)()
    SetConnectedOrganization(value ConnectedOrganizationable)()
    SetConnectedOrganizationId(value *string)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetObjectId(value *string)()
    SetOnPremisesSecurityIdentifier(value *string)()
    SetPrincipalName(value *string)()
    SetType(value *string)()
}
