package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcCrossCloudGovernmentOrganizationMapping 
type CloudPcCrossCloudGovernmentOrganizationMapping struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewCloudPcCrossCloudGovernmentOrganizationMapping instantiates a new cloudPcCrossCloudGovernmentOrganizationMapping and sets the default values.
func NewCloudPcCrossCloudGovernmentOrganizationMapping()(*CloudPcCrossCloudGovernmentOrganizationMapping) {
    m := &CloudPcCrossCloudGovernmentOrganizationMapping{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcCrossCloudGovernmentOrganizationMappingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcCrossCloudGovernmentOrganizationMappingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcCrossCloudGovernmentOrganizationMapping(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcCrossCloudGovernmentOrganizationMapping) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["organizationIdsInUSGovCloud"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetOrganizationIdsInUSGovCloud(res)
        }
        return nil
    }
    return res
}
// GetOrganizationIdsInUSGovCloud gets the organizationIdsInUSGovCloud property value. The tenant ID in the Azure Government cloud corresponding to the GCC tenant in the public cloud. Currently, 1:1 mappings are supported, so this collection can only contain one tenant ID.
func (m *CloudPcCrossCloudGovernmentOrganizationMapping) GetOrganizationIdsInUSGovCloud()([]string) {
    val, err := m.GetBackingStore().Get("organizationIdsInUSGovCloud")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcCrossCloudGovernmentOrganizationMapping) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetOrganizationIdsInUSGovCloud() != nil {
        err = writer.WriteCollectionOfStringValues("organizationIdsInUSGovCloud", m.GetOrganizationIdsInUSGovCloud())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOrganizationIdsInUSGovCloud sets the organizationIdsInUSGovCloud property value. The tenant ID in the Azure Government cloud corresponding to the GCC tenant in the public cloud. Currently, 1:1 mappings are supported, so this collection can only contain one tenant ID.
func (m *CloudPcCrossCloudGovernmentOrganizationMapping) SetOrganizationIdsInUSGovCloud(value []string)() {
    err := m.GetBackingStore().Set("organizationIdsInUSGovCloud", value)
    if err != nil {
        panic(err)
    }
}
// CloudPcCrossCloudGovernmentOrganizationMappingable 
type CloudPcCrossCloudGovernmentOrganizationMappingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOrganizationIdsInUSGovCloud()([]string)
    SetOrganizationIdsInUSGovCloud(value []string)()
}
