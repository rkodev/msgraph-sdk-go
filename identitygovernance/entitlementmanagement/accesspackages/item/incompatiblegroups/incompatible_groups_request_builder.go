package incompatiblegroups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i850ab465043fd50cce4302dbefbecdd24e31178ddb33432fe76a447b5dba5a10 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/incompatiblegroups/ref"
    ibba12159ded5828fb15c5a15642da6ffb7a332123249af9cd4f09214004f8c30 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/incompatiblegroups/count"
)

// IncompatibleGroupsRequestBuilder provides operations to manage the incompatibleGroups property of the microsoft.graph.accessPackage entity.
type IncompatibleGroupsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IncompatibleGroupsRequestBuilderGetQueryParameters get incompatibleGroups from identityGovernance
type IncompatibleGroupsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// IncompatibleGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IncompatibleGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IncompatibleGroupsRequestBuilderGetQueryParameters
}
// NewIncompatibleGroupsRequestBuilderInternal instantiates a new IncompatibleGroupsRequestBuilder and sets the default values.
func NewIncompatibleGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IncompatibleGroupsRequestBuilder) {
    m := &IncompatibleGroupsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/incompatibleGroups{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIncompatibleGroupsRequestBuilder instantiates a new IncompatibleGroupsRequestBuilder and sets the default values.
func NewIncompatibleGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IncompatibleGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIncompatibleGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *IncompatibleGroupsRequestBuilder) Count()(*ibba12159ded5828fb15c5a15642da6ffb7a332123249af9cd4f09214004f8c30.CountRequestBuilder) {
    return ibba12159ded5828fb15c5a15642da6ffb7a332123249af9cd4f09214004f8c30.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get incompatibleGroups from identityGovernance
func (m *IncompatibleGroupsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get incompatibleGroups from identityGovernance
func (m *IncompatibleGroupsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *IncompatibleGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get incompatibleGroups from identityGovernance
func (m *IncompatibleGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *IncompatibleGroupsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.GroupCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateGroupCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.GroupCollectionResponseable), nil
}
// Ref the Ref property
func (m *IncompatibleGroupsRequestBuilder) Ref()(*i850ab465043fd50cce4302dbefbecdd24e31178ddb33432fe76a447b5dba5a10.RefRequestBuilder) {
    return i850ab465043fd50cce4302dbefbecdd24e31178ddb33432fe76a447b5dba5a10.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
