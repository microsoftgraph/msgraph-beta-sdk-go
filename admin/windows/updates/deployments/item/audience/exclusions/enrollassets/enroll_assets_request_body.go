package enrollassets

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/windowsupdates"
)

// EnrollAssetsRequestBody provides operations to call the enrollAssets method.
type EnrollAssetsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    assets []ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdatableAssetable;
    // 
    updateCategory *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdateCategory;
}
// NewEnrollAssetsRequestBody instantiates a new enrollAssetsRequestBody and sets the default values.
func NewEnrollAssetsRequestBody()(*EnrollAssetsRequestBody) {
    m := &EnrollAssetsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEnrollAssetsRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnrollAssetsRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEnrollAssetsRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnrollAssetsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAssets gets the assets property value. 
func (m *EnrollAssetsRequestBody) GetAssets()([]ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdatableAssetable) {
    if m == nil {
        return nil
    } else {
        return m.assets
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnrollAssetsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.CreateUpdatableAssetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdatableAssetable, len(val))
            for i, v := range val {
                res[i] = v.(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdatableAssetable)
            }
            m.SetAssets(res)
        }
        return nil
    }
    res["updateCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.ParseUpdateCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateCategory(val.(*ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdateCategory))
        }
        return nil
    }
    return res
}
// GetUpdateCategory gets the updateCategory property value. 
func (m *EnrollAssetsRequestBody) GetUpdateCategory()(*ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdateCategory) {
    if m == nil {
        return nil
    } else {
        return m.updateCategory
    }
}
// Serialize serializes information the current object
func (m *EnrollAssetsRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAssets() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssets()))
        for i, v := range m.GetAssets() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("assets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUpdateCategory() != nil {
        cast := (*m.GetUpdateCategory()).String()
        err := writer.WriteStringValue("updateCategory", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnrollAssetsRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAssets sets the assets property value. 
func (m *EnrollAssetsRequestBody) SetAssets(value []ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdatableAssetable)() {
    if m != nil {
        m.assets = value
    }
}
// SetUpdateCategory sets the updateCategory property value. 
func (m *EnrollAssetsRequestBody) SetUpdateCategory(value *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdateCategory)() {
    if m != nil {
        m.updateCategory = value
    }
}
