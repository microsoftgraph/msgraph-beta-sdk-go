package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse 
type ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse instantiates a new ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse and sets the default values.
func NewReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse()(*ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse) {
    m := &ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateDestinationSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.DestinationSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.DestinationSummaryable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse) GetValue()([]i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.DestinationSummaryable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.DestinationSummaryable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse) SetValue(value []i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.DestinationSummaryable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponseable 
type ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.DestinationSummaryable)
    SetValue(value []i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.DestinationSummaryable)()
}
