package ibrowser

import (
	"github.com/chromedp/cdproto"
	"github.com/chromedp/cdproto/accessibility"
	"github.com/chromedp/cdproto/animation"
	"github.com/chromedp/cdproto/applicationcache"
	"github.com/chromedp/cdproto/audits"
	"github.com/chromedp/cdproto/backgroundservice"
	"github.com/chromedp/cdproto/browser"
	"github.com/chromedp/cdproto/cachestorage"
	"github.com/chromedp/cdproto/cast"
	"github.com/chromedp/cdproto/css"
	"github.com/chromedp/cdproto/database"
	"github.com/chromedp/cdproto/debugger"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/domdebugger"
	"github.com/chromedp/cdproto/domsnapshot"
	"github.com/chromedp/cdproto/domstorage"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/fetch"
	"github.com/chromedp/cdproto/headlessexperimental"
	"github.com/chromedp/cdproto/heapprofiler"
	"github.com/chromedp/cdproto/indexeddb"
	"github.com/chromedp/cdproto/inspector"
	"github.com/chromedp/cdproto/io"
	"github.com/chromedp/cdproto/layertree"
	"github.com/chromedp/cdproto/log"
	"github.com/chromedp/cdproto/media"
	"github.com/chromedp/cdproto/memory"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/overlay"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/performance"
	"github.com/chromedp/cdproto/profiler"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/cdproto/security"
	"github.com/chromedp/cdproto/serviceworker"
	"github.com/chromedp/cdproto/storage"
	"github.com/chromedp/cdproto/systeminfo"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/cdproto/tethering"
	"github.com/chromedp/cdproto/tracing"
	"github.com/chromedp/cdproto/webaudio"
	"github.com/chromedp/cdproto/webauthn"
	"github.com/sirupsen/logrus"
)

