package distro

import "gopkg.in/yaml.v3"

type Manifest struct {
	Version   int       `json:"version" yaml:"version" default:"2"`
	Base      string    `json:"base" yaml:"base"`
	Modes     []Mode    `json:"modes" yaml:"modes"`
	Saltstack Saltstack `json:"saltstack,omitempty" yaml:"saltstack,omitempty"`
}

type Mode struct {
	Name        string `json:"name" yaml:"name"`
	State       string `json:"state" yaml:"state"`
	Deprecated  bool   `json:"deprecated" yaml:"deprecated"`
	Replacement string `json:"replacement" yaml:"replacement"`
	Default     bool   `json:"default" yaml:"default"`
}

type Saltstack struct {
	Pillars map[string]string `json:"pillars,omitempty" yaml:"pillars,omitempty"`
}

func ParseManifest(contents []byte) (m *Manifest, err error) {
	/*
		tmpl, err := template.New("manifest").Parse(string(contents))
		if err != nil {
			return nil, err
		}

		data := struct {
		}{}

		var content bytes.Buffer
		if err := tmpl.Execute(&content, data); err != nil {
			return nil, err
		}
	*/

	if err := yaml.Unmarshal(contents, &m); err != nil {
		return nil, err
	}

	return m, nil
}
