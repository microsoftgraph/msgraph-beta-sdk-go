// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalPostResponseable instead.
type InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponse struct {
    InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalPostResponse
}
// NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponse instantiates a new InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponse and sets the default values.
func NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponse()(*InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponse) {
    m := &InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponse{
        InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalPostResponse: *NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalPostResponse(),
    }
    return m
}
// CreateInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponse(), nil
}
// Deprecated: This class is obsolete. Use InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalPostResponseable instead.
type InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponseable interface {
    InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
