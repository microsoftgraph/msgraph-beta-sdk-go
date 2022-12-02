package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplicationsItemSynchronizationJobsItemSchemaFilterOperatorsResponse provides operations to call the filterOperators method.
type ApplicationsItemSynchronizationJobsItemSchemaFilterOperatorsResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []FilterOperatorSchemaable
}
// NewApplicationsItemSynchronizationJobsItemSchemaFilterOperatorsResponse instantiates a new ApplicationsItemSynchronizationJobsItemSchemaFilterOperatorsResponse and sets the default values.
func NewApplicationsItemSynchronizationJobsItemSchemaFilterOperatorsResponse()(*ApplicationsItemSynchronizationJobsItemSchemaFilterOperatorsResponse) {
    m := &ApplicationsItemSynchronizationJobsItemSchemaFilterOperatorsResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateApplicationsItemSynchronizationJobsItemSchemaFilterOperatorsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicationsItemSynchronizationJobsItemSchemaFilterOperatorsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplicationsItemSynchronizationJobsItemSchemaFilterOperatorsResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicationsItemSynchronizationJobsItemSchemaFilterOperatorsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFilterOperatorSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FilterOperatorSchemaable, len(val))
            for i, v := range val {
                res[i] = v.(FilterOperatorSchemaable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ApplicationsItemSynchronizationJobsItemSchemaFilterOperatorsResponse) GetValue()([]FilterOperatorSchemaable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ApplicationsItemSynchronizationJobsItemSchemaFilterOperatorsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *ApplicationsItemSynchronizationJobsItemSchemaFilterOperatorsResponse) SetValue(value []FilterOperatorSchemaable)() {
    m.value = value
}
