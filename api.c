#include "api.h"
#include <dlfcn.h>
#include <stdio.h>
#include <stdlib.h>

typedef void *(*__create_peer_connection_factory_fptr__)(void *fd);
typedef void *(*__create_peer_connection_fptr__)(void *factory, void *pc, raw_peer_connection_observer_t *cb);
typedef void *(*__create_audio_track_fptr__)(void *factory, void *track, const char *id, void *source);
typedef void *(*__create_video_track_fptr__)(void *factory, void *track, const char *id, void *source);

typedef void *(*__peer_connection_add_track_fptr__)(void *pc, void *track, size_t size, void *streams, raw_rtc_error_t *err);
typedef int (*__peer_connection_remove_track_fptr__)(void *pc, void *sender, raw_rtc_error_t *err);
typedef void *(*__peer_connection_add_transceiver_fptr__)(void *pc, const char *media_type, raw_rtp_transceiver_init_t *init, raw_rtc_error_t *err);
typedef void (*__peer_connection_create_offer_fptr__)(void *pc, void *observer, raw_create_session_description_observer_t *cb);
typedef void (*__peer_connection_create_answer_fptr__)(void *pc, void *observer, raw_create_session_description_observer_t *cb);
typedef void (*__peer_connection_set_local_description_fptr__)(void *pc, void *observer, raw_set_session_description_observer_t *cb, raw_session_description_t *desc);
typedef void (*__peer_connection_set_remote_description_fptr__)(void *pc, void *observer, raw_set_session_description_observer_t *cb, raw_session_description_t *desc);
typedef int (*__peer_connection_add_ice_candidate_fptr__)(void *pc, raw_ice_candidate_t *candidate, raw_rtc_error_t *err);
typedef void (*__peer_connection_get_receivers_fptr__)(void *pc, size_t *size, void **array);
typedef void (*__peer_connection_get_senders_fptr__)(void *pc, size_t *size, void **array);
typedef void (*__peer_connection_get_transceivers_fptr__)(void *pc, size_t *size, void **array);
typedef void (*__peer_connection_close_fptr__)(void *pc);

typedef const char *(*__media_stream_get_id_fptr__)(void *stream);
typedef int (*__media_stream_add_track_fptr__)(void *stream, void *track);
typedef int (*__media_stream_remove_track_fptr__)(void *stream, void *track);
typedef void (*__media_stream_get_audio_tracks_fptr__)(void *stream, size_t *size, void **array);
typedef void (*__media_stream_get_video_tracks_fptr__)(void *stream, size_t *size, void **array);
typedef void *(*__media_stream_find_audio_track_fptr__)(void *stream, const char *id);
typedef void *(*__media_stream_find_video_track_fptr__)(void *stream, const char *id);

typedef const char *(*__media_stream_track_get_id_fptr__)(void *track);
typedef const char *(*__media_stream_track_get_kind_fptr__)(void *track);
typedef int (*__media_stream_track_get_muted_fptr__)(void *track);
typedef const char *(*__media_stream_track_get_state_fptr__)(void *track);
typedef void *(*__media_stream_track_get_source_fptr__)(void *track);
typedef void (*__media_stream_track_stop_fptr__)(void *track);

typedef const char *(*__rtp_transceiver_get_direction_fptr__)(void *transceiver);
typedef const char *(*__rtp_transceiver_get_mid_fptr__)(void *transceiver);
typedef void *(*__rtp_transceiver_get_receiver_fptr__)(void *transceiver);
typedef void *(*__rtp_transceiver_get_sender_fptr__)(void *transceiver);
typedef int (*__rtp_transceiver_set_direction_fptr__)(void *transceiver, const char *new_direction, raw_rtc_error_t *err);
typedef void (*__rtp_transceiver_stop_fptr__)(void *transceiver);

