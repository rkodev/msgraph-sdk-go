package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsRequestBuilder provides operations to call the applyTags method.
type CasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsRequestBuilderInternal instantiates a new SecurityApplyTagsRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsRequestBuilder) {
    m := &CasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}/queries/{ediscoveryReviewSetQuery%2Did}/security.applyTags", pathParameters),
    }
    return m
}
// NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsRequestBuilder instantiates a new SecurityApplyTagsRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post apply tags to files in an eDiscovery review set. For details, see Tag documents in a review set in eDiscovery.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/security-ediscoveryreviewsetquery-applytags?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsRequestBuilder) Post(ctx context.Context, body CasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsApplyTagsPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation apply tags to files in an eDiscovery review set. For details, see Tag documents in a review set in eDiscovery.
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsRequestBuilder) ToPostRequestInformation(ctx context.Context, body CasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsApplyTagsPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemSecurityApplyTagsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
