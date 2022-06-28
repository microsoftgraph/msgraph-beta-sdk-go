package ediscovery

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddToReviewSetOperation 
type AddToReviewSetOperation struct {
    CaseOperation
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
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
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAddToReviewSetOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAddToReviewSetOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddToReviewSetOperation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AddToReviewSetOperation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
func (m *AddToReviewSetOperation) GetReviewSet()(ReviewSetable) {
    if m == nil {
        return nil
    } else {
        return m.reviewSet
    }
}
// GetSourceCollection gets the sourceCollection property value. The sourceCollection that items are being added from.
func (m *AddToReviewSetOperation) GetSourceCollection()(SourceCollectionable) {
    if m == nil {
        return nil
    } else {
        return m.sourceCollection
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AddToReviewSetOperation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetReviewSet sets the reviewSet property value. The review set to which items matching the source collection query are added to.
func (m *AddToReviewSetOperation) SetReviewSet(value ReviewSetable)() {
    if m != nil {
        m.reviewSet = value
    }
}
// SetSourceCollection sets the sourceCollection property value. The sourceCollection that items are being added from.
func (m *AddToReviewSetOperation) SetSourceCollection(value SourceCollectionable)() {
    if m != nil {
        m.sourceCollection = value
    }
}
