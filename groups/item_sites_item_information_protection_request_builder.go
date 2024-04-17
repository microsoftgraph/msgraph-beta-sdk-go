package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemInformationProtectionRequestBuilder provides operations to manage the informationProtection property of the microsoft.graph.site entity.
type ItemSitesItemInformationProtectionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemInformationProtectionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemInformationProtectionRequestBuilderGetQueryParameters get informationProtection from groups
type ItemSitesItemInformationProtectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemInformationProtectionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemInformationProtectionRequestBuilderGetQueryParameters
}
// ItemSitesItemInformationProtectionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Bitlocker provides operations to manage the bitlocker property of the microsoft.graph.informationProtection entity.
// returns a *ItemSitesItemInformationProtectionBitlockerRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionRequestBuilder) Bitlocker()(*ItemSitesItemInformationProtectionBitlockerRequestBuilder) {
    return NewItemSitesItemInformationProtectionBitlockerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemInformationProtectionRequestBuilderInternal instantiates a new ItemSitesItemInformationProtectionRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionRequestBuilder) {
    m := &ItemSitesItemInformationProtectionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemInformationProtectionRequestBuilder instantiates a new ItemSitesItemInformationProtectionRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// DataLossPreventionPolicies provides operations to manage the dataLossPreventionPolicies property of the microsoft.graph.informationProtection entity.
// returns a *ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionRequestBuilder) DataLossPreventionPolicies()(*ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) {
    return NewItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DecryptBuffer provides operations to call the decryptBuffer method.
// returns a *ItemSitesItemInformationProtectionDecryptBufferRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionRequestBuilder) DecryptBuffer()(*ItemSitesItemInformationProtectionDecryptBufferRequestBuilder) {
    return NewItemSitesItemInformationProtectionDecryptBufferRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property informationProtection for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationProtectionRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderDeleteRequestConfiguration)(error) {
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
// EncryptBuffer provides operations to call the encryptBuffer method.
// returns a *ItemSitesItemInformationProtectionEncryptBufferRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionRequestBuilder) EncryptBuffer()(*ItemSitesItemInformationProtectionEncryptBufferRequestBuilder) {
    return NewItemSitesItemInformationProtectionEncryptBufferRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get informationProtection from groups
// returns a InformationProtectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationProtectionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInformationProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable), nil
}
// Patch update the navigation property informationProtection in groups
// returns a InformationProtectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationProtectionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInformationProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable), nil
}
// Policy provides operations to manage the policy property of the microsoft.graph.informationProtection entity.
// returns a *ItemSitesItemInformationProtectionPolicyRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionRequestBuilder) Policy()(*ItemSitesItemInformationProtectionPolicyRequestBuilder) {
    return NewItemSitesItemInformationProtectionPolicyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SensitivityLabels provides operations to manage the sensitivityLabels property of the microsoft.graph.informationProtection entity.
// returns a *ItemSitesItemInformationProtectionSensitivityLabelsRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionRequestBuilder) SensitivityLabels()(*ItemSitesItemInformationProtectionSensitivityLabelsRequestBuilder) {
    return NewItemSitesItemInformationProtectionSensitivityLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SensitivityPolicySettings provides operations to manage the sensitivityPolicySettings property of the microsoft.graph.informationProtection entity.
// returns a *ItemSitesItemInformationProtectionSensitivityPolicySettingsRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionRequestBuilder) SensitivityPolicySettings()(*ItemSitesItemInformationProtectionSensitivityPolicySettingsRequestBuilder) {
    return NewItemSitesItemInformationProtectionSensitivityPolicySettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SignDigest provides operations to call the signDigest method.
// returns a *ItemSitesItemInformationProtectionSignDigestRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionRequestBuilder) SignDigest()(*ItemSitesItemInformationProtectionSignDigestRequestBuilder) {
    return NewItemSitesItemInformationProtectionSignDigestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ThreatAssessmentRequests provides operations to manage the threatAssessmentRequests property of the microsoft.graph.informationProtection entity.
// returns a *ItemSitesItemInformationProtectionThreatAssessmentRequestsRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionRequestBuilder) ThreatAssessmentRequests()(*ItemSitesItemInformationProtectionThreatAssessmentRequestsRequestBuilder) {
    return NewItemSitesItemInformationProtectionThreatAssessmentRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property informationProtection for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationProtectionRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get informationProtection from groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationProtectionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property informationProtection in groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationProtectionRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// VerifySignature provides operations to call the verifySignature method.
// returns a *ItemSitesItemInformationProtectionVerifySignatureRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionRequestBuilder) VerifySignature()(*ItemSitesItemInformationProtectionVerifySignatureRequestBuilder) {
    return NewItemSitesItemInformationProtectionVerifySignatureRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemInformationProtectionRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemInformationProtectionRequestBuilder) {
    return NewItemSitesItemInformationProtectionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
