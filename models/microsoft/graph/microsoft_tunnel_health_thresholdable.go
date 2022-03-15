package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MicrosoftTunnelHealthThresholdable 
type MicrosoftTunnelHealthThresholdable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDefaultHealthyThreshold()(*int32)
    GetDefaultUnhealthyThreshold()(*int32)
    GetHealthyThreshold()(*int32)
    GetUnhealthyThreshold()(*int32)
    SetDefaultHealthyThreshold(value *int32)()
    SetDefaultUnhealthyThreshold(value *int32)()
    SetHealthyThreshold(value *int32)()
    SetUnhealthyThreshold(value *int32)()
}
