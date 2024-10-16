package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder provides operations to call the getAllCloudCertificationAuthority method.
type CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilderInternal instantiates a new CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder) {
    m := &CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/cloudCertificationAuthority/{cloudCertificationAuthority%2Did}/getAllCloudCertificationAuthority", pathParameters),
    }
    return m
}
// NewCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder instantiates a new CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getAllCloudCertificationAuthority
// Deprecated: This method is obsolete. Use PostAsGetAllCloudCertificationAuthorityPostResponse instead.
// returns a CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder) Post(ctx context.Context, requestConfiguration *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilderPostRequestConfiguration)(CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityResponseable), nil
}
// PostAsGetAllCloudCertificationAuthorityPostResponse invoke action getAllCloudCertificationAuthority
// returns a CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder) PostAsGetAllCloudCertificationAuthorityPostResponse(ctx context.Context, requestConfiguration *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilderPostRequestConfiguration)(CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityPostResponseable), nil
}
// ToPostRequestInformation invoke action getAllCloudCertificationAuthority
// returns a *RequestInformation when successful
func (m *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder when successful
func (m *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder) WithUrl(rawUrl string)(*CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder) {
    return NewCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
