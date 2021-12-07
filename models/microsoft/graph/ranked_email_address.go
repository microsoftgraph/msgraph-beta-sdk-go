package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RankedEmailAddress 
type RankedEmailAddress struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The email address.
    address *string;
    // The rank of the email address. A rank is used as a sort key, in relation to the other returned results. A higher rank value corresponds to a more relevant result. Relevance is determined by communication, collaboration, and business relationship signals.
    rank *float64;
}
// NewRankedEmailAddress instantiates a new rankedEmailAddress and sets the default values.
func NewRankedEmailAddress()(*RankedEmailAddress) {
    m := &RankedEmailAddress{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RankedEmailAddress) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAddress gets the address property value. The email address.
func (m *RankedEmailAddress) GetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// GetRank gets the rank property value. The rank of the email address. A rank is used as a sort key, in relation to the other returned results. A higher rank value corresponds to a more relevant result. Relevance is determined by communication, collaboration, and business relationship signals.
func (m *RankedEmailAddress) GetRank()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.rank
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RankedEmailAddress) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
        }
        return nil
    }
    res["rank"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRank(val)
        }
        return nil
    }
    return res
}
func (m *RankedEmailAddress) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RankedEmailAddress) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("rank", m.GetRank())
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
func (m *RankedEmailAddress) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAddress sets the address property value. The email address.
func (m *RankedEmailAddress) SetAddress(value *string)() {
    if m != nil {
        m.address = value
    }
}
// SetRank sets the rank property value. The rank of the email address. A rank is used as a sort key, in relation to the other returned results. A higher rank value corresponds to a more relevant result. Relevance is determined by communication, collaboration, and business relationship signals.
func (m *RankedEmailAddress) SetRank(value *float64)() {
    if m != nil {
        m.rank = value
    }
}
