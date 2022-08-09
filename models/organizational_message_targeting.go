package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalMessageTargeting contains the groups of devices that will be targeted to receive the organizational message. If a device is part of the excluded group, then it will not receive the message, regardless of the device being part of an included group
type OrganizationalMessageTargeting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The groups that will not receive the message. If a user from an excluded group is part of an included group, it will not receive the message
    excludeIds []string
    // The groups that will be targeted and receive the message
    includeIds []string
    // The OdataType property
    odataType *string
    // Indicates the type of targeting
    targetingType *OrganizationalMessageTargetingType
}
// NewOrganizationalMessageTargeting instantiates a new organizationalMessageTargeting and sets the default values.
func NewOrganizationalMessageTargeting()(*OrganizationalMessageTargeting) {
    m := &OrganizationalMessageTargeting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.organizationalMessageTargeting";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOrganizationalMessageTargetingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationalMessageTargetingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationalMessageTargeting(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationalMessageTargeting) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExcludeIds gets the excludeIds property value. The groups that will not receive the message. If a user from an excluded group is part of an included group, it will not receive the message
func (m *OrganizationalMessageTargeting) GetExcludeIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeIds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalMessageTargeting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["excludeIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeIds(res)
        }
        return nil
    }
    res["includeIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeIds(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["targetingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationalMessageTargetingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetingType(val.(*OrganizationalMessageTargetingType))
        }
        return nil
    }
    return res
}
// GetIncludeIds gets the includeIds property value. The groups that will be targeted and receive the message
func (m *OrganizationalMessageTargeting) GetIncludeIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeIds
    }
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageTargeting) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// GetTargetingType gets the targetingType property value. Indicates the type of targeting
func (m *OrganizationalMessageTargeting) GetTargetingType()(*OrganizationalMessageTargetingType) {
    if m == nil {
        return nil
    } else {
        return m.targetingType
    }
}
// Serialize serializes information the current object
func (m *OrganizationalMessageTargeting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExcludeIds() != nil {
        err := writer.WriteCollectionOfStringValues("excludeIds", m.GetExcludeIds())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeIds() != nil {
        err := writer.WriteCollectionOfStringValues("includeIds", m.GetIncludeIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetTargetingType() != nil {
        cast := (*m.GetTargetingType()).String()
        err := writer.WriteStringValue("targetingType", &cast)
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
func (m *OrganizationalMessageTargeting) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExcludeIds sets the excludeIds property value. The groups that will not receive the message. If a user from an excluded group is part of an included group, it will not receive the message
func (m *OrganizationalMessageTargeting) SetExcludeIds(value []string)() {
    if m != nil {
        m.excludeIds = value
    }
}
// SetIncludeIds sets the includeIds property value. The groups that will be targeted and receive the message
func (m *OrganizationalMessageTargeting) SetIncludeIds(value []string)() {
    if m != nil {
        m.includeIds = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageTargeting) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
// SetTargetingType sets the targetingType property value. Indicates the type of targeting
func (m *OrganizationalMessageTargeting) SetTargetingType(value *OrganizationalMessageTargetingType)() {
    if m != nil {
        m.targetingType = value
    }
}
