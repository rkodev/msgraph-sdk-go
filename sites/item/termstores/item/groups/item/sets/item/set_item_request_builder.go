package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i28411d489ad3cd729f1892303dcb2505651bcd5f6acc3e5cf227bc4118eea651 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/relations"
    i571a44d28bde16f052863de7c5cf0d88446d9f63065509564f1a776efa65f403 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/terms"
    i827e50e0e704e89f39c6ad2c9dd28f427f51f016d3cd9a70c29413bf073622a7 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/children"
    ida7529d11c61cb15d10ee4276f7eef55f69d00bec0595927684536c75e82d402 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/parentgroup"
    i193200b3734d30a01f56084d8677802018b1f9f0bbb205294e6e271745ac81a8 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/terms/item"
    ia984aaf284a0f001442f674627bbd07e2de13686abd06c482db59f9e652b98f2 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/children/item"
    ibb602acd9415bff5d888f98c2aba16005dd06d930f81c6a0360e1614b8bc2aaf "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/groups/item/sets/item/relations/item"
)

// SetItemRequestBuilder provides operations to manage the sets property of the microsoft.graph.termStore.group entity.
type SetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SetItemRequestBuilderDeleteOptions options for Delete
type SetItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SetItemRequestBuilderGetOptions options for Get
type SetItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SetItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SetItemRequestBuilderGetQueryParameters all sets under the group in a term [store].
type SetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SetItemRequestBuilderPatchOptions options for Patch
type SetItemRequestBuilderPatchOptions struct {
    // 
    Body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Children the children property
func (m *SetItemRequestBuilder) Children()(*i827e50e0e704e89f39c6ad2c9dd28f427f51f016d3cd9a70c29413bf073622a7.ChildrenRequestBuilder) {
    return i827e50e0e704e89f39c6ad2c9dd28f427f51f016d3cd9a70c29413bf073622a7.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStores.item.groups.item.sets.item.children.item collection
func (m *SetItemRequestBuilder) ChildrenById(id string)(*ia984aaf284a0f001442f674627bbd07e2de13686abd06c482db59f9e652b98f2.TermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did"] = id
    }
    return ia984aaf284a0f001442f674627bbd07e2de13686abd06c482db59f9e652b98f2.NewTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSetItemRequestBuilderInternal instantiates a new SetItemRequestBuilder and sets the default values.
func NewSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetItemRequestBuilder) {
    m := &SetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/termStores/{store%2Did}/groups/{group%2Did}/sets/{set%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSetItemRequestBuilder instantiates a new SetItemRequestBuilder and sets the default values.
func NewSetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property sets for sites
func (m *SetItemRequestBuilder) CreateDeleteRequestInformation(options *SetItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation all sets under the group in a term [store].
func (m *SetItemRequestBuilder) CreateGetRequestInformation(options *SetItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property sets in sites
func (m *SetItemRequestBuilder) CreatePatchRequestInformation(options *SetItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property sets for sites
func (m *SetItemRequestBuilder) Delete(options *SetItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get all sets under the group in a term [store].
func (m *SetItemRequestBuilder) Get(options *SetItemRequestBuilderGetOptions)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateSetFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable), nil
}
// ParentGroup the parentGroup property
func (m *SetItemRequestBuilder) ParentGroup()(*ida7529d11c61cb15d10ee4276f7eef55f69d00bec0595927684536c75e82d402.ParentGroupRequestBuilder) {
    return ida7529d11c61cb15d10ee4276f7eef55f69d00bec0595927684536c75e82d402.NewParentGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property sets in sites
func (m *SetItemRequestBuilder) Patch(options *SetItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Relations the relations property
func (m *SetItemRequestBuilder) Relations()(*i28411d489ad3cd729f1892303dcb2505651bcd5f6acc3e5cf227bc4118eea651.RelationsRequestBuilder) {
    return i28411d489ad3cd729f1892303dcb2505651bcd5f6acc3e5cf227bc4118eea651.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStores.item.groups.item.sets.item.relations.item collection
func (m *SetItemRequestBuilder) RelationsById(id string)(*ibb602acd9415bff5d888f98c2aba16005dd06d930f81c6a0360e1614b8bc2aaf.RelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation%2Did"] = id
    }
    return ibb602acd9415bff5d888f98c2aba16005dd06d930f81c6a0360e1614b8bc2aaf.NewRelationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Terms the terms property
func (m *SetItemRequestBuilder) Terms()(*i571a44d28bde16f052863de7c5cf0d88446d9f63065509564f1a776efa65f403.TermsRequestBuilder) {
    return i571a44d28bde16f052863de7c5cf0d88446d9f63065509564f1a776efa65f403.NewTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStores.item.groups.item.sets.item.terms.item collection
func (m *SetItemRequestBuilder) TermsById(id string)(*i193200b3734d30a01f56084d8677802018b1f9f0bbb205294e6e271745ac81a8.TermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did"] = id
    }
    return i193200b3734d30a01f56084d8677802018b1f9f0bbb205294e6e271745ac81a8.NewTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
