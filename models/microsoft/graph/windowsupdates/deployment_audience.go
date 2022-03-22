package windowsupdates

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// DeploymentAudience 
type DeploymentAudience struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // Specifies the assets to exclude from the audience.
    exclusions []UpdatableAssetable;
    // Specifies the assets to include in the audience.
    members []UpdatableAssetable;
}
// NewDeploymentAudience instantiates a new deploymentAudience and sets the default values.
func NewDeploymentAudience()(*DeploymentAudience) {
    m := &DeploymentAudience{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// CreateDeploymentAudienceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeploymentAudienceFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeploymentAudience(), nil
}
// GetExclusions gets the exclusions property value. Specifies the assets to exclude from the audience.
func (m *DeploymentAudience) GetExclusions()([]UpdatableAssetable) {
    if m == nil {
        return nil
    } else {
        return m.exclusions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeploymentAudience) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exclusions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUpdatableAssetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UpdatableAssetable, len(val))
            for i, v := range val {
                res[i] = v.(UpdatableAssetable)
            }
            m.SetExclusions(res)
        }
        return nil
    }
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUpdatableAssetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UpdatableAssetable, len(val))
            for i, v := range val {
                res[i] = v.(UpdatableAssetable)
            }
            m.SetMembers(res)
        }
        return nil
    }
    return res
}
// GetMembers gets the members property value. Specifies the assets to include in the audience.
func (m *DeploymentAudience) GetMembers()([]UpdatableAssetable) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// Serialize serializes information the current object
func (m *DeploymentAudience) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetExclusions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExclusions()))
        for i, v := range m.GetExclusions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("exclusions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExclusions sets the exclusions property value. Specifies the assets to exclude from the audience.
func (m *DeploymentAudience) SetExclusions(value []UpdatableAssetable)() {
    if m != nil {
        m.exclusions = value
    }
}
// SetMembers sets the members property value. Specifies the assets to include in the audience.
func (m *DeploymentAudience) SetMembers(value []UpdatableAssetable)() {
    if m != nil {
        m.members = value
    }
}
