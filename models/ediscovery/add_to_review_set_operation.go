package ediscovery

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddToReviewSetOperation 
type AddToReviewSetOperation struct {
    CaseOperation
    // The review set to which items matching the source collection query are added to.
    reviewSet ReviewSetable
    // The sourceCollection that items are being added from.
    sourceCollection SourceCollectionable
}
// NewAddToReviewSetOperation instantiates a new AddToReviewSetOperation and sets the default values.
func NewAddToReviewSetOperation()(*AddToReviewSetOperation) {
    m := &AddToReviewSetOperation{
        CaseOperation: *NewCaseOperation(),
    }
    odataTypeValue := "#microsoft.graph.ediscovery.addToReviewSetOperation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAddToReviewSetOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAddToReviewSetOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddToReviewSetOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddToReviewSetOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    res["reviewSet"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateReviewSetFromDiscriminatorValue , m.SetReviewSet)
    res["sourceCollection"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSourceCollectionFromDiscriminatorValue , m.SetSourceCollection)
    return res
}
// GetReviewSet gets the reviewSet property value. The review set to which items matching the source collection query are added to.
func (m *AddToReviewSetOperation) GetReviewSet()(ReviewSetable) {
    return m.reviewSet
}
// GetSourceCollection gets the sourceCollection property value. The sourceCollection that items are being added from.
func (m *AddToReviewSetOperation) GetSourceCollection()(SourceCollectionable) {
    return m.sourceCollection
}
// Serialize serializes information the current object
func (m *AddToReviewSetOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("reviewSet", m.GetReviewSet())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sourceCollection", m.GetSourceCollection())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetReviewSet sets the reviewSet property value. The review set to which items matching the source collection query are added to.
func (m *AddToReviewSetOperation) SetReviewSet(value ReviewSetable)() {
    m.reviewSet = value
}
// SetSourceCollection sets the sourceCollection property value. The sourceCollection that items are being added from.
func (m *AddToReviewSetOperation) SetSourceCollection(value SourceCollectionable)() {
    m.sourceCollection = value
}
