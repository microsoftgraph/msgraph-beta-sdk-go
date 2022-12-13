package tenantrelationships

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBodyable 
type ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetManagementTemplateStepId()(*string)
    GetStatus()(*string)
    GetTenantId()(*string)
    SetManagementTemplateStepId(value *string)()
    SetStatus(value *string)()
    SetTenantId(value *string)()
}
