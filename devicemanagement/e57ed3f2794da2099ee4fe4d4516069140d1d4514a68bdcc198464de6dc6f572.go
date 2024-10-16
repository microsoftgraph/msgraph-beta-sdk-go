package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder provides operations to call the getAllCloudCertificationAuthorityLeafCertificates method.
type CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilderInternal instantiates a new CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder) {
    m := &CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/cloudCertificationAuthority/{cloudCertificationAuthority%2Did}/getAllCloudCertificationAuthorityLeafCertificates", pathParameters),
    }
    return m
}
// NewCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder instantiates a new CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getAllCloudCertificationAuthorityLeafCertificates
// Deprecated: This method is obsolete. Use PostAsGetAllCloudCertificationAuthorityLeafCertificatesPostResponse instead.
// returns a CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder) Post(ctx context.Context, requestConfiguration *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilderPostRequestConfiguration)(CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesResponseable), nil
}
// PostAsGetAllCloudCertificationAuthorityLeafCertificatesPostResponse invoke action getAllCloudCertificationAuthorityLeafCertificates
// returns a CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder) PostAsGetAllCloudCertificationAuthorityLeafCertificatesPostResponse(ctx context.Context, requestConfiguration *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilderPostRequestConfiguration)(CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesPostResponseable), nil
}
// ToPostRequestInformation invoke action getAllCloudCertificationAuthorityLeafCertificates
// returns a *RequestInformation when successful
func (m *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder when successful
func (m *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder) WithUrl(rawUrl string)(*CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder) {
    return NewCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
