package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder provides operations to call the queryByPlatformType method.
type ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilderInternal instantiates a new ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder and sets the default values.
func NewResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder) {
    m := &ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/resourceAccessProfiles/queryByPlatformType", pathParameters),
    }
    return m
}
// NewResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder instantiates a new ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder and sets the default values.
func NewResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action queryByPlatformType
// Deprecated: This method is obsolete. Use PostAsQueryByPlatformTypePostResponse instead.
// returns a ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder) Post(ctx context.Context, body ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypePostRequestBodyable, requestConfiguration *ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilderPostRequestConfiguration)(ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeResponseable), nil
}
// PostAsQueryByPlatformTypePostResponse invoke action queryByPlatformType
// returns a ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder) PostAsQueryByPlatformTypePostResponse(ctx context.Context, body ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypePostRequestBodyable, requestConfiguration *ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilderPostRequestConfiguration)(ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypePostResponseable), nil
}
// ToPostRequestInformation invoke action queryByPlatformType
// returns a *RequestInformation when successful
func (m *ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder) ToPostRequestInformation(ctx context.Context, body ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypePostRequestBodyable, requestConfiguration *ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder when successful
func (m *ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder) WithUrl(rawUrl string)(*ResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder) {
    return NewResourceaccessprofilesQuerybyplatformtypeQueryByPlatformTypeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
