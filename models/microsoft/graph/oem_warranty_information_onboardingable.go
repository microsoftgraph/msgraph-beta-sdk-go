package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OemWarrantyInformationOnboardingable 
type OemWarrantyInformationOnboardingable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAvailable()(*bool)
    GetEnabled()(*bool)
    GetOemName()(*string)
    SetAvailable(value *bool)()
    SetEnabled(value *bool)()
    SetOemName(value *string)()
}
