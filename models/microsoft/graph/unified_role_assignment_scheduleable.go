package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRoleAssignmentScheduleable 
type UnifiedRoleAssignmentScheduleable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    UnifiedRoleScheduleBaseable
    GetActivatedUsing()(UnifiedRoleEligibilityScheduleable)
    GetAssignmentType()(*string)
    GetMemberType()(*string)
    GetScheduleInfo()(RequestScheduleable)
    SetActivatedUsing(value UnifiedRoleEligibilityScheduleable)()
    SetAssignmentType(value *string)()
    SetMemberType(value *string)()
    SetScheduleInfo(value RequestScheduleable)()
}
