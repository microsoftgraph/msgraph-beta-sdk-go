package enrollassets

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/windowsupdates"
)

// EnrollAssetsRequestBodyable 
type EnrollAssetsRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssets()([]ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdatableAssetable)
    GetUpdateCategory()(*ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdateCategory)
    SetAssets(value []ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdatableAssetable)()
    SetUpdateCategory(value *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdateCategory)()
}
