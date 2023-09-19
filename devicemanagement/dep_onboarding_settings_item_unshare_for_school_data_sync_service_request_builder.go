package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilder provides operations to call the unshareForSchoolDataSyncService method.
type DepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilderInternal instantiates a new UnshareForSchoolDataSyncServiceRequestBuilder and sets the default values.
func NewDepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilder) {
    m := &DepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/unshareForSchoolDataSyncService", pathParameters),
    }
    return m
}
// NewDepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilder instantiates a new UnshareForSchoolDataSyncServiceRequestBuilder and sets the default values.
func NewDepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action unshareForSchoolDataSyncService
func (m *DepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilder) Post(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action unshareForSchoolDataSyncService
func (m *DepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilder) WithUrl(rawUrl string)(*DepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilder) {
    return NewDepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
