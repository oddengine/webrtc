#include "api.h"

#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>

#ifdef _WIN32
#include <Windows.h>
#define dlsym GetProcAddress
#define dlopen(file, mode) LoadLibrary(file)
#define dlerror GetLastError
#else
#include <dlfcn.h>
#define HMODULE void *
#endif

HMODULE handle;

typedef int (*__initialize_library_fptr__)(raw_rtc_constraints_t constraints);

typedef void *(*__create_default_logger_factory_fptr__)(void *fd, void *out, int level);
typedef void *(*__create_default_logger_fptr__)(void *fd, void *factory, const char *scope);
typedef void *(*__create_default_writer_fptr__)(void *fd, raw_default_writer_constraints_t *constraints, raw_default_writer_observer_t *cb);

typedef int (*__writer_open_fptr__)(void *writer, const char *path);
typedef int (*__writer_write_fptr__)(void *writer, const char *message, size_t size);
typedef size_t (*__writer_size_fptr__)(void *writer);
typedef int (*__writer_close_fptr__)(void *writer);

typedef void *(*__create_peer_connection_factory_fptr__)(void *fd);
typedef raw_rtp_capabilities_t (*__get_rtp_sender_capabilities_fptr__)(void *factory, const char *kind);
typedef raw_rtp_capabilities_t (*__get_rtp_receiver_capabilities_fptr__)(void *factory, const char *kind);
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
typedef void (*__rtp_transceiver_set_codec_preferences_fptr__)(void *transceiver, void **codecs, size_t size);
typedef int (*__rtp_transceiver_set_direction_fptr__)(void *transceiver, const char *new_direction, raw_rtc_error_t *err);
typedef void (*__rtp_transceiver_stop_fptr__)(void *transceiver);

typedef void *(*__rtp_receiver_get_track_fptr__)(void *receiver);
typedef void (*__rtp_receiver_get_streams_fptr__)(void *receiver, size_t *size, void **array);
typedef void (*__rtp_receiver_get_parameters_fptr__)(void *receiver, void *parameters);
typedef void (*__rtp_receiver_get_stats_fptr__)(void *receiver, void *stats);

typedef int (*__rtp_sender_set_track_fptr__)(void *sender, void *track);
typedef void *(*__rtp_sender_get_track_fptr__)(void *sender);
typedef void (*__rtp_sender_set_streams_fptr__)(void *sender, size_t size, const char **stream_ids);
typedef void (*__rtp_sender_get_streams_fptr__)(void *sender, size_t *size, void **array);
typedef void (*__rtp_sender_set_parameters_fptr__)(void *sender, void *parameters);
typedef void (*__rtp_sender_get_parameters_fptr__)(void *sender, void *parameters);
typedef void (*__rtp_sender_get_stats_fptr__)(void *sender, void *stats);

__initialize_library_fptr__ __initialize_library__;

__create_default_logger_factory_fptr__ __create_default_logger_factory__;
__create_default_logger_fptr__ __create_default_logger__;
__create_default_writer_fptr__ __create_default_writer__;

__writer_open_fptr__ __writer_open__;
__writer_write_fptr__ __writer_write__;
__writer_size_fptr__ __writer_size__;
__writer_close_fptr__ __writer_close__;

__create_peer_connection_factory_fptr__ __create_peer_connection_factory__;
__get_rtp_sender_capabilities_fptr__ __get_rtp_sender_capabilities__;
__get_rtp_receiver_capabilities_fptr__ __get_rtp_receiver_capabilities__;
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
__rtp_transceiver_set_codec_preferences_fptr__ __rtp_transceiver_set_codec_preferences__;
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

raw_default_writer_observer_t *__logger_writer_observer__;
extern void __onloggerwriterresize__(void *target);

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

