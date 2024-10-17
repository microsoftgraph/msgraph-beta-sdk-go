package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder provides operations to manage the cloudCertificationAuthority property of the microsoft.graph.deviceManagement entity.
type CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderGetQueryParameters collection of CloudCertificationAuthority records associated with account.
type CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderGetQueryParameters
}
// CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChangeCloudCertificationAuthorityStatus provides operations to call the changeCloudCertificationAuthorityStatus method.
// returns a *CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilder when successful
func (m *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) ChangeCloudCertificationAuthorityStatus()(*CloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilder) {
    return NewCloudCertificationAuthorityItemChangeCloudCertificationAuthorityStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CloudCertificationAuthorityLeafCertificate provides operations to manage the cloudCertificationAuthorityLeafCertificate property of the microsoft.graph.cloudCertificationAuthority entity.
// returns a *CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder when successful
func (m *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) CloudCertificationAuthorityLeafCertificate()(*CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder) {
    return NewCloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderInternal instantiates a new CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) {
    m := &CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/cloudCertificationAuthority/{cloudCertificationAuthority%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder instantiates a new CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property cloudCertificationAuthority for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of CloudCertificationAuthority records associated with account.
// returns a CloudCertificationAuthorityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// GetAllCloudCertificationAuthority provides operations to call the getAllCloudCertificationAuthority method.
// returns a *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder when successful
func (m *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) GetAllCloudCertificationAuthority()(*CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilder) {
    return NewCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetAllCloudCertificationAuthorityLeafCertificates provides operations to call the getAllCloudCertificationAuthorityLeafCertificates method.
// returns a *CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder when successful
func (m *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) GetAllCloudCertificationAuthorityLeafCertificates()(*CloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilder) {
    return NewCloudCertificationAuthorityItemGetAllCloudCertificationAuthorityLeafCertificatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCloudCertificationAuthority provides operations to call the getCloudCertificationAuthority method.
// returns a *CloudCertificationAuthorityItemGetCloudCertificationAuthorityRequestBuilder when successful
func (m *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) GetCloudCertificationAuthority()(*CloudCertificationAuthorityItemGetCloudCertificationAuthorityRequestBuilder) {
    return NewCloudCertificationAuthorityItemGetCloudCertificationAuthorityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property cloudCertificationAuthority in deviceManagement
// returns a CloudCertificationAuthorityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityable, requestConfiguration *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// PatchCloudCertificationAuthority provides operations to call the patchCloudCertificationAuthority method.
// returns a *CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilder when successful
func (m *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) PatchCloudCertificationAuthority()(*CloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilder) {
    return NewCloudCertificationAuthorityItemPatchCloudCertificationAuthorityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PostCloudCertificationAuthority provides operations to call the postCloudCertificationAuthority method.
// returns a *CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder when successful
func (m *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) PostCloudCertificationAuthority()(*CloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilder) {
    return NewCloudCertificationAuthorityItemPostCloudCertificationAuthorityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RevokeCloudCertificationAuthorityCertificate provides operations to call the revokeCloudCertificationAuthorityCertificate method.
// returns a *CloudCertificationAuthorityItemRevokeCloudCertificationAuthorityCertificateRequestBuilder when successful
func (m *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) RevokeCloudCertificationAuthorityCertificate()(*CloudCertificationAuthorityItemRevokeCloudCertificationAuthorityCertificateRequestBuilder) {
    return NewCloudCertificationAuthorityItemRevokeCloudCertificationAuthorityCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RevokeLeafCertificate provides operations to call the revokeLeafCertificate method.
// returns a *CloudCertificationAuthorityItemRevokeLeafCertificateRequestBuilder when successful
func (m *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) RevokeLeafCertificate()(*CloudCertificationAuthorityItemRevokeLeafCertificateRequestBuilder) {
    return NewCloudCertificationAuthorityItemRevokeLeafCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property cloudCertificationAuthority for deviceManagement
// returns a *RequestInformation when successful
func (m *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of CloudCertificationAuthority records associated with account.
// returns a *RequestInformation when successful
func (m *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property cloudCertificationAuthority in deviceManagement
// returns a *RequestInformation when successful
func (m *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityable, requestConfiguration *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UploadExternallySignedCertificationAuthorityCertificate provides operations to call the uploadExternallySignedCertificationAuthorityCertificate method.
// returns a *CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificateRequestBuilder when successful
func (m *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) UploadExternallySignedCertificationAuthorityCertificate()(*CloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificateRequestBuilder) {
    return NewCloudCertificationAuthorityItemUploadExternallySignedCertificationAuthorityCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder when successful
func (m *CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) WithUrl(rawUrl string)(*CloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder) {
    return NewCloudCertificationAuthorityCloudCertificationAuthorityItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
