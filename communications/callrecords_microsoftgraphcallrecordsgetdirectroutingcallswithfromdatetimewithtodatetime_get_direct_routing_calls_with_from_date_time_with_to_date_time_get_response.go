package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f "github.com/microsoftgraph/msgraph-beta-sdk-go/models/callrecords"
)

type CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewCallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponse instantiates a new CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponse and sets the default values.
func NewCallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponse()(*CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponse) {
    m := &CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateCallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CreateDirectRoutingLogRowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.DirectRoutingLogRowable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.DirectRoutingLogRowable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []DirectRoutingLogRowable when successful
func (m *CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponse) GetValue()([]iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.DirectRoutingLogRowable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.DirectRoutingLogRowable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponse) SetValue(value []iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.DirectRoutingLogRowable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.DirectRoutingLogRowable)
    SetValue(value []iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.DirectRoutingLogRowable)()
}
