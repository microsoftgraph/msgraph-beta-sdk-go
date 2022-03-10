package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentResourceRoleable 
type AccessPackageAssignmentResourceRoleable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessPackageAssignments()([]AccessPackageAssignmentable)
    GetAccessPackageResourceRole()(AccessPackageResourceRoleable)
    GetAccessPackageResourceScope()(AccessPackageResourceScopeable)
    GetAccessPackageSubject()(AccessPackageSubjectable)
    GetOriginId()(*string)
    GetOriginSystem()(*string)
    GetStatus()(*string)
    SetAccessPackageAssignments(value []AccessPackageAssignmentable)()
    SetAccessPackageResourceRole(value AccessPackageResourceRoleable)()
    SetAccessPackageResourceScope(value AccessPackageResourceScopeable)()
    SetAccessPackageSubject(value AccessPackageSubjectable)()
    SetOriginId(value *string)()
    SetOriginSystem(value *string)()
    SetStatus(value *string)()
}
