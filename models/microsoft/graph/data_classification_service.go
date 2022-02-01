package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DataClassificationService 
type DataClassificationService struct {
    Entity
    // 
    classifyFileJobs []JobResponseBase;
    // 
    classifyTextJobs []JobResponseBase;
    // 
    evaluateDlpPoliciesJobs []JobResponseBase;
    // 
    evaluateLabelJobs []JobResponseBase;
    // 
    exactMatchDataStores []ExactMatchDataStore;
    // 
    exactMatchUploadAgents []ExactMatchUploadAgent;
    // 
    jobs []JobResponseBase;
    // 
    sensitiveTypes []SensitiveType;
    // 
    sensitivityLabels []SensitivityLabel;
}
// NewDataClassificationService instantiates a new dataClassificationService and sets the default values.
func NewDataClassificationService()(*DataClassificationService) {
    m := &DataClassificationService{
        Entity: *NewEntity(),
    }
    return m
}
// GetClassifyFileJobs gets the classifyFileJobs property value. 
func (m *DataClassificationService) GetClassifyFileJobs()([]JobResponseBase) {
    if m == nil {
        return nil
    } else {
        return m.classifyFileJobs
    }
}
// GetClassifyTextJobs gets the classifyTextJobs property value. 
func (m *DataClassificationService) GetClassifyTextJobs()([]JobResponseBase) {
    if m == nil {
        return nil
    } else {
        return m.classifyTextJobs
    }
}
// GetEvaluateDlpPoliciesJobs gets the evaluateDlpPoliciesJobs property value. 
func (m *DataClassificationService) GetEvaluateDlpPoliciesJobs()([]JobResponseBase) {
    if m == nil {
        return nil
    } else {
        return m.evaluateDlpPoliciesJobs
    }
}
// GetEvaluateLabelJobs gets the evaluateLabelJobs property value. 
func (m *DataClassificationService) GetEvaluateLabelJobs()([]JobResponseBase) {
    if m == nil {
        return nil
    } else {
        return m.evaluateLabelJobs
    }
}
// GetExactMatchDataStores gets the exactMatchDataStores property value. 
func (m *DataClassificationService) GetExactMatchDataStores()([]ExactMatchDataStore) {
    if m == nil {
        return nil
    } else {
        return m.exactMatchDataStores
    }
}
// GetExactMatchUploadAgents gets the exactMatchUploadAgents property value. 
func (m *DataClassificationService) GetExactMatchUploadAgents()([]ExactMatchUploadAgent) {
    if m == nil {
        return nil
    } else {
        return m.exactMatchUploadAgents
    }
}
// GetJobs gets the jobs property value. 
func (m *DataClassificationService) GetJobs()([]JobResponseBase) {
    if m == nil {
        return nil
    } else {
        return m.jobs
    }
}
// GetSensitiveTypes gets the sensitiveTypes property value. 
func (m *DataClassificationService) GetSensitiveTypes()([]SensitiveType) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypes
    }
}
// GetSensitivityLabels gets the sensitivityLabels property value. 
func (m *DataClassificationService) GetSensitivityLabels()([]SensitivityLabel) {
    if m == nil {
        return nil
    } else {
        return m.sensitivityLabels
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataClassificationService) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classifyFileJobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJobResponseBase() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobResponseBase, len(val))
            for i, v := range val {
                res[i] = *(v.(*JobResponseBase))
            }
            m.SetClassifyFileJobs(res)
        }
        return nil
    }
    res["classifyTextJobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJobResponseBase() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobResponseBase, len(val))
            for i, v := range val {
                res[i] = *(v.(*JobResponseBase))
            }
            m.SetClassifyTextJobs(res)
        }
        return nil
    }
    res["evaluateDlpPoliciesJobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJobResponseBase() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobResponseBase, len(val))
            for i, v := range val {
                res[i] = *(v.(*JobResponseBase))
            }
            m.SetEvaluateDlpPoliciesJobs(res)
        }
        return nil
    }
    res["evaluateLabelJobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJobResponseBase() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobResponseBase, len(val))
            for i, v := range val {
                res[i] = *(v.(*JobResponseBase))
            }
            m.SetEvaluateLabelJobs(res)
        }
        return nil
    }
    res["exactMatchDataStores"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExactMatchDataStore() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExactMatchDataStore, len(val))
            for i, v := range val {
                res[i] = *(v.(*ExactMatchDataStore))
            }
            m.SetExactMatchDataStores(res)
        }
        return nil
    }
    res["exactMatchUploadAgents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExactMatchUploadAgent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExactMatchUploadAgent, len(val))
            for i, v := range val {
                res[i] = *(v.(*ExactMatchUploadAgent))
            }
            m.SetExactMatchUploadAgents(res)
        }
        return nil
    }
    res["jobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJobResponseBase() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobResponseBase, len(val))
            for i, v := range val {
                res[i] = *(v.(*JobResponseBase))
            }
            m.SetJobs(res)
        }
        return nil
    }
    res["sensitiveTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSensitiveType() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitiveType, len(val))
            for i, v := range val {
                res[i] = *(v.(*SensitiveType))
            }
            m.SetSensitiveTypes(res)
        }
        return nil
    }
    res["sensitivityLabels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSensitivityLabel() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitivityLabel, len(val))
            for i, v := range val {
                res[i] = *(v.(*SensitivityLabel))
            }
            m.SetSensitivityLabels(res)
        }
        return nil
    }
    return res
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
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("classifyFileJobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetClassifyTextJobs() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClassifyTextJobs()))
        for i, v := range m.GetClassifyTextJobs() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("classifyTextJobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvaluateDlpPoliciesJobs() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEvaluateDlpPoliciesJobs()))
        for i, v := range m.GetEvaluateDlpPoliciesJobs() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("evaluateDlpPoliciesJobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvaluateLabelJobs() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEvaluateLabelJobs()))
        for i, v := range m.GetEvaluateLabelJobs() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("evaluateLabelJobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExactMatchDataStores() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExactMatchDataStores()))
        for i, v := range m.GetExactMatchDataStores() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("exactMatchDataStores", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExactMatchUploadAgents() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExactMatchUploadAgents()))
        for i, v := range m.GetExactMatchUploadAgents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("exactMatchUploadAgents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetJobs() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetJobs()))
        for i, v := range m.GetJobs() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("jobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensitiveTypes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSensitiveTypes()))
        for i, v := range m.GetSensitiveTypes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("sensitiveTypes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensitivityLabels() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSensitivityLabels()))
        for i, v := range m.GetSensitivityLabels() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("sensitivityLabels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassifyFileJobs sets the classifyFileJobs property value. 
func (m *DataClassificationService) SetClassifyFileJobs(value []JobResponseBase)() {
    if m != nil {
        m.classifyFileJobs = value
    }
}
// SetClassifyTextJobs sets the classifyTextJobs property value. 
func (m *DataClassificationService) SetClassifyTextJobs(value []JobResponseBase)() {
    if m != nil {
        m.classifyTextJobs = value
    }
}
// SetEvaluateDlpPoliciesJobs sets the evaluateDlpPoliciesJobs property value. 
func (m *DataClassificationService) SetEvaluateDlpPoliciesJobs(value []JobResponseBase)() {
    if m != nil {
        m.evaluateDlpPoliciesJobs = value
    }
}
// SetEvaluateLabelJobs sets the evaluateLabelJobs property value. 
func (m *DataClassificationService) SetEvaluateLabelJobs(value []JobResponseBase)() {
    if m != nil {
        m.evaluateLabelJobs = value
    }
}
// SetExactMatchDataStores sets the exactMatchDataStores property value. 
func (m *DataClassificationService) SetExactMatchDataStores(value []ExactMatchDataStore)() {
    if m != nil {
        m.exactMatchDataStores = value
    }
}
// SetExactMatchUploadAgents sets the exactMatchUploadAgents property value. 
func (m *DataClassificationService) SetExactMatchUploadAgents(value []ExactMatchUploadAgent)() {
    if m != nil {
        m.exactMatchUploadAgents = value
    }
}
// SetJobs sets the jobs property value. 
func (m *DataClassificationService) SetJobs(value []JobResponseBase)() {
    if m != nil {
        m.jobs = value
    }
}
// SetSensitiveTypes sets the sensitiveTypes property value. 
func (m *DataClassificationService) SetSensitiveTypes(value []SensitiveType)() {
    if m != nil {
        m.sensitiveTypes = value
    }
}
// SetSensitivityLabels sets the sensitivityLabels property value. 
func (m *DataClassificationService) SetSensitivityLabels(value []SensitivityLabel)() {
    if m != nil {
        m.sensitivityLabels = value
    }
}
