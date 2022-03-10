package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentityProtectionRootable 
type IdentityProtectionRootable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetRiskDetections()([]RiskDetectionable)
    GetRiskyServicePrincipals()([]RiskyServicePrincipalable)
    GetRiskyUsers()([]RiskyUserable)
    GetServicePrincipalRiskDetections()([]ServicePrincipalRiskDetectionable)
    SetRiskDetections(value []RiskDetectionable)()
    SetRiskyServicePrincipals(value []RiskyServicePrincipalable)()
    SetRiskyUsers(value []RiskyUserable)()
    SetServicePrincipalRiskDetections(value []ServicePrincipalRiskDetectionable)()
}
