package ediscovery

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AddToReviewSetOperation 
type AddToReviewSetOperation struct {
    CaseOperation
    // The review set to which items matching the source collection query are added to.
    reviewSet ReviewSetable;
    // The sourceCollection that items are being added from.
    sourceCollection SourceCollectionable;
}
// NewAddToReviewSetOperation instantiates a new addToReviewSetOperation and sets the default values.
func NewAddToReviewSetOperation()(*AddToReviewSetOperation) {
    m := &AddToReviewSetOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreateAddToReviewSetOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAddToReviewSetOperationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAddToReviewSetOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddToReviewSetOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    res["reviewSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateReviewSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewSet(val.(ReviewSetable))
        }
        return nil
    }
    res["sourceCollection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *AddToReviewSetOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