int InitializeLibrary(const char *file, raw_rtc_constraints_t constraints)
{
    handle = dlopen(file, 1);
    if (handle == NULL)
    {
        printf("Failed to open library: file=%s, "
#ifdef _WIN32
               "eno=%d"
#else
               "err=%s"
#endif
               "\n",
               file, dlerror());
        return -1;
    }

    __initialize_library__ = (__initialize_library_fptr__)dlsym(handle, "InitializeLibrary");

    __create_default_logger_factory__ = (__create_default_logger_factory_fptr__)dlsym(handle, "CreateDefaultLoggerFactory");
    __create_default_logger__ = (__create_default_logger_fptr__)dlsym(handle, "CreateDefaultLogger");
    __create_default_writer__ = (__create_default_writer_fptr__)dlsym(handle, "CreateDefaultWriter");

    __writer_open__ = (__writer_open_fptr__)dlsym(handle, "WriterOpen");
    __writer_write__ = (__writer_write_fptr__)dlsym(handle, "WriterWrite");
    __writer_size__ = (__writer_size_fptr__)dlsym(handle, "WriterSize");
    __writer_close__ = (__writer_close_fptr__)dlsym(handle, "WriterClose");

    __create_peer_connection_factory__ = (__create_peer_connection_factory_fptr__)dlsym(handle, "CreatePeerConnectionFactory");
    __get_rtp_sender_capabilities__ = (__get_rtp_sender_capabilities_fptr__)dlsym(handle, "GetRtpSenderCapabilities");
    __get_rtp_receiver_capabilities__ = (__get_rtp_receiver_capabilities_fptr__)dlsym(handle, "GetRtpReceiverCapabilities");
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
    __rtp_transceiver_set_codec_preferences__ = (__rtp_transceiver_set_codec_preferences_fptr__)dlsym(handle, "RtpTransceiverSetCodecPreferences");
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

    __logger_writer_observer__ = malloc(sizeof(raw_default_writer_observer_t));
    __logger_writer_observer__->onresize = __onloggerwriterresize__;

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

    return __initialize_library__(constraints);
}

void *CreateDefaultLoggerFactory(void *fd, void *out, int level)
{
    // __debugf__(6, "===> CreateDefaultLoggerFactory()");
    return __create_default_logger_factory__(fd, out, level);
}

void *CreateDefaultLogger(void *fd, void *factory, const char *scope)
{
    // __debugf__(6, "===> CreateDefaultLogger()");
    return __create_default_logger__(fd, factory, scope);
}

void *CreateDefaultWriter(void *fd, raw_default_writer_constraints_t *constraints)
{
    // __debugf__(6, "===> CreateDefaultWriter()");
    return __create_default_writer__(fd, constraints, __logger_writer_observer__);
}

int WriterOpen(void *writer, const char *path)
{
    // __debugf__(6, "===> WriterOpen()");
    return __writer_open__(writer, path);
}

int WriterWrite(void *writer, const char *message, size_t size)
{
    // __debugf__(6, "===> WriterWrite()");
    return __writer_write__(writer, message, size);
}

size_t WriterSize(void *writer)
{
    // __debugf__(6, "===> WriterSize()");
    return __writer_size__(writer);
}

int WriterClose(void *writer)
{
    // __debugf__(6, "===> WriterClose()");
    return __writer_close__(writer);
}

void *CreatePeerConnectionFactory(void *fd)
{
    // __debugf__(6, "===> CreatePeerConnectionFactory()");
    return __create_peer_connection_factory__(fd);
}

raw_rtp_capabilities_t GetRtpSenderCapabilities(void *factory, const char *kind)
{
    // __debugf__(6, "===> GetRtpSenderCapabilities()");
    return __get_rtp_sender_capabilities__(factory, kind);
}

raw_rtp_capabilities_t GetRtpReceiverCapabilities(void *factory, const char *kind)
{
    // __debugf__(6, "===> GetRtpReceiverCapabilities()");
    return __get_rtp_receiver_capabilities__(factory, kind);
}

void *CreatePeerConnection(void *factory, void *pc)
{
    // __debugf__(6, "===> CreatePeerConnection()");
    return __create_peer_connection__(factory, pc, __peer_connection_observer__);
}

void *CreateAudioTrack(void *factory, void *track, const char *id, void *source)
{
    // __debugf__(6, "===> CreateAudioTrack(%s)", id);
    return __create_audio_track__(factory, track, id, source);
}

void *CreateVideoTrack(void *factory, void *track, const char *id, void *source)
{
    // __debugf__(6, "===> CreateVideoTrack(%s)", id);
    return __create_video_track__(factory, track, id, source);
}

void *PeerConnectionAddTrack(void *pc, void *track, size_t size, void **streams, raw_rtc_error_t *err)
{
    // __debugf__(6, "===> PeerConnectionAddTrack()");
    return __peer_connection_add_track__(pc, track, size, streams, err);
}

int PeerConnectionRemoveTrack(void *pc, void *sender, raw_rtc_error_t *err)
{
    // __debugf__(6, "===> PeerConnectionRemoveTrack()");
    return __peer_connection_remove_track__(pc, sender, err);
}

