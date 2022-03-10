package externalconnectors

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ExternalConnectionable 
type ExternalConnectionable interface {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetConfiguration()(Configurationable)
    GetConnectorId()(*string)
    GetDescription()(*string)
    GetGroups()([]ExternalGroupable)
    GetIngestedItemsCount()(*int64)
    GetItems()([]ExternalItemable)
    GetName()(*string)
    GetOperations()([]ConnectionOperationable)
    GetQuota()(ConnectionQuotaable)
    GetSchema()(Schemaable)
    GetSearchSettings()(SearchSettingsable)
    GetState()(*ConnectionState)
    SetConfiguration(value Configurationable)()
    SetConnectorId(value *string)()
    SetDescription(value *string)()
    SetGroups(value []ExternalGroupable)()
    SetIngestedItemsCount(value *int64)()
    SetItems(value []ExternalItemable)()
    SetName(value *string)()
    SetOperations(value []ConnectionOperationable)()
    SetQuota(value ConnectionQuotaable)()
    SetSchema(value Schemaable)()
    SetSearchSettings(value SearchSettingsable)()
    SetState(value *ConnectionState)()
}
