package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilder provides operations to call the unshareForSchoolDataSyncService method.
type DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilderInternal instantiates a new DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilder and sets the default values.
func NewDeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilder) {
    m := &DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/unshareForSchoolDataSyncService", pathParameters),
    }
    return m
}
// NewDeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilder instantiates a new DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilder and sets the default values.
func NewDeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action unshareForSchoolDataSyncService
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilder) Post(ctx context.Context, requestConfiguration *DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action unshareForSchoolDataSyncService
// returns a *RequestInformation when successful
func (m *DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilder when successful
func (m *DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilder) WithUrl(rawUrl string)(*DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilder) {
    return NewDeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
