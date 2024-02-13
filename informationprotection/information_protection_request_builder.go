package informationprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// InformationProtectionRequestBuilder provides operations to manage the informationProtection singleton.
type InformationProtectionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InformationProtectionRequestBuilderGetQueryParameters get informationProtection
type InformationProtectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// InformationProtectionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InformationProtectionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InformationProtectionRequestBuilderGetQueryParameters
}
// InformationProtectionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InformationProtectionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Bitlocker provides operations to manage the bitlocker property of the microsoft.graph.informationProtection entity.
// returns a *BitlockerRequestBuilder when successful
func (m *InformationProtectionRequestBuilder) Bitlocker()(*BitlockerRequestBuilder) {
    return NewBitlockerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewInformationProtectionRequestBuilderInternal instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewInformationProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionRequestBuilder) {
    m := &InformationProtectionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/informationProtection{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewInformationProtectionRequestBuilder instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewInformationProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInformationProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// DataLossPreventionPolicies provides operations to manage the dataLossPreventionPolicies property of the microsoft.graph.informationProtection entity.
// returns a *DataLossPreventionPoliciesRequestBuilder when successful
func (m *InformationProtectionRequestBuilder) DataLossPreventionPolicies()(*DataLossPreventionPoliciesRequestBuilder) {
    return NewDataLossPreventionPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DecryptBuffer provides operations to call the decryptBuffer method.
// returns a *DecryptBufferRequestBuilder when successful
func (m *InformationProtectionRequestBuilder) DecryptBuffer()(*DecryptBufferRequestBuilder) {
    return NewDecryptBufferRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EncryptBuffer provides operations to call the encryptBuffer method.
// returns a *EncryptBufferRequestBuilder when successful
func (m *InformationProtectionRequestBuilder) EncryptBuffer()(*EncryptBufferRequestBuilder) {
    return NewEncryptBufferRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get informationProtection
// returns a InformationProtectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *InformationProtectionRequestBuilder) Get(ctx context.Context, requestConfiguration *InformationProtectionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, error) {
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
// Patch update informationProtection
// returns a InformationProtectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *InformationProtectionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, requestConfiguration *InformationProtectionRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, error) {
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
// returns a *PolicyRequestBuilder when successful
func (m *InformationProtectionRequestBuilder) Policy()(*PolicyRequestBuilder) {
    return NewPolicyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SensitivityLabels provides operations to manage the sensitivityLabels property of the microsoft.graph.informationProtection entity.
// returns a *SensitivityLabelsRequestBuilder when successful
func (m *InformationProtectionRequestBuilder) SensitivityLabels()(*SensitivityLabelsRequestBuilder) {
    return NewSensitivityLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SensitivityPolicySettings provides operations to manage the sensitivityPolicySettings property of the microsoft.graph.informationProtection entity.
// returns a *SensitivityPolicySettingsRequestBuilder when successful
func (m *InformationProtectionRequestBuilder) SensitivityPolicySettings()(*SensitivityPolicySettingsRequestBuilder) {
    return NewSensitivityPolicySettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SignDigest provides operations to call the signDigest method.
// returns a *SignDigestRequestBuilder when successful
func (m *InformationProtectionRequestBuilder) SignDigest()(*SignDigestRequestBuilder) {
    return NewSignDigestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ThreatAssessmentRequests provides operations to manage the threatAssessmentRequests property of the microsoft.graph.informationProtection entity.
// returns a *ThreatAssessmentRequestsRequestBuilder when successful
func (m *InformationProtectionRequestBuilder) ThreatAssessmentRequests()(*ThreatAssessmentRequestsRequestBuilder) {
    return NewThreatAssessmentRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get informationProtection
// returns a *RequestInformation when successful
func (m *InformationProtectionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *InformationProtectionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update informationProtection
// returns a *RequestInformation when successful
func (m *InformationProtectionRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, requestConfiguration *InformationProtectionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/informationProtection", m.BaseRequestBuilder.PathParameters)
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
// returns a *VerifySignatureRequestBuilder when successful
func (m *InformationProtectionRequestBuilder) VerifySignature()(*VerifySignatureRequestBuilder) {
    return NewVerifySignatureRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *InformationProtectionRequestBuilder when successful
func (m *InformationProtectionRequestBuilder) WithUrl(rawUrl string)(*InformationProtectionRequestBuilder) {
    return NewInformationProtectionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
