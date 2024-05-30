package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilder provides operations to call the setDefaultProfile method.
type DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilderInternal instantiates a new DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilder and sets the default values.
func NewDeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilder) {
    m := &DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/enrollmentProfiles/{enrollmentProfile%2Did}/setDefaultProfile", pathParameters),
    }
    return m
}
// NewDeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilder instantiates a new DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilder and sets the default values.
func NewDeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action setDefaultProfile
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilder) Post(ctx context.Context, requestConfiguration *DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action setDefaultProfile
// returns a *RequestInformation when successful
func (m *DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilder when successful
func (m *DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilder) WithUrl(rawUrl string)(*DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilder) {
    return NewDeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
