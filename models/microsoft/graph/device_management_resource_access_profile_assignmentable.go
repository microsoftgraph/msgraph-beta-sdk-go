package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementResourceAccessProfileAssignmentable 
type DeviceManagementResourceAccessProfileAssignmentable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetIntent()(*DeviceManagementResourceAccessProfileIntent)
    GetSourceId()(*string)
    GetTarget()(DeviceAndAppManagementAssignmentTargetable)
    SetIntent(value *DeviceManagementResourceAccessProfileIntent)()
    SetSourceId(value *string)()
    SetTarget(value DeviceAndAppManagementAssignmentTargetable)()
}
