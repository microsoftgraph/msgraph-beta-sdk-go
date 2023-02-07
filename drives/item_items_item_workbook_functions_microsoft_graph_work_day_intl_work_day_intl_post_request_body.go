package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The days property
    days ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The holidays property
    holidays ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The startDate property
    startDate ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The weekend property
    weekend ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDays gets the days property value. The days property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) GetDays()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.days
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["days"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDays(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["holidays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHolidays(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["startDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDate(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["weekend"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeekend(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    return res
}
// GetHolidays gets the holidays property value. The holidays property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) GetHolidays()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.holidays
}
// GetStartDate gets the startDate property value. The startDate property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) GetStartDate()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.startDate
}
// GetWeekend gets the weekend property value. The weekend property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) GetWeekend()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.weekend
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("days", m.GetDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("holidays", m.GetHolidays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("startDate", m.GetStartDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("weekend", m.GetWeekend())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDays sets the days property value. The days property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) SetDays(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.days = value
}
// SetHolidays sets the holidays property value. The holidays property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) SetHolidays(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.holidays = value
}
// SetStartDate sets the startDate property value. The startDate property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) SetStartDate(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.startDate = value
}
// SetWeekend sets the weekend property value. The weekend property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) SetWeekend(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.weekend = value
}
