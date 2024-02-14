package compliance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

type EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody instantiates a new EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody and sets the default values.
func NewEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody()(*EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) {
    m := &EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
                if v != nil {
                    res[i] = v.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)
                }
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
                if v != nil {
                    res[i] = v.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)
                }
            }
            m.SetTagsToRemove(res)
        }
        return nil
    }
    return res
}
// GetTagsToAdd gets the tagsToAdd property value. The tagsToAdd property
// returns a []Tagable when successful
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) GetTagsToAdd()([]ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable) {
    val, err := m.GetBackingStore().Get("tagsToAdd")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)
    }
    return nil
}
// GetTagsToRemove gets the tagsToRemove property value. The tagsToRemove property
// returns a []Tagable when successful
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) GetTagsToRemove()([]ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable) {
    val, err := m.GetBackingStore().Get("tagsToRemove")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetTagsToAdd() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTagsToAdd()))
        for i, v := range m.GetTagsToAdd() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("tagsToAdd", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTagsToRemove() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTagsToRemove()))
        for i, v := range m.GetTagsToRemove() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetTagsToAdd sets the tagsToAdd property value. The tagsToAdd property
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) SetTagsToAdd(value []ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)() {
    err := m.GetBackingStore().Set("tagsToAdd", value)
    if err != nil {
        panic(err)
    }
}
// SetTagsToRemove sets the tagsToRemove property value. The tagsToRemove property
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBody) SetTagsToRemove(value []ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)() {
    err := m.GetBackingStore().Set("tagsToRemove", value)
    if err != nil {
        panic(err)
    }
}
type EdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphEdiscoveryApplyTagsApplyTagsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetTagsToAdd()([]ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)
    GetTagsToRemove()([]ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetTagsToAdd(value []ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)()
    SetTagsToRemove(value []ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable)()
}
