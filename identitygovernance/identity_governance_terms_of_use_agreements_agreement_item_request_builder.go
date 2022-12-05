package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder provides operations to manage the agreements property of the microsoft.graph.termsOfUseContainer entity.
type IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderGetQueryParameters represents a tenant's customizable terms of use agreement that's created and managed with Azure Active Directory (Azure AD).
type IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderGetQueryParameters
}
// IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Acceptances provides operations to manage the acceptances property of the microsoft.graph.agreement entity.
func (m *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder) Acceptances()(*IdentityGovernanceTermsOfUseAgreementsItemAcceptancesRequestBuilder) {
    return NewIdentityGovernanceTermsOfUseAgreementsItemAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcceptancesById provides operations to manage the acceptances property of the microsoft.graph.agreement entity.
func (m *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder) AcceptancesById(id string)(*IdentityGovernanceTermsOfUseAgreementsItemAcceptancesAgreementAcceptanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance%2Did"] = id
    }
    return NewIdentityGovernanceTermsOfUseAgreementsItemAcceptancesAgreementAcceptanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderInternal instantiates a new AgreementItemRequestBuilder and sets the default values.
func NewIdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder) {
    m := &IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/termsOfUse/agreements/{agreement%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder instantiates a new AgreementItemRequestBuilder and sets the default values.
func NewIdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property agreements for identityGovernance
func (m *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation represents a tenant's customizable terms of use agreement that's created and managed with Azure Active Directory (Azure AD).
func (m *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property agreements in identityGovernance
func (m *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Agreementable, requestConfiguration *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property agreements for identityGovernance
func (m *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// File provides operations to manage the file property of the microsoft.graph.agreement entity.
func (m *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder) File()(*IdentityGovernanceTermsOfUseAgreementsItemFileRequestBuilder) {
    return NewIdentityGovernanceTermsOfUseAgreementsItemFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Files provides operations to manage the files property of the microsoft.graph.agreement entity.
func (m *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder) Files()(*IdentityGovernanceTermsOfUseAgreementsItemFilesRequestBuilder) {
    return NewIdentityGovernanceTermsOfUseAgreementsItemFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesById provides operations to manage the files property of the microsoft.graph.agreement entity.
func (m *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder) FilesById(id string)(*IdentityGovernanceTermsOfUseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementFileLocalization%2Did"] = id
    }
    return NewIdentityGovernanceTermsOfUseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get represents a tenant's customizable terms of use agreement that's created and managed with Azure Active Directory (Azure AD).
func (m *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Agreementable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAgreementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Agreementable), nil
}
// Patch update the navigation property agreements in identityGovernance
func (m *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Agreementable, requestConfiguration *IdentityGovernanceTermsOfUseAgreementsAgreementItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Agreementable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAgreementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Agreementable), nil
}
