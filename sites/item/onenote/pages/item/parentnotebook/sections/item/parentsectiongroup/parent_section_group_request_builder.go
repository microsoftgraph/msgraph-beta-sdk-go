package parentsectiongroup

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5fb71e163a45f3cba7a6ee0125d330c3fe472fc96091b7e2780934c610604732 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/pages/item/parentnotebook/sections/item/parentsectiongroup/parentnotebook"
    i73e2fbf1b72d5425245f0cc38e3bb88a34289adca87bf92a2bd3c5d6c28fcb9e "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/pages/item/parentnotebook/sections/item/parentsectiongroup/sectiongroups"
    if8ee80c7c51fb2eabd233dcc652e6395db28e77dfd2078f5aac9219660f19f35 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/pages/item/parentnotebook/sections/item/parentsectiongroup/sections"
    i304056de0e80997976d45376195dd3912aa3902d3a32806585f82b14f40af213 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/pages/item/parentnotebook/sections/item/parentsectiongroup/sectiongroups/item"
    if234efa51c36c0a675ebba8283998d34e3b4de1491aa5a568dcf3c0e2fe864f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/pages/item/parentnotebook/sections/item/parentsectiongroup/sections/item"
)

// ParentSectionGroupRequestBuilder builds and executes requests for operations under \sites\{site-id}\onenote\pages\{onenotePage-id}\parentNotebook\sections\{onenoteSection-id}\parentSectionGroup
type ParentSectionGroupRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ParentSectionGroupRequestBuilderDeleteOptions options for Delete
type ParentSectionGroupRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ParentSectionGroupRequestBuilderGetOptions options for Get
type ParentSectionGroupRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ParentSectionGroupRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ParentSectionGroupRequestBuilderGetQueryParameters the section group that contains the section.  Read-only.
type ParentSectionGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// ParentSectionGroupRequestBuilderPatchOptions options for Patch
type ParentSectionGroupRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SectionGroup;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewParentSectionGroupRequestBuilderInternal instantiates a new ParentSectionGroupRequestBuilder and sets the default values.
func NewParentSectionGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentSectionGroupRequestBuilder) {
    m := &ParentSectionGroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/onenote/pages/{onenotePage_id}/parentNotebook/sections/{onenoteSection_id}/parentSectionGroup{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewParentSectionGroupRequestBuilder instantiates a new ParentSectionGroupRequestBuilder and sets the default values.
func NewParentSectionGroupRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentSectionGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParentSectionGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the section group that contains the section.  Read-only.
func (m *ParentSectionGroupRequestBuilder) CreateDeleteRequestInformation(options *ParentSectionGroupRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the section group that contains the section.  Read-only.
func (m *ParentSectionGroupRequestBuilder) CreateGetRequestInformation(options *ParentSectionGroupRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation the section group that contains the section.  Read-only.
func (m *ParentSectionGroupRequestBuilder) CreatePatchRequestInformation(options *ParentSectionGroupRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete the section group that contains the section.  Read-only.
func (m *ParentSectionGroupRequestBuilder) Delete(options *ParentSectionGroupRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the section group that contains the section.  Read-only.
func (m *ParentSectionGroupRequestBuilder) Get(options *ParentSectionGroupRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SectionGroup, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSectionGroup() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SectionGroup), nil
}
func (m *ParentSectionGroupRequestBuilder) ParentNotebook()(*i5fb71e163a45f3cba7a6ee0125d330c3fe472fc96091b7e2780934c610604732.ParentNotebookRequestBuilder) {
    return i5fb71e163a45f3cba7a6ee0125d330c3fe472fc96091b7e2780934c610604732.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ParentSectionGroupRequestBuilder) ParentSectionGroup()(*ParentSectionGroupRequestBuilder) {
    return NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the section group that contains the section.  Read-only.
func (m *ParentSectionGroupRequestBuilder) Patch(options *ParentSectionGroupRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ParentSectionGroupRequestBuilder) SectionGroups()(*i73e2fbf1b72d5425245f0cc38e3bb88a34289adca87bf92a2bd3c5d6c28fcb9e.SectionGroupsRequestBuilder) {
    return i73e2fbf1b72d5425245f0cc38e3bb88a34289adca87bf92a2bd3c5d6c28fcb9e.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.onenote.pages.item.parentNotebook.sections.item.parentSectionGroup.sectionGroups.item collection
func (m *ParentSectionGroupRequestBuilder) SectionGroupsById(id string)(*i304056de0e80997976d45376195dd3912aa3902d3a32806585f82b14f40af213.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return i304056de0e80997976d45376195dd3912aa3902d3a32806585f82b14f40af213.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ParentSectionGroupRequestBuilder) Sections()(*if8ee80c7c51fb2eabd233dcc652e6395db28e77dfd2078f5aac9219660f19f35.SectionsRequestBuilder) {
    return if8ee80c7c51fb2eabd233dcc652e6395db28e77dfd2078f5aac9219660f19f35.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.onenote.pages.item.parentNotebook.sections.item.parentSectionGroup.sections.item collection
func (m *ParentSectionGroupRequestBuilder) SectionsById(id string)(*if234efa51c36c0a675ebba8283998d34e3b4de1491aa5a568dcf3c0e2fe864f4.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id1"] = id
    }
    return if234efa51c36c0a675ebba8283998d34e3b4de1491aa5a568dcf3c0e2fe864f4.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
