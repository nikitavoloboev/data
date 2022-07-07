---
title: LoRaWAN
---

# [LoRaWAN](https://www.thethingsnetwork.org/docs/lorawan/)

[The Things Network](https://www.thethingsnetwork.org/) is nice project together with their [Things Stack server](https://github.com/TheThingsNetwork/lorawan-stack).

## Notes

- [LoraWAN uses Lora radios (all made or licensed by Semtech,) which only implement the physical layer. LoraWAN is a MAC + network software layer. It's similar to 802.15.4 and Zigbee. Thus, if the hardware integrates Lora devices (such as the SX1276), then it can be used to implement LoraWAN or custom protocols using Lora.](https://www.reddit.com/r/raspberry_pi/comments/9mfrym/iot_lora_boards_launch_with_raspberry_pi_microbit/)
- [LoRa is the RF modulation, LoRaWAN is the protocol on top, that is defined and adopted by the LoRa alliance.](https://www.reddit.com/r/IOT/comments/a339h3/im_using_lora_and_not_lorawan_what_am_i_missing/)
- [LoRa is a spread spectrum technology that also has a problem with collisions. Spread spectrum is where the signal bandwidth of your transmission is higher than the information bandwidth, meaning your transmission's spectral efficiency is terrible. But usually spread spectrum systems are designed such that multiple transmissions can occur at the same time without colliding, so the network spectral efficiency is just as good as a narrowband system. But LoRa is a spread spectrum technology that still suffers from collisions, so the network spectral efficiency is really bad.](https://news.ycombinator.com/item?id=21073841)
- [LoRa is like a walkie-talkie. It just broadcasts packets. It's slow (think 300 baud), but long-range (10+ miles). LoRaWAN is like a cellular protocol built on top of LoRa. It's a standard, but it doesn't work by magic, you have to subscribe and pay someone who owns the base stations (called gateways). Your packets don't "directly" get on the internet: your packets are sent only to your Application Gateway on the internet, then you can do anything from there. (Yes, that adds latency.)](https://www.reddit.com/r/Lora/comments/mqg1xn/must_educate_myself/)

## Links

- [Semtech LoRa Technology Overview](http://www.semtech.com/lora) ([GitHub](https://github.com/Lora-net))
- [TTN wiki](https://www.thethingsnetwork.org/docs/) ([Code](https://github.com/TheThingsNetwork/docs))
- [Teaching a new dog old tricks](https://medium.com/@aallan/teaching-a-new-dog-old-tricks-965a578e753b)
- [Nicolas Sornin - Future of LoRa](https://www.youtube.com/watch?v=jNnPTxWRNxs)
- [Introduction to LoRa](https://www.youtube.com/watch?v=8Oxcp9wQQnk)
- [LoRa crash course by Thomas Telkamp](https://www.youtube.com/watch?v=T3dGLqZrjIQ)
- [11 Myths About LoRaWAN](https://www.electronicdesign.com/industrial-automation/11-myths-about-lorawan)
- [LoRaWAN Encryption 101](https://lorawanacademyblog.semtech.com/lorawan-encryption-101)
- [LoRaWAN Academy](https://lorawanacademy.semtech.com/)
- [Building LPWAN Solutions that last for Years (2018)](https://www.youtube.com/watch?v=mEmpbIlo8XQ&list=PLEx5khR4g7PJW7u0GKxRPIQddtu69boT3)
- [Example LoRaWAN application for Mbed-OS](https://github.com/ARMmbed/mbed-os-example-lorawan)
- [Firmware-updates enabled LoRaWAN example application](https://github.com/ARMmbed/mbed-os-example-lorawan-fuota)
- [QMesh](https://github.com/faydr/QMesh) - Synchronized flooded mesh network for voice over LoRa.
- [Things Network Stack for LoRaWAN Frequency Plans](https://github.com/TheThingsNetwork/lorawan-frequency-plans)
- [The Problem with LoRa (2018)](https://hackernoon.com/the-problem-with-lora-af4f5263d378)
- [The Things Network Stack for LoRaWAN](https://github.com/TheThingsNetwork/lorawan-stack) - Open-source LoRaWAN network stack suitable for large, global and geo-distributed public and private networks as well as smaller networks.
- [TTN LoRaWAN v3 stack explanation](https://www.youtube.com/watch?v=CeSvqkxg25c)
- [Driving down the total cost of ownership of LoRaWAN deployments (2019)](https://www.youtube.com/watch?v=HJFkV8qqhr4)
- [Your Primer for LoRa/LoRaWAN (2018)](https://medium.com/iotforall/your-primer-for-lora-lorawan-33f1e0eb4215)
- [Interview with Nicolas Sornin (Semtech) - The Things Conference 2019](https://www.youtube.com/watch?v=MCIzkiMloyE)
- [Interview with Johan Stokking (The Things Network) - The Things Conference 2019](https://www.youtube.com/watch?v=hGf93zY8X1c)
- [Interview with Alistair Fulton (Semtech) - The Things Conference 2019](https://www.youtube.com/watch?v=-tzKMHTdoe8)
- [Extending LoRaWAN's reach - Nicolas Sornin (Semtech) - The Things Conference 2019](https://www.youtube.com/watch?v=pHq7_rgDyFA)
- [From zero to LoRaWAN in a weekend (2018)](https://github.com/ttn-zh/ic880a-gateway/wiki)
- [LoRaWAN and the Ocean Cleanup - Jasper den Hartog (MCS) (2019)](https://www.youtube.com/watch?v=E5huiCbVR5A)
- [Lacuna Space](http://lacuna.space/) - Low-cost, simple and reliable global connections to sensors and mobile equipment.
- [Can someone give me the rundown on LoRa? (2019)](https://www.reddit.com/r/IOT/comments/bzd2q0/can_someone_give_me_the_rundown_on_lora/)
- [LoRaWAN Primer](https://docs.exploratory.engineering/lora/lorawan/)
- [LoRaWAN® distance world record broken, twice. 766 km (476 miles) using 25mW transmission power (2019)](https://www.thethingsnetwork.org/article/lorawan-distance-world-record) ([HN](https://news.ycombinator.com/item?id=20562684))
- [Feather TFT LoRa Sniffer](https://github.com/ImprobableStudios/Feather_TFT_LoRa_Sniffer) - Scans LoRa channels displaying any monitored packets on the TFT display.
- [Compact server for private LoRaWAN networks](https://github.com/gotthardp/lorawan-server)
- [Introduction to LoRaWAN and The Things Network (2019)](https://channel9.msdn.com/Shows/Internet-of-Things-Show/Introduction-to-LoRaWAN-and-The-Things-Network)
- [Technical overview of LoRa and LoRaWAN](https://lora-alliance.org/sites/default/files/2018-04/what-is-lorawan.pdf)
- [The Things Stack](https://thethingsstack.io) - Open Source LoRaWAN Network Server. ([Code](https://github.com/TheThingsIndustries/lorawan-stack-docs)) ([Docs](https://www.thethingsindustries.com/docs/))
- [Securing LoRaWAN with Secure Elements (2019)](https://www.linkedin.com/pulse/securing-lorawan-secure-elements-johan-stokking/)
- [The Things Conference 2020](https://www.youtube.com/watch?v=0eOpMDffbQ0)
- [The LoraWan Pager](https://hackaday.io/project/22038-the-lorawan-pager) - Open telecommunicator device, independent from SIM CARDS and Telecom providers.
- [LoRa-based Device-to-Device Smartphone Communication for Crisis Scenarios (2020)](https://dtn7.github.io/assets/hoechst2020lora.pdf) ([HN](https://news.ycombinator.com/item?id=22725623))
- [Adding machine learning to your LoRaWAN device - Jan Jongboom (2020)](https://www.youtube.com/watch?v=e-v0wnSM6YA)
- [LoRa Mesh Communication without Infrastructure: The Meshtastic Project (ESP32, BLE, GPS) (2020)](https://www.youtube.com/watch?v=TY6m6fS8bxU) ([Reddit](https://www.reddit.com/r/darknetplan/comments/gyalhx/lora_mesh_communication_without_infrastructure/))
- [TTGO LoRa32 development board (2020)](https://www.settorezero.com/wordpress/lora_lorawan_lilygo_ttgo_lora32_esp32/) ([Reddit](https://www.reddit.com/r/esp32/comments/hu32z1/in_this_article_i_explain_how_lora_and_lorawan/))
- [DecodingLora](https://revspace.nl/DecodingLora)
- [We build a \$20 LoRa Satellite Ground Station and we follow the FossaSat-1 launch (2019)](https://www.youtube.com/watch?v=5k0aM-PJzo8) ([HN](https://news.ycombinator.com/item?id=24519340))
- [Arduino Library for The Things Node](https://github.com/TheThingsNetwork/arduino-node-lib)
- [LoRa packet forwarder](https://github.com/Lora-net/packet_forwarder) - Program running on the host of a Lora gateway that forwards RF packets receive by the concentrator to a server through a IP/UDP link, and emits RF packets that are sent by the server.
- [ChirpStack](https://www.chirpstack.io/) - Open-source LoRaWAN Network Server stack. ([Code](https://github.com/brocaar/chirpstack-network-server))
- [Experiments with LoRa radios](https://ds0.me/lora/index.html)
- [Connecting Your LoRaWAN Devices from The Things Stack to AWS IoT Core (2020)](https://aws.amazon.com/blogs/apn/connecting-your-lorawan-devices-from-the-things-stack-to-aws-iot-core/)
- [LORIOT](https://www.loriot.io/) - LoRaWAN Network Server Provider.
- [CellSol](https://github.com/RbtsEvrwhr-Riley/CellSol/) - Long Range Solar Powered Communications Network. LoRa based mesh. ([Docs](https://www.f3.to/cellsol/))
- [BastWAN](https://github.com/ElectronicCats/BastWAN) - Best in the world format Feather and LoRa with a RAK4260 and LoRaWAN.
- [Schema for modeling LoRaWAN devices](https://github.com/lorawan-schema/draft-devices)
- [TTN Mapper](https://ttnmapper.org/) ([Code](https://github.com/ttnmapper/ttnmapper-web))
- [LoFence](https://github.com/kiu/lofence) - LoRaWAN capable IoT electric fence monitoring system running on the The Things Network.
- [Rust LoRaWAN](https://github.com/ivajloip/rust-lorawan) - Fast, zero-copy and lightweight LoRaWAN parsing library in Rust.
- [LoRaWAN (Go)](https://github.com/brocaar/lorawan) - Provides structures and tools to read and write LoraWAN messages from and to a slice of bytes.
- [PaperiNode](https://github.com/RobPo/PaperiNode) - Self powered E-Paper Node for LoRaWAN.
- [Build and configure your own LoraWAN Gateway with The Things Network](https://github.com/IRNAS/ttn-irnas-gw)
- [LoRaWAN computer vision with Edge Impulse & Portenta H7](https://github.com/edgeimpulse/example-portenta-lorawan)
- [LoRa Basics Station](https://github.com/balenalabs/basicstation) - Deploys a LoRaWAN gateway with Basics Station Packet Forward protocol with balena.
- [lorawan-node-simulator](https://github.com/kartben/lorawan-node-simulator) - Simulates LoRaWAN gateways, and endpoints regularly emitting LoRaWAN radio packets.
- [Deploying a LoRaWAN network server on Azure (2021)](https://blog.benjamin-cabe.com/2020/12/01/deploying-a-lorawan-network-server-on-azure)
- [Eveything you need to know about LoRaWAN in 60 minutes - Johan Stokking (2021)](https://www.youtube.com/watch?v=ZsVhYiX4_6o)
- [Economical LoRa nodes (2021)](https://www.reddit.com/r/Lora/comments/lenh3s/looking_for_economical_lora_node/)
- [LoRaWAN library for compatible Arduino board](https://github.com/BeelanMX/Beelan-LoRaWAN)
- [Raspberry Pi Pico: The R2040 meets LoRaWAN (2021)](https://lemariva.com/blog/2021/02/raspberry-pi-pico-rp2040-meets-lorawan)
- [The Things Stack Introduction - Bogdans Afonins (2021)](https://www.youtube.com/watch?v=rK8oJHZ9Q7U)
- [IP2LoRa](https://github.com/airbus-cyber/IP2LoRa) - IP tunnelling over LoRa.
- [Adafruit Tiny LoRa](https://github.com/adafruit/Adafruit_CircuitPython_TinyLoRa) - LoRaWAN/The Things Network, for CircuitPython.
- [LoRa-concentrator](https://github.com/will127534/LoRa-concentrator) - Simple Board for SX1301 and SX125X LoRa Gateway / Concentrator.
- [virtual-lorawan-device](https://github.com/helium/virtual-lorawan-device) - Utility that attaches to a Semtech UDP Host and pretends to be a LoRaWAN Device.
- [Senet and Helium Announce LoRaWAN Network Integration Partnership (2021)](https://www.senetco.com/about/news/senet-and-helium-announce-lorawan-network-integration-partnership/) ([HN](https://news.ycombinator.com/item?id=28617197))
- [LoRaMac-node](https://github.com/Lora-net/LoRaMac-node) - Reference implementation and documentation of a LoRa network node.
- [Picotracker Lora](https://github.com/ImperialSpaceSociety/LoRaMac-node)
- [Generic Node - The Things Industries](https://www.genericnode.com/) - One LoRaWAN node. Endless use cases. ([Code](https://github.com/TheThingsIndustries/generic-node-se))
- [Arduino-LMIC library](https://github.com/mcci-catena/arduino-lmic) - Adapted to run under the Arduino environment.
- [Airtime calculator for LoRaWAN](https://avbentem.github.io/airtime-calculator/ttn/eu868) ([Code](https://github.com/avbentem/airtime-calculator))
- [lorawan-device](https://github.com/Tortoaster/lorawan-device) - LoRaWAN device stack.
- [The Things Stack Application Cookbook](https://github.com/htdvisser/lorawan-stack-application-cookbook) - Guide for building an application or integration for The Things Stack.
- [LoRaWAN on Apache NuttX OS](https://lupyuen.github.io/articles/lorawan3)
- [Casually Chirping Into The World Of LoRaWAN (2022)](https://hackaday.com/2022/01/21/casually-chirping-into-the-world-of-lorawan/)
- [Ebook: Guide to LoRaWAN](https://www.nexpcb.com/lorawan-guide)
- [Private, Secure and Uncensorable Messaging Over a LoRa Mesh (2022)](https://unsigned.io/private-messaging-over-lora/) ([Reddit](https://www.reddit.com/r/darknetplan/comments/tq8v6d/private_secure_and_uncensorable_messaging_over_a/))
- [Reticulum](https://github.com/markqvist/Reticulum) - Self-configuring, encrypted and resilient mesh for LoRa, packet radio, WiFi and everything in between. ([HN](https://news.ycombinator.com/item?id=30870187))
- [RNode Firmware](https://github.com/markqvist/RNode_Firmware) - Firmware for the LoRa interface RNode.
- [Nomad Network](https://github.com/markqvist/NomadNet) - Off-grid, resilient mesh communication with strong encryption, forward secrecy and extreme privacy.
- [TTGO-T-Beam](https://github.com/LilyGO/TTGO-T-Beam)
- [N: Notkia – A Linux phone with LoRa+WiFi+BT connectivity](https://www.hackster.io/reimunotmoe/notkia-f6e772)
- [LoRa Mesh Network Node](https://github.com/brucemack/WARS-Birdhouse)
- [LoRa / LoRaWAN + TTN for MicroPython (ESP32)](https://github.com/fantasticdonkey/uLoRa)
- [Meshtastic](https://www.meshtastic.org/) - Encrypted communications platform for the Lora RF protocol. ([Code](https://github.com/meshtastic/Meshtastic)) ([Docs](https://meshtastic.org/docs/about)) ([HN](https://news.ycombinator.com/item?id=32016142))