typedef void *(*__rtp_receiver_get_track_fptr__)(void *receiver);
typedef void (*__rtp_receiver_get_streams_fptr__)(void *receiver, size_t *size, void **array);
typedef void (*__rtp_receiver_get_parameters_fptr__)(void *receiver, void *parameters);
typedef void (*__rtp_receiver_get_stats_fptr__)(void *receiver, void *stats);

typedef int (*__rtp_sender_set_track_fptr__)(void *sender, void *track);
typedef void *(*__rtp_sender_get_track_fptr__)(void *sender);
typedef void (*__rtp_sender_set_streams_fptr__)(void *sender, size_t size, const char **stream_ids);
typedef void (*__rtp_sender_get_streams_fptr__)(void *sender, raw_array_t *dst);
typedef void (*__rtp_sender_set_parameters_fptr__)(void *sender, void *parameters);
typedef void (*__rtp_sender_get_parameters_fptr__)(void *sender, void *parameters);
typedef void (*__rtp_sender_get_stats_fptr__)(void *sender, void *stats);

void *handle;
__create_peer_connection_factory_fptr__ __create_peer_connection_factory__;
__create_peer_connection_fptr__ __create_peer_connection__;
__create_audio_track_fptr__ __create_audio_track__;
__create_video_track_fptr__ __create_video_track__;

__peer_connection_add_track_fptr__ __peer_connection_add_track__;
__peer_connection_remove_track_fptr__ __peer_connection_remove_track__;
__peer_connection_add_transceiver_fptr__ __peer_connection_add_transceiver__;
__peer_connection_create_offer_fptr__ __peer_connection_create_offer__;
__peer_connection_create_answer_fptr__ __peer_connection_create_answer__;
__peer_connection_set_local_description_fptr__ __peer_connection_set_local_description__;
__peer_connection_set_remote_description_fptr__ __peer_connection_set_remote_description__;
__peer_connection_add_ice_candidate_fptr__ __peer_connection_add_ice_candidate__;
__peer_connection_get_receivers_fptr__ __peer_connection_get_receivers__;
__peer_connection_get_senders_fptr__ __peer_connection_get_senders__;
__peer_connection_get_transceivers_fptr__ __peer_connection_get_transceivers__;
__peer_connection_close_fptr__ __peer_connection_close__;

__media_stream_get_id_fptr__ __media_stream_get_id__;
__media_stream_add_track_fptr__ __media_stream_add_track__;
__media_stream_remove_track_fptr__ __media_stream_remove_track__;
__media_stream_get_audio_tracks_fptr__ __media_stream_get_audio_tracks__;
__media_stream_get_video_tracks_fptr__ __media_stream_get_video_tracks__;
__media_stream_find_audio_track_fptr__ __media_stream_find_audio_track__;
__media_stream_find_video_track_fptr__ __media_stream_find_video_track__;

__media_stream_track_get_id_fptr__ __media_stream_track_get_id__;
__media_stream_track_get_kind_fptr__ __media_stream_track_get_kind__;
__media_stream_track_get_muted_fptr__ __media_stream_track_get_muted__;
__media_stream_track_get_state_fptr__ __media_stream_track_get_state__;
__media_stream_track_get_source_fptr__ __media_stream_track_get_source__;
__media_stream_track_stop_fptr__ __media_stream_track_stop__;

__rtp_transceiver_get_direction_fptr__ __rtp_transceiver_get_direction__;
__rtp_transceiver_get_mid_fptr__ __rtp_transceiver_get_mid__;
__rtp_transceiver_get_receiver_fptr__ __rtp_transceiver_get_receiver__;
__rtp_transceiver_get_sender_fptr__ __rtp_transceiver_get_sender__;
__rtp_transceiver_set_direction_fptr__ __rtp_transceiver_set_direction__;
__rtp_transceiver_stop_fptr__ __rtp_transceiver_stop__;

__rtp_receiver_get_track_fptr__ __rtp_receiver_get_track__;
__rtp_receiver_get_streams_fptr__ __rtp_receiver_get_streams__;
__rtp_receiver_get_parameters_fptr__ __rtp_receiver_get_parameters__;
__rtp_receiver_get_stats_fptr__ __rtp_receiver_get_stats__;

