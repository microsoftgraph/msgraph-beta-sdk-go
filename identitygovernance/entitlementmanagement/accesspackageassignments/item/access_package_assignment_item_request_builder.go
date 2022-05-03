package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2af0bfbcee7907a04a6c5b1b4818035a6651210ffff9b70fda593206c9d69a95 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/target"
    i3681a850a3cd793fe3c063807af12bdde2ee5a63a2cd6b2033458b57d5e3fa51 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentpolicy"
    i6b87ae9babacd24aa7786e6d45bebb8e19e6210a5805c3a06f65b372589cd0a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentresourceroles"
    ia0acba2babd5b1475401d0b6eaac48aff95dc39f020731cafbfd07a3cd647d18 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage"
    iba6b2e1e9a7782cc4bbca9924b10deee1e26627f7845e0b11bbb91671f323380 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/reprocess"
    id7bb72014a00aae74285bd91b1a862ae93a5130eaa1f7e48ba20472c19da7c73 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentrequests"
    i140a5f9ba5238539c561b0796e993a7f43398107e3e3cc9f8bf919276d650b44 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentresourceroles/item"
    i5252ed1a2033545e7629ee16423aea516ec3f2de5b6f26fa3842c58927181b88 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentrequests/item"
)

// AccessPackageAssignmentItemRequestBuilder provides operations to manage the accessPackageAssignments property of the microsoft.graph.entitlementManagement entity.
type AccessPackageAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessPackageAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageAssignmentItemRequestBuilderGetQueryParameters represents the grant of an access package to a subject (user or group).
type AccessPackageAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessPackageAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessPackageAssignmentItemRequestBuilderGetQueryParameters
}
// AccessPackageAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackage the accessPackage property
func (m *AccessPackageAssignmentItemRequestBuilder) AccessPackage()(*ia0acba2babd5b1475401d0b6eaac48aff95dc39f020731cafbfd07a3cd647d18.AccessPackageRequestBuilder) {
    return ia0acba2babd5b1475401d0b6eaac48aff95dc39f020731cafbfd07a3cd647d18.NewAccessPackageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentPolicy the accessPackageAssignmentPolicy property
func (m *AccessPackageAssignmentItemRequestBuilder) AccessPackageAssignmentPolicy()(*i3681a850a3cd793fe3c063807af12bdde2ee5a63a2cd6b2033458b57d5e3fa51.AccessPackageAssignmentPolicyRequestBuilder) {
    return i3681a850a3cd793fe3c063807af12bdde2ee5a63a2cd6b2033458b57d5e3fa51.NewAccessPackageAssignmentPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentRequests the accessPackageAssignmentRequests property
func (m *AccessPackageAssignmentItemRequestBuilder) AccessPackageAssignmentRequests()(*id7bb72014a00aae74285bd91b1a862ae93a5130eaa1f7e48ba20472c19da7c73.AccessPackageAssignmentRequestsRequestBuilder) {
    return id7bb72014a00aae74285bd91b1a862ae93a5130eaa1f7e48ba20472c19da7c73.NewAccessPackageAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignments.item.accessPackageAssignmentRequests.item collection
func (m *AccessPackageAssignmentItemRequestBuilder) AccessPackageAssignmentRequestsById(id string)(*i5252ed1a2033545e7629ee16423aea516ec3f2de5b6f26fa3842c58927181b88.AccessPackageAssignmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentRequest%2Did"] = id
    }
    return i5252ed1a2033545e7629ee16423aea516ec3f2de5b6f26fa3842c58927181b88.NewAccessPackageAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageAssignmentResourceRoles the accessPackageAssignmentResourceRoles property
func (m *AccessPackageAssignmentItemRequestBuilder) AccessPackageAssignmentResourceRoles()(*i6b87ae9babacd24aa7786e6d45bebb8e19e6210a5805c3a06f65b372589cd0a8.AccessPackageAssignmentResourceRolesRequestBuilder) {
    return i6b87ae9babacd24aa7786e6d45bebb8e19e6210a5805c3a06f65b372589cd0a8.NewAccessPackageAssignmentResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentResourceRolesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignments.item.accessPackageAssignmentResourceRoles.item collection
func (m *AccessPackageAssignmentItemRequestBuilder) AccessPackageAssignmentResourceRolesById(id string)(*i140a5f9ba5238539c561b0796e993a7f43398107e3e3cc9f8bf919276d650b44.AccessPackageAssignmentResourceRoleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentResourceRole%2Did"] = id
    }
    return i140a5f9ba5238539c561b0796e993a7f43398107e3e3cc9f8bf919276d650b44.NewAccessPackageAssignmentResourceRoleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAccessPackageAssignmentItemRequestBuilderInternal instantiates a new AccessPackageAssignmentItemRequestBuilder and sets the default values.
func NewAccessPackageAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageAssignmentItemRequestBuilder) {
    m := &AccessPackageAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageAssignmentItemRequestBuilder instantiates a new AccessPackageAssignmentItemRequestBuilder and sets the default values.
func NewAccessPackageAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property accessPackageAssignments for identityGovernance
func (m *AccessPackageAssignmentItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property accessPackageAssignments for identityGovernance
func (m *AccessPackageAssignmentItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *AccessPackageAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration represents the grant of an access package to a subject (user or group).
func (m *AccessPackageAssignmentItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration represents the grant of an access package to a subject (user or group).
func (m *AccessPackageAssignmentItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AccessPackageAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property accessPackageAssignments in identityGovernance
func (m *AccessPackageAssignmentItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property accessPackageAssignments in identityGovernance
func (m *AccessPackageAssignmentItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentable, requestConfiguration *AccessPackageAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property accessPackageAssignments for identityGovernance
func (m *AccessPackageAssignmentItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *AccessPackageAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property accessPackageAssignments for identityGovernance
func (m *AccessPackageAssignmentItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *AccessPackageAssignmentItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// GetWithResponseHandler represents the grant of an access package to a subject (user or group).
func (m *AccessPackageAssignmentItemRequestBuilder) GetWithResponseHandler(requestConfiguration *AccessPackageAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler represents the grant of an access package to a subject (user or group).
func (m *AccessPackageAssignmentItemRequestBuilder) GetWithResponseHandler(requestConfiguration *AccessPackageAssignmentItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentable), nil
}
// PatchWithResponseHandler update the navigation property accessPackageAssignments in identityGovernance
func (m *AccessPackageAssignmentItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentable, requestConfiguration *AccessPackageAssignmentItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property accessPackageAssignments in identityGovernance
func (m *AccessPackageAssignmentItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentable, requestConfiguration *AccessPackageAssignmentItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Reprocess the reprocess property
func (m *AccessPackageAssignmentItemRequestBuilder) Reprocess()(*iba6b2e1e9a7782cc4bbca9924b10deee1e26627f7845e0b11bbb91671f323380.ReprocessRequestBuilder) {
    return iba6b2e1e9a7782cc4bbca9924b10deee1e26627f7845e0b11bbb91671f323380.NewReprocessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Target the target property
func (m *AccessPackageAssignmentItemRequestBuilder) Target()(*i2af0bfbcee7907a04a6c5b1b4818035a6651210ffff9b70fda593206c9d69a95.TargetRequestBuilder) {
    return i2af0bfbcee7907a04a6c5b1b4818035a6651210ffff9b70fda593206c9d69a95.NewTargetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
