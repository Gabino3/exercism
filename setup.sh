#!/bin/bash
wget https://github.com/exercism/cli/releases/download/v2.4.1/exercism-linux-64bit.tgz
tar -xzf exercism-linux-64bit.tgz
rm exercism-linux-64bit.tgz
mkdir $HOME/bin
mv exercism $HOME/bin
export PATH=$HOME/bin:$PATH
echo 'export PATH=$HOME/bin:$PATH' >> $HOME/.bashrc
exercism configure --key=$(cat exercism_api_key)
exercism configure --dir=$(pwd)
mkdir -p $HOME/.config/exercism/
curl http://cli.exercism.io/exercism_completion.bash > $HOME/.config/exercism/exercism_completion.bash
echo 'if [ -f ~/.config/exercism/exercism_completion.bash ]; then
  . ~/.config/exercism/exercism_completion.bash
fi' >> $HOME/.bashrc