__rtp_sender_set_track_fptr__ __rtp_sender_set_track__;
__rtp_sender_get_track_fptr__ __rtp_sender_get_track__;
__rtp_sender_set_streams_fptr__ __rtp_sender_set_streams__;
__rtp_sender_get_streams_fptr__ __rtp_sender_get_streams__;
__rtp_sender_set_parameters_fptr__ __rtp_sender_set_parameters__;
__rtp_sender_get_parameters_fptr__ __rtp_sender_get_parameters__;
__rtp_sender_get_stats_fptr__ __rtp_sender_get_stats__;

raw_peer_connection_observer_t *__peer_connection_observer__;
extern void __onsignalingchange__(void *observer, const char *new_state);
extern void __ondatachannel__(void *observer, void *data_channel);
extern void __onrenegotiationneeded__(void *observer);
extern void __onconnectionchange__(void *observer, const char *new_state);
extern void __oniceconnectionchange__(void *observer, const char *new_state);
extern void __onicegatheringchange__(void *observer, const char *new_state);
extern void __onicecandidate__(void *observer, const char *candidate, const char *sdp_mid, int sdp_mline_index);
extern void __onicecandidateerror__(void *observer, const char *address, int port, const char *url, int error_code, const char *error_text);
extern void __ontrack__(void *observer, void *transceiver);

raw_create_session_description_observer_t *__create_session_description_observer__;
extern void __oncreatesessiondescriptionsuccess__(void *observer, const char *type, const char *sdp);
extern void __oncreatesessiondescriptionfailure__(void *observer, const char *name, const char *message);

raw_set_session_description_observer_t *__set_session_description_observer__;
extern void __onsetsessiondescriptionsuccess__(void *observer);
extern void __onsetsessiondescriptionfailure__(void *observer, const char *name, const char *message);

