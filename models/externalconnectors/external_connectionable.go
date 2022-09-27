package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ExternalConnectionable 
type ExternalConnectionable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivitySettings()(ActivitySettingsable)
    GetComplianceSettings()(ComplianceSettingsable)
    GetConfiguration()(Configurationable)
    GetConnectorId()(*string)
    GetDescription()(*string)
    GetEnabledContentExperiences()(*ContentExperienceType)
    GetGroups()([]ExternalGroupable)
    GetIngestedItemsCount()(*int64)
    GetItems()([]ExternalItemable)
    GetName()(*string)
    GetOperations()([]ConnectionOperationable)
    GetQuota()(ConnectionQuotaable)
    GetSchema()(Schemaable)
    GetSearchSettings()(SearchSettingsable)
    GetState()(*ConnectionState)
    SetActivitySettings(value ActivitySettingsable)()
    SetComplianceSettings(value ComplianceSettingsable)()
    SetConfiguration(value Configurationable)()
    SetConnectorId(value *string)()
    SetDescription(value *string)()
    SetEnabledContentExperiences(value *ContentExperienceType)()
    SetGroups(value []ExternalGroupable)()
    SetIngestedItemsCount(value *int64)()
    SetItems(value []ExternalItemable)()
    SetName(value *string)()
    SetOperations(value []ConnectionOperationable)()
    SetQuota(value ConnectionQuotaable)()
    SetSchema(value Schemaable)()
    SetSearchSettings(value SearchSettingsable)()
}
