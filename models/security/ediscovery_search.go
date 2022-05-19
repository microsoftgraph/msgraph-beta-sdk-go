package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoverySearch provides operations to manage the security singleton.
type EdiscoverySearch struct {
    Search
    // The additionalSources property
    additionalSources []DataSourceable
    // The addToReviewSetOperation property
    addToReviewSetOperation EdiscoveryAddToReviewSetOperationable
    // The custodianSources property
    custodianSources []DataSourceable
    // The dataSourceScopes property
    dataSourceScopes *DataSourceScopes
    // The lastEstimateStatisticsOperation property
    lastEstimateStatisticsOperation EdiscoveryEstimateOperationable
    // The noncustodialSources property
    noncustodialSources []EdiscoveryNoncustodialDataSourceable
}
// NewEdiscoverySearch instantiates a new ediscoverySearch and sets the default values.
func NewEdiscoverySearch()(*EdiscoverySearch) {
    m := &EdiscoverySearch{
        Search: *NewSearch(),
    }
    return m
}
// CreateEdiscoverySearchFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoverySearchFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoverySearch(), nil
}
// GetAdditionalSources gets the additionalSources property value. The additionalSources property
func (m *EdiscoverySearch) GetAdditionalSources()([]DataSourceable) {
    if m == nil {
        return nil
    } else {
        return m.additionalSources
    }
}
// GetAddToReviewSetOperation gets the addToReviewSetOperation property value. The addToReviewSetOperation property
func (m *EdiscoverySearch) GetAddToReviewSetOperation()(EdiscoveryAddToReviewSetOperationable) {
    if m == nil {
        return nil
    } else {
        return m.addToReviewSetOperation
    }
}
// GetCustodianSources gets the custodianSources property value. The custodianSources property
func (m *EdiscoverySearch) GetCustodianSources()([]DataSourceable) {
    if m == nil {
        return nil
    } else {
        return m.custodianSources
    }
}
// GetDataSourceScopes gets the dataSourceScopes property value. The dataSourceScopes property
func (m *EdiscoverySearch) GetDataSourceScopes()(*DataSourceScopes) {
    if m == nil {
        return nil
    } else {
        return m.dataSourceScopes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoverySearch) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Search.GetFieldDeserializers()
    res["additionalSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDataSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DataSourceable, len(val))
            for i, v := range val {
                res[i] = v.(DataSourceable)
            }
            m.SetAdditionalSources(res)
        }
        return nil
    }
    res["addToReviewSetOperation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoveryAddToReviewSetOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddToReviewSetOperation(val.(EdiscoveryAddToReviewSetOperationable))
        }
        return nil
    }
    res["custodianSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDataSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DataSourceable, len(val))
            for i, v := range val {
                res[i] = v.(DataSourceable)
            }
            m.SetCustodianSources(res)
        }
        return nil
    }
    res["dataSourceScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataSourceScopes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataSourceScopes(val.(*DataSourceScopes))
        }
        return nil
    }
    res["lastEstimateStatisticsOperation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoveryEstimateOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastEstimateStatisticsOperation(val.(EdiscoveryEstimateOperationable))
        }
        return nil
    }
    res["noncustodialSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoveryNoncustodialDataSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoveryNoncustodialDataSourceable, len(val))
            for i, v := range val {
                res[i] = v.(EdiscoveryNoncustodialDataSourceable)
            }
            m.SetNoncustodialSources(res)
        }
        return nil
    }
    return res
}
// GetLastEstimateStatisticsOperation gets the lastEstimateStatisticsOperation property value. The lastEstimateStatisticsOperation property
func (m *EdiscoverySearch) GetLastEstimateStatisticsOperation()(EdiscoveryEstimateOperationable) {
    if m == nil {
        return nil
    } else {
        return m.lastEstimateStatisticsOperation
    }
}
// GetNoncustodialSources gets the noncustodialSources property value. The noncustodialSources property
func (m *EdiscoverySearch) GetNoncustodialSources()([]EdiscoveryNoncustodialDataSourceable) {
    if m == nil {
        return nil
    } else {
        return m.noncustodialSources
    }
}
// Serialize serializes information the current object
func (m *EdiscoverySearch) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Search.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdditionalSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdditionalSources()))
        for i, v := range m.GetAdditionalSources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("additionalSources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("addToReviewSetOperation", m.GetAddToReviewSetOperation())
        if err != nil {
            return err
        }
    }
    if m.GetCustodianSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustodianSources()))
        for i, v := range m.GetCustodianSources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("custodianSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDataSourceScopes() != nil {
        cast := (*m.GetDataSourceScopes()).String()
        err = writer.WriteStringValue("dataSourceScopes", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastEstimateStatisticsOperation", m.GetLastEstimateStatisticsOperation())
        if err != nil {
            return err
        }
    }
    if m.GetNoncustodialSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNoncustodialSources()))
        for i, v := range m.GetNoncustodialSources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("noncustodialSources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalSources sets the additionalSources property value. The additionalSources property
func (m *EdiscoverySearch) SetAdditionalSources(value []DataSourceable)() {
    if m != nil {
        m.additionalSources = value
    }
}
// SetAddToReviewSetOperation sets the addToReviewSetOperation property value. The addToReviewSetOperation property
func (m *EdiscoverySearch) SetAddToReviewSetOperation(value EdiscoveryAddToReviewSetOperationable)() {
    if m != nil {
        m.addToReviewSetOperation = value
    }
}
// SetCustodianSources sets the custodianSources property value. The custodianSources property
func (m *EdiscoverySearch) SetCustodianSources(value []DataSourceable)() {
    if m != nil {
        m.custodianSources = value
    }
}
// SetDataSourceScopes sets the dataSourceScopes property value. The dataSourceScopes property
func (m *EdiscoverySearch) SetDataSourceScopes(value *DataSourceScopes)() {
    if m != nil {
        m.dataSourceScopes = value
    }
}
// SetLastEstimateStatisticsOperation sets the lastEstimateStatisticsOperation property value. The lastEstimateStatisticsOperation property
func (m *EdiscoverySearch) SetLastEstimateStatisticsOperation(value EdiscoveryEstimateOperationable)() {
    if m != nil {
        m.lastEstimateStatisticsOperation = value
    }
}
// SetNoncustodialSources sets the noncustodialSources property value. The noncustodialSources property
func (m *EdiscoverySearch) SetNoncustodialSources(value []EdiscoveryNoncustodialDataSourceable)() {
    if m != nil {
        m.noncustodialSources = value
    }
}
