#! /usr/bin/env bash

# variables
set -euo pipefail
tmp_dir=./tmp

## generate-c4: create C4 diagrams via Go and PUML
function task_generate_c4 {
    puml_version=1.2020.10
    rm -rf out
    mkdir out
    mkdir -p bin
    if [ ! -e bin/plantuml.${puml_version}.jar ]; then
        echo "downloading current PlantUML Version ${puml_version}"
        curl -L "https://netix.dl.sourceforge.net/project/plantuml/${puml_version}/plantuml.${puml_version}.jar" -o bin/plantuml.${puml_version}.jar
    fi
    for d in $(cd samples; ls *.go) ; do
        go run samples/$d >> out/$d.puml
    done
    java -Djava.awt.headless=true -jar bin/plantuml.${puml_version}.jar -charset utf-8 out/
}

function task_usage {
    echo "Usage: $0"
    sed -n 's/^##//p' <"$0" | column -t -s ':' |  sed -E $'s/^/\t/'
}

cmd=${1:-}
shift || true
resolved_command=$(echo "task_${cmd}" | sed 's/-/_/g')
if [[ "$(LC_ALL=C type -t "${resolved_command}")" == "function" ]]; then
    pushd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null
    ${resolved_command} "$@"
else
    task_usage
fi
