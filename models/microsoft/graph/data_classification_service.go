package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DataClassificationService struct {
    Entity
    classifyFile []FileClassificationRequest;
    classifyFileJobs []JobResponseBase;
    classifyText []TextClassificationRequest;
    classifyTextJobs []JobResponseBase;
    evaluateDlpPoliciesJobs []JobResponseBase;
    evaluateLabelJobs []JobResponseBase;
    exactMatchDataStores []ExactMatchDataStore;
    exactMatchUploadAgents []ExactMatchUploadAgent;
    jobs []JobResponseBase;
    sensitiveTypes []SensitiveType;
    sensitivityLabels []SensitivityLabel;
}
func NewDataClassificationService()(*DataClassificationService) {
    m := &DataClassificationService{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DataClassificationService) GetClassifyFile()([]FileClassificationRequest) {
    if m == nil {
        return nil
    } else {
        return m.classifyFile
    }
}
func (m *DataClassificationService) GetClassifyFileJobs()([]JobResponseBase) {
    if m == nil {
        return nil
    } else {
        return m.classifyFileJobs
    }
}
func (m *DataClassificationService) GetClassifyText()([]TextClassificationRequest) {
    if m == nil {
        return nil
    } else {
        return m.classifyText
    }
}
func (m *DataClassificationService) GetClassifyTextJobs()([]JobResponseBase) {
    if m == nil {
        return nil
    } else {
        return m.classifyTextJobs
    }
}
func (m *DataClassificationService) GetEvaluateDlpPoliciesJobs()([]JobResponseBase) {
    if m == nil {
        return nil
    } else {
        return m.evaluateDlpPoliciesJobs
    }
}
func (m *DataClassificationService) GetEvaluateLabelJobs()([]JobResponseBase) {
    if m == nil {
        return nil
    } else {
        return m.evaluateLabelJobs
    }
}
func (m *DataClassificationService) GetExactMatchDataStores()([]ExactMatchDataStore) {
    if m == nil {
        return nil
    } else {
        return m.exactMatchDataStores
    }
}
func (m *DataClassificationService) GetExactMatchUploadAgents()([]ExactMatchUploadAgent) {
    if m == nil {
        return nil
    } else {
        return m.exactMatchUploadAgents
    }
}
func (m *DataClassificationService) GetJobs()([]JobResponseBase) {
    if m == nil {
        return nil
    } else {
        return m.jobs
    }
}
func (m *DataClassificationService) GetSensitiveTypes()([]SensitiveType) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypes
    }
}
func (m *DataClassificationService) GetSensitivityLabels()([]SensitivityLabel) {
    if m == nil {
        return nil
    } else {
        return m.sensitivityLabels
    }
}
func (m *DataClassificationService) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classifyFile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFileClassificationRequest() })
        if err != nil {
            return err
        }
        res := make([]FileClassificationRequest, len(val))
        for i, v := range val {
            res[i] = *(v.(*FileClassificationRequest))
        }
        m.SetClassifyFile(res)
        return nil
    }
    res["classifyFileJobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJobResponseBase() })
        if err != nil {
            return err
        }
        res := make([]JobResponseBase, len(val))
        for i, v := range val {
            res[i] = *(v.(*JobResponseBase))
        }
        m.SetClassifyFileJobs(res)
        return nil
    }
    res["classifyText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTextClassificationRequest() })
        if err != nil {
            return err
        }
        res := make([]TextClassificationRequest, len(val))
        for i, v := range val {
            res[i] = *(v.(*TextClassificationRequest))
        }
        m.SetClassifyText(res)
        return nil
    }
    res["classifyTextJobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJobResponseBase() })
        if err != nil {
            return err
        }
        res := make([]JobResponseBase, len(val))
        for i, v := range val {
            res[i] = *(v.(*JobResponseBase))
        }
        m.SetClassifyTextJobs(res)
        return nil
    }
    res["evaluateDlpPoliciesJobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJobResponseBase() })
        if err != nil {
            return err
        }
        res := make([]JobResponseBase, len(val))
        for i, v := range val {
            res[i] = *(v.(*JobResponseBase))
        }
        m.SetEvaluateDlpPoliciesJobs(res)
        return nil
    }
    res["evaluateLabelJobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJobResponseBase() })
        if err != nil {
            return err
        }
        res := make([]JobResponseBase, len(val))
        for i, v := range val {
            res[i] = *(v.(*JobResponseBase))
        }
        m.SetEvaluateLabelJobs(res)
        return nil
    }
    res["exactMatchDataStores"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExactMatchDataStore() })
        if err != nil {
            return err
        }
        res := make([]ExactMatchDataStore, len(val))
        for i, v := range val {
            res[i] = *(v.(*ExactMatchDataStore))
        }
        m.SetExactMatchDataStores(res)
        return nil
    }
    res["exactMatchUploadAgents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExactMatchUploadAgent() })
        if err != nil {
            return err
        }
        res := make([]ExactMatchUploadAgent, len(val))
        for i, v := range val {
            res[i] = *(v.(*ExactMatchUploadAgent))
        }
        m.SetExactMatchUploadAgents(res)
        return nil
    }
    res["jobs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJobResponseBase() })
        if err != nil {
            return err
        }
        res := make([]JobResponseBase, len(val))
        for i, v := range val {
            res[i] = *(v.(*JobResponseBase))
        }
        m.SetJobs(res)
        return nil
    }
    res["sensitiveTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSensitiveType() })
        if err != nil {
            return err
        }
        res := make([]SensitiveType, len(val))
        for i, v := range val {
            res[i] = *(v.(*SensitiveType))
        }
        m.SetSensitiveTypes(res)
        return nil
    }
    res["sensitivityLabels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSensitivityLabel() })
        if err != nil {
            return err
        }
        res := make([]SensitivityLabel, len(val))
        for i, v := range val {
            res[i] = *(v.(*SensitivityLabel))
        }
        m.SetSensitivityLabels(res)
        return nil
    }
    return res
}
func (m *DataClassificationService) IsNil()(bool) {
    return m == nil
}
func (m *DataClassificationService) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClassifyFile()))
        for i, v := range m.GetClassifyFile() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("classifyFile", cast)
        if err != nil {
            return err
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClassifyText()))
        for i, v := range m.GetClassifyText() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("classifyText", cast)
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
func (m *DataClassificationService) SetClassifyFile(value []FileClassificationRequest)() {
    m.classifyFile = value
}
func (m *DataClassificationService) SetClassifyFileJobs(value []JobResponseBase)() {
    m.classifyFileJobs = value
}
func (m *DataClassificationService) SetClassifyText(value []TextClassificationRequest)() {
    m.classifyText = value
}
func (m *DataClassificationService) SetClassifyTextJobs(value []JobResponseBase)() {
    m.classifyTextJobs = value
}
func (m *DataClassificationService) SetEvaluateDlpPoliciesJobs(value []JobResponseBase)() {
    m.evaluateDlpPoliciesJobs = value
}
func (m *DataClassificationService) SetEvaluateLabelJobs(value []JobResponseBase)() {
    m.evaluateLabelJobs = value
}
func (m *DataClassificationService) SetExactMatchDataStores(value []ExactMatchDataStore)() {
    m.exactMatchDataStores = value
}
func (m *DataClassificationService) SetExactMatchUploadAgents(value []ExactMatchUploadAgent)() {
    m.exactMatchUploadAgents = value
}
func (m *DataClassificationService) SetJobs(value []JobResponseBase)() {
    m.jobs = value
}
func (m *DataClassificationService) SetSensitiveTypes(value []SensitiveType)() {
    m.sensitiveTypes = value
}
func (m *DataClassificationService) SetSensitivityLabels(value []SensitivityLabel)() {
    m.sensitivityLabels = value
}
