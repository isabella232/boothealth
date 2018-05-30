Bot for tracking health of the bootnodes
=========================================

How to run?
-------------

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