package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3abc3d9218e91a5fafadab38ade57dead754f7030e54ff7bc0cbad14ddab0a7b "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item/roleassignments/item/subject"
    icaba2f6b01e964573246ad4168789811e2c6315b16dfedbc695d1ef6560520a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item/roleassignments/item/linkedeligibleroleassignment"
    iecec41b98ea15ad4fe92aa166baac5bd2638a6c802ca82a30f2929034826fd0f "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item/roleassignments/item/resource"
    iee4dbea39593613c78ec24e53cdf8d2924237ab40a8e7e05550aae33b0bb3a16 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item/roleassignments/item/roledefinition"
)

// GovernanceRoleAssignmentItemRequestBuilder provides operations to manage the roleAssignments property of the microsoft.graph.governanceResource entity.
type GovernanceRoleAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GovernanceRoleAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GovernanceRoleAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GovernanceRoleAssignmentItemRequestBuilderGetQueryParameters the collection of role assignments for the resource.
type GovernanceRoleAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GovernanceRoleAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GovernanceRoleAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GovernanceRoleAssignmentItemRequestBuilderGetQueryParameters
}
// GovernanceRoleAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GovernanceRoleAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGovernanceRoleAssignmentItemRequestBuilderInternal instantiates a new GovernanceRoleAssignmentItemRequestBuilder and sets the default values.
func NewGovernanceRoleAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GovernanceRoleAssignmentItemRequestBuilder) {
    m := &GovernanceRoleAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/governanceResources/{governanceResource%2Did}/roleAssignments/{governanceRoleAssignment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGovernanceRoleAssignmentItemRequestBuilder instantiates a new GovernanceRoleAssignmentItemRequestBuilder and sets the default values.
func NewGovernanceRoleAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GovernanceRoleAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGovernanceRoleAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignments for governanceResources
func (m *GovernanceRoleAssignmentItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property roleAssignments for governanceResources
func (m *GovernanceRoleAssignmentItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *GovernanceRoleAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of role assignments for the resource.
func (m *GovernanceRoleAssignmentItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the collection of role assignments for the resource.
func (m *GovernanceRoleAssignmentItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GovernanceRoleAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleAssignments in governanceResources
func (m *GovernanceRoleAssignmentItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property roleAssignments in governanceResources
func (m *GovernanceRoleAssignmentItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable, requestConfiguration *GovernanceRoleAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property roleAssignments for governanceResources
func (m *GovernanceRoleAssignmentItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property roleAssignments for governanceResources
func (m *GovernanceRoleAssignmentItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *GovernanceRoleAssignmentItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Get the collection of role assignments for the resource.
func (m *GovernanceRoleAssignmentItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the collection of role assignments for the resource.
func (m *GovernanceRoleAssignmentItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GovernanceRoleAssignmentItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceRoleAssignmentFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable), nil
}
// LinkedEligibleRoleAssignment the linkedEligibleRoleAssignment property
func (m *GovernanceRoleAssignmentItemRequestBuilder) LinkedEligibleRoleAssignment()(*icaba2f6b01e964573246ad4168789811e2c6315b16dfedbc695d1ef6560520a3.LinkedEligibleRoleAssignmentRequestBuilder) {
    return icaba2f6b01e964573246ad4168789811e2c6315b16dfedbc695d1ef6560520a3.NewLinkedEligibleRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property roleAssignments in governanceResources
func (m *GovernanceRoleAssignmentItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property roleAssignments in governanceResources
func (m *GovernanceRoleAssignmentItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable, requestConfiguration *GovernanceRoleAssignmentItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Resource the resource property
func (m *GovernanceRoleAssignmentItemRequestBuilder) Resource()(*iecec41b98ea15ad4fe92aa166baac5bd2638a6c802ca82a30f2929034826fd0f.ResourceRequestBuilder) {
    return iecec41b98ea15ad4fe92aa166baac5bd2638a6c802ca82a30f2929034826fd0f.NewResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinition the roleDefinition property
func (m *GovernanceRoleAssignmentItemRequestBuilder) RoleDefinition()(*iee4dbea39593613c78ec24e53cdf8d2924237ab40a8e7e05550aae33b0bb3a16.RoleDefinitionRequestBuilder) {
    return iee4dbea39593613c78ec24e53cdf8d2924237ab40a8e7e05550aae33b0bb3a16.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Subject the subject property
func (m *GovernanceRoleAssignmentItemRequestBuilder) Subject()(*i3abc3d9218e91a5fafadab38ade57dead754f7030e54ff7bc0cbad14ddab0a7b.SubjectRequestBuilder) {
    return i3abc3d9218e91a5fafadab38ade57dead754f7030e54ff7bc0cbad14ddab0a7b.NewSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