void *PeerConnectionAddTransceiver(void *pc, const char *media_type, raw_rtp_transceiver_init_t *init, raw_rtc_error_t *err)
{
    // __debugf__(6, "===> PeerConnectionAddTransceiver(%s)", media_type);
    return __peer_connection_add_transceiver__(pc, media_type, init, err);
}

void PeerConnectionCreateOffer(void *pc, void *observer)
{
    // __debugf__(6, "===> PeerConnectionCreateOffer()");
    __peer_connection_create_offer__(pc, observer, __create_session_description_observer__);
}

void PeerConnectionCreateAnswer(void *pc, void *observer)
{
    // __debugf__(6, "===> PeerConnectionCreateAnswer()");
    __peer_connection_create_answer__(pc, observer, __create_session_description_observer__);
}

void PeerConnectionSetLocalDescription(void *pc, void *observer, raw_session_description_t *desc)
{
    // __debugf__(6, "===> PeerConnectionSetLocalDescription(type=%s, sdp=\n%s", desc->typ, desc->sdp);
    __peer_connection_set_local_description__(pc, observer, __set_session_description_observer__, desc);
}

void PeerConnectionSetRemoteDescription(void *pc, void *observer, raw_session_description_t *desc)
{
    // __debugf__(6, "===> PeerConnectionSetRemoteDescription(type=%s, sdp=\n%s", desc->typ, desc->sdp);
    __peer_connection_set_remote_description__(pc, observer, __set_session_description_observer__, desc);
}

int PeerConnectionAddIceCandidate(void *pc, raw_ice_candidate_t *candidate, raw_rtc_error_t *err)
{
    // __debugf__(6, "===> PeerConnectionAddIceCandidate(%s)", candidate->candidate);
    return __peer_connection_add_ice_candidate__(pc, candidate, err);
}

void PeerConnectionGetReceivers(void *pc, size_t *size, void **array)
{
    // __debugf__(6, "===> PeerConnectionGetReceivers()");
    __peer_connection_get_receivers__(pc, size, array);
}

void PeerConnectionGetSenders(void *pc, size_t *size, void **array)
{
    // __debugf__(6, "===> PeerConnectionGetSenders()");
    __peer_connection_get_senders__(pc, size, array);
}

void PeerConnectionGetTransceivers(void *pc, size_t *size, void **array)
{
    // __debugf__(6, "===> PeerConnectionGetTransceivers()");
    __peer_connection_get_transceivers__(pc, size, array);
}

void PeerConnectionClose(void *pc)
{
    // __debugf__(6, "===> PeerConnectionClose()");
    __peer_connection_close__(pc);
}

const char *MediaStreamGetID(void *stream)
{
    // __debugf__(6, "===> MediaStreamGetID()");
    return __media_stream_get_id__(stream);
}

int MediaStreamAddTrack(void *stream, void *track)
{
    // __debugf__(6, "===> MediaStreamAddTrack()");
    return __media_stream_add_track__(stream, track);
}

int MediaStreamRemoveTrack(void *stream, void *track)
{
    // __debugf__(6, "===> MediaStreamRemoveTrack()");
    return __media_stream_remove_track__(stream, track);
}

void MediaStreamGetAudioTracks(void *stream, size_t *size, void **array)
{
    // __debugf__(6, "===> MediaStreamGetAudioTracks()");
    __media_stream_get_audio_tracks__(stream, size, array);
}

void MediaStreamGetVideoTracks(void *stream, size_t *size, void **array)
{
    // __debugf__(6, "===> MediaStreamGetVideoTracks()");
    __media_stream_get_video_tracks__(stream, size, array);
}

void *MediaStreamFindAudioTrack(void *stream, const char *id)
{
    // __debugf__(6, "===> MediaStreamFindAudioTrack(%s)", id);
    return __media_stream_find_audio_track__(stream, id);
}

void *MediaStreamFindVideoTrack(void *stream, const char *id)
{
    // __debugf__(6, "===> MediaStreamFindVideoTrack(%s)", id);
    return __media_stream_find_video_track__(stream, id);
}

const char *MediaStreamTrackGetID(void *track)
{
    // __debugf__(6, "===> MediaStreamTrackGetID()");
    return __media_stream_track_get_id__(track);
}

const char *MediaStreamTrackGetKind(void *track)
{
    // __debugf__(6, "===> MediaStreamTrackGetKind()");
    return __media_stream_track_get_kind__(track);
}

