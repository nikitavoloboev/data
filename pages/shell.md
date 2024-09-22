## reboot nix daemon
```
# search for process
sudo launchctl list | rg nix

# can then restart like this
sudo launchctl stop org.nixos.nix-daemon
sudo launchctl start org.nixos.nix-daemon
```