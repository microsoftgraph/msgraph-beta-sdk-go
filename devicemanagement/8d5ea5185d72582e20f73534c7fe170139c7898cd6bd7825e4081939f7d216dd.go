package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder provides operations to manage the cloudCertificationAuthorityLeafCertificate property of the microsoft.graph.deviceManagement entity.
type CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderGetQueryParameters collection of CloudCertificationAuthorityLeafCertificate records associated with account.
type CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderGetQueryParameters
}
// CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderInternal instantiates a new CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder) {
    m := &CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/cloudCertificationAuthorityLeafCertificate/{cloudCertificationAuthorityLeafCertificate%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder instantiates a new CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property cloudCertificationAuthorityLeafCertificate for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get collection of CloudCertificationAuthorityLeafCertificate records associated with account.
// returns a CloudCertificationAuthorityLeafCertificateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityLeafCertificateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudCertificationAuthorityLeafCertificateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityLeafCertificateable), nil
}
// Patch update the navigation property cloudCertificationAuthorityLeafCertificate in deviceManagement
// returns a CloudCertificationAuthorityLeafCertificateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityLeafCertificateable, requestConfiguration *CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityLeafCertificateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudCertificationAuthorityLeafCertificateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityLeafCertificateable), nil
}
// ToDeleteRequestInformation delete navigation property cloudCertificationAuthorityLeafCertificate for deviceManagement
// returns a *RequestInformation when successful
func (m *CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of CloudCertificationAuthorityLeafCertificate records associated with account.
// returns a *RequestInformation when successful
func (m *CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property cloudCertificationAuthorityLeafCertificate in deviceManagement
// returns a *RequestInformation when successful
func (m *CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityLeafCertificateable, requestConfiguration *CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder when successful
func (m *CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder) WithUrl(rawUrl string)(*CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder) {
    return NewCloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
