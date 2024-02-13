package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilder provides operations to call the searchExistingIdentities method.
type ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilderInternal instantiates a new ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilder and sets the default values.
func NewImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilder) {
    m := &ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/importedDeviceIdentities/searchExistingIdentities", pathParameters),
    }
    return m
}
// NewImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilder instantiates a new ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilder and sets the default values.
func NewImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action searchExistingIdentities
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a ImportedDeviceIdentitiesSearchExistingIdentitiesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilder) Post(ctx context.Context, body ImportedDeviceIdentitiesSearchExistingIdentitiesPostRequestBodyable, requestConfiguration *ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilderPostRequestConfiguration)(ImportedDeviceIdentitiesSearchExistingIdentitiesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateImportedDeviceIdentitiesSearchExistingIdentitiesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ImportedDeviceIdentitiesSearchExistingIdentitiesResponseable), nil
}
// PostAsSearchExistingIdentitiesPostResponse invoke action searchExistingIdentities
// returns a ImportedDeviceIdentitiesSearchExistingIdentitiesPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilder) PostAsSearchExistingIdentitiesPostResponse(ctx context.Context, body ImportedDeviceIdentitiesSearchExistingIdentitiesPostRequestBodyable, requestConfiguration *ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilderPostRequestConfiguration)(ImportedDeviceIdentitiesSearchExistingIdentitiesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateImportedDeviceIdentitiesSearchExistingIdentitiesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ImportedDeviceIdentitiesSearchExistingIdentitiesPostResponseable), nil
}
// ToPostRequestInformation invoke action searchExistingIdentities
// returns a *RequestInformation when successful
func (m *ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ImportedDeviceIdentitiesSearchExistingIdentitiesPostRequestBodyable, requestConfiguration *ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilder when successful
func (m *ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilder) WithUrl(rawUrl string)(*ImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilder) {
    return NewImportedDeviceIdentitiesSearchExistingIdentitiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
