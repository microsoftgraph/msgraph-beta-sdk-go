package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileAppRelationshipable 
type MobileAppRelationshipable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetTargetDisplayName()(*string)
    GetTargetDisplayVersion()(*string)
    GetTargetId()(*string)
    GetTargetPublisher()(*string)
    GetTargetType()(*MobileAppRelationshipType)
    SetTargetDisplayName(value *string)()
    SetTargetDisplayVersion(value *string)()
    SetTargetId(value *string)()
    SetTargetPublisher(value *string)()
    SetTargetType(value *MobileAppRelationshipType)()
}
