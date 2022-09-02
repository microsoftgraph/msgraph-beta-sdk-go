package threatsubmission

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i44f875f20192781d7d5d7e4a796fac48b4e14878f9923e86dd403c7edfa045f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/threatsubmission/urlthreats"
    i73a9697fcc2f3912806f4fa3ce0e1bcb5e2e53d5a6779a4fe4ada90a825e98f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/threatsubmission/filethreats"
    iab730a3e5e523ceacfc51a80389a408f3b763d8578932292a4d6486fa46c88e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/threatsubmission/emailthreats"
    if956c41d927fdbd93476646ddf92e2e53ea39cd255c6014671d49ed36a05be59 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/threatsubmission/emailthreatsubmissionpolicies"
    i1cc1873bdc105e1d79034a659428dc9e6f9cacb1aeb842585395edb7e83b9cee "github.com/microsoftgraph/msgraph-beta-sdk-go/security/threatsubmission/emailthreats/item"
    i5f1a4923ee76d1cf452c42ff9b166ea23067d5576e64a4b629b5ce1aa976f6b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/threatsubmission/urlthreats/item"
    ib5a0370dad339fb08a1a2b4167f005a9e8a48731fdadcbe94317bbe45c759d9d "github.com/microsoftgraph/msgraph-beta-sdk-go/security/threatsubmission/emailthreatsubmissionpolicies/item"
    iea4954cd2624be64ac3de2ebf2b2820af5b41742fb51a218c4440447feabbb05 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/threatsubmission/filethreats/item"
)

// ThreatSubmissionRequestBuilder provides operations to manage the threatSubmission property of the microsoft.graph.security entity.
type ThreatSubmissionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ThreatSubmissionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatSubmissionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ThreatSubmissionRequestBuilderGetQueryParameters a threat submission sent to Microsoft; for example, a suspicious email threat, URL threat, or file threat.
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
    m.urlTemplate = "{+baseurl}/security/threatSubmission{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property threatSubmission for security
func (m *ThreatSubmissionRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property threatSubmission for security
func (m *ThreatSubmissionRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ThreatSubmissionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a threat submission sent to Microsoft; for example, a suspicious email threat, URL threat, or file threat.
func (m *ThreatSubmissionRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration a threat submission sent to Microsoft; for example, a suspicious email threat, URL threat, or file threat.
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
// CreatePatchRequestInformation update the navigation property threatSubmission in security
func (m *ThreatSubmissionRequestBuilder) CreatePatchRequestInformation(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property threatSubmission in security
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
// Delete delete navigation property threatSubmission for security
func (m *ThreatSubmissionRequestBuilder) Delete(ctx context.Context, requestConfiguration *ThreatSubmissionRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// EmailThreats the emailThreats property
func (m *ThreatSubmissionRequestBuilder) EmailThreats()(*iab730a3e5e523ceacfc51a80389a408f3b763d8578932292a4d6486fa46c88e1.EmailThreatsRequestBuilder) {
    return iab730a3e5e523ceacfc51a80389a408f3b763d8578932292a4d6486fa46c88e1.NewEmailThreatsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailThreatsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.threatSubmission.emailThreats.item collection
func (m *ThreatSubmissionRequestBuilder) EmailThreatsById(id string)(*i1cc1873bdc105e1d79034a659428dc9e6f9cacb1aeb842585395edb7e83b9cee.EmailThreatSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["emailThreatSubmission%2Did"] = id
    }
    return i1cc1873bdc105e1d79034a659428dc9e6f9cacb1aeb842585395edb7e83b9cee.NewEmailThreatSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// EmailThreatSubmissionPolicies the emailThreatSubmissionPolicies property
func (m *ThreatSubmissionRequestBuilder) EmailThreatSubmissionPolicies()(*if956c41d927fdbd93476646ddf92e2e53ea39cd255c6014671d49ed36a05be59.EmailThreatSubmissionPoliciesRequestBuilder) {
    return if956c41d927fdbd93476646ddf92e2e53ea39cd255c6014671d49ed36a05be59.NewEmailThreatSubmissionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailThreatSubmissionPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.threatSubmission.emailThreatSubmissionPolicies.item collection
func (m *ThreatSubmissionRequestBuilder) EmailThreatSubmissionPoliciesById(id string)(*ib5a0370dad339fb08a1a2b4167f005a9e8a48731fdadcbe94317bbe45c759d9d.EmailThreatSubmissionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["emailThreatSubmissionPolicy%2Did"] = id
    }
    return ib5a0370dad339fb08a1a2b4167f005a9e8a48731fdadcbe94317bbe45c759d9d.NewEmailThreatSubmissionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FileThreats the fileThreats property
func (m *ThreatSubmissionRequestBuilder) FileThreats()(*i73a9697fcc2f3912806f4fa3ce0e1bcb5e2e53d5a6779a4fe4ada90a825e98f5.FileThreatsRequestBuilder) {
    return i73a9697fcc2f3912806f4fa3ce0e1bcb5e2e53d5a6779a4fe4ada90a825e98f5.NewFileThreatsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FileThreatsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.threatSubmission.fileThreats.item collection
func (m *ThreatSubmissionRequestBuilder) FileThreatsById(id string)(*iea4954cd2624be64ac3de2ebf2b2820af5b41742fb51a218c4440447feabbb05.FileThreatSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["fileThreatSubmission%2Did"] = id
    }
    return iea4954cd2624be64ac3de2ebf2b2820af5b41742fb51a218c4440447feabbb05.NewFileThreatSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get a threat submission sent to Microsoft; for example, a suspicious email threat, URL threat, or file threat.
func (m *ThreatSubmissionRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatSubmissionRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateThreatSubmissionRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable), nil
}
// Patch update the navigation property threatSubmission in security
func (m *ThreatSubmissionRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatSubmissionRootable, requestConfiguration *ThreatSubmissionRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// UrlThreats the urlThreats property
func (m *ThreatSubmissionRequestBuilder) UrlThreats()(*i44f875f20192781d7d5d7e4a796fac48b4e14878f9923e86dd403c7edfa045f8.UrlThreatsRequestBuilder) {
    return i44f875f20192781d7d5d7e4a796fac48b4e14878f9923e86dd403c7edfa045f8.NewUrlThreatsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UrlThreatsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.threatSubmission.urlThreats.item collection
func (m *ThreatSubmissionRequestBuilder) UrlThreatsById(id string)(*i5f1a4923ee76d1cf452c42ff9b166ea23067d5576e64a4b629b5ce1aa976f6b2.UrlThreatSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["urlThreatSubmission%2Did"] = id
    }
    return i5f1a4923ee76d1cf452c42ff9b166ea23067d5576e64a4b629b5ce1aa976f6b2.NewUrlThreatSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
