package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Updatesable 
type Updatesable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCatalog()(Catalogable)
    GetDeployments()([]Deploymentable)
    GetUpdatableAssets()([]UpdatableAssetable)
    SetCatalog(value Catalogable)()
    SetDeployments(value []Deploymentable)()
    SetUpdatableAssets(value []UpdatableAssetable)()
}
