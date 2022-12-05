package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemInformationProtectionRequestBuilder provides operations to manage the informationProtection property of the microsoft.graph.user entity.
type UsersItemInformationProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemInformationProtectionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemInformationProtectionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemInformationProtectionRequestBuilderGetQueryParameters get informationProtection from users
type UsersItemInformationProtectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemInformationProtectionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemInformationProtectionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemInformationProtectionRequestBuilderGetQueryParameters
}
// UsersItemInformationProtectionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemInformationProtectionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Bitlocker provides operations to manage the bitlocker property of the microsoft.graph.informationProtection entity.
func (m *UsersItemInformationProtectionRequestBuilder) Bitlocker()(*UsersItemInformationProtectionBitlockerRequestBuilder) {
    return NewUsersItemInformationProtectionBitlockerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUsersItemInformationProtectionRequestBuilderInternal instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewUsersItemInformationProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemInformationProtectionRequestBuilder) {
    m := &UsersItemInformationProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/informationProtection{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemInformationProtectionRequestBuilder instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewUsersItemInformationProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemInformationProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemInformationProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property informationProtection for users
func (m *UsersItemInformationProtectionRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemInformationProtectionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get informationProtection from users
func (m *UsersItemInformationProtectionRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemInformationProtectionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property informationProtection in users
func (m *UsersItemInformationProtectionRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, requestConfiguration *UsersItemInformationProtectionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemInformationProtectionRequestBuilder) DataLossPreventionPolicies()(*UsersItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) {
    return NewUsersItemInformationProtectionDataLossPreventionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DataLossPreventionPoliciesById provides operations to manage the dataLossPreventionPolicies property of the microsoft.graph.informationProtection entity.
func (m *UsersItemInformationProtectionRequestBuilder) DataLossPreventionPoliciesById(id string)(*UsersItemInformationProtectionDataLossPreventionPoliciesDataLossPreventionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataLossPreventionPolicy%2Did"] = id
    }
    return NewUsersItemInformationProtectionDataLossPreventionPoliciesDataLossPreventionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DecryptBuffer provides operations to call the decryptBuffer method.
func (m *UsersItemInformationProtectionRequestBuilder) DecryptBuffer()(*UsersItemInformationProtectionDecryptBufferRequestBuilder) {
    return NewUsersItemInformationProtectionDecryptBufferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property informationProtection for users
func (m *UsersItemInformationProtectionRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemInformationProtectionRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *UsersItemInformationProtectionRequestBuilder) EncryptBuffer()(*UsersItemInformationProtectionEncryptBufferRequestBuilder) {
    return NewUsersItemInformationProtectionEncryptBufferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get informationProtection from users
func (m *UsersItemInformationProtectionRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemInformationProtectionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, error) {
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
// Patch update the navigation property informationProtection in users
func (m *UsersItemInformationProtectionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, requestConfiguration *UsersItemInformationProtectionRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, error) {
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
func (m *UsersItemInformationProtectionRequestBuilder) Policy()(*UsersItemInformationProtectionPolicyRequestBuilder) {
    return NewUsersItemInformationProtectionPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SensitivityLabels provides operations to manage the sensitivityLabels property of the microsoft.graph.informationProtection entity.
func (m *UsersItemInformationProtectionRequestBuilder) SensitivityLabels()(*UsersItemInformationProtectionSensitivityLabelsRequestBuilder) {
    return NewUsersItemInformationProtectionSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SensitivityLabelsById provides operations to manage the sensitivityLabels property of the microsoft.graph.informationProtection entity.
func (m *UsersItemInformationProtectionRequestBuilder) SensitivityLabelsById(id string)(*UsersItemInformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sensitivityLabel%2Did"] = id
    }
    return NewUsersItemInformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SensitivityPolicySettings provides operations to manage the sensitivityPolicySettings property of the microsoft.graph.informationProtection entity.
func (m *UsersItemInformationProtectionRequestBuilder) SensitivityPolicySettings()(*UsersItemInformationProtectionSensitivityPolicySettingsRequestBuilder) {
    return NewUsersItemInformationProtectionSensitivityPolicySettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SignDigest provides operations to call the signDigest method.
func (m *UsersItemInformationProtectionRequestBuilder) SignDigest()(*UsersItemInformationProtectionSignDigestRequestBuilder) {
    return NewUsersItemInformationProtectionSignDigestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreatAssessmentRequests provides operations to manage the threatAssessmentRequests property of the microsoft.graph.informationProtection entity.
func (m *UsersItemInformationProtectionRequestBuilder) ThreatAssessmentRequests()(*UsersItemInformationProtectionThreatAssessmentRequestsRequestBuilder) {
    return NewUsersItemInformationProtectionThreatAssessmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreatAssessmentRequestsById provides operations to manage the threatAssessmentRequests property of the microsoft.graph.informationProtection entity.
func (m *UsersItemInformationProtectionRequestBuilder) ThreatAssessmentRequestsById(id string)(*UsersItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["threatAssessmentRequest%2Did"] = id
    }
    return NewUsersItemInformationProtectionThreatAssessmentRequestsThreatAssessmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// VerifySignature provides operations to call the verifySignature method.
func (m *UsersItemInformationProtectionRequestBuilder) VerifySignature()(*UsersItemInformationProtectionVerifySignatureRequestBuilder) {
    return NewUsersItemInformationProtectionVerifySignatureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