func Message(ev interface{}) {
	cssOn := false
	defaultVale := false
	switch ev.(type) {
	case *cdproto.Message:
		defaultVale = true
		if ev.(*cdproto.Message).Method.String() != "" {
			logrus.Debugfp("", ev)
		}
	case *accessibility.GetPartialAXTreeReturns:
		logrus.Debugln(cdproto.CommandAccessibilityGetPartialAXTree)
	case *accessibility.GetFullAXTreeReturns:
		logrus.Debugln(cdproto.CommandAccessibilityGetFullAXTree)
	case *animation.GetCurrentTimeReturns:
		logrus.Debugln(cdproto.CommandAnimationGetCurrentTime)
	case *animation.GetPlaybackRateReturns:
		logrus.Debugln(cdproto.CommandAnimationGetPlaybackRate)
	case *animation.ResolveAnimationReturns:
		logrus.Debugln(cdproto.CommandAnimationResolveAnimation)
	case *animation.EventAnimationCanceled:
		logrus.Debugln(cdproto.EventAnimationAnimationCanceled)
	case *animation.EventAnimationCreated:
		logrus.Debugln(cdproto.EventAnimationAnimationCreated)
	case *animation.EventAnimationStarted:
		logrus.Debugln(cdproto.EventAnimationAnimationStarted)
	case *applicationcache.GetApplicationCacheForFrameReturns:
		logrus.Debugln(cdproto.CommandApplicationCacheGetApplicationCacheForFrame)
	case *applicationcache.GetFramesWithManifestsReturns:
		logrus.Debugln(cdproto.CommandApplicationCacheGetFramesWithManifests)
	case *applicationcache.GetManifestForFrameReturns:
		logrus.Debugln(cdproto.CommandApplicationCacheGetManifestForFrame)
	case *applicationcache.EventApplicationCacheStatusUpdated:
		logrus.Debugln(cdproto.EventApplicationCacheApplicationCacheStatusUpdated)
	case *applicationcache.EventNetworkStateUpdated:
		logrus.Debugln(cdproto.EventApplicationCacheNetworkStateUpdated)
	case *audits.GetEncodedResponseReturns:
		logrus.Debugln(cdproto.CommandAuditsGetEncodedResponse)
	case *backgroundservice.EventRecordingStateChanged:
		logrus.Debugln(cdproto.EventBackgroundServiceRecordingStateChanged)
	case *backgroundservice.EventBackgroundServiceEventReceived:
		logrus.Debugln(cdproto.EventBackgroundServiceBackgroundServiceEventReceived)
	case *browser.GetVersionReturns:
		logrus.Debugln(cdproto.CommandBrowserGetVersion)
	case *browser.GetBrowserCommandLineReturns:
		logrus.Debugln(cdproto.CommandBrowserGetBrowserCommandLine)
	case *browser.GetHistogramsReturns:
		logrus.Debugln(cdproto.CommandBrowserGetHistograms)
	case *browser.GetHistogramReturns:
		logrus.Debugln(cdproto.CommandBrowserGetHistogram)
	case *browser.GetWindowBoundsReturns:
		logrus.Debugln(cdproto.CommandBrowserGetWindowBounds)
	case *browser.GetWindowForTargetReturns:
		logrus.Debugln(cdproto.CommandBrowserGetWindowForTarget)
	case *cachestorage.RequestCacheNamesReturns:
		logrus.Debugln(cdproto.CommandCacheStorageRequestCacheNames)
	case *cachestorage.RequestCachedResponseReturns:
		logrus.Debugln(cdproto.CommandCacheStorageRequestCachedResponse)
	case *cachestorage.RequestEntriesReturns:
		logrus.Debugln(cdproto.CommandCacheStorageRequestEntries)
	case *cast.EventSinksUpdated:
		logrus.Debugln(cdproto.EventCastSinksUpdated)
	case *cast.EventIssueUpdated:
		logrus.Debugln(cdproto.EventCastIssueUpdated)
	case *dom.CollectClassNamesFromSubtreeReturns:
		logrus.Debugln(cdproto.CommandDOMCollectClassNamesFromSubtree)
	case *dom.CopyToReturns:
		logrus.Debugln(cdproto.CommandDOMCopyTo)
	case *dom.DescribeNodeReturns:
		logrus.Debugln(cdproto.CommandDOMDescribeNode)
	case *dom.GetAttributesReturns:
		logrus.Debugln(cdproto.CommandDOMGetAttributes)
	case *dom.GetBoxModelReturns:
		logrus.Debugln(cdproto.CommandDOMGetBoxModel)
	case *dom.GetContentQuadsReturns:
		logrus.Debugln(cdproto.CommandDOMGetContentQuads)
	case *dom.GetDocumentReturns:
		logrus.Debugln(cdproto.CommandDOMGetDocument)
	case *dom.GetFlattenedDocumentReturns:
		logrus.Debugln(cdproto.CommandDOMGetFlattenedDocument)
	case *dom.GetNodeForLocationReturns:
		logrus.Debugln(cdproto.CommandDOMGetNodeForLocation)
	case *dom.GetOuterHTMLReturns:
		logrus.Debugln(cdproto.CommandDOMGetOuterHTML)
	case *dom.GetRelayoutBoundaryReturns:
		logrus.Debugln(cdproto.CommandDOMGetRelayoutBoundary)
	case *dom.GetSearchResultsReturns:
		logrus.Debugln(cdproto.CommandDOMGetSearchResults)
	case *dom.MoveToReturns:
		logrus.Debugln(cdproto.CommandDOMMoveTo)
	case *dom.PerformSearchReturns:
		logrus.Debugln(cdproto.CommandDOMPerformSearch)
	case *dom.PushNodeByPathToFrontendReturns:
		logrus.Debugln(cdproto.CommandDOMPushNodeByPathToFrontend)
	case *dom.PushNodesByBackendIdsToFrontendReturns:
		logrus.Debugln(cdproto.CommandDOMPushNodesByBackendIdsToFrontend)
	case *dom.QuerySelectorReturns:
		logrus.Debugln(cdproto.CommandDOMQuerySelector)
	case *dom.QuerySelectorAllReturns:
		logrus.Debugln(cdproto.CommandDOMQuerySelectorAll)
	case *dom.RequestNodeReturns:
		logrus.Debugln(cdproto.CommandDOMRequestNode)
	case *dom.ResolveNodeReturns:
		logrus.Debugln(cdproto.CommandDOMResolveNode)
	case *dom.GetNodeStackTracesReturns:
		logrus.Debugln(cdproto.CommandDOMGetNodeStackTraces)
	case *dom.GetFileInfoReturns:
		logrus.Debugln(cdproto.CommandDOMGetFileInfo)
	case *dom.SetNodeNameReturns:
		logrus.Debugln(cdproto.CommandDOMSetNodeName)
	case *dom.GetFrameOwnerReturns:
		logrus.Debugln(cdproto.CommandDOMGetFrameOwner)
	case *dom.EventAttributeModified:
		logrus.Debugln(cdproto.EventDOMAttributeModified)
	case *dom.EventAttributeRemoved:
		logrus.Debugln(cdproto.EventDOMAttributeRemoved)
	case *dom.EventCharacterDataModified:
		logrus.Debugln(cdproto.EventDOMCharacterDataModified)
	case *dom.EventChildNodeCountUpdated:
		logrus.Debugln(cdproto.EventDOMChildNodeCountUpdated)
	case *dom.EventChildNodeInserted:
		logrus.Debugln(cdproto.EventDOMChildNodeInserted)
	case *dom.EventChildNodeRemoved:
		logrus.Debugln(cdproto.EventDOMChildNodeRemoved)
	case *dom.EventDistributedNodesUpdated:
		logrus.Debugln(cdproto.EventDOMDistributedNodesUpdated)
	case *dom.EventDocumentUpdated:
		logrus.Debugln(cdproto.EventDOMDocumentUpdated)
	case *dom.EventInlineStyleInvalidated:
		logrus.Debugln(cdproto.EventDOMInlineStyleInvalidated)
	case *dom.EventPseudoElementAdded:
		logrus.Debugln(cdproto.EventDOMPseudoElementAdded)
	case *dom.EventPseudoElementRemoved:
		logrus.Debugln(cdproto.EventDOMPseudoElementRemoved)
	case *dom.EventSetChildNodes:
		logrus.Debugln(cdproto.EventDOMSetChildNodes)
	case *dom.EventShadowRootPopped:
		logrus.Debugln(cdproto.EventDOMShadowRootPopped)
	case *dom.EventShadowRootPushed:
		logrus.Debugln(cdproto.EventDOMShadowRootPushed)
	case *domdebugger.GetEventListenersReturns:
		logrus.Debugln(cdproto.CommandDOMDebuggerGetEventListeners)
	case *domsnapshot.CaptureSnapshotReturns:
		logrus.Debugln(cdproto.CommandDOMSnapshotCaptureSnapshot)
	case *domstorage.GetDOMStorageItemsReturns:
		logrus.Debugln(cdproto.CommandDOMStorageGetDOMStorageItems)
	case *domstorage.EventDomStorageItemAdded:
		logrus.Debugln(cdproto.EventDOMStorageDomStorageItemAdded)
	case *domstorage.EventDomStorageItemRemoved:
		logrus.Debugln(cdproto.EventDOMStorageDomStorageItemRemoved)
	case *domstorage.EventDomStorageItemUpdated:
		logrus.Debugln(cdproto.EventDOMStorageDomStorageItemUpdated)
	case *domstorage.EventDomStorageItemsCleared:
		logrus.Debugln(cdproto.EventDOMStorageDomStorageItemsCleared)
	case *database.ExecuteSQLReturns:
		logrus.Debugln(cdproto.CommandDatabaseExecuteSQL)
	case *database.GetDatabaseTableNamesReturns:
		logrus.Debugln(cdproto.CommandDatabaseGetDatabaseTableNames)
	case *database.EventAddDatabase:
		logrus.Debugln(cdproto.EventDatabaseAddDatabase)
	case *debugger.EnableReturns:
		logrus.Debugln(cdproto.CommandDebuggerEnable)
	case *debugger.EvaluateOnCallFrameReturns:
		logrus.Debugln(cdproto.CommandDebuggerEvaluateOnCallFrame)
	case *debugger.GetPossibleBreakpointsReturns:
		logrus.Debugln(cdproto.CommandDebuggerGetPossibleBreakpoints)
	case *debugger.GetScriptSourceReturns:
		logrus.Debugln(cdproto.CommandDebuggerGetScriptSource)
	case *debugger.GetStackTraceReturns:
		logrus.Debugln(cdproto.CommandDebuggerGetStackTrace)
	case *debugger.RestartFrameReturns:
		logrus.Debugln(cdproto.CommandDebuggerRestartFrame)
	case *debugger.SearchInContentReturns:
		logrus.Debugln(cdproto.CommandDebuggerSearchInContent)
	case *debugger.SetBreakpointReturns:
		logrus.Debugln(cdproto.CommandDebuggerSetBreakpoint)
	case *debugger.SetInstrumentationBreakpointReturns:
		logrus.Debugln(cdproto.CommandDebuggerSetInstrumentationBreakpoint)
	case *debugger.SetBreakpointByURLReturns:
		logrus.Debugln(cdproto.CommandDebuggerSetBreakpointByURL)
	case *debugger.SetBreakpointOnFunctionCallReturns:
		logrus.Debugln(cdproto.CommandDebuggerSetBreakpointOnFunctionCall)
	case *debugger.SetScriptSourceReturns:
		logrus.Debugln(cdproto.CommandDebuggerSetScriptSource)
	case *debugger.EventBreakpointResolved:
		logrus.Debugln(cdproto.EventDebuggerBreakpointResolved)
	case *debugger.EventPaused:
		logrus.Debugln(cdproto.EventDebuggerPaused)
	case *debugger.EventResumed:
		logrus.Debugln(cdproto.EventDebuggerResumed)
	case *debugger.EventScriptFailedToParse:
		logrus.Debugln(cdproto.EventDebuggerScriptFailedToParse)
	case *debugger.EventScriptParsed:
		logrus.Debugln(cdproto.EventDebuggerScriptParsed)
	case *emulation.CanEmulateReturns:
		logrus.Debugln(cdproto.CommandEmulationCanEmulate)
	case *emulation.SetVirtualTimePolicyReturns:
		logrus.Debugln(cdproto.CommandEmulationSetVirtualTimePolicy)
	case *emulation.EventVirtualTimeBudgetExpired:
		logrus.Debugln(cdproto.EventEmulationVirtualTimeBudgetExpired)
	case *fetch.GetResponseBodyReturns:
		logrus.Debugln(cdproto.CommandFetchGetResponseBody)
	case *fetch.TakeResponseBodyAsStreamReturns:
		logrus.Debugln(cdproto.CommandFetchTakeResponseBodyAsStream)
	case *fetch.EventRequestPaused:
		logrus.Debugln(cdproto.EventFetchRequestPaused)
	case *fetch.EventAuthRequired:
		logrus.Debugln(cdproto.EventFetchAuthRequired)
	case *headlessexperimental.BeginFrameReturns:
		logrus.Debugln(cdproto.CommandHeadlessExperimentalBeginFrame)
	case *heapprofiler.GetHeapObjectIDReturns:
		logrus.Debugln(cdproto.CommandHeapProfilerGetHeapObjectID)
	case *heapprofiler.GetObjectByHeapObjectIDReturns:
		logrus.Debugln(cdproto.CommandHeapProfilerGetObjectByHeapObjectID)
	case *heapprofiler.GetSamplingProfileReturns:
		logrus.Debugln(cdproto.CommandHeapProfilerGetSamplingProfile)
	case *heapprofiler.StopSamplingReturns:
		logrus.Debugln(cdproto.CommandHeapProfilerStopSampling)
	case *heapprofiler.EventAddHeapSnapshotChunk:
		logrus.Debugln(cdproto.EventHeapProfilerAddHeapSnapshotChunk)
	case *heapprofiler.EventHeapStatsUpdate:
		logrus.Debugln(cdproto.EventHeapProfilerHeapStatsUpdate)
	case *heapprofiler.EventLastSeenObjectID:
		logrus.Debugln(cdproto.EventHeapProfilerLastSeenObjectID)
	case *heapprofiler.EventReportHeapSnapshotProgress:
		logrus.Debugln(cdproto.EventHeapProfilerReportHeapSnapshotProgress)
	case *heapprofiler.EventResetProfiles:
		logrus.Debugln(cdproto.EventHeapProfilerResetProfiles)
	case *io.ReadReturns:
		logrus.Debugln(cdproto.CommandIORead)
	case *io.ResolveBlobReturns:
		logrus.Debugln(cdproto.CommandIOResolveBlob)
	case *indexeddb.RequestDataReturns:
		logrus.Debugln(cdproto.CommandIndexedDBRequestData)
	case *indexeddb.GetMetadataReturns:
		logrus.Debugln(cdproto.CommandIndexedDBGetMetadata)
	case *indexeddb.RequestDatabaseReturns:
		logrus.Debugln(cdproto.CommandIndexedDBRequestDatabase)
	case *indexeddb.RequestDatabaseNamesReturns:
		logrus.Debugln(cdproto.CommandIndexedDBRequestDatabaseNames)
	case *inspector.EventDetached:
		logrus.Debugln(cdproto.EventInspectorDetached)
	case *inspector.EventTargetCrashed:
		logrus.Debugln(cdproto.EventInspectorTargetCrashed)
	case *inspector.EventTargetReloadedAfterCrash:
		logrus.Debugln(cdproto.EventInspectorTargetReloadedAfterCrash)
	case *layertree.CompositingReasonsReturns:
		logrus.Debugln(cdproto.CommandLayerTreeCompositingReasons)
	case *layertree.LoadSnapshotReturns:
		logrus.Debugln(cdproto.CommandLayerTreeLoadSnapshot)
	case *layertree.MakeSnapshotReturns:
		logrus.Debugln(cdproto.CommandLayerTreeMakeSnapshot)
	case *layertree.ProfileSnapshotReturns:
		logrus.Debugln(cdproto.CommandLayerTreeProfileSnapshot)
	case *layertree.ReplaySnapshotReturns:
		logrus.Debugln(cdproto.CommandLayerTreeReplaySnapshot)
	case *layertree.SnapshotCommandLogReturns:
		logrus.Debugln(cdproto.CommandLayerTreeSnapshotCommandLog)
	case *layertree.EventLayerPainted:
		logrus.Debugln(cdproto.EventLayerTreeLayerPainted)
	case *layertree.EventLayerTreeDidChange:
		logrus.Debugln(cdproto.EventLayerTreeLayerTreeDidChange)
	case *log.EventEntryAdded:
		// logrus.Debugln(cdproto.EventLogEntryAdded)
	case *media.EventPlayerPropertiesChanged:
		logrus.Debugln(cdproto.EventMediaPlayerPropertiesChanged)
	case *media.EventPlayerEventsAdded:
		logrus.Debugln(cdproto.EventMediaPlayerEventsAdded)
	case *media.EventPlayersCreated:
		logrus.Debugln(cdproto.EventMediaPlayersCreated)
	case *memory.GetDOMCountersReturns:
		logrus.Debugln(cdproto.CommandMemoryGetDOMCounters)
	case *memory.GetAllTimeSamplingProfileReturns:
		logrus.Debugln(cdproto.CommandMemoryGetAllTimeSamplingProfile)
	case *memory.GetBrowserSamplingProfileReturns:
		logrus.Debugln(cdproto.CommandMemoryGetBrowserSamplingProfile)
	case *memory.GetSamplingProfileReturns:
		logrus.Debugln(cdproto.CommandMemoryGetSamplingProfile)
	case *network.GetAllCookiesReturns:
		logrus.Debugln(cdproto.CommandNetworkGetAllCookies)
	case *network.GetCertificateReturns:
		logrus.Debugln(cdproto.CommandNetworkGetCertificate)
	case *network.GetCookiesReturns:
		logrus.Debugln(cdproto.CommandNetworkGetCookies)
	case *network.GetResponseBodyReturns:
		logrus.Debugln(cdproto.CommandNetworkGetResponseBody)
	case *network.GetRequestPostDataReturns:
		logrus.Debugln(cdproto.CommandNetworkGetRequestPostData)
	case *network.GetResponseBodyForInterceptionReturns:
		logrus.Debugln(cdproto.CommandNetworkGetResponseBodyForInterception)
	case *network.TakeResponseBodyForInterceptionAsStreamReturns:
		logrus.Debugln(cdproto.CommandNetworkTakeResponseBodyForInterceptionAsStream)
	case *network.SearchInResponseBodyReturns:
		logrus.Debugln(cdproto.CommandNetworkSearchInResponseBody)
	case *network.SetCookieReturns:
		logrus.Debugln(cdproto.CommandNetworkSetCookie)
	case *network.EventDataReceived:
		logrus.Debugln(cdproto.EventNetworkDataReceived)
	case *network.EventEventSourceMessageReceived:
		logrus.Debugln(cdproto.EventNetworkEventSourceMessageReceived)
	case *network.EventLoadingFailed:
		logrus.Debugln(cdproto.EventNetworkLoadingFailed)
	case *network.EventLoadingFinished:
		logrus.Debugln(cdproto.EventNetworkLoadingFinished)
	case *network.EventRequestServedFromCache:
		logrus.Debugln(cdproto.EventNetworkRequestServedFromCache)
	case *network.EventRequestWillBeSent:
		logrus.Debugln(cdproto.EventNetworkRequestWillBeSent)
	case *network.EventResourceChangedPriority:
		logrus.Debugln(cdproto.EventNetworkResourceChangedPriority)
	case *network.EventSignedExchangeReceived:
		logrus.Debugln(cdproto.EventNetworkSignedExchangeReceived)
	case *network.EventResponseReceived:
		logrus.Debugln(cdproto.EventNetworkResponseReceived)
	case *network.EventWebSocketClosed:
		logrus.Debugln(cdproto.EventNetworkWebSocketClosed)
	case *network.EventWebSocketCreated:
		logrus.Debugln(cdproto.EventNetworkWebSocketCreated)
	case *network.EventWebSocketFrameError:
		logrus.Debugln(cdproto.EventNetworkWebSocketFrameError)
	case *network.EventWebSocketFrameReceived:
		logrus.Debugln(cdproto.EventNetworkWebSocketFrameReceived)
	case *network.EventWebSocketFrameSent:
		logrus.Debugln(cdproto.EventNetworkWebSocketFrameSent)
	case *network.EventWebSocketHandshakeResponseReceived:
		logrus.Debugln(cdproto.EventNetworkWebSocketHandshakeResponseReceived)
	case *network.EventWebSocketWillSendHandshakeRequest:
		logrus.Debugln(cdproto.EventNetworkWebSocketWillSendHandshakeRequest)
	case *network.EventRequestWillBeSentExtraInfo:
		logrus.Debugln(cdproto.EventNetworkRequestWillBeSentExtraInfo)
	case *network.EventResponseReceivedExtraInfo:
		logrus.Debugln(cdproto.EventNetworkResponseReceivedExtraInfo)
	case *overlay.GetHighlightObjectForTestReturns:
		logrus.Debugln(cdproto.CommandOverlayGetHighlightObjectForTest)
	case *overlay.EventInspectNodeRequested:
		logrus.Debugln(cdproto.EventOverlayInspectNodeRequested)
	case *overlay.EventNodeHighlightRequested:
		logrus.Debugln(cdproto.EventOverlayNodeHighlightRequested)
	case *overlay.EventScreenshotRequested:
		logrus.Debugln(cdproto.EventOverlayScreenshotRequested)
	case *overlay.EventInspectModeCanceled:
		logrus.Debugln(cdproto.EventOverlayInspectModeCanceled)
	case *page.AddScriptToEvaluateOnNewDocumentReturns:
		logrus.Debugln(cdproto.CommandPageAddScriptToEvaluateOnNewDocument)
	case *page.CaptureScreenshotReturns:
		logrus.Debugln(cdproto.CommandPageCaptureScreenshot)
	case *page.CaptureSnapshotReturns:
		logrus.Debugln(cdproto.CommandPageCaptureSnapshot)
	case *page.CreateIsolatedWorldReturns:
		logrus.Debugln(cdproto.CommandPageCreateIsolatedWorld)
	case *page.GetAppManifestReturns:
		logrus.Debugln(cdproto.CommandPageGetAppManifest)
	case *page.GetInstallabilityErrorsReturns:
		logrus.Debugln(cdproto.CommandPageGetInstallabilityErrors)
	case *page.GetFrameTreeReturns:
		logrus.Debugln(cdproto.CommandPageGetFrameTree)
	case *page.GetLayoutMetricsReturns:
		logrus.Debugln(cdproto.CommandPageGetLayoutMetrics)
	case *page.GetNavigationHistoryReturns:
		logrus.Debugln(cdproto.CommandPageGetNavigationHistory)
	case *page.GetResourceContentReturns:
		logrus.Debugln(cdproto.CommandPageGetResourceContent)
	case *page.GetResourceTreeReturns:
		logrus.Debugln(cdproto.CommandPageGetResourceTree)
	case *page.NavigateReturns:
		logrus.Debugln(cdproto.CommandPageNavigate)
	case *page.PrintToPDFReturns:
		logrus.Debugln(cdproto.CommandPagePrintToPDF)
	case *page.SearchInResourceReturns:
		logrus.Debugln(cdproto.CommandPageSearchInResource)
	case *page.EventDomContentEventFired:
		logrus.Debugln(cdproto.EventPageDomContentEventFired)
	case *page.EventFileChooserOpened:
		logrus.Debugln(cdproto.EventPageFileChooserOpened)
	case *page.EventFrameAttached:
		logrus.Debugln(cdproto.EventPageFrameAttached)
	case *page.EventFrameDetached:
		logrus.Debugln(cdproto.EventPageFrameDetached)
	case *page.EventFrameNavigated:
		logrus.Debugln(cdproto.EventPageFrameNavigated)
	case *page.EventFrameResized:
		logrus.Debugln(cdproto.EventPageFrameResized)
	case *page.EventFrameRequestedNavigation:
		logrus.Debugln(cdproto.EventPageFrameRequestedNavigation)
	case *page.EventFrameStartedLoading:
		logrus.Debugln(cdproto.EventPageFrameStartedLoading)
	case *page.EventFrameStoppedLoading:
		logrus.Debugln(cdproto.EventPageFrameStoppedLoading)
	case *page.EventDownloadWillBegin:
		logrus.Debugln(cdproto.EventPageDownloadWillBegin)
	case *page.EventInterstitialHidden:
		logrus.Debugln(cdproto.EventPageInterstitialHidden)
	case *page.EventInterstitialShown:
		logrus.Debugln(cdproto.EventPageInterstitialShown)
	case *page.EventJavascriptDialogClosed:
		logrus.Debugln(cdproto.EventPageJavascriptDialogClosed)
	case *page.EventJavascriptDialogOpening:
		logrus.Debugln(cdproto.EventPageJavascriptDialogOpening)
	case *page.EventLifecycleEvent:
		logrus.Debugln(cdproto.EventPageLifecycleEvent)
	case *page.EventLoadEventFired:
		logrus.Debugln(cdproto.EventPageLoadEventFired)
	case *page.EventNavigatedWithinDocument:
		logrus.Debugln(cdproto.EventPageNavigatedWithinDocument)
	case *page.EventScreencastFrame:
		logrus.Debugln(cdproto.EventPageScreencastFrame)
	case *page.EventScreencastVisibilityChanged:
		logrus.Debugln(cdproto.EventPageScreencastVisibilityChanged)
	case *page.EventWindowOpen:
		logrus.Debugln(cdproto.EventPageWindowOpen)
	case *page.EventCompilationCacheProduced:
		logrus.Debugln(cdproto.EventPageCompilationCacheProduced)
	case *performance.GetMetricsReturns:
		logrus.Debugln(cdproto.CommandPerformanceGetMetrics)
	case *performance.EventMetrics:
		logrus.Debugln(cdproto.EventPerformanceMetrics)
	case *profiler.GetBestEffortCoverageReturns:
		logrus.Debugln(cdproto.CommandProfilerGetBestEffortCoverage)
	case *profiler.StopReturns:
		logrus.Debugln(cdproto.CommandProfilerStop)
	case *profiler.TakePreciseCoverageReturns:
		logrus.Debugln(cdproto.CommandProfilerTakePreciseCoverage)
	case *profiler.TakeTypeProfileReturns:
		logrus.Debugln(cdproto.CommandProfilerTakeTypeProfile)
	case *profiler.EventConsoleProfileFinished:
		logrus.Debugln(cdproto.EventProfilerConsoleProfileFinished)
	case *profiler.EventConsoleProfileStarted:
		logrus.Debugln(cdproto.EventProfilerConsoleProfileStarted)
	case *runtime.AwaitPromiseReturns:
		logrus.Debugln(cdproto.CommandRuntimeAwaitPromise)
	case *runtime.CallFunctionOnReturns:
		logrus.Debugln(cdproto.CommandRuntimeCallFunctionOn)
	case *runtime.CompileScriptReturns:
		logrus.Debugln(cdproto.CommandRuntimeCompileScript)
	case *runtime.EvaluateReturns:
		logrus.Debugln(cdproto.CommandRuntimeEvaluate)
	case *runtime.GetIsolateIDReturns:
		logrus.Debugln(cdproto.CommandRuntimeGetIsolateID)
	case *runtime.GetHeapUsageReturns:
		logrus.Debugln(cdproto.CommandRuntimeGetHeapUsage)
	case *runtime.GetPropertiesReturns:
		logrus.Debugln(cdproto.CommandRuntimeGetProperties)
	case *runtime.GlobalLexicalScopeNamesReturns:
		logrus.Debugln(cdproto.CommandRuntimeGlobalLexicalScopeNames)
	case *runtime.QueryObjectsReturns:
		logrus.Debugln(cdproto.CommandRuntimeQueryObjects)
	case *runtime.RunScriptReturns:
		logrus.Debugln(cdproto.CommandRuntimeRunScript)
	case *runtime.EventBindingCalled:
		logrus.Debugln(cdproto.EventRuntimeBindingCalled)
	case *runtime.EventConsoleAPICalled:
		logrus.Debugln(cdproto.EventRuntimeConsoleAPICalled)
	case *runtime.EventExceptionRevoked:
		logrus.Debugln(cdproto.EventRuntimeExceptionRevoked)
	case *runtime.EventExceptionThrown:
		logrus.Debugln(cdproto.EventRuntimeExceptionThrown)
	case *runtime.EventExecutionContextCreated:
		//	logrus.Debugln(cdproto.EventRuntimeExecutionContextCreated)
	case *runtime.EventExecutionContextDestroyed:
		logrus.Debugln(cdproto.EventRuntimeExecutionContextDestroyed)
	case *runtime.EventExecutionContextsCleared:
		logrus.Debugln(cdproto.EventRuntimeExecutionContextsCleared)
	case *runtime.EventInspectRequested:
		logrus.Debugln(cdproto.EventRuntimeInspectRequested)
	case *security.EventVisibleSecurityStateChanged:
		logrus.Debugln(cdproto.EventSecurityVisibleSecurityStateChanged)
	case *security.EventSecurityStateChanged:
		logrus.Debugln(cdproto.EventSecuritySecurityStateChanged)
	case *serviceworker.EventWorkerErrorReported:
		logrus.Debugln(cdproto.EventServiceWorkerWorkerErrorReported)
	case *serviceworker.EventWorkerRegistrationUpdated:
		logrus.Debugln(cdproto.EventServiceWorkerWorkerRegistrationUpdated)
	case *serviceworker.EventWorkerVersionUpdated:
		logrus.Debugln(cdproto.EventServiceWorkerWorkerVersionUpdated)
	case *storage.GetUsageAndQuotaReturns:
		logrus.Debugln(cdproto.CommandStorageGetUsageAndQuota)
	case *storage.EventCacheStorageContentUpdated:
		logrus.Debugln(cdproto.EventStorageCacheStorageContentUpdated)
	case *storage.EventCacheStorageListUpdated:
		logrus.Debugln(cdproto.EventStorageCacheStorageListUpdated)
	case *storage.EventIndexedDBContentUpdated:
		logrus.Debugln(cdproto.EventStorageIndexedDBContentUpdated)
	case *storage.EventIndexedDBListUpdated:
		logrus.Debugln(cdproto.EventStorageIndexedDBListUpdated)
	case *systeminfo.GetInfoReturns:
		logrus.Debugln(cdproto.CommandSystemInfoGetInfo)
	case *systeminfo.GetProcessInfoReturns:
		logrus.Debugln(cdproto.CommandSystemInfoGetProcessInfo)
	case *target.AttachToTargetReturns:
		logrus.Debugln(cdproto.CommandTargetAttachToTarget)
	case *target.AttachToBrowserTargetReturns:
		logrus.Debugln(cdproto.CommandTargetAttachToBrowserTarget)
	case *target.CloseTargetReturns:
		logrus.Debugln(cdproto.CommandTargetCloseTarget)
	case *target.CreateBrowserContextReturns:
		logrus.Debugln(cdproto.CommandTargetCreateBrowserContext)
	case *target.GetBrowserContextsReturns:
		logrus.Debugln(cdproto.CommandTargetGetBrowserContexts)
	case *target.CreateTargetReturns:
		logrus.Debugln(cdproto.CommandTargetCreateTarget)
	case *target.GetTargetInfoReturns:
		logrus.Debugln(cdproto.CommandTargetGetTargetInfo)
	case *target.GetTargetsReturns:
		logrus.Debugln(cdproto.CommandTargetGetTargets)
	case *target.EventAttachedToTarget:
		logrus.Debugln(cdproto.EventTargetAttachedToTarget)
	case *target.EventDetachedFromTarget:
		logrus.Debugln(cdproto.EventTargetDetachedFromTarget)
	case *target.EventReceivedMessageFromTarget:
		logrus.Debugln(cdproto.EventTargetReceivedMessageFromTarget)
	case *target.EventTargetCreated:
		// 	logrus.Debugln(cdproto.EventTargetTargetCreated)
	case *target.EventTargetDestroyed:
		logrus.Debugln(cdproto.EventTargetTargetDestroyed)
	case *target.EventTargetCrashed:
		logrus.Debugln(cdproto.EventTargetTargetCrashed)
	case *target.EventTargetInfoChanged:
		// 	logrus.Debugln(cdproto.EventTargetTargetInfoChanged)
	case *tethering.EventAccepted:
		logrus.Debugln(cdproto.EventTetheringAccepted)
	case *tracing.GetCategoriesReturns:
		logrus.Debugln(cdproto.CommandTracingGetCategories)
	case *tracing.RequestMemoryDumpReturns:
		logrus.Debugln(cdproto.CommandTracingRequestMemoryDump)
	case *tracing.EventBufferUsage:
		logrus.Debugln(cdproto.EventTracingBufferUsage)
	case *tracing.EventDataCollected:
		logrus.Debugln(cdproto.EventTracingDataCollected)
	case *tracing.EventTracingComplete:
		logrus.Debugln(cdproto.EventTracingTracingComplete)
	case *webaudio.GetRealtimeDataReturns:
		logrus.Debugln(cdproto.CommandWebAudioGetRealtimeData)
	case *webaudio.EventContextCreated:
		logrus.Debugln(cdproto.EventWebAudioContextCreated)
	case *webaudio.EventContextWillBeDestroyed:
		logrus.Debugln(cdproto.EventWebAudioContextWillBeDestroyed)
	case *webaudio.EventContextChanged:
		logrus.Debugln(cdproto.EventWebAudioContextChanged)
	case *webaudio.EventAudioListenerCreated:
		logrus.Debugln(cdproto.EventWebAudioAudioListenerCreated)
	case *webaudio.EventAudioListenerWillBeDestroyed:
		logrus.Debugln(cdproto.EventWebAudioAudioListenerWillBeDestroyed)
	case *webaudio.EventAudioNodeCreated:
		logrus.Debugln(cdproto.EventWebAudioAudioNodeCreated)
	case *webaudio.EventAudioNodeWillBeDestroyed:
		logrus.Debugln(cdproto.EventWebAudioAudioNodeWillBeDestroyed)
	case *webaudio.EventAudioParamCreated:
		logrus.Debugln(cdproto.EventWebAudioAudioParamCreated)
	case *webaudio.EventAudioParamWillBeDestroyed:
		logrus.Debugln(cdproto.EventWebAudioAudioParamWillBeDestroyed)
	case *webaudio.EventNodesConnected:
		logrus.Debugln(cdproto.EventWebAudioNodesConnected)
	case *webaudio.EventNodesDisconnected:
		logrus.Debugln(cdproto.EventWebAudioNodesDisconnected)
	case *webaudio.EventNodeParamConnected:
		logrus.Debugln(cdproto.EventWebAudioNodeParamConnected)
	case *webaudio.EventNodeParamDisconnected:
		logrus.Debugln(cdproto.EventWebAudioNodeParamDisconnected)
	case *webauthn.AddVirtualAuthenticatorReturns:
		logrus.Debugln(cdproto.CommandWebAuthnAddVirtualAuthenticator)
	case *webauthn.GetCredentialReturns:
		logrus.Debugln(cdproto.CommandWebAuthnGetCredential)
	case *webauthn.GetCredentialsReturns:
		logrus.Debugln(cdproto.CommandWebAuthnGetCredentials)
	case *css.AddRuleReturns:
		if cssOn {
			logrus.Debugln(cdproto.CommandCSSAddRule)
		}
		cssOn = true
	case *css.CollectClassNamesReturns:
		if cssOn {
			logrus.Debugln(cdproto.CommandCSSCollectClassNames)
		}
		cssOn = true
	case *css.CreateStyleSheetReturns:
		if cssOn {
			logrus.Debugln(cdproto.CommandCSSCreateStyleSheet)
		}
		cssOn = true
	case *css.GetBackgroundColorsReturns:
		if cssOn {
			logrus.Debugln(cdproto.CommandCSSGetBackgroundColors)
		}
		cssOn = true
	case *css.GetComputedStyleForNodeReturns:
		if cssOn {
			logrus.Debugln(cdproto.CommandCSSGetComputedStyleForNode)
		}
		cssOn = true
	case *css.GetInlineStylesForNodeReturns:
		if cssOn {
			logrus.Debugln(cdproto.CommandCSSGetInlineStylesForNode)
		}
		cssOn = true
	case *css.GetMatchedStylesForNodeReturns:
		if cssOn {
			logrus.Debugln(cdproto.CommandCSSGetMatchedStylesForNode)
		}
		cssOn = true
	case *css.GetMediaQueriesReturns:
		if cssOn {
			logrus.Debugln(cdproto.CommandCSSGetMediaQueries)
		}
		cssOn = true
	case *css.GetPlatformFontsForNodeReturns:
		if cssOn {
			logrus.Debugln(cdproto.CommandCSSGetPlatformFontsForNode)
		}
		cssOn = true
	case *css.GetStyleSheetTextReturns:
		if cssOn {
			logrus.Debugln(cdproto.CommandCSSGetStyleSheetText)
		}
		cssOn = true
	case *css.SetKeyframeKeyReturns:
		if cssOn {
			logrus.Debugln(cdproto.CommandCSSSetKeyframeKey)
		}
		cssOn = true
	case *css.SetMediaTextReturns:
		if cssOn {
			logrus.Debugln(cdproto.CommandCSSSetMediaText)
		}
		cssOn = true
	case *css.SetRuleSelectorReturns:
		if cssOn {
			logrus.Debugln(cdproto.CommandCSSSetRuleSelector)
		}
		cssOn = true
	case *css.SetStyleSheetTextReturns:
		if cssOn {
			logrus.Debugln(cdproto.CommandCSSSetStyleSheetText)
		}
		cssOn = true
	case *css.SetStyleTextsReturns:
		if cssOn {
			logrus.Debugln(cdproto.CommandCSSSetStyleTexts)
		}
		cssOn = true
	case *css.StopRuleUsageTrackingReturns:
		if cssOn {
			logrus.Debugln(cdproto.CommandCSSStopRuleUsageTracking)
		}
		cssOn = true
	case *css.TakeCoverageDeltaReturns:
		if cssOn {
			logrus.Debugln(cdproto.CommandCSSTakeCoverageDelta)
		}
		cssOn = true
	case *css.EventFontsUpdated:
		if cssOn {
			logrus.Debugln(cdproto.EventCSSFontsUpdated)
		}
		cssOn = true
	case *css.EventMediaQueryResultChanged:
		if cssOn {
			logrus.Debugln(cdproto.EventCSSMediaQueryResultChanged)
		}
		cssOn = true
	case *css.EventStyleSheetAdded:
		if cssOn {
			logrus.Debugln(cdproto.EventCSSStyleSheetAdded)
		}
		cssOn = true
	case *css.EventStyleSheetChanged:
		if cssOn {
			logrus.Debugln(cdproto.EventCSSStyleSheetChanged)
		}
		cssOn = true
	case *css.EventStyleSheetRemoved:
		if cssOn {
			logrus.Debugln(cdproto.EventCSSStyleSheetRemoved)
		}
		cssOn = true
	default:
		defaultVale = true
		logrus.Errorfp("unknow event.", ev)
	}

	if !defaultVale {
		if cssOn {
			return
		}
		// logrus.Warnfp("", ev)
	}
}