int MediaStreamTrackGetMuted(void *track)
{
    // __debugf__(6, "===> MediaStreamTrackGetMuted()");
    return __media_stream_track_get_muted__(track);
}

const char *MediaStreamTrackGetState(void *track)
{
    // __debugf__(6, "===> MediaStreamTrackGetState()");
    return __media_stream_track_get_state__(track);
}

void *MediaStreamTrackGetSource(void *track)
{
    // __debugf__(6, "===> MediaStreamTrackGetSource()");
    return __media_stream_track_get_source__(track);
}

void MediaStreamTrackStop(void *track)
{
    // __debugf__(6, "===> MediaStreamTrackStop()");
    __media_stream_track_stop__(track);
}

const char *RtpTransceiverGetDirection(void *transceiver)
{
    // __debugf__(6, "===> RtpTransceiverGetDirection()");
    return __rtp_transceiver_get_direction__(transceiver);
}

const char *RtpTransceiverGetMid(void *transceiver)
{
    // __debugf__(6, "===> RtpTransceiverGetMid()");
    return __rtp_transceiver_get_mid__(transceiver);
}

void *RtpTransceiverGetReceiver(void *transceiver)
{
    // __debugf__(6, "===> RtpTransceiverGetReceiver()");
    return __rtp_transceiver_get_receiver__(transceiver);
}

void *RtpTransceiverGetSender(void *transceiver)
{
    // __debugf__(6, "===> RtpTransceiverGetSender()");
    return __rtp_transceiver_get_sender__(transceiver);
}

void RtpTransceiverSetCodecPreferences(void *transceiver, void **codecs, size_t size)
{
    // __debugf__(6, "===> RtpTransceiverSetCodecPreferences()");
    return __rtp_transceiver_set_codec_preferences__(transceiver, codecs, size);
}

int RtpTransceiverSetDirection(void *transceiver, const char *new_direction, raw_rtc_error_t *err)
{
    // __debugf__(6, "===> RtpTransceiverSetDirection()");
    return __rtp_transceiver_set_direction__(transceiver, new_direction, err);
}

void RtpTransceiverStop(void *transceiver)
{
    // __debugf__(6, "===> RtpTransceiverStop()");
    __rtp_transceiver_stop__(transceiver);
}

void *RtpReceiverGetTrack(void *receiver)
{
    // __debugf__(6, "===> RtpReceiverGetTrack()");
    return __rtp_receiver_get_track__(receiver);
}

void RtpReceiverGetStreams(void *receiver, size_t *size, void **array)
{
    // __debugf__(6, "===> RtpReceiverGetStreams()");
    __rtp_receiver_get_streams__(receiver, size, array);
}

void RtpReceiverGetParameters(void *receiver, void *parameters)
{
    // __debugf__(6, "===> RtpReceiverGetParameters()");
    __rtp_receiver_get_parameters__(receiver, parameters);
}

void RtpReceiverGetStats(void *receiver, void *stats)
{
    // __debugf__(6, "===> RtpReceiverGetStats()");
    __rtp_receiver_get_stats__(receiver, stats);
}

int RtpSenderSetTrack(void *sender, void *track)
{
    // __debugf__(6, "===> RtpSenderSetTrack()");
    return __rtp_sender_set_track__(sender, track);
}

void *RtpSenderGetTrack(void *sender)
{
    // __debugf__(6, "===> RtpSenderGetTrack()");
    return __rtp_sender_get_track__(sender);
}

void RtpSenderSetStreams(void *sender, size_t size, const char **stream_ids)
{
    // __debugf__(6, "===> RtpSenderSetStreams()");
    __rtp_sender_set_streams__(sender, size, stream_ids);
}

void RtpSenderGetStreams(void *sender, size_t *size, void **array)
{
    // __debugf__(6, "===> RtpSenderGetStreams()");
    __rtp_sender_get_streams__(sender, size, array);
}

void RtpSenderSetParameters(void *sender, void *parameters)
{
    // __debugf__(6, "===> RtpSenderSetParameters()");
    __rtp_sender_set_parameters__(sender, parameters);
}

void RtpSenderGetParameters(void *sender, void *parameters)
{
    // __debugf__(6, "===> RtpSenderGetParameters()");
    __rtp_sender_get_parameters__(sender, parameters);
}

void RtpSenderGetStats(void *sender, void *stats)
{
    // __debugf__(6, "===> RtpSenderGetStats()");
    __rtp_sender_get_stats__(sender, stats);
}
