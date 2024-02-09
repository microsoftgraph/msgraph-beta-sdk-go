package ediscovery

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AddToReviewSetOperation struct {
    CaseOperation
}
// NewAddToReviewSetOperation instantiates a new AddToReviewSetOperation and sets the default values.
func NewAddToReviewSetOperation()(*AddToReviewSetOperation) {
    m := &AddToReviewSetOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreateAddToReviewSetOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAddToReviewSetOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddToReviewSetOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AddToReviewSetOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    res["reviewSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateReviewSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewSet(val.(ReviewSetable))
        }
        return nil
    }
    res["sourceCollection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSourceCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceCollection(val.(SourceCollectionable))
        }
        return nil
    }
    return res
}
// GetReviewSet gets the reviewSet property value. The review set to which items matching the source collection query are added to.
// returns a ReviewSetable when successful
func (m *AddToReviewSetOperation) GetReviewSet()(ReviewSetable) {
    val, err := m.GetBackingStore().Get("reviewSet")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ReviewSetable)
    }
    return nil
}
// GetSourceCollection gets the sourceCollection property value. The sourceCollection that items are being added from.
// returns a SourceCollectionable when successful
func (m *AddToReviewSetOperation) GetSourceCollection()(SourceCollectionable) {
    val, err := m.GetBackingStore().Get("sourceCollection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SourceCollectionable)
    }
    return nil
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
    err := m.GetBackingStore().Set("reviewSet", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceCollection sets the sourceCollection property value. The sourceCollection that items are being added from.
func (m *AddToReviewSetOperation) SetSourceCollection(value SourceCollectionable)() {
    err := m.GetBackingStore().Set("sourceCollection", value)
    if err != nil {
        panic(err)
    }
}
type AddToReviewSetOperationable interface {
    CaseOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReviewSet()(ReviewSetable)
    GetSourceCollection()(SourceCollectionable)
    SetReviewSet(value ReviewSetable)()
    SetSourceCollection(value SourceCollectionable)()
}
