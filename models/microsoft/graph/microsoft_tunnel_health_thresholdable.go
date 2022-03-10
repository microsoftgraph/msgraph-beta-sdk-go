package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MicrosoftTunnelHealthThresholdable 
type MicrosoftTunnelHealthThresholdable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDefaultHealthyThreshold()(*int64)
    GetDefaultUnhealthyThreshold()(*int64)
    GetHealthyThreshold()(*int64)
    GetUnhealthyThreshold()(*int64)
    SetDefaultHealthyThreshold(value *int64)()
    SetDefaultUnhealthyThreshold(value *int64)()
    SetHealthyThreshold(value *int64)()
    SetUnhealthyThreshold(value *int64)()
}
