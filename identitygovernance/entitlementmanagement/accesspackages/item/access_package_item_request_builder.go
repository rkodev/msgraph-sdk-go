package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i00a97a45cdd179280c0c26b9ae2b3dfae5b403807c291dab7636b2b17fe890e1 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackagesincompatiblewith"
    i161d6a842b6fc7e35133da25937154a2f211684ebd0fc91e42d9c04c808dcbb9 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/getapplicablepolicyrequirements"
    i34cc2ccdd89a0dc322fc7b60944a65cdd8e38db88f780aaf2b601ca7a951852b "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/catalog"
    ic04b51b80331d0ef154e9a47d3a48addcc3e7870db77e3d10035ccebc69ab2d2 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/incompatibleaccesspackages"
    ie6d4f486c0867dff170afaea36b1c76fdbe68d2685cea59f8b28ed53c73a0f82 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/incompatiblegroups"
    if00f060904d770b29264fd177fd1a312e5b13ff355804e4b535010a6ca327b7e "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/assignmentpolicies"
    i0d939cbe12e283ad757f5173cee0e7b391d9057bb5290e21f5b3564b156341ed "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/assignmentpolicies/item"
    i2d34790c26609e5a034ef3cf47c96ab70e63b650b4f8dd426d00f8699c0776e0 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/incompatiblegroups/item"
    ibfbb41f855b81a5c1c77bd65b370db99999f07d7f021ba23fba4e1cac05c76f6 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackagesincompatiblewith/item"
    ie3628ed2d9d3bb5cbeb9941c961182395063206941d1fa942fe03bb74b75dc66 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/incompatibleaccesspackages/item"
)

// AccessPackageItemRequestBuilder provides operations to manage the accessPackages property of the microsoft.graph.entitlementManagement entity.
type AccessPackageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessPackageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageItemRequestBuilderGetQueryParameters access packages define the collection of resource roles and the policies for which subjects can request or be assigned access to those resources.
type AccessPackageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessPackageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessPackageItemRequestBuilderGetQueryParameters
}
// AccessPackageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackagesIncompatibleWith the accessPackagesIncompatibleWith property
func (m *AccessPackageItemRequestBuilder) AccessPackagesIncompatibleWith()(*i00a97a45cdd179280c0c26b9ae2b3dfae5b403807c291dab7636b2b17fe890e1.AccessPackagesIncompatibleWithRequestBuilder) {
    return i00a97a45cdd179280c0c26b9ae2b3dfae5b403807c291dab7636b2b17fe890e1.NewAccessPackagesIncompatibleWithRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackagesIncompatibleWithById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item.accessPackagesIncompatibleWith.item collection
func (m *AccessPackageItemRequestBuilder) AccessPackagesIncompatibleWithById(id string)(*ibfbb41f855b81a5c1c77bd65b370db99999f07d7f021ba23fba4e1cac05c76f6.AccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did1"] = id
    }
    return ibfbb41f855b81a5c1c77bd65b370db99999f07d7f021ba23fba4e1cac05c76f6.NewAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AssignmentPolicies the assignmentPolicies property
func (m *AccessPackageItemRequestBuilder) AssignmentPolicies()(*if00f060904d770b29264fd177fd1a312e5b13ff355804e4b535010a6ca327b7e.AssignmentPoliciesRequestBuilder) {
    return if00f060904d770b29264fd177fd1a312e5b13ff355804e4b535010a6ca327b7e.NewAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item.assignmentPolicies.item collection
func (m *AccessPackageItemRequestBuilder) AssignmentPoliciesById(id string)(*i0d939cbe12e283ad757f5173cee0e7b391d9057bb5290e21f5b3564b156341ed.AccessPackageAssignmentPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy%2Did"] = id
    }
    return i0d939cbe12e283ad757f5173cee0e7b391d9057bb5290e21f5b3564b156341ed.NewAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Catalog the catalog property
func (m *AccessPackageItemRequestBuilder) Catalog()(*i34cc2ccdd89a0dc322fc7b60944a65cdd8e38db88f780aaf2b601ca7a951852b.CatalogRequestBuilder) {
    return i34cc2ccdd89a0dc322fc7b60944a65cdd8e38db88f780aaf2b601ca7a951852b.NewCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessPackageItemRequestBuilderInternal instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewAccessPackageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageItemRequestBuilder) {
    m := &AccessPackageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageItemRequestBuilder instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewAccessPackageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackages for identityGovernance
func (m *AccessPackageItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property accessPackages for identityGovernance
func (m *AccessPackageItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *AccessPackageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation access packages define the collection of resource roles and the policies for which subjects can request or be assigned access to those resources.
func (m *AccessPackageItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration access packages define the collection of resource roles and the policies for which subjects can request or be assigned access to those resources.
func (m *AccessPackageItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AccessPackageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property accessPackages in identityGovernance
func (m *AccessPackageItemRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property accessPackages in identityGovernance
func (m *AccessPackageItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, requestConfiguration *AccessPackageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property accessPackages for identityGovernance
func (m *AccessPackageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AccessPackageItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get access packages define the collection of resource roles and the policies for which subjects can request or be assigned access to those resources.
func (m *AccessPackageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessPackageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable), nil
}
// GetApplicablePolicyRequirements the getApplicablePolicyRequirements property
func (m *AccessPackageItemRequestBuilder) GetApplicablePolicyRequirements()(*i161d6a842b6fc7e35133da25937154a2f211684ebd0fc91e42d9c04c808dcbb9.GetApplicablePolicyRequirementsRequestBuilder) {
    return i161d6a842b6fc7e35133da25937154a2f211684ebd0fc91e42d9c04c808dcbb9.NewGetApplicablePolicyRequirementsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncompatibleAccessPackages the incompatibleAccessPackages property
func (m *AccessPackageItemRequestBuilder) IncompatibleAccessPackages()(*ic04b51b80331d0ef154e9a47d3a48addcc3e7870db77e3d10035ccebc69ab2d2.IncompatibleAccessPackagesRequestBuilder) {
    return ic04b51b80331d0ef154e9a47d3a48addcc3e7870db77e3d10035ccebc69ab2d2.NewIncompatibleAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncompatibleAccessPackagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item.incompatibleAccessPackages.item collection
func (m *AccessPackageItemRequestBuilder) IncompatibleAccessPackagesById(id string)(*ie3628ed2d9d3bb5cbeb9941c961182395063206941d1fa942fe03bb74b75dc66.AccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did1"] = id
    }
    return ie3628ed2d9d3bb5cbeb9941c961182395063206941d1fa942fe03bb74b75dc66.NewAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// IncompatibleGroups the incompatibleGroups property
func (m *AccessPackageItemRequestBuilder) IncompatibleGroups()(*ie6d4f486c0867dff170afaea36b1c76fdbe68d2685cea59f8b28ed53c73a0f82.IncompatibleGroupsRequestBuilder) {
    return ie6d4f486c0867dff170afaea36b1c76fdbe68d2685cea59f8b28ed53c73a0f82.NewIncompatibleGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncompatibleGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item.incompatibleGroups.item collection
func (m *AccessPackageItemRequestBuilder) IncompatibleGroupsById(id string)(*i2d34790c26609e5a034ef3cf47c96ab70e63b650b4f8dd426d00f8699c0776e0.GroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group%2Did"] = id
    }
    return i2d34790c26609e5a034ef3cf47c96ab70e63b650b4f8dd426d00f8699c0776e0.NewGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property accessPackages in identityGovernance
func (m *AccessPackageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, requestConfiguration *AccessPackageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable), nil
}
