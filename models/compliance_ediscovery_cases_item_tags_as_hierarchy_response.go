package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComplianceEdiscoveryCasesItemTagsAsHierarchyResponse provides operations to call the asHierarchy method.
type ComplianceEdiscoveryCasesItemTagsAsHierarchyResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewComplianceEdiscoveryCasesItemTagsAsHierarchyResponse instantiates a new ComplianceEdiscoveryCasesItemTagsAsHierarchyResponse and sets the default values.
func NewComplianceEdiscoveryCasesItemTagsAsHierarchyResponse()(*ComplianceEdiscoveryCasesItemTagsAsHierarchyResponse) {
    m := &ComplianceEdiscoveryCasesItemTagsAsHierarchyResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateComplianceEdiscoveryCasesItemTagsAsHierarchyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComplianceEdiscoveryCasesItemTagsAsHierarchyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComplianceEdiscoveryCasesItemTagsAsHierarchyResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComplianceEdiscoveryCasesItemTagsAsHierarchyResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ComplianceEdiscoveryCasesItemTagsAsHierarchyResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
