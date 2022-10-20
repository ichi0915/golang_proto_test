package migrate

import "strings"

func GetProfiles(KustomizeData Kustomize) string {

	return GetClouddriverProfile(KustomizeData) +
		GetDeckProfile(KustomizeData) +
		GetEchoProfile(KustomizeData) +
		GetFiatProfile(KustomizeData) +
		GetFront50Profile(KustomizeData) +
		GetGateProfile(KustomizeData) +
		GetIgorProfile(KustomizeData) +
		GetKayentaProfile(KustomizeData) +
		GetOrcaProfile(KustomizeData) +
		GetRoscoProfile(KustomizeData)
}

func GetClouddriverProfile(KustomizeData Kustomize) string {
	str := `
	clouddriver: {} # is the contents of ~/.hal/default/profiles/clouddriver.yml`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Clouddriver*/

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetDeckProfile(KustomizeData Kustomize) string {
	str := `
	# deck has a special key "settings-local.js" for the contents of settings-local.js
	deck:
	  # settings-local.js - contents of ~/.hal/default/profiles/settings-local.js
	  # Use the | YAML symbol to indicate a block-style multiline string
	  settings-local.js: |
	    window.spinnakerSettings.feature.kustomizeEnabled = true;
	    window.spinnakerSettings.feature.artifactsRewrite = true;`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetEchoProfile(KustomizeData Kustomize) string {
	str := `
	echo: {}    # is the contents of ~/.hal/default/profiles/echo.yml`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Echo*/

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetFiatProfile(KustomizeData Kustomize) string {
	str := `
	fiat: {}    # is the contents of ~/.hal/default/profiles/fiat.yml`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Fiat*/

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetFront50Profile(KustomizeData Kustomize) string {
	str := `
	front50: {} # is the contents of ~/.hal/default/profiles/front50.yml`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Front50*/

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetGateProfile(KustomizeData Kustomize) string {
	str := `
	gate: {}    # is the contents of ~/.hal/default/profiles/gate.yml`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Gate*/

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetIgorProfile(KustomizeData Kustomize) string {
	str := `
	igor: {}    # is the contents of ~/.hal/default/profiles/igor.yml`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Igor*/

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetKayentaProfile(KustomizeData Kustomize) string {
	str := `
	kayenta: {} # is the contents of ~/.hal/default/profiles/kayenta.yml`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Kayenta*/

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetOrcaProfile(KustomizeData Kustomize) string {
	str := `
	orca: {}    # is the contents of ~/.hal/default/profiles/orca.yml`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Orca*/

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetRoscoProfile(KustomizeData Kustomize) string {
	str := `
	rosco: {}   # is the contents of ~/.hal/default/profiles/rosco.yml`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Rosco*/

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}
