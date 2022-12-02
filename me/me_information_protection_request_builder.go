package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeInformationProtectionRequestBuilder provides operations to manage the informationProtection property of the microsoft.graph.user entity.
type MeInformationProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeInformationProtectionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeInformationProtectionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeInformationProtectionRequestBuilderGetQueryParameters get informationProtection from me
type MeInformationProtectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeInformationProtectionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeInformationProtectionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeInformationProtectionRequestBuilderGetQueryParameters
}
// MeInformationProtectionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeInformationProtectionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Bitlocker provides operations to manage the bitlocker property of the microsoft.graph.informationProtection entity.
func (m *MeInformationProtectionRequestBuilder) Bitlocker()(*MeInformationProtectionBitlockerRequestBuilder) {
    return NewMeInformationProtectionBitlockerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMeInformationProtectionRequestBuilderInternal instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewMeInformationProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeInformationProtectionRequestBuilder) {
    m := &MeInformationProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/informationProtection{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeInformationProtectionRequestBuilder instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewMeInformationProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeInformationProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeInformationProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property informationProtection for me
func (m *MeInformationProtectionRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MeInformationProtectionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get informationProtection from me
func (m *MeInformationProtectionRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeInformationProtectionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property informationProtection in me
func (m *MeInformationProtectionRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, requestConfiguration *MeInformationProtectionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DataLossPreventionPolicies provides operations to manage the dataLossPreventionPolicies property of the microsoft.graph.informationProtection entity.
func (m *MeInformationProtectionRequestBuilder) DataLossPreventionPolicies()(*MeInformationProtectionDataLossPreventionPoliciesRequestBuilder) {
    return NewMeInformationProtectionDataLossPreventionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DataLossPreventionPoliciesById provides operations to manage the dataLossPreventionPolicies property of the microsoft.graph.informationProtection entity.
func (m *MeInformationProtectionRequestBuilder) DataLossPreventionPoliciesById(id string)(*MeInformationProtectionDataLossPreventionPoliciesDataLossPreventionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataLossPreventionPolicy%2Did"] = id
    }
    return NewMeInformationProtectionDataLossPreventionPoliciesDataLossPreventionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DecryptBuffer provides operations to call the decryptBuffer method.
func (m *MeInformationProtectionRequestBuilder) DecryptBuffer()(*MeInformationProtectionDecryptBufferRequestBuilder) {
    return NewMeInformationProtectionDecryptBufferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property informationProtection for me
func (m *MeInformationProtectionRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeInformationProtectionRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// EncryptBuffer provides operations to call the encryptBuffer method.
func (m *MeInformationProtectionRequestBuilder) EncryptBuffer()(*MeInformationProtectionEncryptBufferRequestBuilder) {
    return NewMeInformationProtectionEncryptBufferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get informationProtection from me
func (m *MeInformationProtectionRequestBuilder) Get(ctx context.Context, requestConfiguration *MeInformationProtectionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInformationProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable), nil
}
// Patch update the navigation property informationProtection in me
func (m *MeInformationProtectionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, requestConfiguration *MeInformationProtectionRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInformationProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable), nil
}
// Policy provides operations to manage the policy property of the microsoft.graph.informationProtection entity.
func (m *MeInformationProtectionRequestBuilder) Policy()(*MeInformationProtectionPolicyRequestBuilder) {
    return NewMeInformationProtectionPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SensitivityLabels provides operations to manage the sensitivityLabels property of the microsoft.graph.informationProtection entity.
func (m *MeInformationProtectionRequestBuilder) SensitivityLabels()(*MeInformationProtectionSensitivityLabelsRequestBuilder) {
    return NewMeInformationProtectionSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SensitivityLabelsById provides operations to manage the sensitivityLabels property of the microsoft.graph.informationProtection entity.
func (m *MeInformationProtectionRequestBuilder) SensitivityLabelsById(id string)(*MeInformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sensitivityLabel%2Did"] = id
    }
    return NewMeInformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SensitivityPolicySettings provides operations to manage the sensitivityPolicySettings property of the microsoft.graph.informationProtection entity.
func (m *MeInformationProtectionRequestBuilder) SensitivityPolicySettings()(*MeInformationProtectionSensitivityPolicySettingsRequestBuilder) {
    return NewMeInformationProtectionSensitivityPolicySettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SignDigest provides operations to call the signDigest method.
func (m *MeInformationProtectionRequestBuilder) SignDigest()(*MeInformationProtectionSignDigestRequestBuilder) {
    return NewMeInformationProtectionSignDigestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreatAssessmentRequests provides operations to manage the threatAssessmentRequests property of the microsoft.graph.informationProtection entity.
func (m *MeInformationProtectionRequestBuilder) ThreatAssessmentRequests()(*MeInformationProtectionThreatAssessmentRequestsRequestBuilder) {
    return NewMeInformationProtectionThreatAssessmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreatAssessmentRequestsById provides operations to manage the threatAssessmentRequests property of the microsoft.graph.informationProtection entity.
func (m *MeInformationProtectionRequestBuilder) ThreatAssessmentRequestsById(id string)(*MeInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["threatAssessmentRequest%2Did"] = id
    }
    return NewMeInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// VerifySignature provides operations to call the verifySignature method.
func (m *MeInformationProtectionRequestBuilder) VerifySignature()(*MeInformationProtectionVerifySignatureRequestBuilder) {
    return NewMeInformationProtectionVerifySignatureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
