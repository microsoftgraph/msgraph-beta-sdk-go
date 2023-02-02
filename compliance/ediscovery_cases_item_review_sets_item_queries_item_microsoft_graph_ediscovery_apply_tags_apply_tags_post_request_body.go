package compliance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody 
type EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The tagsToAdd property
    tagsToAdd []ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable
    // The tagsToRemove property
    tagsToRemove []ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable
}
// NewEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody instantiates a new EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody and sets the default values.
func NewEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody()(*EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) {
    m := &EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["tagsToAdd"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable, len(val))
            for i, v := range val {
                res[i] = v.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)
            }
            m.SetTagsToAdd(res)
        }
        return nil
    }
    res["tagsToRemove"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable, len(val))
            for i, v := range val {
                res[i] = v.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)
            }
            m.SetTagsToRemove(res)
        }
        return nil
    }
    return res
}
// GetTagsToAdd gets the tagsToAdd property value. The tagsToAdd property
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) GetTagsToAdd()([]ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable) {
    return m.tagsToAdd
}
// GetTagsToRemove gets the tagsToRemove property value. The tagsToRemove property
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) GetTagsToRemove()([]ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable) {
    return m.tagsToRemove
}
// Serialize serializes information the current object
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetTagsToAdd() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTagsToAdd()))
        for i, v := range m.GetTagsToAdd() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("tagsToAdd", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTagsToRemove() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTagsToRemove()))
        for i, v := range m.GetTagsToRemove() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("tagsToRemove", cast)
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
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTagsToAdd sets the tagsToAdd property value. The tagsToAdd property
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) SetTagsToAdd(value []ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)() {
    m.tagsToAdd = value
}
// SetTagsToRemove sets the tagsToRemove property value. The tagsToRemove property
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) SetTagsToRemove(value []ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)() {
    m.tagsToRemove = value
}
