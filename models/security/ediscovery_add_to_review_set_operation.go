package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryAddToReviewSetOperation provides operations to manage the security singleton.
type EdiscoveryAddToReviewSetOperation struct {
    CaseOperation
    // The reviewSet property
    reviewSet EdiscoveryReviewSetable
    // The search property
    search EdiscoverySearchable
}
// NewEdiscoveryAddToReviewSetOperation instantiates a new ediscoveryAddToReviewSetOperation and sets the default values.
func NewEdiscoveryAddToReviewSetOperation()(*EdiscoveryAddToReviewSetOperation) {
    m := &EdiscoveryAddToReviewSetOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreateEdiscoveryAddToReviewSetOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryAddToReviewSetOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryAddToReviewSetOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryAddToReviewSetOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    res["reviewSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoveryReviewSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewSet(val.(EdiscoveryReviewSetable))
        }
        return nil
    }
    res["search"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoverySearchFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearch(val.(EdiscoverySearchable))
        }
        return nil
    }
    return res
}
// GetReviewSet gets the reviewSet property value. The reviewSet property
func (m *EdiscoveryAddToReviewSetOperation) GetReviewSet()(EdiscoveryReviewSetable) {
    if m == nil {
        return nil
    } else {
        return m.reviewSet
    }
}
// GetSearch gets the search property value. The search property
func (m *EdiscoveryAddToReviewSetOperation) GetSearch()(EdiscoverySearchable) {
    if m == nil {
        return nil
    } else {
        return m.search
    }
}
// Serialize serializes information the current object
func (m *EdiscoveryAddToReviewSetOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("search", m.GetSearch())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetReviewSet sets the reviewSet property value. The reviewSet property
func (m *EdiscoveryAddToReviewSetOperation) SetReviewSet(value EdiscoveryReviewSetable)() {
    if m != nil {
        m.reviewSet = value
    }
}
// SetSearch sets the search property value. The search property
func (m *EdiscoveryAddToReviewSetOperation) SetSearch(value EdiscoverySearchable)() {
    if m != nil {
        m.search = value
    }
}
