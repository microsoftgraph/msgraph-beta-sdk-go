package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemInformationprotectionInformationProtectionRequestBuilder provides operations to manage the informationProtection property of the microsoft.graph.site entity.
type ItemSitesItemInformationprotectionInformationProtectionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemInformationprotectionInformationProtectionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationprotectionInformationProtectionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemInformationprotectionInformationProtectionRequestBuilderGetQueryParameters get informationProtection from groups
type ItemSitesItemInformationprotectionInformationProtectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemInformationprotectionInformationProtectionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationprotectionInformationProtectionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemInformationprotectionInformationProtectionRequestBuilderGetQueryParameters
}
// ItemSitesItemInformationprotectionInformationProtectionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationprotectionInformationProtectionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Bitlocker provides operations to manage the bitlocker property of the microsoft.graph.informationProtection entity.
// returns a *ItemSitesItemInformationprotectionBitlockerRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) Bitlocker()(*ItemSitesItemInformationprotectionBitlockerRequestBuilder) {
    return NewItemSitesItemInformationprotectionBitlockerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemInformationprotectionInformationProtectionRequestBuilderInternal instantiates a new ItemSitesItemInformationprotectionInformationProtectionRequestBuilder and sets the default values.
func NewItemSitesItemInformationprotectionInformationProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) {
    m := &ItemSitesItemInformationprotectionInformationProtectionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemInformationprotectionInformationProtectionRequestBuilder instantiates a new ItemSitesItemInformationprotectionInformationProtectionRequestBuilder and sets the default values.
func NewItemSitesItemInformationprotectionInformationProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationprotectionInformationProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// DataLossPreventionPolicies provides operations to manage the dataLossPreventionPolicies property of the microsoft.graph.informationProtection entity.
// returns a *ItemSitesItemInformationprotectionDatalosspreventionpoliciesDataLossPreventionPoliciesRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) DataLossPreventionPolicies()(*ItemSitesItemInformationprotectionDatalosspreventionpoliciesDataLossPreventionPoliciesRequestBuilder) {
    return NewItemSitesItemInformationprotectionDatalosspreventionpoliciesDataLossPreventionPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DecryptBuffer provides operations to call the decryptBuffer method.
// returns a *ItemSitesItemInformationprotectionDecryptbufferDecryptBufferRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) DecryptBuffer()(*ItemSitesItemInformationprotectionDecryptbufferDecryptBufferRequestBuilder) {
    return NewItemSitesItemInformationprotectionDecryptbufferDecryptBufferRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property informationProtection for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemInformationprotectionInformationProtectionRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *ItemSitesItemInformationprotectionEncryptbufferEncryptBufferRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) EncryptBuffer()(*ItemSitesItemInformationprotectionEncryptbufferEncryptBufferRequestBuilder) {
    return NewItemSitesItemInformationprotectionEncryptbufferEncryptBufferRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get informationProtection from groups
// returns a InformationProtectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemInformationprotectionInformationProtectionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, error) {
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
func (m *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, requestConfiguration *ItemSitesItemInformationprotectionInformationProtectionRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, error) {
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
// returns a *ItemSitesItemInformationprotectionPolicyRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) Policy()(*ItemSitesItemInformationprotectionPolicyRequestBuilder) {
    return NewItemSitesItemInformationprotectionPolicyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SensitivityLabels provides operations to manage the sensitivityLabels property of the microsoft.graph.informationProtection entity.
// returns a *ItemSitesItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) SensitivityLabels()(*ItemSitesItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilder) {
    return NewItemSitesItemInformationprotectionSensitivitylabelsSensitivityLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SensitivityPolicySettings provides operations to manage the sensitivityPolicySettings property of the microsoft.graph.informationProtection entity.
// returns a *ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) SensitivityPolicySettings()(*ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) {
    return NewItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SignDigest provides operations to call the signDigest method.
// returns a *ItemSitesItemInformationprotectionSigndigestSignDigestRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) SignDigest()(*ItemSitesItemInformationprotectionSigndigestSignDigestRequestBuilder) {
    return NewItemSitesItemInformationprotectionSigndigestSignDigestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ThreatAssessmentRequests provides operations to manage the threatAssessmentRequests property of the microsoft.graph.informationProtection entity.
// returns a *ItemSitesItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) ThreatAssessmentRequests()(*ItemSitesItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilder) {
    return NewItemSitesItemInformationprotectionThreatassessmentrequestsThreatAssessmentRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property informationProtection for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationprotectionInformationProtectionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationprotectionInformationProtectionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, requestConfiguration *ItemSitesItemInformationprotectionInformationProtectionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemInformationprotectionVerifysignatureVerifySignatureRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) VerifySignature()(*ItemSitesItemInformationprotectionVerifysignatureVerifySignatureRequestBuilder) {
    return NewItemSitesItemInformationprotectionVerifysignatureVerifySignatureRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemInformationprotectionInformationProtectionRequestBuilder) {
    return NewItemSitesItemInformationprotectionInformationProtectionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
