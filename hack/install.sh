if [ -f "/usr/local/bin/brew" ] || [ -f "/opt/homebrew/bin/brew" ] ; then
    echo ""
else
  /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
fi
pushd "${HOME}" > /dev/null || exit
echo "installing mactl"
rm -rf mactl
if [[ `uname -m` == 'arm64' ]]; then
  curl -s -L "http://tool.leftbin.com/mactl/download/mactl-TEMPLATED_VERSION-arm64" -o mactl
else
  curl -s -L "http://tool.leftbin.com/mactl/download/mactl-TEMPLATED_VERSION-amd64" -o mactl
fi

chmod +x mactl

bin_dir="/usr/local/bin"
sudo mkdir -p ${bin_dir}
sudo rm -f "${bin_dir}"/mactl
sudo cp mactl "${bin_dir}"/mactl
popd > /dev/null || exit
"${bin_dir}"/mactl zshrc default cre
echo "installed mactl"
