package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder provides operations to call the exportMobileConfig method.
type DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilderInternal instantiates a new DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder and sets the default values.
func NewDeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder) {
    m := &DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/enrollmentProfiles/{enrollmentProfile%2Did}/exportMobileConfig()", pathParameters),
    }
    return m
}
// NewDeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder instantiates a new DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder and sets the default values.
func NewDeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilderInternal(urlParams, requestAdapter)
}
// Get exports the mobile configuration
// Deprecated: This method is obsolete. Use GetAsExportMobileConfigGetResponse instead.
// returns a DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder) Get(ctx context.Context, requestConfiguration *DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilderGetRequestConfiguration)(DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigResponseable), nil
}
// GetAsExportMobileConfigGetResponse exports the mobile configuration
// returns a DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder) GetAsExportMobileConfigGetResponse(ctx context.Context, requestConfiguration *DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilderGetRequestConfiguration)(DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigGetResponseable), nil
}
// ToGetRequestInformation exports the mobile configuration
// returns a *RequestInformation when successful
func (m *DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder when successful
func (m *DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder) WithUrl(rawUrl string)(*DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder) {
    return NewDeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
