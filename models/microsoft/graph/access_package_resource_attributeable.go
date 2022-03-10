package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageResourceAttributeable 
type AccessPackageResourceAttributeable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAttributeDestination()(AccessPackageResourceAttributeDestinationable)
    GetAttributeName()(*string)
    GetAttributeSource()(AccessPackageResourceAttributeSourceable)
    GetId()(*string)
    GetIsEditable()(*bool)
    GetIsPersistedOnAssignmentRemoval()(*bool)
    SetAttributeDestination(value AccessPackageResourceAttributeDestinationable)()
    SetAttributeName(value *string)()
    SetAttributeSource(value AccessPackageResourceAttributeSourceable)()
    SetId(value *string)()
    SetIsEditable(value *bool)()
    SetIsPersistedOnAssignmentRemoval(value *bool)()
}
