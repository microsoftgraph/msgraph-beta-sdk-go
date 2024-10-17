package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilder provides operations to call the changeCloudCertificationAuthorityStatus method.
type CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilderInternal instantiates a new CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilder) {
    m := &CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/cloudCertificationAuthority/{cloudCertificationAuthority%2Did}/changeCloudCertificationAuthorityStatus", pathParameters),
    }
    return m
}
// NewCloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilder instantiates a new CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action changeCloudCertificationAuthorityStatus
// returns a CloudCertificationAuthorityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilder) Post(ctx context.Context, body CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBodyable, requestConfiguration *CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudCertificationAuthorityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityable), nil
}
// ToPostRequestInformation invoke action changeCloudCertificationAuthorityStatus
// returns a *RequestInformation when successful
func (m *CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusPostRequestBodyable, requestConfiguration *CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilder when successful
func (m *CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilder) WithUrl(rawUrl string)(*CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilder) {
    return NewCloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