int LoadLibrary(const char *file)
{
    handle = dlopen(file, 1);
    if (handle == NULL)
    {
        printf("Failed to open library: file=%s, err=%s\n", file, dlerror());
        return -1;
    }
    __create_peer_connection_factory__ = (__create_peer_connection_factory_fptr__)dlsym(handle, "CreatePeerConnectionFactory");
    __create_peer_connection__ = (__create_peer_connection_fptr__)dlsym(handle, "CreatePeerConnection");
    __create_audio_track__ = (__create_audio_track_fptr__)dlsym(handle, "CreateAudioTrack");
    __create_video_track__ = (__create_video_track_fptr__)dlsym(handle, "CreateVideoTrack");

    __peer_connection_add_track__ = (__peer_connection_add_track_fptr__)dlsym(handle, "PeerConnectionAddTrack");
    __peer_connection_remove_track__ = (__peer_connection_remove_track_fptr__)dlsym(handle, "PeerConnectionRemoveTrack");
    __peer_connection_add_transceiver__ = (__peer_connection_add_transceiver_fptr__)dlsym(handle, "PeerConnectionAddTransceiver");
    __peer_connection_create_offer__ = (__peer_connection_create_offer_fptr__)dlsym(handle, "PeerConnectionCreateOffer");
    __peer_connection_create_answer__ = (__peer_connection_create_answer_fptr__)dlsym(handle, "PeerConnectionCreateAnswer");
    __peer_connection_set_local_description__ = (__peer_connection_set_local_description_fptr__)dlsym(handle, "PeerConnectionSetLocalDescription");
    __peer_connection_set_remote_description__ = (__peer_connection_set_remote_description_fptr__)dlsym(handle, "PeerConnectionSetRemoteDescription");
    __peer_connection_add_ice_candidate__ = (__peer_connection_add_ice_candidate_fptr__)dlsym(handle, "PeerConnectionAddIceCandidate");
    __peer_connection_get_receivers__ = (__peer_connection_get_receivers_fptr__)dlsym(handle, "PeerConnectionGetReceivers");
    __peer_connection_get_senders__ = (__peer_connection_get_senders_fptr__)dlsym(handle, "PeerConnectionGetSenders");
    __peer_connection_get_transceivers__ = (__peer_connection_get_transceivers_fptr__)dlsym(handle, "PeerConnectionGetTransceivers");
    __peer_connection_close__ = (__peer_connection_close_fptr__)dlsym(handle, "PeerConnectionClose");

    __media_stream_get_id__ = (__media_stream_get_id_fptr__)dlsym(handle, "MediaStreamGetID");
    __media_stream_add_track__ = (__media_stream_add_track_fptr__)dlsym(handle, "MediaStreamAddTrack");
    __media_stream_remove_track__ = (__media_stream_remove_track_fptr__)dlsym(handle, "MediaStreamRemoveTrack");
    __media_stream_get_audio_tracks__ = (__media_stream_get_audio_tracks_fptr__)dlsym(handle, "MediaStreamGetAudioTracks");
    __media_stream_get_video_tracks__ = (__media_stream_get_video_tracks_fptr__)dlsym(handle, "MediaStreamGetVideoTracks");
    __media_stream_find_audio_track__ = (__media_stream_find_audio_track_fptr__)dlsym(handle, "MediaStreamGetAudioTrackByID");
    __media_stream_find_video_track__ = (__media_stream_find_video_track_fptr__)dlsym(handle, "MediaStreamGetVideoTrackByID");

    __media_stream_track_get_id__ = (__media_stream_track_get_id_fptr__)dlsym(handle, "MediaStreamTrackGetID");
    __media_stream_track_get_kind__ = (__media_stream_track_get_kind_fptr__)dlsym(handle, "MediaStreamTrackGetKind");
    __media_stream_track_get_muted__ = (__media_stream_track_get_muted_fptr__)dlsym(handle, "MediaStreamTrackGetMuted");
    __media_stream_track_get_state__ = (__media_stream_track_get_state_fptr__)dlsym(handle, "MediaStreamTrackGetState");
    __media_stream_track_get_source__ = (__media_stream_track_get_source_fptr__)dlsym(handle, "MediaStreamTrackGetSource");
    __media_stream_track_stop__ = (__media_stream_track_stop_fptr__)dlsym(handle, "MediaStreamTrackStop");

    __rtp_transceiver_get_direction__ = (__rtp_transceiver_get_direction_fptr__)dlsym(handle, "RtpTransceiverGetDirection");
    __rtp_transceiver_get_mid__ = (__rtp_transceiver_get_mid_fptr__)dlsym(handle, "RtpTransceiverGetMid");
    __rtp_transceiver_get_receiver__ = (__rtp_transceiver_get_receiver_fptr__)dlsym(handle, "RtpTransceiverGetReceiver");
    __rtp_transceiver_get_sender__ = (__rtp_transceiver_get_sender_fptr__)dlsym(handle, "RtpTransceiverGetSender");
    __rtp_transceiver_set_direction__ = (__rtp_transceiver_set_direction_fptr__)dlsym(handle, "RtpTransceiverSetDirection");
    __rtp_transceiver_stop__ = (__rtp_transceiver_stop_fptr__)dlsym(handle, "RtpTransceiverStop");

    __rtp_receiver_get_track__ = (__rtp_receiver_get_track_fptr__)dlsym(handle, "RtpReceiverGetTrack");
    __rtp_receiver_get_streams__ = (__rtp_receiver_get_streams_fptr__)dlsym(handle, "RtpReceiverGetStreams");
    __rtp_receiver_get_parameters__ = (__rtp_receiver_get_parameters_fptr__)dlsym(handle, "RtpReceiverGetParameters");
    __rtp_receiver_get_stats__ = (__rtp_receiver_get_stats_fptr__)dlsym(handle, "RtpReceiverGetStats");

    __rtp_sender_set_track__ = (__rtp_sender_set_track_fptr__)dlsym(handle, "RtpSenderSetTrack");
    __rtp_sender_get_track__ = (__rtp_sender_get_track_fptr__)dlsym(handle, "RtpSenderGetTrack");
    __rtp_sender_set_streams__ = (__rtp_sender_set_streams_fptr__)dlsym(handle, "RtpSenderSetStreams");
    __rtp_sender_get_streams__ = (__rtp_sender_get_streams_fptr__)dlsym(handle, "RtpSenderGetStreams");
    __rtp_sender_set_parameters__ = (__rtp_sender_set_parameters_fptr__)dlsym(handle, "RtpSenderSetParameters");
    __rtp_sender_get_parameters__ = (__rtp_sender_get_parameters_fptr__)dlsym(handle, "RtpSenderGetParameters");
    __rtp_sender_get_stats__ = (__rtp_sender_get_stats_fptr__)dlsym(handle, "RtpSenderGetStats");

    __peer_connection_observer__ = malloc(sizeof(raw_peer_connection_observer_t));
    __peer_connection_observer__->onsignalingchange = __onsignalingchange__;
    __peer_connection_observer__->ondatachannel = __ondatachannel__;
    __peer_connection_observer__->onrenegotiationneeded = __onrenegotiationneeded__;
    __peer_connection_observer__->onconnectionchange = __onconnectionchange__;
    __peer_connection_observer__->oniceconnectionchange = __oniceconnectionchange__;
    __peer_connection_observer__->onicegatheringchange = __onicegatheringchange__;
    __peer_connection_observer__->onicecandidate = __onicecandidate__;
    __peer_connection_observer__->onicecandidateerror = __onicecandidateerror__;
    __peer_connection_observer__->ontrack = __ontrack__;

    __create_session_description_observer__ = malloc(sizeof(raw_create_session_description_observer_t));
    __create_session_description_observer__->onsuccess = __oncreatesessiondescriptionsuccess__;
    __create_session_description_observer__->onfailure = __oncreatesessiondescriptionfailure__;

    __set_session_description_observer__ = malloc(sizeof(raw_set_session_description_observer_t));
    __set_session_description_observer__->onsuccess = __onsetsessiondescriptionsuccess__;
    __set_session_description_observer__->onfailure = __onsetsessiondescriptionfailure__;
    return 0;
}

