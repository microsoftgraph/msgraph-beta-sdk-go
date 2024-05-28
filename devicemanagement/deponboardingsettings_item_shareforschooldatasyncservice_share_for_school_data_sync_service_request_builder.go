package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilder provides operations to call the shareForSchoolDataSyncService method.
type DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilderInternal instantiates a new DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilder and sets the default values.
func NewDeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilder) {
    m := &DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/shareForSchoolDataSyncService", pathParameters),
    }
    return m
}
// NewDeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilder instantiates a new DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilder and sets the default values.
func NewDeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action shareForSchoolDataSyncService
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilder) Post(ctx context.Context, requestConfiguration *DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action shareForSchoolDataSyncService
// returns a *RequestInformation when successful
func (m *DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilder when successful
func (m *DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilder) WithUrl(rawUrl string)(*DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilder) {
    return NewDeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
