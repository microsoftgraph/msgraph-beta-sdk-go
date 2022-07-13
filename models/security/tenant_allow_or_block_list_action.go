package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantAllowOrBlockListAction 
type TenantAllowOrBlockListAction struct {
    // The action property
    action *TenantAllowBlockListAction
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The expirationDateTime property
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The note property
    note *string
    // The results property
    results []TenantAllowBlockListEntryResultable
}
// NewTenantAllowOrBlockListAction instantiates a new tenantAllowOrBlockListAction and sets the default values.
func NewTenantAllowOrBlockListAction()(*TenantAllowOrBlockListAction) {
    m := &TenantAllowOrBlockListAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTenantAllowOrBlockListActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantAllowOrBlockListActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantAllowOrBlockListAction(), nil
}
// GetAction gets the action property value. The action property
func (m *TenantAllowOrBlockListAction) GetAction()(*TenantAllowBlockListAction) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantAllowOrBlockListAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. The expirationDateTime property
func (m *TenantAllowOrBlockListAction) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantAllowOrBlockListAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTenantAllowBlockListAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*TenantAllowBlockListAction))
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["note"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNote(val)
        }
        return nil
    }
    res["results"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantAllowBlockListEntryResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantAllowBlockListEntryResultable, len(val))
            for i, v := range val {
                res[i] = v.(TenantAllowBlockListEntryResultable)
            }
            m.SetResults(res)
        }
        return nil
    }
    return res
}
// GetNote gets the note property value. The note property
func (m *TenantAllowOrBlockListAction) GetNote()(*string) {
    if m == nil {
        return nil
    } else {
        return m.note
    }
}
// GetResults gets the results property value. The results property
func (m *TenantAllowOrBlockListAction) GetResults()([]TenantAllowBlockListEntryResultable) {
    if m == nil {
        return nil
    } else {
        return m.results
    }
}
// Serialize serializes information the current object
func (m *TenantAllowOrBlockListAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err := writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("note", m.GetNote())
        if err != nil {
            return err
        }
    }
    if m.GetResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResults()))
        for i, v := range m.GetResults() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("results", cast)
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
// SetAction sets the action property value. The action property
func (m *TenantAllowOrBlockListAction) SetAction(value *TenantAllowBlockListAction)() {
    if m != nil {
        m.action = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantAllowOrBlockListAction) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The expirationDateTime property
func (m *TenantAllowOrBlockListAction) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetNote sets the note property value. The note property
func (m *TenantAllowOrBlockListAction) SetNote(value *string)() {
    if m != nil {
        m.note = value
    }
}
// SetResults sets the results property value. The results property
func (m *TenantAllowOrBlockListAction) SetResults(value []TenantAllowBlockListEntryResultable)() {
    if m != nil {
        m.results = value
    }
}
