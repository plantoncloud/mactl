package docker

const (
	DaemonConfig = `
{
  "builder": {
    "gc": {
      "defaultKeepStorage": "20GB",
      "enabled": true
    }
  },
  "experimental": true,
  "features": {
    "buildkit": true
  }
}
`
	VmConfigTemplate = `
{
  "acceptCanaryUpdates": false,
  "activeOrganizationName": "",
  "analyticsEnabled": true,
  "autoDownloadUpdates": false,
  "autoStart": true,
  "backupData": false,
  "cpus": {{.Cpu}},
  "credentialHelper": "docker-credential-osxkeychain",
  "dataFolder": "{{.HomeDir}}/Library/Containers/com.docker.docker/Data/vms/0/data",
  "disableHardwareAcceleration": false,
  "disableTips": true,
  "disableUpdate": false,
  "diskFlush": "os",
  "diskQcowCompactAfter": 262144,
  "diskQcowKeepErased": 262144,
  "diskQcowRuntimeAsserts": false,
  "diskSizeMiB": {{.DiskSizeInMb}},
  "diskStats": "",
  "diskTRIM": true,
  "displayRestartDialog": true,
  "displaySwitchVersionPack": true,
  "displayedDeprecate1012": false,
  "displayedDeprecate1013": false,
  "displayedElectronPopup": [],
  "displayedTutorial": true,
  "dns": "8.8.8.8",
  "dockerAppLaunchPath": "/Applications/Docker.app",
  "filesharingDirectories": [
    "/Users",
    "/Volumes",
    "/private",
    "/tmp",
    "/var/folders"
  ],
  "kubernetesEnabled": false,
  "kubernetesInitialInstallPerformed": false,
  "lastLoginDate": 0,
  "latestBannerKey": "",
  "licenseTermsVersion": 2,
  "memoryMiB": {{.MemoryInMb}},
  "openUIOnStartupDisabled": false,
  "overrideProxyExclude": "",
  "overrideProxyHttp": "",
  "overrideProxyHttps": "",
  "proxyHttpMode": "system",
  "settingsVersion": 15,
  "showKubernetesSystemContainers": false,
  "socksProxyPort": 0,
  "swapMiB": 1024,
  "synchronizedDirectories": [],
  "tipLastId": 18,
  "tipLastViewedTime": 1637802627556,
  "updateAvailableTime": 0,
  "updatePopupAppearanceTime": 0,
  "updateSkippedBuild": "",
  "useCredentialHelper": true,
  "useDnsForwarder": true,
  "useNightlyBuildUpdates": false,
  "useVirtualizationFramework": false,
  "useVpnkit": true,
  "useWindowsContainers": false,
  "versionPack": "default",
  "vpnKitAllowedBindAddresses": "0.0.0.0",
  "vpnKitMTU": 1500,
  "vpnKitMaxConnections": 2000,
  "vpnKitMaxPortIdleTime": 300,
  "vpnKitTransparentProxy": true,
  "vpnkitCIDR": "192.168.65.0/24",
  "wslEngineEnabled": false
}
`
)
