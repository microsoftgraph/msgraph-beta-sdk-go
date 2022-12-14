package compliance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody provides operations to call the addToReviewSet method.
type EdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The additionalDataOptions property
    additionalDataOptions *ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.AdditionalDataOptions
    // The sourceCollection property
    sourceCollection ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable
}
// NewEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody instantiates a new EdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody and sets the default values.
func NewEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody()(*EdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) {
    m := &EdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAdditionalDataOptions gets the additionalDataOptions property value. The additionalDataOptions property
func (m *EdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) GetAdditionalDataOptions()(*ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.AdditionalDataOptions) {
    return m.additionalDataOptions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *EdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) GetSourceCollection()(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable) {
    return m.sourceCollection
}
// Serialize serializes information the current object
func (m *EdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAdditionalDataOptions sets the additionalDataOptions property value. The additionalDataOptions property
func (m *EdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) SetAdditionalDataOptions(value *ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.AdditionalDataOptions)() {
    m.additionalDataOptions = value
}
// SetSourceCollection sets the sourceCollection property value. The sourceCollection property
func (m *EdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) SetSourceCollection(value ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable)() {
    m.sourceCollection = value
}
