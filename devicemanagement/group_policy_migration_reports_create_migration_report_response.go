package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type GroupPolicyMigrationReportsCreateMigrationReportResponse struct {
    GroupPolicyMigrationReportsCreateMigrationReportPostResponse
}
// NewGroupPolicyMigrationReportsCreateMigrationReportResponse instantiates a new GroupPolicyMigrationReportsCreateMigrationReportResponse and sets the default values.
func NewGroupPolicyMigrationReportsCreateMigrationReportResponse()(*GroupPolicyMigrationReportsCreateMigrationReportResponse) {
    m := &GroupPolicyMigrationReportsCreateMigrationReportResponse{
        GroupPolicyMigrationReportsCreateMigrationReportPostResponse: *NewGroupPolicyMigrationReportsCreateMigrationReportPostResponse(),
    }
    return m
}
// CreateGroupPolicyMigrationReportsCreateMigrationReportResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupPolicyMigrationReportsCreateMigrationReportResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyMigrationReportsCreateMigrationReportResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type GroupPolicyMigrationReportsCreateMigrationReportResponseable interface {
    GroupPolicyMigrationReportsCreateMigrationReportPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
