package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder provides operations to manage the symantecCodeSigningCertificate property of the microsoft.graph.deviceAppManagement entity.
type SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderGetQueryParameters the WinPhone Symantec Code Signing Certificate.
type SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderGetQueryParameters
}
// SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderInternal instantiates a new SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder and sets the default values.
func NewSymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder) {
    m := &SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/symantecCodeSigningCertificate{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewSymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder instantiates a new SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder and sets the default values.
func NewSymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property symantecCodeSigningCertificate for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder) Delete(ctx context.Context, requestConfiguration *SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the WinPhone Symantec Code Signing Certificate.
// returns a SymantecCodeSigningCertificateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder) Get(ctx context.Context, requestConfiguration *SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SymantecCodeSigningCertificateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSymantecCodeSigningCertificateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SymantecCodeSigningCertificateable), nil
}
// Patch update the navigation property symantecCodeSigningCertificate in deviceAppManagement
// returns a SymantecCodeSigningCertificateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SymantecCodeSigningCertificateable, requestConfiguration *SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SymantecCodeSigningCertificateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSymantecCodeSigningCertificateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SymantecCodeSigningCertificateable), nil
}
// ToDeleteRequestInformation delete navigation property symantecCodeSigningCertificate for deviceAppManagement
// returns a *RequestInformation when successful
func (m *SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the WinPhone Symantec Code Signing Certificate.
// returns a *RequestInformation when successful
func (m *SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property symantecCodeSigningCertificate in deviceAppManagement
// returns a *RequestInformation when successful
func (m *SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SymantecCodeSigningCertificateable, requestConfiguration *SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder when successful
func (m *SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder) WithUrl(rawUrl string)(*SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder) {
    return NewSymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
