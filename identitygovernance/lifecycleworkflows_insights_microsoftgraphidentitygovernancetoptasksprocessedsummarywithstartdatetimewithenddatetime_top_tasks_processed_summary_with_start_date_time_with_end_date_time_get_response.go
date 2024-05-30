package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

type LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewLifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponse instantiates a new LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponse and sets the default values.
func NewLifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponse()(*LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponse) {
    m := &LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateLifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateTopTasksInsightsSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TopTasksInsightsSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TopTasksInsightsSummaryable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []TopTasksInsightsSummaryable when successful
func (m *LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponse) GetValue()([]i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TopTasksInsightsSummaryable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TopTasksInsightsSummaryable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponse) SetValue(value []i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TopTasksInsightsSummaryable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TopTasksInsightsSummaryable)
    SetValue(value []i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TopTasksInsightsSummaryable)()
}
