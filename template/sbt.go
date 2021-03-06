/*
 *  Copyright 2016 Cisco Systems, Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package template

import "github.com/cisco/elsy/helpers"

var sbtTemplateV1 = template{
	name: "sbt",
	composeYmlTmpl: `
{{if .ScratchVolumes}}
sbtscratch:
  image: busybox
  command: /bin/true
  volumes:
    {{.ScratchVolumes}}
{{end}}
sbt: &sbt
{{if .TemplateImage}}
  image: {{.TemplateImage}}
{{else}}
  image: paulcichonski/sbt
{{end}}
  volumes:
    - ./:/opt/project
  working_dir: /opt/project
  entrypoint: sbt
  volumes_from:
    - lc_shared_sbtdata
{{if .ScratchVolumes}}
    - sbtscratch
{{end}}
test:
  <<: *sbt
  entrypoint: [sbt, test]
package:
  <<: *sbt
  command: [assembly]
publish:
  <<: *sbt
  entrypoint: /bin/true
clean:
  <<: *sbt
  entrypoint: [sbt, clean]
`,
	scratchVolumes: `
    - /opt/project/target/resolution-cache
    - /opt/project/target/scala-2.11/classes
    - /opt/project/target/scala-2.11/test-classes
    - /opt/project/target/streams
    - /opt/project/project/project
    - /opt/project/project/target
`}

var sbtTemplateV2 = template{
	name: "sbt",
	composeYmlTmpl: `
version: '2'
services:
  {{if .ScratchVolumes}}
  sbtscratch:
    image: busybox
    command: /bin/true
    volumes:
      {{.ScratchVolumes}}
  {{end}}
  sbt: &sbt
{{if .TemplateImage}}
    image: {{.TemplateImage}}
{{else}}
    image: paulcichonski/sbt
{{end}}
    volumes:
      - ./:/opt/project
    working_dir: /opt/project
    entrypoint: sbt
    volumes_from:
      - container:lc_shared_sbtdata
  {{if .ScratchVolumes}}
      - sbtscratch
  {{end}}
  test:
    <<: *sbt
    entrypoint: [sbt, test]
  package:
    <<: *sbt
    command: [assembly]
  publish:
    <<: *sbt
    entrypoint: /bin/true
  clean:
    <<: *sbt
    entrypoint: [sbt, clean]
`,
	scratchVolumes: `
      - /opt/project/target/resolution-cache
      - /opt/project/target/scala-2.11/classes
      - /opt/project/target/scala-2.11/test-classes
      - /opt/project/target/streams
      - /opt/project/project/project
      - /opt/project/project/target
`}

func init() {
	addSharedExternalDataContainer("sbt", helpers.DockerDataContainer{
		Image:     "busybox:latest",
		Name:      "lc_shared_sbtdata",
		Volumes:   []string{"/root/.ivy2"},
		Resilient: true,
	})

	addV1(sbtTemplateV1)
	addV2(sbtTemplateV2)
}
