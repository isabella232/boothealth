Bot for tracking health of the bootnodes
=========================================

The only mandatory parameter is nursery nodes. For testing cluster we use:

1. enode://1b843c7697f6fc42a1f606fb3cfaac54e025f06789dc20ad9278be3388967cf21e3a1b1e4be51faecd66c2c3adef12e942b4fcdeb8727657abe60636efb6224f@206.189.6.46:30404
2. enode://b29100c8468e3e6604817174a15e4d71627458b0dcdbeea169ab2eb4ab2bbc6f24adbb175826726cec69db8fdba6c0dd60b3da598e530ede562180d300728659@206.189.6.48:30404


Image can be built with:

```bash
make image
```

Or downloaded from `statusteam/boothleath` repository.

To execute health check run (also it can be run as a daemon):

```
docker run -ti statusteam/boothealth:latest -n enode://1b843c7697f6fc42a1f606fb3cfaac54e025f06789dc20ad9278be3388967cf21e3a1b1e4be51faecd66c2c3adef12e942b4fcdeb8727657abe60636efb6224f@206.189.6.46:30404 -n enode://b29100c8468e3e6604817174a15e4d71627458b0dcdbeea169ab2eb4ab2bbc6f24adbb175826726cec69db8fdba6c0dd60b3da598e530ede562180d300728659@206.189.6.48:30404
```

Latency for finding a peer can be viewed in logs:

```
INFO [05-30|09:12:49] Started search.                          topic=whisper limit=5
INFO [05-30|09:12:53] Discovered node                          total=1 node=enode://8457590884c77f46af7f7b48e66b30d931c7e9c1d9d399e58551a4e6c7d881f4de97ca570da139a5fb4e5761f7b8badd5409136cd3dfd515a3f90e1be68d04e1@206.189.6.48:30304   latency=3.826192682s
INFO [05-30|09:12:53] Discovered node                          total=2 node=enode://f54f588dc1a307c44fe05dd031b0bea340b2ee75c967e9ba87f48ee0341ca854717808f96bb467ff25745b475c0c7cc87bd0a23af879fdf2da3eb284a21c42b9@206.189.6.48:30305   latency=4.114527495s
INFO [05-30|09:12:53] Discovered node                          total=3 node=enode://878432d37c737d4299628fcdb0212a6f9b1bc3e872f8b4eea076e36ddcba7f613b59794330eebc418f026b0830c8fb3b19ffdfa2ff910d48817727eae3c9dea2@206.189.50.97:30305  latency=4.350209285s
INFO [05-30|09:12:54] Discovered node                          total=4 node=enode://03c3bfb149750d85b05882e51ea071ad93068023ef76b92dc8760b2963c92c52c1a249584cb457f79c5c10d68e89372bdda2342c22e5818ebb1ba6b264c57ed0@206.189.56.154:30304 latency=4.861431408s
```

Or by using metrics exposed by default on 8080 port.

```
curl 0.0.0.0:8080/metrics | grep peers_discovery
```

```
peers_discovery_latency_bucket{peers="1",le="0.5"} 0
peers_discovery_latency_bucket{peers="1",le="2"} 0
peers_discovery_latency_bucket{peers="1",le="5"} 4
peers_discovery_latency_bucket{peers="1",le="10"} 4
peers_discovery_latency_bucket{peers="1",le="20"} 4
peers_discovery_latency_bucket{peers="1",le="30"} 4
peers_discovery_latency_bucket{peers="1",le="+Inf"} 4
peers_discovery_latency_sum{peers="1"} 15.150799482
peers_discovery_latency_count{peers="1"} 4
peers_discovery_latency_bucket{peers="2",le="0.5"} 0
peers_discovery_latency_bucket{peers="2",le="2"} 0
peers_discovery_latency_bucket{peers="2",le="5"} 3
peers_discovery_latency_bucket{peers="2",le="10"} 4
peers_discovery_latency_bucket{peers="2",le="20"} 4
peers_discovery_latency_bucket{peers="2",le="30"} 4
peers_discovery_latency_bucket{peers="2",le="+Inf"} 4
peers_discovery_latency_sum{peers="2"} 18.084040674
peers_discovery_latency_count{peers="2"} 4
```