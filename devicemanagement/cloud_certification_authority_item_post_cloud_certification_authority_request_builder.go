package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder provides operations to call the postCloudCertificationAuthority method.
type CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilderInternal instantiates a new CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder) {
    m := &CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/cloudCertificationAuthority/{cloudCertificationAuthority%2Did}/postCloudCertificationAuthority", pathParameters),
    }
    return m
}
// NewCloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder instantiates a new CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action postCloudCertificationAuthority
// Deprecated: This method is obsolete. Use PostAsPostCloudCertificationAuthorityPostResponse instead.
// returns a CloudCertificationAuthorityItemPostCloudCertificationAuthorityResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder) Post(ctx context.Context, requestConfiguration *CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilderPostRequestConfiguration)(CloudCertificationAuthorityItemPostCloudCertificationAuthorityResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCloudCertificationAuthorityItemPostCloudCertificationAuthorityResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CloudCertificationAuthorityItemPostCloudCertificationAuthorityResponseable), nil
}
// PostAsPostCloudCertificationAuthorityPostResponse invoke action postCloudCertificationAuthority
// returns a CloudCertificationAuthorityItemPostCloudCertificationAuthorityPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder) PostAsPostCloudCertificationAuthorityPostResponse(ctx context.Context, requestConfiguration *CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilderPostRequestConfiguration)(CloudCertificationAuthorityItemPostCloudCertificationAuthorityPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCloudCertificationAuthorityItemPostCloudCertificationAuthorityPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CloudCertificationAuthorityItemPostCloudCertificationAuthorityPostResponseable), nil
}
// ToPostRequestInformation invoke action postCloudCertificationAuthority
// returns a *RequestInformation when successful
func (m *CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder when successful
func (m *CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder) WithUrl(rawUrl string)(*CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder) {
    return NewCloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
