# [Dendrite](https://github.com/matrix-org/dendrite)
Dendrite is Matrix homeserver written in Go.

It is a rewrite of [synapse](https://github.com/matrix-org/synapse) that was written in Python which could not scale.

## Notes
- You don't need to run Kafka unless working on distributed stuff.
- Dendrite is composed of several microservices (room server, sync api, media api...).
	- Kafka is used to transfer messages between those microservices.

## Links
- [Dendrite design](https://github.com/matrix-org/dendrite/blob/master/DESIGN.md)
- [Dendrite checklist](https://docs.google.com/spreadsheets/d/1tkMNpIpPjvuDJWjPFbw_xzNzOHBA-Hp50Rkpcr43xTw) - API to finish for release.