package threatsubmission

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i839b7ce1a3161f2b8802e112d0f8dcf56ba63102027be1d561b73ba8add1e4d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/threatsubmission/filethreats"
    ia521c4b7c7b153efe6f509a08ab3716be2eada40c7fc1403d39369608784180e "github.com/microsoftgraph/msgraph-beta-sdk-go/threatsubmission/emailthreatsubmissionpolicies"
    icb54636bfde3eca5af86d2c7097ebddd5013dfd9b34d52115fc1508cc0711117 "github.com/microsoftgraph/msgraph-beta-sdk-go/threatsubmission/urlthreats"
    ide2b52488ab580fd95e34b4b7b93698b9dcd9676a2360c62d4b7b1cb2e4083be "github.com/microsoftgraph/msgraph-beta-sdk-go/threatsubmission/emailthreats"
    i0ad7e6cb7a41139cbe17cb23b2d6a9e0470437b2b26bebd90ece9c5d68cc6097 "github.com/microsoftgraph/msgraph-beta-sdk-go/threatsubmission/emailthreatsubmissionpolicies/item"
    i7b75ce8bd8b760a9c342833943cdc0ede5c3293076bca0d3248d39af8b6b1d37 "github.com/microsoftgraph/msgraph-beta-sdk-go/threatsubmission/emailthreats/item"
    i7d0964b12524c2873a7a2125a01e7b576f0db892b9ceeaa5fca68319cff98e86 "github.com/microsoftgraph/msgraph-beta-sdk-go/threatsubmission/urlthreats/item"
    i9b322beb715796b7bc65df518a61a4b597828f6be25783bc8d137a7e3650a2cd "github.com/microsoftgraph/msgraph-beta-sdk-go/threatsubmission/filethreats/item"
)

// ThreatSubmissionRequestBuilder provides operations to manage the threatSubmissionRoot singleton.
type ThreatSubmissionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ThreatSubmissionRequestBuilderGetQueryParameters get threatSubmission
type ThreatSubmissionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatSubmissionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatSubmissionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatSubmissionRequestBuilderGetQueryParameters
}
// ThreatSubmissionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatSubmissionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewThreatSubmissionRequestBuilderInternal instantiates a new ThreatSubmissionRequestBuilder and sets the default values.
func NewThreatSubmissionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatSubmissionRequestBuilder) {
    m := &ThreatSubmissionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/threatSubmission{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewThreatSubmissionRequestBuilder instantiates a new ThreatSubmissionRequestBuilder and sets the default values.
func NewThreatSubmissionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatSubmissionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatSubmissionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get threatSubmission
func (m *ThreatSubmissionRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get threatSubmission
func (m *ThreatSubmissionRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ThreatSubmissionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update threatSubmission
func (m *ThreatSubmissionRequestBuilder) CreatePatchRequestInformation(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update threatSubmission
func (m *ThreatSubmissionRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, requestConfiguration *ThreatSubmissionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// EmailThreats the emailThreats property
func (m *ThreatSubmissionRequestBuilder) EmailThreats()(*ide2b52488ab580fd95e34b4b7b93698b9dcd9676a2360c62d4b7b1cb2e4083be.EmailThreatsRequestBuilder) {
    return ide2b52488ab580fd95e34b4b7b93698b9dcd9676a2360c62d4b7b1cb2e4083be.NewEmailThreatsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailThreatsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.threatSubmission.emailThreats.item collection
func (m *ThreatSubmissionRequestBuilder) EmailThreatsById(id string)(*i7b75ce8bd8b760a9c342833943cdc0ede5c3293076bca0d3248d39af8b6b1d37.EmailThreatSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["emailThreatSubmission%2Did"] = id
    }
    return i7b75ce8bd8b760a9c342833943cdc0ede5c3293076bca0d3248d39af8b6b1d37.NewEmailThreatSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// EmailThreatSubmissionPolicies the emailThreatSubmissionPolicies property
func (m *ThreatSubmissionRequestBuilder) EmailThreatSubmissionPolicies()(*ia521c4b7c7b153efe6f509a08ab3716be2eada40c7fc1403d39369608784180e.EmailThreatSubmissionPoliciesRequestBuilder) {
    return ia521c4b7c7b153efe6f509a08ab3716be2eada40c7fc1403d39369608784180e.NewEmailThreatSubmissionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailThreatSubmissionPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.threatSubmission.emailThreatSubmissionPolicies.item collection
func (m *ThreatSubmissionRequestBuilder) EmailThreatSubmissionPoliciesById(id string)(*i0ad7e6cb7a41139cbe17cb23b2d6a9e0470437b2b26bebd90ece9c5d68cc6097.EmailThreatSubmissionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["emailThreatSubmissionPolicy%2Did"] = id
    }
    return i0ad7e6cb7a41139cbe17cb23b2d6a9e0470437b2b26bebd90ece9c5d68cc6097.NewEmailThreatSubmissionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FileThreats the fileThreats property
func (m *ThreatSubmissionRequestBuilder) FileThreats()(*i839b7ce1a3161f2b8802e112d0f8dcf56ba63102027be1d561b73ba8add1e4d1.FileThreatsRequestBuilder) {
    return i839b7ce1a3161f2b8802e112d0f8dcf56ba63102027be1d561b73ba8add1e4d1.NewFileThreatsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FileThreatsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.threatSubmission.fileThreats.item collection
func (m *ThreatSubmissionRequestBuilder) FileThreatsById(id string)(*i9b322beb715796b7bc65df518a61a4b597828f6be25783bc8d137a7e3650a2cd.FileThreatSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["fileThreatSubmission%2Did"] = id
    }
    return i9b322beb715796b7bc65df518a61a4b597828f6be25783bc8d137a7e3650a2cd.NewFileThreatSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get threatSubmission
func (m *ThreatSubmissionRequestBuilder) Get()(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get threatSubmission
func (m *ThreatSubmissionRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ThreatSubmissionRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateThreatSubmissionRootFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable), nil
}
// Patch update threatSubmission
func (m *ThreatSubmissionRequestBuilder) Patch(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update threatSubmission
func (m *ThreatSubmissionRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, requestConfiguration *ThreatSubmissionRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// UrlThreats the urlThreats property
func (m *ThreatSubmissionRequestBuilder) UrlThreats()(*icb54636bfde3eca5af86d2c7097ebddd5013dfd9b34d52115fc1508cc0711117.UrlThreatsRequestBuilder) {
    return icb54636bfde3eca5af86d2c7097ebddd5013dfd9b34d52115fc1508cc0711117.NewUrlThreatsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UrlThreatsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.threatSubmission.urlThreats.item collection
func (m *ThreatSubmissionRequestBuilder) UrlThreatsById(id string)(*i7d0964b12524c2873a7a2125a01e7b576f0db892b9ceeaa5fca68319cff98e86.UrlThreatSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["urlThreatSubmission%2Did"] = id
    }
    return i7d0964b12524c2873a7a2125a01e7b576f0db892b9ceeaa5fca68319cff98e86.NewUrlThreatSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
