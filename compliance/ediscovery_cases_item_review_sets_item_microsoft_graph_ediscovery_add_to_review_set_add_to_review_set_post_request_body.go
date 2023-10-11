package compliance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody 
type EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody instantiates a new EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody and sets the default values.
func NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody()(*EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody) {
    m := &EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetAdditionalDataOptions gets the additionalDataOptions property value. The additionalDataOptions property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody) GetAdditionalDataOptions()(*ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.AdditionalDataOptions) {
    val, err := m.GetBackingStore().Get("additionalDataOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.AdditionalDataOptions)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["additionalDataOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ParseAdditionalDataOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalDataOptions(val.(*ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.AdditionalDataOptions))
        }
        return nil
    }
    res["sourceCollection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateSourceCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceCollection(val.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable))
        }
        return nil
    }
    return res
}
// GetSourceCollection gets the sourceCollection property value. The sourceCollection property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody) GetSourceCollection()(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable) {
    val, err := m.GetBackingStore().Get("sourceCollection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAdditionalDataOptions() != nil {
        cast := (*m.GetAdditionalDataOptions()).String()
        err := writer.WriteStringValue("additionalDataOptions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sourceCollection", m.GetSourceCollection())
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
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalDataOptions sets the additionalDataOptions property value. The additionalDataOptions property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody) SetAdditionalDataOptions(value *ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.AdditionalDataOptions)() {
    err := m.GetBackingStore().Set("additionalDataOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetSourceCollection sets the sourceCollection property value. The sourceCollection property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBody) SetSourceCollection(value ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable)() {
    err := m.GetBackingStore().Set("sourceCollection", value)
    if err != nil {
        panic(err)
    }
}
// EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBodyable 
type EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryAddToReviewSetAddToReviewSetPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalDataOptions()(*ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.AdditionalDataOptions)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetSourceCollection()(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable)
    SetAdditionalDataOptions(value *ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.AdditionalDataOptions)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetSourceCollection(value ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable)()
}
