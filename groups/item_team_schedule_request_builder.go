package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTeamScheduleRequestBuilder provides operations to manage the schedule property of the microsoft.graph.team entity.
type ItemTeamScheduleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamScheduleRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamScheduleRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamScheduleRequestBuilderGetQueryParameters retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
type ItemTeamScheduleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamScheduleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamScheduleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamScheduleRequestBuilderGetQueryParameters
}
// ItemTeamScheduleRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamScheduleRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamScheduleRequestBuilderInternal instantiates a new ScheduleRequestBuilder and sets the default values.
func NewItemTeamScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamScheduleRequestBuilder) {
    m := &ItemTeamScheduleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/schedule{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemTeamScheduleRequestBuilder instantiates a new ScheduleRequestBuilder and sets the default values.
func NewItemTeamScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property schedule for groups
func (m *ItemTeamScheduleRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamScheduleRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/schedule-get?view=graph-rest-1.0
func (m *ItemTeamScheduleRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamScheduleRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable), nil
}
// OfferShiftRequests provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) OfferShiftRequests()(*ItemTeamScheduleOfferShiftRequestsRequestBuilder) {
    return NewItemTeamScheduleOfferShiftRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OfferShiftRequestsById provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) OfferShiftRequestsById(id string)(*ItemTeamScheduleOfferShiftRequestsOfferShiftRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest%2Did"] = id
    }
    return NewItemTeamScheduleOfferShiftRequestsOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// OpenShiftChangeRequests provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) OpenShiftChangeRequests()(*ItemTeamScheduleOpenShiftChangeRequestsRequestBuilder) {
    return NewItemTeamScheduleOpenShiftChangeRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OpenShiftChangeRequestsById provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*ItemTeamScheduleOpenShiftChangeRequestsOpenShiftChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest%2Did"] = id
    }
    return NewItemTeamScheduleOpenShiftChangeRequestsOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// OpenShifts provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) OpenShifts()(*ItemTeamScheduleOpenShiftsRequestBuilder) {
    return NewItemTeamScheduleOpenShiftsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OpenShiftsById provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) OpenShiftsById(id string)(*ItemTeamScheduleOpenShiftsOpenShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift%2Did"] = id
    }
    return NewItemTeamScheduleOpenShiftsOpenShiftItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Put update the navigation property schedule in groups
func (m *ItemTeamScheduleRequestBuilder) Put(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, requestConfiguration *ItemTeamScheduleRequestBuilderPutRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable), nil
}
// SchedulingGroups provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) SchedulingGroups()(*ItemTeamScheduleSchedulingGroupsRequestBuilder) {
    return NewItemTeamScheduleSchedulingGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SchedulingGroupsById provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) SchedulingGroupsById(id string)(*ItemTeamScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup%2Did"] = id
    }
    return NewItemTeamScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Share provides operations to call the share method.
func (m *ItemTeamScheduleRequestBuilder) Share()(*ItemTeamScheduleShareRequestBuilder) {
    return NewItemTeamScheduleShareRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Shifts provides operations to manage the shifts property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) Shifts()(*ItemTeamScheduleShiftsRequestBuilder) {
    return NewItemTeamScheduleShiftsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShiftsById provides operations to manage the shifts property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) ShiftsById(id string)(*ItemTeamScheduleShiftsShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift%2Did"] = id
    }
    return NewItemTeamScheduleShiftsShiftItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// SwapShiftsChangeRequests provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) SwapShiftsChangeRequests()(*ItemTeamScheduleSwapShiftsChangeRequestsRequestBuilder) {
    return NewItemTeamScheduleSwapShiftsChangeRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SwapShiftsChangeRequestsById provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*ItemTeamScheduleSwapShiftsChangeRequestsSwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest%2Did"] = id
    }
    return NewItemTeamScheduleSwapShiftsChangeRequestsSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// TimeOffReasons provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) TimeOffReasons()(*ItemTeamScheduleTimeOffReasonsRequestBuilder) {
    return NewItemTeamScheduleTimeOffReasonsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TimeOffReasonsById provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) TimeOffReasonsById(id string)(*ItemTeamScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason%2Did"] = id
    }
    return NewItemTeamScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// TimeOffRequests provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) TimeOffRequests()(*ItemTeamScheduleTimeOffRequestsRequestBuilder) {
    return NewItemTeamScheduleTimeOffRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TimeOffRequestsById provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) TimeOffRequestsById(id string)(*ItemTeamScheduleTimeOffRequestsTimeOffRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest%2Did"] = id
    }
    return NewItemTeamScheduleTimeOffRequestsTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// TimesOff provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) TimesOff()(*ItemTeamScheduleTimesOffRequestBuilder) {
    return NewItemTeamScheduleTimesOffRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TimesOffById provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
func (m *ItemTeamScheduleRequestBuilder) TimesOffById(id string)(*ItemTeamScheduleTimesOffTimeOffItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff%2Did"] = id
    }
    return NewItemTeamScheduleTimesOffTimeOffItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property schedule for groups
func (m *ItemTeamScheduleRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamScheduleRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
func (m *ItemTeamScheduleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPutRequestInformation update the navigation property schedule in groups
func (m *ItemTeamScheduleRequestBuilder) ToPutRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, requestConfiguration *ItemTeamScheduleRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers.Add("Accept", "application/json")
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
