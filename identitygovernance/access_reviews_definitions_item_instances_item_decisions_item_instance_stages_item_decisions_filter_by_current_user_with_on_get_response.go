package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponse instantiates a new AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponse and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponse()(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponse) {
    m := &AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []AccessReviewInstanceDecisionItemable when successful
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponse) GetValue()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponse) SetValue(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable)
    SetValue(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable)()
}
