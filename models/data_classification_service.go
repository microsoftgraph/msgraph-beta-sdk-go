package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DataClassificationService 
type DataClassificationService struct {
    Entity
    // The classifyFileJobs property
    classifyFileJobs []JobResponseBaseable
    // The classifyTextJobs property
    classifyTextJobs []JobResponseBaseable
    // The evaluateDlpPoliciesJobs property
    evaluateDlpPoliciesJobs []JobResponseBaseable
    // The evaluateLabelJobs property
    evaluateLabelJobs []JobResponseBaseable
    // The exactMatchDataStores property
    exactMatchDataStores []ExactMatchDataStoreable
    // The exactMatchUploadAgents property
    exactMatchUploadAgents []ExactMatchUploadAgentable
    // The jobs property
    jobs []JobResponseBaseable
    // The sensitiveTypes property
    sensitiveTypes []SensitiveTypeable
    // The sensitivityLabels property
    sensitivityLabels []SensitivityLabelable
}
// NewDataClassificationService instantiates a new DataClassificationService and sets the default values.
func NewDataClassificationService()(*DataClassificationService) {
    m := &DataClassificationService{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.dataClassificationService";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDataClassificationServiceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDataClassificationServiceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataClassificationService(), nil
}
// GetClassifyFileJobs gets the classifyFileJobs property value. The classifyFileJobs property
func (m *DataClassificationService) GetClassifyFileJobs()([]JobResponseBaseable) {
    return m.classifyFileJobs
}
// GetClassifyTextJobs gets the classifyTextJobs property value. The classifyTextJobs property
func (m *DataClassificationService) GetClassifyTextJobs()([]JobResponseBaseable) {
    return m.classifyTextJobs
}
// GetEvaluateDlpPoliciesJobs gets the evaluateDlpPoliciesJobs property value. The evaluateDlpPoliciesJobs property
func (m *DataClassificationService) GetEvaluateDlpPoliciesJobs()([]JobResponseBaseable) {
    return m.evaluateDlpPoliciesJobs
}
// GetEvaluateLabelJobs gets the evaluateLabelJobs property value. The evaluateLabelJobs property
func (m *DataClassificationService) GetEvaluateLabelJobs()([]JobResponseBaseable) {
    return m.evaluateLabelJobs
}
// GetExactMatchDataStores gets the exactMatchDataStores property value. The exactMatchDataStores property
func (m *DataClassificationService) GetExactMatchDataStores()([]ExactMatchDataStoreable) {
    return m.exactMatchDataStores
}
// GetExactMatchUploadAgents gets the exactMatchUploadAgents property value. The exactMatchUploadAgents property
func (m *DataClassificationService) GetExactMatchUploadAgents()([]ExactMatchUploadAgentable) {
    return m.exactMatchUploadAgents
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataClassificationService) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classifyFileJobs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateJobResponseBaseFromDiscriminatorValue , m.SetClassifyFileJobs)
    res["classifyTextJobs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateJobResponseBaseFromDiscriminatorValue , m.SetClassifyTextJobs)
    res["evaluateDlpPoliciesJobs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateJobResponseBaseFromDiscriminatorValue , m.SetEvaluateDlpPoliciesJobs)
    res["evaluateLabelJobs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateJobResponseBaseFromDiscriminatorValue , m.SetEvaluateLabelJobs)
    res["exactMatchDataStores"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateExactMatchDataStoreFromDiscriminatorValue , m.SetExactMatchDataStores)
    res["exactMatchUploadAgents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateExactMatchUploadAgentFromDiscriminatorValue , m.SetExactMatchUploadAgents)
    res["jobs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateJobResponseBaseFromDiscriminatorValue , m.SetJobs)
    res["sensitiveTypes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSensitiveTypeFromDiscriminatorValue , m.SetSensitiveTypes)
    res["sensitivityLabels"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSensitivityLabelFromDiscriminatorValue , m.SetSensitivityLabels)
    return res
}
// GetJobs gets the jobs property value. The jobs property
func (m *DataClassificationService) GetJobs()([]JobResponseBaseable) {
    return m.jobs
}
// GetSensitiveTypes gets the sensitiveTypes property value. The sensitiveTypes property
func (m *DataClassificationService) GetSensitiveTypes()([]SensitiveTypeable) {
    return m.sensitiveTypes
}
// GetSensitivityLabels gets the sensitivityLabels property value. The sensitivityLabels property
func (m *DataClassificationService) GetSensitivityLabels()([]SensitivityLabelable) {
    return m.sensitivityLabels
}
// Serialize serializes information the current object
func (m *DataClassificationService) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassifyFileJobs() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetClassifyFileJobs())
        err = writer.WriteCollectionOfObjectValues("classifyFileJobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetClassifyTextJobs() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetClassifyTextJobs())
        err = writer.WriteCollectionOfObjectValues("classifyTextJobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvaluateDlpPoliciesJobs() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEvaluateDlpPoliciesJobs())
        err = writer.WriteCollectionOfObjectValues("evaluateDlpPoliciesJobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvaluateLabelJobs() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEvaluateLabelJobs())
        err = writer.WriteCollectionOfObjectValues("evaluateLabelJobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExactMatchDataStores() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExactMatchDataStores())
        err = writer.WriteCollectionOfObjectValues("exactMatchDataStores", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExactMatchUploadAgents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExactMatchUploadAgents())
        err = writer.WriteCollectionOfObjectValues("exactMatchUploadAgents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetJobs() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetJobs())
        err = writer.WriteCollectionOfObjectValues("jobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensitiveTypes() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSensitiveTypes())
        err = writer.WriteCollectionOfObjectValues("sensitiveTypes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensitivityLabels() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSensitivityLabels())
        err = writer.WriteCollectionOfObjectValues("sensitivityLabels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassifyFileJobs sets the classifyFileJobs property value. The classifyFileJobs property
func (m *DataClassificationService) SetClassifyFileJobs(value []JobResponseBaseable)() {
    m.classifyFileJobs = value
}
// SetClassifyTextJobs sets the classifyTextJobs property value. The classifyTextJobs property
func (m *DataClassificationService) SetClassifyTextJobs(value []JobResponseBaseable)() {
    m.classifyTextJobs = value
}
// SetEvaluateDlpPoliciesJobs sets the evaluateDlpPoliciesJobs property value. The evaluateDlpPoliciesJobs property
func (m *DataClassificationService) SetEvaluateDlpPoliciesJobs(value []JobResponseBaseable)() {
    m.evaluateDlpPoliciesJobs = value
}
// SetEvaluateLabelJobs sets the evaluateLabelJobs property value. The evaluateLabelJobs property
func (m *DataClassificationService) SetEvaluateLabelJobs(value []JobResponseBaseable)() {
    m.evaluateLabelJobs = value
}
// SetExactMatchDataStores sets the exactMatchDataStores property value. The exactMatchDataStores property
func (m *DataClassificationService) SetExactMatchDataStores(value []ExactMatchDataStoreable)() {
    m.exactMatchDataStores = value
}
// SetExactMatchUploadAgents sets the exactMatchUploadAgents property value. The exactMatchUploadAgents property
func (m *DataClassificationService) SetExactMatchUploadAgents(value []ExactMatchUploadAgentable)() {
    m.exactMatchUploadAgents = value
}
// SetJobs sets the jobs property value. The jobs property
func (m *DataClassificationService) SetJobs(value []JobResponseBaseable)() {
    m.jobs = value
}
// SetSensitiveTypes sets the sensitiveTypes property value. The sensitiveTypes property
func (m *DataClassificationService) SetSensitiveTypes(value []SensitiveTypeable)() {
    m.sensitiveTypes = value
}
// SetSensitivityLabels sets the sensitivityLabels property value. The sensitivityLabels property
func (m *DataClassificationService) SetSensitivityLabels(value []SensitivityLabelable)() {
    m.sensitivityLabels = value
}
