package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DataClassificationServiceable 
type DataClassificationServiceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetClassifyFileJobs()([]JobResponseBaseable)
    GetClassifyTextJobs()([]JobResponseBaseable)
    GetEvaluateDlpPoliciesJobs()([]JobResponseBaseable)
    GetEvaluateLabelJobs()([]JobResponseBaseable)
    GetExactMatchDataStores()([]ExactMatchDataStoreable)
    GetExactMatchUploadAgents()([]ExactMatchUploadAgentable)
    GetJobs()([]JobResponseBaseable)
    GetSensitiveTypes()([]SensitiveTypeable)
    GetSensitivityLabels()([]SensitivityLabelable)
    SetClassifyFileJobs(value []JobResponseBaseable)()
    SetClassifyTextJobs(value []JobResponseBaseable)()
    SetEvaluateDlpPoliciesJobs(value []JobResponseBaseable)()
    SetEvaluateLabelJobs(value []JobResponseBaseable)()
    SetExactMatchDataStores(value []ExactMatchDataStoreable)()
    SetExactMatchUploadAgents(value []ExactMatchUploadAgentable)()
    SetJobs(value []JobResponseBaseable)()
    SetSensitiveTypes(value []SensitiveTypeable)()
    SetSensitivityLabels(value []SensitivityLabelable)()
}
