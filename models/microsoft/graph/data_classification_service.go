package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new dataClassificationService and sets the default values.
func NewDataClassificationService()(*DataClassificationService) {
    m := &DataClassificationService{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the classifyFileJobs property value. 
func (m *DataClassificationService) GetClassifyFileJobs()([]JobResponseBase) {
    if m == nil {
        return nil
    } else {
        return m.classifyFileJobs
    }
}
// Gets the classifyTextJobs property value. 
func (m *DataClassificationService) GetClassifyTextJobs()([]JobResponseBase) {
    if m == nil {
        return nil
    } else {
        return m.classifyTextJobs
    }
}
// Gets the evaluateDlpPoliciesJobs property value. 
func (m *DataClassificationService) GetEvaluateDlpPoliciesJobs()([]JobResponseBase) {
    if m == nil {
        return nil
    } else {
        return m.evaluateDlpPoliciesJobs
    }
}
// Gets the evaluateLabelJobs property value. 
func (m *DataClassificationService) GetEvaluateLabelJobs()([]JobResponseBase) {
    if m == nil {
        return nil
    } else {
        return m.evaluateLabelJobs
    }
}
// Gets the exactMatchDataStores property value. 
func (m *DataClassificationService) GetExactMatchDataStores()([]ExactMatchDataStore) {
    if m == nil {
        return nil
    } else {
        return m.exactMatchDataStores
    }
}
// Gets the exactMatchUploadAgents property value. 
func (m *DataClassificationService) GetExactMatchUploadAgents()([]ExactMatchUploadAgent) {
    if m == nil {
        return nil
    } else {
        return m.exactMatchUploadAgents
    }
}
// Gets the jobs property value. 
func (m *DataClassificationService) GetJobs()([]JobResponseBase) {
    if m == nil {
        return nil
    } else {
        return m.jobs
    }
}
// Gets the sensitiveTypes property value. 
func (m *DataClassificationService) GetSensitiveTypes()([]SensitiveType) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypes
    }
}
// Gets the sensitivityLabels property value. 
func (m *DataClassificationService) GetSensitivityLabels()([]SensitivityLabel) {
    if m == nil {
        return nil
    } else {
        return m.sensitivityLabels
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DataClassificationService) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
// Sets the classifyFileJobs property value. 
// Parameters:
//  - value : Value to set for the classifyFileJobs property.
func (m *DataClassificationService) SetClassifyFileJobs(value []JobResponseBase)() {
    m.classifyFileJobs = value
}
// Sets the classifyTextJobs property value. 
// Parameters:
//  - value : Value to set for the classifyTextJobs property.
func (m *DataClassificationService) SetClassifyTextJobs(value []JobResponseBase)() {
    m.classifyTextJobs = value
}
// Sets the evaluateDlpPoliciesJobs property value. 
// Parameters:
//  - value : Value to set for the evaluateDlpPoliciesJobs property.
func (m *DataClassificationService) SetEvaluateDlpPoliciesJobs(value []JobResponseBase)() {
    m.evaluateDlpPoliciesJobs = value
}
// Sets the evaluateLabelJobs property value. 
// Parameters:
//  - value : Value to set for the evaluateLabelJobs property.
func (m *DataClassificationService) SetEvaluateLabelJobs(value []JobResponseBase)() {
    m.evaluateLabelJobs = value
}
// Sets the exactMatchDataStores property value. 
// Parameters:
//  - value : Value to set for the exactMatchDataStores property.
func (m *DataClassificationService) SetExactMatchDataStores(value []ExactMatchDataStore)() {
    m.exactMatchDataStores = value
}
// Sets the exactMatchUploadAgents property value. 
// Parameters:
//  - value : Value to set for the exactMatchUploadAgents property.
func (m *DataClassificationService) SetExactMatchUploadAgents(value []ExactMatchUploadAgent)() {
    m.exactMatchUploadAgents = value
}
// Sets the jobs property value. 
// Parameters:
//  - value : Value to set for the jobs property.
func (m *DataClassificationService) SetJobs(value []JobResponseBase)() {
    m.jobs = value
}
// Sets the sensitiveTypes property value. 
// Parameters:
//  - value : Value to set for the sensitiveTypes property.
func (m *DataClassificationService) SetSensitiveTypes(value []SensitiveType)() {
    m.sensitiveTypes = value
}
// Sets the sensitivityLabels property value. 
// Parameters:
//  - value : Value to set for the sensitivityLabels property.
func (m *DataClassificationService) SetSensitivityLabels(value []SensitivityLabel)() {
    m.sensitivityLabels = value
}
