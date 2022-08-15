# macOS

[macOS is incredible as a personal OS](https://github.com/nikitavoloboev/my-mac-os). Would love to join Apple and contribute to the OS at some point. [Karabiner](apps/karabiner/karabiner.md) is life changing.

## Clean install

You can clean install by going to Recovery mode (restart with `cmd+r` pressed). Then Disk Utility > Select disk > Erase (Format it) > Close Disk Utility > Select option Reinstall MacOS (Choose macOS ver. to install).

## Notes

- In save dialogues I can press these keys:
  - `Return` or `⌘ + S` = Save
  - `ESC` = Cancel
- I can also press `/` or `~` to quickly go to some directory from a save dialogue. And I can press `⌘ + ↑` to go to `parent directory`.
- Recovery mode: Power off the machine, press the power button and immediately hold Cmd-R.
- [Both Windows and MacOS are at a point where clean installs are unnecessary.](https://www.reddit.com/r/MacOS/comments/90g4h9/is_it_worth_the_effort_to_do_a_clean_install_of/)
  - I can appreciate someone wanting to do a clean install if they've installed and removed many apps and just want to clear out everything spread around all the system and hidden folders, even if it doesn't really affect performance and won't save a ton of disk space. There is something cathartic about a clean install.
- `/usr/local/bin` is a good place to put raw binaries available in the path, that are not installed with Nix.
- [`defaults write NSGlobalDomain KeyRepeat -int 1` setup keyboard repeat.](https://twitter.com/jordwalke/status/1230582824224165888)
- [Can select text from middle of link's text by holding down alt while you drag and select with the mouse](https://twitter.com/MBoffin/status/1218668903586394112)
- [plutil tool support the generation of Swift or Objective-C code directly from plists. For example: plutil -convert swift.](https://twitter.com/dmartincy/status/1295029196503351298)
- [To code sign binaries ad hoc, run `codesign -s - <path_to_binary>`.](https://github.com/golang/go/issues/42684) This will give users a gatekeeper warning but they could still run the binary. To sign so users can run binary without warning, you need Apple developer account.
- [I basically install nothing except GUI apps on my Mac. I don’t even have a bashrc (or any shell rc for that matter). I never open the terminal on Mac. I live in the VM for dev work, and use GUI apps like calendar, email, messages, browser, etc. on Mac.](https://twitter.com/mitchellh/status/1449151623092060164)
- [`ls -lha /Library/Developer/CommandLineTools/usr/bin` to see what xcode command line developer tools installs.](https://twitter.com/nikitonsky/status/1453019511502909441)
- [Can right-click on a 2-factor authentication code to set it up with iCloud Keychain.](https://twitter.com/rafahari/status/1456013646144933893)
- [macOS has network quality command: networkQuality](https://twitter.com/justsitandgrin/status/1460286405578563595) ([Reddit](https://www.reddit.com/r/MacOS/comments/qxa933/builtin_network_bandwidth_test_tool_on_macos/))
- [Just because a root certificate is in the built-in iOS/macOS trust store doesn't mean that it is trusted. Apple applies additional constraints via configuration updates to maintain a high-level of security.](https://twitter.com/BasileBailey/status/1494801237694300161)

## Code

### macOS Defaults

```bash
# Remove dock animation. https://www.reddit.com/r/apple/comments/6xg9xq/tip_of_the_day_one_thing_i_cant_live_without_in/
defaults write com.apple.dock autohide-delay -int 0
defaults write com.apple.dock autohide-time-modifier -float 0.4
killall Dock

# Revert
defaults delete com.apple.dock autohide-delay
defaults delete com.apple.dock autohide-time-modifier
killall Dock
```

```bash
# Turn internal keyboard off. https://discussions.apple.com/thread/5044946?answerId=26556362022#26556362022
sudo kextunload /System/Library/Extensions/AppleUSBTopCase.kext/Contents/PlugIns/AppleUSBTCKeyboard.kext/

# Turn internal keyboard on
sudo kextload /System/Library/Extensions/AppleUSBTopCase.kext/Contents/PlugIns/AppleUSBTCKeyboard.kext/
```

## Links

- [macOS developer tutorials](https://www.raywenderlich.com/category/macos)
- [A Pro’s Guide to the Best Secret Mac Features](https://matthewpalmer.net/blog/2018/04/14/ultimate-pro-guide-best-secret-mac-features/index.html)
- [macOS open source](https://opensource.apple.com/)
- [Create bootable USB macOS installer](https://macdaddy.io/create-bootable-usb-macos-installer/)
- [macOS Security and Privacy Guide](https://github.com/drduh/macOS-Security-and-Privacy-Guide)
- [Quick Look plugins](https://github.com/sindresorhus/quick-look-plugins)
- [cron is dead, long live launchd! (2017)](http://blog.jan-ahrens.eu/2017/01/13/cron-is-dead-long-live-launchd.html)
- [Control Mac Keyboard Brightness](https://github.com/pirate/mac-keyboard-brightness) - Programmatically flash the keyboard lights and control display brightness on Macs.
- [maclaunchmaclaunch](https://github.com/HazCod/maclaunch) - Manage your macOS startup items.
- [Objective-See](https://www.objective-see.com/) - Simple, yet effective macOS security tools.
- [Caliguvara - another approach to Touch Bar Presets (2019)](https://community.folivora.ai/t/caliguvara-c2-my-first-update-great-new-calendar-widgets-a-totally-new-way-to-share-your-music-design-fixes-and-more/6996)
- [AppUpdater](https://github.com/mxcl/AppUpdater) - Simple app-updater for macOS, checks your GitHub releases for a binary asset once a day and silently updates your app.
- [Brooklyn](https://github.com/pedrommcarrasco/Brooklyn) - Screensaver inspired by Apple's Event on October 30, 2018.
- [How macOS versions evolved and changed over the time](https://www.reddit.com/r/MacOS/comments/bc1lvk/how_macos_versions_evolved_and_changed_over_the/)
- [PureDarwin](https://github.com/PureDarwin/PureDarwin) - Community project to make Darwin more usable. ([HN](https://news.ycombinator.com/item?id=23799331)) ([Web](http://www.puredarwin.org/)) ([HN](https://news.ycombinator.com/item?id=31943198))
- [Curated list of shell commands and tools specific to macOS](https://github.com/herrbischoff/awesome-macos-command-line)
- [The new Mac Pro is a design remix (2019)](https://www.arun.is/blog/mac-pro/)
- [BlockBlock](https://www.objective-see.com/products/blockblock.html) - Continually monitors common persistence locations and displays an alert whenever a persistent component is added to the OS.
- [Impact](https://github.com/ChimeHQ/Impact) - Crash detection and recording library for Apple platforms.
- [Designing LookUp for macOS (2019)](https://medium.com/lookup-design/designing-lookup-for-macos-bf5b8fea1a01)
- [macOS screenshot tips](https://twitter.com/CoreyGinnivan/status/1187209574303973376)
- [gon](https://github.com/mitchellh/gon) - CLI and Go Library for macOS Notarization.
- [Shipping a Catalyst App - Peter Steinberger (2019)](https://www.youtube.com/watch?v=Xo3zGlyxXcI)
- [Open-source components of macOS](https://github.com/apple-open-source/macos)
- [macOS Headers](https://github.com/w0lfschild/macOS_headers) - Consistently maintained dump of most macOS Headers.
- [Every macOS white paper](https://github.com/0xmachos/mac-white-papers)
- [AppMover](https://github.com/OskarGroth/AppMover) - Framework for moving your application bundle to Applications folder on launch.
- [Lilu](https://github.com/acidanthera/Lilu) - Arbitrary kext and process patching on macOS.
- [osx-hid-inspector](https://github.com/pqrs-org/osx-hid-inspector) - Command line tool for macOS for inspecting human input devices (HID).
- [macOS Kernel Extensions are officially deprecated](https://developer.apple.com/support/kernel-extensions/) ([HN](https://news.ycombinator.com/item?id=22251076))
- [mas-cli](https://github.com/mas-cli/mas) - Simple command line interface for the Mac App Store. Designed for scripting and automation.
- [apply-user-defaults](https://github.com/zero-sh/apply-user-defaults) - Small utility to set macOS user defaults declaratively from a YAML file.
- [Zero.sh](https://github.com/zero-sh/zero.sh) - Radically simple personal bootstrapping tool for macOS.
- [DefaultApp](https://tyler.io/default-app-for-mac-ios/) - Template for starting macOS projects. ([Code](https://github.com/tylerhall/DefaultApp)) ([HN](https://news.ycombinator.com/item?id=22582456))
- [Awesome macOS](https://github.com/iCHAIT/awesome-macOS)
- [macOS and iOS Security Related Tools](https://github.com/ashishb/osx-and-ios-security-awesome)
- [BlackHole](https://github.com/ExistentialAudio/BlackHole) - Modern MacOS virtual audio driver that allows applications to pass audio to other applications with zero additional latency.
- [xcnotary](https://github.com/akeru-inc/xcnotary) - Missing macOS app notarization helper, built with Rust. ([HN](https://news.ycombinator.com/item?id=22743659))
- [skhd](https://github.com/koekeishiya/skhd) - Simple hotkey daemon for macOS.
- [Proxy Audio Driver](https://github.com/briankendall/proxy-audio-device) - Virtual audio driver for macOS to sends all audio to another output.
- [Icons.app](https://github.com/SAP/macOS-icon-generator) - App for macOS which is designed to generate consistent sized icons of an existing application in various states, jiggling (shaking) etc.
- [OpenCore](https://dortania.github.io/OpenCore-Desktop-Guide/) - Open-source, unconventional, first-in-class piece of software designed to intercept kernel loading to insert a highly advanced rootkit, designed to be an alternative to Clover. ([Code](https://github.com/dortania/OpenCore-Desktop-Guide)) ([HN](https://news.ycombinator.com/item?id=22916281))
- [Creating a macOS App with SwiftUI](https://developer.apple.com/tutorials/swiftui/creating-a-macos-app)
- [Syphon](https://github.com/Syphon/Syphon-Framework) - macOS technology to allow applications to share video and still images with one another in realtime, instantly.
- [Mac Bare Metal](https://flow.swiss/mac-bare-metal) - Enterprise-class IaaS for macOS.
- [React Native for macOS](https://github.com/microsoft/react-native-macos) - Build native macOS apps with React. ([HN](https://news.ycombinator.com/item?id=23160075))
- [My Mac App Store Debate (2020)](https://inessential.com/2020/05/12/my_mac_app_store_debate)
- [Kernel debugging macOS with SIP (2020)](https://www.offensive-security.com/offsec/kernel-debugging-macos-with-sip/)
- [Mac App Store Sandbox Escape (2020)](https://saagarjha.com/blog/2020/05/20/mac-app-store-sandbox-escape/)
- [macOS 10.15: Slow by Design (2020)](https://sigpipe.macromates.com/2020/macos-catalina-slow-by-design/) ([HN](https://news.ycombinator.com/item?id=23273247))
- [Catalina is checking notarization of unsigned executables (2020)](https://lapcatsoftware.com/articles/catalina-executables.html) ([HN](https://news.ycombinator.com/item?id=23281564))
- [Push-Button Installer of macOS Guests in VirtualBox for Windows, Linux, macOS](https://github.com/myspaghetti/macos-virtualbox) ([HN](https://news.ycombinator.com/item?id=23284987))
- [macOS in a Docker Container](https://github.com/sickcodes/Docker-OSX) ([HN](https://news.ycombinator.com/item?id=23419101)) ([Reddit](https://www.reddit.com/r/jailbreak/comments/gwg3e4/free_release_dockerosx_run_xcode_on_linux_sign/))
- [netboot.nix](https://github.com/grahamc/netboot.nix) - Create full netboot images in 15 seconds.
- [Swizzle](https://github.com/NSExceptional/Swizzle) - Extensible tweak to create simple tweaks for any app, from within any app.
- [What’s New in macOS Big Sur: Human Interface Guidelines (2020)](https://developer.apple.com/design/human-interface-guidelines/macos/overview/whats-new-in-macos/) ([HN](https://news.ycombinator.com/item?id=23606052))
- [About the Rosetta Translation Environment (2020)](https://developer.apple.com/documentation/apple_silicon/about_the_rosetta_translation_environment) ([HN](https://news.ycombinator.com/item?id=23613995))
- [The End of OS X (2020)](https://stratechery.com/2020/the-end-of-os-x/) ([HN](https://news.ycombinator.com/item?id=23617629))
- [Common Mac OS X Cursors as PNGs](https://tobiasahlin.com/blog/common-mac-os-x-lion-cursors/)
- [Apple’s Relentless Strategy, Execution, and Point of View (2020)](https://medium.learningbyshipping.com/apples-relentless-strategy-and-execution-7544a76aa26) ([HN](https://news.ycombinator.com/item?id=23670722))
- [Extract the system libraries on macOS Big Sur (2020)](https://lapcatsoftware.com/articles/bigsur.html)
- [Apple Technologies Search](https://developer.apple.com/documentation/technologies)
- [AquaticPrime](https://github.com/bdrister/AquaticPrime) - Mac software licensing code using cryptographically signed license files.
- [SimpleVM](https://github.com/KhaosT/SimpleVM) - Sample code for Virtualization framework. ([Fork](https://github.com/danczar/SimpleVM))
- [macOS 11 Big Sur UI Kit](https://products.ls.graphics/macos/)
- [The Art Of Mac Malware](https://taomm.org/)
- [Compromising the macOS Kernel through Safari by Chaining Six Vulnerabilities](https://github.com/sslab-gatech/pwn2own2020)
- [Mac keyboard shortcuts](https://support.apple.com/en-us/HT201236) ([HN](https://news.ycombinator.com/item?id=24080378))
- [Ask HN: What feature did you find after years of using macOS? (2020)](https://news.ycombinator.com/item?id=24091707)
- [Sinter](https://github.com/trailofbits/sinter) - User-mode application authorization system for MacOS written in Swift. ([Article](https://blog.trailofbits.com/2020/08/12/sinter-new-user-mode-security-enforcement-for-macos/))
- [MacPorts package manager](https://www.macports.org/) ([Code](https://github.com/macports/macports-ports))
- [macOS icon pack](https://macosicons.com/) - Beautiful open source icons for Big Sur. ([Code](https://github.com/elrumo/macOS_Big_Sur_icons_replacements))
- [Thoughts on macOS Package Managers (2019)](https://saagarjha.com/blog/2019/04/26/thoughts-on-macos-package-managers/)
- [Hard to discover tips and apps for macOS (2020)](https://thume.ca/2020/09/04/macos-tips/) ([HN](https://news.ycombinator.com/item?id=24391899)) ([Lobsters](https://lobste.rs/s/kpar4v/hard_discover_tips_apps_for_making_macos))
- [fastmac](https://github.com/fastai/fastmac/) - MacOS instance or Linux shell. ([HN](https://news.ycombinator.com/item?id=24452384))
- [What should I do after getting my (first) own new Macbook? (2020)](https://lobste.rs/s/uzdehw/what_should_i_do_after_getting_my_first_own)
- [macOS Simple KVM](https://github.com/foxlet/macOS-Simple-KVM) - Tools to set up a quick macOS VM in QEMU, accelerated by KVM.
- [HN: Apple’s T2 security chip jailbreak (2020)](https://news.ycombinator.com/item?id=24636166)
- [Apple's macOS SDKs](https://github.com/alexey-lysiuk/macos-sdk)
- [macOS11 Big Sur UI Kit](https://applypixels.com/resource/big-sur-ui)
- [Apple Platform Versions](https://github.com/phatblat/ApplePlatformVersions) - Recent history of platforms developed by Apple, including Apple-managed build tools for these platforms.
- [Sketch — Part of your world: Why we’re proud to build a truly native Mac app (2020)](https://www.sketch.com/blog/2020/10/26/part-of-your-world-why-we-re-proud-to-build-a-truly-native-mac-app/) ([HN](https://news.ycombinator.com/item?id=24899391)) ([HN 2](https://news.ycombinator.com/item?id=24912325)) ([Tweet](https://twitter.com/amix3k/status/1321404287566680064))
- [macOS 11 boot volume layout (2020)](https://eclecticlight.co/2020/09/16/boot-volume-layout/) ([HN](https://news.ycombinator.com/item?id=24957783))
- [Booting a macOS Apple Silicon Kernel in QEMU (2020)](https://worthdoingbadly.com/xnuqemu3/) ([HN](https://news.ycombinator.com/item?id=25064593))
- [macOS defaults](https://macos-defaults.com/) - List of macOS defaults commands with demos. ([Code](https://github.com/yannbertrand/macos-defaults))
- [Does Apple really log every app you run? A technical look (2020)](https://blog.jacopo.io/en/post/apple-ocsp/) ([HN](https://news.ycombinator.com/item?id=25095438)) ([Reddit](https://www.reddit.com/r/apple/comments/juajrb/does_apple_really_log_every_app_you_run_a/))
- [Does it ARM](https://doesitarm.com/) - Apps that are reported to support Apple Silicon. ([Code](https://github.com/ThatGuySam/doesitarm))
- [Black Friday Deals macOS/iOS dev](https://github.com/mRs-/Black-Friday-Deals)
- [create-dmg](https://github.com/create-dmg/create-dmg) - Shell script to build fancy DMGs.
- [Mach-O learning tool](https://github.com/nico/lssym) - Toy program to learn more about the mach-o file format.
- [Popover](https://github.com/iSapozhnik/Popover) - Custom macOS Popover.
- [Reverse Engineering on macOS](https://github.com/steven-michaud/reverse-engineering-on-osx)
- [Virtualization.framework tool (vftool)](https://github.com/evansm7/vftool) - Runs Linux virtual machines in macOS. ([HN](https://news.ycombinator.com/item?id=25382529))
- [macOS Setup Guide](https://sourabhbajaj.com/mac-setup/) ([Code](https://github.com/sb2nov/mac-setup))
- [Some Differences between macOS and Common Unix Systems (2020)](https://www.dyx.name/posts/macunix.html)
- [The Mac that saved Apple](https://sixcolors.com/post/2020/12/20-macs-for-2020-1-imac-g3/) ([HN](https://news.ycombinator.com/item?id=25566642))
- [macOS app in plain C](https://github.com/jimon/osx_app_in_plain_c)
- [Reverse-Engineering Apple Dictionary (2020)](https://fmentzer.github.io/posts/2020/dictionary/)
- [Project Mendacius](https://github.com/PraneetNeuro/Project-Mendacius) - GUI based virtualization tool for running Linux on macOS Big Sur.
- [dmgdist](https://github.com/insidegui/dmgdist) - Automate the process of creating, uploading and notarizing the DMG of a Mac app.
- [Apple Open Source Modules of Big Sur](https://opensource.apple.com/release/macos-1101.html) ([HN](https://news.ycombinator.com/item?id=25675763))
- [Apple ld64 Linker](https://opensource.apple.com/source/ld64/ld64-136/doc/design/linker.html)
- [jolk](https://github.com/kendfinger/jolk) - macOS System Executable Analyzer. ([Analyzing a new release of macOS with jolk](https://blog.endfinger.io/apple/macos/technical/jolk/2021/01/10/jolk.html))
- [MacHack](https://github.com/kendfinger/MacHack) - List of built-in tools in macOS that you probably didn't know about.
- [Shield](https://github.com/theevilbit/Shield) - App to protect against process injection on macOS. ([Article](https://theevilbit.github.io/shield/))
- [VMCLI](https://github.com/gyf304/vmcli) - Set of utilities to help you manage VMs with Virtualization.framework. ([HN](https://news.ycombinator.com/item?id=25786640))
- [m1n1: Experimentation playground for Apple Silicon](https://github.com/AsahiLinux/m1n1)
- [SplitConfigurations](https://github.com/KevinGutowski/SplitConfigurations) - Up the basics of a Big Sur app layout. Includes splitview and toolbarItems.
- [Preloader for Linux on M1](https://github.com/corellium/preloader-m1)
- [TinyLinux](https://github.com/niw/TinyLinux) - Tiny minimum implementation of Virtualization framework to boot Linux.
- [DyldExtractor](https://github.com/arandomdev/DyldExtractor) - Extract Binaries from Apple's Dyld Shared Cache.
- [ipsw](https://github.com/blacktop/ipsw) - iOS/macOS Research Swiss Army Knife. ([Web](https://blacktop.github.io/ipsw/))
- [Hacking native ARM64 binaries to run on the iOS Simulator (2021)](https://bogo.wtf/arm64-to-sim.html) ([Code](https://github.com/bogo/arm64-to-sim))
- [strongarm](https://github.com/datatheorem/strongarm) - Full-featured, cross-platform ARM64 Mach-O analysis library.
- [macbac](https://github.com/hazcod/macbac) - Lists, controls and schedules efficient APFS snapshots for your convenience.
- [UTM](https://mac.getutm.app/) - Virtual machines for Mac.
- [Use Touch ID for Sudo on Mac (2021)](https://davidwalsh.name/touch-sudo) ([HN](https://news.ycombinator.com/item?id=26302139))
- [Linux Desktop on Apple Silicon/M1 in Practice](https://gist.github.com/akihikodaki/87df4149e7ca87f18dc56807ec5a1bc5) ([HN](https://news.ycombinator.com/item?id=26324590))
- [Reverse-engineering Rosetta 2 (2021)](https://ffri.github.io/ProjectChampollion/part1/) ([Code](https://github.com/FFRI/ProjectChampollion))
- [Apple M1 Microarchitecture Research](https://dougallj.github.io/applecpu/firestorm.html) ([HN](https://news.ycombinator.com/item?id=26375139))
- [Mathias Bynens' Sensible macOS Defaults](https://github.com/mathiasbynens/dotfiles/blob/master/.macos) ([HN](https://news.ycombinator.com/item?id=26513528))
- [xhyve](https://github.com/machyve/xhyve) - Lightweight macOS virtualization solution.
- [Declarative macOS configuration options (2021)](https://lobste.rs/s/q8qgr7/any_advice_on_declarative_macos)
- [macOS scripts](https://github.com/0xmachos/macos-scripts) - Various scripts for macOS tasks.
- [Pure Rust Implementation of Apple Code Signing (2021)](https://gregoryszorc.com/blog/2021/04/14/pure-rust-implementation-of-apple-code-signing/)
- [Apple’s Notarizing](https://www.cdfinder.de/guide/blog/apple_hell.html) ([HN](https://news.ycombinator.com/item?id=26993857))
- [macOS dev setup](https://github.com/donnemartin/dev-setup)
- [Lima](https://github.com/AkihiroSuda/lima) - Linux-on-Mac ("macOS subsystem for Linux", "containerd for Mac"). ([HN](https://news.ycombinator.com/item?id=27151993))
- [Inspect Apple macOS software updates](https://github.com/hjuutilainen/sus-inspector)
- [MacVM](https://github.com/KhaosT/MacVM) - MacOS VM for Apple Silicon using Virtualization API.
- [My 2021 New Mac Setup](https://www.swyx.io/new-mac-setup-2021/)
- [macOS extensions are moving away from the kernel (2021)](https://eclecticlight.co/2021/07/07/extensions-are-moving-away-from-the-kernel/) ([HN](https://news.ycombinator.com/item?id=27758181))
- [iOS and macOS Performance Tuning Book (2017)](https://www.oreilly.com/library/view/ios-and-macostm/9780133085501/)
- [node-mac-permissions](https://github.com/codebytere/node-mac-permissions) - Native node module to manage system permissions on macOS.
- [Swift Programming for macOS](https://gavinw.me/swift-macos/) ([Code](https://github.com/wigging/swift-macos))
- [Organize and Index Your Screenshots (OCR) on macOS (2021)](https://alexn.org/blog/2020/11/11/organize-index-screenshots-ocr-macos.html)
- [The journey to controlling external monitors on M1 Macs (2021)](https://alinpanaitiu.com/blog/journey-to-ddc-on-m1-macs/) ([HN](https://news.ycombinator.com/item?id=28011861))
- [dyld-shared-cache-extractor](https://github.com/keith/dyld-shared-cache-extractor) - CLI for extracting libraries from Apple's dyld shared cache file.
- [ArchTest](https://github.com/below/ArchTest) - Why does this not compile on Apple Silicon?
- [The macOS Sandbox File Limit (2021)](https://buckleyisms.com/blog/anecdotes-about-the-macos-sandbox-file-limit/) ([HN](https://news.ycombinator.com/item?id=28105814))
- [Asahi Linux for Apple M1 progress report, August 2021](https://asahilinux.org/2021/08/progress-report-august-2021/) ([HN](https://news.ycombinator.com/item?id=28180135))
- [macOS 11's hidden security improvements (2021)](https://blog.malwarebytes.com/mac/2021/08/macos-11s-hidden-security-improvements/) ([HN](https://news.ycombinator.com/item?id=28250340))
- [Santa](https://github.com/google/santa) - Binary authorization system for macOS. ([Docs](https://santa.dev/))
- [Rudolph](https://github.com/airbnb/rudolph) - Control server counterpart of Santa, and is used to rapidly deploy configurations to Santa agents.
- [macOS persistence – Beyond the good ol' LaunchAgents](https://theevilbit.github.io/beyond/) ([HN](https://news.ycombinator.com/item?id=28498058))
- [macOS Security Compliance](https://github.com/usnistgov/macos_security) - Open source effort to provide a programmatic approach to generating security guidance.
- [Apple M1 Exploration](https://drive.google.com/file/d/1WrMYCZMnhsGP4o3H33ioAUKL_bjuJSPt/view)
- [WhatsNewViewController](https://github.com/Jonathan-Gander/WhatsNewViewController) - Nice way to present your new app features.
- [macOS Cross toolchain for Linux and \*BSD](https://github.com/tpoechtrager/osxcross)
- [Guide to moving to M1 mac (2021)](https://twitter.com/vvoyer/status/1451116237132533765)
- [macOS Orb](https://github.com/CircleCI-Public/macos-orb) - Convenient tools and settings for utilizing MacOS on CircleCI.
- [macOS Monterey: The MacStories Review (2021)](https://www.macstories.net/stories/macos-monterey-the-macstories-review/)
- [How macOS is more reliable, and doesn’t need reinstalling (2021)](https://eclecticlight.co/2021/10/29/how-macos-is-more-reliable-and-doesnt-need-reinstalling/)
- [Faster Mac Dev Tools with Custom Allocators (2021)](https://eisel.me/devtool-allocators) ([HN](https://news.ycombinator.com/item?id=29068828))
- [perfrecord](https://github.com/mstange/perfrecord) - macOS-only command line CPU profiler that displays the result in the Firefox profiler.
- [vz](https://github.com/Code-Hex/vz) - Create virtual machines and run Linux-based operating systems in Go using Apple Virtualization.framework.
- [diskspace](https://github.com/scriptingosx/diskspace) - macOS command line tool to return the available disk space on APFS volumes.
- [PlayCover](https://github.com/iVoider/PlayCover) - Run iOS apps & games on M1 Mac with mouse, keyboard and controller support.
- [Sideload iOS apps regardless of security settings](https://github.com/EricRabil/m1-ios-sideloader)
- [node-mac-userdefaults](https://github.com/codebytere/node-mac-userdefaults) - Native Node.js module that provides an interface to the user’s defaults database on macOS.
- [Resources about macOS/iOS system security](https://github.com/houjingyi233/macOS-iOS-system-security)
- [asitop](https://github.com/tlkh/asitop) - Performance monitoring CLI tool for Apple Silicon. ([Web](https://tlkh.github.io/asitop/))
- [OSX-KVM](https://github.com/kholia/OSX-KVM) - Run macOS on QEMU/KVM. ([HN](https://news.ycombinator.com/item?id=29426862))
- [macOS Optimizer](https://github.com/sickcodes/osx-optimizer) - Shell scripts to speed up your mac boot time, accelerate loading, and prevent unnecessary throttling.
- [How to sync multiple macs (work/personal) (2021)](https://twitter.com/tenderlove/status/1461461938592878594)
- [Mach-O Binaries (2015)](http://www.m4b.io/reverse/engineering/mach/binaries/2015/03/29/mach-binaries.html)
- [Explainer: .DS_Store Files (2021)](https://eclecticlight.co/2021/11/27/explainer-ds_store-files/) ([HN](https://news.ycombinator.com/item?id=29358932))
- [Interview with Hansen Hsu, engineer at Apple during transition from OS 9 to OS X (2021)](https://corecursive.com/cocoa-culture-with-hansen-hsu/) ([Lobsters](https://lobste.rs/s/0subb0/cocoa_culture)) ([HN](https://news.ycombinator.com/item?id=29424245))
- [optool](https://github.com/alexzielenski/optool) - Command Line Tool for interacting with MachO binaries on macOS/iOS.
- [mkuser](https://github.com/freegeek-pdx/mkuser) - Make user accounts for macOS with many advanced options.
- [Setup a New Developer Computer](https://github.com/vendasta/setup-new-computer-script) ([HN](https://news.ycombinator.com/item?id=29535432))
- [Entitlement AND Hardened Runtime Check](https://github.com/cedowens/EntitlementCheck) - Python3 script for macOS to check for binaries with problematic/interesting entitlements.
- [macFUSE](https://github.com/osxfuse/osxfuse) - Allows you to extend macOS via third party file systems.
- [Tuning Your Code’s Performance for Apple Silicon](https://developer.apple.com/documentation/apple-silicon/tuning-your-code-s-performance-for-apple-silicon) ([HN](https://news.ycombinator.com/item?id=29719544))
- [macOS Setup after 15 Years of Linux (2021)](https://hookrace.net/blog/macos-setup/) ([HN](https://news.ycombinator.com/item?id=29742551))
- [Machium](https://github.com/PsychoBird/Machium) - Debugger for Apple Silicon. ([Article](https://psychobird.github.io/Machium/Machium.html))
- [Apple Data Formats and Knowledge](https://github.com/hack-different/apple-knowledge) - Collection of reverse engineered Apple formats, protocols, or other interesting bits.
- [Running macOS in a Virtual Machine on Apple Silicon Macs](https://developer.apple.com/documentation/virtualization/running_macos_in_a_virtual_machine_on_apple_silicon_macs) ([Tweet](https://twitter.com/stroughtonsmith/status/1492535749488619520)) ([HN](https://news.ycombinator.com/item?id=30317260))
- [Apple Internals](https://github.com/mroi/apple-internals) - Information and tools to understand the internals of Apple’s operating systems.
- [Apple's custom NVMes are amazingly fast – if you don't care about data integrity (2022)](https://twitter.com/marcan42/status/1494213855387734019) ([HN](https://news.ycombinator.com/item?id=30370551)) ([Reddit](https://www.reddit.com/r/apple/comments/sun6pa/apples_custom_nvme_drives_are_amazingly_fast_if/)) ([Lobsters](https://lobste.rs/s/j5x2pm/apple_s_custom_nvme_drives_are_amazingly))
- [macOS Packer Templates for Cirrus CI](https://github.com/cirruslabs/osx-images)
- [NanoMDM](https://github.com/micromdm/nanomdm) - Minimalist Apple MDM server heavily inspired by MicroMDM.
- [eBPF for Docker Desktop on macOS](https://github.com/singe/ebpf-docker-for-mac)
- [duti](https://github.com/moretension/duti) - Command-line utility capable of setting default applications for various document types on macOS.
- [mpw-emu: Emulating 1998-Vintage Mac Compilers (2022)](https://wuffs.org/blog/emulating-mac-compilers)
- [xpcspy](https://github.com/hot3eed/xpcspy) - Bidirectional XPC message interception and more.
- [macOS Apple released code](https://github.com/apple-oss-distributions/distribution-macOS)
- [Apple T2 XPC](https://github.com/duo-labs/apple-t2-xpc) - Tools to explore the XPC interface of Apple's T2 chip.
- [aperture-node](https://github.com/wulkano/aperture-node) - Record the screen on macOS from Node.js.
- [macOS GateKeeper Helper](https://github.com/wynioux/macOS-GateKeeper-Helper) - Simple macOS GateKeeper script.
- [My macOS keyboard shortcuts (2022)](https://www.jamieonkeys.dev/posts/keyboard-shortcuts/) ([HN](https://news.ycombinator.com/item?id=30876934))
- [Asahi Linux installer](https://github.com/AsahiLinux/asahi-installer)
- [Infinite Mac](https://system7.app/)
- [Mac pfctl Port Forwarding](https://salferrarello.com/mac-pfctl-port-forwarding/)
- [MACVZ](https://github.com/mac-vz/macvz) - macOS Virtualization for Linux.
- [Network Quality Server](https://github.com/network-quality/server) - Share example servers that can be used by the networkQuality command line tool available in macOS 12.
- [SwiftInMemoryLoading](https://github.com/slyd0g/SwiftInMemoryLoading) - Swift implementation of in-memory Mach-O loading on macOS.
- [fetch-them-macos-headers](https://github.com/ziglang/fetch-them-macos-headers) - Utility for fetching minimal macOS libc headers.
- [SwiftAuthorizationSample](https://github.com/trilemma-dev/SwiftAuthorizationSample) - Swift sample app for running privileged operations on macOS using a helper tool.
- [Exploring Mach-O (2022)](https://gpanders.com/blog/exploring-mach-o-part-1/)
- [launch-tui](https://github.com/wecraftforfun/launch-tui) - Small TUI app to manage launchD.
- [apple-tools](https://github.com/meme/apple-tools) - Collection of tools for working with Apple software/hardware.
- [sudo-touchid](https://github.com/artginzburg/sudo-touchid) - Native and reliable TouchID support for sudo.
- [SendKeys](https://github.com/socsieng/sendkeys) - macOS command line application used to automate the keystrokes and mouse events.
- [How macOS manages M1 CPU cores (2022)](https://eclecticlight.co/2022/04/25/how-macos-manages-m1-cpu-cores/) ([HN](https://news.ycombinator.com/item?id=31151393))
- [macOS cross compiler toolchains](https://github.com/messense/homebrew-macos-cross-toolchains)
- [macchanger](https://github.com/acrogenesis/macchanger) - Easily change your mac address.
- [Moving a macOS window by clicking anywhere on it (2022)](https://mmazzarolo.com/blog/2022-04-16-drag-window-by-clicking-anywhere-on-macos/) ([HN](https://news.ycombinator.com/item?id=31273893))
- [Tart](https://github.com/cirruslabs/tart) - macOS VMs on Apple Silicon to use in CI and other automations.
- [Scripting OS X Blog](https://scriptingosx.com/) ([Twitter](https://twitter.com/scriptingosx))
- [macOS tips and tricks](https://saurabhs.org/macos-tips)
- [HvDecompile](https://github.com/zhuowei/HvDecompile) - Decompiling macOS Hypervisor.framework by hand.
- [Rediscovering the Mac: An iPad User's Journey into macOS with the M1 Max MacBook Pro (2022)](https://www.macstories.net/stories/rediscovering-the-mac/)
- [Create macOS or Linux virtual machines (2022)](https://developer.apple.com/videos/play/wwdc2022/10002/) ([HN](https://news.ycombinator.com/item?id=31645000))
- [Quick look at user-mode file systems on macOS Ventura (2022)](https://threedots.ovh/blog/2022/06/quick-look-at-user-mode-file-systems-on-macos-ventura/) ([Lobsters](https://lobste.rs/s/bm0ifk/quick_look_at_user_mode_file_systems_on))
- [Quick Tip: Enable Touch ID for sudo](https://sixcolors.com/post/2020/11/quick-tip-enable-touch-id-for-sudo/)
- [macOS-Login-Items](https://github.com/puffyCid/macos-loginitems) - Simple macOS LoginItems parser (and library) written in Rust.
- [macOS-fseventsd](https://github.com/puffyCid/macos-fseventsd) - Simple macOS File System Events Disk Log Stream (FsEventsd) parser (and library) written in Rust.
- [macOS-launchd](https://github.com/puffyCid/macos-launchd) - Simple macOS launchd parser (and library) written in Rust.
- [VirtualBuddy](https://github.com/insidegui/VirtualBuddy) - Virtualize macOS 12 and later on Apple Silicon. ([HN](https://news.ycombinator.com/item?id=31878876))
- [Hardening macOS (2022)](https://www.bejarano.io/hardening-macos/) ([HN](https://news.ycombinator.com/item?id=31864974))
- [macholib](https://github.com/ronaldoussoren/macholib) - Analyze and edit Mach-O headers, the executable format used by macOS.
- [rust-macho](https://github.com/flier/rust-macho) - Mach-O File Format Parser for Rust.
- [os-signpost](https://github.com/explosion/os-signpost) - Wrapper for the macOS signpost API.
- [macOS: App sandboxing via sandbox-exec (2020)](https://www.karltarvas.com/2020/10/25/macos-app-sandboxing-via-sandbox-exec.html) ([HN](https://news.ycombinator.com/item?id=31973232))
- [macschema](https://github.com/progrium/macschema) - Toolchain for generating JSON definitions of Apple APIs.
- [qemu-vmnet](https://github.com/adnsio/qemu-vmnet) - Native macOS networking for QEMU using vmnet.framework and socket networking.
- [Apple’s Virtualization framework is a great, free way to test new macOS betas (2022)](https://arstechnica.com/gadgets/2022/07/how-to-use-free-virtualization-apps-to-safely-test-the-macos-ventura-betas/)
- [Advice on keeping macOS clean (2022)](https://twitter.com/ShriramKMurthi/status/1553954621294489600)
- [Running Linux microVMs on macOS (M1/M2) (2022)](https://slp.prose.sh/running-microvms-on-m1)
- [libSandy](https://github.com/opa334/libSandy) - Securely extend the sandbox of system processes and user applications.
- [PlayCover](https://github.com/PlayCover/PlayCover) - Run iOS apps and games on Apple Silicon Macs with mouse, keyboard and controller support.
- [Completely Open-Source Implementation of Apple Code Signing and Notarization (2022)](https://gregoryszorc.com/blog/2022/08/08/achieving-a-completely-open-source-implementation-of-apple-code-signing-and-notarization/) ([HN](https://news.ycombinator.com/item?id=32386762))
- [ldid](https://github.com/ProcursusTeam/ldid) - Link Identity Editor. Put real or fake signatures in a Mach-O.
- [macOS Hardening](https://github.com/beerisgood/macOS_Hardening)
