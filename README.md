<h1 align="center">
  <img src="logo.png" alt="MediaMTX / rtsp-simple-server">

  <br>
  <br>

  [![Test](https://github.com/bluenviron/mediamtx/workflows/test/badge.svg)](https://github.com/bluenviron/mediamtx/actions?query=workflow:test)
  [![Lint](https://github.com/bluenviron/mediamtx/workflows/lint/badge.svg)](https://github.com/bluenviron/mediamtx/actions?query=workflow:lint)
  [![CodeCov](https://codecov.io/gh/bluenviron/mediamtx/branch/main/graph/badge.svg)](https://app.codecov.io/gh/bluenviron/mediamtx/branch/main)
  [![Release](https://img.shields.io/github/v/release/bluenviron/mediamtx)](https://github.com/bluenviron/mediamtx/releases)
  [![Docker Hub](https://img.shields.io/badge/docker-bluenviron/mediamtx-blue)](https://hub.docker.com/r/bluenviron/mediamtx)
  [![API Documentation](https://img.shields.io/badge/api-documentation-blue)](https://bluenviron.github.io/mediamtx)
</h1>

<br>

_MediaMTX_ (formerly _rtsp-simple-server_) is a ready-to-use and zero-dependency real-time media server and media proxy that allows to publish, read, proxy, record and playback video and audio streams. It has been conceived as a "media router" that routes media streams from one end to the other.

Live streams can be published to the server with:

|protocol|variants|video codecs|audio codecs|
|--------|--------|------------|------------|
|[SRT clients](#srt-clients)||H265, H264, MPEG-4 Video (H263, Xvid), MPEG-1/2 Video|Opus, MPEG-4 Audio (AAC), MPEG-1/2 Audio (MP3), AC-3|
|[SRT cameras and servers](#srt-cameras-and-servers)||H265, H264, MPEG-4 Video (H263, Xvid), MPEG-1/2 Video|Opus, MPEG-4 Audio (AAC), MPEG-1/2 Audio (MP3), AC-3|
|[WebRTC clients](#webrtc-clients)|Browser-based, WHIP|AV1, VP9, VP8, H264|Opus, G722, G711 (PCMA, PCMU)|
|[WebRTC servers](#webrtc-servers)|WHEP|AV1, VP9, VP8, H264|Opus, G722, G711 (PCMA, PCMU)|
|[RTSP clients](#rtsp-clients)|UDP, TCP, RTSPS|AV1, VP9, VP8, H265, H264, MPEG-4 Video (H263, Xvid), MPEG-1/2 Video, M-JPEG and any RTP-compatible codec|Opus, MPEG-4 Audio (AAC), MPEG-1/2 Audio (MP3), AC-3, G726, G722, G711 (PCMA, PCMU), LPCM and any RTP-compatible codec|
|[RTSP cameras and servers](#rtsp-cameras-and-servers)|UDP, UDP-Multicast, TCP, RTSPS|AV1, VP9, VP8, H265, H264, MPEG-4 Video (H263, Xvid), MPEG-1/2 Video, M-JPEG and any RTP-compatible codec|Opus, MPEG-4 Audio (AAC), MPEG-1/2 Audio (MP3), AC-3, G726, G722, G711 (PCMA, PCMU), LPCM and any RTP-compatible codec|
|[RTMP clients](#rtmp-clients)|RTMP, RTMPS, Enhanced RTMP|AV1, VP9, H265, H264|MPEG-4 Audio (AAC), MPEG-1/2 Audio (MP3), G711 (PCMA, PCMU), LPCM|
|[RTMP cameras and servers](#rtmp-cameras-and-servers)|RTMP, RTMPS, Enhanced RTMP|H264|MPEG-4 Audio (AAC), MPEG-1/2 Audio (MP3)|
|[HLS cameras and servers](#hls-cameras-and-servers)|Low-Latency HLS, MP4-based HLS, legacy HLS|AV1, VP9, H265, H264|Opus, MPEG-4 Audio (AAC)|
|[UDP/MPEG-TS](#udpmpeg-ts)|Unicast, broadcast, multicast|H265, H264, MPEG-4 Video (H263, Xvid), MPEG-1/2 Video|Opus, MPEG-4 Audio (AAC), MPEG-1/2 Audio (MP3), AC-3|
|[Raspberry Pi Cameras](#raspberry-pi-cameras)||H264||

And can be read from the server with:

|protocol|variants|video codecs|audio codecs|
|--------|--------|------------|------------|
|[SRT](#srt)||H265, H264, MPEG-4 Video (H263, Xvid), MPEG-1/2 Video|Opus, MPEG-4 Audio (AAC), MPEG-1/2 Audio (MP3), AC-3|
|[WebRTC](#webrtc)|Browser-based, WHEP|AV1, VP9, VP8, H264|Opus, G722, G711 (PCMA, PCMU)|
|[RTSP](#rtsp)|UDP, UDP-Multicast, TCP, RTSPS|AV1, VP9, VP8, H265, H264, MPEG-4 Video (H263, Xvid), MPEG-1/2 Video, M-JPEG and any RTP-compatible codec|Opus, MPEG-4 Audio (AAC), MPEG-1/2 Audio (MP3), AC-3, G726, G722, G711 (PCMA, PCMU), LPCM and any RTP-compatible codec|
|[RTMP](#rtmp)|RTMP, RTMPS, Enhanced RTMP|H264|MPEG-4 Audio (AAC), MPEG-1/2 Audio (MP3)|
|[HLS](#hls)|Low-Latency HLS, MP4-based HLS, legacy HLS|AV1, VP9, H265, H264|Opus, MPEG-4 Audio (AAC)|

And can be recorded and played back with:

|format|video codecs|audio codecs|
|------|------------|------------|
|[fMP4](#record-streams-to-disk)|AV1, VP9, H265, H264, MPEG-4 Video (H263, Xvid), MPEG-1/2 Video, M-JPEG|Opus, MPEG-4 Audio (AAC), MPEG-1/2 Audio (MP3), AC-3, G711 (PCMA, PCMU), LPCM|
|[MPEG-TS](#record-streams-to-disk)|H265, H264, MPEG-4 Video (H263, Xvid), MPEG-1/2 Video|Opus, MPEG-4 Audio (AAC), MPEG-1/2 Audio (MP3), AC-3|

**Features**

* Publish live streams to the server
* Read live streams from the server
* Streams are automatically converted from a protocol to another
* Serve multiple streams at once in separate paths
* Record streams to disk
* Playback recorded streams
* Authenticate users; use internal or external authentication
* Redirect readers to other RTSP servers (load balancing)
* Control the server through the Control API
* Reload the configuration without disconnecting existing clients (hot reloading)
* Read Prometheus-compatible metrics
* Run hooks (external commands) when clients connect, disconnect, read or publish streams
* Compatible with Linux, Windows and macOS, does not require any dependency or interpreter, it's a single executable

**Note about rtsp-simple-server**

_rtsp-simple-server_ has been rebranded as _MediaMTX_. The reason is pretty obvious: this project started as a RTSP server but has evolved into a much more versatile product that is not tied to the RTSP protocol anymore. Nothing will change regarding license, features and backward compatibility.

## Table of contents

* [Installation](#installation)
  * [Standalone binary](#standalone-binary)
  * [Docker image](#docker-image)
  * [Arch Linux package](#arch-linux-package)
  * [OpenWrt binary](#openwrt-binary)
* [Basic usage](#basic-usage)
* [Publish to the server](#publish-to-the-server)
  * [By software](#by-software)
    * [FFmpeg](#ffmpeg)
    * [GStreamer](#gstreamer)
    * [OBS Studio](#obs-studio)
    * [OpenCV](#opencv)
    * [Web browsers](#web-browsers)
  * [By device](#by-device)
    * [Generic webcam](#generic-webcam)
    * [Raspberry Pi Cameras](#raspberry-pi-cameras)
  * [By protocol](#by-protocol)
    * [SRT clients](#srt-clients)
    * [SRT cameras and servers](#srt-cameras-and-servers)
    * [WebRTC clients](#webrtc-clients)
    * [WebRTC servers](#webrtc-servers)
    * [RTSP clients](#rtsp-clients)
    * [RTSP cameras and servers](#rtsp-cameras-and-servers)
    * [RTMP clients](#rtmp-clients)
    * [RTMP cameras and servers](#rtmp-cameras-and-servers)
    * [HLS cameras and servers](#hls-cameras-and-servers)
    * [UDP/MPEG-TS](#udpmpeg-ts)
* [Read from the server](#read-from-the-server)
  * [By software](#by-software-1)
    * [FFmpeg](#ffmpeg-1)
    * [GStreamer](#gstreamer-1)
    * [VLC](#vlc)
    * [Web browsers](#web-browsers-1)
  * [By protocol](#by-protocol-1)
    * [SRT](#srt)
    * [WebRTC](#webrtc)
    * [RTSP](#rtsp)
    * [RTMP](#rtmp)
    * [HLS](#hls)
* [Other features](#other-features)
  * [Configuration](#configuration)
  * [Authentication](#authentication)
  * [Encrypt the configuration](#encrypt-the-configuration)
  * [Remuxing, re-encoding, compression](#remuxing-re-encoding-compression)
  * [Record streams to disk](#record-streams-to-disk)
  * [Playback recorded streams](#playback-recorded-streams)
  * [Forward streams to other servers](#forward-streams-to-other-servers)
  * [Proxy requests to other servers](#proxy-requests-to-other-servers)
  * [On-demand publishing](#on-demand-publishing)
  * [Start on boot](#start-on-boot)
    * [Linux](#linux)
    * [OpenWrt](#openwrt)
    * [Windows](#windows)
  * [Hooks](#hooks)
  * [Control API](#control-api)
  * [Metrics](#metrics)
  * [pprof](#pprof)
  * [SRT-specific features](#srt-specific-features)
    * [Standard stream ID syntax](#standard-stream-id-syntax)
  * [WebRTC-specific features](#webrtc-specific-features)
    * [Connectivity issues](#connectivity-issues)
  * [RTSP-specific features](#rtsp-specific-features)
    * [Transport protocols](#transport-protocols)
    * [Encryption](#encryption)
    * [Corrupted frames](#corrupted-frames)
  * [RTMP-specific features](#rtmp-specific-features)
    * [Encryption](#encryption-1)
* [Compile from source](#compile-from-source)
  * [Standard](#standard)
  * [Raspberry Pi](#raspberry-pi)
  * [OpenWrt](#openwrt-1)
  * [Cross compile](#cross-compile)
  * [Compile for all supported platforms](#compile-for-all-supported-platforms)
* [Specifications](#specifications)
* [Related projects](#related-projects)

## Installation

There are several installation methods available: standalone binary, Docker image, Arch Linux package and OpenWrt binary.

### Standalone binary

1. Download and extract a standalone binary from the [release page](https://github.com/bluenviron/mediamtx/releases) that corresponds to your operating system and architecture.

2. Start the server:

   ```sh
   ./mediamtx
   ```

### Docker image

Download and launch the image:

```
docker run --rm -it --network=host bluenviron/mediamtx:latest
```

Available images:

|name|FFmpeg included|RPI Camera support|
|----|---------------|------------------|
|bluenviron/mediamtx:latest|:x:|:x:|
|bluenviron/mediamtx:latest-ffmpeg|:heavy_check_mark:|:x:|
|bluenviron/mediamtx:latest-rpi|:x:|:heavy_check_mark:|
|bluenviron/mediamtx:latest-ffmpeg-rpi|:heavy_check_mark:|:heavy_check_mark:|

The `--network=host` flag is mandatory since Docker can change the source port of UDP packets for routing reasons, and this doesn't allow the RTSP server to identify the senders of the packets. This issue can be avoided by disabling the UDP transport protocol:

```
docker run --rm -it \
-e MTX_PROTOCOLS=tcp \
-e MTX_WEBRTCADDITIONALHOSTS=192.168.x.x \
-p 8554:8554 \
-p 1935:1935 \
-p 8888:8888 \
-p 8889:8889 \
-p 8890:8890/udp \
-p 8189:8189/udp \
bluenviron/mediamtx
```

set `MTX_WEBRTCADDITIONALHOSTS` to your local IP address.

### Arch Linux package

If you are running the Arch Linux distribution, run:

```sh
git clone https://aur.archlinux.org/mediamtx.git
cd mediamtx
makepkg -si
```

### OpenWrt binary

If the architecture of the OpenWrt device is amd64, armv6, armv7 or arm64, use the [standalone binary method](#standalone-binary) and download a Linux binary that corresponds to your architecture.

Otherwise, [compile the server from source](#openwrt-1).

## Basic usage

1. Publish a stream. For instance, you can publish a video/audio file with _FFmpeg_:

   ```sh
   ffmpeg -re -stream_loop -1 -i file.ts -c copy -f rtsp rtsp://localhost:8554/mystream
   ```

   or _GStreamer_:

   ```sh
   gst-launch-1.0 rtspclientsink name=s location=rtsp://localhost:8554/mystream filesrc location=file.mp4 \
   ! qtdemux name=d d.video_0 ! queue ! s.sink_0 d.audio_0 ! queue ! s.sink_1
   ```

2. Open the stream. For instance, you can open the stream with _VLC_:

   ```sh
   vlc --network-caching=50 rtsp://localhost:8554/mystream
   ```

   or _GStreamer_:

   ```sh
   gst-play-1.0 rtsp://localhost:8554/mystream
   ```

   or _FFmpeg_:

   ```sh
   ffmpeg -i rtsp://localhost:8554/mystream -c copy output.mp4
   ```

## Publish to the server

### By software

#### FFmpeg

FFmpeg can publish a stream to the server in multiple ways (SRT client, SRT server, RTSP client, RTMP client, UDP/MPEG-TS, WebRTC with WHIP). The recommended one consists in publishing as a [RTSP client](#rtsp-clients):

```
ffmpeg -re -stream_loop -1 -i file.ts -c copy -f rtsp rtsp://localhost:8554/mystream
```

The RTSP protocol supports multiple underlying transport protocols, each with its own characteristics (see [RTSP-specific features](#rtsp-specific-features)). You can set the transport protocol by using the `rtsp_transport` flag, for instance, in order to use TCP:

```sh
ffmpeg -re -stream_loop -1 -i file.ts -c copy -f rtsp -rtsp_transport tcp rtsp://localhost:8554/mystream
```

The resulting stream will be available in path `/mystream`.

#### GStreamer

GStreamer can publish a stream to the server in multiple ways (SRT client, SRT server, RTSP client, RTMP client, UDP/MPEG-TS, WebRTC with WHIP). The recommended one consists in publishing as a [RTSP client](#rtsp-clients):

```sh
gst-launch-1.0 rtspclientsink name=s location=rtsp://localhost:8554/mystream \
filesrc location=file.mp4 ! qtdemux name=d \
d.video_0 ! queue ! s.sink_0 \
d.audio_0 ! queue ! s.sink_1
```

If the stream is video only:

```sh
gst-launch-1.0 filesrc location=file.mp4 ! qtdemux name=d \
d.video_0 ! rtspclientsink location=rtsp://localhost:8554/mystream
```

The RTSP protocol supports multiple underlying transport protocols, each with its own characteristics (see [RTSP-specific features](#rtsp-specific-features)). You can set the transport protocol by using the `protocols` flag:

```sh
gst-launch-1.0 filesrc location=file.mp4 ! qtdemux name=d \
d.video_0 ! rtspclientsink protocols=tcp name=s location=rtsp://localhost:8554/mystream
```

The resulting stream will be available in path `/mystream`.

#### OBS Studio

OBS Studio can publish to the server in multiple ways (SRT client, RTMP client, WebRTC client). The recommended one consists in publishing as a [RTMP client](#rtmp-clients). In `Settings -> Stream` (or in the Auto-configuration Wizard), use the following parameters:

* Service: `Custom...`
* Server: `rtmp://localhost`
* Stream key: `mystream`

If credentials are in use, use the following parameters:

* Service: `Custom...`
* Server: `rtmp://localhost`
* Stream key: `mystream?user=myuser&pass=mypass`

Save the configuration and click `Start streaming`.

If you want to generate a stream that can be read with WebRTC, open `Settings -> Output -> Recording` and use the following parameters:

* FFmpeg output type: `Output to URL`
* File path or URL: `rtsp://localhost:8554/mystream`
* Container format: `rtsp`
* Check `show all codecs (even if potentically incompatible`
* Video encoder: `h264_nvenc (libx264)`
* Video encoder settings (if any): `bf=0`
* Audio track: `1`
* Audio encoder: `libopus`

Then use the button `Start Recording` (instead of `Start Streaming`) to start streaming.

Latest versions of OBS Studio can publish to the server with the [WebRTC / WHIP protocol](#webrtc). Use the following parameters:

* Service: `WHIP`
* Server: `http://localhost:8889/mystream/whip`

Save the configuration and click `Start streaming`.

The resulting stream will be available in path `/mystream`.

#### OpenCV

OpenCV can publish to the server through its GStreamer plugin, as a [RTSP client](#rtsp-clients). It must be compiled with GStreamer support, by following this procedure:

```sh
sudo apt install -y libgstreamer1.0-dev libgstreamer-plugins-base1.0-dev gstreamer1.0-plugins-ugly gstreamer1.0-rtsp python3-dev python3-numpy
git clone --depth=1 -b 4.5.4 https://github.com/opencv/opencv
cd opencv
mkdir build && cd build
cmake -D CMAKE_INSTALL_PREFIX=/usr -D WITH_GSTREAMER=ON ..
make -j$(nproc)
sudo make install
```

You can check that OpenCV has been installed correctly by running:

```sh
python3 -c 'import cv2; print(cv2.getBuildInformation())'
```

Check that the output contains `GStreamer: YES`.

Videos can be published with `VideoWriter`:

```python
from datetime import datetime
from time import sleep, time

import cv2
import numpy as np

fps = 15
width = 800
height = 600
colors = [
    (0, 0, 255),
    (255, 0, 0),
    (0, 255, 0),
]

out = cv2.VideoWriter('appsrc ! videoconvert' + \
    ' ! video/x-raw,format=I420' + \
    ' ! x264enc speed-preset=ultrafast bitrate=600 key-int-max=' + str(fps * 2) + \
    ' ! video/x-h264,profile=baseline' + \
    ' ! rtspclientsink location=rtsp://localhost:8554/mystream',
    cv2.CAP_GSTREAMER, 0, fps, (width, height), True)
if not out.isOpened():
    raise Exception("can't open video writer")

curcolor = 0
start = time()

while True:
    frame = np.zeros((height, width, 3), np.uint8)

    # create a rectangle
    color = colors[curcolor]
    curcolor += 1
    curcolor %= len(colors)
    for y in range(0, int(frame.shape[0] / 2)):
        for x in range(0, int(frame.shape[1] / 2)):
            frame[y][x] = color

    out.write(frame)
    print("%s frame written to the server" % datetime.now())

    now = time()
    diff = (1 / fps) - now - start
    if diff > 0:
        sleep(diff)
    start = now
```

The resulting stream will be available in path `/mystream`.

#### Web browsers

Web browsers can publish a stream to the server by using the [WebRTC protocol](#webrtc). Start the server and open the web page:

```
http://localhost:8889/mystream/publish
```

The resulting stream will be available in path `/mystream`.

This web page can be embedded into another web page by using an iframe:

```html
<iframe src="http://mediamtx-ip:8889/mystream/publish" scrolling="no"></iframe>
```

For more advanced setups, you can create and serve a custom web page by starting from the [source code of the publish page](internal/servers/webrtc/publish_index.html).

### By device

#### Generic webcam

If the OS is Linux-based, edit `mediamtx.yml` and replace everything inside section `paths` with the following content:

```yml
paths:
  cam:
    runOnInit: ffmpeg -f v4l2 -i /dev/video0 -pix_fmt yuv420p -preset ultrafast -b:v 600k -f rtsp rtsp://localhost:$RTSP_PORT/$MTX_PATH
    runOnInitRestart: yes
```

If the OS is Windows:

```yml
paths:
  cam:
    runOnInit: ffmpeg -f dshow -i video="USB2.0 HD UVC WebCam" -pix_fmt yuv420p -c:v libx264 -preset ultrafast -b:v 600k -f rtsp rtsp://localhost:$RTSP_PORT/$MTX_PATH
    runOnInitRestart: yes
```

Where `USB2.0 HD UVC WebCam` is the name of a webcam, that can be obtained with:

```sh
ffmpeg -list_devices true -f dshow -i dummy
```

The resulting stream will be available in path `/cam`.

#### Raspberry Pi Cameras

_MediaMTX_ natively supports the Raspberry Pi Camera, enabling high-quality and low-latency video streaming from the camera to any user, for any purpose. There are a couple of requirements:

1. The server must run on a Raspberry Pi, with Raspberry Pi OS bullseye or newer as operative system. Both 32 bit and 64 bit operative systems are supported.

2. Make sure that the legacy camera stack is disabled. Type `sudo raspi-config`, then go to `Interfacing options`, `enable/disable legacy camera support`, choose `no`. Reboot the system.

If you want to run the standard (non-Docker) version of the server:

1. Make sure that the following packages are installed:

   * `libcamera0` (&ge; 0.0.5)
   * `libfreetype6`

2. download the server executable. If you're using 64-bit version of the operative system, make sure to pick the `arm64` variant.

3. edit `mediamtx.yml` and replace everything inside section `paths` with the following content:

   ```yml
   paths:
     cam:
       source: rpiCamera
   ```

The resulting stream will be available in path `/cam`.

If you want to run the server inside Docker, you need to use the `latest-rpi` image (that already contains required libraries) and launch the container with some additional flags:

```sh
docker run --rm -it \
--network=host \
--privileged \
--tmpfs /dev/shm:exec \
-v /run/udev:/run/udev:ro \
-e MTX_PATHS_CAM_SOURCE=rpiCamera \
bluenviron/mediamtx:latest-rpi
```

Be aware that the Docker image is not compatible with cameras that requires a custom `libcamera` (like some ArduCam products), since it comes with a standard `libcamera` included.

Camera settings can be changed by using the `rpiCamera*` parameters:

```yml
paths:
  cam:
    source: rpiCamera
    rpiCameraWidth: 1920
    rpiCameraHeight: 1080
```

All available parameters are listed in the [sample configuration file](/mediamtx.yml).

In order to add audio from a USB microfone, install GStreamer and alsa-utils:

```sh
sudo apt install -y gstreamer1.0-tools gstreamer1.0-rtsp gstreamer1.0-alsa alsa-utils
```

list available audio cards with:

```sh
arecord -L
```

Sample output:

```
surround51:CARD=ICH5,DEV=0
    Intel ICH5, Intel ICH5
    5.1 Surround output to Front, Center, Rear and Subwoofer speakers
default:CARD=U0x46d0x809
    USB Device 0x46d:0x809, USB Audio
    Default Audio Device
```

Find the audio card of the microfone and take note of its name, for instance `default:CARD=U0x46d0x809`. Then use GStreamer inside `runOnReady` to read the video stream, add audio and publish the new stream to another path:

```yml
paths:
  cam:
    source: rpiCamera
    runOnReady: >
      gst-launch-1.0
      rtspclientsink name=s location=rtsp://localhost:$RTSP_PORT/cam_with_audio
      rtspsrc location=rtsp://127.0.0.1:$RTSP_PORT/$MTX_PATH latency=0 ! rtph264depay ! s.
      alsasrc device=default:CARD=U0x46d0x809 ! opusenc bitrate=16000 ! s.
    runOnReadyRestart: yes
  cam_with_audio:
```

The resulting stream will be available in path `/cam_with_audio`.

### By protocol

#### SRT clients

SRT is a protocol that allows to publish and read live data stream, providing encryption, integrity and a retransmission mechanism. It is usually used to transfer media streams encoded with MPEG-TS. In order to publish a stream to the server with the SRT protocol, use this URL:

```
srt://localhost:8890?streamid=publish:mystream&pkt_size=1316
```

Replace `mystream` with any name you want. The resulting stream will be available in path `/mystream`.

If credentials are enabled, append username and password to `streamid`;

```
srt://localhost:8890?streamid=publish:mystream:user:pass&pkt_size=1316
```

If you need to use the standard stream ID syntax instead of the custom one in use by this server, see [Standard stream ID syntax](#standard-stream-id-syntax).

If you want to publish a stream by using a client in listening mode (i.e. with `mode=listener` appended to the URL), read the next section.

Known clients that can publish with SRT are [FFmpeg](#ffmpeg), [GStreamer](#gstreamer), [OBS Studio](#obs-studio).

#### SRT cameras and servers

In order to ingest into the server a SRT stream from an existing server, camera or client in listening mode (i.e. with `mode=listener` appended to the URL), add the corresponding URL into the `source` parameter of a path:

```yml
paths:
  proxied:
    # url of the source stream, in the format srt://host:port?streamid=streamid&other_parameters
    source: srt://original-url
```

#### WebRTC clients

WebRTC is an API that makes use of a set of protocols and methods to connect two clients together and allow them to exchange real-time media or data streams. You can publish a stream with WebRTC and a web browser by visiting:

```
http://localhost:8889/mystream/publish
```

The resulting stream will be available in path `/mystream`.

WHIP is a WebRTC extensions that allows to publish streams by using a URL, without passing through a web page. This allows to use WebRTC as a general purpose streaming protocol. If you are using a software that supports WHIP (for instance, latest versions of OBS Studio), you can publish a stream to the server by using this URL:

```
http://localhost:8889/mystream/whip
```

Depending on the network it may be difficult to establish a connection between server and clients, see [WebRTC-specific features](#webrtc-specific-features) for remediations.

Known clients that can publish with WebRTC and WHIP are [FFmpeg](#ffmpeg), [GStreamer](#gstreamer), [OBS Studio](#obs-studio).

#### WebRTC servers

In order to ingest into the server a WebRTC stream from an existing server, add the corresponding WHEP URL into the `source` parameter of a path:

```yml
paths:
  proxied:
    # url of the source stream, in the format whep://host:port/path (HTTP) or wheps:// (HTTPS)
    source: wheps://host:port/path
```

#### RTSP clients

RTSP is a protocol that allows to publish and read streams. It supports different underlying transport protocols and allows to encrypt streams in transit (see [RTSP-specific features](#rtsp-specific-features)). In order to publish a stream to the server with the RTSP protocol, use this URL:

```
rtsp://localhost:8554/mystream
```

The resulting stream will be available in path `/mystream`.

Known clients that can publish with RTSP are [FFmpeg](#ffmpeg), [GStreamer](#gstreamer), [OBS Studio](#obs-studio).

#### RTSP cameras and servers

Most IP cameras expose their video stream by using a RTSP server that is embedded into the camera itself. In particular, cameras that are compliant to ONVIF profile S or T meet this requirement. You can use _MediaMTX_ to connect to one or multiple existing RTSP servers and read their video streams:

```yml
paths:
  proxied:
    # url of the source stream, in the format rtsp://user:pass@host:port/path
    source: rtsp://original-url
```

The resulting stream will be available in path `/proxied`.

The server supports any number of source streams (count is just limited by available hardware resources) it's enough to add additional entries to the paths section:

```yml
paths:
  proxied1:
    source: rtsp://url1

  proxied2:
    source: rtsp://url1
```

#### RTMP clients

RTMP is a protocol that allows to read and publish streams, but is less versatile and less efficient than RTSP and WebRTC (doesn't support UDP, doesn't support most RTSP codecs, doesn't support feedback mechanism). Streams can be published to the server by using the URL:

```
rtmp://localhost/mystream
```

The resulting stream will be available in path `/mystream`.

In case authentication is enabled, credentials can be passed to the server by using the `user` and `pass` query parameters:

```
rtmp://localhost/mystream?user=myuser&pass=mypass
```

Known clients that can publish with RTMP are [FFmpeg](#ffmpeg), [GStreamer](#gstreamer), [OBS Studio](#obs-studio).

#### RTMP cameras and servers

You can use _MediaMTX_ to connect to one or multiple existing RTMP servers and read their video streams:

```yml
paths:
  proxied:
    # url of the source stream, in the format rtmp://user:pass@host:port/path
    source: rtmp://original-url
```

The resulting stream will be available in path `/proxied`.

#### HLS cameras and servers

HLS is a streaming protocol that works by splitting streams into segments, and by serving these segments and a playlist with the HTTP protocol. You can use _MediaMTX_ to connect to one or multiple existing HLS servers and read their video streams:

```yml
paths:
  proxied:
    # url of the playlist of the stream, in the format http://user:pass@host:port/path
    source: http://original-url/stream/index.m3u8
```

The resulting stream will be available in path `/proxied`.

#### UDP/MPEG-TS

The server supports ingesting UDP/MPEG-TS packets (i.e. MPEG-TS packets sent with UDP). Packets can be unicast, broadcast or multicast. For instance, you can generate a multicast UDP/MPEG-TS stream with GStreamer:

```
gst-launch-1.0 -v mpegtsmux name=mux alignment=1 ! udpsink host=238.0.0.1 port=1234 \
videotestsrc ! video/x-raw,width=1280,height=720,format=I420 ! x264enc speed-preset=ultrafast bitrate=3000 key-int-max=60 ! video/x-h264,profile=high ! mux. \
audiotestsrc ! audioconvert ! avenc_aac ! mux.
```

or FFmpeg:

```
ffmpeg -re -f lavfi -i testsrc=size=1280x720:rate=30 \
-pix_fmt yuv420p -c:v libx264 -preset ultrafast -b:v 600k \
-f mpegts udp://238.0.0.1:1234?pkt_size=1316
```

Edit `mediamtx.yml` and replace everything inside section `paths` with the following content:

```yml
paths:
  mypath:
    source: udp://238.0.0.1:1234
```

The resulting stream will be available in path `/mypath`.

Known clients that can publish with WebRTC and WHIP are [FFmpeg](#ffmpeg) and [GStreamer](#gstreamer).

## Read from the server

### By software

#### FFmpeg

FFmpeg can read a stream from the server in multiple ways (RTSP, RTMP, HLS, WebRTC with WHEP, SRT). The recommended one consists in reading with [RTSP](#rtsp):

```sh
ffmpeg -i rtsp://localhost:8554/mystream -c copy output.mp4
```

The RTSP protocol supports multiple underlying transport protocols, each with its own characteristics (see [RTSP-specific features](#rtsp-specific-features)). You can set the transport protocol by using the `rtsp_transport` flag:

```sh
ffmpeg -rtsp_transport tcp -i rtsp://localhost:8554/mystream -c copy output.mp4
```

#### GStreamer

GStreamer can read a stream from the server in multiple ways (RTSP, RTMP, HLS, WebRTC with WHEP, SRT). The recommended one consists in reading with [RTSP](#rtsp):

```sh
gst-launch-1.0 rtspsrc location=rtsp://127.0.0.1:8554/mystream latency=0 ! decodebin ! autovideosink
```

The RTSP protocol supports multiple underlying transport protocols, each with its own characteristics (see [RTSP-specific features](#rtsp-specific-features)). You can change the transport protocol by using the `protocols` flag:

```sh
gst-launch-1.0 rtspsrc protocols=tcp location=rtsp://127.0.0.1:8554/mystream latency=0 ! decodebin ! autovideosink
```

If encryption is enabled, set `tls-validation-flags` to `0`:

```sh
gst-launch-1.0 rtspsrc tls-validation-flags=0 location=rtsps://ip:8322/...
```

#### VLC

VLC can read a stream from the server in multiple ways (RTSP, RTMP, HLS, SRT). The recommended one consists in reading with [RTSP](#rtsp):

```sh
vlc --network-caching=50 rtsp://localhost:8554/mystream
```

The RTSP protocol supports multiple underlying transport protocols, each with its own characteristics (see [RTSP-specific features](#rtsp-specific-features)).

In order to use the TCP transport protocol, use the `--rtsp_tcp` flag:

```sh
vlc --network-caching=50 --rtsp-tcp rtsp://localhost:8554/mystream
```

In order to use the UDP-multicast transport protocol, append `?vlcmulticast` to the URL:

```sh
vlc --network-caching=50 rtsp://localhost:8554/mystream?vlcmulticast
```

##### Ubuntu bug

The VLC shipped with Ubuntu 21.10 doesn't support playing RTSP due to a license issue (see [here](https://bugs.debian.org/cgi-bin/bugreport.cgi?bug=982299) and [here](https://stackoverflow.com/questions/69766748/cvlc-cannot-play-rtsp-omxplayer-instead-can)). To fix the issue, remove the default VLC instance and install the snap version:

```
sudo apt purge -y vlc
snap install vlc
```

##### Encrypted streams

At the moment VLC doesn't support reading encrypted RTSP streams. However, you can use a proxy like [stunnel](https://www.stunnel.org) or [nginx](https://nginx.org/) or a local _MediaMTX_ instance to decrypt streams before reading them.

#### Web browsers

Web browsers can read a stream from the server in multiple ways (WebRTC or HLS).

You can read a stream by using the [WebRTC protocol](#webrtc-1) by visiting the web page:

```
http://localhost:8889/mystream
```

This web page can be embedded into another web page by using an iframe:

```html
<iframe src="http://mediamtx-ip:8889/mystream" scrolling="no"></iframe>
```

For more advanced setups, you can create and serve a custom web page by starting from the [source code of the read page](internal/servers/webrtc/read_index.html).

Web browsers can also read a stream with the [HLS protocol](#hls). Latency is higher but there are less problems related to connectivity between server and clients, furthermore the server load can be balanced by using a common HTTP CDN (like CloudFront or Cloudflare), and this allows to handle readers in the order of millions. Visit the web page:

```
http://localhost:8888/mystream
```

This web page can be embedded into another web page by using an iframe:

```html
<iframe src="http://mediamtx-ip:8888/mystream" scrolling="no"></iframe>
```

### By protocol

#### SRT

SRT is a protocol that allows to publish and read live data stream, providing encryption, integrity and a retransmission mechanism. It is usually used to transfer media streams encoded with MPEG-TS. In order to read a stream from the server with the SRT protocol, use this URL:

```
srt://localhost:8890?streamid=read:mystream
```

Replace `mystream` with the path name.

If credentials are enabled, append username and password to `streamid`;

```
srt://localhost:8890?streamid=read:mystream:user:pass
```

If you need to use the standard stream ID syntax instead of the custom one in use by this server, see [Standard stream ID syntax](#standard-stream-id-syntax).

Known clients that can read with SRT are [FFmpeg](#ffmpeg-1), [GStreamer](#gstreamer-1) and [VLC](#vlc).

#### WebRTC

WebRTC is an API that makes use of a set of protocols and methods to connect two clients together and allow them to exchange real-time media or data streams. You can read a stream with WebRTC and a web browser by visiting:

```
http://localhost:8889/mystream
```

WHEP is a WebRTC extensions that allows to read streams by using a URL, without passing through a web page. This allows to use WebRTC as a general purpose streaming protocol. If you are using a software that supports WHEP, you can read a stream from the server by using this URL:

```
http://localhost:8889/mystream/whep
```

Depending on the network it may be difficult to establish a connection between server and clients, see [WebRTC-specific features](#webrtc-specific-features) for remediations.

Known clients that can read with WebRTC and WHEP are [FFmpeg](#ffmpeg-1), [GStreamer](#gstreamer-1) and [web browsers](#web-browsers-1).

#### RTSP

RTSP is a protocol that allows to publish and read streams. It supports different underlying transport protocols and allows to encrypt streams in transit (see [RTSP-specific features](#rtsp-specific-features)). In order to read a stream with the RTSP protocol, use this URL:

```
rtsp://localhost:8554/mystream
```

Known clients that can read with RTSP are [FFmpeg](#ffmpeg-1), [GStreamer](#gstreamer-1) and [VLC](#vlc).

##### Latency

The RTSP protocol doesn't introduce any latency by itself. Latency is usually introduced by clients, that put frames in a buffer to compensate network fluctuations. In order to decrease latency, the best way consists in tuning the client. For instance, in VLC, latency can be decreased by decreasing the Network caching parameter, that is available in the "Open network stream" dialog or alternatively can be set with the command line:

```
vlc --network-caching=50 rtsp://...
```

#### RTMP

RTMP is a protocol that allows to read and publish streams, but is less versatile and less efficient than RTSP and WebRTC (doesn't support UDP, doesn't support most RTSP codecs, doesn't support feedback mechanism). Streams can be read from the server by using the URL:

```
rtmp://localhost/mystream
```

In case authentication is enabled, credentials can be passed to the server by using the `user` and `pass` query parameters:

```
rtmp://localhost/mystream?user=myuser&pass=mypass
```

Known clients that can read with RTMP are [FFmpeg](#ffmpeg-1), [GStreamer](#gstreamer-1) and [VLC](#vlc).

#### HLS

HLS is a protocol that works by splitting streams into segments, and by serving these segments and a playlist with the HTTP protocol. You can use _MediaMTX_ to generate a HLS stream, that is accessible through a web page:

```
http://localhost:8888/mystream
```

and can also be accessed without using the browsers, by software that supports the HLS protocol (for instance VLC or _MediaMTX_ itself) by using this URL:

```
http://localhost:8888/mystream/index.m3u8
```

Although the server can produce HLS with a variety of video and audio codecs (that are listed at the beginning of the README), not all browsers can read all codecs.

You can check what codecs your browser can read by [using this tool](https://jsfiddle.net/g1qyf4ea).

If you want to support most browsers, you can to re-encode the stream by using the H264 and AAC codecs, for instance by using FFmpeg:

```sh
ffmpeg -i rtsp://original-source \
-pix_fmt yuv420p -c:v libx264 -preset ultrafast -b:v 600k \
-c:a aac -b:a 160k \
-f rtsp rtsp://localhost:8554/mystream
```

Known clients that can read with HLS are [FFmpeg](#ffmpeg-1), [GStreamer](#gstreamer-1), [VLC](#vlc) and [web browsers](#web-browsers-1).

##### LL-HLS

Low-Latency HLS is a recently standardized variant of the protocol that allows to greatly reduce playback latency. It works by splitting segments into parts, that are served before the segment is complete. LL-HLS is enabled by default. If the stream is not shown correctly, try tuning the hlsPartDuration parameter, for instance:

```yml
hlsPartDuration: 500ms
```

##### Compatibility with Apple devices

In order to correctly display Low-Latency HLS streams in Safari running on Apple devices (iOS or macOS), a TLS certificate is needed and can be generated with OpenSSL:

```sh
openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

Set the `hlsEncryption`, `hlsServerKey` and `hlsServerCert` parameters in the configuration file:

```yml
hlsEncryption: yes
hlsServerKey: server.key
hlsServerCert: server.crt
```

Keep also in mind that not all H264 video streams can be played on Apple Devices due to some intrinsic properties (distance between I-Frames, profile). If the video can't be played correctly, you can either:

* re-encode it by following instructions in this README
* disable the Low-latency variant of HLS and go back to the legacy variant:

  ```yml
    hlsVariant: mpegts
  ```

##### Latency

in HLS, latency is introduced since a client must wait for the server to generate segments before downloading them. This latency amounts to 500ms-3s when the low-latency HLS variant is enabled (and it is by default), otherwise amounts to 1-15secs.

To decrease the latency, you can:

* try decreasing the hlsPartDuration parameter
* try decreasing the hlsSegmentDuration parameter
* The segment duration is influenced by the interval between the IDR frames of the video track. An IDR frame is a frame that can be decoded independently from the others. The server changes the segment duration in order to include at least one IDR frame into each segment. Therefore, you need to decrease the interval between the IDR frames. This can be done in two ways:

  * if the stream is being hardware-generated (i.e. by a camera), there's usually a setting called Key-Frame Interval in the camera configuration page
  * otherwise, the stream must be re-encoded. It's possible to tune the IDR frame interval by using ffmpeg's -g option:

    ```sh
    ffmpeg -i rtsp://original-stream -pix_fmt yuv420p -c:v libx264 -preset ultrafast -b:v 600k -max_muxing_queue_size 1024 -g 30 -f rtsp rtsp://localhost:$RTSP_PORT/compressed
    ```

## Other features

### Configuration

All the configuration parameters are listed and commented in the [configuration file](mediamtx.yml).

There are 3 ways to change the configuration:

1. By editing the `mediamtx.yml` file, that is

   * included into the release bundle
   * available in the root folder of the Docker image (`/mediamtx.yml`); it can be overridden in this way:

     ```
     docker run --rm -it --network=host -v $PWD/mediamtx.yml:/mediamtx.yml bluenviron/mediamtx
     ```

   The configuration can be changed dynamically when the server is running (hot reloading) by writing to the configuration file. Changes are detected and applied without disconnecting existing clients, whenever it's possible.

2. By overriding configuration parameters with environment variables, in the format `MTX_PARAMNAME`, where `PARAMNAME` is the uppercase name of a parameter. For instance, the `rtspAddress` parameter can be overridden in the following way:

   ```
   MTX_RTSPADDRESS="127.0.0.1:8554" ./mediamtx
   ```

   Parameters that have array as value can be overridden by setting a comma-separated list. For example:

   ```
   MTX_PROTOCOLS="tcp,udp"
   ```

   Parameters in maps can be overridden by using underscores, in the following way:

   ```
   MTX_PATHS_TEST_SOURCE=rtsp://myurl ./mediamtx
   ```

   This method is particularly useful when using Docker; any configuration parameter can be changed by passing environment variables with the `-e` flag:

   ```
   docker run --rm -it --network=host -e MTX_PATHS_TEST_SOURCE=rtsp://myurl bluenviron/mediamtx
   ```

3. By using the [Control API](#control-api).

### Authentication

Edit `mediamtx.yml` and set `publishUser` and `publishPass`:

```yml
pathDefaults:
  publishUser: myuser
  publishPass: mypass
```

Only publishers that provide both username and password will be able to proceed:

```
ffmpeg -re -stream_loop -1 -i file.ts -c copy -f rtsp rtsp://myuser:mypass@localhost:8554/mystream
```

It's possible to setup authentication for readers too:

```yml
pathDefaults:
  readUser: myuser
  readPass: mypass
```

If storing plain credentials in the configuration file is a security problem, username and passwords can be stored as hashed strings. The Argon2 and SHA256 hashing algorithms are supported.

To use Argon2, the string must be hashed using Argon2id (recommended) or Argon2i:

```
echo -n "mypass" | argon2 saltItWithSalt -id -l 32 -e
```

Then stored with the `argon2:` prefix:

```yml
pathDefaults:
  readUser: argon2:$argon2id$v=19$m=4096,t=3,p=1$MTIzNDU2Nzg$OGGO0eCMN0ievb4YGSzvS/H+Vajx1pcbUmtLp2tRqRU
  readPass: argon2:$argon2i$v=19$m=4096,t=3,p=1$MTIzNDU2Nzg$oct3kOiFywTdDdt19kT07hdvmsPTvt9zxAUho2DLqZw
```

To use SHA256, the string must be hashed with SHA256 and encoded with base64:

```
echo -n "mypass" | openssl dgst -binary -sha256 | openssl base64
```

Then stored with the `sha256:` prefix:

```yml
pathDefaults:
  readUser: sha256:j1tsRqDEw9xvq/D7/9tMx6Jh/jMhk3UfjwIB2f1zgMo=
  readPass: sha256:BdSWkrdV+ZxFBLUQQY7+7uv9RmiSVA8nrPmjGjJtZQQ=
```

**WARNING**: enable encryption or use a VPN to ensure that no one is intercepting the credentials in transit.

Authentication can be delegated to an external HTTP server:

```yml
externalAuthenticationURL: http://myauthserver/auth
```

Each time a user needs to be authenticated, the specified URL will be requested with the POST method and this payload:

```json
{
  "ip": "ip",
  "user": "user",
  "password": "password",
  "path": "path",
  "protocol": "rtsp|rtmp|hls|webrtc",
  "id": "id",
  "action": "read|publish",
  "query": "query"
}
```

If the URL returns a status code that begins with `20` (i.e. `200`), authentication is successful, otherwise it fails.

Please be aware that it's perfectly normal for the authentication server to receive requests with empty users and passwords, i.e.:

```json
{
  "user": "",
  "password": "",
}
```

This happens because a RTSP client doesn't provide credentials until it is asked to. In order to receive the credentials, the authentication server must reply with status code `401`, then the client will send credentials.

### Encrypt the configuration

The configuration file can be entirely encrypted for security purposes.

An online encryption tool is [available here](https://play.golang.org/p/rX29jwObNe4).

The encryption procedure is the following:

1. NaCL's `crypto_secretbox` function is applied to the content of the configuration. NaCL is a cryptographic library available for [C/C++](https://nacl.cr.yp.to/secretbox.html), [Go](https://pkg.go.dev/golang.org/x/crypto/nacl/secretbox), [C#](https://github.com/somdoron/NaCl.net) and many other languages;

2. The string is prefixed with the nonce;

3. The string is encoded with base64.

After performing the encryption, put the base64-encoded result into the configuration file, and launch the server with the `MTX_CONFKEY` variable:

```
MTX_CONFKEY=mykey ./mediamtx
```

### Remuxing, re-encoding, compression

To change the format, codec or compression of a stream, use _FFmpeg_ or _GStreamer_ together with _MediaMTX_. For instance, to re-encode an existing stream, that is available in the `/original` path, and publish the resulting stream in the `/compressed` path, edit `mediamtx.yml` and replace everything inside section `paths` with the following content:

```yml
paths:
  compressed:
  original:
    runOnReady: >
      ffmpeg -i rtsp://localhost:$RTSP_PORT/$MTX_PATH
        -pix_fmt yuv420p -c:v libx264 -preset ultrafast -b:v 600k
        -max_muxing_queue_size 1024 -f rtsp rtsp://localhost:$RTSP_PORT/compressed
    runOnReadyRestart: yes
```

### Record streams to disk

To save available streams to disk, set the `record` and the `recordPath` parameter in the configuration file:

```yml
pathDefaults:
  # Record streams to disk.
  record: yes
  # Path of recording segments.
  # Extension is added automatically.
  # Available variables are %path (path name), %Y %m %d %H %M %S %f %s (time in strftime format)
  recordPath: ./recordings/%path/%Y-%m-%d_%H-%M-%S-%f
```

All available recording parameters are listed in the [sample configuration file](/mediamtx.yml).

Be aware that not all codecs can be saved with all formats, as described in the compatibility matrix at the beginning of the README.

To upload recordings to a remote location, you can use _MediaMTX_ together with [rclone](https://github.com/rclone/rclone), a command line tool that provides file synchronization capabilities with a huge variety of services (including S3, FTP, SMB, Google Drive):

1. Download and install [rclone](https://github.com/rclone/rclone).

2. Configure _rclone_:

   ```
   rclone config
   ```

3. Place `rclone` into the `runOnInit` and `runOnRecordSegmentComplete` hooks:

   ```yml
   pathDefaults:
     # this is needed to sync segments after a crash.
     # replace myconfig with the name of the rclone config.
     runOnInit: rclone sync -v ./recordings myconfig:/my-path/recordings

     # this is called when a segment has been finalized.
     # replace myconfig with the name of the rclone config.
     runOnRecordSegmentComplete: rclone sync -v --min-age=1ms ./recordings myconfig:/my-path/recordings
   ```

   If you want to delete local segments after they are uploaded, replace `rclone sync` with `rclone move`.

### Playback recorded streams

Existing recordings can be served to users through a dedicated HTTP server, that can be enabled inside the configuration:

```yml
playback: yes
playbackAddress: :9996
```

The server provides an endpoint to list recorded timespans:

```
http://localhost:9996/list?path=[mypath]
```

Where [mypath] is the name of a path. The server will return a list of timespans in JSON format:

```json
[
  {
    "start": "2006-01-02T15:04:05Z07:00",
    "duration": "60.0"
  },
  {
    "start": "2006-01-02T15:07:05Z07:00",
    "duration": "32.33"
  }
]
```

The server provides an endpoint for downloading recordings:

```
http://localhost:9996/get?path=[mypath]&start=[start_date]&duration=[duration]
```

Where:

* [mypath] is the path name
* [start_date] is the start date in RFC3339 format
* [duration] is the maximum duration of the recording in seconds

All parameters must be [url-encoded](https://www.urlencoder.org/). For instance:

```
http://localhost:9996/get?path=stream2&start=2024-01-14T16%3A33%3A17%2B00%3A00&duration=200.5
```

The resulting stream uses the fMP4 format, that is natively compatible with any browser, therefore its URL can be directly inserted into a \<video> tag:

```html
<video controls>
  <source src="http://localhost:9996/get?path=[mypath]&start=[start_date]&duration=[duration]" type="video/mp4" />
</video>
```

### Forward streams to other servers

To forward incoming streams to another server, use _FFmpeg_ inside the `runOnReady` parameter:

```yml
pathDefaults:
  runOnReady: >
    ffmpeg -i rtsp://localhost:$RTSP_PORT/$MTX_PATH
    -c copy
    -f rtsp rtsp://other-server:8554/another-path
  runOnReadyRestart: yes
```

### Proxy requests to other servers

The server allows to proxy incoming requests to other servers or cameras. This is useful to expose servers or cameras behind a NAT. Edit `mediamtx.yml` and replace everything inside section `paths` with the following content:

```yml
paths:
  "~^proxy_(.+)$":
    # If path name is a regular expression, $G1, G2, etc will be replaced
    # with regular expression groups.
    source: rtsp://other-server:8554/$G1
    sourceOnDemand: yes
```

All requests addressed to `rtsp://server:8854/proxy_a` will be forwarded to `rtsp://other-server:8854/a` and so on.

### On-demand publishing

Edit `mediamtx.yml` and replace everything inside section `paths` with the following content:

```yml
paths:
  ondemand:
    runOnDemand: ffmpeg -re -stream_loop -1 -i file.ts -c copy -f rtsp rtsp://localhost:$RTSP_PORT/$MTX_PATH
    runOnDemandRestart: yes
```

The command inserted into `runOnDemand` will start only when a client requests the path `ondemand`, therefore the file will start streaming only when requested.

### Start on boot

#### Linux

On most Linux distributions (including Ubuntu and Debian, but not OpenWrt), _systemd_ is in charge of managing services and starting them on boot.

Move the server executable and configuration in global folders:

```sh
sudo mv mediamtx /usr/local/bin/
sudo mv mediamtx.yml /usr/local/etc/
```

Create a _systemd_ service:

```sh
sudo tee /etc/systemd/system/mediamtx.service >/dev/null << EOF
[Unit]
Wants=network.target
[Service]
ExecStart=/usr/local/bin/mediamtx /usr/local/etc/mediamtx.yml
[Install]
WantedBy=multi-user.target
EOF
```

Enable and start the service:

```sh
sudo systemctl daemon-reload
sudo systemctl enable mediamtx
sudo systemctl start mediamtx
```

#### OpenWrt

Move the server executable and configuration in global folders:

```sh
mv mediamtx /usr/bin/
mkdir -p /usr/etc && mv mediamtx.yml /usr/etc/
```

Create a procd service:

```sh
tee /etc/init.d/mediamtx >/dev/null << EOF
#!/bin/sh /etc/rc.common
USE_PROCD=1
START=95
STOP=01
start_service() {
    procd_open_instance
    procd_set_param command /usr/bin/mediamtx
    procd_set_param stdout 1
    procd_set_param stderr 1
    procd_close_instance
}
EOF
```

Enable and start the service:

```sh
chmod +x /etc/init.d/mediamtx
/etc/init.d/mediamtx enable
/etc/init.d/mediamtx start
```

Read the server logs:

```sh
logread
```

#### Windows

Download the [WinSW v2 executable](https://github.com/winsw/winsw/releases/download/v2.11.0/WinSW-x64.exe) and place it into the same folder of `mediamtx.exe`.

In the same folder, create a file named `WinSW-x64.xml` with this content:

```xml
<service>
  <id>mediamtx</id>
  <name>mediamtx</name>
  <description></description>
  <executable>%BASE%/mediamtx.exe</executable>
</service>
```

Open a terminal, navigate to the folder and run:

```
WinSW-x64 install
```

The server is now installed as a system service and will start at boot time.

### Hooks

The server allows to specify commands that are executed when a certain event happens, allowing the propagation of events to external software.

`runOnConnect` allows to run a command when a client connects to the server:

```yml
# Command to run when a client connects to the server.
# This is terminated with SIGINT when a client disconnects from the server.
# The following environment variables are available:
# * RTSP_PORT: RTSP server port
# * MTX_CONN_TYPE: connection type
# * MTX_CONN_ID: connection ID
runOnConnect: curl http://my-custom-server/webhook?conn_type=$MTX_CONN_TYPE&conn_id=$MTX_CONN_ID
# Restart the command if it exits.
runOnConnectRestart: no
```

`runOnDisconnect` allows to run a command when a client disconnects from the server:

```yml
# Command to run when a client disconnects from the server.
# Environment variables are the same of runOnConnect.
runOnDisconnect: curl http://my-custom-server/webhook?conn_type=$MTX_CONN_TYPE&conn_id=$MTX_CONN_ID
```

`runOnInit` allows to run a command when a path is initialized. This can be used to publish a stream when the server is launched:

```yml
paths:
  mypath:
    # Command to run when this path is initialized.
    # This can be used to publish a stream when the server is launched.
    # The following environment variables are available:
    # * MTX_PATH: path name
    # * RTSP_PORT: RTSP server port
    # * G1, G2, ...: regular expression groups, if path name is
    #   a regular expression.
    runOnInit: ffmpeg -i my_file.mp4 -c copy -f rtsp rtsp://localhost:8554/mypath
    # Restart the command if it exits.
    runOnInitRestart: no
```

`runOnDemand` allows to run a command when a path is requested by a reader. This can be used to publish a stream on demand:

```yml
pathDefaults:
  # Command to run when this path is requested by a reader
  # and no one is publishing to this path yet.
  # This is terminated with SIGINT when there are no readers anymore.
  # The following environment variables are available:
  # * MTX_PATH: path name
  # * MTX_QUERY: query parameters (passed by first reader)
  # * RTSP_PORT: RTSP server port
  # * G1, G2, ...: regular expression groups, if path name is
  #   a regular expression.
  runOnDemand: ffmpeg -i my_file.mp4 -c copy -f rtsp rtsp://localhost:8554/mypath
  # Restart the command if it exits.
  runOnDemandRestart: no
```

`runOnUnDemand` allows to run a command when there are no readers anymore:

```yml
pathDefaults:
  # Command to run when there are no readers anymore.
  # Environment variables are the same of runOnDemand.
  runOnUnDemand:
```

`runOnReady` allows to run a command when a stream is ready to be read:

```yml
pathDefaults:
  # Command to run when the stream is ready to be read, whenever it is
  # published by a client or pulled from a server / camera.
  # This is terminated with SIGINT when the stream is not ready anymore.
  # The following environment variables are available:
  # * MTX_PATH: path name
  # * MTX_QUERY: query parameters (passed by publisher)
  # * MTX_SOURCE_TYPE: source type
  # * MTX_SOURCE_ID: source ID
  # * RTSP_PORT: RTSP server port
  # * G1, G2, ...: regular expression groups, if path name is
  #   a regular expression.
  runOnReady: curl http://my-custom-server/webhook?path=$MTX_PATH&source_type=$MTX_SOURCE_TYPE&source_id=$MTX_SOURCE_ID
  # Restart the command if it exits.
  runOnReadyRestart: no
```

`runOnNotReady` allows to run a command when a stream is not available anymore:

```yml
pathDefaults:
  # Command to run when the stream is not available anymore.
  # Environment variables are the same of runOnReady.
  runOnNotReady: curl http://my-custom-server/webhook?path=$MTX_PATH&source_type=$MTX_SOURCE_TYPE&source_id=$MTX_SOURCE_ID
```

`runOnRead` allows to run a command when a client starts reading:

```yml
pathDefaults:
  # Command to run when a client starts reading.
  # This is terminated with SIGINT when a client stops reading.
  # The following environment variables are available:
  # * MTX_PATH: path name
  # * MTX_QUERY: query parameters (passed by reader)
  # * MTX_READER_TYPE: reader type
  # * MTX_READER_ID: reader ID
  # * RTSP_PORT: RTSP server port
  # * G1, G2, ...: regular expression groups, if path name is
  #   a regular expression.
  runOnRead: curl http://my-custom-server/webhook?path=$MTX_PATH&reader_type=$MTX_READER_TYPE&reader_id=$MTX_READER_ID
  # Restart the command if it exits.
  runOnReadRestart: no
```

`runOnUnread` allows to run a command when a client stops reading:

```yml
pathDefaults:
  # Command to run when a client stops reading.
  # Environment variables are the same of runOnRead.
  runOnUnread: curl http://my-custom-server/webhook?path=$MTX_PATH&reader_type=$MTX_READER_TYPE&reader_id=$MTX_READER_ID
```

`runOnRecordSegmentCreate` allows to run a command when a recording segment is created:

```yml
pathDefaults:
  # Command to run when a recording segment is created.
  # The following environment variables are available:
  # * MTX_PATH: path name
  # * RTSP_PORT: RTSP server port
  # * G1, G2, ...: regular expression groups, if path name is
  #   a regular expression.
  # * MTX_SEGMENT_PATH: segment file path
  runOnRecordSegmentCreate: curl http://my-custom-server/webhook?path=$MTX_PATH&segment_path=$MTX_SEGMENT_PATH
```

`runOnRecordSegmentComplete` allows to run a command when a recording segment is complete:

```yml
pathDefaults:
  # Command to run when a recording segment is complete.
  # The following environment variables are available:
  # * MTX_PATH: path name
  # * RTSP_PORT: RTSP server port
  # * G1, G2, ...: regular expression groups, if path name is
  #   a regular expression.
  # * MTX_SEGMENT_PATH: segment file path
  runOnRecordSegmentComplete: curl http://my-custom-server/webhook?path=$MTX_PATH&segment_path=$MTX_SEGMENT_PATH
```

### Control API

The server can be queried and controlled with an API, that must be enabled by setting the `api` parameter in the configuration:

```yml
api: yes
```

The API listens on `apiAddress`, that by default is `127.0.0.1:9997`; for instance, to obtain a list of active paths, run:

```
curl http://127.0.0.1:9997/v2/paths/list
```

Full documentation of the Control API is available on the [dedicated site](https://bluenviron.github.io/mediamtx/).

### Metrics

A metrics exporter, compatible with [Prometheus](https://prometheus.io/), can be enabled with the parameter `metrics: yes`; then the server can be queried for metrics with Prometheus or with a simple HTTP request:

```
curl localhost:9998/metrics
```

Obtaining:

```ini
# metrics of every path
paths{name="[path_name]",state="[state]"} 1
paths_bytes_received{name="[path_name]",state="[state]"} 1234
paths_bytes_sent{name="[path_name]",state="[state]"} 1234

# metrics of every HLS muxer
hls_muxers{name="[name]"} 1
hls_muxers_bytes_sent{name="[name]"} 187

# metrics of every RTSP connection
rtsp_conns{id="[id]"} 1
rtsp_conns_bytes_received{id="[id]"} 1234
rtsp_conns_bytes_sent{id="[id]"} 187

# metrics of every RTSP session
rtsp_sessions{id="[id]",state="idle"} 1
rtsp_sessions_bytes_received{id="[id]",state="[state]"} 1234
rtsp_sessions_bytes_sent{id="[id]",state="[state]"} 187

# metrics of every RTSPS connection
rtsps_conns{id="[id]"} 1
rtsps_conns_bytes_received{id="[id]"} 1234
rtsps_conns_bytes_sent{id="[id]"} 187

# metrics of every RTSPS session
rtsps_sessions{id="[id]",state="[state]"} 1
rtsps_sessions_bytes_received{id="[id]",state="[state]"} 1234
rtsps_sessions_bytes_sent{id="[id]",state="[state]"} 187

# metrics of every RTMP connection
rtmp_conns{id="[id]",state="[state]"} 1
rtmp_conns_bytes_received{id="[id]",state="[state]"} 1234
rtmp_conns_bytes_sent{id="[id]",state="[state]"} 187

# metrics of every RTMPS connection
rtmps_conns{id="[id]",state="[state]"} 1
rtmps_conns_bytes_received{id="[id]",state="[state]"} 1234
rtmps_conns_bytes_sent{id="[id]",state="[state]"} 187

# metrics of every SRT connection
srt_conns{id="[id]",state="[state]"} 1
srt_conns_bytes_received{id="[id]",state="[state]"} 1234
srt_conns_bytes_sent{id="[id]",state="[state]"} 187

# metrics of every WebRTC session
webrtc_sessions{id="[id]",state="[state]"} 1
webrtc_sessions_bytes_received{id="[id]",state="[state]"} 1234
webrtc_sessions_bytes_sent{id="[id]",state="[state]"} 187
```

### pprof

A performance monitor, compatible with pprof, can be enabled with the parameter `pprof: yes`; then the server can be queried for metrics with pprof-compatible tools, like:

```
go tool pprof -text http://localhost:9999/debug/pprof/goroutine
go tool pprof -text http://localhost:9999/debug/pprof/heap
go tool pprof -text http://localhost:9999/debug/pprof/profile?seconds=30
```

### SRT-specific features

#### Standard stream ID syntax

In SRT, the stream ID is a string that is sent to the counterpart in order to advertise what action the caller is gonna do (publish or read), the path and the credentials. All these informations have to be encoded into a single string. This server supports two stream ID syntaxes, a custom one (that is the one reported in rest of the README) and also a [standard one](https://github.com/Haivision/srt/blob/master/docs/features/access-control.md) proposed by the authors of the protocol and sometimes enforced by some hardware. The standard syntax can be used in this way:

```
srt://localhost:8890?streamid=#!::m=publish,r=mypath,u=myuser,s=mypass&pkt_size=1316
```

Where:

* key `m` contains the action (`publish` or `request`)
* key `r` contains the path
* key `u` contains the username
* key `s` contains the password

### WebRTC-specific features

#### Connectivity issues

If the server is hosted inside a container or is behind a NAT, additional configuration is required in order to allow the two WebRTC parts (server and client) to establish a connection.

Make sure that `webrtcAdditionalHosts` includes your public IPs, that are IPs that can be used by clients to reach the server. If clients are on the same LAN as the server, then insert the LAN address of the server. If clients are coming from the internet, insert the public IP address of the server, or alternatively a DNS name, if you have one. You can insert multiple values to support all scenarios:

```yml
webrtcAdditionalHosts: [192.168.x.x, 1.2.3.4, my-dns.example.org, ...]
```

If there's a NAT / container between server and clients, it must be configured to route all incoming UDP packets on port 8189 to the server. If you're using Docker, this can be achieved with the flag:

```sh
docker run --rm -it \
-p 8189:8189/udp
....
bluenviron/mediamtx
```

If you still have problems, maybe the UDP protocol is blocked by a firewall. Enable the local TCP listener:

```yml
# any port of choice
webrtcLocalTCPAddress: :8189
```

If there's a NAT / container between server and clients, it must be configured to route all incoming TCP packets on port 8189 to the server.

If you still have problems, enable a STUN server:

```yml
# STUN servers allows to obtain and share the public IP of the server.
webrtcICEServers2:
  - url: stun:stun.l.google.com:19302
```

If you really still have problems, you can force all WebRTC/ICE connections to pass through a TURN server, like coturn, that must be configured externally. The server address and credentials must be set in the configuration file:

```yml
# TURN/TURNS servers forces all traffic through them.
webrtcICEServers2:
- url: turn:host:port
  username: user
  password: password
```

Where user and pass are the username and password of the server. Note that port is not optional.

If the server uses a secret-based authentication (for instance, coturn with the use-auth-secret option), it must be configured by using AUTH_SECRET as username, and the secret as password:

```yml
webrtcICEServers2:
- url: turn:host:port
  username: AUTH_SECRET
  password: secret
```

where secret is the secret of the TURN server. MediaMTX will generate a set of credentials by using the secret, and credentials will be sent to clients before the WebRTC/ICE connection is established.

### RTSP-specific features

#### Transport protocols

The RTSP protocol supports different underlying transport protocols, that are chosen by clients during the handshake with the server:

* UDP: the most performant, but doesn't work when there's a NAT/firewall between server and clients. It doesn't support encryption.
* UDP-multicast: allows to save bandwidth when clients are all in the same LAN, by sending packets once to a fixed multicast IP. It doesn't support encryption.
* TCP: the most versatile, does support encryption.

The default transport protocol is UDP. To change the transport protocol, you have to tune the configuration of your client of choice.

#### Encryption

Incoming and outgoing RTSP streams can be encrypted with TLS, obtaining the RTSPS protocol. A TLS certificate is needed and can be generated with OpenSSL:

```sh
openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

Edit `mediamtx.yml`, and set the `protocols`, `encryption`, `serverKey` and serverCert parameters:

```yml
protocols: [tcp]
encryption: optional
serverKey: server.key
serverCert: server.crt
```

Streams can be published and read with the `rtsps` scheme and the `8322` port:

```
rtsps://localhost:8322/mystream
```

#### Corrupted frames

In some scenarios, when publishing or reading from the server with RTSP, frames can get corrupted. This can be caused by multiple reasons:

* the write queue of the server is too small and can't keep up with the stream throughput. A solution consists in increasing its size:

  ```yml
  writeQueueSize: 1024
  ```

* The stream throughput is too big and the stream can't be transmitted correctly with the UDP transport protocol. UDP is more performant, faster and more efficient than TCP, but doesn't have a retransmission mechanism, that is needed in case of streams that need a large bandwidth. A solution consists in switching to TCP:

  ```yml
  protocols: [tcp]
  ```

  In case the source is a camera:

  ```yml
  paths:
    test:
      source: rtsp://..
      rtspTransport: tcp
   ```

* The stream throughput is too big to be handled by the network between server and readers. Upgrade the network or decrease the stream bitrate by re-encoding it.

### RTMP-specific features

#### Encryption

RTMP connections can be encrypted with TLS, obtaining the RTMPS protocol. A TLS certificate is needed and can be generated with OpenSSL:

```yml
openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

Edit mediamtx.yml, and set the `rtmpEncryption`, `rtmpServerKey` and `rtmpServerCert` parameters:

```yml
rtmpEncryption: optional
rtmpServerKey: server.key
rtmpServerCert: server.crt
```

Streams can be published and read with the rtmps scheme and the 1937 port:

```
rtmps://localhost:1937/...
```

Be aware that RTMPS is currently unsupported by all major players. However, you can use a proxy like [stunnel](https://www.stunnel.org) or [nginx](https://nginx.org/) or a dedicated _MediaMTX_ instance to decrypt streams before reading them.

## Compile from source

### Standard

Install git and Go &ge; 1.21. Clone the repository, enter into the folder and start the building process:

```sh
git clone https://github.com/bluenviron/mediamtx
cd mediamtx
CGO_ENABLED=0 go build .
```

The command will produce the `mediamtx` binary.

### Raspberry Pi

The server can be compiled with native support for the Raspberry Pi Camera. Compilation must be performed on a Raspberry Pi, with the following dependencies:

* Go &ge; 1.21
* `libcamera-dev`
* `libfreetype-dev`
* `xxd`

Download the repository, open a terminal in it and run:

```sh
cd internal/protocols/rpicamera/exe
make
cd ../../../../
go build -tags rpicamera .
```

The command will produce the `mediamtx` binary.

### OpenWrt

The compilation procedure is the same as the standard one. On the OpenWrt device, install git and Go:

```sh
opkg update
opkg install golang git git-http
```

Clone the repository, enter into the folder and start the building process:

```sh
git clone https://github.com/bluenviron/mediamtx
cd mediamtx
CGO_ENABLED=0 go build .
```

The command will produce the `mediamtx` binary.

If the OpenWrt device doesn't have enough resources to compile, you can [cross compile](#cross-compile) from another machine.

### Cross compile

Cross compilation allows to build an executable for a target machine from another machine with different operating system or architecture. This is useful in case the target machine doesn't have enough resources for compilation or if you don't want to install the compilation dependencies on it.

On the machine you want to use to compile, install git and Go &ge; 1.21. Clone the repository, enter into the folder and start the building process:

```sh
git clone https://github.com/bluenviron/mediamtx
cd mediamtx
CGO_ENABLED=0 GOOS=my_os GOARCH=my_arch go build .
```

Replace `my_os` and `my_arch` with the operating system and architecture of your target machine. A list of all supported combinations can be obtained with:

```sh
go tool dist list
```

For instance:

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build .
```

In case of the `arm` architecture, there's an additional flag available, `GOARM`, that allows to set the ARM version:

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GOARM=7 go build .
```

In case of the `mips` architecture, there's an additional flag available, `GOMIPS`, that allows to set additional parameters:

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=mips GOMIPS=softfloat go build .
```

The command will produce the `mediamtx` binary.

### Compile for all supported platforms

Install Docker and launch:

```sh
make binaries
```

The command will produce tarballs in folder `binaries/`.

## Specifications

|name|area|
|----|----|
|[RTSP / RTP / RTCP specifications](https://github.com/bluenviron/gortsplib#specifications)|RTSP|
|[HLS specifications](https://github.com/bluenviron/gohlslib#specifications)|HLS|
|[RTMP](https://rtmp.veriskope.com/pdf/rtmp_specification_1.0.pdf)|RTMP|
|[Enhanced RTMP](https://raw.githubusercontent.com/veovera/enhanced-rtmp/main/enhanced-rtmp-v1.pdf)|RTMP|
|[WebRTC: Real-Time Communication in Browsers](https://www.w3.org/TR/webrtc/)|WebRTC|
|[WebRTC HTTP Ingestion Protocol (WHIP)](https://datatracker.ietf.org/doc/draft-ietf-wish-whip/)|WebRTC|
|[WebRTC HTTP Egress Protocol (WHEP)](https://datatracker.ietf.org/doc/draft-murillo-whep/)|WebRTC|
|[The SRT Protocol](https://haivision.github.io/srt-rfc/draft-sharabayko-srt.html)|SRT|
|[Codec specifications](https://github.com/bluenviron/mediacommon#specifications)|codecs|
|[Golang project layout](https://github.com/golang-standards/project-layout)|project layout|

## Related projects

* [gortsplib (RTSP library used internally)](https://github.com/bluenviron/gortsplib)
* [gohlslib (HLS library used internally)](https://github.com/bluenviron/gohlslib)
* [mediacommon (codecs and formats library used internally)](https://github.com/bluenviron/mediacommon)
* [datarhei/gosrt (SRT library used internally)](https://github.com/datarhei/gosrt)
* [pion/webrtc (WebRTC library used internally)](https://github.com/pion/webrtc)
* [pion/sdp (SDP library used internally)](https://github.com/pion/sdp)
* [pion/rtp (RTP library used internally)](https://github.com/pion/rtp)
* [pion/rtcp (RTCP library used internally)](https://github.com/pion/rtcp)
* [notedit/rtmp (RTMP library used internally)](https://github.com/notedit/rtmp)
* [go-astits (MPEG-TS library used internally)](https://github.com/asticode/go-astits)
* [go-mp4 (MP4 library used internally)](https://github.com/abema/go-mp4)
* [hls.js (browser-side HLS library used internally)](https://github.com/video-dev/hls.js)