void *CreatePeerConnectionFactory(void *fd)
{
    printf("===> CreatePeerConnectionFactory()\n");
    return (*__create_peer_connection_factory__)(fd);
}

void *CreatePeerConnection(void *factory, void *fd)
{
    printf("===> CreatePeerConnection()\n");
    return (*__create_peer_connection__)(factory, fd, __peer_connection_observer__);
}

void *CreateAudioTrack(void *factory, void *fd, const char *id, void *source)
{
    printf("===> CreateAudioTrack(%s)\n", id);
    return (*__create_audio_track__)(factory, fd, id, source);
}

void *CreateVideoTrack(void *factory, void *fd, const char *id, void *source)
{
    printf("===> CreateVideoTrack(%s)\n", id);
    return (*__create_video_track__)(factory, fd, id, source);
}

void *PeerConnectionAddTrack(void *pc, void *track, size_t size, void **streams, raw_rtc_error_t *err)
{
    printf("===> PeerConnectionAddTrack()\n");
    return (*__peer_connection_add_track__)(pc, track, size, streams, err);
}

int PeerConnectionRemoveTrack(void *pc, void *sender, raw_rtc_error_t *err)
{
    printf("===> PeerConnectionRemoveTrack()\n");
    return (*__peer_connection_remove_track__)(pc, sender, err);
}

void *PeerConnectionAddTransceiver(void *pc, const char *media_type, raw_rtp_transceiver_init_t *init, raw_rtc_error_t *err)
{
    printf("===> PeerConnectionAddTransceiver(%s)\n", media_type);
    return (*__peer_connection_add_transceiver__)(pc, media_type, init, err);
}

void PeerConnectionCreateOffer(void *pc, void *observer)
{
    printf("===> PeerConnectionCreateOffer()\n");
    (*__peer_connection_create_offer__)(pc, observer, __create_session_description_observer__);
}

