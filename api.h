#ifndef RAWRTC_API_H_
#define RAWRTC_API_H_

#include <string.h>

#include "include/api.h"
#include "include/ice_candidate.h"
#include "include/rtc_error.h"
#include "include/rtp_parameters.h"
#include "include/rtp_transceiver.h"
#include "include/session_description.h"

int InitializeLibrary(const char *file, raw_rtc_constraints_t constraints);
void *Calloc(size_t size);
void Free(void *p);

void LogInfo(const char *message);
void LogWarn(const char *message);
void LogError(const char *message);

void *CreatePeerConnectionFactory(void *fd);
raw_rtp_capabilities_t GetRtpSenderCapabilities(void *factory, const char *kind);
raw_rtp_capabilities_t GetRtpReceiverCapabilities(void *factory, const char *kind);
void *CreatePeerConnection(void *factory, void *fd, raw_rtc_configuration_t *configuration);
void *CreateAudioTrack(void *factory, void *fd, const char *id, void *source);
void *CreateVideoTrack(void *factory, void *fd, const char *id, void *source);

void *PeerConnectionAddTrack(void *pc, void *track, size_t size, void **streams, raw_rtc_error_t *err);
int PeerConnectionRemoveTrack(void *pc, void *sender, raw_rtc_error_t *err);
void *PeerConnectionAddTransceiver(void *pc, const char *media_type, raw_rtp_transceiver_init_t *init, raw_rtc_error_t *err);
void PeerConnectionCreateOffer(void *pc, void *observer);
void PeerConnectionCreateAnswer(void *pc, void *observer);
void PeerConnectionSetLocalDescription(void *pc, void *observer, raw_session_description_t *desc);
void PeerConnectionSetRemoteDescription(void *pc, void *observer, raw_session_description_t *desc);
int PeerConnectionAddIceCandidate(void *pc, raw_ice_candidate_t *candidate, raw_rtc_error_t *err);
void PeerConnectionGetReceivers(void *pc, size_t *size, void **array);
void PeerConnectionGetSenders(void *pc, size_t *size, void **array);
void PeerConnectionGetTransceivers(void *pc, size_t *size, void **array);
void PeerConnectionClose(void *pc);
void PeerConnectionRelease(void *pc);

void *CreateCreateSessionDescriptionObserver(void *fd);
void CreateSessionDescriptionObserverRelease(void *observer);

void *CreateSetSessionDescriptionObserver(void *fd);
void SetSessionDescriptionObserverRelease(void *observer);

const char *MediaStreamGetID(void *stream);
int MediaStreamAddTrack(void *stream, void *track);
int MediaStreamRemoveTrack(void *stream, void *track);
void MediaStreamGetAudioTracks(void *stream, size_t *size, void **array);
void MediaStreamGetVideoTracks(void *stream, size_t *size, void **array);
void *MediaStreamFindAudioTrack(void *stream, const char *id);
void *MediaStreamFindVideoTrack(void *stream, const char *id);
void MediaStreamRelease(void *stream);

const char *MediaStreamTrackGetID(void *track);
const char *MediaStreamTrackGetKind(void *track);
int MediaStreamTrackGetMuted(void *track);
const char *MediaStreamTrackGetState(void *track);
void *MediaStreamTrackGetSource(void *track);
void MediaStreamTrackStop(void *track);
void MediaStreamTrackRelease(void *track);

void MediaSourceRelease(void *source);

const char *RtpTransceiverGetDirection(void *transceiver);
const char *RtpTransceiverGetMid(void *transceiver);
void *RtpTransceiverGetReceiver(void *transceiver);
void *RtpTransceiverGetSender(void *transceiver);
void RtpTransceiverSetCodecPreferences(void *transceiver, void **codecs, size_t size);
int RtpTransceiverSetDirection(void *transceiver, const char *new_direction, raw_rtc_error_t *err);
void RtpTransceiverStop(void *transceiver);
void RtpTransceiverRelease(void *transceiver);

void *RtpReceiverGetTrack(void *receiver);
void RtpReceiverGetStreams(void *receiver, size_t *size, void **array);
void RtpReceiverGetParameters(void *receiver, void *parameters);
void RtpReceiverGetStats(void *receiver, void *stats);
void RtpReceiverRelease(void *receiver);

int RtpSenderSetTrack(void *sender, void *track);
void *RtpSenderGetTrack(void *sender);
void RtpSenderSetStreams(void *sender, size_t size, const char **stream_ids);
void RtpSenderGetStreams(void *sender, size_t *size, void **array);
void RtpSenderSetParameters(void *sender, void *parameters);
void RtpSenderGetParameters(void *sender, void *parameters);
void RtpSenderGetStats(void *sender, void *stats);
void RtpSenderRelease(void *sender);

#endif // RAWRTC_API_H_
