package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilder provides operations to call the patchCloudCertificationAuthority method.
type CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilderInternal instantiates a new CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilder) {
    m := &CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/cloudCertificationAuthority/{cloudCertificationAuthority%2Did}/patchCloudCertificationAuthority", pathParameters),
    }
    return m
}
// NewCloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilder instantiates a new CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action patchCloudCertificationAuthority
// returns a CloudCertificationAuthorityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilder) Post(ctx context.Context, requestConfiguration *CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation invoke action patchCloudCertificationAuthority
// returns a *RequestInformation when successful
func (m *CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilder when successful
func (m *CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilder) WithUrl(rawUrl string)(*CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilder) {
    return NewCloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