void PeerConnectionCreateAnswer(void *pc, void *observer)
{
    printf("===> PeerConnectionCreateAnswer()\n");
    (*__peer_connection_create_answer__)(pc, observer, __create_session_description_observer__);
}

void PeerConnectionSetLocalDescription(void *pc, void *observer, raw_session_description_t *desc)
{
    printf("===> PeerConnectionSetLocalDescription(type=%s, sdp=\n%s\n", desc->typ, desc->sdp);
    (*__peer_connection_set_local_description__)(pc, observer, __set_session_description_observer__, desc);
}

void PeerConnectionSetRemoteDescription(void *pc, void *observer, raw_session_description_t *desc)
{
    printf("===> PeerConnectionSetRemoteDescription(type=%s, sdp=\n%s\n", desc->typ, desc->sdp);
    (*__peer_connection_set_remote_description__)(pc, observer, __set_session_description_observer__, desc);
}

int PeerConnectionAddIceCandidate(void *pc, raw_ice_candidate_t *candidate, raw_rtc_error_t *err)
{
    printf("===> PeerConnectionAddIceCandidate(%s)\n", candidate->candidate);
    return (*__peer_connection_add_ice_candidate__)(pc, candidate, err);
}

void PeerConnectionGetReceivers(void *pc, size_t *size, void **array)
{
    printf("===> PeerConnectionGetReceivers()\n");
    return (*__peer_connection_get_receivers__)(pc, size, array);
}

void PeerConnectionGetSenders(void *pc, size_t *size, void **array)
{
    printf("===> PeerConnectionGetSenders()\n");
    return (*__peer_connection_get_senders__)(pc, size, array);
}

void PeerConnectionGetTransceivers(void *pc, size_t *size, void **array)
{
    printf("===> PeerConnectionGetTransceivers()\n");
    return (*__peer_connection_get_transceivers__)(pc, size, array);
}

void PeerConnectionClose(void *pc)
{
    printf("===> PeerConnectionClose()\n");
    (*__peer_connection_close__)(pc);
}

const char *MediaStreamGetID(void *stream)
{
    printf("===> MediaStreamGetID()\n");
    return (*__media_stream_get_id__)(stream);
}

int MediaStreamAddTrack(void *stream, void *track)
{
    printf("===> MediaStreamAddTrack()\n");
    return (*__media_stream_add_track__)(stream, track);
}

int MediaStreamRemoveTrack(void *stream, void *track)
{
    printf("===> MediaStreamRemoveTrack()\n");
    return (*__media_stream_remove_track__)(stream, track);
}

void MediaStreamGetAudioTracks(void *stream, size_t *size, void **array)
{
    printf("===> MediaStreamGetAudioTracks()\n");
    (*__media_stream_get_audio_tracks__)(stream, size, array);
}

void MediaStreamGetVideoTracks(void *stream, size_t *size, void **array)
{
    printf("===> MediaStreamGetVideoTracks()\n");
    (*__media_stream_get_video_tracks__)(stream, size, array);
}

void *MediaStreamFindAudioTrack(void *stream, const char *id)
{
    printf("===> MediaStreamFindAudioTrack(%s)\n", id);
    (*__media_stream_find_audio_track__)(stream, id);
}

void *MediaStreamFindVideoTrack(void *stream, const char *id)
{
    printf("===> MediaStreamFindVideoTrack(%s)\n", id);
    (*__media_stream_find_video_track__)(stream, id);
}

const char *MediaStreamTrackGetID(void *track)
{
    printf("===> MediaStreamTrackGetID()\n");
    (*__media_stream_track_get_id__)(track);
}

const char *MediaStreamTrackGetKind(void *track)
{
    printf("===> MediaStreamTrackGetKind()\n");
    (*__media_stream_track_get_kind__)(track);
}

int MediaStreamTrackGetMuted(void *track)
{
    printf("===> MediaStreamTrackGetMuted()\n");
    (*__media_stream_track_get_muted__)(track);
}

