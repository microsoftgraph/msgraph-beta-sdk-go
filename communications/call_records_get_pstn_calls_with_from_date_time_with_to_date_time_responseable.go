package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f "github.com/microsoftgraph/msgraph-beta-sdk-go/models/callrecords"
)

// CallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeResponseable 
type CallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.PstnCallLogRowable)
    SetValue(value []iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.PstnCallLogRowable)()
}
