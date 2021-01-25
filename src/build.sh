#!/bin/bash
go build .
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/dmitrii

oss=( linux darwin )
archs=( amd64 386 )
plugins_dir="${HOME}/.terraform.d/plugins/registry.terraform.io/dmitrii"

install_plugin() {
  plugin=$1
  version=0.0.1
  plugin_name=terraform-provider-$(basename "${plugin}")
  echo "Installing Terraform plugin ${plugin}..."
  for os in "${oss[@]}"
  do
    for arch in "${archs[@]}"
    do
      file="${plugin}_v${version}-${os}-${arch}"
      plugin_dst="${plugins_dir}/${plugin}/${version}/${os}_${arch}/${plugin_name}_v${version}"
      mkdir -p "$(dirname "${plugin_dst}")"
      echo "location: ${plugin_location}"
      cp "${plugin_name}" "${plugin_dst}"
      echo "Copied to ${plugin_dst}"
    done
  done
}

install_plugin "vkstatus"
