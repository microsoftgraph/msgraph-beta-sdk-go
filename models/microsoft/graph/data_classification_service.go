package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DataClassificationService provides operations to manage the dataClassificationService singleton.
type DataClassificationService struct {
    Entity
    // 
    classifyFileJobs []JobResponseBaseable;
    // 
    classifyTextJobs []JobResponseBaseable;
    // 
    evaluateDlpPoliciesJobs []JobResponseBaseable;
    // 
    evaluateLabelJobs []JobResponseBaseable;
    // 
    exactMatchDataStores []ExactMatchDataStoreable;
    // 
    exactMatchUploadAgents []ExactMatchUploadAgentable;
    // 
    jobs []JobResponseBaseable;
    // 
    sensitiveTypes []SensitiveTypeable;
    // 
    sensitivityLabels []SensitivityLabelable;
}
// NewDataClassificationService instantiates a new dataClassificationService and sets the default values.
func NewDataClassificationService()(*DataClassificationService) {
    m := &DataClassificationService{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDataClassificationServiceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDataClassificationServiceFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDataClassificationService(), nil
}
// GetClassifyFileJobs gets the classifyFileJobs property value. 
func (m *DataClassificationService) GetClassifyFileJobs()([]JobResponseBaseable) {
    if m == nil {
        return nil
    } else {
        return m.classifyFileJobs
    }
}
// GetClassifyTextJobs gets the classifyTextJobs property value. 
func (m *DataClassificationService) GetClassifyTextJobs()([]JobResponseBaseable) {
    if m == nil {
        return nil
    } else {
        return m.classifyTextJobs
    }
}
// GetEvaluateDlpPoliciesJobs gets the evaluateDlpPoliciesJobs property value. 
func (m *DataClassificationService) GetEvaluateDlpPoliciesJobs()([]JobResponseBaseable) {
    if m == nil {
        return nil
    } else {
        return m.evaluateDlpPoliciesJobs
    }
}
// GetEvaluateLabelJobs gets the evaluateLabelJobs property value. 
func (m *DataClassificationService) GetEvaluateLabelJobs()([]JobResponseBaseable) {
    if m == nil {
        return nil
    } else {
        return m.evaluateLabelJobs
    }
}
// GetExactMatchDataStores gets the exactMatchDataStores property value. 
func (m *DataClassificationService) GetExactMatchDataStores()([]ExactMatchDataStoreable) {
    if m == nil {
        return nil
    } else {
        return m.exactMatchDataStores
    }
}
// GetExactMatchUploadAgents gets the exactMatchUploadAgents property value. 
func (m *DataClassificationService) GetExactMatchUploadAgents()([]ExactMatchUploadAgentable) {
    if m == nil {
        return nil
    } else {
        return m.exactMatchUploadAgents
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataClassificationService) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classifyFileJobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJobResponseBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobResponseBaseable, len(val))
            for i, v := range val {
                res[i] = v.(JobResponseBaseable)
            }
            m.SetClassifyFileJobs(res)
        }
        return nil
    }
    res["classifyTextJobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJobResponseBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobResponseBaseable, len(val))
            for i, v := range val {
                res[i] = v.(JobResponseBaseable)
            }
            m.SetClassifyTextJobs(res)
        }
        return nil
    }
    res["evaluateDlpPoliciesJobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJobResponseBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobResponseBaseable, len(val))
            for i, v := range val {
                res[i] = v.(JobResponseBaseable)
            }
            m.SetEvaluateDlpPoliciesJobs(res)
        }
        return nil
    }
    res["evaluateLabelJobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJobResponseBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobResponseBaseable, len(val))
            for i, v := range val {
                res[i] = v.(JobResponseBaseable)
            }
            m.SetEvaluateLabelJobs(res)
        }
        return nil
    }
    res["exactMatchDataStores"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExactMatchDataStoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExactMatchDataStoreable, len(val))
            for i, v := range val {
                res[i] = v.(ExactMatchDataStoreable)
            }
            m.SetExactMatchDataStores(res)
        }
        return nil
    }
    res["exactMatchUploadAgents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExactMatchUploadAgentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExactMatchUploadAgentable, len(val))
            for i, v := range val {
                res[i] = v.(ExactMatchUploadAgentable)
            }
            m.SetExactMatchUploadAgents(res)
        }
        return nil
    }
    res["jobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJobResponseBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobResponseBaseable, len(val))
            for i, v := range val {
                res[i] = v.(JobResponseBaseable)
            }
            m.SetJobs(res)
        }
        return nil
    }
    res["sensitiveTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensitiveTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitiveTypeable, len(val))
            for i, v := range val {
                res[i] = v.(SensitiveTypeable)
            }
            m.SetSensitiveTypes(res)
        }
        return nil
    }
    res["sensitivityLabels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensitivityLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitivityLabelable, len(val))
            for i, v := range val {
                res[i] = v.(SensitivityLabelable)
            }
            m.SetSensitivityLabels(res)
        }
        return nil
    }
    return res
}
// GetJobs gets the jobs property value. 
func (m *DataClassificationService) GetJobs()([]JobResponseBaseable) {
    if m == nil {
        return nil
    } else {
        return m.jobs
    }
}
// GetSensitiveTypes gets the sensitiveTypes property value. 
func (m *DataClassificationService) GetSensitiveTypes()([]SensitiveTypeable) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypes
    }
}
// GetSensitivityLabels gets the sensitivityLabels property value. 
func (m *DataClassificationService) GetSensitivityLabels()([]SensitivityLabelable) {
    if m == nil {
        return nil
    } else {
        return m.sensitivityLabels
    }
}
func (m *DataClassificationService) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DataClassificationService) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassifyFileJobs() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClassifyFileJobs()))
        for i, v := range m.GetClassifyFileJobs() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("classifyFileJobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetClassifyTextJobs() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClassifyTextJobs()))
        for i, v := range m.GetClassifyTextJobs() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("classifyTextJobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvaluateDlpPoliciesJobs() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEvaluateDlpPoliciesJobs()))
        for i, v := range m.GetEvaluateDlpPoliciesJobs() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("evaluateDlpPoliciesJobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvaluateLabelJobs() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEvaluateLabelJobs()))
        for i, v := range m.GetEvaluateLabelJobs() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("evaluateLabelJobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExactMatchDataStores() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExactMatchDataStores()))
        for i, v := range m.GetExactMatchDataStores() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("exactMatchDataStores", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExactMatchUploadAgents() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExactMatchUploadAgents()))
        for i, v := range m.GetExactMatchUploadAgents() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("exactMatchUploadAgents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetJobs() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetJobs()))
        for i, v := range m.GetJobs() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("jobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensitiveTypes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSensitiveTypes()))
        for i, v := range m.GetSensitiveTypes() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sensitiveTypes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensitivityLabels() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSensitivityLabels()))
        for i, v := range m.GetSensitivityLabels() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sensitivityLabels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassifyFileJobs sets the classifyFileJobs property value. 
func (m *DataClassificationService) SetClassifyFileJobs(value []JobResponseBaseable)() {
    if m != nil {
        m.classifyFileJobs = value
    }
}
// SetClassifyTextJobs sets the classifyTextJobs property value. 
func (m *DataClassificationService) SetClassifyTextJobs(value []JobResponseBaseable)() {
    if m != nil {
        m.classifyTextJobs = value
    }
}
// SetEvaluateDlpPoliciesJobs sets the evaluateDlpPoliciesJobs property value. 
func (m *DataClassificationService) SetEvaluateDlpPoliciesJobs(value []JobResponseBaseable)() {
    if m != nil {
        m.evaluateDlpPoliciesJobs = value
    }
}
// SetEvaluateLabelJobs sets the evaluateLabelJobs property value. 
func (m *DataClassificationService) SetEvaluateLabelJobs(value []JobResponseBaseable)() {
    if m != nil {
        m.evaluateLabelJobs = value
    }
}
// SetExactMatchDataStores sets the exactMatchDataStores property value. 
func (m *DataClassificationService) SetExactMatchDataStores(value []ExactMatchDataStoreable)() {
    if m != nil {
        m.exactMatchDataStores = value
    }
}
// SetExactMatchUploadAgents sets the exactMatchUploadAgents property value. 
func (m *DataClassificationService) SetExactMatchUploadAgents(value []ExactMatchUploadAgentable)() {
    if m != nil {
        m.exactMatchUploadAgents = value
    }
}
// SetJobs sets the jobs property value. 
func (m *DataClassificationService) SetJobs(value []JobResponseBaseable)() {
    if m != nil {
        m.jobs = value
    }
}
// SetSensitiveTypes sets the sensitiveTypes property value. 
func (m *DataClassificationService) SetSensitiveTypes(value []SensitiveTypeable)() {
    if m != nil {
        m.sensitiveTypes = value
    }
}
// SetSensitivityLabels sets the sensitivityLabels property value. 
func (m *DataClassificationService) SetSensitivityLabels(value []SensitivityLabelable)() {
    if m != nil {
        m.sensitivityLabels = value
    }
}