const char *MediaStreamTrackGetState(void *track)
{
    printf("===> MediaStreamTrackGetState()\n");
    (*__media_stream_track_get_state__)(track);
}

void *MediaStreamTrackGetSource(void *track)
{
    printf("===> MediaStreamTrackGetSource()\n");
    (*__media_stream_track_get_source__)(track);
}

void MediaStreamTrackStop(void *track)
{
    printf("===> MediaStreamTrackStop()\n");
    (*__media_stream_track_stop__)(track);
}

const char *RtpTransceiverGetDirection(void *transceiver)
{
    printf("===> RtpTransceiverGetDirection()\n");
    (*__rtp_transceiver_get_direction__)(transceiver);
}

const char *RtpTransceiverGetMid(void *transceiver)
{
    printf("===> RtpTransceiverGetMid()\n");
    (*__rtp_transceiver_get_mid__)(transceiver);
}

void *RtpTransceiverGetReceiver(void *transceiver)
{
    printf("===> RtpTransceiverGetReceiver()\n");
    (*__rtp_transceiver_get_receiver__)(transceiver);
}

void *RtpTransceiverGetSender(void *transceiver)
{
    printf("===> RtpTransceiverGetSender()\n");
    (*__rtp_transceiver_get_sender__)(transceiver);
}

int RtpTransceiverSetDirection(void *transceiver, const char *new_direction, raw_rtc_error_t *err)
{
    printf("===> RtpTransceiverSetDirection()\n");
    (*__rtp_transceiver_set_direction__)(transceiver, new_direction, err);
}

void RtpTransceiverStop(void *transceiver)
{
    printf("===> RtpTransceiverStop()\n");
    (*__rtp_transceiver_stop__)(transceiver);
}

void *RtpReceiverGetTrack(void *receiver)
{
    printf("===> RtpReceiverGetTrack()\n");
    (*__rtp_receiver_get_track__)(receiver);
}

void RtpReceiverGetStreams(void *receiver, size_t *size, void **array)
{
    printf("===> RtpReceiverGetStreams()\n");
    (*__rtp_receiver_get_streams__)(receiver, size, array);
}

void RtpReceiverGetParameters(void *receiver, void *parameters)
{
    printf("===> RtpReceiverGetParameters()\n");
    (*__rtp_receiver_get_parameters__)(receiver, parameters);
}

void RtpReceiverGetStats(void *receiver, void *stats)
{
    printf("===> RtpReceiverGetStats()\n");
    (*__rtp_receiver_get_stats__)(receiver, stats);
}

int RtpSenderSetTrack(void *sender, void *track)
{
    printf("===> RtpSenderSetTrack()\n");
    (*__rtp_sender_set_track__)(sender, track);
}

void *RtpSenderGetTrack(void *sender)
{
    printf("===> RtpSenderGetTrack()\n");
    (*__rtp_sender_get_track__)(sender);
}

void RtpSenderSetStreams(void *sender, size_t size, const char **stream_ids)
{
    printf("===> RtpSenderSetStreams()\n");
    (*__rtp_sender_set_streams__)(sender, size, stream_ids);
}

void RtpSenderGetStreams(void *sender, raw_array_t *dst)
{
    printf("===> RtpSenderGetStreams()\n");
    (*__rtp_sender_get_streams__)(sender, dst);
}

void RtpSenderSetParameters(void *sender, void *parameters)
{
    printf("===> RtpSenderSetParameters()\n");
    (*__rtp_sender_set_parameters__)(sender, parameters);
}

void RtpSenderGetParameters(void *sender, void *parameters)
{
    printf("===> RtpSenderGetParameters()\n");
    (*__rtp_sender_get_parameters__)(sender, parameters);
}

void RtpSenderGetStats(void *sender, void *stats)
{
    printf("===> RtpSenderGetStats()\n");
    (*__rtp_sender_get_stats__)(sender, stats);
}
