package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemInformationProtectionRequestBuilder provides operations to manage the informationProtection property of the microsoft.graph.site entity.
type ItemSitesItemInformationProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
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
func (m *ItemSitesItemInformationProtectionRequestBuilder) Bitlocker()(*ItemSitesItemInformationProtectionBitlockerRequestBuilder) {
    return NewItemSitesItemInformationProtectionBitlockerRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemSitesItemInformationProtectionRequestBuilderInternal instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionRequestBuilder) {
    m := &ItemSitesItemInformationProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemInformationProtectionRequestBuilder instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// DataLossPreventionPolicies provides operations to manage the dataLossPreventionPolicies property of the microsoft.graph.informationProtection entity.
func (m *ItemSitesItemInformationProtectionRequestBuilder) DataLossPreventionPolicies()(*ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) {
    return NewItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DataLossPreventionPoliciesById provides operations to manage the dataLossPreventionPolicies property of the microsoft.graph.informationProtection entity.
func (m *ItemSitesItemInformationProtectionRequestBuilder) DataLossPreventionPoliciesById(id string)(*ItemSitesItemInformationProtectionDataLossPreventionPoliciesDataLossPreventionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataLossPreventionPolicy%2Did"] = id
    }
    return NewItemSitesItemInformationProtectionDataLossPreventionPoliciesDataLossPreventionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Delete delete navigation property informationProtection for groups
func (m *ItemSitesItemInformationProtectionRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get informationProtection from groups
func (m *ItemSitesItemInformationProtectionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInformationProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable), nil
}
// MicrosoftGraphDecryptBuffer provides operations to call the decryptBuffer method.
func (m *ItemSitesItemInformationProtectionRequestBuilder) MicrosoftGraphDecryptBuffer()(*ItemSitesItemInformationProtectionMicrosoftGraphDecryptBufferRequestBuilder) {
    return NewItemSitesItemInformationProtectionMicrosoftGraphDecryptBufferRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEncryptBuffer provides operations to call the encryptBuffer method.
func (m *ItemSitesItemInformationProtectionRequestBuilder) MicrosoftGraphEncryptBuffer()(*ItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferRequestBuilder) {
    return NewItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSignDigest provides operations to call the signDigest method.
func (m *ItemSitesItemInformationProtectionRequestBuilder) MicrosoftGraphSignDigest()(*ItemSitesItemInformationProtectionMicrosoftGraphSignDigestRequestBuilder) {
    return NewItemSitesItemInformationProtectionMicrosoftGraphSignDigestRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphVerifySignature provides operations to call the verifySignature method.
func (m *ItemSitesItemInformationProtectionRequestBuilder) MicrosoftGraphVerifySignature()(*ItemSitesItemInformationProtectionMicrosoftGraphVerifySignatureRequestBuilder) {
    return NewItemSitesItemInformationProtectionMicrosoftGraphVerifySignatureRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property informationProtection in groups
func (m *ItemSitesItemInformationProtectionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInformationProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable), nil
}
// Policy provides operations to manage the policy property of the microsoft.graph.informationProtection entity.
func (m *ItemSitesItemInformationProtectionRequestBuilder) Policy()(*ItemSitesItemInformationProtectionPolicyRequestBuilder) {
    return NewItemSitesItemInformationProtectionPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SensitivityLabels provides operations to manage the sensitivityLabels property of the microsoft.graph.informationProtection entity.
func (m *ItemSitesItemInformationProtectionRequestBuilder) SensitivityLabels()(*ItemSitesItemInformationProtectionSensitivityLabelsRequestBuilder) {
    return NewItemSitesItemInformationProtectionSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SensitivityLabelsById provides operations to manage the sensitivityLabels property of the microsoft.graph.informationProtection entity.
func (m *ItemSitesItemInformationProtectionRequestBuilder) SensitivityLabelsById(id string)(*ItemSitesItemInformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sensitivityLabel%2Did"] = id
    }
    return NewItemSitesItemInformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SensitivityPolicySettings provides operations to manage the sensitivityPolicySettings property of the microsoft.graph.informationProtection entity.
func (m *ItemSitesItemInformationProtectionRequestBuilder) SensitivityPolicySettings()(*ItemSitesItemInformationProtectionSensitivityPolicySettingsRequestBuilder) {
    return NewItemSitesItemInformationProtectionSensitivityPolicySettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ThreatAssessmentRequests provides operations to manage the threatAssessmentRequests property of the microsoft.graph.informationProtection entity.
func (m *ItemSitesItemInformationProtectionRequestBuilder) ThreatAssessmentRequests()(*ItemSitesItemInformationProtectionThreatAssessmentRequestsRequestBuilder) {
    return NewItemSitesItemInformationProtectionThreatAssessmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ThreatAssessmentRequestsById provides operations to manage the threatAssessmentRequests property of the microsoft.graph.informationProtection entity.
func (m *ItemSitesItemInformationProtectionRequestBuilder) ThreatAssessmentRequestsById(id string)(*ItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["threatAssessmentRequest%2Did"] = id
    }
    return NewItemSitesItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property informationProtection for groups
func (m *ItemSitesItemInformationProtectionRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get informationProtection from groups
func (m *ItemSitesItemInformationProtectionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property informationProtection in groups
func (m *ItemSitesItemInformationProtectionRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, requestConfiguration *ItemSitesItemInformationProtectionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
